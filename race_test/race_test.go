package race

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
)

func TestStateDBNotSet(t *testing.T) {
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	addr := common.HexToAddress("0x00")
	nonce := statedb.GetNonce(addr)
	go statedb.SetNonce(addr, 12)
	if nonce == statedb.GetNonce(addr) {
		t.Fatal("Race")
	}
}

func TestStateDBRaceNonce(t *testing.T) {
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	addr := common.HexToAddress("0x00")
	statedb.SetNonce(addr, 10)
	nonce := statedb.GetNonce(addr)
	go statedb.SetNonce(addr, 12)
	if nonce == statedb.GetNonce(addr) {
		t.Fatal("Race")
	}
}
