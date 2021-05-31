package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	SK       = "0xcdfbe6f7602f67a97602e3e9fc24cde1cdffa88acd47745c0b84c5ff55891e1b"
	ADDR     = "0xb02A2EdA1b317FBd16760128836B0Ac59B560e9D"
	Contract = "0x51eF031Ba1B2128971bd67A3ba5dda9BfF8706c8"
)

func main() {
	cl, sk := getRealBackend()
	fakeTx := types.NewTransaction(1, common.HexToAddress(ADDR), big.NewInt(0), 21000, big.NewInt(1), nil)
	noneip, err := types.SignTx(fakeTx, types.HomesteadSigner{}, sk)
	if err != nil {
		panic(err)
	}
	if err := cl.SendTransaction(context.Background(), noneip); err != nil {
		fmt.Println(err) // Prints "only replay-protected (EIP-155) transactions allowed over RPC"
	}
	eip155, err := types.SignTx(fakeTx, types.NewEIP155Signer(big.NewInt(1337)), sk)
	if err != nil {
		panic(err)
	}
	if err := cl.SendTransaction(context.Background(), eip155); err != nil {
		panic(err)
	}
}

func getRealBackend() (*ethclient.Client, *ecdsa.PrivateKey) {
	// eth.sendTransaction({from:personal.listAccounts[0], to:"0xb02A2EdA1b317FBd16760128836B0Ac59B560e9D", value: "100000000000000"})

	sk := crypto.ToECDSAUnsafe(common.FromHex(SK))
	if crypto.PubkeyToAddress(sk.PublicKey).Hex() != ADDR {
		panic(fmt.Sprintf("wrong address want %s got %s", crypto.PubkeyToAddress(sk.PublicKey).Hex(), ADDR))
	}
	cl, err := ethclient.Dial("/tmp/geth.ipc")
	if err != nil {
		panic(err)
	}
	return cl, sk
}

func makeCorpus() [][]byte {

	dir, err := ioutil.ReadDir("corpus")
	if err != nil {
		panic(err)
	}
	corpus := make([][]byte, len(dir))
	for i, info := range dir {
		in, err := ioutil.ReadFile("corpus/" + info.Name())
		if err != nil {
			panic(err)
		}
		corpus[i] = in
	}
	return corpus
}
