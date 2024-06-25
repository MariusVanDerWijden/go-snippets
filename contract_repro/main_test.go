package geth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/MariusVanDerWijden/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/net/context"
)

func TestTellor(t *testing.T) {
	backend, sk := getRealBackend()
	transactor, err := bind.NewKeyedTransactorWithChainID(sk, big.NewInt(1337))
	if err != nil {
		t.Fatal(err)
	}
	_, tx, _, err := DeployERC20(transactor, backend)
	if err != nil {
		t.Fatal(err)
	}
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	addr, err := bind.WaitDeployed(ctx, backend, tx)
	if err != nil {
		t.Error(err)
	}
	_ = addr
}

func TestTellor(t *testing.T) {
	backend, sk := getSimBackend()
	transactor, err := bind.NewKeyedTransactorWithChainID(sk, big.NewInt(1337))
	if err != nil {
		t.Fatal(err)
	}
	supply := big.NewInt(1000)
	_, tx, _, err := DeploySmallContract(transactor, backend, supply)
	if err != nil {
		t.Fatal(err)
	}
	addr, err := bind.WaitDeployed(context.Background(), backend, tx)
	if err != nil {
		t.Error(err)
	}
	_ = addr
}

var (
	SK   = "0xcdfbe6f7602f67a97602e3e9fc24cde1cdffa88acd47745c0b84c5ff55891e1b"
	ADDR = "0xb02A2EdA1b317FBd16760128836B0Ac59B560e9D"
)

func getSimBackend() (*backends.SimulatedBackend, *ecdsa.PrivateKey) {
	sk, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	faucetAddr := crypto.PubkeyToAddress(sk.PublicKey)
	addr := map[common.Address]core.GenesisAccount{
		faucetAddr: {Balance: new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(9))},
	}
	alloc := core.GenesisAlloc(addr)
	return backends.NewSimulatedBackend(alloc, 80000000), sk
}

func getRealBackend() (*ethclient.Client, *ecdsa.PrivateKey) {
	// eth.sendTransaction({from:personal.listAccounts[0], to:"0xb02A2EdA1b317FBd16760128836B0Ac59B560e9D", value: "100000000000000"})

	sk := crypto.ToECDSAUnsafe(common.FromHex(SK))
	if crypto.PubkeyToAddress(sk.PublicKey).Hex() != ADDR {
		panic(fmt.Sprintf("wrong address want %s got %s", crypto.PubkeyToAddress(sk.PublicKey).Hex(), ADDR))
	}
	cl, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}
	return cl, sk
}
