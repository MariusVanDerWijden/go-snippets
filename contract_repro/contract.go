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

// TellorABI is the input ABI used to generate the binding from.
const TellorABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minerIndex\",\"type\":\"uint256\"}],\"name\":\"beginDispute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNewCurrentVariables\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256\",\"name\":\"_difficutly\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNewVariablesOnDeck\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"idsOnDeck\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"tipsOnDeck\",\"type\":\"uint256[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTopRequestIDs\",\"outputs\":[{\"internalType\":\"uint256[5]\",\"name\":\"_requestIds\",\"type\":\"uint256[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_propNewTellorAddress\",\"type\":\"address\"}],\"name\":\"proposeFork\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_pendingOwner\",\"type\":\"address\"}],\"name\":\"proposeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"requestStakingWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nonce\",\"type\":\"string\"},{\"internalType\":\"uint256[5]\",\"name\":\"_requestId\",\"type\":\"uint256[5]\"},{\"internalType\":\"uint256[5]\",\"name\":\"_value\",\"type\":\"uint256[5]\"}],\"name\":\"submitMiningSolution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"tallyVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"unlockDisputeFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"updateTellor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_supportsDispute\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TellorBin is the compiled bytecode used for deploying new contracts.
var TellorBin = "0x608060405234801561001057600080fd5b506113af806100206000396000f3fe608060405234801561001057600080fd5b50600436106101825760003560e01c8063710bf322116100d85780639a7077ab1161008c578063c9d27afe11610066578063c9d27afe14610536578063f458ab981461055b578063fe1cd15d1461057857610182565b80639a7077ab14610489578063a9059cbb146104f5578063bed9d8611461052e57610182565b80638581af19116100bd5780638581af191461043b57806395d89b41146104645780639a01ca131461046c57610182565b8063710bf322146103e5578063752d49a11461041857610182565b806328449c3a1161013a5780634350283e116101145780634350283e1461034c5780634d318b0e146103c05780634e71e0c8146103dd57610182565b806328449c3a146102d1578063313ce567146102d95780634049f198146102f757610182565b80630d2d76a21161016b5780630d2d76a21461025157806323b872dd1461025b57806326b7d9f61461029e57610182565b806306fdde0314610187578063095ea7b314610204575b600080fd5b61018f6105b8565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101c95781810151838201526020016101b1565b50505050905090810190601f1680156101f65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61023d6004803603604081101561021a57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356105ef565b604080519115158252519081900360200190f35b6102596106b2565b005b61023d6004803603606081101561027157600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813581169160208101359091169060400135610735565b610259600480360360208110156102b457600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610801565b6102596108a1565b6102e161090a565b6040805160ff9092168252519081900360200190f35b6102ff61090f565b604051848152602081018460a080838360005b8381101561032a578181015183820152602001610312565b5050505090500183815260200182815260200194505050505060405180910390f35b610259600480360361016081101561036357600080fd5b81019060208101813564010000000081111561037e57600080fd5b82018360208201111561039057600080fd5b803590602001918460018302840111640100000000831117156103b257600080fd5b919350915060a08101610934565b610259600480360360208110156103d657600080fd5b5035610a4a565b610259610aba565b610259600480360360208110156103fb57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610b23565b6102596004803603604081101561042e57600080fd5b5080359060200135610ba8565b6102596004803603606081101561045157600080fd5b5080359060208101359060400135610c3b565b61018f610cd6565b6102596004803603602081101561048257600080fd5b5035610d0d565b610491610d7d565b604051808360a080838360005b838110156104b657818101518382015260200161049e565b5050505090500182600560200280838360005b838110156104e15781810151838201526020016104c9565b505050509050019250505060405180910390f35b61023d6004803603604081101561050b57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610d9f565b610259610e2f565b6102596004803603604081101561054c57600080fd5b50803590602001351515610e98565b6102596004803603602081101561057157600080fd5b5035610f10565b610580610f80565b604051808260a080838360005b838110156105a557818101518382015260200161058d565b5050505090500191505060405180910390f35b60408051808201909152600f81527f54656c6c6f722054726962757465730000000000000000000000000000000000602082015290565b604080517f850dcc3200000000000000000000000000000000000000000000000000000000815260006004820181905273ffffffffffffffffffffffffffffffffffffffff8516602483015260448201849052915173455e70b77f59219c699e206ba3dc9f45faa1841d9163850dcc32916064808301926020929190829003018186803b15801561067f57600080fd5b505af4158015610693573d6000803e3d6000fd5b505050506040513d60208110156106a957600080fd5b50519392505050565b604080517f820a2d660000000000000000000000000000000000000000000000000000000081526000600482018190529151734fdb593ead869b38240459f7116f78cb9a4d178f9263820a2d669260248082019391829003018186803b15801561071b57600080fd5b505af415801561072f573d6000803e3d6000fd5b50505050565b604080517fca50189900000000000000000000000000000000000000000000000000000000815260006004820181905273ffffffffffffffffffffffffffffffffffffffff80871660248401528516604483015260648201849052915173455e70b77f59219c699e206ba3dc9f45faa1841d9163ca501899916084808301926020929190829003018186803b1580156107cd57600080fd5b505af41580156107e1573d6000803e3d6000fd5b505050506040513d60208110156107f757600080fd5b5051949350505050565b604080517f694bf49f00000000000000000000000000000000000000000000000000000000815260006004820181905273ffffffffffffffffffffffffffffffffffffffff84166024830152915173d0332a131edfcaf2e86f4af16e0294f213fc105e9263694bf49f9260448082019391829003018186803b15801561088657600080fd5b505af415801561089a573d6000803e3d6000fd5b5050505050565b604080517fc9cf5e4c0000000000000000000000000000000000000000000000000000000081526000600482018190529151734fdb593ead869b38240459f7116f78cb9a4d178f9263c9cf5e4c9260248082019391829003018186803b15801561071b57600080fd5b601290565b600061091961135c565b6000806109266000610f97565b935093509350935090919293565b600073d4cdaa978fdcd627bb2fc08e1b99f87332485d1a63a4bc40679091868686866040518663ffffffff1660e01b8152600401808681526020018060200184600560200280828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01690910190508360a080828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01690910183810383528681526020019050868680828437600081840152601f19601f820116905080830192505050965050505050505060006040518083038186803b158015610a2c57600080fd5b505af4158015610a40573d6000803e3d6000fd5b5050505050505050565b604080517fdef6fac700000000000000000000000000000000000000000000000000000000815260006004820181905260248201849052915173d0332a131edfcaf2e86f4af16e0294f213fc105e9263def6fac79260448082019391829003018186803b15801561088657600080fd5b604080517f314691ff000000000000000000000000000000000000000000000000000000008152600060048201819052915173d4cdaa978fdcd627bb2fc08e1b99f87332485d1a9263314691ff9260248082019391829003018186803b15801561071b57600080fd5b604080517f291f8b7300000000000000000000000000000000000000000000000000000000815260006004820181905273ffffffffffffffffffffffffffffffffffffffff84166024830152915173d4cdaa978fdcd627bb2fc08e1b99f87332485d1a9263291f8b739260448082019391829003018186803b15801561088657600080fd5b604080517f02e8f21b0000000000000000000000000000000000000000000000000000000081526000600482018190526024820185905260448201849052915173d4cdaa978fdcd627bb2fc08e1b99f87332485d1a926302e8f21b9260648082019391829003018186803b158015610c1f57600080fd5b505af4158015610c33573d6000803e3d6000fd5b505050505050565b604080517fca9a4ea5000000000000000000000000000000000000000000000000000000008152600060048201819052602482018690526044820185905260648201849052915173d0332a131edfcaf2e86f4af16e0294f213fc105e9263ca9a4ea59260848082019391829003018186803b158015610cb957600080fd5b505af4158015610ccd573d6000803e3d6000fd5b50505050505050565b60408051808201909152600381527f5452420000000000000000000000000000000000000000000000000000000000602082015290565b604080517f97f5f97200000000000000000000000000000000000000000000000000000000815260006004820181905260248201849052915173d0332a131edfcaf2e86f4af16e0294f213fc105e926397f5f9729260448082019391829003018186803b15801561088657600080fd5b610d8561135c565b610d8d61135c565b610d976000611066565b915091509091565b604080517fc84b96f500000000000000000000000000000000000000000000000000000000815260006004820181905273ffffffffffffffffffffffffffffffffffffffff8516602483015260448201849052915173455e70b77f59219c699e206ba3dc9f45faa1841d9163c84b96f5916064808301926020929190829003018186803b15801561067f57600080fd5b604080517f44bacc4b0000000000000000000000000000000000000000000000000000000081526000600482018190529151734fdb593ead869b38240459f7116f78cb9a4d178f926344bacc4b9260248082019391829003018186803b15801561071b57600080fd5b604080517f2da0706e000000000000000000000000000000000000000000000000000000008152600060048201819052602482018590528315156044830152915173d0332a131edfcaf2e86f4af16e0294f213fc105e92632da0706e9260648082019391829003018186803b158015610c1f57600080fd5b604080517f22048ecf00000000000000000000000000000000000000000000000000000000815260006004820181905260248201849052915173d0332a131edfcaf2e86f4af16e0294f213fc105e926322048ecf9260448082019391829003018186803b15801561088657600080fd5b610f8861135c565b610f92600061111b565b905090565b6000610fa161135c565b600080805b6005811015610fdc57856035018160058110610fbe57fe5b6002020154848260058110610fcf57fe5b6020020152600101610fa6565b50508354604080517f646966666963756c7479000000000000000000000000000000000000000000008152815190819003600a01812060009081528288016020818152848320547f63757272656e74546f74616c54697073000000000000000000000000000000008552855194859003601001909420835252919091205491945091509193509193565b61106e61135c565b61107661135c565b61107f8361111b565b915060005b60058110156111155783604801600084836005811061109f57fe5b60200201518152602001908152602001600020600401600060405180807f746f74616c5469700000000000000000000000000000000000000000000000008152506008019050604051809103902081526020019081526020016000205482826005811061110857fe5b6020020152600101611084565b50915091565b61112361135c565b61112b61135c565b61113361135c565b6040805161066081019182905261116c91600187019060339082845b81548152602001906001019080831161114f575050505050611208565b909250905060005b60058110156112005782816005811061118957fe5b6020020151156111cf578460430160008383600581106111a557fe5b60200201518152602001908152602001600020548482600581106111c557fe5b60200201526111f8565b8460350181600403600581106111e157fe5b60020201548482600581106111f257fe5b60200201525b600101611174565b505050919050565b61121061135c565b61121861135c565b60208301516000805b600581101561129b5785816001016033811061123957fe5b602002015185826005811061124a57fe5b60200201526001810184826005811061125f57fe5b60200201528285826005811061127157fe5b602002015110156112935784816005811061128857fe5b602002015192508091505b600101611221565b5060065b603381101561135457828682603381106112b557fe5b6020020151111561134c578581603381106112cc57fe5b60200201518583600581106112dd57fe5b6020020152808483600581106112ef57fe5b602002015285816033811061130057fe5b6020020151925060005b600581101561134a578386826005811061132057fe5b602002015110156113425785816005811061133757fe5b602002015193508092505b60010161130a565b505b60010161129f565b505050915091565b6040518060a00160405280600590602082028038833950919291505056fea265627a7a7231582002c6db9a1e63d4da189817d449dd0b78f56053fd0ac422f7689a27f90b39258a64736f6c63430005100032"

// DeployTellor deploys a new Ethereum contract, binding an instance of Tellor to it.
func DeployTellor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tellor, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TellorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tellor{TellorCaller: TellorCaller{contract: contract}, TellorTransactor: TellorTransactor{contract: contract}, TellorFilterer: TellorFilterer{contract: contract}}, nil
}

// Tellor is an auto generated Go binding around an Ethereum contract.
type Tellor struct {
	TellorCaller     // Read-only binding to the contract
	TellorTransactor // Write-only binding to the contract
	TellorFilterer   // Log filterer for contract events
}

// TellorCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorSession struct {
	Contract     *Tellor           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorCallerSession struct {
	Contract *TellorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TellorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorTransactorSession struct {
	Contract     *TellorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorRaw struct {
	Contract *Tellor // Generic contract binding to access the raw methods on
}

// TellorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorCallerRaw struct {
	Contract *TellorCaller // Generic read-only contract binding to access the raw methods on
}

// TellorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorTransactorRaw struct {
	Contract *TellorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellor creates a new instance of Tellor, bound to a specific deployed contract.
func NewTellor(address common.Address, backend bind.ContractBackend) (*Tellor, error) {
	contract, err := bindTellor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tellor{TellorCaller: TellorCaller{contract: contract}, TellorTransactor: TellorTransactor{contract: contract}, TellorFilterer: TellorFilterer{contract: contract}}, nil
}

// NewTellorCaller creates a new read-only instance of Tellor, bound to a specific deployed contract.
func NewTellorCaller(address common.Address, caller bind.ContractCaller) (*TellorCaller, error) {
	contract, err := bindTellor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorCaller{contract: contract}, nil
}

// NewTellorTransactor creates a new write-only instance of Tellor, bound to a specific deployed contract.
func NewTellorTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorTransactor, error) {
	contract, err := bindTellor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorTransactor{contract: contract}, nil
}

// NewTellorFilterer creates a new log filterer instance of Tellor, bound to a specific deployed contract.
func NewTellorFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorFilterer, error) {
	contract, err := bindTellor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorFilterer{contract: contract}, nil
}

// bindTellor binds a generic wrapper to an already deployed contract.
func bindTellor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tellor *TellorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tellor.Contract.TellorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tellor *TellorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tellor.Contract.TellorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tellor *TellorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tellor.Contract.TellorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tellor *TellorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tellor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tellor *TellorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tellor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tellor *TellorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tellor.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Tellor *TellorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Tellor *TellorSession) Decimals() (uint8, error) {
	return _Tellor.Contract.Decimals(&_Tellor.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Tellor *TellorCallerSession) Decimals() (uint8, error) {
	return _Tellor.Contract.Decimals(&_Tellor.CallOpts)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficutly, uint256 _tip)
func (_Tellor *TellorCaller) GetNewCurrentVariables(opts *bind.CallOpts) (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "getNewCurrentVariables")

	outstruct := new(struct {
		Challenge  [32]byte
		RequestIds [5]*big.Int
		Difficutly *big.Int
		Tip        *big.Int
	})

	outstruct.Challenge = out[0].([32]byte)
	outstruct.RequestIds = out[1].([5]*big.Int)
	outstruct.Difficutly = out[2].(*big.Int)
	outstruct.Tip = out[3].(*big.Int)

	return *outstruct, err

}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficutly, uint256 _tip)
func (_Tellor *TellorSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	return _Tellor.Contract.GetNewCurrentVariables(&_Tellor.CallOpts)
}

// GetNewCurrentVariables is a free data retrieval call binding the contract method 0x4049f198.
//
// Solidity: function getNewCurrentVariables() view returns(bytes32 _challenge, uint256[5] _requestIds, uint256 _difficutly, uint256 _tip)
func (_Tellor *TellorCallerSession) GetNewCurrentVariables() (struct {
	Challenge  [32]byte
	RequestIds [5]*big.Int
	Difficutly *big.Int
	Tip        *big.Int
}, error) {
	return _Tellor.Contract.GetNewCurrentVariables(&_Tellor.CallOpts)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_Tellor *TellorCaller) GetNewVariablesOnDeck(opts *bind.CallOpts) (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "getNewVariablesOnDeck")

	outstruct := new(struct {
		IdsOnDeck  [5]*big.Int
		TipsOnDeck [5]*big.Int
	})

	outstruct.IdsOnDeck = out[0].([5]*big.Int)
	outstruct.TipsOnDeck = out[1].([5]*big.Int)

	return *outstruct, err

}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_Tellor *TellorSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _Tellor.Contract.GetNewVariablesOnDeck(&_Tellor.CallOpts)
}

// GetNewVariablesOnDeck is a free data retrieval call binding the contract method 0x9a7077ab.
//
// Solidity: function getNewVariablesOnDeck() view returns(uint256[5] idsOnDeck, uint256[5] tipsOnDeck)
func (_Tellor *TellorCallerSession) GetNewVariablesOnDeck() (struct {
	IdsOnDeck  [5]*big.Int
	TipsOnDeck [5]*big.Int
}, error) {
	return _Tellor.Contract.GetNewVariablesOnDeck(&_Tellor.CallOpts)
}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_Tellor *TellorCaller) GetTopRequestIDs(opts *bind.CallOpts) ([5]*big.Int, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "getTopRequestIDs")

	if err != nil {
		return *new([5]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([5]*big.Int)).(*[5]*big.Int)

	return out0, err

}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_Tellor *TellorSession) GetTopRequestIDs() ([5]*big.Int, error) {
	return _Tellor.Contract.GetTopRequestIDs(&_Tellor.CallOpts)
}

// GetTopRequestIDs is a free data retrieval call binding the contract method 0xfe1cd15d.
//
// Solidity: function getTopRequestIDs() view returns(uint256[5] _requestIds)
func (_Tellor *TellorCallerSession) GetTopRequestIDs() ([5]*big.Int, error) {
	return _Tellor.Contract.GetTopRequestIDs(&_Tellor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Tellor *TellorCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Tellor *TellorSession) Name() (string, error) {
	return _Tellor.Contract.Name(&_Tellor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Tellor *TellorCallerSession) Name() (string, error) {
	return _Tellor.Contract.Name(&_Tellor.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Tellor *TellorCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tellor.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Tellor *TellorSession) Symbol() (string, error) {
	return _Tellor.Contract.Symbol(&_Tellor.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Tellor *TellorCallerSession) Symbol() (string, error) {
	return _Tellor.Contract.Symbol(&_Tellor.CallOpts)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_Tellor *TellorTransactor) AddTip(opts *bind.TransactOpts, _requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "addTip", _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_Tellor *TellorSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.AddTip(&_Tellor.TransactOpts, _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _tip) returns()
func (_Tellor *TellorTransactorSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.AddTip(&_Tellor.TransactOpts, _requestId, _tip)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Tellor *TellorTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Tellor *TellorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.Approve(&_Tellor.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Tellor *TellorTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.Approve(&_Tellor.TransactOpts, _spender, _amount)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_Tellor *TellorTransactor) BeginDispute(opts *bind.TransactOpts, _requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "beginDispute", _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_Tellor *TellorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.BeginDispute(&_Tellor.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(uint256 _requestId, uint256 _timestamp, uint256 _minerIndex) returns()
func (_Tellor *TellorTransactorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.BeginDispute(&_Tellor.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Tellor *TellorTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Tellor *TellorSession) ClaimOwnership() (*types.Transaction, error) {
	return _Tellor.Contract.ClaimOwnership(&_Tellor.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Tellor *TellorTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _Tellor.Contract.ClaimOwnership(&_Tellor.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Tellor *TellorTransactor) DepositStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "depositStake")
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Tellor *TellorSession) DepositStake() (*types.Transaction, error) {
	return _Tellor.Contract.DepositStake(&_Tellor.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Tellor *TellorTransactorSession) DepositStake() (*types.Transaction, error) {
	return _Tellor.Contract.DepositStake(&_Tellor.TransactOpts)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_Tellor *TellorTransactor) ProposeFork(opts *bind.TransactOpts, _propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "proposeFork", _propNewTellorAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_Tellor *TellorSession) ProposeFork(_propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _Tellor.Contract.ProposeFork(&_Tellor.TransactOpts, _propNewTellorAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(address _propNewTellorAddress) returns()
func (_Tellor *TellorTransactorSession) ProposeFork(_propNewTellorAddress common.Address) (*types.Transaction, error) {
	return _Tellor.Contract.ProposeFork(&_Tellor.TransactOpts, _propNewTellorAddress)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _pendingOwner) returns()
func (_Tellor *TellorTransactor) ProposeOwnership(opts *bind.TransactOpts, _pendingOwner common.Address) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "proposeOwnership", _pendingOwner)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _pendingOwner) returns()
func (_Tellor *TellorSession) ProposeOwnership(_pendingOwner common.Address) (*types.Transaction, error) {
	return _Tellor.Contract.ProposeOwnership(&_Tellor.TransactOpts, _pendingOwner)
}

// ProposeOwnership is a paid mutator transaction binding the contract method 0x710bf322.
//
// Solidity: function proposeOwnership(address _pendingOwner) returns()
func (_Tellor *TellorTransactorSession) ProposeOwnership(_pendingOwner common.Address) (*types.Transaction, error) {
	return _Tellor.Contract.ProposeOwnership(&_Tellor.TransactOpts, _pendingOwner)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Tellor *TellorTransactor) RequestStakingWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "requestStakingWithdraw")
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Tellor *TellorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _Tellor.Contract.RequestStakingWithdraw(&_Tellor.TransactOpts)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Tellor *TellorTransactorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _Tellor.Contract.RequestStakingWithdraw(&_Tellor.TransactOpts)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_Tellor *TellorTransactor) SubmitMiningSolution(opts *bind.TransactOpts, _nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "submitMiningSolution", _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_Tellor *TellorSession) SubmitMiningSolution(_nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.SubmitMiningSolution(&_Tellor.TransactOpts, _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x4350283e.
//
// Solidity: function submitMiningSolution(string _nonce, uint256[5] _requestId, uint256[5] _value) returns()
func (_Tellor *TellorTransactorSession) SubmitMiningSolution(_nonce string, _requestId [5]*big.Int, _value [5]*big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.SubmitMiningSolution(&_Tellor.TransactOpts, _nonce, _requestId, _value)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_Tellor *TellorTransactor) TallyVotes(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "tallyVotes", _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_Tellor *TellorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.TallyVotes(&_Tellor.TransactOpts, _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(uint256 _disputeId) returns()
func (_Tellor *TellorTransactorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.TallyVotes(&_Tellor.TransactOpts, _disputeId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Tellor *TellorTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Tellor *TellorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.Transfer(&_Tellor.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Tellor *TellorTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.Transfer(&_Tellor.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Tellor *TellorTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Tellor *TellorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.TransferFrom(&_Tellor.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Tellor *TellorTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.TransferFrom(&_Tellor.TransactOpts, _from, _to, _amount)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_Tellor *TellorTransactor) UnlockDisputeFee(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "unlockDisputeFee", _disputeId)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_Tellor *TellorSession) UnlockDisputeFee(_disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.UnlockDisputeFee(&_Tellor.TransactOpts, _disputeId)
}

// UnlockDisputeFee is a paid mutator transaction binding the contract method 0x9a01ca13.
//
// Solidity: function unlockDisputeFee(uint256 _disputeId) returns()
func (_Tellor *TellorTransactorSession) UnlockDisputeFee(_disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.UnlockDisputeFee(&_Tellor.TransactOpts, _disputeId)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_Tellor *TellorTransactor) UpdateTellor(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "updateTellor", _disputeId)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_Tellor *TellorSession) UpdateTellor(_disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.UpdateTellor(&_Tellor.TransactOpts, _disputeId)
}

// UpdateTellor is a paid mutator transaction binding the contract method 0xf458ab98.
//
// Solidity: function updateTellor(uint256 _disputeId) returns()
func (_Tellor *TellorTransactorSession) UpdateTellor(_disputeId *big.Int) (*types.Transaction, error) {
	return _Tellor.Contract.UpdateTellor(&_Tellor.TransactOpts, _disputeId)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_Tellor *TellorTransactor) Vote(opts *bind.TransactOpts, _disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "vote", _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_Tellor *TellorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Tellor.Contract.Vote(&_Tellor.TransactOpts, _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _disputeId, bool _supportsDispute) returns()
func (_Tellor *TellorTransactorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Tellor.Contract.Vote(&_Tellor.TransactOpts, _disputeId, _supportsDispute)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Tellor *TellorTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tellor.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Tellor *TellorSession) WithdrawStake() (*types.Transaction, error) {
	return _Tellor.Contract.WithdrawStake(&_Tellor.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Tellor *TellorTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _Tellor.Contract.WithdrawStake(&_Tellor.TransactOpts)
}
