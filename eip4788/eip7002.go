package eip4788

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
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
)

var deploymentCode7002 = "0x61049d5f556101af80600f5f395ff33373fffffffffffffffffffffffffffffffffffffffe1460a0573615156028575f545f5260205ff35b36603814156101ab5760115f54600182026001905f5b5f82111560595781019083028483029004916001019190603e565b9093900434106101ab57600154600101600155600354806003026004013381556001015f35815560010160203590553360601b5f5260385f601437604c5fa0600101600355005b6003546002548082038060101160b4575060105b5f5b8181146101585780604c02838201600302600401805490600101805490600101549160601b83528260140152807fffffffffffffffffffffffffffffffff0000000000000000000000000000000016826034015260401c906044018160381c81600701538160301c81600601538160281c81600501538160201c81600401538160181c81600301538160101c81600201538160081c81600101535360010160b6565b910180921461016a5790600255610175565b90505f6002555f6003555b5f548061049d141561018457505f5b6001546002828201116101995750505f61019f565b01600290035b5f555f600155604c025ff35b5f5ffd"
var code7002 = "0x3373fffffffffffffffffffffffffffffffffffffffe1460c7573615156028575f545f5260205ff35b36603814156101f05760115f54807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff146101f057600182026001905f5b5f821115608057810190830284830290049160010191906065565b9093900434106101f057600154600101600155600354806003026004013381556001015f35815560010160203590553360601b5f5260385f601437604c5fa0600101600355005b6003546002548082038060101160db575060105b5f5b81811461017f5780604c02838201600302600401805490600101805490600101549160601b83528260140152807fffffffffffffffffffffffffffffffff0000000000000000000000000000000016826034015260401c906044018160381c81600701538160301c81600601538160281c81600501538160201c81600401538160181c81600301538160101c81600201538160081c81600101535360010160dd565b9101809214610191579060025561019c565b90505f6002555f6003555b5f54807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff14156101c957505f5b6001546002828201116101de5750505f6101e4565b01600290035b5f555f600155604c025ff35b5f5ffd"

const (
	slot_excess  = 0
	slot_count   = 1
	queue_head   = 2
	queue_tail   = 3
	queue_offset = 4
	RECORD_SIZE  = 76
)

func get7002(statedb, statedb2 *state.StateDB) {
	addr := common.HexToAddress(precompileAddress)
	store := common.HexToAddress("0xfffffffffffffffffffffffffffffffffffffffe")
	out, logs, err := testCode7002(store, []byte{}, new(big.Int), statedb)
	if err != nil {
		panic(err)
	}
	_ = logs

	logger := logger.NewMarkdownLogger(nil, os.Stdout)
	config := vm.Config{Tracer: logger}
	out2, _, err := runtime.Call(addr, []byte{}, &runtime.Config{ChainConfig: params.AllDevChainProtocolChanges, Origin: store, State: statedb2, Debug: true, EVMConfig: config})
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(out, out2) {
		panic(fmt.Sprintf("out missmatch %v \n%v\n", out, out2))
	}
}

func set7002(input []byte, value *big.Int, statedb, statedb2 *state.StateDB) error {
	addr := common.HexToAddress(precompileAddress)
	caller := common.HexToAddress("0x1")

	out, _, err1 := testCode7002(caller, input, value, statedb)

	logger := logger.NewMarkdownLogger(nil, os.Stdout)
	config := vm.Config{Tracer: logger}
	out2, _, err2 := runtime.Call(addr, input, &runtime.Config{ChainConfig: params.AllDevChainProtocolChanges, Origin: caller, Value: value, Time: 0, State: statedb2, Debug: true, EVMConfig: config})
	if err1 != nil && err2 != nil {
		return errors.New("two errors")
	}
	if err1 != nil || err2 != nil {
		panic(fmt.Sprintf("%v%v", err1, err2))
	}
	if !bytes.Equal(out, out2) {
		panic(fmt.Sprintf("out missmatch %v \n%v\n", out, out2))
	}
	return nil
}

func CallGet7002() int {
	addr := common.HexToAddress(precompileAddress)
	statedb, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewDatabase(memorydb.New())), nil)
	statedb.SetCode(addr, common.FromHex(code7002))
	statedb2 := statedb.Copy()
	// Store
	get7002(statedb, statedb2)

	hash1, _ := statedb.Commit(0, false)
	hash2, _ := statedb2.Commit(0, false)
	if !bytes.Equal(hash1[:], hash2[:]) {
		panic(fmt.Sprintf("roots missmatch %v \n%v\n", hash1, hash2))
	}
	return 0
}

func CallSet7002(input []byte, value *big.Int) int {
	if len(input) > 57 {
		input = input[0:56]
	}
	addr := common.HexToAddress(precompileAddress)
	caller := common.HexToAddress("0x1")
	statedb, _ := state.New(types.EmptyRootHash, state.NewDatabase(rawdb.NewDatabase(memorydb.New())), nil)
	statedb.SetCode(addr, common.FromHex(code7002))
	statedb.CreateAccount(caller)
	statedb.SetBalance(caller, value)
	statedb2 := statedb.Copy()
	// Store
	if err := set7002(input, value, statedb, statedb2); err != nil {
		return 0
	}
	// Do the value transfer
	statedb.SetBalance(caller, big.NewInt(0))
	statedb.SetBalance(addr, value)
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

		panic(fmt.Sprintf("roots missmatch %v \n%v\n", hash1, hash2))
	}

	if len(input) == 56 {
		return 1
	}
	return 0
}

// https://github.com/lightclient/sys-asm/blob/main/src/withdrawals/main.eas
func testCode7002(caller common.Address, calldata []byte, value *big.Int, state *state.StateDB) ([]byte, []byte, error) {
	addr := common.HexToAddress(precompileAddress)

	if caller == common.Address(common.FromHex("0xfffffffffffffffffffffffffffffffffffffffe")) {
		tail_idx := new(big.Int).SetBytes(state.GetState(addr, common.BytesToHash(binary.BigEndian.AppendUint32([]byte{}, uint32(queue_tail)))).Bytes()) // sload(queue_tail)
		head_idx := new(big.Int).SetBytes(state.GetState(addr, common.BytesToHash(binary.BigEndian.AppendUint32([]byte{}, uint32(queue_head)))).Bytes()) // sload(queue_head)
		count := new(big.Int).Sub(tail_idx, head_idx).Uint64()
		if count > 16 {
			count = 16
		}
		var memory []byte
		for i := 0; i < int(count); i++ {
			offset := i * RECORD_SIZE
			addr_slot := int(head_idx.Uint64())*3 + offset + queue_offset
			address := state.GetState(addr, common.BytesToHash(binary.BigEndian.AppendUint32([]byte{}, uint32(addr_slot)))).Bytes() // sload(addr[i])
			pk := state.GetState(addr, common.BytesToHash(binary.BigEndian.AppendUint32([]byte{}, uint32(addr_slot+1)))).Bytes()    // sload(pk_0:32[i])
			pk_am := state.GetState(addr, common.BytesToHash(binary.BigEndian.AppendUint32([]byte{}, uint32(addr_slot+2)))).Bytes() // sload(pk2_am[i])
			memory = append(memory, address[12:32]...)
			memory = append(memory, pk...)
			memory = append(memory, pk_am[0:24]...)
			memory = append(memory, make([]byte, 20)...)
		}
		if head_idx.Uint64()+count == tail_idx.Uint64() {
			// reset queue
			state.SetState(addr, Uint64ToHash(uint64(queue_head)), common.Hash{}) // sstore(queue_head, 0)
			state.SetState(addr, Uint64ToHash(uint64(queue_tail)), common.Hash{}) // sstore(queue_tail, 0)
		} else {
			newHead := head_idx.Add(head_idx, big.NewInt(int64(count)))
			state.SetState(addr, Uint64ToHash(uint64(queue_tail)), common.BigToHash(newHead)) // sstore(queue_head, new_head)
		}
		// update new excess withdrawal requests
		excess := new(big.Int).SetBytes(state.GetState(addr, common.BytesToHash(binary.BigEndian.AppendUint32([]byte{}, uint32(slot_excess)))).Bytes()) // sload(queue_head)
		if excess.Uint64() == 1181 {
			excess.SetUint64(0)
		}
		countSlot := new(big.Int).SetBytes(state.GetState(addr, common.BytesToHash(binary.BigEndian.AppendUint32([]byte{}, uint32(slot_count)))).Bytes()) // sload(slot_count)
		if excess.Uint64()+countSlot.Uint64() > 2 {
			// compute excess
			countExcess := countSlot.Add(countSlot, excess)
			countExcess.Sub(countExcess, big.NewInt(2))
			excess.Set(countExcess)
		} else {
			excess.SetUint64(0)
		}
		state.SetState(addr, Uint64ToHash(uint64(slot_excess)), common.BigToHash(excess))   // sstore(slot_excess, excess)
		state.SetState(addr, Uint64ToHash(uint64(slot_count)), common.BigToHash(countSlot)) // sstore(slot_count, countSlot)
		return nil, memory, nil
	} else {
		if len(calldata) == 0 {
			excess_reqs := state.GetState(addr, common.BytesToHash(binary.BigEndian.AppendUint32([]byte{}, uint32(slot_excess)))) // sload(excess_reqs)
			return excess_reqs[:], nil, nil
		} else {
			if len(calldata) != 56 {
				return nil, nil, errors.New("invalid size")
			}
			excess_reqs := state.GetState(addr, common.BytesToHash(binary.BigEndian.AppendUint32([]byte{}, uint32(slot_excess)))) // sload(excess_reqs)
			req_fee := calcReqFee(big.NewInt(1), new(big.Int).SetBytes(excess_reqs.Bytes()), big.NewInt(17))
			if value.Cmp(req_fee) < 0 {
				return nil, nil, errors.New("to little fee")
			}
			// request can pay, increment withdrawal count
			req_count := state.GetState(addr, common.BytesToHash(binary.BigEndian.AppendUint32([]byte{}, uint32(slot_count)))) // sload(slot_count)
			newCount := new(big.Int).Add(new(big.Int).SetBytes(req_count[:]), common.Big1)
			state.SetState(addr, Uint64ToHash(uint64(slot_count)), common.BigToHash(newCount)) // sstore(slot_count, newCount)
			fmt.Printf("Setting %x to %x\n", Uint64ToHash(uint64(slot_count)), common.BigToHash(newCount))
			// insert req into queue
			tail_idx := new(big.Int).SetBytes(state.GetState(addr, common.BytesToHash(binary.BigEndian.AppendUint32([]byte{}, uint32(queue_tail)))).Bytes()) // sload(queue_tail)
			slot := new(big.Int).Add(new(big.Int).Mul(tail_idx, big.NewInt(3)), big.NewInt(queue_offset))                                                    // 3 * tail_idx + queue_offset
			state.SetState(addr, common.BigToHash(slot), common.BytesToHash(caller.Bytes()))                                                                 // sstore(slot, caller)
			fmt.Printf("Setting %x to %x\n", common.BigToHash(slot), common.BytesToHash(caller.Bytes()))
			slot = slot.Add(slot, common.Big1)                                               // slot += 1
			state.SetState(addr, common.BigToHash(slot), common.BytesToHash(calldata[0:32])) // sstore(slot, pk[0:32])
			fmt.Printf("Setting %x to %x\n", common.BigToHash(slot), common.BytesToHash(calldata[0:32]))
			slot = slot.Add(slot, common.Big1)                                                                                                                 // slot += 1
			state.SetState(addr, common.BigToHash(slot), common.BytesToHash(append(calldata[32:], []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}...))) // sstore(slot, pk[32:48]..amount)
			fmt.Printf("Setting %x to %x\n", common.BigToHash(slot), common.BytesToHash(calldata[32:]))
			// assemble log data
			var logData []byte
			logData = append(logData, caller.Bytes()...)
			logData = append(logData, calldata...)
			// store queue tail
			tail_idx.Add(tail_idx, big.NewInt(1))
			state.SetState(addr, common.BigToHash(big.NewInt(queue_tail)), common.BytesToHash(tail_idx.Bytes())) // sstore(slot, pk[32:48]..amount)
			fmt.Printf("Setting %x to %x\n", common.BigToHash(big.NewInt(queue_tail)), common.BytesToHash(tail_idx.Bytes()))
			return nil, logData, nil
		}
	}
}

/*
def fake_exponential(factor: int, numerator: int, denominator: int) -> int:
    i = 1
    output = 0
    numerator_accum = factor * denominator
    while numerator_accum > 0:
        output += numerator_accum
        numerator_accum = (numerator_accum * numerator) // (denominator * i)
        i += 1
    return output // denominator
*/

func calcReqFee(factor *big.Int, numerator *big.Int, denominator *big.Int) *big.Int {
	i := big.NewInt(1)
	output := big.NewInt(0)
	numerator_accum := factor.Mul(factor, denominator)
	for {
		if numerator_accum.Cmp(common.Big0) <= 0 {
			return output.Div(output, denominator)
		}
		output.Add(output, numerator_accum)
		numerator_accum = numerator_accum.Div(numerator_accum.Mul(numerator_accum, numerator), new(big.Int).Mul(denominator, i))
		i.Add(i, common.Big1)
	}
}
