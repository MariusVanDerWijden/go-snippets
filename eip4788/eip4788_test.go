package eip4788

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/core/vm/runtime"
	"github.com/ethereum/go-ethereum/eth/tracers/logger"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/params"
)

func TestBytecodeMatch(t *testing.T) {
	orig := origByteCode()
	newb := newByteCode()
	if !bytes.Equal(orig, newb) {
		panic(common.Bytes2Hex(newb))
	}
	fmt.Println("success")
}

func TestGoCode(t *testing.T) {
	caller := common.HexToAddress("0x0")
	store := common.HexToAddress("0xfffffffffffffffffffffffffffffffffffffffe")
	calldata := Uint64ToHash(0x12345678)
	timestamp := uint64(0x1234)
	statedb, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewDatabase(memorydb.New())), nil)
	// Store
	_, err := testCode4788(store, calldata[:], timestamp, statedb)
	if err != nil {
		t.Fatal(err)
	}
	// Retrieve
	retr := Uint64ToHash(timestamp)
	out2, err2 := testCode4788(caller, retr[:], 0, statedb)
	if err2 != nil {
		t.Fatal(err2)
	}
	if !bytes.Equal(out2, calldata[:]) {
		panic(out2)
	}
}

func TestBytecode(t *testing.T) {
	input := Uint64ToHash(0x12345678)
	timestamp := uint64(0x1234)
	// Store
	statedb, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewDatabase(memorydb.New())), nil)
	statedb.SetCode(common.BytesToAddress([]byte("contract")), origByteCode())

	logger := logger.NewMarkdownLogger(nil, os.Stdout)
	store := common.HexToAddress("0xfffffffffffffffffffffffffffffffffffffffe")
	out, _, err := runtime.Call(common.BytesToAddress([]byte("contract")), input[:], &runtime.Config{ChainConfig: params.AllDevChainProtocolChanges, Origin: store, Time: timestamp, State: statedb, Debug: true, EVMConfig: vm.Config{Tracer: logger}})
	if err != nil {
		t.Fatal(err)
	}
	_ = out
	// Retrieve
	caller := common.HexToAddress("0x0")
	retr := Uint64ToHash(timestamp)
	out2, _, err2 := runtime.Call(common.BytesToAddress([]byte("contract")), retr[:], &runtime.Config{ChainConfig: params.AllDevChainProtocolChanges, Origin: caller, Time: timestamp, State: statedb, Debug: true, EVMConfig: vm.Config{Tracer: logger}})
	if err2 != nil {
		fmt.Println(out2)
		t.Fatal(err2)
	}
	if !bytes.Equal(out2, input[:]) {
		panic(out2)
	}
}

func TestFuzzSet(t *testing.T) {
	calldata := Uint64ToHash(0x12345678)
	FuzzSet([]byte{0x01}, 0x12345678)
	FuzzSet(calldata[:], 0x12345678)
	FuzzSet(append(calldata[:], 0x01), 0x12345678)
	FuzzSet(append(calldata[:], calldata[:]...), 0x12345678)
	FuzzSet(append(calldata[:], []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}...), 0x12345678)

	FuzzSet([]byte(string("00000000000000000000")), 0x12345678)
}

func TestFuzzGet(t *testing.T) {
	calldata := Uint64ToHash(0x12345678)
	//FuzzGet([]byte{0x01})
	FuzzGet(calldata[:])
	FuzzGet(append(calldata[:], 0x01))
	FuzzGet(append(calldata[:], calldata[:]...))
	FuzzGet(append(calldata[:], []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}...))

	FuzzGet([]byte(string("00000000000000000000")))
}

func TestGetCrash(t *testing.T) {
	input := []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
	FuzzGet(input)
}

func FuzzFuzzSet(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte, timestamp uint64) {
		FuzzSet(input, timestamp)
	})
}

func FuzzFuzzGet(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		FuzzGet(input)
	})
}
