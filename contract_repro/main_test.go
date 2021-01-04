package geth

import (
	"crypto/ecdsa"
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/net/context"
)

func TestTellor(t *testing.T) {
	backend, sk := getSimBackend()
	transactor := bind.NewKeyedTransactor(sk)
	_, tx, tellor, err := DeployTellor(transactor, backend)
	if err != nil {
		t.Fail()
	}
	backend.Commit()
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	addr, err := bind.WaitDeployed(ctx, backend, tx)
	if err != nil {
		t.Error(err)
	}
	_ = addr
	_ = tellor
}

func TestGivenTellor(t *testing.T) {
	key, err := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	sim := backends.NewSimulatedBackend(core.GenesisAlloc{auth.From: {Balance: big.NewInt(math.MaxInt64)}}, math.MaxInt64)

	address, tx, _, err := DeployTellor(auth, sim)
	if err != nil {
		t.Error(err)
	}
	sim.Commit()
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	addr, err := bind.WaitDeployed(ctx, sim, tx)
	if err != nil {
		t.Error(err)
	}
	_ = addr
	_ = address
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
		common.BytesToAddress([]byte{1}): {Balance: big.NewInt(1)}, // ECRecover
		common.BytesToAddress([]byte{2}): {Balance: big.NewInt(1)}, // SHA256
		common.BytesToAddress([]byte{3}): {Balance: big.NewInt(1)}, // RIPEMD
		common.BytesToAddress([]byte{4}): {Balance: big.NewInt(1)}, // Identity
		common.BytesToAddress([]byte{5}): {Balance: big.NewInt(1)}, // ModExp
		common.BytesToAddress([]byte{6}): {Balance: big.NewInt(1)}, // ECAdd
		common.BytesToAddress([]byte{7}): {Balance: big.NewInt(1)}, // ECScalarMul
		common.BytesToAddress([]byte{8}): {Balance: big.NewInt(1)}, // ECPairing
		faucetAddr:                       {Balance: new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(9))},
	}
	alloc := core.GenesisAlloc(addr)
	return backends.NewSimulatedBackend(alloc, 80000000), sk
}
