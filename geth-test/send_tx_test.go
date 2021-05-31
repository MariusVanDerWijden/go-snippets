package geth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func fundSender(sender common.Address, sk *ecdsa.PrivateKey, client *backends.SimulatedBackend) {
	txfund := types.NewTransaction(0, sender, big.NewInt(11000000000000000), 21000, big.NewInt(1), nil)

	chainID := big.NewInt(1337)
	privateKey := sk

	signedTx, err := types.SignTx(txfund, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		panic(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		panic(err)
	}
}

func TestTransfer(t *testing.T) {

	client, faucetSK := getSimBackend()
	// Create a new sender address
	sk, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	fromAddress := crypto.PubkeyToAddress(sk.PublicKey)
	// Prefund with 0.011 ETH
	fundSender(fromAddress, faucetSK, client)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000)
	gasLimit := uint64(21000)
	// Set gas price to 18 gwei
	gasPrice := big.NewInt(184000014590)
	fmt.Println(int64(gasLimit)*gasPrice.Int64() + value.Int64())
	toAddress := fromAddress
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// SimBackend has chainid 1337
	chainID := big.NewInt(1337)
	privateKey := sk

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)

	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(err)
		return
	}
	// We have to advance the simulated blockchain
	client.Commit()
	// Wait for the chain result
	rec, err := bind.WaitMined(context.Background(), client, signedTx)
	if err != nil {
		panic(err)
	}
	if rec.Status != types.ReceiptStatusSuccessful {
		panic("tx failed")
	}
	panic("success")
}
