package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	var (
		protected   int
		unprotected int
	)
	// Dial a node
	cl, err := ethclient.Dial("/home/asdf/.ethereum/geth.ipc")
	if err != nil {
		panic(err)
	}
	head, err := cl.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}
	senders := make(map[common.Address]int)
	fmt.Printf("head: %v\n", head)
	for i := 0; i < 10000; i++ {
		block, err := cl.BlockByNumber(context.Background(), new(big.Int).SetUint64(head-uint64(i)))
		if err != nil {
			panic(err)
		}
		for _, tx := range block.Transactions() {
			if tx.Protected() {
				protected++
			} else {
				unprotected++
				msg, err := tx.AsMessage(types.HomesteadSigner{})
				if err != nil {
					fmt.Println(err)
				}
				senders[msg.From()]++
			}
		}
	}
	fmt.Printf("prot: %v , unprot: %v\n", protected, unprotected)
	fmt.Printf("Unique senders: %v\n", len(senders))
	fmt.Printf("%v\n", senders)
}
