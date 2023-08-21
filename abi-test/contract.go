// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package geth

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

// BConstABI is the input ABI used to generate the binding from.
const BConstABI = "[{\"inputs\":[],\"name\":\"BONE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BPOW_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_TOTAL_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_POOL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_IN_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_OUT_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TOTAL_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION_NUMBER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WEIGHT_CHANGE_PCT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WEIGHT_UPDATE_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BConstFuncSigs maps the 4-byte function signature to its string representation.
var BConstFuncSigs = map[string]string{
	"c36596a6": "BONE()",
	"189d00ca": "BPOW_PRECISION()",
	"ed61bdc0": "DEFAULT_TOTAL_WEIGHT()",
	"c6580d12": "EXIT_FEE()",
	"9381cd2b": "INIT_POOL_SUPPLY()",
	"b0e0d136": "MAX_BOUND_TOKENS()",
	"bc694ea2": "MAX_BPOW_BASE()",
	"bc063e1a": "MAX_FEE()",
	"ec093021": "MAX_IN_RATIO()",
	"992e2a92": "MAX_OUT_RATIO()",
	"09a3bbe4": "MAX_TOTAL_WEIGHT()",
	"e4a28a52": "MAX_WEIGHT()",
	"867378c5": "MIN_BALANCE()",
	"b7b800a4": "MIN_BOUND_TOKENS()",
	"ba019dab": "MIN_BPOW_BASE()",
	"76c7a3c7": "MIN_FEE()",
	"218b5382": "MIN_WEIGHT()",
	"8025e303": "VERSION_NUMBER()",
	"6852b5cd": "WEIGHT_CHANGE_PCT()",
	"7e6d71a2": "WEIGHT_UPDATE_DELAY()",
}

// BConstBin is the compiled bytecode used for deploying new contracts.
var BConstBin = "0x608060405234801561001057600080fd5b506102e8806100206000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c8063b0e0d136116100ad578063c36596a611610071578063c36596a6146101b3578063c6580d12146101bb578063e4a28a52146101c3578063ec093021146101cb578063ed61bdc0146101c35761012c565b8063b0e0d13614610193578063b7b800a41461019b578063ba019dab14610173578063bc063e1a146101a3578063bc694ea2146101ab5761012c565b80637e6d71a2116100f45780637e6d71a21461016b5780638025e30314610173578063867378c51461017b5780639381cd2b14610183578063992e2a921461018b5761012c565b806309a3bbe414610131578063189d00ca1461014b578063218b5382146101535780636852b5cd1461015b57806376c7a3c714610163575b600080fd5b6101396101d3565b60408051918252519081900360200190f35b6101396101e0565b6101396101f4565b610139610204565b610139610214565b610139610226565b61013961022c565b610139610231565b610139610245565b610139610252565b61013961025e565b610139610263565b610139610268565b610139610278565b610139610284565b610139610290565b610139610295565b6101396102a2565b680168d28e3f0028000081565b6402540be400670de0b6b3a76400005b0481565b6004670de0b6b3a76400006101f0565b6064670de0b6b3a76400006101f0565b620f4240670de0b6b3a76400006101f0565b610e1081565b600181565b64e8d4a51000670de0b6b3a76400006101f0565b68056bc75e2d6310000081565b6704a03ce68d21555681565b600a81565b600281565b600a670de0b6b3a76400006101f0565b671bc16d674ec7ffff81565b670de0b6b3a764000081565b600081565b68015af1d78b58c4000081565b6002670de0b6b3a76400006101f056fea26469706673582212208206c21e240beb9fa2ff55e2bfa462087e2a0804926ec681c0960e99a0d368ad64736f6c63430007000033"

// DeployBConst deploys a new Ethereum contract, binding an instance of BConst to it.
func DeployBConst(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BConst, error) {
	parsed, err := abi.JSON(strings.NewReader(BConstABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BConstBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BConst{BConstCaller: BConstCaller{contract: contract}, BConstTransactor: BConstTransactor{contract: contract}, BConstFilterer: BConstFilterer{contract: contract}}, nil
}

// BConst is an auto generated Go binding around an Ethereum contract.
type BConst struct {
	BConstCaller     // Read-only binding to the contract
	BConstTransactor // Write-only binding to the contract
	BConstFilterer   // Log filterer for contract events
}

// BConstCaller is an auto generated read-only Go binding around an Ethereum contract.
type BConstCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BConstTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BConstTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BConstFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BConstFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BConstSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BConstSession struct {
	Contract     *BConst           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BConstCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BConstCallerSession struct {
	Contract *BConstCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BConstTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BConstTransactorSession struct {
	Contract     *BConstTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BConstRaw is an auto generated low-level Go binding around an Ethereum contract.
type BConstRaw struct {
	Contract *BConst // Generic contract binding to access the raw methods on
}

// BConstCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BConstCallerRaw struct {
	Contract *BConstCaller // Generic read-only contract binding to access the raw methods on
}

// BConstTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BConstTransactorRaw struct {
	Contract *BConstTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBConst creates a new instance of BConst, bound to a specific deployed contract.
func NewBConst(address common.Address, backend bind.ContractBackend) (*BConst, error) {
	contract, err := bindBConst(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BConst{BConstCaller: BConstCaller{contract: contract}, BConstTransactor: BConstTransactor{contract: contract}, BConstFilterer: BConstFilterer{contract: contract}}, nil
}

// NewBConstCaller creates a new read-only instance of BConst, bound to a specific deployed contract.
func NewBConstCaller(address common.Address, caller bind.ContractCaller) (*BConstCaller, error) {
	contract, err := bindBConst(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BConstCaller{contract: contract}, nil
}

// NewBConstTransactor creates a new write-only instance of BConst, bound to a specific deployed contract.
func NewBConstTransactor(address common.Address, transactor bind.ContractTransactor) (*BConstTransactor, error) {
	contract, err := bindBConst(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BConstTransactor{contract: contract}, nil
}

// NewBConstFilterer creates a new log filterer instance of BConst, bound to a specific deployed contract.
func NewBConstFilterer(address common.Address, filterer bind.ContractFilterer) (*BConstFilterer, error) {
	contract, err := bindBConst(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BConstFilterer{contract: contract}, nil
}

// bindBConst binds a generic wrapper to an already deployed contract.
func bindBConst(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BConstABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BConst *BConstRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BConst.Contract.BConstCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BConst *BConstRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BConst.Contract.BConstTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BConst *BConstRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BConst.Contract.BConstTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BConst *BConstCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BConst.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BConst *BConstTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BConst.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BConst *BConstTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BConst.Contract.contract.Transact(opts, method, params...)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BConst *BConstCaller) BONE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "BONE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BConst *BConstSession) BONE() (*big.Int, error) {
	return _BConst.Contract.BONE(&_BConst.CallOpts)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BConst *BConstCallerSession) BONE() (*big.Int, error) {
	return _BConst.Contract.BONE(&_BConst.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BConst *BConstCaller) BPOWPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "BPOW_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BConst *BConstSession) BPOWPRECISION() (*big.Int, error) {
	return _BConst.Contract.BPOWPRECISION(&_BConst.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BConst *BConstCallerSession) BPOWPRECISION() (*big.Int, error) {
	return _BConst.Contract.BPOWPRECISION(&_BConst.CallOpts)
}

// DEFAULTTOTALWEIGHT is a free data retrieval call binding the contract method 0xed61bdc0.
//
// Solidity: function DEFAULT_TOTAL_WEIGHT() view returns(uint256)
func (_BConst *BConstCaller) DEFAULTTOTALWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "DEFAULT_TOTAL_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTTOTALWEIGHT is a free data retrieval call binding the contract method 0xed61bdc0.
//
// Solidity: function DEFAULT_TOTAL_WEIGHT() view returns(uint256)
func (_BConst *BConstSession) DEFAULTTOTALWEIGHT() (*big.Int, error) {
	return _BConst.Contract.DEFAULTTOTALWEIGHT(&_BConst.CallOpts)
}

// DEFAULTTOTALWEIGHT is a free data retrieval call binding the contract method 0xed61bdc0.
//
// Solidity: function DEFAULT_TOTAL_WEIGHT() view returns(uint256)
func (_BConst *BConstCallerSession) DEFAULTTOTALWEIGHT() (*big.Int, error) {
	return _BConst.Contract.DEFAULTTOTALWEIGHT(&_BConst.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BConst *BConstCaller) EXITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "EXIT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BConst *BConstSession) EXITFEE() (*big.Int, error) {
	return _BConst.Contract.EXITFEE(&_BConst.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BConst *BConstCallerSession) EXITFEE() (*big.Int, error) {
	return _BConst.Contract.EXITFEE(&_BConst.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BConst *BConstCaller) INITPOOLSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "INIT_POOL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BConst *BConstSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _BConst.Contract.INITPOOLSUPPLY(&_BConst.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BConst *BConstCallerSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _BConst.Contract.INITPOOLSUPPLY(&_BConst.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BConst *BConstCaller) MAXBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MAX_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BConst *BConstSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _BConst.Contract.MAXBOUNDTOKENS(&_BConst.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BConst *BConstCallerSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _BConst.Contract.MAXBOUNDTOKENS(&_BConst.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BConst *BConstCaller) MAXBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MAX_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BConst *BConstSession) MAXBPOWBASE() (*big.Int, error) {
	return _BConst.Contract.MAXBPOWBASE(&_BConst.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BConst *BConstCallerSession) MAXBPOWBASE() (*big.Int, error) {
	return _BConst.Contract.MAXBPOWBASE(&_BConst.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BConst *BConstCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BConst *BConstSession) MAXFEE() (*big.Int, error) {
	return _BConst.Contract.MAXFEE(&_BConst.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BConst *BConstCallerSession) MAXFEE() (*big.Int, error) {
	return _BConst.Contract.MAXFEE(&_BConst.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BConst *BConstCaller) MAXINRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MAX_IN_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BConst *BConstSession) MAXINRATIO() (*big.Int, error) {
	return _BConst.Contract.MAXINRATIO(&_BConst.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BConst *BConstCallerSession) MAXINRATIO() (*big.Int, error) {
	return _BConst.Contract.MAXINRATIO(&_BConst.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BConst *BConstCaller) MAXOUTRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MAX_OUT_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BConst *BConstSession) MAXOUTRATIO() (*big.Int, error) {
	return _BConst.Contract.MAXOUTRATIO(&_BConst.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BConst *BConstCallerSession) MAXOUTRATIO() (*big.Int, error) {
	return _BConst.Contract.MAXOUTRATIO(&_BConst.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BConst *BConstCaller) MAXTOTALWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MAX_TOTAL_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BConst *BConstSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _BConst.Contract.MAXTOTALWEIGHT(&_BConst.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BConst *BConstCallerSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _BConst.Contract.MAXTOTALWEIGHT(&_BConst.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BConst *BConstCaller) MAXWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MAX_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BConst *BConstSession) MAXWEIGHT() (*big.Int, error) {
	return _BConst.Contract.MAXWEIGHT(&_BConst.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BConst *BConstCallerSession) MAXWEIGHT() (*big.Int, error) {
	return _BConst.Contract.MAXWEIGHT(&_BConst.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BConst *BConstCaller) MINBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MIN_BALANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BConst *BConstSession) MINBALANCE() (*big.Int, error) {
	return _BConst.Contract.MINBALANCE(&_BConst.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BConst *BConstCallerSession) MINBALANCE() (*big.Int, error) {
	return _BConst.Contract.MINBALANCE(&_BConst.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BConst *BConstCaller) MINBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MIN_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BConst *BConstSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _BConst.Contract.MINBOUNDTOKENS(&_BConst.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BConst *BConstCallerSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _BConst.Contract.MINBOUNDTOKENS(&_BConst.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BConst *BConstCaller) MINBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MIN_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BConst *BConstSession) MINBPOWBASE() (*big.Int, error) {
	return _BConst.Contract.MINBPOWBASE(&_BConst.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BConst *BConstCallerSession) MINBPOWBASE() (*big.Int, error) {
	return _BConst.Contract.MINBPOWBASE(&_BConst.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BConst *BConstCaller) MINFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MIN_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BConst *BConstSession) MINFEE() (*big.Int, error) {
	return _BConst.Contract.MINFEE(&_BConst.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BConst *BConstCallerSession) MINFEE() (*big.Int, error) {
	return _BConst.Contract.MINFEE(&_BConst.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BConst *BConstCaller) MINWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "MIN_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BConst *BConstSession) MINWEIGHT() (*big.Int, error) {
	return _BConst.Contract.MINWEIGHT(&_BConst.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BConst *BConstCallerSession) MINWEIGHT() (*big.Int, error) {
	return _BConst.Contract.MINWEIGHT(&_BConst.CallOpts)
}

// VERSIONNUMBER is a free data retrieval call binding the contract method 0x8025e303.
//
// Solidity: function VERSION_NUMBER() view returns(uint256)
func (_BConst *BConstCaller) VERSIONNUMBER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "VERSION_NUMBER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERSIONNUMBER is a free data retrieval call binding the contract method 0x8025e303.
//
// Solidity: function VERSION_NUMBER() view returns(uint256)
func (_BConst *BConstSession) VERSIONNUMBER() (*big.Int, error) {
	return _BConst.Contract.VERSIONNUMBER(&_BConst.CallOpts)
}

// VERSIONNUMBER is a free data retrieval call binding the contract method 0x8025e303.
//
// Solidity: function VERSION_NUMBER() view returns(uint256)
func (_BConst *BConstCallerSession) VERSIONNUMBER() (*big.Int, error) {
	return _BConst.Contract.VERSIONNUMBER(&_BConst.CallOpts)
}

// WEIGHTCHANGEPCT is a free data retrieval call binding the contract method 0x6852b5cd.
//
// Solidity: function WEIGHT_CHANGE_PCT() view returns(uint256)
func (_BConst *BConstCaller) WEIGHTCHANGEPCT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "WEIGHT_CHANGE_PCT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WEIGHTCHANGEPCT is a free data retrieval call binding the contract method 0x6852b5cd.
//
// Solidity: function WEIGHT_CHANGE_PCT() view returns(uint256)
func (_BConst *BConstSession) WEIGHTCHANGEPCT() (*big.Int, error) {
	return _BConst.Contract.WEIGHTCHANGEPCT(&_BConst.CallOpts)
}

// WEIGHTCHANGEPCT is a free data retrieval call binding the contract method 0x6852b5cd.
//
// Solidity: function WEIGHT_CHANGE_PCT() view returns(uint256)
func (_BConst *BConstCallerSession) WEIGHTCHANGEPCT() (*big.Int, error) {
	return _BConst.Contract.WEIGHTCHANGEPCT(&_BConst.CallOpts)
}

// WEIGHTUPDATEDELAY is a free data retrieval call binding the contract method 0x7e6d71a2.
//
// Solidity: function WEIGHT_UPDATE_DELAY() view returns(uint256)
func (_BConst *BConstCaller) WEIGHTUPDATEDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BConst.contract.Call(opts, &out, "WEIGHT_UPDATE_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WEIGHTUPDATEDELAY is a free data retrieval call binding the contract method 0x7e6d71a2.
//
// Solidity: function WEIGHT_UPDATE_DELAY() view returns(uint256)
func (_BConst *BConstSession) WEIGHTUPDATEDELAY() (*big.Int, error) {
	return _BConst.Contract.WEIGHTUPDATEDELAY(&_BConst.CallOpts)
}

// WEIGHTUPDATEDELAY is a free data retrieval call binding the contract method 0x7e6d71a2.
//
// Solidity: function WEIGHT_UPDATE_DELAY() view returns(uint256)
func (_BConst *BConstCallerSession) WEIGHTUPDATEDELAY() (*big.Int, error) {
	return _BConst.Contract.WEIGHTUPDATEDELAY(&_BConst.CallOpts)
}

// BNumABI is the input ABI used to generate the binding from.
const BNumABI = "[{\"inputs\":[],\"name\":\"BONE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BPOW_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_TOTAL_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_POOL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_IN_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_OUT_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TOTAL_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION_NUMBER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WEIGHT_CHANGE_PCT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WEIGHT_UPDATE_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"badd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"bdiv\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"bfloor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"bmul\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"base\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"}],\"name\":\"bpow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"base\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"precision\",\"type\":\"uint256\"}],\"name\":\"bpowApprox\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"bpowi\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"bsub\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"bsubSign\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"btoi\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// BNumFuncSigs maps the 4-byte function signature to its string representation.
var BNumFuncSigs = map[string]string{
	"c36596a6": "BONE()",
	"189d00ca": "BPOW_PRECISION()",
	"ed61bdc0": "DEFAULT_TOTAL_WEIGHT()",
	"c6580d12": "EXIT_FEE()",
	"9381cd2b": "INIT_POOL_SUPPLY()",
	"b0e0d136": "MAX_BOUND_TOKENS()",
	"bc694ea2": "MAX_BPOW_BASE()",
	"bc063e1a": "MAX_FEE()",
	"ec093021": "MAX_IN_RATIO()",
	"992e2a92": "MAX_OUT_RATIO()",
	"09a3bbe4": "MAX_TOTAL_WEIGHT()",
	"e4a28a52": "MAX_WEIGHT()",
	"867378c5": "MIN_BALANCE()",
	"b7b800a4": "MIN_BOUND_TOKENS()",
	"ba019dab": "MIN_BPOW_BASE()",
	"76c7a3c7": "MIN_FEE()",
	"218b5382": "MIN_WEIGHT()",
	"8025e303": "VERSION_NUMBER()",
	"6852b5cd": "WEIGHT_CHANGE_PCT()",
	"7e6d71a2": "WEIGHT_UPDATE_DELAY()",
	"0b71b95c": "badd(uint256,uint256)",
	"7673eb11": "bdiv(uint256,uint256)",
	"317223bc": "bfloor(uint256)",
	"8c051bf3": "bmul(uint256,uint256)",
	"1d0a82df": "bpow(uint256,uint256)",
	"d011f32a": "bpowApprox(uint256,uint256,uint256)",
	"cd660299": "bpowi(uint256,uint256)",
	"3cf3c7d4": "bsub(uint256,uint256)",
	"743480cc": "bsubSign(uint256,uint256)",
	"859ae48e": "btoi(uint256)",
}

// BNumBin is the compiled bytecode used for deploying new contracts.
var BNumBin = "0x608060405234801561001057600080fd5b50610a2a806100206000396000f3fe608060405234801561001057600080fd5b50600436106101da5760003560e01c80638c051bf311610104578063bc694ea2116100a2578063d011f32a11610071578063d011f32a146103b9578063e4a28a52146103e2578063ec093021146103ea578063ed61bdc0146103e2576101da565b8063bc694ea21461037e578063c36596a614610386578063c6580d121461038e578063cd66029914610396576101da565b8063b0e0d136116100de578063b0e0d13614610366578063b7b800a41461036e578063ba019dab14610306578063bc063e1a14610376576101da565b80638c051bf3146103335780639381cd2b14610356578063992e2a921461035e576101da565b80636852b5cd1161017c5780637e6d71a21161014b5780637e6d71a2146102fe5780638025e30314610306578063859ae48e1461030e578063867378c51461032b576101da565b80636852b5cd1461028f578063743480cc146102975780637673eb11146102d357806376c7a3c7146102f6576101da565b80631d0a82df116101b85780631d0a82df14610224578063218b538214610247578063317223bc1461024f5780633cf3c7d41461026c576101da565b806309a3bbe4146101df5780630b71b95c146101f9578063189d00ca1461021c575b600080fd5b6101e76103f2565b60408051918252519081900360200190f35b6101e76004803603604081101561020f57600080fd5b50803590602001356103ff565b6101e7610455565b6101e76004803603604081101561023a57600080fd5b5080359060200135610469565b6101e7610577565b6101e76004803603602081101561026557600080fd5b5035610587565b6101e76004803603604081101561028257600080fd5b50803590602001356105a2565b6101e7610604565b6102ba600480360360408110156102ad57600080fd5b5080359060200135610614565b6040805192835290151560208301528051918290030190f35b6101e7600480360360408110156102e957600080fd5b5080359060200135610639565b6101e761074c565b6101e761075e565b6101e7610764565b6101e76004803603602081101561032457600080fd5b5035610769565b6101e7610777565b6101e76004803603604081101561034957600080fd5b508035906020013561078b565b6101e761084d565b6101e761085a565b6101e7610866565b6101e761086b565b6101e7610870565b6101e7610880565b6101e761088c565b6101e7610898565b6101e7600480360360408110156103ac57600080fd5b508035906020013561089d565b6101e7600480360360608110156103cf57600080fd5b50803590602081013590604001356108f4565b6101e76109d7565b6101e76109e4565b680168d28e3f0028000081565b60008282018381101561044c576040805162461bcd60e51b815260206004820152601060248201526f4552525f4144445f4f564552464c4f5760801b604482015290519081900360640190fd5b90505b92915050565b6402540be400670de0b6b3a76400005b0481565b600060018310156104b9576040805162461bcd60e51b81526020600482015260156024820152744552525f42504f575f424153455f544f4f5f4c4f5760581b604482015290519081900360640190fd5b671bc16d674ec7ffff83111561050f576040805162461bcd60e51b815260206004820152601660248201527508aa4a4be84a09eaebe8482a68abea89e9ebe90928e960531b604482015290519081900360640190fd5b600061051a83610587565b9050600061052884836105a2565b9050600061053e8661053985610769565b61089d565b90508161054f57925061044f915050565b600061056087846305f5e1006108f4565b905061056c828261078b565b979650505050505050565b6004670de0b6b3a7640000610465565b6000670de0b6b3a764000061059b83610769565b0292915050565b60008060006105b18585610614565b9150915080156105fc576040805162461bcd60e51b81526020600482015260116024820152704552525f5355425f554e444552464c4f5760781b604482015290519081900360640190fd5b509392505050565b6064670de0b6b3a7640000610465565b60008082841061062a5750508082036000610632565b505081810360015b9250929050565b60008161067c576040805162461bcd60e51b815260206004820152600c60248201526b4552525f4449565f5a45524f60a01b604482015290519081900360640190fd5b670de0b6b3a764000083028315806106a45750670de0b6b3a76400008482816106a157fe5b04145b6106e8576040805162461bcd60e51b815260206004820152601060248201526f11549497d1125597d25395115493905360821b604482015290519081900360640190fd5b60028304810181811015610736576040805162461bcd60e51b815260206004820152601060248201526f11549497d1125597d25395115493905360821b604482015290519081900360640190fd5b600084828161074157fe5b049695505050505050565b620f4240670de0b6b3a7640000610465565b610e1081565b600181565b670de0b6b3a7640000900490565b64e8d4a51000670de0b6b3a7640000610465565b60008282028315806107a55750828482816107a257fe5b04145b6107e9576040805162461bcd60e51b815260206004820152601060248201526f4552525f4d554c5f4f564552464c4f5760801b604482015290519081900360640190fd5b6706f05b59d3b2000081018181101561083c576040805162461bcd60e51b815260206004820152601060248201526f4552525f4d554c5f4f564552464c4f5760801b604482015290519081900360640190fd5b6000670de0b6b3a764000082610741565b68056bc75e2d6310000081565b6704a03ce68d21555681565b600a81565b600281565b600a670de0b6b3a7640000610465565b671bc16d674ec7ffff81565b670de0b6b3a764000081565b600081565b600080600283066108b657670de0b6b3a76400006108b8565b835b90506002830492505b821561044c576108d1848561078b565b935060028306156108e9576108e6818561078b565b90505b6002830492506108c1565b600082818061090b87670de0b6b3a7640000610614565b9092509050670de0b6b3a764000080600060015b8884106109c8576000670de0b6b3a7640000820290506000806109538a61094e85670de0b6b3a76400006105a2565b610614565b9150915061096a87610965848c61078b565b61078b565b96506109768784610639565b965086610985575050506109c8565b871561098f579315935b8015610999579315935b84156109b0576109a986886105a2565b95506109bd565b6109ba86886103ff565b95505b50505060010161091f565b50909998505050505050505050565b68015af1d78b58c4000081565b6002670de0b6b3a764000061046556fea26469706673582212209b95ce8789dbd202639cbec9da08755a32e78a4275bce0621da2b201c621471364736f6c63430007000033"

// DeployBNum deploys a new Ethereum contract, binding an instance of BNum to it.
func DeployBNum(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BNum, error) {
	parsed, err := abi.JSON(strings.NewReader(BNumABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BNumBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BNum{BNumCaller: BNumCaller{contract: contract}, BNumTransactor: BNumTransactor{contract: contract}, BNumFilterer: BNumFilterer{contract: contract}}, nil
}

// BNum is an auto generated Go binding around an Ethereum contract.
type BNum struct {
	BNumCaller     // Read-only binding to the contract
	BNumTransactor // Write-only binding to the contract
	BNumFilterer   // Log filterer for contract events
}

// BNumCaller is an auto generated read-only Go binding around an Ethereum contract.
type BNumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BNumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BNumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BNumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BNumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BNumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BNumSession struct {
	Contract     *BNum             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BNumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BNumCallerSession struct {
	Contract *BNumCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BNumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BNumTransactorSession struct {
	Contract     *BNumTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BNumRaw is an auto generated low-level Go binding around an Ethereum contract.
type BNumRaw struct {
	Contract *BNum // Generic contract binding to access the raw methods on
}

// BNumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BNumCallerRaw struct {
	Contract *BNumCaller // Generic read-only contract binding to access the raw methods on
}

// BNumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BNumTransactorRaw struct {
	Contract *BNumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBNum creates a new instance of BNum, bound to a specific deployed contract.
func NewBNum(address common.Address, backend bind.ContractBackend) (*BNum, error) {
	contract, err := bindBNum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BNum{BNumCaller: BNumCaller{contract: contract}, BNumTransactor: BNumTransactor{contract: contract}, BNumFilterer: BNumFilterer{contract: contract}}, nil
}

// NewBNumCaller creates a new read-only instance of BNum, bound to a specific deployed contract.
func NewBNumCaller(address common.Address, caller bind.ContractCaller) (*BNumCaller, error) {
	contract, err := bindBNum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BNumCaller{contract: contract}, nil
}

// NewBNumTransactor creates a new write-only instance of BNum, bound to a specific deployed contract.
func NewBNumTransactor(address common.Address, transactor bind.ContractTransactor) (*BNumTransactor, error) {
	contract, err := bindBNum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BNumTransactor{contract: contract}, nil
}

// NewBNumFilterer creates a new log filterer instance of BNum, bound to a specific deployed contract.
func NewBNumFilterer(address common.Address, filterer bind.ContractFilterer) (*BNumFilterer, error) {
	contract, err := bindBNum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BNumFilterer{contract: contract}, nil
}

// bindBNum binds a generic wrapper to an already deployed contract.
func bindBNum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BNumABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BNum *BNumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BNum.Contract.BNumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BNum *BNumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BNum.Contract.BNumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BNum *BNumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BNum.Contract.BNumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BNum *BNumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BNum.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BNum *BNumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BNum.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BNum *BNumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BNum.Contract.contract.Transact(opts, method, params...)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BNum *BNumCaller) BONE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "BONE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BNum *BNumSession) BONE() (*big.Int, error) {
	return _BNum.Contract.BONE(&_BNum.CallOpts)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BNum *BNumCallerSession) BONE() (*big.Int, error) {
	return _BNum.Contract.BONE(&_BNum.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BNum *BNumCaller) BPOWPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "BPOW_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BNum *BNumSession) BPOWPRECISION() (*big.Int, error) {
	return _BNum.Contract.BPOWPRECISION(&_BNum.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BNum *BNumCallerSession) BPOWPRECISION() (*big.Int, error) {
	return _BNum.Contract.BPOWPRECISION(&_BNum.CallOpts)
}

// DEFAULTTOTALWEIGHT is a free data retrieval call binding the contract method 0xed61bdc0.
//
// Solidity: function DEFAULT_TOTAL_WEIGHT() view returns(uint256)
func (_BNum *BNumCaller) DEFAULTTOTALWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "DEFAULT_TOTAL_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTTOTALWEIGHT is a free data retrieval call binding the contract method 0xed61bdc0.
//
// Solidity: function DEFAULT_TOTAL_WEIGHT() view returns(uint256)
func (_BNum *BNumSession) DEFAULTTOTALWEIGHT() (*big.Int, error) {
	return _BNum.Contract.DEFAULTTOTALWEIGHT(&_BNum.CallOpts)
}

// DEFAULTTOTALWEIGHT is a free data retrieval call binding the contract method 0xed61bdc0.
//
// Solidity: function DEFAULT_TOTAL_WEIGHT() view returns(uint256)
func (_BNum *BNumCallerSession) DEFAULTTOTALWEIGHT() (*big.Int, error) {
	return _BNum.Contract.DEFAULTTOTALWEIGHT(&_BNum.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BNum *BNumCaller) EXITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "EXIT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BNum *BNumSession) EXITFEE() (*big.Int, error) {
	return _BNum.Contract.EXITFEE(&_BNum.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BNum *BNumCallerSession) EXITFEE() (*big.Int, error) {
	return _BNum.Contract.EXITFEE(&_BNum.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BNum *BNumCaller) INITPOOLSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "INIT_POOL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BNum *BNumSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _BNum.Contract.INITPOOLSUPPLY(&_BNum.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BNum *BNumCallerSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _BNum.Contract.INITPOOLSUPPLY(&_BNum.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BNum *BNumCaller) MAXBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MAX_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BNum *BNumSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _BNum.Contract.MAXBOUNDTOKENS(&_BNum.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BNum *BNumCallerSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _BNum.Contract.MAXBOUNDTOKENS(&_BNum.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BNum *BNumCaller) MAXBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MAX_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BNum *BNumSession) MAXBPOWBASE() (*big.Int, error) {
	return _BNum.Contract.MAXBPOWBASE(&_BNum.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BNum *BNumCallerSession) MAXBPOWBASE() (*big.Int, error) {
	return _BNum.Contract.MAXBPOWBASE(&_BNum.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BNum *BNumCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BNum *BNumSession) MAXFEE() (*big.Int, error) {
	return _BNum.Contract.MAXFEE(&_BNum.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BNum *BNumCallerSession) MAXFEE() (*big.Int, error) {
	return _BNum.Contract.MAXFEE(&_BNum.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BNum *BNumCaller) MAXINRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MAX_IN_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BNum *BNumSession) MAXINRATIO() (*big.Int, error) {
	return _BNum.Contract.MAXINRATIO(&_BNum.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BNum *BNumCallerSession) MAXINRATIO() (*big.Int, error) {
	return _BNum.Contract.MAXINRATIO(&_BNum.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BNum *BNumCaller) MAXOUTRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MAX_OUT_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BNum *BNumSession) MAXOUTRATIO() (*big.Int, error) {
	return _BNum.Contract.MAXOUTRATIO(&_BNum.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BNum *BNumCallerSession) MAXOUTRATIO() (*big.Int, error) {
	return _BNum.Contract.MAXOUTRATIO(&_BNum.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BNum *BNumCaller) MAXTOTALWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MAX_TOTAL_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BNum *BNumSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _BNum.Contract.MAXTOTALWEIGHT(&_BNum.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BNum *BNumCallerSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _BNum.Contract.MAXTOTALWEIGHT(&_BNum.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BNum *BNumCaller) MAXWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MAX_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BNum *BNumSession) MAXWEIGHT() (*big.Int, error) {
	return _BNum.Contract.MAXWEIGHT(&_BNum.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BNum *BNumCallerSession) MAXWEIGHT() (*big.Int, error) {
	return _BNum.Contract.MAXWEIGHT(&_BNum.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BNum *BNumCaller) MINBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MIN_BALANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BNum *BNumSession) MINBALANCE() (*big.Int, error) {
	return _BNum.Contract.MINBALANCE(&_BNum.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BNum *BNumCallerSession) MINBALANCE() (*big.Int, error) {
	return _BNum.Contract.MINBALANCE(&_BNum.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BNum *BNumCaller) MINBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MIN_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BNum *BNumSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _BNum.Contract.MINBOUNDTOKENS(&_BNum.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BNum *BNumCallerSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _BNum.Contract.MINBOUNDTOKENS(&_BNum.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BNum *BNumCaller) MINBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MIN_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BNum *BNumSession) MINBPOWBASE() (*big.Int, error) {
	return _BNum.Contract.MINBPOWBASE(&_BNum.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BNum *BNumCallerSession) MINBPOWBASE() (*big.Int, error) {
	return _BNum.Contract.MINBPOWBASE(&_BNum.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BNum *BNumCaller) MINFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MIN_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BNum *BNumSession) MINFEE() (*big.Int, error) {
	return _BNum.Contract.MINFEE(&_BNum.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BNum *BNumCallerSession) MINFEE() (*big.Int, error) {
	return _BNum.Contract.MINFEE(&_BNum.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BNum *BNumCaller) MINWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "MIN_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BNum *BNumSession) MINWEIGHT() (*big.Int, error) {
	return _BNum.Contract.MINWEIGHT(&_BNum.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BNum *BNumCallerSession) MINWEIGHT() (*big.Int, error) {
	return _BNum.Contract.MINWEIGHT(&_BNum.CallOpts)
}

// VERSIONNUMBER is a free data retrieval call binding the contract method 0x8025e303.
//
// Solidity: function VERSION_NUMBER() view returns(uint256)
func (_BNum *BNumCaller) VERSIONNUMBER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "VERSION_NUMBER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERSIONNUMBER is a free data retrieval call binding the contract method 0x8025e303.
//
// Solidity: function VERSION_NUMBER() view returns(uint256)
func (_BNum *BNumSession) VERSIONNUMBER() (*big.Int, error) {
	return _BNum.Contract.VERSIONNUMBER(&_BNum.CallOpts)
}

// VERSIONNUMBER is a free data retrieval call binding the contract method 0x8025e303.
//
// Solidity: function VERSION_NUMBER() view returns(uint256)
func (_BNum *BNumCallerSession) VERSIONNUMBER() (*big.Int, error) {
	return _BNum.Contract.VERSIONNUMBER(&_BNum.CallOpts)
}

// WEIGHTCHANGEPCT is a free data retrieval call binding the contract method 0x6852b5cd.
//
// Solidity: function WEIGHT_CHANGE_PCT() view returns(uint256)
func (_BNum *BNumCaller) WEIGHTCHANGEPCT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "WEIGHT_CHANGE_PCT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WEIGHTCHANGEPCT is a free data retrieval call binding the contract method 0x6852b5cd.
//
// Solidity: function WEIGHT_CHANGE_PCT() view returns(uint256)
func (_BNum *BNumSession) WEIGHTCHANGEPCT() (*big.Int, error) {
	return _BNum.Contract.WEIGHTCHANGEPCT(&_BNum.CallOpts)
}

// WEIGHTCHANGEPCT is a free data retrieval call binding the contract method 0x6852b5cd.
//
// Solidity: function WEIGHT_CHANGE_PCT() view returns(uint256)
func (_BNum *BNumCallerSession) WEIGHTCHANGEPCT() (*big.Int, error) {
	return _BNum.Contract.WEIGHTCHANGEPCT(&_BNum.CallOpts)
}

// WEIGHTUPDATEDELAY is a free data retrieval call binding the contract method 0x7e6d71a2.
//
// Solidity: function WEIGHT_UPDATE_DELAY() view returns(uint256)
func (_BNum *BNumCaller) WEIGHTUPDATEDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "WEIGHT_UPDATE_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WEIGHTUPDATEDELAY is a free data retrieval call binding the contract method 0x7e6d71a2.
//
// Solidity: function WEIGHT_UPDATE_DELAY() view returns(uint256)
func (_BNum *BNumSession) WEIGHTUPDATEDELAY() (*big.Int, error) {
	return _BNum.Contract.WEIGHTUPDATEDELAY(&_BNum.CallOpts)
}

// WEIGHTUPDATEDELAY is a free data retrieval call binding the contract method 0x7e6d71a2.
//
// Solidity: function WEIGHT_UPDATE_DELAY() view returns(uint256)
func (_BNum *BNumCallerSession) WEIGHTUPDATEDELAY() (*big.Int, error) {
	return _BNum.Contract.WEIGHTUPDATEDELAY(&_BNum.CallOpts)
}

// Badd is a free data retrieval call binding the contract method 0x0b71b95c.
//
// Solidity: function badd(uint256 a, uint256 b) pure returns(uint256)
func (_BNum *BNumCaller) Badd(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "badd", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Badd is a free data retrieval call binding the contract method 0x0b71b95c.
//
// Solidity: function badd(uint256 a, uint256 b) pure returns(uint256)
func (_BNum *BNumSession) Badd(a *big.Int, b *big.Int) (*big.Int, error) {
	return _BNum.Contract.Badd(&_BNum.CallOpts, a, b)
}

// Badd is a free data retrieval call binding the contract method 0x0b71b95c.
//
// Solidity: function badd(uint256 a, uint256 b) pure returns(uint256)
func (_BNum *BNumCallerSession) Badd(a *big.Int, b *big.Int) (*big.Int, error) {
	return _BNum.Contract.Badd(&_BNum.CallOpts, a, b)
}

// Bdiv is a free data retrieval call binding the contract method 0x7673eb11.
//
// Solidity: function bdiv(uint256 a, uint256 b) pure returns(uint256)
func (_BNum *BNumCaller) Bdiv(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "bdiv", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bdiv is a free data retrieval call binding the contract method 0x7673eb11.
//
// Solidity: function bdiv(uint256 a, uint256 b) pure returns(uint256)
func (_BNum *BNumSession) Bdiv(a *big.Int, b *big.Int) (*big.Int, error) {
	return _BNum.Contract.Bdiv(&_BNum.CallOpts, a, b)
}

// Bdiv is a free data retrieval call binding the contract method 0x7673eb11.
//
// Solidity: function bdiv(uint256 a, uint256 b) pure returns(uint256)
func (_BNum *BNumCallerSession) Bdiv(a *big.Int, b *big.Int) (*big.Int, error) {
	return _BNum.Contract.Bdiv(&_BNum.CallOpts, a, b)
}

// Bfloor is a free data retrieval call binding the contract method 0x317223bc.
//
// Solidity: function bfloor(uint256 a) pure returns(uint256)
func (_BNum *BNumCaller) Bfloor(opts *bind.CallOpts, a *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "bfloor", a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bfloor is a free data retrieval call binding the contract method 0x317223bc.
//
// Solidity: function bfloor(uint256 a) pure returns(uint256)
func (_BNum *BNumSession) Bfloor(a *big.Int) (*big.Int, error) {
	return _BNum.Contract.Bfloor(&_BNum.CallOpts, a)
}

// Bfloor is a free data retrieval call binding the contract method 0x317223bc.
//
// Solidity: function bfloor(uint256 a) pure returns(uint256)
func (_BNum *BNumCallerSession) Bfloor(a *big.Int) (*big.Int, error) {
	return _BNum.Contract.Bfloor(&_BNum.CallOpts, a)
}

// Bmul is a free data retrieval call binding the contract method 0x8c051bf3.
//
// Solidity: function bmul(uint256 a, uint256 b) pure returns(uint256)
func (_BNum *BNumCaller) Bmul(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "bmul", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bmul is a free data retrieval call binding the contract method 0x8c051bf3.
//
// Solidity: function bmul(uint256 a, uint256 b) pure returns(uint256)
func (_BNum *BNumSession) Bmul(a *big.Int, b *big.Int) (*big.Int, error) {
	return _BNum.Contract.Bmul(&_BNum.CallOpts, a, b)
}

// Bmul is a free data retrieval call binding the contract method 0x8c051bf3.
//
// Solidity: function bmul(uint256 a, uint256 b) pure returns(uint256)
func (_BNum *BNumCallerSession) Bmul(a *big.Int, b *big.Int) (*big.Int, error) {
	return _BNum.Contract.Bmul(&_BNum.CallOpts, a, b)
}

// Bpow is a free data retrieval call binding the contract method 0x1d0a82df.
//
// Solidity: function bpow(uint256 base, uint256 exp) pure returns(uint256)
func (_BNum *BNumCaller) Bpow(opts *bind.CallOpts, base *big.Int, exp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "bpow", base, exp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bpow is a free data retrieval call binding the contract method 0x1d0a82df.
//
// Solidity: function bpow(uint256 base, uint256 exp) pure returns(uint256)
func (_BNum *BNumSession) Bpow(base *big.Int, exp *big.Int) (*big.Int, error) {
	return _BNum.Contract.Bpow(&_BNum.CallOpts, base, exp)
}

// Bpow is a free data retrieval call binding the contract method 0x1d0a82df.
//
// Solidity: function bpow(uint256 base, uint256 exp) pure returns(uint256)
func (_BNum *BNumCallerSession) Bpow(base *big.Int, exp *big.Int) (*big.Int, error) {
	return _BNum.Contract.Bpow(&_BNum.CallOpts, base, exp)
}

// BpowApprox is a free data retrieval call binding the contract method 0xd011f32a.
//
// Solidity: function bpowApprox(uint256 base, uint256 exp, uint256 precision) pure returns(uint256)
func (_BNum *BNumCaller) BpowApprox(opts *bind.CallOpts, base *big.Int, exp *big.Int, precision *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "bpowApprox", base, exp, precision)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BpowApprox is a free data retrieval call binding the contract method 0xd011f32a.
//
// Solidity: function bpowApprox(uint256 base, uint256 exp, uint256 precision) pure returns(uint256)
func (_BNum *BNumSession) BpowApprox(base *big.Int, exp *big.Int, precision *big.Int) (*big.Int, error) {
	return _BNum.Contract.BpowApprox(&_BNum.CallOpts, base, exp, precision)
}

// BpowApprox is a free data retrieval call binding the contract method 0xd011f32a.
//
// Solidity: function bpowApprox(uint256 base, uint256 exp, uint256 precision) pure returns(uint256)
func (_BNum *BNumCallerSession) BpowApprox(base *big.Int, exp *big.Int, precision *big.Int) (*big.Int, error) {
	return _BNum.Contract.BpowApprox(&_BNum.CallOpts, base, exp, precision)
}

// Bpowi is a free data retrieval call binding the contract method 0xcd660299.
//
// Solidity: function bpowi(uint256 a, uint256 n) pure returns(uint256)
func (_BNum *BNumCaller) Bpowi(opts *bind.CallOpts, a *big.Int, n *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "bpowi", a, n)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bpowi is a free data retrieval call binding the contract method 0xcd660299.
//
// Solidity: function bpowi(uint256 a, uint256 n) pure returns(uint256)
func (_BNum *BNumSession) Bpowi(a *big.Int, n *big.Int) (*big.Int, error) {
	return _BNum.Contract.Bpowi(&_BNum.CallOpts, a, n)
}

// Bpowi is a free data retrieval call binding the contract method 0xcd660299.
//
// Solidity: function bpowi(uint256 a, uint256 n) pure returns(uint256)
func (_BNum *BNumCallerSession) Bpowi(a *big.Int, n *big.Int) (*big.Int, error) {
	return _BNum.Contract.Bpowi(&_BNum.CallOpts, a, n)
}

// Bsub is a free data retrieval call binding the contract method 0x3cf3c7d4.
//
// Solidity: function bsub(uint256 a, uint256 b) pure returns(uint256)
func (_BNum *BNumCaller) Bsub(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "bsub", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bsub is a free data retrieval call binding the contract method 0x3cf3c7d4.
//
// Solidity: function bsub(uint256 a, uint256 b) pure returns(uint256)
func (_BNum *BNumSession) Bsub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _BNum.Contract.Bsub(&_BNum.CallOpts, a, b)
}

// Bsub is a free data retrieval call binding the contract method 0x3cf3c7d4.
//
// Solidity: function bsub(uint256 a, uint256 b) pure returns(uint256)
func (_BNum *BNumCallerSession) Bsub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _BNum.Contract.Bsub(&_BNum.CallOpts, a, b)
}

// BsubSign is a free data retrieval call binding the contract method 0x743480cc.
//
// Solidity: function bsubSign(uint256 a, uint256 b) pure returns(uint256, bool)
func (_BNum *BNumCaller) BsubSign(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, bool, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "bsubSign", a, b)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// BsubSign is a free data retrieval call binding the contract method 0x743480cc.
//
// Solidity: function bsubSign(uint256 a, uint256 b) pure returns(uint256, bool)
func (_BNum *BNumSession) BsubSign(a *big.Int, b *big.Int) (*big.Int, bool, error) {
	return _BNum.Contract.BsubSign(&_BNum.CallOpts, a, b)
}

// BsubSign is a free data retrieval call binding the contract method 0x743480cc.
//
// Solidity: function bsubSign(uint256 a, uint256 b) pure returns(uint256, bool)
func (_BNum *BNumCallerSession) BsubSign(a *big.Int, b *big.Int) (*big.Int, bool, error) {
	return _BNum.Contract.BsubSign(&_BNum.CallOpts, a, b)
}

// Btoi is a free data retrieval call binding the contract method 0x859ae48e.
//
// Solidity: function btoi(uint256 a) pure returns(uint256)
func (_BNum *BNumCaller) Btoi(opts *bind.CallOpts, a *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BNum.contract.Call(opts, &out, "btoi", a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Btoi is a free data retrieval call binding the contract method 0x859ae48e.
//
// Solidity: function btoi(uint256 a) pure returns(uint256)
func (_BNum *BNumSession) Btoi(a *big.Int) (*big.Int, error) {
	return _BNum.Contract.Btoi(&_BNum.CallOpts, a)
}

// Btoi is a free data retrieval call binding the contract method 0x859ae48e.
//
// Solidity: function btoi(uint256 a) pure returns(uint256)
func (_BNum *BNumCallerSession) Btoi(a *big.Int) (*big.Int, error) {
	return _BNum.Contract.Btoi(&_BNum.CallOpts, a)
}
