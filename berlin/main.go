package main

import "github.com/ethereum/go-ethereum/ethclient"

func main() {
	url := "http://164.90.161.52:8545/"
	cl, err := ethclient.Dial(url)
}
