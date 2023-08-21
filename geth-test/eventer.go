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

// EventerMetaData contains all meta data concerning the Eventer contract.
var EventerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"AnonEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int8\",\"name\":\"out1\",\"type\":\"int8\"},{\"indexed\":true,\"internalType\":\"int8\",\"name\":\"out2\",\"type\":\"int8\"}],\"name\":\"TestInt8\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"anonEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061026f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063529c2b1f14610046578063a9cc471814610050578063bf819c201461005a575b600080fd5b61004e610064565b005b61005861009f565b005b6100626100e2565b005b7f5aabb901501aff3228b2a4010c287e9d9fe99f2a3d5036dd6cdd2829b567f64f3333604051610095929190610184565b60405180910390a1565b60006100e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100d7906101ad565b60405180910390fd5b565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe7f8f50d21be7587a4814a9d4c10b7c8d1eea6389adbd44cb59ddaba790fd2ecbbd60405160405180910390a3565b61015b816101de565b82525050565b600061016e600c836101cd565b915061017982610210565b602082019050919050565b60006040820190506101996000830185610152565b6101a66020830184610152565b9392505050565b600060208201905081810360008301526101c681610161565b9050919050565b600082825260208201905092915050565b60006101e9826101f0565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b7f6572726f7220737472696e67000000000000000000000000000000000000000060008201525056fea2646970667358221220f36aeaa995e75d46d14b6dae3d3e12137688d47d5b5bf8de31f5680b9359d82c64736f6c63430008070033",
}

// Eventer is an auto generated Go binding around an Ethereum contract.
type Eventer struct {
	abi abi.ABI
}

// EventerInstance represents a deployed instance of the Eventer contract.
type EventerInstance struct {
	Eventer
	address common.Address
}

func (i *EventerInstance) Address() common.Address {
	return i.address
}

// NewEventer creates a new instance of Eventer.
func NewEventer() (*Eventer, error) {
	parsed, err := EventerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	return &Eventer{abi: *parsed}, nil
}

func (_Eventer *Eventer) PackConstructor() ([]byte, error) {
	return _Eventer.abi.Pack("")
}

// AnonEvent is a free data retrieval call binding the contract method 0x529c2b1f.
//
// Solidity: function anonEvent() returns()
func (_Eventer *Eventer) PackAnonEvent() ([]byte, error) {
	return _Eventer.abi.Pack("anonEvent")
}

func (_Eventer *Eventer) UnpackAnonEvent(data []byte) (struct{}, error) {
	out, err := _Eventer.abi.Unpack("anonEvent", data)
	_ = out
	outstruct := new(struct{})
	if err != nil {
		return *outstruct, err
	}

	return *outstruct, err

}

// Fail is a free data retrieval call binding the contract method 0xa9cc4718.
//
// Solidity: function fail() returns()
func (_Eventer *Eventer) PackFail() ([]byte, error) {
	return _Eventer.abi.Pack("fail")
}

func (_Eventer *Eventer) UnpackFail(data []byte) (struct{}, error) {
	out, err := _Eventer.abi.Unpack("fail", data)
	_ = out
	outstruct := new(struct{})
	if err != nil {
		return *outstruct, err
	}

	return *outstruct, err

}

// GetEvent is a free data retrieval call binding the contract method 0xbf819c20.
//
// Solidity: function getEvent() returns()
func (_Eventer *Eventer) PackGetEvent() ([]byte, error) {
	return _Eventer.abi.Pack("getEvent")
}

func (_Eventer *Eventer) UnpackGetEvent(data []byte) (struct{}, error) {
	out, err := _Eventer.abi.Unpack("getEvent", data)
	_ = out
	outstruct := new(struct{})
	if err != nil {
		return *outstruct, err
	}

	return *outstruct, err

}

// EventerAnonEvent represents a AnonEvent event raised by the Eventer contract.
type EventerAnonEvent struct {
	Arg0 common.Address
	Arg1 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

func (_Eventer *Eventer) AnonEventEventID() common.Hash {
	return common.HexToHash("0x5aabb901501aff3228b2a4010c287e9d9fe99f2a3d5036dd6cdd2829b567f64f")
}

func (_Eventer *Eventer) UnpackAnonEventEvent(log types.Log) (*EventerAnonEvent, error) {
	event := "AnonEvent"
	if log.Topics[0] != _Eventer.abi.Events[event].ID {
		return nil, fmt.Errorf("event signature mismatch")
	}
	out := new(EventerAnonEvent)
	if len(log.Data) > 0 {
		if err := _Eventer.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range _Eventer.abi.Events[event].Inputs {
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

// EventerTestInt8 represents a TestInt8 event raised by the Eventer contract.
type EventerTestInt8 struct {
	Out1 int8
	Out2 int8
	Raw  types.Log // Blockchain specific contextual infos
}

func (_Eventer *Eventer) TestInt8EventID() common.Hash {
	return common.HexToHash("0x8f50d21be7587a4814a9d4c10b7c8d1eea6389adbd44cb59ddaba790fd2ecbbd")
}

func (_Eventer *Eventer) UnpackTestInt8Event(log types.Log) (*EventerTestInt8, error) {
	event := "TestInt8"
	if log.Topics[0] != _Eventer.abi.Events[event].ID {
		return nil, fmt.Errorf("event signature mismatch")
	}
	out := new(EventerTestInt8)
	if len(log.Data) > 0 {
		if err := _Eventer.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range _Eventer.abi.Events[event].Inputs {
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
