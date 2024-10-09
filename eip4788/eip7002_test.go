package eip4788

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
)

func FuzzEIP7002(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte, valueB uint64) {
		value := new(big.Int).SetUint64(valueB)
		CallSet7002(input, value)
	})
}

func TestEIP7002(t *testing.T) {
	var (
		input = []byte("000100A2200000708010201100101170\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")
		val   = 20
	)
	value := new(big.Int).SetUint64(uint64(val))
	CallSet7002(input, value)
}

func FuzzEIP7002GetSet(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		var (
			valueSum = new(big.Int)
			inputs   [][]byte
			valueB   []uint64
		)
		for i := 0; i < len(input)%(56+8); i++ {
			if i*(56+8)+56+8 > len(input) {
				break
			}
			inputs = append(inputs, input[i*(56+8):i*(56+8)+56])
			valueB = append(valueB, binary.BigEndian.Uint64(input[i*(56+8)+56:i*(56+8)+56+8]))
		}

		for _, val := range valueB {
			valueSum.Add(valueSum, new(big.Int).SetUint64(val))
		}
		addr := common.HexToAddress(precompileAddress)
		caller := common.HexToAddress("0x1")
		statedb, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewDatabase(memorydb.New())), nil)
		statedb.SetCode(addr, common.FromHex(code7002))
		statedb.CreateAccount(caller)
		statedb.SetBalance(caller, valueSum)
		statedb2 := statedb.Copy()
		// Store
		for i := 0; i < min(len(inputs), len(valueB)); i++ {
			value := new(big.Int).SetUint64(valueB[i])
			if err := set7002(input, value, statedb, statedb2); err != nil {
				continue
			}
		}
		// Do the value transfer
		statedb.SetBalance(caller, big.NewInt(0))
		statedb.SetBalance(addr, valueSum)
		statedb2.SetBalance(addr, valueSum)
		statedb2.SetBalance(caller, big.NewInt(0)) // need to do this, to touch the caller, otherwise root mismatch
		// Should be equal
		hash1, _ := statedb.Commit(0, true)
		hash2, _ := statedb2.Commit(0, true)

		if !bytes.Equal(hash1[:], hash2[:]) {
			fmt.Printf("%x\n", statedb.GetBalance(caller))
			fmt.Printf("%x\n", statedb2.GetBalance(caller))
			fmt.Printf("%x\n", statedb.GetBalance(addr))
			fmt.Printf("%x\n", statedb2.GetBalance(addr))
			fmt.Printf("%#v\n", statedb.GetOrNewStateObject(caller))
			fmt.Printf("%#v\n", statedb2.GetOrNewStateObject(caller))
			fmt.Printf("%s\n", statedb.Dump(nil))
			fmt.Printf("%s\n", statedb2.Dump(nil))

			panic(fmt.Sprintf("roots mismatch %v \n%v\n", hash1, hash2))
		}
		// now we do the get
		get7002(statedb, statedb2)
		hash1, _ = statedb.Commit(0, false)
		hash2, _ = statedb2.Commit(0, false)
		if !bytes.Equal(hash1[:], hash2[:]) {
			panic(fmt.Sprintf("roots mismatch after get %v \n%v\n", hash1, hash2))
		}
	})
}
