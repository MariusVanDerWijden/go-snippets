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

// T is an auto generated low-level Go binding around an user-defined struct.
type T struct {
	X *big.Int
	Y *big.Int
}

// TupleTest2A is an auto generated low-level Go binding around an user-defined struct.
type TupleTest2A struct {
	A *big.Int
	B []*big.Int
}

// TupleTest2S is an auto generated low-level Go binding around an user-defined struct.
type TupleTest2S struct {
	A *big.Int
	B []*big.Int
	C []T
}

// TupleTestS is an auto generated low-level Go binding around an user-defined struct.
type TupleTestS struct {
	A *big.Int
	B []*big.Int
	C []T
}

// ArrayABI is the input ABI used to generate the binding from.
const ArrayABI = "[{\"inputs\":[],\"name\":\"get_array\",\"outputs\":[{\"internalType\":\"uint8[2][4]\",\"name\":\"\",\"type\":\"uint8[2][4]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArrayFuncSigs maps the 4-byte function signature to its string representation.
var ArrayFuncSigs = map[string]string{
	"c75d70ed": "get_array()",
}

// ArrayBin is the compiled bytecode used for deploying new contracts.
var ArrayBin = "0x608060405234801561001057600080fd5b50610187806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063c75d70ed14610030575b600080fd5b61003861004e565b60405161004591906100f0565b60405180910390f35b6100566100a5565b61005e6100a5565b805160009052602080820180516001905260408301805160029052606084018051600390528451600490850152915160059084015251600690830152516007910152919050565b60405180608001604052806004905b6100bc6100d2565b8152602001906001900390816100b45790505090565b60405180604001604052806002906020820280368337509192915050565b610100810181836000805b6004811015610147578251829085905b600283101561012e57805160ff168252600192909201916020918201910161010b565b50505060409390930192602092909201916001016100fb565b505050509291505056fea264697066735822122047b11e0d9619d5deff5c7629e48ad60567de8e292d37f8dd3dbcc2ad8483533464736f6c63430006050033"

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
func (_Array *ArrayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_Array *ArrayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
// Solidity: function get_array() pure returns(uint8[2][4])
func (_Array *ArrayCaller) GetArray(opts *bind.CallOpts) ([4][2]uint8, error) {
	var out []interface{}
	err := _Array.contract.Call(opts, &out, "get_array")

	if err != nil {
		return *new([4][2]uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([4][2]uint8)).(*[4][2]uint8)

	return out0, err

}

// GetArray is a free data retrieval call binding the contract method 0xc75d70ed.
//
// Solidity: function get_array() pure returns(uint8[2][4])
func (_Array *ArraySession) GetArray() ([4][2]uint8, error) {
	return _Array.Contract.GetArray(&_Array.CallOpts)
}

// GetArray is a free data retrieval call binding the contract method 0xc75d70ed.
//
// Solidity: function get_array() pure returns(uint8[2][4])
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
var BigBoardBin = "0x608060405234801561001057600080fd5b506102ec806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80630350e9281461003b578063810851a814610045575b600080fd5b610043610063565b005b61004d6100f6565b60405161005a919061022a565b60405180910390f35b60005b60038110156100f35760005b600a8110156100ea5760005b60098110156100e15780820183026000846003811061009957fe5b600a020183600a81106100a857fe5b0182600981106100b457fe5b602091828204019190066101000a81548160ff021916908360ff160217905550808060010191505061007e565b50600101610072565b50600101610066565b50565b6100fe6101b0565b60408051606081019091526000600381835b828210156101a757604080516101408101909152600a8084028601906000835b8282101561019457604080516101208101918290529085840190600990826000855b825461010083900a900460ff16815260206001928301818104948501949093039092029101808411610152579050505050505081526020019060010190610130565b5050505081526020019060010190610110565b50505050905090565b60405180606001604052806003905b6101c76101dd565b8152602001906001900390816101bf5790505090565b604051806101400160405280600a905b6101f561020b565b8152602001906001900390816101ed5790505090565b6040518061012001604052806009906020820280368337509192915050565b6121c08101818360005b60038110156102ad57815160009084905b600a83101561029357805160009083905b600983101561027957805160ff1682526001929092019160209182019101610256565b505050600192909201916101209190910190602001610245565b505050610b40929092019160209190910190600101610234565b5050509291505056fea26469706673582212206966b3c2980316e0efe4a006bcea57a750572866e8cdff7265adfe0c78dc0f4c64736f6c63430006050033"

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
func (_BigBoard *BigBoardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_BigBoard *BigBoardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
// Solidity: function get_board_state() view returns(uint8[9][10][3])
func (_BigBoard *BigBoardCaller) GetBoardState(opts *bind.CallOpts) ([3][10][9]uint8, error) {
	var out []interface{}
	err := _BigBoard.contract.Call(opts, &out, "get_board_state")

	if err != nil {
		return *new([3][10][9]uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([3][10][9]uint8)).(*[3][10][9]uint8)

	return out0, err

}

// GetBoardState is a free data retrieval call binding the contract method 0x810851a8.
//
// Solidity: function get_board_state() view returns(uint8[9][10][3])
func (_BigBoard *BigBoardSession) GetBoardState() ([3][10][9]uint8, error) {
	return _BigBoard.Contract.GetBoardState(&_BigBoard.CallOpts)
}

// GetBoardState is a free data retrieval call binding the contract method 0x810851a8.
//
// Solidity: function get_board_state() view returns(uint8[9][10][3])
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
var EventerBin = "0x608060405234801561001057600080fd5b50610165806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063529c2b1f14610046578063a9cc471814610050578063bf819c2014610058575b600080fd5b61004e610060565b005b61004e61009b565b61004e6100bc565b7f5aabb901501aff3228b2a4010c287e9d9fe99f2a3d5036dd6cdd2829b567f64f33336040516100919291906100ef565b60405180910390a1565b60405162461bcd60e51b81526004016100b390610109565b60405180910390fd5b60405160021990600119907f8f50d21be7587a4814a9d4c10b7c8d1eea6389adbd44cb59ddaba790fd2ecbbd90600090a3565b6001600160a01b0392831681529116602082015260400190565b6020808252600c908201526b6572726f7220737472696e6760a01b60408201526060019056fea2646970667358221220fca32549fa1266a0e1db054b2ab47dae3253a74a5ac25cc5a3039dc41877ce8364736f6c63430006050033"

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
var ReceiveFallbackBin = "0x608060405234801561001057600080fd5b5060bf8061001f6000396000f3fe608060405236603f577f681ad3cb4f20e39dd01bc0e45f0bc6f95ea2de196277ed6ec4f7357e2e2044b833604051603591906075565b60405180910390a1005b348015604a57600080fd5b507fdd0c5778ddb0b0501649417aeffed44057db3d5d530231758c8dc5ef6eeb789433604051603591905b6001600160a01b039190911681526020019056fea2646970667358221220f5119e1076f2453921518f753ebcecb5d3244e7fca85f5e75bdbaaa951981e9964736f6c63430006050033"

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
func (_ReceiveFallback *ReceiveFallbackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_ReceiveFallback *ReceiveFallbackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ReceiveFallback *ReceiveFallbackTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ReceiveFallback.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ReceiveFallback *ReceiveFallbackSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ReceiveFallback.Contract.Fallback(&_ReceiveFallback.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_ReceiveFallback *ReceiveFallbackTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ReceiveFallback.Contract.Fallback(&_ReceiveFallback.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ReceiveFallback *ReceiveFallbackTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiveFallback.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ReceiveFallback *ReceiveFallbackSession) Receive() (*types.Transaction, error) {
	return _ReceiveFallback.Contract.Receive(&_ReceiveFallback.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ReceiveFallback *ReceiveFallbackTransactorSession) Receive() (*types.Transaction, error) {
	return _ReceiveFallback.Contract.Receive(&_ReceiveFallback.TransactOpts)
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

// TupleTestABI is the input ABI used to generate the binding from.
const TupleTestABI = "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"b\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structT[]\",\"name\":\"c\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structTupleTest.S\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structT\",\"name\":\"\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"evF\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"b\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structT[]\",\"name\":\"c\",\"type\":\"tuple[]\"}],\"internalType\":\"structTupleTest.S\",\"name\":\"s\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structT\",\"name\":\"t\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"u\",\"type\":\"uint256\"}],\"name\":\"f\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"g\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"b\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structT[]\",\"name\":\"c\",\"type\":\"tuple[]\"}],\"internalType\":\"structTupleTest.S\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structT\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// TupleTestFuncSigs maps the 4-byte function signature to its string representation.
var TupleTestFuncSigs = map[string]string{
	"6f2be728": "f((uint256,uint256[],(uint256,uint256)[]),(uint256,uint256),uint256)",
	"e2179b8e": "g()",
}

// TupleTestBin is the compiled bytecode used for deploying new contracts.
var TupleTestBin = "0x608060405234801561001057600080fd5b50610437806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80636f2be7281461003b578063e2179b8e14610050575b600080fd5b61004e6100493660046101ac565b610070565b005b6100586100b0565b604051610067939291906102e3565b60405180910390f35b7f5490ef71cb8b76b3f0788e975052f0f43ba07bc17bae7ea289ee00e50b4d84858383836040516100a3939291906102e3565b60405180910390a1505050565b6100b86100c7565b6100c06100e8565b6000909192565b60405180606001604052806000815260200160608152602001606081525090565b604051806040016040528060008152602001600081525090565b600082601f830112610112578081fd5b8135610125610120826103d8565b6103b1565b818152915060208083019084810160408085028701830188101561014857600080fd5b60005b8581101561016f5761015d898461017b565b8552938301939181019160010161014b565b50505050505092915050565b60006040828403121561018c578081fd5b61019660406103b1565b9050813581526020820135602082015292915050565b6000806000608084860312156101c0578283fd5b833567ffffffffffffffff808211156101d7578485fd5b818601606081890312156101e9578586fd5b6101f360606103b1565b9250803583526020808201358381111561020b578788fd5b8083018a601f82011261021c578889fd5b8035915061022c610120836103d8565b82815283810190828501858502840186018e1015610248578b8cfd5b8b93505b8484101561026a57803583526001939093019291850191850161024c565b5087850152505050604082013583811115610283578788fd5b61028f8a828501610102565b6040860152508396506102a489828a0161017b565b955050505050606084013590509250925092565b60006102c483836102d4565b505060400190565b815260200190565b80518252602090810151910152565b60006080825260e0820185516080840152602080870151606060a086015282815161030e81866103f8565b9284019450859291505b8083101561033d5761032b8286516102cc565b91508385019450600183019250610318565b5060408901519350607f198682030160c0870152809150835161036081836103f8565b8695909350840191505b8085101561038f5761037d8383516102b8565b9250838201915060018501945061036a565b5050809350506103a1818501876102d4565b5050826060830152949350505050565b60405181810167ffffffffffffffff811182821017156103d057600080fd5b604052919050565b600067ffffffffffffffff8211156103ee578081fd5b5060209081020190565b9081526020019056fea2646970667358221220ee9a7ac23122a6ad55c39bfd06a9b1ee8a0663a6649e934329787416ec4262ac64736f6c63430006050033"

// DeployTupleTest deploys a new Ethereum contract, binding an instance of TupleTest to it.
func DeployTupleTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TupleTest, error) {
	parsed, err := abi.JSON(strings.NewReader(TupleTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TupleTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TupleTest{TupleTestCaller: TupleTestCaller{contract: contract}, TupleTestTransactor: TupleTestTransactor{contract: contract}, TupleTestFilterer: TupleTestFilterer{contract: contract}}, nil
}

// TupleTest is an auto generated Go binding around an Ethereum contract.
type TupleTest struct {
	TupleTestCaller     // Read-only binding to the contract
	TupleTestTransactor // Write-only binding to the contract
	TupleTestFilterer   // Log filterer for contract events
}

// TupleTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type TupleTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TupleTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TupleTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TupleTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TupleTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TupleTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TupleTestSession struct {
	Contract     *TupleTest        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TupleTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TupleTestCallerSession struct {
	Contract *TupleTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TupleTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TupleTestTransactorSession struct {
	Contract     *TupleTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TupleTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type TupleTestRaw struct {
	Contract *TupleTest // Generic contract binding to access the raw methods on
}

// TupleTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TupleTestCallerRaw struct {
	Contract *TupleTestCaller // Generic read-only contract binding to access the raw methods on
}

// TupleTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TupleTestTransactorRaw struct {
	Contract *TupleTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTupleTest creates a new instance of TupleTest, bound to a specific deployed contract.
func NewTupleTest(address common.Address, backend bind.ContractBackend) (*TupleTest, error) {
	contract, err := bindTupleTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TupleTest{TupleTestCaller: TupleTestCaller{contract: contract}, TupleTestTransactor: TupleTestTransactor{contract: contract}, TupleTestFilterer: TupleTestFilterer{contract: contract}}, nil
}

// NewTupleTestCaller creates a new read-only instance of TupleTest, bound to a specific deployed contract.
func NewTupleTestCaller(address common.Address, caller bind.ContractCaller) (*TupleTestCaller, error) {
	contract, err := bindTupleTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TupleTestCaller{contract: contract}, nil
}

// NewTupleTestTransactor creates a new write-only instance of TupleTest, bound to a specific deployed contract.
func NewTupleTestTransactor(address common.Address, transactor bind.ContractTransactor) (*TupleTestTransactor, error) {
	contract, err := bindTupleTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TupleTestTransactor{contract: contract}, nil
}

// NewTupleTestFilterer creates a new log filterer instance of TupleTest, bound to a specific deployed contract.
func NewTupleTestFilterer(address common.Address, filterer bind.ContractFilterer) (*TupleTestFilterer, error) {
	contract, err := bindTupleTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TupleTestFilterer{contract: contract}, nil
}

// bindTupleTest binds a generic wrapper to an already deployed contract.
func bindTupleTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TupleTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TupleTest *TupleTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TupleTest.Contract.TupleTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TupleTest *TupleTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TupleTest.Contract.TupleTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TupleTest *TupleTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TupleTest.Contract.TupleTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TupleTest *TupleTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TupleTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TupleTest *TupleTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TupleTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TupleTest *TupleTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TupleTest.Contract.contract.Transact(opts, method, params...)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() pure returns((uint256,uint256[],(uint256,uint256)[]), (uint256,uint256), uint256)
func (_TupleTest *TupleTestCaller) G(opts *bind.CallOpts) (TupleTestS, T, *big.Int, error) {
	var out []interface{}
	err := _TupleTest.contract.Call(opts, &out, "g")

	if err != nil {
		return *new(TupleTestS), *new(T), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(TupleTestS)).(*TupleTestS)
	out1 := *abi.ConvertType(out[1], new(T)).(*T)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() pure returns((uint256,uint256[],(uint256,uint256)[]), (uint256,uint256), uint256)
func (_TupleTest *TupleTestSession) G() (TupleTestS, T, *big.Int, error) {
	return _TupleTest.Contract.G(&_TupleTest.CallOpts)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() pure returns((uint256,uint256[],(uint256,uint256)[]), (uint256,uint256), uint256)
func (_TupleTest *TupleTestCallerSession) G() (TupleTestS, T, *big.Int, error) {
	return _TupleTest.Contract.G(&_TupleTest.CallOpts)
}

// F is a paid mutator transaction binding the contract method 0x6f2be728.
//
// Solidity: function f((uint256,uint256[],(uint256,uint256)[]) s, (uint256,uint256) t, uint256 u) returns()
func (_TupleTest *TupleTestTransactor) F(opts *bind.TransactOpts, s TupleTestS, t T, u *big.Int) (*types.Transaction, error) {
	return _TupleTest.contract.Transact(opts, "f", s, t, u)
}

// F is a paid mutator transaction binding the contract method 0x6f2be728.
//
// Solidity: function f((uint256,uint256[],(uint256,uint256)[]) s, (uint256,uint256) t, uint256 u) returns()
func (_TupleTest *TupleTestSession) F(s TupleTestS, t T, u *big.Int) (*types.Transaction, error) {
	return _TupleTest.Contract.F(&_TupleTest.TransactOpts, s, t, u)
}

// F is a paid mutator transaction binding the contract method 0x6f2be728.
//
// Solidity: function f((uint256,uint256[],(uint256,uint256)[]) s, (uint256,uint256) t, uint256 u) returns()
func (_TupleTest *TupleTestTransactorSession) F(s TupleTestS, t T, u *big.Int) (*types.Transaction, error) {
	return _TupleTest.Contract.F(&_TupleTest.TransactOpts, s, t, u)
}

// TupleTestEvFIterator is returned from FilterEvF and is used to iterate over the raw logs and unpacked data for EvF events raised by the TupleTest contract.
type TupleTestEvFIterator struct {
	Event *TupleTestEvF // Event containing the contract specifics and raw log

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
func (it *TupleTestEvFIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TupleTestEvF)
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
		it.Event = new(TupleTestEvF)
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
func (it *TupleTestEvFIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TupleTestEvFIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TupleTestEvF represents a EvF event raised by the TupleTest contract.
type TupleTestEvF struct {
	Arg0 TupleTestS
	Arg1 T
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEvF is a free log retrieval operation binding the contract event 0x5490ef71cb8b76b3f0788e975052f0f43ba07bc17bae7ea289ee00e50b4d8485.
//
// Solidity: event evF((uint256,uint256[],(uint256,uint256)[]) arg0, (uint256,uint256) arg1, uint256 arg2)
func (_TupleTest *TupleTestFilterer) FilterEvF(opts *bind.FilterOpts) (*TupleTestEvFIterator, error) {

	logs, sub, err := _TupleTest.contract.FilterLogs(opts, "evF")
	if err != nil {
		return nil, err
	}
	return &TupleTestEvFIterator{contract: _TupleTest.contract, event: "evF", logs: logs, sub: sub}, nil
}

// WatchEvF is a free log subscription operation binding the contract event 0x5490ef71cb8b76b3f0788e975052f0f43ba07bc17bae7ea289ee00e50b4d8485.
//
// Solidity: event evF((uint256,uint256[],(uint256,uint256)[]) arg0, (uint256,uint256) arg1, uint256 arg2)
func (_TupleTest *TupleTestFilterer) WatchEvF(opts *bind.WatchOpts, sink chan<- *TupleTestEvF) (event.Subscription, error) {

	logs, sub, err := _TupleTest.contract.WatchLogs(opts, "evF")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TupleTestEvF)
				if err := _TupleTest.contract.UnpackLog(event, "evF", log); err != nil {
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

// ParseEvF is a log parse operation binding the contract event 0x5490ef71cb8b76b3f0788e975052f0f43ba07bc17bae7ea289ee00e50b4d8485.
//
// Solidity: event evF((uint256,uint256[],(uint256,uint256)[]) arg0, (uint256,uint256) arg1, uint256 arg2)
func (_TupleTest *TupleTestFilterer) ParseEvF(log types.Log) (*TupleTestEvF, error) {
	event := new(TupleTestEvF)
	if err := _TupleTest.contract.UnpackLog(event, "evF", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TupleTest2ABI is the input ABI used to generate the binding from.
const TupleTest2ABI = "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"b\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structT[]\",\"name\":\"c\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structTupleTest2.S\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structT\",\"name\":\"\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"evF\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structT\",\"name\":\"t\",\"type\":\"tuple\"}],\"name\":\"a\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structT\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structT[]\",\"name\":\"t\",\"type\":\"tuple[]\"}],\"name\":\"b\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structT[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"b\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structT[]\",\"name\":\"c\",\"type\":\"tuple[]\"}],\"internalType\":\"structTupleTest2.S\",\"name\":\"s\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structT\",\"name\":\"t\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"u\",\"type\":\"uint256\"}],\"name\":\"f\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"g\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"b\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structT[]\",\"name\":\"c\",\"type\":\"tuple[]\"}],\"internalType\":\"structTupleTest2.S\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structT\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"b\",\"type\":\"uint256[]\"}],\"internalType\":\"structTupleTest2.A\",\"name\":\"a\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"method\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TupleTest2FuncSigs maps the 4-byte function signature to its string representation.
var TupleTest2FuncSigs = map[string]string{
	"7705ed20": "a((uint256,uint256))",
	"9346002a": "b((uint256,uint256)[])",
	"6f2be728": "f((uint256,uint256[],(uint256,uint256)[]),(uint256,uint256),uint256)",
	"e2179b8e": "g()",
	"ab41fe7a": "method((uint256,uint256[]),uint256)",
}

// TupleTest2Bin is the compiled bytecode used for deploying new contracts.
var TupleTest2Bin = "0x608060405234801561001057600080fd5b506106b7806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80636f2be7281461005c5780637705ed20146100715780639346002a1461009a578063ab41fe7a146100ba578063e2179b8e146100cd575b600080fd5b61006f61006a36600461043f565b6100e4565b005b61008461007f3660046104f5565b6100e9565b604051610091919061061d565b60405180910390f35b6100ad6100a8366004610382565b6100f5565b6040516100919190610572565b61006f6100c83660046103bd565b6100f8565b6100d56100fc565b60405161009193929190610585565b505050565b6100f1610234565b5090565b90565b5050565b61010461024e565b61010c610234565b6000610116610234565b506040805180820182526001815260026020820152815160038082526080820190935290916060919081602001602082028036833701905050905060018160008151811061016057fe5b60200260200101818152505060028160018151811061017b57fe5b60200260200101818152505060038160028151811061019657fe5b60209081029190910101526040805160028082526060828101909352816020015b6101bf610234565b8152602001906001900390816101b757905050905082816000815181106101e257fe5b602002602001018190525082816001815181106101fb57fe5b602002602001018190525061020e61024e565b506040805160608101825260018152602081019390935282015294909350600392509050565b604051806040016040528060008152602001600081525090565b60405180606001604052806000815260200160608152602001606081525090565b600082601f83011261027f578081fd5b813561029261028d82610658565b610631565b81815291506020808301908481016040808502870183018810156102b557600080fd5b60005b858110156102dc576102ca8984610351565b855293830193918101916001016102b8565b50505050505092915050565b600082601f8301126102f8578081fd5b813561030661028d82610658565b81815291506020808301908481018184028601820187101561032757600080fd5b60005b848110156103465781358452928201929082019060010161032a565b505050505092915050565b600060408284031215610362578081fd5b61036c6040610631565b9050813581526020820135602082015292915050565b600060208284031215610393578081fd5b813567ffffffffffffffff8111156103a9578182fd5b6103b58482850161026f565b949350505050565b600080604083850312156103cf578081fd5b823567ffffffffffffffff808211156103e6578283fd5b818501604081880312156103f8578384fd5b6104026040610631565b925080358352602081013582811115610419578485fd5b610425888284016102e8565b602085015250505080925050602083013590509250929050565b600080600060808486031215610453578081fd5b833567ffffffffffffffff8082111561046a578283fd5b8186016060818903121561047c578384fd5b6104866060610631565b92508035835260208101358281111561049d578485fd5b6104a9898284016102e8565b6020850152506040810135828111156104c0578485fd5b6104cc8982840161026f565b6040850152505050809350506104e58560208601610351565b9150606084013590509250925092565b600060408284031215610506578081fd5b6105108383610351565b9392505050565b815260200190565b6000815180845260208085019450808401835b8381101561055857610545878351610563565b6040969096019590820190600101610532565b509495945050505050565b80518252602090810151910152565b600060208252610510602083018461051f565b60006080825260e0820185516080840152602080870151606060a08601528281516105b08186610678565b9284019450859291505b808310156105df576105cd828651610517565b915083850194506001830192506105ba565b506040890151868203607f190160c088015293506105fd818561051f565b9450505061060d81850187610563565b5050826060830152949350505050565b6040810161062b8284610563565b92915050565b60405181810167ffffffffffffffff8111828210171561065057600080fd5b604052919050565b600067ffffffffffffffff82111561066e578081fd5b5060209081020190565b9081526020019056fea26469706673582212209aebd32fbbdfa65d7bf142b42fe5d51646f0123926374286ab3db59e8a4968f164736f6c63430006050033"

// DeployTupleTest2 deploys a new Ethereum contract, binding an instance of TupleTest2 to it.
func DeployTupleTest2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TupleTest2, error) {
	parsed, err := abi.JSON(strings.NewReader(TupleTest2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TupleTest2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TupleTest2{TupleTest2Caller: TupleTest2Caller{contract: contract}, TupleTest2Transactor: TupleTest2Transactor{contract: contract}, TupleTest2Filterer: TupleTest2Filterer{contract: contract}}, nil
}

// TupleTest2 is an auto generated Go binding around an Ethereum contract.
type TupleTest2 struct {
	TupleTest2Caller     // Read-only binding to the contract
	TupleTest2Transactor // Write-only binding to the contract
	TupleTest2Filterer   // Log filterer for contract events
}

// TupleTest2Caller is an auto generated read-only Go binding around an Ethereum contract.
type TupleTest2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TupleTest2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TupleTest2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TupleTest2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TupleTest2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TupleTest2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TupleTest2Session struct {
	Contract     *TupleTest2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TupleTest2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TupleTest2CallerSession struct {
	Contract *TupleTest2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TupleTest2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TupleTest2TransactorSession struct {
	Contract     *TupleTest2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TupleTest2Raw is an auto generated low-level Go binding around an Ethereum contract.
type TupleTest2Raw struct {
	Contract *TupleTest2 // Generic contract binding to access the raw methods on
}

// TupleTest2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TupleTest2CallerRaw struct {
	Contract *TupleTest2Caller // Generic read-only contract binding to access the raw methods on
}

// TupleTest2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TupleTest2TransactorRaw struct {
	Contract *TupleTest2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTupleTest2 creates a new instance of TupleTest2, bound to a specific deployed contract.
func NewTupleTest2(address common.Address, backend bind.ContractBackend) (*TupleTest2, error) {
	contract, err := bindTupleTest2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TupleTest2{TupleTest2Caller: TupleTest2Caller{contract: contract}, TupleTest2Transactor: TupleTest2Transactor{contract: contract}, TupleTest2Filterer: TupleTest2Filterer{contract: contract}}, nil
}

// NewTupleTest2Caller creates a new read-only instance of TupleTest2, bound to a specific deployed contract.
func NewTupleTest2Caller(address common.Address, caller bind.ContractCaller) (*TupleTest2Caller, error) {
	contract, err := bindTupleTest2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TupleTest2Caller{contract: contract}, nil
}

// NewTupleTest2Transactor creates a new write-only instance of TupleTest2, bound to a specific deployed contract.
func NewTupleTest2Transactor(address common.Address, transactor bind.ContractTransactor) (*TupleTest2Transactor, error) {
	contract, err := bindTupleTest2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TupleTest2Transactor{contract: contract}, nil
}

// NewTupleTest2Filterer creates a new log filterer instance of TupleTest2, bound to a specific deployed contract.
func NewTupleTest2Filterer(address common.Address, filterer bind.ContractFilterer) (*TupleTest2Filterer, error) {
	contract, err := bindTupleTest2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TupleTest2Filterer{contract: contract}, nil
}

// bindTupleTest2 binds a generic wrapper to an already deployed contract.
func bindTupleTest2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TupleTest2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TupleTest2 *TupleTest2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TupleTest2.Contract.TupleTest2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TupleTest2 *TupleTest2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TupleTest2.Contract.TupleTest2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TupleTest2 *TupleTest2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TupleTest2.Contract.TupleTest2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TupleTest2 *TupleTest2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TupleTest2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TupleTest2 *TupleTest2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TupleTest2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TupleTest2 *TupleTest2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TupleTest2.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0x7705ed20.
//
// Solidity: function a((uint256,uint256) t) view returns((uint256,uint256))
func (_TupleTest2 *TupleTest2Caller) A(opts *bind.CallOpts, t T) (T, error) {
	var out []interface{}
	err := _TupleTest2.contract.Call(opts, &out, "a", t)

	if err != nil {
		return *new(T), err
	}

	out0 := *abi.ConvertType(out[0], new(T)).(*T)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0x7705ed20.
//
// Solidity: function a((uint256,uint256) t) view returns((uint256,uint256))
func (_TupleTest2 *TupleTest2Session) A(t T) (T, error) {
	return _TupleTest2.Contract.A(&_TupleTest2.CallOpts, t)
}

// A is a free data retrieval call binding the contract method 0x7705ed20.
//
// Solidity: function a((uint256,uint256) t) view returns((uint256,uint256))
func (_TupleTest2 *TupleTest2CallerSession) A(t T) (T, error) {
	return _TupleTest2.Contract.A(&_TupleTest2.CallOpts, t)
}

// B is a free data retrieval call binding the contract method 0x9346002a.
//
// Solidity: function b((uint256,uint256)[] t) view returns((uint256,uint256)[])
func (_TupleTest2 *TupleTest2Caller) B(opts *bind.CallOpts, t []T) ([]T, error) {
	var out []interface{}
	err := _TupleTest2.contract.Call(opts, &out, "b", t)

	if err != nil {
		return *new([]T), err
	}

	out0 := *abi.ConvertType(out[0], new([]T)).(*[]T)

	return out0, err

}

// B is a free data retrieval call binding the contract method 0x9346002a.
//
// Solidity: function b((uint256,uint256)[] t) view returns((uint256,uint256)[])
func (_TupleTest2 *TupleTest2Session) B(t []T) ([]T, error) {
	return _TupleTest2.Contract.B(&_TupleTest2.CallOpts, t)
}

// B is a free data retrieval call binding the contract method 0x9346002a.
//
// Solidity: function b((uint256,uint256)[] t) view returns((uint256,uint256)[])
func (_TupleTest2 *TupleTest2CallerSession) B(t []T) ([]T, error) {
	return _TupleTest2.Contract.B(&_TupleTest2.CallOpts, t)
}

// F is a free data retrieval call binding the contract method 0x6f2be728.
//
// Solidity: function f((uint256,uint256[],(uint256,uint256)[]) s, (uint256,uint256) t, uint256 u) view returns()
func (_TupleTest2 *TupleTest2Caller) F(opts *bind.CallOpts, s TupleTest2S, t T, u *big.Int) error {
	var out []interface{}
	err := _TupleTest2.contract.Call(opts, &out, "f", s, t, u)

	if err != nil {
		return err
	}

	return err

}

// F is a free data retrieval call binding the contract method 0x6f2be728.
//
// Solidity: function f((uint256,uint256[],(uint256,uint256)[]) s, (uint256,uint256) t, uint256 u) view returns()
func (_TupleTest2 *TupleTest2Session) F(s TupleTest2S, t T, u *big.Int) error {
	return _TupleTest2.Contract.F(&_TupleTest2.CallOpts, s, t, u)
}

// F is a free data retrieval call binding the contract method 0x6f2be728.
//
// Solidity: function f((uint256,uint256[],(uint256,uint256)[]) s, (uint256,uint256) t, uint256 u) view returns()
func (_TupleTest2 *TupleTest2CallerSession) F(s TupleTest2S, t T, u *big.Int) error {
	return _TupleTest2.Contract.F(&_TupleTest2.CallOpts, s, t, u)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() view returns((uint256,uint256[],(uint256,uint256)[]), (uint256,uint256), uint256)
func (_TupleTest2 *TupleTest2Caller) G(opts *bind.CallOpts) (TupleTest2S, T, *big.Int, error) {
	var out []interface{}
	err := _TupleTest2.contract.Call(opts, &out, "g")

	if err != nil {
		return *new(TupleTest2S), *new(T), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(TupleTest2S)).(*TupleTest2S)
	out1 := *abi.ConvertType(out[1], new(T)).(*T)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() view returns((uint256,uint256[],(uint256,uint256)[]), (uint256,uint256), uint256)
func (_TupleTest2 *TupleTest2Session) G() (TupleTest2S, T, *big.Int, error) {
	return _TupleTest2.Contract.G(&_TupleTest2.CallOpts)
}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() view returns((uint256,uint256[],(uint256,uint256)[]), (uint256,uint256), uint256)
func (_TupleTest2 *TupleTest2CallerSession) G() (TupleTest2S, T, *big.Int, error) {
	return _TupleTest2.Contract.G(&_TupleTest2.CallOpts)
}

// Method is a free data retrieval call binding the contract method 0xab41fe7a.
//
// Solidity: function method((uint256,uint256[]) a, uint256 b) view returns()
func (_TupleTest2 *TupleTest2Caller) Method(opts *bind.CallOpts, a TupleTest2A, b *big.Int) error {
	var out []interface{}
	err := _TupleTest2.contract.Call(opts, &out, "method", a, b)

	if err != nil {
		return err
	}

	return err

}

// Method is a free data retrieval call binding the contract method 0xab41fe7a.
//
// Solidity: function method((uint256,uint256[]) a, uint256 b) view returns()
func (_TupleTest2 *TupleTest2Session) Method(a TupleTest2A, b *big.Int) error {
	return _TupleTest2.Contract.Method(&_TupleTest2.CallOpts, a, b)
}

// Method is a free data retrieval call binding the contract method 0xab41fe7a.
//
// Solidity: function method((uint256,uint256[]) a, uint256 b) view returns()
func (_TupleTest2 *TupleTest2CallerSession) Method(a TupleTest2A, b *big.Int) error {
	return _TupleTest2.Contract.Method(&_TupleTest2.CallOpts, a, b)
}

// TupleTest2EvFIterator is returned from FilterEvF and is used to iterate over the raw logs and unpacked data for EvF events raised by the TupleTest2 contract.
type TupleTest2EvFIterator struct {
	Event *TupleTest2EvF // Event containing the contract specifics and raw log

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
func (it *TupleTest2EvFIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TupleTest2EvF)
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
		it.Event = new(TupleTest2EvF)
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
func (it *TupleTest2EvFIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TupleTest2EvFIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TupleTest2EvF represents a EvF event raised by the TupleTest2 contract.
type TupleTest2EvF struct {
	Arg0 TupleTest2S
	Arg1 T
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEvF is a free log retrieval operation binding the contract event 0x5490ef71cb8b76b3f0788e975052f0f43ba07bc17bae7ea289ee00e50b4d8485.
//
// Solidity: event evF((uint256,uint256[],(uint256,uint256)[]) arg0, (uint256,uint256) arg1, uint256 arg2)
func (_TupleTest2 *TupleTest2Filterer) FilterEvF(opts *bind.FilterOpts) (*TupleTest2EvFIterator, error) {

	logs, sub, err := _TupleTest2.contract.FilterLogs(opts, "evF")
	if err != nil {
		return nil, err
	}
	return &TupleTest2EvFIterator{contract: _TupleTest2.contract, event: "evF", logs: logs, sub: sub}, nil
}

// WatchEvF is a free log subscription operation binding the contract event 0x5490ef71cb8b76b3f0788e975052f0f43ba07bc17bae7ea289ee00e50b4d8485.
//
// Solidity: event evF((uint256,uint256[],(uint256,uint256)[]) arg0, (uint256,uint256) arg1, uint256 arg2)
func (_TupleTest2 *TupleTest2Filterer) WatchEvF(opts *bind.WatchOpts, sink chan<- *TupleTest2EvF) (event.Subscription, error) {

	logs, sub, err := _TupleTest2.contract.WatchLogs(opts, "evF")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TupleTest2EvF)
				if err := _TupleTest2.contract.UnpackLog(event, "evF", log); err != nil {
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

// ParseEvF is a log parse operation binding the contract event 0x5490ef71cb8b76b3f0788e975052f0f43ba07bc17bae7ea289ee00e50b4d8485.
//
// Solidity: event evF((uint256,uint256[],(uint256,uint256)[]) arg0, (uint256,uint256) arg1, uint256 arg2)
func (_TupleTest2 *TupleTest2Filterer) ParseEvF(log types.Log) (*TupleTest2EvF, error) {
	event := new(TupleTest2EvF)
	if err := _TupleTest2.contract.UnpackLog(event, "evF", log); err != nil {
		return nil, err
	}
	return event, nil
}

// VoidTestABI is the input ABI used to generate the binding from.
const VoidTestABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"method\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VoidTestFuncSigs maps the 4-byte function signature to its string representation.
var VoidTestFuncSigs = map[string]string{
	"9a40e3f6": "method(uint256)",
}

// VoidTestBin is the compiled bytecode used for deploying new contracts.
var VoidTestBin = "0x6080604052348015600f57600080fd5b50608e8061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80639a40e3f614602d575b600080fd5b603c60383660046041565b603e565b005b50565b6000602082840312156051578081fd5b503591905056fea2646970667358221220625eb4d9417df0e5db175f3180458cf2899e8f4054d555f5337613c00a80fb3f64736f6c63430006050033"

// DeployVoidTest deploys a new Ethereum contract, binding an instance of VoidTest to it.
func DeployVoidTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VoidTest, error) {
	parsed, err := abi.JSON(strings.NewReader(VoidTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VoidTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VoidTest{VoidTestCaller: VoidTestCaller{contract: contract}, VoidTestTransactor: VoidTestTransactor{contract: contract}, VoidTestFilterer: VoidTestFilterer{contract: contract}}, nil
}

// VoidTest is an auto generated Go binding around an Ethereum contract.
type VoidTest struct {
	VoidTestCaller     // Read-only binding to the contract
	VoidTestTransactor // Write-only binding to the contract
	VoidTestFilterer   // Log filterer for contract events
}

// VoidTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoidTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoidTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoidTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoidTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoidTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoidTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoidTestSession struct {
	Contract     *VoidTest         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoidTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoidTestCallerSession struct {
	Contract *VoidTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// VoidTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoidTestTransactorSession struct {
	Contract     *VoidTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VoidTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoidTestRaw struct {
	Contract *VoidTest // Generic contract binding to access the raw methods on
}

// VoidTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoidTestCallerRaw struct {
	Contract *VoidTestCaller // Generic read-only contract binding to access the raw methods on
}

// VoidTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoidTestTransactorRaw struct {
	Contract *VoidTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoidTest creates a new instance of VoidTest, bound to a specific deployed contract.
func NewVoidTest(address common.Address, backend bind.ContractBackend) (*VoidTest, error) {
	contract, err := bindVoidTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VoidTest{VoidTestCaller: VoidTestCaller{contract: contract}, VoidTestTransactor: VoidTestTransactor{contract: contract}, VoidTestFilterer: VoidTestFilterer{contract: contract}}, nil
}

// NewVoidTestCaller creates a new read-only instance of VoidTest, bound to a specific deployed contract.
func NewVoidTestCaller(address common.Address, caller bind.ContractCaller) (*VoidTestCaller, error) {
	contract, err := bindVoidTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoidTestCaller{contract: contract}, nil
}

// NewVoidTestTransactor creates a new write-only instance of VoidTest, bound to a specific deployed contract.
func NewVoidTestTransactor(address common.Address, transactor bind.ContractTransactor) (*VoidTestTransactor, error) {
	contract, err := bindVoidTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoidTestTransactor{contract: contract}, nil
}

// NewVoidTestFilterer creates a new log filterer instance of VoidTest, bound to a specific deployed contract.
func NewVoidTestFilterer(address common.Address, filterer bind.ContractFilterer) (*VoidTestFilterer, error) {
	contract, err := bindVoidTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoidTestFilterer{contract: contract}, nil
}

// bindVoidTest binds a generic wrapper to an already deployed contract.
func bindVoidTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VoidTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoidTest *VoidTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoidTest.Contract.VoidTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoidTest *VoidTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoidTest.Contract.VoidTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoidTest *VoidTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoidTest.Contract.VoidTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoidTest *VoidTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoidTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoidTest *VoidTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoidTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoidTest *VoidTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoidTest.Contract.contract.Transact(opts, method, params...)
}

// Method is a free data retrieval call binding the contract method 0x9a40e3f6.
//
// Solidity: function method(uint256 a) view returns()
func (_VoidTest *VoidTestCaller) Method(opts *bind.CallOpts, a *big.Int) error {
	var out []interface{}
	err := _VoidTest.contract.Call(opts, &out, "method", a)

	if err != nil {
		return err
	}

	return err

}

// Method is a free data retrieval call binding the contract method 0x9a40e3f6.
//
// Solidity: function method(uint256 a) view returns()
func (_VoidTest *VoidTestSession) Method(a *big.Int) error {
	return _VoidTest.Contract.Method(&_VoidTest.CallOpts, a)
}

// Method is a free data retrieval call binding the contract method 0x9a40e3f6.
//
// Solidity: function method(uint256 a) view returns()
func (_VoidTest *VoidTestCallerSession) Method(a *big.Int) error {
	return _VoidTest.Contract.Method(&_VoidTest.CallOpts, a)
}
