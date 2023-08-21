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

// KeyinfoMetaData contains all meta data concerning the Keyinfo contract.
var KeyinfoMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"user\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"party1\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"party2\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"KeyShareUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"newPubkey\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"user\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PubkeyForwarded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"user\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"party1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"party2\",\"type\":\"bytes\"}],\"name\":\"addKeyShareInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newPubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"user\",\"type\":\"bytes32\"}],\"name\":\"forwardPubkey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"pubkeyForwards\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"pubkeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"party1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"party2\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userkeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610c53806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806307873ddc1461005c5780632ac8597c1461008c5780634e6a698d146100be578063545239fe146100ee578063b8b641811461011e575b600080fd5b6100766004803603810190610071919061071f565b61014e565b604051610083919061098d565b60405180910390f35b6100a660048036038101906100a1919061081a565b61027d565b6040516100b5939291906109ca565b60405180910390f35b6100d860048036038101906100d3919061081a565b6103cd565b6040516100e591906109a8565b60405180910390f35b610108600480360381019061010391906107da565b610483565b60405161011591906109a8565b60405180910390f35b61013860048036038101906101339190610863565b61053c565b604051610145919061098d565b60405180910390f35b60008360405180606001604052808581526020018481526020014281525060008660405161017c9190610976565b908152602001604051809103902060008201518160000190805190602001906101a69291906105e2565b5060208201518160010190805190602001906101c39291906105e2565b5060408201518160020155905050600160008781526020019081526020016000208590806001815401808255809150506001900390600052602060002001600090919091909150908051906020019061021d9291906105e2565b50858560405161022d9190610976565b60405180910390207f909956b054ea2ed1b865ac0e9bf20662c97d78d97205883668215fa0f7ca1b8e868642604051610268939291906109ca565b60405180910390a36001915050949350505050565b6000818051602081018201805184825260208301602085012081835280955050505050506000915090508060000180546102b690610b09565b80601f01602080910402602001604051908101604052809291908181526020018280546102e290610b09565b801561032f5780601f106103045761010080835404028352916020019161032f565b820191906000526020600020905b81548152906001019060200180831161031257829003601f168201915b50505050509080600101805461034490610b09565b80601f016020809104026020016040519081016040528092919081815260200182805461037090610b09565b80156103bd5780601f10610392576101008083540402835291602001916103bd565b820191906000526020600020905b8154815290600101906020018083116103a057829003601f168201915b5050505050908060020154905083565b600281805160208101820180518482526020830160208501208183528095505050505050600091509050805461040290610b09565b80601f016020809104026020016040519081016040528092919081815260200182805461042e90610b09565b801561047b5780601f106104505761010080835404028352916020019161047b565b820191906000526020600020905b81548152906001019060200180831161045e57829003601f168201915b505050505081565b6001602052816000526040600020818154811061049f57600080fd5b906000526020600020016000915091505080546104bb90610b09565b80601f01602080910402602001604051908101604052809291908181526020018280546104e790610b09565b80156105345780601f1061050957610100808354040283529160200191610534565b820191906000526020600020905b81548152906001019060200180831161051757829003601f168201915b505050505081565b600083836002866040516105509190610976565b908152602001604051809103902090805190602001906105719291906105e2565b5082846040516105819190610976565b6040518091039020866040516105979190610976565b60405180910390207f439b679fc9f542ae812b84dc519a97855cda3b311b9b926d034e924bdf0cf1a0426040516105ce9190610a0f565b60405180910390a460019150509392505050565b8280546105ee90610b09565b90600052602060002090601f0160209004810192826106105760008555610657565b82601f1061062957805160ff1916838001178555610657565b82800160010185558215610657579182015b8281111561065657825182559160200191906001019061063b565b5b5090506106649190610668565b5090565b5b80821115610681576000816000905550600101610669565b5090565b600061069861069384610a4f565b610a2a565b9050828152602081018484840111156106b4576106b3610bcf565b5b6106bf848285610ac7565b509392505050565b6000813590506106d681610bef565b92915050565b600082601f8301126106f1576106f0610bca565b5b8135610701848260208601610685565b91505092915050565b60008135905061071981610c06565b92915050565b6000806000806080858703121561073957610738610bd9565b5b6000610747878288016106c7565b945050602085013567ffffffffffffffff81111561076857610767610bd4565b5b610774878288016106dc565b935050604085013567ffffffffffffffff81111561079557610794610bd4565b5b6107a1878288016106dc565b925050606085013567ffffffffffffffff8111156107c2576107c1610bd4565b5b6107ce878288016106dc565b91505092959194509250565b600080604083850312156107f1576107f0610bd9565b5b60006107ff858286016106c7565b92505060206108108582860161070a565b9150509250929050565b6000602082840312156108305761082f610bd9565b5b600082013567ffffffffffffffff81111561084e5761084d610bd4565b5b61085a848285016106dc565b91505092915050565b60008060006060848603121561087c5761087b610bd9565b5b600084013567ffffffffffffffff81111561089a57610899610bd4565b5b6108a6868287016106dc565b935050602084013567ffffffffffffffff8111156108c7576108c6610bd4565b5b6108d3868287016106dc565b92505060406108e4868287016106c7565b9150509250925092565b6108f781610aa7565b82525050565b600061090882610a80565b6109128185610a8b565b9350610922818560208601610ad6565b61092b81610bde565b840191505092915050565b600061094182610a80565b61094b8185610a9c565b935061095b818560208601610ad6565b80840191505092915050565b61097081610abd565b82525050565b60006109828284610936565b915081905092915050565b60006020820190506109a260008301846108ee565b92915050565b600060208201905081810360008301526109c281846108fd565b905092915050565b600060608201905081810360008301526109e481866108fd565b905081810360208301526109f881856108fd565b9050610a076040830184610967565b949350505050565b6000602082019050610a246000830184610967565b92915050565b6000610a34610a45565b9050610a408282610b3b565b919050565b6000604051905090565b600067ffffffffffffffff821115610a6a57610a69610b9b565b5b610a7382610bde565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b60008115159050919050565b6000819050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610af4578082015181840152602081019050610ad9565b83811115610b03576000848401525b50505050565b60006002820490506001821680610b2157607f821691505b60208210811415610b3557610b34610b6c565b5b50919050565b610b4482610bde565b810181811067ffffffffffffffff82111715610b6357610b62610b9b565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b610bf881610ab3565b8114610c0357600080fd5b50565b610c0f81610abd565b8114610c1a57600080fd5b5056fea264697066735822122006e8821ab05ef0c2654a6bbae499ebb9a39b8cea9a0ce28f0994581abbc91f0764736f6c63430008070033",
}

// Keyinfo is an auto generated Go binding around an Ethereum contract.
type Keyinfo struct {
	abi abi.ABI
}

// KeyinfoInstance represents a deployed instance of the Keyinfo contract.
type KeyinfoInstance struct {
	Keyinfo
	address common.Address
}

func (i *KeyinfoInstance) Address() common.Address {
	return i.address
}

// NewKeyinfo creates a new instance of Keyinfo.
func NewKeyinfo() (*Keyinfo, error) {
	parsed, err := KeyinfoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	return &Keyinfo{abi: *parsed}, nil
}

func (_Keyinfo *Keyinfo) PackConstructor() ([]byte, error) {
	return _Keyinfo.abi.Pack("")
}

// AddKeyShareInfo is a free data retrieval call binding the contract method 0x07873ddc.
//
// Solidity: function addKeyShareInfo(bytes32 user, bytes pubkey, bytes party1, bytes party2) returns(bool success)
func (_Keyinfo *Keyinfo) PackAddKeyShareInfo(user [32]byte, pubkey []byte, party1 []byte, party2 []byte) ([]byte, error) {
	return _Keyinfo.abi.Pack("addKeyShareInfo", user, pubkey, party1, party2)
}

func (_Keyinfo *Keyinfo) UnpackAddKeyShareInfo(data []byte) (bool, error) {
	out, err := _Keyinfo.abi.Unpack("addKeyShareInfo", data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ForwardPubkey is a free data retrieval call binding the contract method 0xb8b64181.
//
// Solidity: function forwardPubkey(bytes pubkey, bytes newPubkey, bytes32 user) returns(bool success)
func (_Keyinfo *Keyinfo) PackForwardPubkey(pubkey []byte, newPubkey []byte, user [32]byte) ([]byte, error) {
	return _Keyinfo.abi.Pack("forwardPubkey", pubkey, newPubkey, user)
}

func (_Keyinfo *Keyinfo) UnpackForwardPubkey(data []byte) (bool, error) {
	out, err := _Keyinfo.abi.Unpack("forwardPubkey", data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PubkeyForwards is a free data retrieval call binding the contract method 0x4e6a698d.
//
// Solidity: function pubkeyForwards(bytes ) view returns(bytes)
func (_Keyinfo *Keyinfo) PackPubkeyForwards(arg0 []byte) ([]byte, error) {
	return _Keyinfo.abi.Pack("pubkeyForwards", arg0)
}

func (_Keyinfo *Keyinfo) UnpackPubkeyForwards(data []byte) ([]byte, error) {
	out, err := _Keyinfo.abi.Unpack("pubkeyForwards", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Pubkeys is a free data retrieval call binding the contract method 0x2ac8597c.
//
// Solidity: function pubkeys(bytes ) view returns(bytes party1, bytes party2, uint256 timestamp)
func (_Keyinfo *Keyinfo) PackPubkeys(arg0 []byte) ([]byte, error) {
	return _Keyinfo.abi.Pack("pubkeys", arg0)
}

func (_Keyinfo *Keyinfo) UnpackPubkeys(data []byte) (struct {
	Party1    []byte
	Party2    []byte
	Timestamp *big.Int
}, error) {
	out, err := _Keyinfo.abi.Unpack("pubkeys", data)

	outstruct := new(struct {
		Party1    []byte
		Party2    []byte
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Party1 = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Party2 = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Userkeys is a free data retrieval call binding the contract method 0x545239fe.
//
// Solidity: function userkeys(bytes32 , uint256 ) view returns(bytes)
func (_Keyinfo *Keyinfo) PackUserkeys(arg0 [32]byte, arg1 *big.Int) ([]byte, error) {
	return _Keyinfo.abi.Pack("userkeys", arg0, arg1)
}

func (_Keyinfo *Keyinfo) UnpackUserkeys(data []byte) ([]byte, error) {
	out, err := _Keyinfo.abi.Unpack("userkeys", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// KeyinfoKeyShareUpdated represents a KeyShareUpdated event raised by the Keyinfo contract.
type KeyinfoKeyShareUpdated struct {
	Pubkey    common.Hash
	User      [32]byte
	Party1    []byte
	Party2    []byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

func (_Keyinfo *Keyinfo) KeyShareUpdatedEventID() common.Hash {
	return common.HexToHash("0x909956b054ea2ed1b865ac0e9bf20662c97d78d97205883668215fa0f7ca1b8e")
}

func (_Keyinfo *Keyinfo) UnpackKeyShareUpdatedEvent(log types.Log) (*KeyinfoKeyShareUpdated, error) {
	event := "KeyShareUpdated"
	if log.Topics[0] != _Keyinfo.abi.Events[event].ID {
		return nil, fmt.Errorf("event signature mismatch")
	}
	out := new(KeyinfoKeyShareUpdated)
	if len(log.Data) > 0 {
		if err := _Keyinfo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range _Keyinfo.abi.Events[event].Inputs {
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

// KeyinfoPubkeyForwarded represents a PubkeyForwarded event raised by the Keyinfo contract.
type KeyinfoPubkeyForwarded struct {
	Pubkey    common.Hash
	NewPubkey common.Hash
	User      [32]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

func (_Keyinfo *Keyinfo) PubkeyForwardedEventID() common.Hash {
	return common.HexToHash("0x439b679fc9f542ae812b84dc519a97855cda3b311b9b926d034e924bdf0cf1a0")
}

func (_Keyinfo *Keyinfo) UnpackPubkeyForwardedEvent(log types.Log) (*KeyinfoPubkeyForwarded, error) {
	event := "PubkeyForwarded"
	if log.Topics[0] != _Keyinfo.abi.Events[event].ID {
		return nil, fmt.Errorf("event signature mismatch")
	}
	out := new(KeyinfoPubkeyForwarded)
	if len(log.Data) > 0 {
		if err := _Keyinfo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range _Keyinfo.abi.Events[event].Inputs {
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
