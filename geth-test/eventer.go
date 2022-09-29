// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package geth

import (
	"errors"
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

// EventerABI is the input ABI used to generate the binding from.
// Deprecated: Use EventerMetaData.ABI instead.
var EventerABI = EventerMetaData.ABI

// EventerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EventerMetaData.Bin instead.
var EventerBin = EventerMetaData.Bin

// DeployEventer deploys a new Ethereum contract, binding an instance of Eventer to it.
func DeployEventer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Eventer, error) {
	parsed, err := EventerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EventerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Eventer{EventerCaller: EventerCaller{contract: contract}, EventerTransactor: EventerTransactor{contract: contract}, EventerFilterer: EventerFilterer{contract: contract}}, nil
}

// Eventer is an auto generated Go binding around an Ethereum contract.
type Eventer struct {
	EventerCaller     // Read-only binding to the contract
	EventerTransactor // Write-only binding to the contract
	EventerFilterer   // Log filterer for contract events
}

// EventerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventerSession struct {
	Contract     *Eventer          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventerCallerSession struct {
	Contract *EventerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// EventerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventerTransactorSession struct {
	Contract     *EventerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EventerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventerRaw struct {
	Contract *Eventer // Generic contract binding to access the raw methods on
}

// EventerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventerCallerRaw struct {
	Contract *EventerCaller // Generic read-only contract binding to access the raw methods on
}

// EventerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventerTransactorRaw struct {
	Contract *EventerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventer creates a new instance of Eventer, bound to a specific deployed contract.
func NewEventer(address common.Address, backend bind.ContractBackend) (*Eventer, error) {
	contract, err := bindEventer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eventer{EventerCaller: EventerCaller{contract: contract}, EventerTransactor: EventerTransactor{contract: contract}, EventerFilterer: EventerFilterer{contract: contract}}, nil
}

// NewEventerCaller creates a new read-only instance of Eventer, bound to a specific deployed contract.
func NewEventerCaller(address common.Address, caller bind.ContractCaller) (*EventerCaller, error) {
	contract, err := bindEventer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventerCaller{contract: contract}, nil
}

// NewEventerTransactor creates a new write-only instance of Eventer, bound to a specific deployed contract.
func NewEventerTransactor(address common.Address, transactor bind.ContractTransactor) (*EventerTransactor, error) {
	contract, err := bindEventer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventerTransactor{contract: contract}, nil
}

// NewEventerFilterer creates a new log filterer instance of Eventer, bound to a specific deployed contract.
func NewEventerFilterer(address common.Address, filterer bind.ContractFilterer) (*EventerFilterer, error) {
	contract, err := bindEventer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventerFilterer{contract: contract}, nil
}

// bindEventer binds a generic wrapper to an already deployed contract.
func bindEventer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EventerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eventer *EventerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eventer.Contract.EventerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eventer *EventerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eventer.Contract.EventerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eventer *EventerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eventer.Contract.EventerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eventer *EventerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eventer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eventer *EventerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eventer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eventer *EventerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eventer.Contract.contract.Transact(opts, method, params...)
}

// AnonEvent is a paid mutator transaction binding the contract method 0x529c2b1f.
//
// Solidity: function anonEvent() returns()
func (_Eventer *EventerTransactor) AnonEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eventer.contract.Transact(opts, "anonEvent")
}

// AnonEvent is a paid mutator transaction binding the contract method 0x529c2b1f.
//
// Solidity: function anonEvent() returns()
func (_Eventer *EventerSession) AnonEvent() (*types.Transaction, error) {
	return _Eventer.Contract.AnonEvent(&_Eventer.TransactOpts)
}

// AnonEvent is a paid mutator transaction binding the contract method 0x529c2b1f.
//
// Solidity: function anonEvent() returns()
func (_Eventer *EventerTransactorSession) AnonEvent() (*types.Transaction, error) {
	return _Eventer.Contract.AnonEvent(&_Eventer.TransactOpts)
}

// Fail is a paid mutator transaction binding the contract method 0xa9cc4718.
//
// Solidity: function fail() returns()
func (_Eventer *EventerTransactor) Fail(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eventer.contract.Transact(opts, "fail")
}

// Fail is a paid mutator transaction binding the contract method 0xa9cc4718.
//
// Solidity: function fail() returns()
func (_Eventer *EventerSession) Fail() (*types.Transaction, error) {
	return _Eventer.Contract.Fail(&_Eventer.TransactOpts)
}

// Fail is a paid mutator transaction binding the contract method 0xa9cc4718.
//
// Solidity: function fail() returns()
func (_Eventer *EventerTransactorSession) Fail() (*types.Transaction, error) {
	return _Eventer.Contract.Fail(&_Eventer.TransactOpts)
}

// GetEvent is a paid mutator transaction binding the contract method 0xbf819c20.
//
// Solidity: function getEvent() returns()
func (_Eventer *EventerTransactor) GetEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eventer.contract.Transact(opts, "getEvent")
}

// GetEvent is a paid mutator transaction binding the contract method 0xbf819c20.
//
// Solidity: function getEvent() returns()
func (_Eventer *EventerSession) GetEvent() (*types.Transaction, error) {
	return _Eventer.Contract.GetEvent(&_Eventer.TransactOpts)
}

// GetEvent is a paid mutator transaction binding the contract method 0xbf819c20.
//
// Solidity: function getEvent() returns()
func (_Eventer *EventerTransactorSession) GetEvent() (*types.Transaction, error) {
	return _Eventer.Contract.GetEvent(&_Eventer.TransactOpts)
}

// EventerAnonEventIterator is returned from FilterAnonEvent and is used to iterate over the raw logs and unpacked data for AnonEvent events raised by the Eventer contract.
type EventerAnonEventIterator struct {
	Event *EventerAnonEvent // Event containing the contract specifics and raw log

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
func (it *EventerAnonEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventerAnonEvent)
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
		it.Event = new(EventerAnonEvent)
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
func (it *EventerAnonEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventerAnonEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventerAnonEvent represents a AnonEvent event raised by the Eventer contract.
type EventerAnonEvent struct {
	Arg0 common.Address
	Arg1 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAnonEvent is a free log retrieval operation binding the contract event 0x5aabb901501aff3228b2a4010c287e9d9fe99f2a3d5036dd6cdd2829b567f64f.
//
// Solidity: event AnonEvent(address arg0, address arg1)
func (_Eventer *EventerFilterer) FilterAnonEvent(opts *bind.FilterOpts) (*EventerAnonEventIterator, error) {

	logs, sub, err := _Eventer.contract.FilterLogs(opts, "AnonEvent")
	if err != nil {
		return nil, err
	}
	return &EventerAnonEventIterator{contract: _Eventer.contract, event: "AnonEvent", logs: logs, sub: sub}, nil
}

// WatchAnonEvent is a free log subscription operation binding the contract event 0x5aabb901501aff3228b2a4010c287e9d9fe99f2a3d5036dd6cdd2829b567f64f.
//
// Solidity: event AnonEvent(address arg0, address arg1)
func (_Eventer *EventerFilterer) WatchAnonEvent(opts *bind.WatchOpts, sink chan<- *EventerAnonEvent) (event.Subscription, error) {

	logs, sub, err := _Eventer.contract.WatchLogs(opts, "AnonEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventerAnonEvent)
				if err := _Eventer.contract.UnpackLog(event, "AnonEvent", log); err != nil {
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

// ParseAnonEvent is a log parse operation binding the contract event 0x5aabb901501aff3228b2a4010c287e9d9fe99f2a3d5036dd6cdd2829b567f64f.
//
// Solidity: event AnonEvent(address arg0, address arg1)
func (_Eventer *EventerFilterer) ParseAnonEvent(log types.Log) (*EventerAnonEvent, error) {
	event := new(EventerAnonEvent)
	if err := _Eventer.contract.UnpackLog(event, "AnonEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventerTestInt8Iterator is returned from FilterTestInt8 and is used to iterate over the raw logs and unpacked data for TestInt8 events raised by the Eventer contract.
type EventerTestInt8Iterator struct {
	Event *EventerTestInt8 // Event containing the contract specifics and raw log

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
func (it *EventerTestInt8Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventerTestInt8)
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
		it.Event = new(EventerTestInt8)
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
func (it *EventerTestInt8Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventerTestInt8Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventerTestInt8 represents a TestInt8 event raised by the Eventer contract.
type EventerTestInt8 struct {
	Out1 int8
	Out2 int8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTestInt8 is a free log retrieval operation binding the contract event 0x8f50d21be7587a4814a9d4c10b7c8d1eea6389adbd44cb59ddaba790fd2ecbbd.
//
// Solidity: event TestInt8(int8 indexed out1, int8 indexed out2)
func (_Eventer *EventerFilterer) FilterTestInt8(opts *bind.FilterOpts, out1 []int8, out2 []int8) (*EventerTestInt8Iterator, error) {

	var out1Rule []interface{}
	for _, out1Item := range out1 {
		out1Rule = append(out1Rule, out1Item)
	}
	var out2Rule []interface{}
	for _, out2Item := range out2 {
		out2Rule = append(out2Rule, out2Item)
	}

	logs, sub, err := _Eventer.contract.FilterLogs(opts, "TestInt8", out1Rule, out2Rule)
	if err != nil {
		return nil, err
	}
	return &EventerTestInt8Iterator{contract: _Eventer.contract, event: "TestInt8", logs: logs, sub: sub}, nil
}

// WatchTestInt8 is a free log subscription operation binding the contract event 0x8f50d21be7587a4814a9d4c10b7c8d1eea6389adbd44cb59ddaba790fd2ecbbd.
//
// Solidity: event TestInt8(int8 indexed out1, int8 indexed out2)
func (_Eventer *EventerFilterer) WatchTestInt8(opts *bind.WatchOpts, sink chan<- *EventerTestInt8, out1 []int8, out2 []int8) (event.Subscription, error) {

	var out1Rule []interface{}
	for _, out1Item := range out1 {
		out1Rule = append(out1Rule, out1Item)
	}
	var out2Rule []interface{}
	for _, out2Item := range out2 {
		out2Rule = append(out2Rule, out2Item)
	}

	logs, sub, err := _Eventer.contract.WatchLogs(opts, "TestInt8", out1Rule, out2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventerTestInt8)
				if err := _Eventer.contract.UnpackLog(event, "TestInt8", log); err != nil {
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

// ParseTestInt8 is a log parse operation binding the contract event 0x8f50d21be7587a4814a9d4c10b7c8d1eea6389adbd44cb59ddaba790fd2ecbbd.
//
// Solidity: event TestInt8(int8 indexed out1, int8 indexed out2)
func (_Eventer *EventerFilterer) ParseTestInt8(log types.Log) (*EventerTestInt8, error) {
	event := new(EventerTestInt8)
	if err := _Eventer.contract.UnpackLog(event, "TestInt8", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
