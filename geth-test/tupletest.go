// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package geth

import (
	"errors"
	"fmt"
	"math/big"
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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// T is an auto generated low-level Go binding around an user-defined struct.
type T struct {
	X *big.Int
	Y *big.Int
}

// TupleTestS is an auto generated low-level Go binding around an user-defined struct.
type TupleTestS struct {
	A *big.Int
	B []*big.Int
	C []T
}

// TupleTestMetaData contains all meta data concerning the TupleTest contract.
var TupleTestMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"b\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structT[]\",\"name\":\"c\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structTupleTest.S\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structT\",\"name\":\"\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"evF\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"b\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structT[]\",\"name\":\"c\",\"type\":\"tuple[]\"}],\"internalType\":\"structTupleTest.S\",\"name\":\"s\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structT\",\"name\":\"t\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"u\",\"type\":\"uint256\"}],\"name\":\"f\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"g\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"b\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structT[]\",\"name\":\"c\",\"type\":\"tuple[]\"}],\"internalType\":\"structTupleTest.S\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structT\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610787806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80636f2be7281461003b578063e2179b8e14610057575b600080fd5b61005560048036038101906100509190610346565b610077565b005b61005f6100b7565b60405161006e93929190610574565b60405180910390f35b7f5490ef71cb8b76b3f0788e975052f0f43ba07bc17bae7ea289ee00e50b4d84858383836040516100aa93929190610574565b60405180910390a1505050565b6100bf6100ce565b6100c76100ef565b6000909192565b60405180606001604052806000815260200160608152602001606081525090565b604051806040016040528060008152602001600081525090565b600061011c610117846105d7565b6105b2565b9050808382526020820190508285604086028201111561013f5761013e61071a565b5b60005b8581101561016f578161015588826102e1565b845260208401935060408301925050600181019050610142565b5050509392505050565b600061018c61018784610603565b6105b2565b905080838252602082019050828560208602820111156101af576101ae61071a565b5b60005b858110156101df57816101c58882610331565b8452602084019350602083019250506001810190506101b2565b5050509392505050565b600082601f8301126101fe576101fd61070b565b5b813561020e848260208601610109565b91505092915050565b600082601f83011261022c5761022b61070b565b5b813561023c848260208601610179565b91505092915050565b60006060828403121561025b5761025a610710565b5b61026560606105b2565b9050600061027584828501610331565b600083015250602082013567ffffffffffffffff81111561029957610298610715565b5b6102a584828501610217565b602083015250604082013567ffffffffffffffff8111156102c9576102c8610715565b5b6102d5848285016101e9565b60408301525092915050565b6000604082840312156102f7576102f6610710565b5b61030160406105b2565b9050600061031184828501610331565b600083015250602061032584828501610331565b60208301525092915050565b6000813590506103408161073a565b92915050565b60008060006080848603121561035f5761035e610724565b5b600084013567ffffffffffffffff81111561037d5761037c61071f565b5b61038986828701610245565b935050602061039a868287016102e1565b92505060606103ab86828701610331565b9150509250925092565b60006103c183836104f8565b60408301905092915050565b60006103d98383610556565b60208301905092915050565b60006103f08261064f565b6103fa818561067f565b93506104058361062f565b8060005b8381101561043657815161041d88826103b5565b975061042883610665565b925050600181019050610409565b5085935050505092915050565b600061044e8261065a565b6104588185610690565b93506104638361063f565b8060005b8381101561049457815161047b88826103cd565b975061048683610672565b925050600181019050610467565b5085935050505092915050565b60006060830160008301516104b96000860182610556565b50602083015184820360208601526104d18282610443565b915050604083015184820360408601526104eb82826103e5565b9150508091505092915050565b60408201600082015161050e6000850182610556565b5060208201516105216020850182610556565b50505050565b60408201600082015161053d6000850182610556565b5060208201516105506020850182610556565b50505050565b61055f816106a1565b82525050565b61056e816106a1565b82525050565b6000608082019050818103600083015261058e81866104a1565b905061059d6020830185610527565b6105aa6060830184610565565b949350505050565b60006105bc6105cd565b90506105c882826106ab565b919050565b6000604051905090565b600067ffffffffffffffff8211156105f2576105f16106dc565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561061e5761061d6106dc565b5b602082029050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b6000819050919050565b6106b482610729565b810181811067ffffffffffffffff821117156106d3576106d26106dc565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b610743816106a1565b811461074e57600080fd5b5056fea2646970667358221220e3bbca2b4baa0a3e0ff7e111d7b2b2daa848dfeb1cd49bf215461ba18e8c5c8164736f6c63430008070033",
}

// TupleTest is an auto generated Go binding around an Ethereum contract.
type TupleTest struct {
	abi abi.ABI
}

// TupleTestInstance represents a deployed instance of the TupleTest contract.
type TupleTestInstance struct {
	TupleTest
	address common.Address
}

func (i *TupleTestInstance) Address() common.Address {
	return i.address
}

// NewTupleTest creates a new instance of TupleTest.
func NewTupleTest() (*TupleTest, error) {
	parsed, err := TupleTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	return &TupleTest{abi: *parsed}, nil
}

func (_TupleTest *TupleTest) PackConstructor() ([]byte, error) {
	return _TupleTest.abi.Pack("")
}

// F is a free data retrieval call binding the contract method 0x6f2be728.
//
// Solidity: function f((uint256,uint256[],(uint256,uint256)[]) s, (uint256,uint256) t, uint256 u) returns()
func (_TupleTest *TupleTest) PackF(s TupleTestS, t T, u *big.Int) ([]byte, error) {
	return _TupleTest.abi.Pack("f", s, t, u)
}

func (_TupleTest *TupleTest) UnpackF(data []byte) (struct{}, error) {
	out, err := _TupleTest.abi.Unpack("f", data)
	_ = out
	outstruct := new(struct{})
	if err != nil {
		return *outstruct, err
	}

	return *outstruct, err

}

// G is a free data retrieval call binding the contract method 0xe2179b8e.
//
// Solidity: function g() pure returns((uint256,uint256[],(uint256,uint256)[]), (uint256,uint256), uint256)
func (_TupleTest *TupleTest) PackG() ([]byte, error) {
	return _TupleTest.abi.Pack("g")
}

func (_TupleTest *TupleTest) UnpackG(data []byte) (struct {
	Arg  TupleTestS
	Arg0 T
	Arg1 *big.Int
}, error) {
	out, err := _TupleTest.abi.Unpack("g", data)

	outstruct := new(struct {
		Arg  TupleTestS
		Arg0 T
		Arg1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Arg = *abi.ConvertType(out[0], new(TupleTestS)).(*TupleTestS)
	outstruct.Arg0 = *abi.ConvertType(out[1], new(T)).(*T)
	outstruct.Arg1 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TupleTestEvF represents a EvF event raised by the TupleTest contract.
type TupleTestEvF struct {
	Arg0 TupleTestS
	Arg1 T
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

func (_TupleTest *TupleTest) EvFEventID() common.Hash {
	return common.HexToHash("0x5490ef71cb8b76b3f0788e975052f0f43ba07bc17bae7ea289ee00e50b4d8485")
}

func (_TupleTest *TupleTest) UnpackEvFEvent(log types.Log) (*TupleTestEvF, error) {
	event := "EvF"
	if log.Topics[0] != _TupleTest.abi.Events[event].ID {
		return nil, fmt.Errorf("event signature mismatch")
	}
	out := new(TupleTestEvF)
	if len(log.Data) > 0 {
		if err := _TupleTest.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range _TupleTest.abi.Events[event].Inputs {
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
