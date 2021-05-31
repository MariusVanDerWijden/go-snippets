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

// CallProxyABI is the input ABI used to generate the binding from.
const CallProxyABI = "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTarget\",\"type\":\"address\"}],\"name\":\"setTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CallProxyFuncSigs maps the 4-byte function signature to its string representation.
var CallProxyFuncSigs = map[string]string{
	"776d1a01": "setTarget(address)",
}

// CallProxyBin is the compiled bytecode used for deploying new contracts.
var CallProxyBin = "0x6080604052348015600f57600080fd5b5060bf8061001e6000396000f3fe608060405260043610601c5760003560e01c8063776d1a01146036575b60008054368283378182368434855af150503d81823e3d81f35b348015604157600080fd5b50606560048036036020811015605657600080fd5b50356001600160a01b03166067565b005b600080546001600160a01b0319166001600160a01b039290921691909117905556fea26469706673582212203e9b42776313ec41db0c042e3a419af90a45e3d9a527b742d8395a8ebe1593b064736f6c63430007000033"

// DeployCallProxy deploys a new Ethereum contract, binding an instance of CallProxy to it.
func DeployCallProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CallProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(CallProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CallProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CallProxy{CallProxyCaller: CallProxyCaller{contract: contract}, CallProxyTransactor: CallProxyTransactor{contract: contract}, CallProxyFilterer: CallProxyFilterer{contract: contract}}, nil
}

// CallProxy is an auto generated Go binding around an Ethereum contract.
type CallProxy struct {
	CallProxyCaller     // Read-only binding to the contract
	CallProxyTransactor // Write-only binding to the contract
	CallProxyFilterer   // Log filterer for contract events
}

// CallProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type CallProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CallProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CallProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CallProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CallProxySession struct {
	Contract     *CallProxy        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CallProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CallProxyCallerSession struct {
	Contract *CallProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CallProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CallProxyTransactorSession struct {
	Contract     *CallProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CallProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type CallProxyRaw struct {
	Contract *CallProxy // Generic contract binding to access the raw methods on
}

// CallProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CallProxyCallerRaw struct {
	Contract *CallProxyCaller // Generic read-only contract binding to access the raw methods on
}

// CallProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CallProxyTransactorRaw struct {
	Contract *CallProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCallProxy creates a new instance of CallProxy, bound to a specific deployed contract.
func NewCallProxy(address common.Address, backend bind.ContractBackend) (*CallProxy, error) {
	contract, err := bindCallProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CallProxy{CallProxyCaller: CallProxyCaller{contract: contract}, CallProxyTransactor: CallProxyTransactor{contract: contract}, CallProxyFilterer: CallProxyFilterer{contract: contract}}, nil
}

// NewCallProxyCaller creates a new read-only instance of CallProxy, bound to a specific deployed contract.
func NewCallProxyCaller(address common.Address, caller bind.ContractCaller) (*CallProxyCaller, error) {
	contract, err := bindCallProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CallProxyCaller{contract: contract}, nil
}

// NewCallProxyTransactor creates a new write-only instance of CallProxy, bound to a specific deployed contract.
func NewCallProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*CallProxyTransactor, error) {
	contract, err := bindCallProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CallProxyTransactor{contract: contract}, nil
}

// NewCallProxyFilterer creates a new log filterer instance of CallProxy, bound to a specific deployed contract.
func NewCallProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*CallProxyFilterer, error) {
	contract, err := bindCallProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CallProxyFilterer{contract: contract}, nil
}

// bindCallProxy binds a generic wrapper to an already deployed contract.
func bindCallProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CallProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CallProxy *CallProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CallProxy.Contract.CallProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CallProxy *CallProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CallProxy.Contract.CallProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CallProxy *CallProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CallProxy.Contract.CallProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CallProxy *CallProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CallProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CallProxy *CallProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CallProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CallProxy *CallProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CallProxy.Contract.contract.Transact(opts, method, params...)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address newTarget) returns()
func (_CallProxy *CallProxyTransactor) SetTarget(opts *bind.TransactOpts, newTarget common.Address) (*types.Transaction, error) {
	return _CallProxy.contract.Transact(opts, "setTarget", newTarget)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address newTarget) returns()
func (_CallProxy *CallProxySession) SetTarget(newTarget common.Address) (*types.Transaction, error) {
	return _CallProxy.Contract.SetTarget(&_CallProxy.TransactOpts, newTarget)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address newTarget) returns()
func (_CallProxy *CallProxyTransactorSession) SetTarget(newTarget common.Address) (*types.Transaction, error) {
	return _CallProxy.Contract.SetTarget(&_CallProxy.TransactOpts, newTarget)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CallProxy *CallProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _CallProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CallProxy *CallProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CallProxy.Contract.Fallback(&_CallProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CallProxy *CallProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CallProxy.Contract.Fallback(&_CallProxy.TransactOpts, calldata)
}

// CcallProxyABI is the input ABI used to generate the binding from.
const CcallProxyABI = "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTarget\",\"type\":\"address\"}],\"name\":\"setTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CcallProxyFuncSigs maps the 4-byte function signature to its string representation.
var CcallProxyFuncSigs = map[string]string{
	"776d1a01": "setTarget(address)",
}

// CcallProxyBin is the compiled bytecode used for deploying new contracts.
var CcallProxyBin = "0x6080604052348015600f57600080fd5b5060bf8061001e6000396000f3fe608060405260043610601c5760003560e01c8063776d1a01146036575b60008054368283378182368434855af250503d81823e3d81f35b348015604157600080fd5b50606560048036036020811015605657600080fd5b50356001600160a01b03166067565b005b600080546001600160a01b0319166001600160a01b039290921691909117905556fea26469706673582212201007b5fe46ab1717286fce54256c3de37a2304218108b9cf8ae216e2fd73eaeb64736f6c63430007000033"

// DeployCcallProxy deploys a new Ethereum contract, binding an instance of CcallProxy to it.
func DeployCcallProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CcallProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(CcallProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CcallProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CcallProxy{CcallProxyCaller: CcallProxyCaller{contract: contract}, CcallProxyTransactor: CcallProxyTransactor{contract: contract}, CcallProxyFilterer: CcallProxyFilterer{contract: contract}}, nil
}

// CcallProxy is an auto generated Go binding around an Ethereum contract.
type CcallProxy struct {
	CcallProxyCaller     // Read-only binding to the contract
	CcallProxyTransactor // Write-only binding to the contract
	CcallProxyFilterer   // Log filterer for contract events
}

// CcallProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type CcallProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CcallProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CcallProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CcallProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CcallProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CcallProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CcallProxySession struct {
	Contract     *CcallProxy       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CcallProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CcallProxyCallerSession struct {
	Contract *CcallProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CcallProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CcallProxyTransactorSession struct {
	Contract     *CcallProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CcallProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type CcallProxyRaw struct {
	Contract *CcallProxy // Generic contract binding to access the raw methods on
}

// CcallProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CcallProxyCallerRaw struct {
	Contract *CcallProxyCaller // Generic read-only contract binding to access the raw methods on
}

// CcallProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CcallProxyTransactorRaw struct {
	Contract *CcallProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCcallProxy creates a new instance of CcallProxy, bound to a specific deployed contract.
func NewCcallProxy(address common.Address, backend bind.ContractBackend) (*CcallProxy, error) {
	contract, err := bindCcallProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CcallProxy{CcallProxyCaller: CcallProxyCaller{contract: contract}, CcallProxyTransactor: CcallProxyTransactor{contract: contract}, CcallProxyFilterer: CcallProxyFilterer{contract: contract}}, nil
}

// NewCcallProxyCaller creates a new read-only instance of CcallProxy, bound to a specific deployed contract.
func NewCcallProxyCaller(address common.Address, caller bind.ContractCaller) (*CcallProxyCaller, error) {
	contract, err := bindCcallProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CcallProxyCaller{contract: contract}, nil
}

// NewCcallProxyTransactor creates a new write-only instance of CcallProxy, bound to a specific deployed contract.
func NewCcallProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*CcallProxyTransactor, error) {
	contract, err := bindCcallProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CcallProxyTransactor{contract: contract}, nil
}

// NewCcallProxyFilterer creates a new log filterer instance of CcallProxy, bound to a specific deployed contract.
func NewCcallProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*CcallProxyFilterer, error) {
	contract, err := bindCcallProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CcallProxyFilterer{contract: contract}, nil
}

// bindCcallProxy binds a generic wrapper to an already deployed contract.
func bindCcallProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CcallProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CcallProxy *CcallProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CcallProxy.Contract.CcallProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CcallProxy *CcallProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CcallProxy.Contract.CcallProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CcallProxy *CcallProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CcallProxy.Contract.CcallProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CcallProxy *CcallProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CcallProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CcallProxy *CcallProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CcallProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CcallProxy *CcallProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CcallProxy.Contract.contract.Transact(opts, method, params...)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address newTarget) returns()
func (_CcallProxy *CcallProxyTransactor) SetTarget(opts *bind.TransactOpts, newTarget common.Address) (*types.Transaction, error) {
	return _CcallProxy.contract.Transact(opts, "setTarget", newTarget)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address newTarget) returns()
func (_CcallProxy *CcallProxySession) SetTarget(newTarget common.Address) (*types.Transaction, error) {
	return _CcallProxy.Contract.SetTarget(&_CcallProxy.TransactOpts, newTarget)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address newTarget) returns()
func (_CcallProxy *CcallProxyTransactorSession) SetTarget(newTarget common.Address) (*types.Transaction, error) {
	return _CcallProxy.Contract.SetTarget(&_CcallProxy.TransactOpts, newTarget)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CcallProxy *CcallProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _CcallProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CcallProxy *CcallProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CcallProxy.Contract.Fallback(&_CcallProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CcallProxy *CcallProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CcallProxy.Contract.Fallback(&_CcallProxy.TransactOpts, calldata)
}

// DcallProxyABI is the input ABI used to generate the binding from.
const DcallProxyABI = "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTarget\",\"type\":\"address\"}],\"name\":\"setTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DcallProxyFuncSigs maps the 4-byte function signature to its string representation.
var DcallProxyFuncSigs = map[string]string{
	"776d1a01": "setTarget(address)",
}

// DcallProxyBin is the compiled bytecode used for deploying new contracts.
var DcallProxyBin = "0x6080604052348015600f57600080fd5b5060be8061001e6000396000f3fe608060405260043610601c5760003560e01c8063776d1a01146035575b600080543682833781823684845af450503d81823e3d81f35b348015604057600080fd5b50606460048036036020811015605557600080fd5b50356001600160a01b03166066565b005b600080546001600160a01b0319166001600160a01b039290921691909117905556fea264697066735822122037397d372e10f092bc3636054286ebfad26a6057759d56871e044cef9ca7131364736f6c63430007000033"

// DeployDcallProxy deploys a new Ethereum contract, binding an instance of DcallProxy to it.
func DeployDcallProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DcallProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(DcallProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DcallProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DcallProxy{DcallProxyCaller: DcallProxyCaller{contract: contract}, DcallProxyTransactor: DcallProxyTransactor{contract: contract}, DcallProxyFilterer: DcallProxyFilterer{contract: contract}}, nil
}

// DcallProxy is an auto generated Go binding around an Ethereum contract.
type DcallProxy struct {
	DcallProxyCaller     // Read-only binding to the contract
	DcallProxyTransactor // Write-only binding to the contract
	DcallProxyFilterer   // Log filterer for contract events
}

// DcallProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type DcallProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DcallProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DcallProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DcallProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DcallProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DcallProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DcallProxySession struct {
	Contract     *DcallProxy       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DcallProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DcallProxyCallerSession struct {
	Contract *DcallProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DcallProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DcallProxyTransactorSession struct {
	Contract     *DcallProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DcallProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type DcallProxyRaw struct {
	Contract *DcallProxy // Generic contract binding to access the raw methods on
}

// DcallProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DcallProxyCallerRaw struct {
	Contract *DcallProxyCaller // Generic read-only contract binding to access the raw methods on
}

// DcallProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DcallProxyTransactorRaw struct {
	Contract *DcallProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDcallProxy creates a new instance of DcallProxy, bound to a specific deployed contract.
func NewDcallProxy(address common.Address, backend bind.ContractBackend) (*DcallProxy, error) {
	contract, err := bindDcallProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DcallProxy{DcallProxyCaller: DcallProxyCaller{contract: contract}, DcallProxyTransactor: DcallProxyTransactor{contract: contract}, DcallProxyFilterer: DcallProxyFilterer{contract: contract}}, nil
}

// NewDcallProxyCaller creates a new read-only instance of DcallProxy, bound to a specific deployed contract.
func NewDcallProxyCaller(address common.Address, caller bind.ContractCaller) (*DcallProxyCaller, error) {
	contract, err := bindDcallProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DcallProxyCaller{contract: contract}, nil
}

// NewDcallProxyTransactor creates a new write-only instance of DcallProxy, bound to a specific deployed contract.
func NewDcallProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*DcallProxyTransactor, error) {
	contract, err := bindDcallProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DcallProxyTransactor{contract: contract}, nil
}

// NewDcallProxyFilterer creates a new log filterer instance of DcallProxy, bound to a specific deployed contract.
func NewDcallProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*DcallProxyFilterer, error) {
	contract, err := bindDcallProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DcallProxyFilterer{contract: contract}, nil
}

// bindDcallProxy binds a generic wrapper to an already deployed contract.
func bindDcallProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DcallProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DcallProxy *DcallProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DcallProxy.Contract.DcallProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DcallProxy *DcallProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DcallProxy.Contract.DcallProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DcallProxy *DcallProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DcallProxy.Contract.DcallProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DcallProxy *DcallProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DcallProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DcallProxy *DcallProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DcallProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DcallProxy *DcallProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DcallProxy.Contract.contract.Transact(opts, method, params...)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address newTarget) returns()
func (_DcallProxy *DcallProxyTransactor) SetTarget(opts *bind.TransactOpts, newTarget common.Address) (*types.Transaction, error) {
	return _DcallProxy.contract.Transact(opts, "setTarget", newTarget)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address newTarget) returns()
func (_DcallProxy *DcallProxySession) SetTarget(newTarget common.Address) (*types.Transaction, error) {
	return _DcallProxy.Contract.SetTarget(&_DcallProxy.TransactOpts, newTarget)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address newTarget) returns()
func (_DcallProxy *DcallProxyTransactorSession) SetTarget(newTarget common.Address) (*types.Transaction, error) {
	return _DcallProxy.Contract.SetTarget(&_DcallProxy.TransactOpts, newTarget)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DcallProxy *DcallProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _DcallProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DcallProxy *DcallProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DcallProxy.Contract.Fallback(&_DcallProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_DcallProxy *DcallProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DcallProxy.Contract.Fallback(&_DcallProxy.TransactOpts, calldata)
}

// DeployerABI is the input ABI used to generate the binding from.
const DeployerABI = "[{\"inputs\":[],\"name\":\"deploy\",\"outputs\":[{\"internalType\":\"address[5]\",\"name\":\"\",\"type\":\"address[5]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DeployerFuncSigs maps the 4-byte function signature to its string representation.
var DeployerFuncSigs = map[string]string{
	"775c300c": "deploy()",
}

// DeployerBin is the compiled bytecode used for deploying new contracts.
var DeployerBin = "0x608060405234801561001057600080fd5b506106dc806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063775c300c14610030575b600080fd5b610038610070565b604051808260a080838360005b8381101561005d578181015183820152602001610045565b5050505090500191505060405180910390f35b610078610197565b610080610197565b60405161008c906101b5565b604051809103906000f0801580156100a8573d6000803e3d6000fd5b506001600160a01b031681526040516100c0906101c2565b604051809103906000f0801580156100dc573d6000803e3d6000fd5b506001600160a01b031660208201526040516100f7906101ce565b604051809103906000f080158015610113573d6000803e3d6000fd5b506001600160a01b031660408083019190915251610130906101da565b604051809103906000f08015801561014c573d6000803e3d6000fd5b506001600160a01b03166060820152604051610167906101e6565b604051809103906000f080158015610183573d6000803e3d6000fd5b506001600160a01b03166080820152919050565b6040518060a001604052806005906020820280368337509192915050565b610142806101f383390190565b60dc8061033583390190565b60dd8061041183390190565b60dd806104ee83390190565b60dc806105cb8339019056fe608060405234801561001057600080fd5b50610122806100206000396000f3fe608060405260043610603a5760003560e01c80633405d6bf14603f57806361da1439146061578063df8db75f14608d578063e2e52ec11460a7575b600080fd5b605f60048036036040811015605357600080fd5b508035906020013560c7565b005b607b60048036036020811015607557600080fd5b503560cf565b60408051918252519081900360200190f35b607b6004803603602081101560a157600080fd5b503560dc565b605f6004803603604081101560bb57600080fd5b508035906020013560e8565b808255600080fd5b6000815460005260206000f35b60008154600052600080fd5b905556fea2646970667358221220d1860a7c6440b5ce125750011e12a85bde2e9c3ae0d061c63cb8138c86c3cbdb64736f6c634300070000336080604052348015600f57600080fd5b5060be8061001e6000396000f3fe608060405260043610601c5760003560e01c8063776d1a01146035575b600080543682833781823684845af450503d81823e3d81f35b348015604057600080fd5b50606460048036036020811015605557600080fd5b50356001600160a01b03166066565b005b600080546001600160a01b0319166001600160a01b039290921691909117905556fea264697066735822122037397d372e10f092bc3636054286ebfad26a6057759d56871e044cef9ca7131364736f6c634300070000336080604052348015600f57600080fd5b5060bf8061001e6000396000f3fe608060405260043610601c5760003560e01c8063776d1a01146036575b60008054368283378182368434855af150503d81823e3d81f35b348015604157600080fd5b50606560048036036020811015605657600080fd5b50356001600160a01b03166067565b005b600080546001600160a01b0319166001600160a01b039290921691909117905556fea26469706673582212203e9b42776313ec41db0c042e3a419af90a45e3d9a527b742d8395a8ebe1593b064736f6c634300070000336080604052348015600f57600080fd5b5060bf8061001e6000396000f3fe608060405260043610601c5760003560e01c8063776d1a01146036575b60008054368283378182368434855af250503d81823e3d81f35b348015604157600080fd5b50606560048036036020811015605657600080fd5b50356001600160a01b03166067565b005b600080546001600160a01b0319166001600160a01b039290921691909117905556fea26469706673582212201007b5fe46ab1717286fce54256c3de37a2304218108b9cf8ae216e2fd73eaeb64736f6c634300070000336080604052348015600f57600080fd5b5060be8061001e6000396000f3fe608060405260043610601c5760003560e01c8063776d1a01146035575b600080543682833781823684845afa50503d81823e3d81f35b348015604057600080fd5b50606460048036036020811015605557600080fd5b50356001600160a01b03166066565b005b600080546001600160a01b0319166001600160a01b039290921691909117905556fea2646970667358221220c3813fa3ef946ecb2fff63101d0408b82cf1c28b827ce6d67ce0685e73b8467764736f6c63430007000033a26469706673582212207c47c4f530df88173ed1ced99df8725ff5cd472758f4a405959feafc8be3d0b164736f6c63430007000033"

// DeployDeployer deploys a new Ethereum contract, binding an instance of Deployer to it.
func DeployDeployer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Deployer, error) {
	parsed, err := abi.JSON(strings.NewReader(DeployerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DeployerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Deployer{DeployerCaller: DeployerCaller{contract: contract}, DeployerTransactor: DeployerTransactor{contract: contract}, DeployerFilterer: DeployerFilterer{contract: contract}}, nil
}

// Deployer is an auto generated Go binding around an Ethereum contract.
type Deployer struct {
	DeployerCaller     // Read-only binding to the contract
	DeployerTransactor // Write-only binding to the contract
	DeployerFilterer   // Log filterer for contract events
}

// DeployerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeployerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeployerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeployerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeployerSession struct {
	Contract     *Deployer         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeployerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeployerCallerSession struct {
	Contract *DeployerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DeployerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeployerTransactorSession struct {
	Contract     *DeployerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DeployerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeployerRaw struct {
	Contract *Deployer // Generic contract binding to access the raw methods on
}

// DeployerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeployerCallerRaw struct {
	Contract *DeployerCaller // Generic read-only contract binding to access the raw methods on
}

// DeployerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeployerTransactorRaw struct {
	Contract *DeployerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeployer creates a new instance of Deployer, bound to a specific deployed contract.
func NewDeployer(address common.Address, backend bind.ContractBackend) (*Deployer, error) {
	contract, err := bindDeployer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Deployer{DeployerCaller: DeployerCaller{contract: contract}, DeployerTransactor: DeployerTransactor{contract: contract}, DeployerFilterer: DeployerFilterer{contract: contract}}, nil
}

// NewDeployerCaller creates a new read-only instance of Deployer, bound to a specific deployed contract.
func NewDeployerCaller(address common.Address, caller bind.ContractCaller) (*DeployerCaller, error) {
	contract, err := bindDeployer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeployerCaller{contract: contract}, nil
}

// NewDeployerTransactor creates a new write-only instance of Deployer, bound to a specific deployed contract.
func NewDeployerTransactor(address common.Address, transactor bind.ContractTransactor) (*DeployerTransactor, error) {
	contract, err := bindDeployer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeployerTransactor{contract: contract}, nil
}

// NewDeployerFilterer creates a new log filterer instance of Deployer, bound to a specific deployed contract.
func NewDeployerFilterer(address common.Address, filterer bind.ContractFilterer) (*DeployerFilterer, error) {
	contract, err := bindDeployer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeployerFilterer{contract: contract}, nil
}

// bindDeployer binds a generic wrapper to an already deployed contract.
func bindDeployer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeployerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deployer *DeployerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deployer.Contract.DeployerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deployer *DeployerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deployer.Contract.DeployerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deployer *DeployerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deployer.Contract.DeployerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deployer *DeployerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deployer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deployer *DeployerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deployer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deployer *DeployerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deployer.Contract.contract.Transact(opts, method, params...)
}

// Deploy is a paid mutator transaction binding the contract method 0x775c300c.
//
// Solidity: function deploy() returns(address[5])
func (_Deployer *DeployerTransactor) Deploy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deployer.contract.Transact(opts, "deploy")
}

// Deploy is a paid mutator transaction binding the contract method 0x775c300c.
//
// Solidity: function deploy() returns(address[5])
func (_Deployer *DeployerSession) Deploy() (*types.Transaction, error) {
	return _Deployer.Contract.Deploy(&_Deployer.TransactOpts)
}

// Deploy is a paid mutator transaction binding the contract method 0x775c300c.
//
// Solidity: function deploy() returns(address[5])
func (_Deployer *DeployerTransactorSession) Deploy() (*types.Transaction, error) {
	return _Deployer.Contract.Deploy(&_Deployer.TransactOpts)
}

// IOABI is the input ABI used to generate the binding from.
const IOABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"read\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"readRevert\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"write\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"writeRevert\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// IOFuncSigs maps the 4-byte function signature to its string representation.
var IOFuncSigs = map[string]string{
	"61da1439": "read(bytes32)",
	"df8db75f": "readRevert(bytes32)",
	"e2e52ec1": "write(bytes32,bytes32)",
	"3405d6bf": "writeRevert(bytes32,bytes32)",
}

// IOBin is the compiled bytecode used for deploying new contracts.
var IOBin = "0x608060405234801561001057600080fd5b50610122806100206000396000f3fe608060405260043610603a5760003560e01c80633405d6bf14603f57806361da1439146061578063df8db75f14608d578063e2e52ec11460a7575b600080fd5b605f60048036036040811015605357600080fd5b508035906020013560c7565b005b607b60048036036020811015607557600080fd5b503560cf565b60408051918252519081900360200190f35b607b6004803603602081101560a157600080fd5b503560dc565b605f6004803603604081101560bb57600080fd5b508035906020013560e8565b808255600080fd5b6000815460005260206000f35b60008154600052600080fd5b905556fea2646970667358221220d1860a7c6440b5ce125750011e12a85bde2e9c3ae0d061c63cb8138c86c3cbdb64736f6c63430007000033"

// DeployIO deploys a new Ethereum contract, binding an instance of IO to it.
func DeployIO(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IO, error) {
	parsed, err := abi.JSON(strings.NewReader(IOABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IOBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IO{IOCaller: IOCaller{contract: contract}, IOTransactor: IOTransactor{contract: contract}, IOFilterer: IOFilterer{contract: contract}}, nil
}

// IO is an auto generated Go binding around an Ethereum contract.
type IO struct {
	IOCaller     // Read-only binding to the contract
	IOTransactor // Write-only binding to the contract
	IOFilterer   // Log filterer for contract events
}

// IOCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOSession struct {
	Contract     *IO               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOCallerSession struct {
	Contract *IOCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IOTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOTransactorSession struct {
	Contract     *IOTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IORaw is an auto generated low-level Go binding around an Ethereum contract.
type IORaw struct {
	Contract *IO // Generic contract binding to access the raw methods on
}

// IOCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOCallerRaw struct {
	Contract *IOCaller // Generic read-only contract binding to access the raw methods on
}

// IOTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOTransactorRaw struct {
	Contract *IOTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIO creates a new instance of IO, bound to a specific deployed contract.
func NewIO(address common.Address, backend bind.ContractBackend) (*IO, error) {
	contract, err := bindIO(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IO{IOCaller: IOCaller{contract: contract}, IOTransactor: IOTransactor{contract: contract}, IOFilterer: IOFilterer{contract: contract}}, nil
}

// NewIOCaller creates a new read-only instance of IO, bound to a specific deployed contract.
func NewIOCaller(address common.Address, caller bind.ContractCaller) (*IOCaller, error) {
	contract, err := bindIO(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOCaller{contract: contract}, nil
}

// NewIOTransactor creates a new write-only instance of IO, bound to a specific deployed contract.
func NewIOTransactor(address common.Address, transactor bind.ContractTransactor) (*IOTransactor, error) {
	contract, err := bindIO(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOTransactor{contract: contract}, nil
}

// NewIOFilterer creates a new log filterer instance of IO, bound to a specific deployed contract.
func NewIOFilterer(address common.Address, filterer bind.ContractFilterer) (*IOFilterer, error) {
	contract, err := bindIO(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOFilterer{contract: contract}, nil
}

// bindIO binds a generic wrapper to an already deployed contract.
func bindIO(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IO *IORaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IO.Contract.IOCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IO *IORaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IO.Contract.IOTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IO *IORaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IO.Contract.IOTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IO *IOCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IO.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IO *IOTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IO.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IO *IOTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IO.Contract.contract.Transact(opts, method, params...)
}

// Read is a paid mutator transaction binding the contract method 0x61da1439.
//
// Solidity: function read(bytes32 slot) payable returns(bytes32)
func (_IO *IOTransactor) Read(opts *bind.TransactOpts, slot [32]byte) (*types.Transaction, error) {
	return _IO.contract.Transact(opts, "read", slot)
}

// Read is a paid mutator transaction binding the contract method 0x61da1439.
//
// Solidity: function read(bytes32 slot) payable returns(bytes32)
func (_IO *IOSession) Read(slot [32]byte) (*types.Transaction, error) {
	return _IO.Contract.Read(&_IO.TransactOpts, slot)
}

// Read is a paid mutator transaction binding the contract method 0x61da1439.
//
// Solidity: function read(bytes32 slot) payable returns(bytes32)
func (_IO *IOTransactorSession) Read(slot [32]byte) (*types.Transaction, error) {
	return _IO.Contract.Read(&_IO.TransactOpts, slot)
}

// ReadRevert is a paid mutator transaction binding the contract method 0xdf8db75f.
//
// Solidity: function readRevert(bytes32 slot) payable returns(bytes32)
func (_IO *IOTransactor) ReadRevert(opts *bind.TransactOpts, slot [32]byte) (*types.Transaction, error) {
	return _IO.contract.Transact(opts, "readRevert", slot)
}

// ReadRevert is a paid mutator transaction binding the contract method 0xdf8db75f.
//
// Solidity: function readRevert(bytes32 slot) payable returns(bytes32)
func (_IO *IOSession) ReadRevert(slot [32]byte) (*types.Transaction, error) {
	return _IO.Contract.ReadRevert(&_IO.TransactOpts, slot)
}

// ReadRevert is a paid mutator transaction binding the contract method 0xdf8db75f.
//
// Solidity: function readRevert(bytes32 slot) payable returns(bytes32)
func (_IO *IOTransactorSession) ReadRevert(slot [32]byte) (*types.Transaction, error) {
	return _IO.Contract.ReadRevert(&_IO.TransactOpts, slot)
}

// Write is a paid mutator transaction binding the contract method 0xe2e52ec1.
//
// Solidity: function write(bytes32 slot, bytes32 value) payable returns()
func (_IO *IOTransactor) Write(opts *bind.TransactOpts, slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _IO.contract.Transact(opts, "write", slot, value)
}

// Write is a paid mutator transaction binding the contract method 0xe2e52ec1.
//
// Solidity: function write(bytes32 slot, bytes32 value) payable returns()
func (_IO *IOSession) Write(slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _IO.Contract.Write(&_IO.TransactOpts, slot, value)
}

// Write is a paid mutator transaction binding the contract method 0xe2e52ec1.
//
// Solidity: function write(bytes32 slot, bytes32 value) payable returns()
func (_IO *IOTransactorSession) Write(slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _IO.Contract.Write(&_IO.TransactOpts, slot, value)
}

// WriteRevert is a paid mutator transaction binding the contract method 0x3405d6bf.
//
// Solidity: function writeRevert(bytes32 slot, bytes32 value) payable returns()
func (_IO *IOTransactor) WriteRevert(opts *bind.TransactOpts, slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _IO.contract.Transact(opts, "writeRevert", slot, value)
}

// WriteRevert is a paid mutator transaction binding the contract method 0x3405d6bf.
//
// Solidity: function writeRevert(bytes32 slot, bytes32 value) payable returns()
func (_IO *IOSession) WriteRevert(slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _IO.Contract.WriteRevert(&_IO.TransactOpts, slot, value)
}

// WriteRevert is a paid mutator transaction binding the contract method 0x3405d6bf.
//
// Solidity: function writeRevert(bytes32 slot, bytes32 value) payable returns()
func (_IO *IOTransactorSession) WriteRevert(slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _IO.Contract.WriteRevert(&_IO.TransactOpts, slot, value)
}

// ScallProxyABI is the input ABI used to generate the binding from.
const ScallProxyABI = "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTarget\",\"type\":\"address\"}],\"name\":\"setTarget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ScallProxyFuncSigs maps the 4-byte function signature to its string representation.
var ScallProxyFuncSigs = map[string]string{
	"776d1a01": "setTarget(address)",
}

// ScallProxyBin is the compiled bytecode used for deploying new contracts.
var ScallProxyBin = "0x6080604052348015600f57600080fd5b5060be8061001e6000396000f3fe608060405260043610601c5760003560e01c8063776d1a01146035575b600080543682833781823684845afa50503d81823e3d81f35b348015604057600080fd5b50606460048036036020811015605557600080fd5b50356001600160a01b03166066565b005b600080546001600160a01b0319166001600160a01b039290921691909117905556fea2646970667358221220c3813fa3ef946ecb2fff63101d0408b82cf1c28b827ce6d67ce0685e73b8467764736f6c63430007000033"

// DeployScallProxy deploys a new Ethereum contract, binding an instance of ScallProxy to it.
func DeployScallProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ScallProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(ScallProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ScallProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ScallProxy{ScallProxyCaller: ScallProxyCaller{contract: contract}, ScallProxyTransactor: ScallProxyTransactor{contract: contract}, ScallProxyFilterer: ScallProxyFilterer{contract: contract}}, nil
}

// ScallProxy is an auto generated Go binding around an Ethereum contract.
type ScallProxy struct {
	ScallProxyCaller     // Read-only binding to the contract
	ScallProxyTransactor // Write-only binding to the contract
	ScallProxyFilterer   // Log filterer for contract events
}

// ScallProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ScallProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScallProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScallProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScallProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScallProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScallProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScallProxySession struct {
	Contract     *ScallProxy       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScallProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScallProxyCallerSession struct {
	Contract *ScallProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ScallProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScallProxyTransactorSession struct {
	Contract     *ScallProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ScallProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScallProxyRaw struct {
	Contract *ScallProxy // Generic contract binding to access the raw methods on
}

// ScallProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScallProxyCallerRaw struct {
	Contract *ScallProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ScallProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScallProxyTransactorRaw struct {
	Contract *ScallProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScallProxy creates a new instance of ScallProxy, bound to a specific deployed contract.
func NewScallProxy(address common.Address, backend bind.ContractBackend) (*ScallProxy, error) {
	contract, err := bindScallProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ScallProxy{ScallProxyCaller: ScallProxyCaller{contract: contract}, ScallProxyTransactor: ScallProxyTransactor{contract: contract}, ScallProxyFilterer: ScallProxyFilterer{contract: contract}}, nil
}

// NewScallProxyCaller creates a new read-only instance of ScallProxy, bound to a specific deployed contract.
func NewScallProxyCaller(address common.Address, caller bind.ContractCaller) (*ScallProxyCaller, error) {
	contract, err := bindScallProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScallProxyCaller{contract: contract}, nil
}

// NewScallProxyTransactor creates a new write-only instance of ScallProxy, bound to a specific deployed contract.
func NewScallProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ScallProxyTransactor, error) {
	contract, err := bindScallProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScallProxyTransactor{contract: contract}, nil
}

// NewScallProxyFilterer creates a new log filterer instance of ScallProxy, bound to a specific deployed contract.
func NewScallProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ScallProxyFilterer, error) {
	contract, err := bindScallProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScallProxyFilterer{contract: contract}, nil
}

// bindScallProxy binds a generic wrapper to an already deployed contract.
func bindScallProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ScallProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScallProxy *ScallProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScallProxy.Contract.ScallProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScallProxy *ScallProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScallProxy.Contract.ScallProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScallProxy *ScallProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScallProxy.Contract.ScallProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScallProxy *ScallProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScallProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScallProxy *ScallProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScallProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScallProxy *ScallProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScallProxy.Contract.contract.Transact(opts, method, params...)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address newTarget) returns()
func (_ScallProxy *ScallProxyTransactor) SetTarget(opts *bind.TransactOpts, newTarget common.Address) (*types.Transaction, error) {
	return _ScallProxy.contract.Transact(opts, "setTarget", newTarget)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address newTarget) returns()
func (_ScallProxy *ScallProxySession) SetTarget(newTarget common.Address) (*types.Transaction, error) {
	return _ScallProxy.Contract.SetTarget(&_ScallProxy.TransactOpts, newTarget)
}

// SetTarget is a paid mutator transaction binding the contract method 0x776d1a01.
//
// Solidity: function setTarget(address newTarget) returns()
func (_ScallProxy *ScallProxyTransactorSession) SetTarget(newTarget common.Address) (*types.Transaction, error) {
	return _ScallProxy.Contract.SetTarget(&_ScallProxy.TransactOpts, newTarget)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ScallProxy *ScallProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ScallProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ScallProxy *ScallProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ScallProxy.Contract.Fallback(&_ScallProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ScallProxy *ScallProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ScallProxy.Contract.Fallback(&_ScallProxy.TransactOpts, calldata)
}
