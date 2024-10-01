package eip4788

import (
	"encoding/binary"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
)

var deploymentCode7002 = "0x61049d5f556101af80600f5f395ff33373fffffffffffffffffffffffffffffffffffffffe1460a0573615156028575f545f5260205ff35b36603814156101ab5760115f54600182026001905f5b5f82111560595781019083028483029004916001019190603e565b9093900434106101ab57600154600101600155600354806003026004013381556001015f35815560010160203590553360601b5f5260385f601437604c5fa0600101600355005b6003546002548082038060101160b4575060105b5f5b8181146101585780604c02838201600302600401805490600101805490600101549160601b83528260140152807fffffffffffffffffffffffffffffffff0000000000000000000000000000000016826034015260401c906044018160381c81600701538160301c81600601538160281c81600501538160201c81600401538160181c81600301538160101c81600201538160081c81600101535360010160b6565b910180921461016a5790600255610175565b90505f6002555f6003555b5f548061049d141561018457505f5b6001546002828201116101995750505f61019f565b01600290035b5f555f600155604c025ff35b5f5ffd"
var code7002 = "0x3373fffffffffffffffffffffffffffffffffffffffe1460a0573615156028575f545f5260205ff35b36603814156101ab5760115f54600182026001905f5b5f82111560595781019083028483029004916001019190603e565b9093900434106101ab57600154600101600155600354806003026004013381556001015f35815560010160203590553360601b5f5260385f601437604c5fa0600101600355005b6003546002548082038060101160b4575060105b5f5b8181146101585780604c02838201600302600401805490600101805490600101549160601b83528260140152807fffffffffffffffffffffffffffffffff0000000000000000000000000000000016826034015260401c906044018160381c81600701538160301c81600601538160281c81600501538160201c81600401538160181c81600301538160101c81600201538160081c81600101535360010160b6565b910180921461016a5790600255610175565b90505f6002555f6003555b5f548061049d141561018457505f5b6001546002828201116101995750505f61019f565b01600290035b5f555f600155604c025ff35b5f5ffd"

const (
	slot_excess  = 0
	slot_count   = 1
	queue_head   = 2
	queue_tail   = 3
	queue_offset = 4
)

// https://github.com/lightclient/sys-asm/blob/main/src/withdrawals/main.eas
func testCode7002(caller common.Address, calldata []byte, value *big.Int, state *state.StateDB) ([]byte, []byte, error) {
	addr := common.HexToAddress(precompileAddress)

	if caller == common.Address(common.FromHex("0xfffffffffffffffffffffffffffffffffffffffe")) {
		tail_idx := new(big.Int).SetBytes(state.GetState(addr, common.BytesToHash(binary.BigEndian.AppendUint32([]byte{}, uint32(queue_tail)))).Bytes()) // sload(queue_tail)
		head_idx := new(big.Int).SetBytes(state.GetState(addr, common.BytesToHash(binary.BigEndian.AppendUint32([]byte{}, uint32(queue_head)))).Bytes()) // sload(queue_head)
		count := new(big.Int).Sub(tail_idx, head_idx)
		if count.Cmp(big.NewInt(16)) > 0 {

		}
	} else {
		if len(calldata) == 0 {
			excess_reqs := state.GetState(addr, common.BytesToHash(binary.BigEndian.AppendUint32([]byte{}, uint32(slot_excess)))) // sload(excess_reqs)
			return excess_reqs[:], nil, nil
		} else {
			if len(calldata) != 56 {
				return nil, nil, errors.New("invalid size")
			}
			excess_reqs := state.GetState(addr, common.BytesToHash(binary.BigEndian.AppendUint32([]byte{}, uint32(slot_excess)))) // sload(excess_reqs)
			req_fee := calcReqFee(new(big.Int).SetBytes(excess_reqs.Bytes()), big.NewInt(17), big.NewInt(1))
			if value.Cmp(req_fee) < 0 {
				return nil, nil, errors.New("to little fee")
			}
			// request can pay, increment withdrawal count
			req_count := state.GetState(addr, common.BytesToHash(binary.BigEndian.AppendUint32([]byte{}, uint32(slot_count)))) // sload(slot_count)
			newCount := new(big.Int).Add(new(big.Int).SetBytes(req_count[:]), common.Big1)
			state.SetState(addr, Uint64ToHash(uint64(slot_count)), common.BigToHash(newCount)) // sstore(slot_count, newCount)
			// insert req into queue
			tail_idx := new(big.Int).SetBytes(state.GetState(addr, common.BytesToHash(binary.BigEndian.AppendUint32([]byte{}, uint32(queue_tail)))).Bytes()) // sload(queue_tail)
			slot := new(big.Int).Add(new(big.Int).Mul(tail_idx, big.NewInt(3)), big.NewInt(queue_offset))                                                    // 3 * tail_idx + queue_offset
			state.SetState(addr, common.BigToHash(slot), common.BytesToHash(caller.Bytes()))                                                                 // sstore(slot, caller)
			slot = slot.Add(slot, common.Big1)                                                                                                               // slot += 1
			state.SetState(addr, common.BigToHash(slot), common.BytesToHash(calldata[0:32]))                                                                 // sstore(slot, pk[0:32])
			slot = slot.Add(slot, common.Big1)                                                                                                               // slot += 1
			state.SetState(addr, common.BigToHash(slot), common.BytesToHash(calldata[32:]))                                                                  // sstore(slot, pk[32:48]..amount)
			// assemble log data
			var logData []byte
			logData = append(logData, caller.Bytes()...)
			logData = append(logData, calldata...)
			// store queue tail
			tail_idx.Add(tail_idx, big.NewInt(1))
			state.SetState(addr, common.BigToHash(big.NewInt(queue_tail)), common.Hash(tail_idx.Bytes())) // sstore(slot, pk[32:48]..amount)
			return nil, logData, nil
		}
	}
	return nil, nil, errors.New("should never happen")
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
	i := big.NewInt(0)
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
