package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

func to1559Tx(tx *types.Transaction, chainID, tip, feeCap, gasPrice *big.Int) *types.Transaction {
	v, r, s := tx.RawSignatureValues()
	newTx := types.DynamicFeeTx{
		ChainID:    chainID,
		Nonce:      tx.Nonce(),
		Tip:        tip,
		FeeCap:     feeCap,
		Gas:        tx.Gas(),
		To:         tx.To(),
		Value:      tx.Value(),
		Data:       tx.Data(),
		AccessList: tx.AccessList(),
		V:          v,
		R:          r,
		S:          s,
	}
	return types.NewTx(&newTx)
}
