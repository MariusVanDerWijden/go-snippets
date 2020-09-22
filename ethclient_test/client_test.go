package main

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

func TestDebugTraceCall(t *testing.T) {
	client, _ := rpc.Dial("/tmp/geth.ipc")
	res, err := TraceCallRPC(client)
	if err != nil {
		t.Fail()
	}
	panic(fmt.Sprintf("Success: %v", res))
}

type traceResult struct {
	gas         int         `json:"gas"`
	failed      bool        `json:failed`
	returnValue string      `json:returnValue`
	structLogs  []types.Log `json:structLogs`
}

func TraceCallRPC(client *rpc.Client) (traceResult, error) {
	var result traceResult
	err := client.CallContext(context.Background(), &result, "debug_traceCall", nil, "pending")
	return result, err
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

func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	pending := big.NewInt(-1)
	if number.Cmp(pending) == 0 {
		return "pending"
	}
	return hexutil.EncodeBig(number)
}
