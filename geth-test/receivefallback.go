// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package geth

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ReceiveFallbackMetaData contains all meta data concerning the ReceiveFallback contract.
var ReceiveFallbackMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Fallback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Receive\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061010e806100206000396000f3fe608060405236603f577f681ad3cb4f20e39dd01bc0e45f0bc6f95ea2de196277ed6ec4f7357e2e2044b83360405160359190608f565b60405180910390a1005b348015604a57600080fd5b507fdd0c5778ddb0b0501649417aeffed44057db3d5d530231758c8dc5ef6eeb78943360405160789190608f565b60405180910390a1005b60898160a8565b82525050565b600060208201905060a260008301846082565b92915050565b600060b18260b8565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff8216905091905056fea2646970667358221220ff87533080fe8ad33a417e9dcec6f37cc5f89614a47474a3b547ef07fc7743a464736f6c63430008070033",
}

// ReceiveFallback is an auto generated Go binding around an Ethereum contract.
type ReceiveFallback struct {
	abi abi.ABI
}

// ReceiveFallbackInstance represents a deployed instance of the ReceiveFallback contract.
type ReceiveFallbackInstance struct {
	ReceiveFallback
	address common.Address
}

func (i *ReceiveFallbackInstance) Address() common.Address {
	return i.address
}

// NewReceiveFallback creates a new instance of ReceiveFallback.
func NewReceiveFallback() (*ReceiveFallback, error) {
	parsed, err := ReceiveFallbackMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	return &ReceiveFallback{abi: *parsed}, nil
}

func (_ReceiveFallback *ReceiveFallback) PackConstructor() ([]byte, error) {
	return _ReceiveFallback.abi.Pack("")
}

// ReceiveFallbackFallback represents a Fallback event raised by the ReceiveFallback contract.
type ReceiveFallbackFallback struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

func (_ReceiveFallback *ReceiveFallback) FallbackEventID() common.Hash {
	return common.HexToHash("0xdd0c5778ddb0b0501649417aeffed44057db3d5d530231758c8dc5ef6eeb7894")
}

func (_ReceiveFallback *ReceiveFallback) UnpackFallbackEvent(log types.Log) (*ReceiveFallbackFallback, error) {
	event := "Fallback"
	if log.Topics[0] != _ReceiveFallback.abi.Events[event].ID {
		return nil, fmt.Errorf("event signature mismatch")
	}
	out := new(ReceiveFallbackFallback)
	if len(log.Data) > 0 {
		if err := _ReceiveFallback.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range _ReceiveFallback.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ReceiveFallbackReceive represents a Receive event raised by the ReceiveFallback contract.
type ReceiveFallbackReceive struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

func (_ReceiveFallback *ReceiveFallback) ReceiveEventID() common.Hash {
	return common.HexToHash("0x681ad3cb4f20e39dd01bc0e45f0bc6f95ea2de196277ed6ec4f7357e2e2044b8")
}

func (_ReceiveFallback *ReceiveFallback) UnpackReceiveEvent(log types.Log) (*ReceiveFallbackReceive, error) {
	event := "Receive"
	if log.Topics[0] != _ReceiveFallback.abi.Events[event].ID {
		return nil, fmt.Errorf("event signature mismatch")
	}
	out := new(ReceiveFallbackReceive)
	if len(log.Data) > 0 {
		if err := _ReceiveFallback.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range _ReceiveFallback.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}
