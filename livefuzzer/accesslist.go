package main

import (
	"math/rand"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

// TODO can later be replaced by a call to ethclient.CreateAccessList
func createAccessList(client *rpc.Client, msg ethereum.CallMsg) (*types.AccessList, error) {
	type accessListResult struct {
		Accesslist *types.AccessList `json:"accessList"`
		Error      string            `json:"error,omitempty"`
		GasUsed    hexutil.Uint64    `json:"gasUsed"`
	}
	var result accessListResult
	if err := client.Call(&result, "eth_createAccessList", toCallArg(msg)); err != nil {
		return nil, err
	}
	return result.Accesslist, nil
}

func mutateAccessList(list types.AccessList) *types.AccessList {
	switch rand.Int31n(5) {
	case 0:
		// Leave the accesslist as is
		return &list
	case 1:
		// delete the access list
		return &types.AccessList{}
	case 2:
		// empty the access list
		return &types.AccessList{}
	case 3:
		// add a random entry and random slots to the list
		addr := randomAddress()
		keys := []common.Hash{}
		for i := 0; i < rand.Intn(10); i++ {
			h := randomHash()
			keys = append(keys, h)
		}
		tuple := types.AccessTuple{Address: addr, StorageKeys: keys}
		newList := types.AccessList(append([]types.AccessTuple{tuple}, list...))
		return &newList
	case 4:
		// replace a random entry and random slots of it in the list
		slot := list[rand.Int31n(int32(len(list)))]
		addr := randomAddress()
		keys := []common.Hash{}
		if len(slot.StorageKeys) == 0 {
			break
		}
		for i := 0; i < rand.Intn(len(slot.StorageKeys)); i++ {
			h := randomHash()
			keys = append(keys, h)
		}
		tuple := types.AccessTuple{Address: addr, StorageKeys: keys}
		newList := types.AccessList(append([]types.AccessTuple{tuple}, list...))
		return &newList
	case 5:
		// replace a random slot in an existing entry
		keyIdx := rand.Int31n(int32(len(list)))
		slotIdx := rand.Int31n(int32(len(list[keyIdx].StorageKeys)))
		h := randomHash()
		list[keyIdx].StorageKeys[slotIdx] = h
	case 6:
		var accesslist []types.AccessTuple
		for i := 0; i < rand.Int(); i++ {
			addr := randomAddress()
			keys := []common.Hash{}
			// create a fully random access list
			for q := 0; q < rand.Int(); q++ {
				h := randomHash()
				keys = append(keys, h)
			}
			tuple := types.AccessTuple{Address: addr, StorageKeys: keys}
			accesslist = append(accesslist, tuple)
		}
		newList := types.AccessList(accesslist)
		return &newList
	}
	return &list
}

func toCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["data"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
}
