package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-snippets/ethclient_test/coolcontract"
)

func main() {
	// Dial a node
	cl, err := ethclient.Dial("/tmp/geth.ipc")
	if err != nil {
		panic(err)
	}
	var transactOpts *bind.TransactOpts

	ctx := context.Background()
	{
		// Retrieve a block by number
		block, err := cl.BlockByNumber(ctx, big.NewInt(123))
		// Get Balance of an account (nil means at newest block)
		addr := common.HexToAddress("0xb02A2EdA1b317FBd16760128836B0Ac59B560e9D")
		balance, err := cl.BalanceAt(ctx, addr, nil)
		// Send transaction
		tx := new(types.Transaction)
		err = cl.SendTransaction(ctx, tx)
		// Get sync progress of the node
		progress, err := cl.SyncProgress(ctx)

		_ = block
		_ = balance
		_ = progress
		_ = err
	}

	// Account MGMT
	{

		// Use secret key hex string to sign a raw transaction
		SK := "0x0000"
		sk := crypto.ToECDSAUnsafe(common.FromHex(SK))
		nonce, err := cl.NonceAt(ctx, crypto.PubkeyToAddress(sk.PublicKey), nil)
		to := common.HexToAddress("0xb02A2EdA1b317FBd16760128836B0Ac59B560e9D")
		amount := big.NewInt(10 * params.GWei)
		gasLimit := uint64(21000)
		gasPrice := big.NewInt(10 * params.GWei)
		data := []byte{}
		tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(nil), sk)
		// If you have a bind.TransactOpts object you can also use the following
		transactOpts = bind.NewKeyedTransactor(sk)
		sigTx, err := transactOpts.Signer(types.NewEIP155Signer(nil), crypto.PubkeyToAddress(sk.PublicKey), tx)

		// Open Keystore
		ks := keystore.NewKeyStore("/home/matematik/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
		// Create an account
		acc, err := ks.NewAccount("password")
		// List all accounts
		accs := ks.Accounts()
		// Unlock an account
		ks.Unlock(accs[0], "password")
		// Create a TransactOpts object
		ksOpts, err := bind.NewKeyStoreTransactor(ks, accs[0])

		_ = err
		_ = signedTx
		_ = sigTx
		_ = ksOpts
		_ = acc
	}
	// Contract Bindings
	{
		backend, err := ethclient.Dial("/tmp/geth.ipc")
		// Deploy a new contract
		addr, tx, ctr, err := coolcontract.DeployCoolContract(transactOpts, backend)
		// Check if the contract was deployed successfully
		_, err = bind.WaitDeployed(ctx, backend, tx)
		// Call a pure/view function
		callOpts := &bind.CallOpts{Context: ctx, Pending: false}
		bal, err := ctr.SeeBalance(callOpts)

		// Filter for a Deposited event
		filterOpts := &bind.FilterOpts{Context: ctx, Start: 9000000, End: nil}
		itr, err := ctr.FilterDeposited(filterOpts)
		// Loop over all found events
		for itr.Next() {
			event := itr.Event
			// Print out all caller addresses
			fmt.Printf(event.Addr.Hex())
		}

		// Watch for a Deposited event
		watchOpts := &bind.WatchOpts{Context: ctx, Start: nil}
		// Setup a channel for results
		channel := make(chan *coolcontract.CoolContractDeposited)
		// Start a goroutine which watches new events
		go func() {
			sub, err := ctr.WatchDeposited(watchOpts, channel)
			defer sub.Unsubscribe()
		}()
		// Receive events from the channel
		event := <-channel

		log := *new(types.Log)
		event, err := ctr.ParseDeposited(log)

		_ = addr
		_ = ctr
		_ = err
		_ = bal
		_ = event
	}
}
