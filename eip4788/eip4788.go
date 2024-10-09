package eip4788

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/core/vm/runtime"
	"github.com/ethereum/go-ethereum/eth/tracers/logger"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/goevmlab/ops"
	"github.com/holiman/goevmlab/program"
	"github.com/holiman/uint256"
)

func FuzzSet(origInput []byte, timestamp uint64) int {
	input := origInput
	if len(input) > 32 {
		input = input[:32]
	} else if len(input) < 32 {
		tmp := common.BytesToHash(input)
		input = tmp[:]
	}
	addr := common.HexToAddress(precompileAddress)
	store := common.HexToAddress("0xfffffffffffffffffffffffffffffffffffffffe")
	statedb, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewDatabase(memorydb.New())), nil)
	statedb.SetCode(addr, origByteCode())
	statedb2 := statedb.Copy()
	// Store
	out, err := testCode4788(store, input, timestamp, statedb)
	if err != nil {
		panic(err)
	}

	logger := logger.NewMarkdownLogger(nil, os.Stdout)
	config := vm.Config{Tracer: logger}
	out2, _, err := runtime.Call(addr, input, &runtime.Config{ChainConfig: params.AllDevChainProtocolChanges, Origin: store, Time: timestamp, State: statedb2, Debug: true, EVMConfig: config})
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(out, out2) {
		panic(fmt.Sprintf("out missmatch %v \n%v\n", out, out2))
	}
	hash1, _ := statedb.Commit(0, false)
	hash2, _ := statedb2.Commit(0, false)
	if !bytes.Equal(hash1[:], hash2[:]) {
		panic(fmt.Sprintf("roots missmatch %v \n%v\n", hash1, hash2))
	}
	if len(origInput) == 32 {
		return 1
	}
	return 0
}

func FuzzGet(input []byte) int {
	addr := common.HexToAddress(precompileAddress)
	caller := common.HexToAddress("0x1")
	statedb, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewDatabase(memorydb.New())), nil)
	statedb.SetCode(addr, origByteCode())
	cd := new(uint256.Int).SetBytes(input)
	t := new(uint256.Int).Mod(cd, uint256.NewInt(0x018000))
	// TODO only set one state
	statedb.SetState(addr, t.Bytes32(), common.BytesToHash(input))
	statedb.CreateAccount(caller)
	statedb2 := statedb.Copy()
	// Store
	out, err1 := testCode4788(caller, input, 0, statedb)

	logger := logger.NewMarkdownLogger(nil, os.Stdout)
	config := vm.Config{Tracer: logger}
	out2, _, err2 := runtime.Call(addr, input, &runtime.Config{ChainConfig: params.AllDevChainProtocolChanges, Origin: caller, Time: 0, State: statedb2, Debug: true, EVMConfig: config})
	if err1 != nil && err2 != nil {
		return 0
	}
	if err1 != nil || err2 != nil {
		panic(fmt.Sprintf("%v%v", err1, err2))
	}
	if !bytes.Equal(out, out2) {
		panic(fmt.Sprintf("out missmatch %v \n%v\n", out, out2))
	}

	// Should be equal
	hash1, _ := statedb.Commit(0, false)
	hash2, _ := statedb2.Commit(0, false)
	if !bytes.Equal(hash1[:], hash2[:]) {
		panic(fmt.Sprintf("roots missmatch %v \n%v\n", hash1, hash2))
	}

	if len(input) == 32 {
		return 1
	}
	return 0
}

func FuzzSetGet(input []byte) int {
	return 0
}

func origByteCode() []byte {
	return common.FromHex("3373fffffffffffffffffffffffffffffffffffffffe14604d57602036146024575f5ffd5b5f35801560495762001fff810690815414603c575f5ffd5b62001fff01545f5260205ff35b5f5ffd5b62001fff42064281555f359062001fff015500")
}

var precompileAddress = "0xfffffffffffffffffffffffffffffffffffffffe"

func testCode4788(caller common.Address, calldata []byte, timestamp uint64, state *state.StateDB) ([]byte, error) {
	addr := common.HexToAddress(precompileAddress)
	if caller == common.Address(common.FromHex("0xfffffffffffffffffffffffffffffffffffffffe")) {
		timeKey := Uint64ToHash(timestamp % 8191)
		state.SetState(addr, timeKey, Uint64ToHash(timestamp)) // sstore(timeKey, timestamp)
		rootKey := Uint64ToHash(timestamp%8191 + 8191)
		state.SetState(addr, rootKey, common.BytesToHash(calldata)) // sstore(rootKey, calldata)
		fmt.Printf("Writing %v to %v\n", timeKey, Uint64ToHash(timestamp))
		fmt.Printf("Writing %v to %v\n", rootKey, common.BytesToHash(calldata))
		return nil, nil // stop
	}
	if len(calldata) != 0x20 {
		return nil, errors.New("len(calldata) != 0x20") // revert(0,0)
	}
	cd := new(uint256.Int).SetBytes(calldata)
	t := new(uint256.Int).Mod(cd, uint256.NewInt(8191))
	loaded := state.GetState(addr, t.Bytes32()) // sload(t)
	if !bytes.Equal(loaded[:], calldata) || bytes.Equal(loaded[:], common.Hash{}.Bytes()) {
		fmt.Printf("loaded: %v \ncalldata: %v \nt.Bytes: %v\n", loaded, calldata, t.Bytes())
		return nil, errors.New("loaded != calldata") // revert(0,0)
	}
	t2 := new(uint256.Int).Add(t, uint256.NewInt(8191))
	loaded2 := state.GetState(addr, t2.Bytes32()) // sload(t2)
	fmt.Printf("Loaded %v from %v\n", loaded, t.Hex())
	fmt.Printf("Loaded %v from %v\n", loaded2, t2.Hex())
	return loaded2[:], nil // mstore(loaded2) return (20, 0)
}

func newByteCode() []byte {
	prog := program.NewProgram()

	prog.Op(ops.CALLER)
	prog.Push(common.FromHex("0xfffffffffffffffffffffffffffffffffffffffe"))
	prog.Op(ops.EQ)
	prog.Push(0x4d)
	prog.Op(ops.JUMPI)

	prog.Push(0x20)
	prog.Op(ops.CALLDATASIZE)
	prog.Op(ops.EQ)
	prog.Push(0x24)
	prog.Op(ops.JUMPI)

	prog.Push0()
	prog.Push0()
	prog.Op(ops.REVERT)

	prog.Op(ops.JUMPDEST)
	prog.Push0()
	prog.Op(ops.CALLDATALOAD)
	prog.Op(ops.DUP1)
	prog.Op(ops.ISZERO)
	prog.Push(0x49)
	prog.Op(ops.JUMPI)

	prog.Op(ops.PUSH3)
	prog.AddAll([]byte{0x00, 0x1f, 0xff})
	prog.Op(ops.DUP2)
	prog.Op(ops.MOD)
	prog.Op(ops.SWAP1)
	prog.Op(ops.DUP2)
	prog.Op(ops.SLOAD)
	prog.Op(ops.EQ)
	prog.Push(0x3c)
	prog.Op(ops.JUMPI)

	prog.Push0()
	prog.Push0()
	prog.Op(ops.REVERT)

	prog.Op(ops.JUMPDEST)
	prog.Op(ops.PUSH3)
	prog.AddAll([]byte{0x00, 0x1f, 0xff})
	prog.Op(ops.ADD)
	prog.Op(ops.SLOAD)
	prog.Push0()
	prog.Op(ops.MSTORE)
	prog.Push(0x20)
	prog.Push0()
	prog.Op(ops.RETURN)

	prog.Op(ops.JUMPDEST)
	prog.Push0()
	prog.Push0()
	prog.Op(ops.REVERT)

	prog.Op(ops.JUMPDEST)
	prog.Op(ops.PUSH3)
	prog.AddAll([]byte{0x00, 0x1f, 0xff})
	prog.Op(ops.TIMESTAMP)
	prog.Op(ops.MOD)
	prog.Op(ops.TIMESTAMP)
	prog.Op(ops.DUP2)
	prog.Op(ops.SSTORE)
	prog.Push0()
	prog.Op(ops.CALLDATALOAD)
	prog.Op(ops.SWAP1)
	prog.Op(ops.PUSH3)
	prog.AddAll([]byte{0x00, 0x1f, 0xff})
	//prog.Push(0x001fff)
	prog.Op(ops.ADD)
	prog.Op(ops.SSTORE)
	prog.Op(ops.STOP)

	return prog.Bytecode()
}

func Uint64ToHash(u uint64) common.Hash {
	var h common.Hash
	binary.BigEndian.PutUint64(h[24:], u)
	return h
}
