package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

func getRealBackend() *ethclient.Client {
	cl, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}
	return cl
}
