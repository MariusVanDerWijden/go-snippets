package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	ttdString := "58_750_000_000_000_000_000_000"
	//ttdString := "42150821634901480"

	blocksToCheck := 100000

	ttd, _ := new(big.Int).SetString(ttdString, 0)
	//https://mainnet.infura.io/v3/a1743f084f8a46bfb3696389eeb9f217
	client, err := rpc.Dial("/mnt/ethereum/geth.ipc")
	if err != nil {
		panic(err)
	}
	dataset := buildDataset(client, uint64(blocksToCheck))
	linearRegression(dataset, ttd)
}

type block struct {
	number int
	d      *big.Int
	td     *big.Int
	time   *big.Int
}

func linearRegression(dataset []block, ttd *big.Int) {
	latest := dataset[len(dataset)-1]
	leftTillTTD := new(big.Int).Sub(ttd, latest.td)
	avgDiff := new(big.Float)
	for _, b := range dataset {
		avgDiff.Add(avgDiff, big.NewFloat(float64(b.d.Uint64())))
	}
	avgDiff.Quo(avgDiff, big.NewFloat(float64(len(dataset))))

	avgTime := new(big.Float)
	for i := 0; i < len(dataset)-1; i++ {
		blockTime := new(big.Float).Sub(big.NewFloat(float64(dataset[i+1].time.Uint64())), big.NewFloat(float64(dataset[i].time.Uint64())))
		avgTime.Add(avgTime, blockTime)
	}
	avgTime.Quo(avgTime, big.NewFloat(float64(len(dataset))))

	blocksToMine := new(big.Float).Quo(new(big.Float).SetInt(leftTillTTD), avgDiff)
	timeNeeded := new(big.Float).Mul(blocksToMine, avgTime)

	currentTime := big.NewInt(time.Now().Unix())
	timeNeed, _ := timeNeeded.Int(new(big.Int))
	projectedTime := new(big.Int).Add(currentTime, timeNeed)
	projectedDate := time.Unix(projectedTime.Int64(), 0)

	fmt.Printf("Difficulty Left: %v\n", leftTillTTD.String())
	fmt.Printf("Average Difficulty: %v\n", avgDiff.String())
	fmt.Printf("Average Time per Block: %v\n", avgTime.String())
	fmt.Printf("Blocks to mine until TTD: %v\n", blocksToMine.String())
	fmt.Printf("Time needed to mine until TTD: %v\n", timeNeeded.String())
	fmt.Printf("Projected Date of TTD hit: %v\n", projectedDate)
}

func buildDataset(client *rpc.Client, n uint64) []block {
	cl := ethclient.NewClient(client)
	end, err := cl.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}
	start := end - n

	fmt.Printf("Building dataset from %v to %v, %v blocks\n", start, end, n)
	var bls []block
	for i := start; i < end; i++ {
		b, err := HeaderByNumber(client, big.NewInt(int64(i)))
		if err != nil {
			panic(err)
		}
		bls = append(bls, block{
			number: int(i),
			d:      b.Difficulty.ToInt(),
			td:     b.TotalDifficulty.ToInt(),
			time:   big.NewInt(int64(b.Time)),
		})
	}
	return bls
}

func HeaderByNumber(client *rpc.Client, number *big.Int) (Header, error) {
	var head Header
	err := client.CallContext(context.Background(), &head, "eth_getBlockByNumber", hexutil.EncodeBig(number), false)
	return head, err
}

type Header struct {
	Difficulty *hexutil.Big `json:"difficulty"       gencodec:"required"`
	Number     *hexutil.Big `json:"number"           gencodec:"required"`

	// BaseFee was added by EIP-1559 and is ignored in legacy headers.
	BaseFee         *hexutil.Big `json:"baseFeePerGas" rlp:"optional"`
	TotalDifficulty *hexutil.Big `json:"totalDifficulty" rlp:"optional"`

	Time hexutil.Uint64 `json:"timestamp"        gencodec:"required"`
}
