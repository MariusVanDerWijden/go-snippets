package main

import (
	"context"
	"math/big"

	"github.com/MariusVanDerWijden/go-snippets/ethclient_test/coolcontract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
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

		_ = addr
		_ = ctr
		_ = err
		_ = bal
	}
}
