// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// WarmSStoreMetaData contains all meta data concerning the WarmSStore contract.
var WarmSStoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"runs\",\"type\":\"uint256\"}],\"name\":\"Sstore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"run\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"values\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040525f600155348015610013575f80fd5b5061027e806100215f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c80635e383d211461004357806364c421d114610073578063c04062261461008f575b5f80fd5b61005d6004803603810190610058919061014e565b6100ad565b60405161006a9190610188565b60405180910390f35b61008d6004803603810190610088919061014e565b6100c1565b005b610097610111565b6040516100a49190610188565b60405180910390f35b5f602052805f5260405f205f915090505481565b6001805f8282546100d291906101ce565b925050819055505f5b8181101561010d576001545f808381526020019081526020015f2081905550808061010590610201565b9150506100db565b5050565b60015481565b5f80fd5b5f819050919050565b61012d8161011b565b8114610137575f80fd5b50565b5f8135905061014881610124565b92915050565b5f6020828403121561016357610162610117565b5b5f6101708482850161013a565b91505092915050565b6101828161011b565b82525050565b5f60208201905061019b5f830184610179565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6101d88261011b565b91506101e38361011b565b92508282019050808211156101fb576101fa6101a1565b5b92915050565b5f61020b8261011b565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361023d5761023c6101a1565b5b60018201905091905056fea2646970667358221220c4b41bded80148284f70ae066a88523aececacc660354117571b02535bf41ffa64736f6c63430008140033",
}

// WarmSStoreABI is the input ABI used to generate the binding from.
// Deprecated: Use WarmSStoreMetaData.ABI instead.
var WarmSStoreABI = WarmSStoreMetaData.ABI

// WarmSStoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WarmSStoreMetaData.Bin instead.
var WarmSStoreBin = WarmSStoreMetaData.Bin

// DeployWarmSStore deploys a new Ethereum contract, binding an instance of WarmSStore to it.
func DeployWarmSStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WarmSStore, error) {
	parsed, err := WarmSStoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WarmSStoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WarmSStore{WarmSStoreCaller: WarmSStoreCaller{contract: contract}, WarmSStoreTransactor: WarmSStoreTransactor{contract: contract}, WarmSStoreFilterer: WarmSStoreFilterer{contract: contract}}, nil
}

// WarmSStore is an auto generated Go binding around an Ethereum contract.
type WarmSStore struct {
	WarmSStoreCaller     // Read-only binding to the contract
	WarmSStoreTransactor // Write-only binding to the contract
	WarmSStoreFilterer   // Log filterer for contract events
}

// WarmSStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type WarmSStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WarmSStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WarmSStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WarmSStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WarmSStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WarmSStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WarmSStoreSession struct {
	Contract     *WarmSStore       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WarmSStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WarmSStoreCallerSession struct {
	Contract *WarmSStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// WarmSStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WarmSStoreTransactorSession struct {
	Contract     *WarmSStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// WarmSStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type WarmSStoreRaw struct {
	Contract *WarmSStore // Generic contract binding to access the raw methods on
}

// WarmSStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WarmSStoreCallerRaw struct {
	Contract *WarmSStoreCaller // Generic read-only contract binding to access the raw methods on
}

// WarmSStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WarmSStoreTransactorRaw struct {
	Contract *WarmSStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWarmSStore creates a new instance of WarmSStore, bound to a specific deployed contract.
func NewWarmSStore(address common.Address, backend bind.ContractBackend) (*WarmSStore, error) {
	contract, err := bindWarmSStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WarmSStore{WarmSStoreCaller: WarmSStoreCaller{contract: contract}, WarmSStoreTransactor: WarmSStoreTransactor{contract: contract}, WarmSStoreFilterer: WarmSStoreFilterer{contract: contract}}, nil
}

// NewWarmSStoreCaller creates a new read-only instance of WarmSStore, bound to a specific deployed contract.
func NewWarmSStoreCaller(address common.Address, caller bind.ContractCaller) (*WarmSStoreCaller, error) {
	contract, err := bindWarmSStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WarmSStoreCaller{contract: contract}, nil
}

// NewWarmSStoreTransactor creates a new write-only instance of WarmSStore, bound to a specific deployed contract.
func NewWarmSStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*WarmSStoreTransactor, error) {
	contract, err := bindWarmSStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WarmSStoreTransactor{contract: contract}, nil
}

// NewWarmSStoreFilterer creates a new log filterer instance of WarmSStore, bound to a specific deployed contract.
func NewWarmSStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*WarmSStoreFilterer, error) {
	contract, err := bindWarmSStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WarmSStoreFilterer{contract: contract}, nil
}

// bindWarmSStore binds a generic wrapper to an already deployed contract.
func bindWarmSStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WarmSStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WarmSStore *WarmSStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WarmSStore.Contract.WarmSStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WarmSStore *WarmSStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WarmSStore.Contract.WarmSStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WarmSStore *WarmSStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WarmSStore.Contract.WarmSStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WarmSStore *WarmSStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WarmSStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WarmSStore *WarmSStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WarmSStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WarmSStore *WarmSStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WarmSStore.Contract.contract.Transact(opts, method, params...)
}

// Run is a free data retrieval call binding the contract method 0xc0406226.
//
// Solidity: function run() view returns(uint256)
func (_WarmSStore *WarmSStoreCaller) Run(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WarmSStore.contract.Call(opts, &out, "run")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Run is a free data retrieval call binding the contract method 0xc0406226.
//
// Solidity: function run() view returns(uint256)
func (_WarmSStore *WarmSStoreSession) Run() (*big.Int, error) {
	return _WarmSStore.Contract.Run(&_WarmSStore.CallOpts)
}

// Run is a free data retrieval call binding the contract method 0xc0406226.
//
// Solidity: function run() view returns(uint256)
func (_WarmSStore *WarmSStoreCallerSession) Run() (*big.Int, error) {
	return _WarmSStore.Contract.Run(&_WarmSStore.CallOpts)
}

// Values is a free data retrieval call binding the contract method 0x5e383d21.
//
// Solidity: function values(uint256 ) view returns(uint256)
func (_WarmSStore *WarmSStoreCaller) Values(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WarmSStore.contract.Call(opts, &out, "values", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Values is a free data retrieval call binding the contract method 0x5e383d21.
//
// Solidity: function values(uint256 ) view returns(uint256)
func (_WarmSStore *WarmSStoreSession) Values(arg0 *big.Int) (*big.Int, error) {
	return _WarmSStore.Contract.Values(&_WarmSStore.CallOpts, arg0)
}

// Values is a free data retrieval call binding the contract method 0x5e383d21.
//
// Solidity: function values(uint256 ) view returns(uint256)
func (_WarmSStore *WarmSStoreCallerSession) Values(arg0 *big.Int) (*big.Int, error) {
	return _WarmSStore.Contract.Values(&_WarmSStore.CallOpts, arg0)
}

// Sstore is a paid mutator transaction binding the contract method 0x64c421d1.
//
// Solidity: function Sstore(uint256 runs) returns()
func (_WarmSStore *WarmSStoreTransactor) Sstore(opts *bind.TransactOpts, runs *big.Int) (*types.Transaction, error) {
	return _WarmSStore.contract.Transact(opts, "Sstore", runs)
}

// Sstore is a paid mutator transaction binding the contract method 0x64c421d1.
//
// Solidity: function Sstore(uint256 runs) returns()
func (_WarmSStore *WarmSStoreSession) Sstore(runs *big.Int) (*types.Transaction, error) {
	return _WarmSStore.Contract.Sstore(&_WarmSStore.TransactOpts, runs)
}

// Sstore is a paid mutator transaction binding the contract method 0x64c421d1.
//
// Solidity: function Sstore(uint256 runs) returns()
func (_WarmSStore *WarmSStoreTransactorSession) Sstore(runs *big.Int) (*types.Transaction, error) {
	return _WarmSStore.Contract.Sstore(&_WarmSStore.TransactOpts, runs)
}
