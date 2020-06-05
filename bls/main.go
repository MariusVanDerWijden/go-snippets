package main

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	SK       = "0xcdfbe6f7602f67a97602e3e9fc24cde1cdffa88acd47745c0b84c5ff55891e1b"
	ADDR     = "0xb02A2EdA1b317FBd16760128836B0Ac59B560e9D"
	Contract = "0x51eF031Ba1B2128971bd67A3ba5dda9BfF8706c8"
)

func main() {

	getRealBackend()
}

func getRealBackend() (*ethclient.Client, *ecdsa.PrivateKey) {
	// eth.sendTransaction({from:personal.listAccounts[0], to:"0xb02A2EdA1b317FBd16760128836B0Ac59B560e9D", value: "100000000000000"})

	sk := crypto.ToECDSAUnsafe(common.FromHex(SK))
	if crypto.PubkeyToAddress(sk.PublicKey).Hex() != ADDR {
		panic(fmt.Sprintf("wrong address want %s got %s", crypto.PubkeyToAddress(sk.PublicKey).Hex(), ADDR))
	}
	cl, err := ethclient.Dial("/home/matematik/.yolo/geth.ipc")
	if err != nil {
		panic(err)
	}
	return cl, sk
}
