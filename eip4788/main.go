package main

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/goevmlab/ops"
	"github.com/holiman/goevmlab/program"
)

//

func main() {
	orig := origByteCode()
	newb := newByteCode()
	if !bytes.Equal(orig, newb) {
		panic(common.Bytes2Hex(newb))
	}
	fmt.Println("success")
}

func origByteCode() []byte {
	return common.FromHex("3373fffffffffffffffffffffffffffffffffffffffe14604457602036146024575f5ffd5b620180005f350680545f35146037575f5ffd5b6201800001545f5260205ff35b6201800042064281555f359062018000015500")
}

func testCode(caller []byte, calldata []byte, timestamp int) {
	if bytes.Equal(caller, common.Hex2Bytes("0xfffffffffffffffffffffffffffffffffffffffe")) {
		t := timestamp % 0x018000
		sstore(t, timestamp)
		t2 := timestamp + 0x018000
		sstore(t2, calldata)
		return // stop
	}
	if len(calldata) != 0x20 {
		return // revert(0,0)
	}

}

func newByteCode() []byte {
	prog := program.NewProgram()

	prog.Op(ops.CALLER)
	prog.Push(common.FromHex("0xfffffffffffffffffffffffffffffffffffffffe"))
	prog.Op(ops.EQ)
	prog.Push(0x44)
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
	prog.Push(0x018000)
	prog.Push0()
	prog.Op(ops.CALLDATALOAD)
	prog.Op(ops.MOD)
	prog.Op(ops.DUP1)
	prog.Op(ops.SLOAD)
	prog.Push0()
	prog.Op(ops.CALLDATALOAD)
	prog.Op(ops.EQ)
	prog.Push(0x37)
	prog.Op(ops.JUMPI)

	prog.Push0()
	prog.Push0()
	prog.Op(ops.REVERT)

	prog.Op(ops.JUMPDEST)
	prog.Push(0x018000)
	prog.Op(ops.ADD)
	prog.Op(ops.SLOAD)
	prog.Push0()
	prog.Op(ops.MSTORE)
	prog.Push(0x20)
	prog.Push0()
	prog.Op(ops.RETURN)

	prog.Op(ops.JUMPDEST)
	prog.Push(0x018000)
	prog.Op(ops.TIMESTAMP)
	prog.Op(ops.MOD)
	prog.Op(ops.TIMESTAMP)
	prog.Op(ops.DUP2)
	prog.Op(ops.SSTORE)
	prog.Push0()
	prog.Op(ops.CALLDATALOAD)
	prog.Op(ops.SWAP1)
	prog.Push(0x018000)
	prog.Op(ops.ADD)
	prog.Op(ops.SSTORE)
	prog.Op(ops.STOP)

	return prog.Bytecode()
}
