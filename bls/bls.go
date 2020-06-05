// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CallBLSABI is the input ABI used to generate the binding from.
const CallBLSABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"Bytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Uint\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"addr\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"callPrec\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CallBLSFuncSigs maps the 4-byte function signature to its string representation.
var CallBLSFuncSigs = map[string]string{
	"f6532808": "callPrec(uint256,bytes)",
}

// CallBLSBin is the compiled bytecode used for deploying new contracts.
var CallBLSBin = "0x608060405234801561001057600080fd5b506102ba806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063f653280814610030575b600080fd5b6100dd6004803603604081101561004657600080fd5b8135919081019060408101602082013564010000000081111561006857600080fd5b82018360208201111561007a57600080fd5b8035906020019184600183028401116401000000008311171561009c57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506100df945050505050565b005b600080600080600085519050600087600a14806100fc575087600b145b80610107575087600c145b806101125750876011145b1561011f575060806101a5565b87600d148061012e575087600e145b80610139575087600f145b806101445750876012145b1561015257506101006101a5565b8760101415610163575060206101a5565b6040805162461bcd60e51b8152602060048201526012602482015271496e76616c696420707265636f6d70696c6560701b604482015290519081900360640190fd5b6040518181848a8c5afa6101b857600080fd5b8051965060208101519550604081015194506060810151935050600080516020610265833981519152866040518082815260200191505060405180910390a16040805186815290516000805160206102658339815191529181900360200190a16040805185815290516000805160206102658339815191529181900360200190a16040805184815290516000805160206102658339815191529181900360200190a1505050505050505056fe432cc6516825015332dc662a569788a7b0790ea2c6a53a601c743c54cbdafe60a2646970667358221220f1a1502801d53d989cf97153ee69b0b94a58be06ab006d8084046caabe46848564736f6c63430006050033"

// DeployCallBLS deploys a new Ethereum contract, binding an instance of CallBLS to it.
func DeployCallBLS(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CallBLS, error) {
	parsed, err := abi.JSON(strings.NewReader(CallBLSABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CallBLSBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CallBLS{CallBLSCaller: CallBLSCaller{contract: contract}, CallBLSTransactor: CallBLSTransactor{contract: contract}, CallBLSFilterer: CallBLSFilterer{contract: contract}}, nil
}

// CallBLS is an auto generated Go binding around an Ethereum contract.
type CallBLS struct {
	CallBLSCaller     // Read-only binding to the contract
	CallBLSTransactor // Write-only binding to the contract
	CallBLSFilterer   // Log filterer for contract events
}

// CallBLSCaller is an auto generated read-only Go binding around an Ethereum contract.
type CallBLSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallBLSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CallBLSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallBLSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CallBLSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallBLSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CallBLSSession struct {
	Contract     *CallBLS          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallBLSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CallBLSCallerSession struct {
	Contract *CallBLSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CallBLSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CallBLSTransactorSession struct {
	Contract     *CallBLSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CallBLSRaw is an auto generated low-level Go binding around an Ethereum contract.
type CallBLSRaw struct {
	Contract *CallBLS // Generic contract binding to access the raw methods on
}

// CallBLSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CallBLSCallerRaw struct {
	Contract *CallBLSCaller // Generic read-only contract binding to access the raw methods on
}

// CallBLSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CallBLSTransactorRaw struct {
	Contract *CallBLSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCallBLS creates a new instance of CallBLS, bound to a specific deployed contract.
func NewCallBLS(address common.Address, backend bind.ContractBackend) (*CallBLS, error) {
	contract, err := bindCallBLS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CallBLS{CallBLSCaller: CallBLSCaller{contract: contract}, CallBLSTransactor: CallBLSTransactor{contract: contract}, CallBLSFilterer: CallBLSFilterer{contract: contract}}, nil
}

// NewCallBLSCaller creates a new read-only instance of CallBLS, bound to a specific deployed contract.
func NewCallBLSCaller(address common.Address, caller bind.ContractCaller) (*CallBLSCaller, error) {
	contract, err := bindCallBLS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CallBLSCaller{contract: contract}, nil
}

// NewCallBLSTransactor creates a new write-only instance of CallBLS, bound to a specific deployed contract.
func NewCallBLSTransactor(address common.Address, transactor bind.ContractTransactor) (*CallBLSTransactor, error) {
	contract, err := bindCallBLS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CallBLSTransactor{contract: contract}, nil
}

// NewCallBLSFilterer creates a new log filterer instance of CallBLS, bound to a specific deployed contract.
func NewCallBLSFilterer(address common.Address, filterer bind.ContractFilterer) (*CallBLSFilterer, error) {
	contract, err := bindCallBLS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CallBLSFilterer{contract: contract}, nil
}

// bindCallBLS binds a generic wrapper to an already deployed contract.
func bindCallBLS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CallBLSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CallBLS *CallBLSRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CallBLS.Contract.CallBLSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CallBLS *CallBLSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CallBLS.Contract.CallBLSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CallBLS *CallBLSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CallBLS.Contract.CallBLSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CallBLS *CallBLSCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CallBLS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CallBLS *CallBLSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CallBLS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CallBLS *CallBLSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CallBLS.Contract.contract.Transact(opts, method, params...)
}

// CallPrec is a paid mutator transaction binding the contract method 0xf6532808.
//
// Solidity: function callPrec(uint256 addr, bytes input) returns()
func (_CallBLS *CallBLSTransactor) CallPrec(opts *bind.TransactOpts, addr *big.Int, input []byte) (*types.Transaction, error) {
	return _CallBLS.contract.Transact(opts, "callPrec", addr, input)
}

// CallPrec is a paid mutator transaction binding the contract method 0xf6532808.
//
// Solidity: function callPrec(uint256 addr, bytes input) returns()
func (_CallBLS *CallBLSSession) CallPrec(addr *big.Int, input []byte) (*types.Transaction, error) {
	return _CallBLS.Contract.CallPrec(&_CallBLS.TransactOpts, addr, input)
}

// CallPrec is a paid mutator transaction binding the contract method 0xf6532808.
//
// Solidity: function callPrec(uint256 addr, bytes input) returns()
func (_CallBLS *CallBLSTransactorSession) CallPrec(addr *big.Int, input []byte) (*types.Transaction, error) {
	return _CallBLS.Contract.CallPrec(&_CallBLS.TransactOpts, addr, input)
}

// CallBLSBytesIterator is returned from FilterBytes and is used to iterate over the raw logs and unpacked data for Bytes events raised by the CallBLS contract.
type CallBLSBytesIterator struct {
	Event *CallBLSBytes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CallBLSBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CallBLSBytes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CallBLSBytes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CallBLSBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CallBLSBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CallBLSBytes represents a Bytes event raised by the CallBLS contract.
type CallBLSBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBytes is a free log retrieval operation binding the contract event 0x24ace672d0b56bb74d7b34faba262f0ca4bc2f84bc4494f66422a1c5f810d76b.
//
// Solidity: event Bytes(bytes arg0)
func (_CallBLS *CallBLSFilterer) FilterBytes(opts *bind.FilterOpts) (*CallBLSBytesIterator, error) {

	logs, sub, err := _CallBLS.contract.FilterLogs(opts, "Bytes")
	if err != nil {
		return nil, err
	}
	return &CallBLSBytesIterator{contract: _CallBLS.contract, event: "Bytes", logs: logs, sub: sub}, nil
}

// WatchBytes is a free log subscription operation binding the contract event 0x24ace672d0b56bb74d7b34faba262f0ca4bc2f84bc4494f66422a1c5f810d76b.
//
// Solidity: event Bytes(bytes arg0)
func (_CallBLS *CallBLSFilterer) WatchBytes(opts *bind.WatchOpts, sink chan<- *CallBLSBytes) (event.Subscription, error) {

	logs, sub, err := _CallBLS.contract.WatchLogs(opts, "Bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CallBLSBytes)
				if err := _CallBLS.contract.UnpackLog(event, "Bytes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBytes is a log parse operation binding the contract event 0x24ace672d0b56bb74d7b34faba262f0ca4bc2f84bc4494f66422a1c5f810d76b.
//
// Solidity: event Bytes(bytes arg0)
func (_CallBLS *CallBLSFilterer) ParseBytes(log types.Log) (*CallBLSBytes, error) {
	event := new(CallBLSBytes)
	if err := _CallBLS.contract.UnpackLog(event, "Bytes", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CallBLSUintIterator is returned from FilterUint and is used to iterate over the raw logs and unpacked data for Uint events raised by the CallBLS contract.
type CallBLSUintIterator struct {
	Event *CallBLSUint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CallBLSUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CallBLSUint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CallBLSUint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CallBLSUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CallBLSUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CallBLSUint represents a Uint event raised by the CallBLS contract.
type CallBLSUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUint is a free log retrieval operation binding the contract event 0x432cc6516825015332dc662a569788a7b0790ea2c6a53a601c743c54cbdafe60.
//
// Solidity: event Uint(uint256 arg0)
func (_CallBLS *CallBLSFilterer) FilterUint(opts *bind.FilterOpts) (*CallBLSUintIterator, error) {

	logs, sub, err := _CallBLS.contract.FilterLogs(opts, "Uint")
	if err != nil {
		return nil, err
	}
	return &CallBLSUintIterator{contract: _CallBLS.contract, event: "Uint", logs: logs, sub: sub}, nil
}

// WatchUint is a free log subscription operation binding the contract event 0x432cc6516825015332dc662a569788a7b0790ea2c6a53a601c743c54cbdafe60.
//
// Solidity: event Uint(uint256 arg0)
func (_CallBLS *CallBLSFilterer) WatchUint(opts *bind.WatchOpts, sink chan<- *CallBLSUint) (event.Subscription, error) {

	logs, sub, err := _CallBLS.contract.WatchLogs(opts, "Uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CallBLSUint)
				if err := _CallBLS.contract.UnpackLog(event, "Uint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUint is a log parse operation binding the contract event 0x432cc6516825015332dc662a569788a7b0790ea2c6a53a601c743c54cbdafe60.
//
// Solidity: event Uint(uint256 arg0)
func (_CallBLS *CallBLSFilterer) ParseUint(log types.Log) (*CallBLSUint, error) {
	event := new(CallBLSUint)
	if err := _CallBLS.contract.UnpackLog(event, "Uint", log); err != nil {
		return nil, err
	}
	return event, nil
}
