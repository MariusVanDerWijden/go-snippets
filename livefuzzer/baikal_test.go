package main

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

func TestBaikal(t *testing.T) {
	backend, key := getRealBackend()
	SendBaikalTransactions(backend, key, 100)
}

func TestUnstuck(t *testing.T) {

	backend, key := getRealBackend()
	unstuck(key, backend, common.Address{}, new(big.Int), new(big.Int).Mul(big.NewInt(110), big.NewInt(params.GWei)))
}
