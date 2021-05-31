package main

import (
	"context"
	crand "crypto/rand"
	"math/big"

	"github.com/MariusVanDerWijden/FuzzyVM/filler"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

func SendBaikalTransactions(N int) {
	backend, key := getRealBackend()

	rnd := make([]byte, 10000)
	crand.Read(rnd)
	f := filler.NewFiller(rnd)

	sender := common.HexToAddress(ADDR)
	chainid, err := backend.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	nonce, err := backend.NonceAt(context.Background(), sender, nil)
	if err != nil {
		panic(err)
	}
	for i := 0; i < N; i++ {

		tx := randomTx(f, nonce)
		if f.Bool() {
			tx = to1559Tx(tx, chainid, f.BigInt(), f.BigInt(), f.BigInt())
		}
		signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainid), key)
		if err != nil {
			panic(err)
		}
		err = backend.SendTransaction(context.Background(), signedTx)
		if err == nil {
			nonce++
		}
	}
}

func randomTx(f *filler.Filler, nonce uint64) *types.Transaction {
	code := randomCode(f)
	if len(code) > 4096 {
		code = code[:4096]
	}
	if f.Bool() {
		return types.NewContractCreation(nonce, big.NewInt(10), 500000, new(big.Int).Mul(big.NewInt(100), big.NewInt(params.GWei)), code)
	}
	return types.NewTransaction(nonce, randomAddress(), big.NewInt(10), 500000, new(big.Int).Mul(big.NewInt(100), big.NewInt(params.GWei)), code)
}
