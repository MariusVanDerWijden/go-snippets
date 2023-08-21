package geth

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestStruct(t *testing.T) {
	// Generate a new random account and a funded simulator
	key, _ := crypto.GenerateKey()
	auth, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))

	sim := backends.NewSimulatedBackend(core.GenesisAlloc{auth.From: {Balance: big.NewInt(10000000000)}}, 10000000)
	defer sim.Close()

	// Deploy a structs method invoker contract and execute its default method
	_, _, structs, err := DeployBNum(auth, sim)
	if err != nil {
		t.Fatalf("Failed to deploy defaulter contract: %v", err)
	}
	sim.Commit()
	out, err := structs.Bpowi(nil, big.NewInt(9999), big.NewInt(1))
	if err != nil {
		t.Fatalf("Failed to deploy defaulter contract: %v", err)
	}
	panic(out)
}
