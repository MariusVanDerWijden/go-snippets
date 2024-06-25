package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	txfuzz "github.com/MariusVanDerWijden/tx-fuzz"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	address = "http://127.0.0.1:8545"
)

func main() {
	client, key := getRealBackend()
	backend := ethclient.NewClient(client)
	transactor, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	if err != nil {
		panic(err)
	}
	_, _, store, err := DeployWarmSStore(transactor, backend)
	if err != nil {
		panic(err)
	}
	lastGood := 100
	for i := 0; i < 1000; i++ {
		runs := lastGood + 1
		transactor.GasLimit = 30_000_000
		fmt.Printf("Using runs: %v\n", runs)
		tx, err := store.Sstore(transactor, big.NewInt(int64(runs)))
		if err != nil {
			fmt.Println(err)
			lastGood = lastGood - 2
			continue
		} else {
			lastGood = runs
		}
		rec, err := bind.WaitMined(context.Background(), backend, tx)
		if err != nil {
			panic(err)
		}
		if rec.Status == types.ReceiptStatusFailed {
			lastGood -= 10
		} else if rec.GasUsed < 29_000_000 {
			lastGood += 100
		}
	}
}

func getRealBackend() (*rpc.Client, *ecdsa.PrivateKey) {
	// eth.sendTransaction({from:personal.listAccounts[0], to:"0xb02A2EdA1b317FBd16760128836B0Ac59B560e9D", value: "100000000000000"})

	sk := crypto.ToECDSAUnsafe(common.FromHex(txfuzz.SK))
	if crypto.PubkeyToAddress(sk.PublicKey).Hex() != txfuzz.ADDR {
		panic(fmt.Sprintf("wrong address want %s got %s", crypto.PubkeyToAddress(sk.PublicKey).Hex(), txfuzz.ADDR))
	}
	cl, err := rpc.Dial(address)
	if err != nil {
		panic(err)
	}
	return cl, sk
}
