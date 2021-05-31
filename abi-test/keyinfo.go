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

// KeyInfoABI is the input ABI used to generate the binding from.
const KeyInfoABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"user\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"party1\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"party2\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"KeyShareUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"newPubkey\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"user\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PubkeyForwarded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"user\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"party1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"party2\",\"type\":\"bytes\"}],\"name\":\"addKeyShareInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newPubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"user\",\"type\":\"bytes32\"}],\"name\":\"forwardPubkey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"pubkeyForwards\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"pubkeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"party1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"party2\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userkeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// KeyInfoFuncSigs maps the 4-byte function signature to its string representation.
var KeyInfoFuncSigs = map[string]string{
	"07873ddc": "addKeyShareInfo(bytes32,bytes,bytes,bytes)",
	"b8b64181": "forwardPubkey(bytes,bytes,bytes32)",
	"4e6a698d": "pubkeyForwards(bytes)",
	"2ac8597c": "pubkeys(bytes)",
	"545239fe": "userkeys(bytes32,uint256)",
}

// KeyInfoBin is the compiled bytecode used for deploying new contracts.
var KeyInfoBin = "0x608060405234801561001057600080fd5b50610d20806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806307873ddc1461005c5780632ac8597c146102255780634e6a698d146103ae578063545239fe146104c7578063b8b64181146104ea575b600080fd5b6102116004803603608081101561007257600080fd5b81359190810190604081016020820135600160201b81111561009357600080fd5b8201836020820111156100a557600080fd5b803590602001918460018302840111600160201b831117156100c657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561011857600080fd5b82018360208201111561012a57600080fd5b803590602001918460018302840111600160201b8311171561014b57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561019d57600080fd5b8201836020820111156101af57600080fd5b803590602001918460018302840111600160201b831117156101d057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610615945050505050565b604080519115158252519081900360200190f35b6102c96004803603602081101561023b57600080fd5b810190602081018135600160201b81111561025557600080fd5b82018360208201111561026757600080fd5b803590602001918460018302840111600160201b8311171561028857600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061086f945050505050565b604051808060200180602001848152602001838103835286818151815260200191508051906020019080838360005b838110156103105781810151838201526020016102f8565b50505050905090810190601f16801561033d5780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b83811015610370578181015183820152602001610358565b50505050905090810190601f16801561039d5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b610452600480360360208110156103c457600080fd5b810190602081018135600160201b8111156103de57600080fd5b8201836020820111156103f057600080fd5b803590602001918460018302840111600160201b8311171561041157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506109c0945050505050565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561048c578181015183820152602001610474565b50505050905090810190601f1680156104b95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610452600480360360408110156104dd57600080fd5b5080359060200135610a66565b6102116004803603606081101561050057600080fd5b810190602081018135600160201b81111561051a57600080fd5b82018360208201111561052c57600080fd5b803590602001918460018302840111600160201b8311171561054d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561059f57600080fd5b8201836020820111156105b157600080fd5b803590602001918460018302840111600160201b831117156105d257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610ae7915050565b6000836040518060600160405280858152602001848152602001428152506000866040518082805190602001908083835b602083106106655780518252601f199092019160209182019101610646565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381019093208451805191946106a694508593500190610c4f565b5060208281015180516106bf9260018501920190610c4f565b506040918201516002909101556000878152600160208181529282208054918201815582529082902087516106fc93919092019190880190610c4f565b5085856040518082805190602001908083835b6020831061072e5780518252601f19909201916020918201910161070f565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207f909956b054ea2ed1b865ac0e9bf20662c97d78d97205883668215fa0f7ca1b8e868642604051808060200180602001848152602001838103835286818151815260200191508051906020019080838360005b838110156107c65781810151838201526020016107ae565b50505050905090810190601f1680156107f35780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b8381101561082657818101518382015260200161080e565b50505050905090810190601f1680156108535780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a350600195945050505050565b805160208183018101805160008252928201938201939093209190925280546040805160026001841615610100026000190190931692909204601f81018590048502830185019091528082529192909183918301828280156109125780601f106108e757610100808354040283529160200191610912565b820191906000526020600020905b8154815290600101906020018083116108f557829003601f168201915b505050505090806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109b05780601f10610985576101008083540402835291602001916109b0565b820191906000526020600020905b81548152906001019060200180831161099357829003601f168201915b5050505050908060020154905083565b80516020818301810180516002808352938301948301949094209390528254604080516001831615610100026000190190921693909304601f8101839004830282018301909352828152929190830182828015610a5e5780601f10610a3357610100808354040283529160200191610a5e565b820191906000526020600020905b815481529060010190602001808311610a4157829003601f168201915b505050505081565b60016020528160005260406000208181548110610a7f57fe5b600091825260209182902001805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152945090925090830182828015610a5e5780601f10610a3357610100808354040283529160200191610a5e565b600083836002866040518082805190602001908083835b60208310610b1d5780518252601f199092019160209182019101610afe565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381019093208451610b5e9591949190910192509050610c4f565b5082846040518082805190602001908083835b60208310610b905780518252601f199092019160209182019101610b71565b51815160209384036101000a60001901801990921691161790526040519190930181900381208b519095508b945090928392508401908083835b60208310610be95780518252601f199092019160209182019101610bca565b51815160209384036101000a60001901801990921691161790526040805192909401829003822042835293519395507f439b679fc9f542ae812b84dc519a97855cda3b311b9b926d034e924bdf0cf1a094509083900301919050a4506001949350505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610c9057805160ff1916838001178555610cbd565b82800160010185558215610cbd579182015b82811115610cbd578251825591602001919060010190610ca2565b50610cc9929150610ccd565b5090565b610ce791905b80821115610cc95760008155600101610cd3565b9056fea26469706673582212203dba6e37dddc0623f6ea55ec4392e0b85260e2c1ba3bebc44b69716fb408ddd164736f6c63430006050033"

// DeployKeyInfo deploys a new Ethereum contract, binding an instance of KeyInfo to it.
func DeployKeyInfo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KeyInfo, error) {
	parsed, err := abi.JSON(strings.NewReader(KeyInfoABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KeyInfoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KeyInfo{KeyInfoCaller: KeyInfoCaller{contract: contract}, KeyInfoTransactor: KeyInfoTransactor{contract: contract}, KeyInfoFilterer: KeyInfoFilterer{contract: contract}}, nil
}

// KeyInfo is an auto generated Go binding around an Ethereum contract.
type KeyInfo struct {
	KeyInfoCaller     // Read-only binding to the contract
	KeyInfoTransactor // Write-only binding to the contract
	KeyInfoFilterer   // Log filterer for contract events
}

// KeyInfoCaller is an auto generated read-only Go binding around an Ethereum contract.
type KeyInfoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyInfoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KeyInfoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyInfoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KeyInfoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KeyInfoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KeyInfoSession struct {
	Contract     *KeyInfo          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KeyInfoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KeyInfoCallerSession struct {
	Contract *KeyInfoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// KeyInfoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KeyInfoTransactorSession struct {
	Contract     *KeyInfoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// KeyInfoRaw is an auto generated low-level Go binding around an Ethereum contract.
type KeyInfoRaw struct {
	Contract *KeyInfo // Generic contract binding to access the raw methods on
}

// KeyInfoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KeyInfoCallerRaw struct {
	Contract *KeyInfoCaller // Generic read-only contract binding to access the raw methods on
}

// KeyInfoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KeyInfoTransactorRaw struct {
	Contract *KeyInfoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKeyInfo creates a new instance of KeyInfo, bound to a specific deployed contract.
func NewKeyInfo(address common.Address, backend bind.ContractBackend) (*KeyInfo, error) {
	contract, err := bindKeyInfo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeyInfo{KeyInfoCaller: KeyInfoCaller{contract: contract}, KeyInfoTransactor: KeyInfoTransactor{contract: contract}, KeyInfoFilterer: KeyInfoFilterer{contract: contract}}, nil
}

// NewKeyInfoCaller creates a new read-only instance of KeyInfo, bound to a specific deployed contract.
func NewKeyInfoCaller(address common.Address, caller bind.ContractCaller) (*KeyInfoCaller, error) {
	contract, err := bindKeyInfo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeyInfoCaller{contract: contract}, nil
}

// NewKeyInfoTransactor creates a new write-only instance of KeyInfo, bound to a specific deployed contract.
func NewKeyInfoTransactor(address common.Address, transactor bind.ContractTransactor) (*KeyInfoTransactor, error) {
	contract, err := bindKeyInfo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeyInfoTransactor{contract: contract}, nil
}

// NewKeyInfoFilterer creates a new log filterer instance of KeyInfo, bound to a specific deployed contract.
func NewKeyInfoFilterer(address common.Address, filterer bind.ContractFilterer) (*KeyInfoFilterer, error) {
	contract, err := bindKeyInfo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeyInfoFilterer{contract: contract}, nil
}

// bindKeyInfo binds a generic wrapper to an already deployed contract.
func bindKeyInfo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KeyInfoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyInfo *KeyInfoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyInfo.Contract.KeyInfoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyInfo *KeyInfoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyInfo.Contract.KeyInfoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyInfo *KeyInfoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyInfo.Contract.KeyInfoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KeyInfo *KeyInfoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeyInfo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KeyInfo *KeyInfoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeyInfo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KeyInfo *KeyInfoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeyInfo.Contract.contract.Transact(opts, method, params...)
}

// PubkeyForwards is a free data retrieval call binding the contract method 0x4e6a698d.
//
// Solidity: function pubkeyForwards(bytes ) view returns(bytes)
func (_KeyInfo *KeyInfoCaller) PubkeyForwards(opts *bind.CallOpts, arg0 []byte) ([]byte, error) {
	var out []interface{}
	err := _KeyInfo.contract.Call(opts, &out, "pubkeyForwards", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PubkeyForwards is a free data retrieval call binding the contract method 0x4e6a698d.
//
// Solidity: function pubkeyForwards(bytes ) view returns(bytes)
func (_KeyInfo *KeyInfoSession) PubkeyForwards(arg0 []byte) ([]byte, error) {
	return _KeyInfo.Contract.PubkeyForwards(&_KeyInfo.CallOpts, arg0)
}

// PubkeyForwards is a free data retrieval call binding the contract method 0x4e6a698d.
//
// Solidity: function pubkeyForwards(bytes ) view returns(bytes)
func (_KeyInfo *KeyInfoCallerSession) PubkeyForwards(arg0 []byte) ([]byte, error) {
	return _KeyInfo.Contract.PubkeyForwards(&_KeyInfo.CallOpts, arg0)
}

// Pubkeys is a free data retrieval call binding the contract method 0x2ac8597c.
//
// Solidity: function pubkeys(bytes ) view returns(bytes party1, bytes party2, uint256 timestamp)
func (_KeyInfo *KeyInfoCaller) Pubkeys(opts *bind.CallOpts, arg0 []byte) (struct {
	Party1    []byte
	Party2    []byte
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _KeyInfo.contract.Call(opts, &out, "pubkeys", arg0)

	outstruct := new(struct {
		Party1    []byte
		Party2    []byte
		Timestamp *big.Int
	})

	outstruct.Party1 = out[0].([]byte)
	outstruct.Party2 = out[1].([]byte)
	outstruct.Timestamp = out[2].(*big.Int)

	return *outstruct, err

}

// Pubkeys is a free data retrieval call binding the contract method 0x2ac8597c.
//
// Solidity: function pubkeys(bytes ) view returns(bytes party1, bytes party2, uint256 timestamp)
func (_KeyInfo *KeyInfoSession) Pubkeys(arg0 []byte) (struct {
	Party1    []byte
	Party2    []byte
	Timestamp *big.Int
}, error) {
	return _KeyInfo.Contract.Pubkeys(&_KeyInfo.CallOpts, arg0)
}

// Pubkeys is a free data retrieval call binding the contract method 0x2ac8597c.
//
// Solidity: function pubkeys(bytes ) view returns(bytes party1, bytes party2, uint256 timestamp)
func (_KeyInfo *KeyInfoCallerSession) Pubkeys(arg0 []byte) (struct {
	Party1    []byte
	Party2    []byte
	Timestamp *big.Int
}, error) {
	return _KeyInfo.Contract.Pubkeys(&_KeyInfo.CallOpts, arg0)
}

// Userkeys is a free data retrieval call binding the contract method 0x545239fe.
//
// Solidity: function userkeys(bytes32 , uint256 ) view returns(bytes)
func (_KeyInfo *KeyInfoCaller) Userkeys(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _KeyInfo.contract.Call(opts, &out, "userkeys", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Userkeys is a free data retrieval call binding the contract method 0x545239fe.
//
// Solidity: function userkeys(bytes32 , uint256 ) view returns(bytes)
func (_KeyInfo *KeyInfoSession) Userkeys(arg0 [32]byte, arg1 *big.Int) ([]byte, error) {
	return _KeyInfo.Contract.Userkeys(&_KeyInfo.CallOpts, arg0, arg1)
}

// Userkeys is a free data retrieval call binding the contract method 0x545239fe.
//
// Solidity: function userkeys(bytes32 , uint256 ) view returns(bytes)
func (_KeyInfo *KeyInfoCallerSession) Userkeys(arg0 [32]byte, arg1 *big.Int) ([]byte, error) {
	return _KeyInfo.Contract.Userkeys(&_KeyInfo.CallOpts, arg0, arg1)
}

// AddKeyShareInfo is a paid mutator transaction binding the contract method 0x07873ddc.
//
// Solidity: function addKeyShareInfo(bytes32 user, bytes pubkey, bytes party1, bytes party2) returns(bool success)
func (_KeyInfo *KeyInfoTransactor) AddKeyShareInfo(opts *bind.TransactOpts, user [32]byte, pubkey []byte, party1 []byte, party2 []byte) (*types.Transaction, error) {
	return _KeyInfo.contract.Transact(opts, "addKeyShareInfo", user, pubkey, party1, party2)
}

// AddKeyShareInfo is a paid mutator transaction binding the contract method 0x07873ddc.
//
// Solidity: function addKeyShareInfo(bytes32 user, bytes pubkey, bytes party1, bytes party2) returns(bool success)
func (_KeyInfo *KeyInfoSession) AddKeyShareInfo(user [32]byte, pubkey []byte, party1 []byte, party2 []byte) (*types.Transaction, error) {
	return _KeyInfo.Contract.AddKeyShareInfo(&_KeyInfo.TransactOpts, user, pubkey, party1, party2)
}

// AddKeyShareInfo is a paid mutator transaction binding the contract method 0x07873ddc.
//
// Solidity: function addKeyShareInfo(bytes32 user, bytes pubkey, bytes party1, bytes party2) returns(bool success)
func (_KeyInfo *KeyInfoTransactorSession) AddKeyShareInfo(user [32]byte, pubkey []byte, party1 []byte, party2 []byte) (*types.Transaction, error) {
	return _KeyInfo.Contract.AddKeyShareInfo(&_KeyInfo.TransactOpts, user, pubkey, party1, party2)
}

// ForwardPubkey is a paid mutator transaction binding the contract method 0xb8b64181.
//
// Solidity: function forwardPubkey(bytes pubkey, bytes newPubkey, bytes32 user) returns(bool success)
func (_KeyInfo *KeyInfoTransactor) ForwardPubkey(opts *bind.TransactOpts, pubkey []byte, newPubkey []byte, user [32]byte) (*types.Transaction, error) {
	return _KeyInfo.contract.Transact(opts, "forwardPubkey", pubkey, newPubkey, user)
}

// ForwardPubkey is a paid mutator transaction binding the contract method 0xb8b64181.
//
// Solidity: function forwardPubkey(bytes pubkey, bytes newPubkey, bytes32 user) returns(bool success)
func (_KeyInfo *KeyInfoSession) ForwardPubkey(pubkey []byte, newPubkey []byte, user [32]byte) (*types.Transaction, error) {
	return _KeyInfo.Contract.ForwardPubkey(&_KeyInfo.TransactOpts, pubkey, newPubkey, user)
}

// ForwardPubkey is a paid mutator transaction binding the contract method 0xb8b64181.
//
// Solidity: function forwardPubkey(bytes pubkey, bytes newPubkey, bytes32 user) returns(bool success)
func (_KeyInfo *KeyInfoTransactorSession) ForwardPubkey(pubkey []byte, newPubkey []byte, user [32]byte) (*types.Transaction, error) {
	return _KeyInfo.Contract.ForwardPubkey(&_KeyInfo.TransactOpts, pubkey, newPubkey, user)
}

// KeyInfoKeyShareUpdatedIterator is returned from FilterKeyShareUpdated and is used to iterate over the raw logs and unpacked data for KeyShareUpdated events raised by the KeyInfo contract.
type KeyInfoKeyShareUpdatedIterator struct {
	Event *KeyInfoKeyShareUpdated // Event containing the contract specifics and raw log

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
func (it *KeyInfoKeyShareUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyInfoKeyShareUpdated)
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
		it.Event = new(KeyInfoKeyShareUpdated)
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
func (it *KeyInfoKeyShareUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyInfoKeyShareUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyInfoKeyShareUpdated represents a KeyShareUpdated event raised by the KeyInfo contract.
type KeyInfoKeyShareUpdated struct {
	Pubkey    common.Hash
	User      [32]byte
	Party1    []byte
	Party2    []byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterKeyShareUpdated is a free log retrieval operation binding the contract event 0x909956b054ea2ed1b865ac0e9bf20662c97d78d97205883668215fa0f7ca1b8e.
//
// Solidity: event KeyShareUpdated(bytes indexed pubkey, bytes32 indexed user, bytes party1, bytes party2, uint256 timestamp)
func (_KeyInfo *KeyInfoFilterer) FilterKeyShareUpdated(opts *bind.FilterOpts, pubkey [][]byte, user [][32]byte) (*KeyInfoKeyShareUpdatedIterator, error) {

	var pubkeyRule []interface{}
	for _, pubkeyItem := range pubkey {
		pubkeyRule = append(pubkeyRule, pubkeyItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _KeyInfo.contract.FilterLogs(opts, "KeyShareUpdated", pubkeyRule, userRule)
	if err != nil {
		return nil, err
	}
	return &KeyInfoKeyShareUpdatedIterator{contract: _KeyInfo.contract, event: "KeyShareUpdated", logs: logs, sub: sub}, nil
}

// WatchKeyShareUpdated is a free log subscription operation binding the contract event 0x909956b054ea2ed1b865ac0e9bf20662c97d78d97205883668215fa0f7ca1b8e.
//
// Solidity: event KeyShareUpdated(bytes indexed pubkey, bytes32 indexed user, bytes party1, bytes party2, uint256 timestamp)
func (_KeyInfo *KeyInfoFilterer) WatchKeyShareUpdated(opts *bind.WatchOpts, sink chan<- *KeyInfoKeyShareUpdated, pubkey [][]byte, user [][32]byte) (event.Subscription, error) {

	var pubkeyRule []interface{}
	for _, pubkeyItem := range pubkey {
		pubkeyRule = append(pubkeyRule, pubkeyItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _KeyInfo.contract.WatchLogs(opts, "KeyShareUpdated", pubkeyRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyInfoKeyShareUpdated)
				if err := _KeyInfo.contract.UnpackLog(event, "KeyShareUpdated", log); err != nil {
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

// ParseKeyShareUpdated is a log parse operation binding the contract event 0x909956b054ea2ed1b865ac0e9bf20662c97d78d97205883668215fa0f7ca1b8e.
//
// Solidity: event KeyShareUpdated(bytes indexed pubkey, bytes32 indexed user, bytes party1, bytes party2, uint256 timestamp)
func (_KeyInfo *KeyInfoFilterer) ParseKeyShareUpdated(log types.Log) (*KeyInfoKeyShareUpdated, error) {
	event := new(KeyInfoKeyShareUpdated)
	if err := _KeyInfo.contract.UnpackLog(event, "KeyShareUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// KeyInfoPubkeyForwardedIterator is returned from FilterPubkeyForwarded and is used to iterate over the raw logs and unpacked data for PubkeyForwarded events raised by the KeyInfo contract.
type KeyInfoPubkeyForwardedIterator struct {
	Event *KeyInfoPubkeyForwarded // Event containing the contract specifics and raw log

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
func (it *KeyInfoPubkeyForwardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeyInfoPubkeyForwarded)
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
		it.Event = new(KeyInfoPubkeyForwarded)
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
func (it *KeyInfoPubkeyForwardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KeyInfoPubkeyForwardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KeyInfoPubkeyForwarded represents a PubkeyForwarded event raised by the KeyInfo contract.
type KeyInfoPubkeyForwarded struct {
	Pubkey    common.Hash
	NewPubkey common.Hash
	User      [32]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPubkeyForwarded is a free log retrieval operation binding the contract event 0x439b679fc9f542ae812b84dc519a97855cda3b311b9b926d034e924bdf0cf1a0.
//
// Solidity: event PubkeyForwarded(bytes indexed pubkey, bytes indexed newPubkey, bytes32 indexed user, uint256 timestamp)
func (_KeyInfo *KeyInfoFilterer) FilterPubkeyForwarded(opts *bind.FilterOpts, pubkey [][]byte, newPubkey [][]byte, user [][32]byte) (*KeyInfoPubkeyForwardedIterator, error) {

	var pubkeyRule []interface{}
	for _, pubkeyItem := range pubkey {
		pubkeyRule = append(pubkeyRule, pubkeyItem)
	}
	var newPubkeyRule []interface{}
	for _, newPubkeyItem := range newPubkey {
		newPubkeyRule = append(newPubkeyRule, newPubkeyItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _KeyInfo.contract.FilterLogs(opts, "PubkeyForwarded", pubkeyRule, newPubkeyRule, userRule)
	if err != nil {
		return nil, err
	}
	return &KeyInfoPubkeyForwardedIterator{contract: _KeyInfo.contract, event: "PubkeyForwarded", logs: logs, sub: sub}, nil
}

// WatchPubkeyForwarded is a free log subscription operation binding the contract event 0x439b679fc9f542ae812b84dc519a97855cda3b311b9b926d034e924bdf0cf1a0.
//
// Solidity: event PubkeyForwarded(bytes indexed pubkey, bytes indexed newPubkey, bytes32 indexed user, uint256 timestamp)
func (_KeyInfo *KeyInfoFilterer) WatchPubkeyForwarded(opts *bind.WatchOpts, sink chan<- *KeyInfoPubkeyForwarded, pubkey [][]byte, newPubkey [][]byte, user [][32]byte) (event.Subscription, error) {

	var pubkeyRule []interface{}
	for _, pubkeyItem := range pubkey {
		pubkeyRule = append(pubkeyRule, pubkeyItem)
	}
	var newPubkeyRule []interface{}
	for _, newPubkeyItem := range newPubkey {
		newPubkeyRule = append(newPubkeyRule, newPubkeyItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _KeyInfo.contract.WatchLogs(opts, "PubkeyForwarded", pubkeyRule, newPubkeyRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KeyInfoPubkeyForwarded)
				if err := _KeyInfo.contract.UnpackLog(event, "PubkeyForwarded", log); err != nil {
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

// ParsePubkeyForwarded is a log parse operation binding the contract event 0x439b679fc9f542ae812b84dc519a97855cda3b311b9b926d034e924bdf0cf1a0.
//
// Solidity: event PubkeyForwarded(bytes indexed pubkey, bytes indexed newPubkey, bytes32 indexed user, uint256 timestamp)
func (_KeyInfo *KeyInfoFilterer) ParsePubkeyForwarded(log types.Log) (*KeyInfoPubkeyForwarded, error) {
	event := new(KeyInfoPubkeyForwarded)
	if err := _KeyInfo.contract.UnpackLog(event, "PubkeyForwarded", log); err != nil {
		return nil, err
	}
	return event, nil
}
