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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ArrayABI is the input ABI used to generate the binding from.
const ArrayABI = "[{\"inputs\":[],\"name\":\"get_array\",\"outputs\":[{\"internalType\":\"uint8[2][4]\",\"name\":\"\",\"type\":\"uint8[2][4]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArrayFuncSigs maps the 4-byte function signature to its string representation.
var ArrayFuncSigs = map[string]string{
	"c75d70ed": "get_array()",
}

// ArrayBin is the compiled bytecode used for deploying new contracts.
var ArrayBin = "0x608060405234801561001057600080fd5b5061016a806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063c75d70ed14610030575b600080fd5b610038610092565b6040516000826004835b818410156100825760208402830151604080838360005b83811015610071578181015183820152602001610059565b505050509050019260010192610042565b9250505091505060405180910390f35b61009a6100e9565b6100a26100e9565b805160009052602080820180516001905260408301805160029052606084018051600390528451600490850152915160059084015251600690830152516007910152919050565b60405180608001604052806004905b610100610116565b8152602001906001900390816100f85790505090565b6040518060400160405280600290602082028036833750919291505056fea2646970667358221220ce9451a2cb130dc17cb76dc3b286d550a95cf59afe2bdfb7905cd6c4a4267c1564736f6c63430006050033"

// DeployArray deploys a new Ethereum contract, binding an instance of Array to it.
func DeployArray(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Array, error) {
	parsed, err := abi.JSON(strings.NewReader(ArrayABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArrayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Array{ArrayCaller: ArrayCaller{contract: contract}, ArrayTransactor: ArrayTransactor{contract: contract}, ArrayFilterer: ArrayFilterer{contract: contract}}, nil
}

// Array is an auto generated Go binding around an Ethereum contract.
type Array struct {
	ArrayCaller     // Read-only binding to the contract
	ArrayTransactor // Write-only binding to the contract
	ArrayFilterer   // Log filterer for contract events
}

// ArrayCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArrayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArrayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArrayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArrayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArrayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArraySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArraySession struct {
	Contract     *Array            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArrayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArrayCallerSession struct {
	Contract *ArrayCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ArrayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArrayTransactorSession struct {
	Contract     *ArrayTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArrayRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArrayRaw struct {
	Contract *Array // Generic contract binding to access the raw methods on
}

// ArrayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArrayCallerRaw struct {
	Contract *ArrayCaller // Generic read-only contract binding to access the raw methods on
}

// ArrayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArrayTransactorRaw struct {
	Contract *ArrayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArray creates a new instance of Array, bound to a specific deployed contract.
func NewArray(address common.Address, backend bind.ContractBackend) (*Array, error) {
	contract, err := bindArray(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Array{ArrayCaller: ArrayCaller{contract: contract}, ArrayTransactor: ArrayTransactor{contract: contract}, ArrayFilterer: ArrayFilterer{contract: contract}}, nil
}

// NewArrayCaller creates a new read-only instance of Array, bound to a specific deployed contract.
func NewArrayCaller(address common.Address, caller bind.ContractCaller) (*ArrayCaller, error) {
	contract, err := bindArray(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArrayCaller{contract: contract}, nil
}

// NewArrayTransactor creates a new write-only instance of Array, bound to a specific deployed contract.
func NewArrayTransactor(address common.Address, transactor bind.ContractTransactor) (*ArrayTransactor, error) {
	contract, err := bindArray(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArrayTransactor{contract: contract}, nil
}

// NewArrayFilterer creates a new log filterer instance of Array, bound to a specific deployed contract.
func NewArrayFilterer(address common.Address, filterer bind.ContractFilterer) (*ArrayFilterer, error) {
	contract, err := bindArray(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArrayFilterer{contract: contract}, nil
}

// bindArray binds a generic wrapper to an already deployed contract.
func bindArray(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArrayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Array *ArrayRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Array.Contract.ArrayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Array *ArrayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Array.Contract.ArrayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Array *ArrayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Array.Contract.ArrayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Array *ArrayCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Array.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Array *ArrayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Array.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Array *ArrayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Array.Contract.contract.Transact(opts, method, params...)
}

// GetArray is a free data retrieval call binding the contract method 0xc75d70ed.
//
// Solidity: function get_array() constant returns(uint8[2][4])
func (_Array *ArrayCaller) GetArray(opts *bind.CallOpts) ([4][2]uint8, error) {
	var (
		ret0 = new([4][2]uint8)
	)
	out := ret0
	err := _Array.contract.Call(opts, out, "get_array")
	return *ret0, err
}

// GetArray is a free data retrieval call binding the contract method 0xc75d70ed.
//
// Solidity: function get_array() constant returns(uint8[2][4])
func (_Array *ArraySession) GetArray() ([4][2]uint8, error) {
	return _Array.Contract.GetArray(&_Array.CallOpts)
}

// GetArray is a free data retrieval call binding the contract method 0xc75d70ed.
//
// Solidity: function get_array() constant returns(uint8[2][4])
func (_Array *ArrayCallerSession) GetArray() ([4][2]uint8, error) {
	return _Array.Contract.GetArray(&_Array.CallOpts)
}

// BigBoardABI is the input ABI used to generate the binding from.
const BigBoardABI = "[{\"inputs\":[],\"name\":\"fill1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_board_state\",\"outputs\":[{\"internalType\":\"uint8[9][10][3]\",\"name\":\"\",\"type\":\"uint8[9][10][3]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BigBoardFuncSigs maps the 4-byte function signature to its string representation.
var BigBoardFuncSigs = map[string]string{
	"0350e928": "fill1()",
	"810851a8": "get_board_state()",
}

// BigBoardBin is the compiled bytecode used for deploying new contracts.
var BigBoardBin = "0x608060405234801561001057600080fd5b506102c8806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80630350e9281461003b578063810851a814610045575b600080fd5b6100436100cb565b005b61004d61015e565b6040516000826003835b818410156100bb5782846020020151600a6000925b818410156100ad576020840283015161012080838360005b8381101561009c578181015183820152602001610084565b50505050905001926001019261006c565b925050509260010192610057565b9250505091505060405180910390f35b60005b600381101561015b5760005b600a8110156101525760005b60098110156101495780820183026000846003811061010157fe5b600a020183600a811061011057fe5b01826009811061011c57fe5b602091828204019190066101000a81548160ff021916908360ff16021790555080806001019150506100e6565b506001016100da565b506001016100ce565b50565b610166610218565b60408051606081019091526000600381835b8282101561020f57604080516101408101909152600a8084028601906000835b828210156101fc57604080516101208101918290529085840190600990826000855b825461010083900a900460ff168152602060019283018181049485019490930390920291018084116101ba579050505050505081526020019060010190610198565b5050505081526020019060010190610178565b50505050905090565b60405180606001604052806003905b61022f610245565b8152602001906001900390816102275790505090565b604051806101400160405280600a905b61025d610273565b8152602001906001900390816102555790505090565b604051806101200160405280600990602082028036833750919291505056fea2646970667358221220fa825c09acbb2d35a020b3d9830ef25f4b046cdccef6cb8aa7b23ced4bcece3464736f6c63430006050033"

// DeployBigBoard deploys a new Ethereum contract, binding an instance of BigBoard to it.
func DeployBigBoard(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BigBoard, error) {
	parsed, err := abi.JSON(strings.NewReader(BigBoardABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BigBoardBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BigBoard{BigBoardCaller: BigBoardCaller{contract: contract}, BigBoardTransactor: BigBoardTransactor{contract: contract}, BigBoardFilterer: BigBoardFilterer{contract: contract}}, nil
}

// BigBoard is an auto generated Go binding around an Ethereum contract.
type BigBoard struct {
	BigBoardCaller     // Read-only binding to the contract
	BigBoardTransactor // Write-only binding to the contract
	BigBoardFilterer   // Log filterer for contract events
}

// BigBoardCaller is an auto generated read-only Go binding around an Ethereum contract.
type BigBoardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BigBoardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BigBoardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BigBoardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BigBoardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BigBoardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BigBoardSession struct {
	Contract     *BigBoard         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BigBoardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BigBoardCallerSession struct {
	Contract *BigBoardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BigBoardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BigBoardTransactorSession struct {
	Contract     *BigBoardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BigBoardRaw is an auto generated low-level Go binding around an Ethereum contract.
type BigBoardRaw struct {
	Contract *BigBoard // Generic contract binding to access the raw methods on
}

// BigBoardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BigBoardCallerRaw struct {
	Contract *BigBoardCaller // Generic read-only contract binding to access the raw methods on
}

// BigBoardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BigBoardTransactorRaw struct {
	Contract *BigBoardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBigBoard creates a new instance of BigBoard, bound to a specific deployed contract.
func NewBigBoard(address common.Address, backend bind.ContractBackend) (*BigBoard, error) {
	contract, err := bindBigBoard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BigBoard{BigBoardCaller: BigBoardCaller{contract: contract}, BigBoardTransactor: BigBoardTransactor{contract: contract}, BigBoardFilterer: BigBoardFilterer{contract: contract}}, nil
}

// NewBigBoardCaller creates a new read-only instance of BigBoard, bound to a specific deployed contract.
func NewBigBoardCaller(address common.Address, caller bind.ContractCaller) (*BigBoardCaller, error) {
	contract, err := bindBigBoard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BigBoardCaller{contract: contract}, nil
}

// NewBigBoardTransactor creates a new write-only instance of BigBoard, bound to a specific deployed contract.
func NewBigBoardTransactor(address common.Address, transactor bind.ContractTransactor) (*BigBoardTransactor, error) {
	contract, err := bindBigBoard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BigBoardTransactor{contract: contract}, nil
}

// NewBigBoardFilterer creates a new log filterer instance of BigBoard, bound to a specific deployed contract.
func NewBigBoardFilterer(address common.Address, filterer bind.ContractFilterer) (*BigBoardFilterer, error) {
	contract, err := bindBigBoard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BigBoardFilterer{contract: contract}, nil
}

// bindBigBoard binds a generic wrapper to an already deployed contract.
func bindBigBoard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BigBoardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BigBoard *BigBoardRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BigBoard.Contract.BigBoardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BigBoard *BigBoardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BigBoard.Contract.BigBoardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BigBoard *BigBoardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BigBoard.Contract.BigBoardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BigBoard *BigBoardCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BigBoard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BigBoard *BigBoardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BigBoard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BigBoard *BigBoardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BigBoard.Contract.contract.Transact(opts, method, params...)
}

// GetBoardState is a free data retrieval call binding the contract method 0x810851a8.
//
// Solidity: function get_board_state() constant returns(uint8[9][10][3])
func (_BigBoard *BigBoardCaller) GetBoardState(opts *bind.CallOpts) ([3][10][9]uint8, error) {
	var (
		ret0 = new([3][10][9]uint8)
	)
	out := ret0
	err := _BigBoard.contract.Call(opts, out, "get_board_state")
	return *ret0, err
}

// GetBoardState is a free data retrieval call binding the contract method 0x810851a8.
//
// Solidity: function get_board_state() constant returns(uint8[9][10][3])
func (_BigBoard *BigBoardSession) GetBoardState() ([3][10][9]uint8, error) {
	return _BigBoard.Contract.GetBoardState(&_BigBoard.CallOpts)
}

// GetBoardState is a free data retrieval call binding the contract method 0x810851a8.
//
// Solidity: function get_board_state() constant returns(uint8[9][10][3])
func (_BigBoard *BigBoardCallerSession) GetBoardState() ([3][10][9]uint8, error) {
	return _BigBoard.Contract.GetBoardState(&_BigBoard.CallOpts)
}

// Fill1 is a paid mutator transaction binding the contract method 0x0350e928.
//
// Solidity: function fill1() returns()
func (_BigBoard *BigBoardTransactor) Fill1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BigBoard.contract.Transact(opts, "fill1")
}

// Fill1 is a paid mutator transaction binding the contract method 0x0350e928.
//
// Solidity: function fill1() returns()
func (_BigBoard *BigBoardSession) Fill1() (*types.Transaction, error) {
	return _BigBoard.Contract.Fill1(&_BigBoard.TransactOpts)
}

// Fill1 is a paid mutator transaction binding the contract method 0x0350e928.
//
// Solidity: function fill1() returns()
func (_BigBoard *BigBoardTransactorSession) Fill1() (*types.Transaction, error) {
	return _BigBoard.Contract.Fill1(&_BigBoard.TransactOpts)
}

// EventerABI is the input ABI used to generate the binding from.
const EventerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"AnonEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int8\",\"name\":\"out1\",\"type\":\"int8\"},{\"indexed\":true,\"internalType\":\"int8\",\"name\":\"out2\",\"type\":\"int8\"}],\"name\":\"TestInt8\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"anonEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fail\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EventerFuncSigs maps the 4-byte function signature to its string representation.
var EventerFuncSigs = map[string]string{
	"529c2b1f": "anonEvent()",
	"a9cc4718": "fail()",
	"bf819c20": "getEvent()",
}

// EventerBin is the compiled bytecode used for deploying new contracts.
var EventerBin = "0x608060405234801561001057600080fd5b50610141806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063529c2b1f14610046578063a9cc471814610050578063bf819c2014610058575b600080fd5b61004e610060565b005b61004e61009c565b61004e6100d8565b6040805133808252602082015281517f5aabb901501aff3228b2a4010c287e9d9fe99f2a3d5036dd6cdd2829b567f64f929181900390910190a1565b6040805162461bcd60e51b815260206004820152600c60248201526b6572726f7220737472696e6760a01b604482015290519081900360640190fd5b60405160021990600119907f8f50d21be7587a4814a9d4c10b7c8d1eea6389adbd44cb59ddaba790fd2ecbbd90600090a356fea26469706673582212202c4f5fd37a174d1ddfb1cbfb8dd8102707bc18efb96aa3a8f739f6c02e2557ed64736f6c63430006050033"

// DeployEventer deploys a new Ethereum contract, binding an instance of Eventer to it.
func DeployEventer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Eventer, error) {
	parsed, err := abi.JSON(strings.NewReader(EventerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EventerBin), backend)
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
	parsed, err := abi.JSON(strings.NewReader(EventerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eventer *EventerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_Eventer *EventerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
// Solidity: event AnonEvent(address , address )
func (_Eventer *EventerFilterer) FilterAnonEvent(opts *bind.FilterOpts) (*EventerAnonEventIterator, error) {

	logs, sub, err := _Eventer.contract.FilterLogs(opts, "AnonEvent")
	if err != nil {
		return nil, err
	}
	return &EventerAnonEventIterator{contract: _Eventer.contract, event: "AnonEvent", logs: logs, sub: sub}, nil
}

// WatchAnonEvent is a free log subscription operation binding the contract event 0x5aabb901501aff3228b2a4010c287e9d9fe99f2a3d5036dd6cdd2829b567f64f.
//
// Solidity: event AnonEvent(address , address )
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
// Solidity: event AnonEvent(address , address )
func (_Eventer *EventerFilterer) ParseAnonEvent(log types.Log) (*EventerAnonEvent, error) {
	event := new(EventerAnonEvent)
	if err := _Eventer.contract.UnpackLog(event, "AnonEvent", log); err != nil {
		return nil, err
	}
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
	return event, nil
}

// ReceiveFallbackABI is the input ABI used to generate the binding from.
const ReceiveFallbackABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Fallback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Receive\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ReceiveFallbackBin is the compiled bytecode used for deploying new contracts.
var ReceiveFallbackBin = "0x6080604052348015600f57600080fd5b5060b48061001e6000396000f3fe608060405236603d576040805133815290517f681ad3cb4f20e39dd01bc0e45f0bc6f95ea2de196277ed6ec4f7357e2e2044b89181900360200190a1005b348015604857600080fd5b506040805133815290517fdd0c5778ddb0b0501649417aeffed44057db3d5d530231758c8dc5ef6eeb78949181900360200190a100fea2646970667358221220f1e64bd208f5b27009f728fb0c5c2289ea5d8142cd42796c117b40360e96ebe464736f6c63430006050033"

// DeployReceiveFallback deploys a new Ethereum contract, binding an instance of ReceiveFallback to it.
func DeployReceiveFallback(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReceiveFallback, error) {
	parsed, err := abi.JSON(strings.NewReader(ReceiveFallbackABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReceiveFallbackBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReceiveFallback{ReceiveFallbackCaller: ReceiveFallbackCaller{contract: contract}, ReceiveFallbackTransactor: ReceiveFallbackTransactor{contract: contract}, ReceiveFallbackFilterer: ReceiveFallbackFilterer{contract: contract}}, nil
}

// ReceiveFallback is an auto generated Go binding around an Ethereum contract.
type ReceiveFallback struct {
	ReceiveFallbackCaller     // Read-only binding to the contract
	ReceiveFallbackTransactor // Write-only binding to the contract
	ReceiveFallbackFilterer   // Log filterer for contract events
}

// ReceiveFallbackCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReceiveFallbackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiveFallbackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReceiveFallbackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiveFallbackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReceiveFallbackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiveFallbackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReceiveFallbackSession struct {
	Contract     *ReceiveFallback  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReceiveFallbackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReceiveFallbackCallerSession struct {
	Contract *ReceiveFallbackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReceiveFallbackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReceiveFallbackTransactorSession struct {
	Contract     *ReceiveFallbackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReceiveFallbackRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReceiveFallbackRaw struct {
	Contract *ReceiveFallback // Generic contract binding to access the raw methods on
}

// ReceiveFallbackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReceiveFallbackCallerRaw struct {
	Contract *ReceiveFallbackCaller // Generic read-only contract binding to access the raw methods on
}

// ReceiveFallbackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReceiveFallbackTransactorRaw struct {
	Contract *ReceiveFallbackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReceiveFallback creates a new instance of ReceiveFallback, bound to a specific deployed contract.
func NewReceiveFallback(address common.Address, backend bind.ContractBackend) (*ReceiveFallback, error) {
	contract, err := bindReceiveFallback(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReceiveFallback{ReceiveFallbackCaller: ReceiveFallbackCaller{contract: contract}, ReceiveFallbackTransactor: ReceiveFallbackTransactor{contract: contract}, ReceiveFallbackFilterer: ReceiveFallbackFilterer{contract: contract}}, nil
}

// NewReceiveFallbackCaller creates a new read-only instance of ReceiveFallback, bound to a specific deployed contract.
func NewReceiveFallbackCaller(address common.Address, caller bind.ContractCaller) (*ReceiveFallbackCaller, error) {
	contract, err := bindReceiveFallback(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiveFallbackCaller{contract: contract}, nil
}

// NewReceiveFallbackTransactor creates a new write-only instance of ReceiveFallback, bound to a specific deployed contract.
func NewReceiveFallbackTransactor(address common.Address, transactor bind.ContractTransactor) (*ReceiveFallbackTransactor, error) {
	contract, err := bindReceiveFallback(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiveFallbackTransactor{contract: contract}, nil
}

// NewReceiveFallbackFilterer creates a new log filterer instance of ReceiveFallback, bound to a specific deployed contract.
func NewReceiveFallbackFilterer(address common.Address, filterer bind.ContractFilterer) (*ReceiveFallbackFilterer, error) {
	contract, err := bindReceiveFallback(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReceiveFallbackFilterer{contract: contract}, nil
}

// bindReceiveFallback binds a generic wrapper to an already deployed contract.
func bindReceiveFallback(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReceiveFallbackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiveFallback *ReceiveFallbackRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReceiveFallback.Contract.ReceiveFallbackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiveFallback *ReceiveFallbackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiveFallback.Contract.ReceiveFallbackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiveFallback *ReceiveFallbackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiveFallback.Contract.ReceiveFallbackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiveFallback *ReceiveFallbackCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReceiveFallback.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiveFallback *ReceiveFallbackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiveFallback.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiveFallback *ReceiveFallbackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiveFallback.Contract.contract.Transact(opts, method, params...)
}

// ReceiveFallbackFallbackIterator is returned from FilterFallback and is used to iterate over the raw logs and unpacked data for Fallback events raised by the ReceiveFallback contract.
type ReceiveFallbackFallbackIterator struct {
	Event *ReceiveFallbackFallback // Event containing the contract specifics and raw log

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
func (it *ReceiveFallbackFallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiveFallbackFallback)
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
		it.Event = new(ReceiveFallbackFallback)
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
func (it *ReceiveFallbackFallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiveFallbackFallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiveFallbackFallback represents a Fallback event raised by the ReceiveFallback contract.
type ReceiveFallbackFallback struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFallback is a free log retrieval operation binding the contract event 0xdd0c5778ddb0b0501649417aeffed44057db3d5d530231758c8dc5ef6eeb7894.
//
// Solidity: event Fallback(address sender)
func (_ReceiveFallback *ReceiveFallbackFilterer) FilterFallback(opts *bind.FilterOpts) (*ReceiveFallbackFallbackIterator, error) {

	logs, sub, err := _ReceiveFallback.contract.FilterLogs(opts, "Fallback")
	if err != nil {
		return nil, err
	}
	return &ReceiveFallbackFallbackIterator{contract: _ReceiveFallback.contract, event: "Fallback", logs: logs, sub: sub}, nil
}

// WatchFallback is a free log subscription operation binding the contract event 0xdd0c5778ddb0b0501649417aeffed44057db3d5d530231758c8dc5ef6eeb7894.
//
// Solidity: event Fallback(address sender)
func (_ReceiveFallback *ReceiveFallbackFilterer) WatchFallback(opts *bind.WatchOpts, sink chan<- *ReceiveFallbackFallback) (event.Subscription, error) {

	logs, sub, err := _ReceiveFallback.contract.WatchLogs(opts, "Fallback")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiveFallbackFallback)
				if err := _ReceiveFallback.contract.UnpackLog(event, "Fallback", log); err != nil {
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

// ParseFallback is a log parse operation binding the contract event 0xdd0c5778ddb0b0501649417aeffed44057db3d5d530231758c8dc5ef6eeb7894.
//
// Solidity: event Fallback(address sender)
func (_ReceiveFallback *ReceiveFallbackFilterer) ParseFallback(log types.Log) (*ReceiveFallbackFallback, error) {
	event := new(ReceiveFallbackFallback)
	if err := _ReceiveFallback.contract.UnpackLog(event, "Fallback", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReceiveFallbackReceiveIterator is returned from FilterReceive and is used to iterate over the raw logs and unpacked data for Receive events raised by the ReceiveFallback contract.
type ReceiveFallbackReceiveIterator struct {
	Event *ReceiveFallbackReceive // Event containing the contract specifics and raw log

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
func (it *ReceiveFallbackReceiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReceiveFallbackReceive)
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
		it.Event = new(ReceiveFallbackReceive)
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
func (it *ReceiveFallbackReceiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReceiveFallbackReceiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReceiveFallbackReceive represents a Receive event raised by the ReceiveFallback contract.
type ReceiveFallbackReceive struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceive is a free log retrieval operation binding the contract event 0x681ad3cb4f20e39dd01bc0e45f0bc6f95ea2de196277ed6ec4f7357e2e2044b8.
//
// Solidity: event Receive(address sender)
func (_ReceiveFallback *ReceiveFallbackFilterer) FilterReceive(opts *bind.FilterOpts) (*ReceiveFallbackReceiveIterator, error) {

	logs, sub, err := _ReceiveFallback.contract.FilterLogs(opts, "Receive")
	if err != nil {
		return nil, err
	}
	return &ReceiveFallbackReceiveIterator{contract: _ReceiveFallback.contract, event: "Receive", logs: logs, sub: sub}, nil
}

// WatchReceive is a free log subscription operation binding the contract event 0x681ad3cb4f20e39dd01bc0e45f0bc6f95ea2de196277ed6ec4f7357e2e2044b8.
//
// Solidity: event Receive(address sender)
func (_ReceiveFallback *ReceiveFallbackFilterer) WatchReceive(opts *bind.WatchOpts, sink chan<- *ReceiveFallbackReceive) (event.Subscription, error) {

	logs, sub, err := _ReceiveFallback.contract.WatchLogs(opts, "Receive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReceiveFallbackReceive)
				if err := _ReceiveFallback.contract.UnpackLog(event, "Receive", log); err != nil {
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

// ParseReceive is a log parse operation binding the contract event 0x681ad3cb4f20e39dd01bc0e45f0bc6f95ea2de196277ed6ec4f7357e2e2044b8.
//
// Solidity: event Receive(address sender)
func (_ReceiveFallback *ReceiveFallbackFilterer) ParseReceive(log types.Log) (*ReceiveFallbackReceive, error) {
	event := new(ReceiveFallbackReceive)
	if err := _ReceiveFallback.contract.UnpackLog(event, "Receive", log); err != nil {
		return nil, err
	}
	return event, nil
}
