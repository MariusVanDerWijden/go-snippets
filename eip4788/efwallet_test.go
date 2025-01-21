package eip4788

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/goevmlab/ops"
	"github.com/holiman/goevmlab/program"
)

var walletCode = "0x608060405273ffffffffffffffffffffffffffffffffffffffff600054167fa619486e0000000000000000000000000000000000000000000000000000000060003514156050578060005260206000f35b3660008037600080366000845af43d6000803e60008114156070573d6000fd5b3d6000f3fea264697066735822122003d1488ee65e08fa41e58e888a9865554c535f2c77126a82cb4c0f917f31441364736f6c63430007060033"

func newProxyByteCode() []byte {
	prog := program.NewProgram()

	prog.Push(0x80)                                                                                 // 80
	prog.Push(0x40)                                                                                 // 80 40
	prog.Op(ops.MSTORE)                                                                             //
	prog.Push(common.HexToAddress("0xffffffffffffffffffffffffffffffffffffffff"))                    // 0xfff..
	prog.Push(0x00)                                                                                 // 0xfff... 0x00
	prog.Op(ops.SLOAD)                                                                              // 0xfff... slot[0x00]
	prog.Op(ops.AND)                                                                                // slot[0x00]
	prog.Push(common.Hex2Bytes("a619486e00000000000000000000000000000000000000000000000000000000")) // slot[0x00] a619..
	prog.Push(0x00)                                                                                 // slot[0x00] 0xa619.. 0x00
	prog.Op(ops.CALLDATALOAD)                                                                       // slot[0x00] 0xa619.. calldata[0x00]
	prog.Op(ops.EQ)                                                                                 // slot[0x00] eq?
	prog.Op(ops.ISZERO)                                                                             // slot[0x00] isZero
	prog.Push(0x50)                                                                                 // slot[0x00] isZero 0x50
	prog.Op(ops.JUMPI)                                                                              // slot[0x00]

	prog.Op(ops.DUP1)   // slot[0x00] slot[0x00]
	prog.Push(0x00)     // slot[0x00] slot[0x00] 0x00
	prog.Op(ops.MSTORE) // slot[0x00]
	prog.Push(0x20)     // slot[0x00] 0x20
	prog.Push(0x00)     // slot[0x00] 0x20 0x00
	prog.Op(ops.RETURN)

	prog.Op(ops.JUMPDEST)       // slot[0x00]
	prog.Op(ops.CALLDATASIZE)   // slot[0x00], calldatasize
	prog.Push(0x00)             // slot[0x00], calldatasize, 0x00
	prog.Op(ops.DUP1)           // slot[0x00], calldatasize, 0x00 0x00
	prog.Op(ops.CALLDATACOPY)   // slot[0x00]
	prog.Push(0x00)             // slot[0x00], 0x00
	prog.Op(ops.DUP1)           // slot[0x00], 0x00, 0x00
	prog.Op(ops.CALLDATASIZE)   // slot[0x00], 0x00, 0x00, calldatasize
	prog.Push(0x00)             // slot[0x00], 0x00, 0x00, calldatasize, 0x00
	prog.Op(ops.DUP5)           // slot[0x00], 0x00, 0x00, calldatasize, 0x00, slot[0x00]
	prog.Op(ops.GAS)            // slot[0x00], 0x00, 0x00, calldatasize, 0x00, slot[0x00], gas
	prog.Op(ops.DELEGATECALL)   // slot[0x00], success
	prog.Op(ops.RETURNDATASIZE) // slot[0x00], success, returndatasize
	prog.Push(0x00)             // slot[0x00], success, returndatasize, 0x00
	prog.Op(ops.DUP1)           // slot[0x00], success, returndatasize, 0x00, 0x00
	prog.Op(ops.RETURNDATACOPY) // slot[0x00], success, returndatasize, 0x00, 0x00
	prog.Push(0x00)             // slot[0x00], success, 0x00
	prog.Op(ops.DUP2)           // slot[0x00], success, 0x00, success
	prog.Op(ops.EQ)             // slot[0x00], success, eq
	prog.Op(ops.ISZERO)         // slot[0x00], success, iszero
	prog.Push(0x70)             // slot[0x00], success, iszero, 0x70
	prog.Op(ops.JUMPI)          // slot[0x00], success

	prog.Op(ops.RETURNDATASIZE) // slot[0x00], success, returndatasize
	prog.Push(0x00)             // slot[0x00], success, returndatasize, 0x00
	prog.Op(ops.REVERT)         // slot[0x00], success

	prog.Op(ops.JUMPDEST)       // slot[0x00], success
	prog.Op(ops.RETURNDATASIZE) // slot[0x00], success, returndatasize
	prog.Push(0x00)             // slot[0x00], success, returndatasize, 0x00
	prog.Op(ops.RETURN)         // slot[0x00], success

	// additional data

	prog.Op(0xfe)
	prog.Op(ops.LOG2)
	prog.Push(0x6970667358)
	prog.Op(0x22)
	prog.Op(ops.SLT)
	prog.Op(ops.KECCAK256)
	prog.Op(ops.SUB)
	prog.Op(0xd1)
	prog.Op(ops.BASEFEE)
	prog.Op(ops.DUP15)
	prog.Op(0xe6)
	prog.Op(0x5e) // mcopy
	prog.Op(ops.ADDMOD)
	prog.Op(ops.STATICCALL)
	prog.Op(ops.COINBASE)
	prog.Op(0xe5)
	prog.Op(ops.DUP15)
	prog.Op(ops.DUP9)
	prog.Op(ops.DUP11)
	prog.Op(ops.SWAP9)
	prog.Push(common.Hex2Bytes("554c535f2c77"))
	prog.Op(ops.SLT)
	prog.Push(common.Hex2Bytes("82cb4c0f917f3144136473"))
	prog.AddAll(common.Hex2Bytes("6f6c63430007060033"))

	return prog.Bytecode()
}

func TestProxyByteCode(t *testing.T) {
	bc := newProxyByteCode()
	if !bytes.Equal(bc, common.FromHex(walletCode)) {
		panic(common.Bytes2Hex(bc))
	}
}
