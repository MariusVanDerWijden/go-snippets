package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/goevmlab/fuzzing"
	"github.com/holiman/goevmlab/ops"
	"github.com/holiman/goevmlab/program"
)

var (
	SK       = "0xcdfbe6f7602f67a97602e3e9fc24cde1cdffa88acd47745c0b84c5ff55891e1b"
	ADDR     = "0xb02A2EdA1b317FBd16760128836B0Ac59B560e9D"
	Contract = "0x51eF031Ba1B2128971bd67A3ba5dda9BfF8706c8"
	service  *eth.Ethereum
)

func sendTx(sk *ecdsa.PrivateKey, backend *ethclient.Client) {
	to := common.HexToAddress("0xEa8AeFde57181b47A186C570e49f6c0598306527")
	sender := common.HexToAddress(ADDR)
	nonce, err := backend.PendingNonceAt(context.Background(), sender)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Nonce: %v\n", nonce)
	gp, _ := backend.SuggestGasPrice(context.Background())
	tx := types.NewTransaction(nonce, to, big.NewInt(100000000000000000), 500000, gp, nil)
	signedTx, _ := types.SignTx(tx, types.HomesteadSigner{}, sk)
	backend.SendTransaction(context.Background(), signedTx)

}

func main() {
	cl, sk := getRealBackend()
	/*
		io, _, _, _, _, err := deploy(sk, cl)
		if err != nil {
			panic(err)
		}
		_ = io
	*/
	sendTx(sk, cl)
	/*
		for i := 0; i < 10000; i++ {
			/*
				if err := ioFuzz(io, sk, cl); err != nil {
					panic(err)
				}

			if err := jumpSubFuzz(cl, sk); err != nil {
				fmt.Printf("%v \n", err)
			}
			time.Sleep(50 * time.Millisecond)
		}*/
}

func ioFuzz(io *IO, sk *ecdsa.PrivateKey, cl *ethclient.Client) error {
	txopts := bind.NewKeyedTransactor(sk)
	var slot [32]byte
	var value [32]byte
	rand.Read(slot[:])
	rand.Read(value[:])
	if err := execAddListAndWait(sk, cl, func() (*types.Transaction, error) { return io.Read(txopts, slot) }); err != nil {
		return err
	}
	/*
		if err := execAddListAndWait(sk, cl, func() (*types.Transaction, error) { return io.ReadRevert(txopts, slot) }); err != nil {
			return err
		}*/
	if err := execAddListAndWait(sk, cl, func() (*types.Transaction, error) { return io.Write(txopts, slot, value) }); err != nil {
		return err
	}
	/*
		if err := execAddListAndWait(sk, cl, func() (*types.Transaction, error) { return io.WriteRevert(txopts, slot, value) }); err != nil {
			return err
		}*/
	return nil
}

func execAddListAndWait(sk *ecdsa.PrivateKey, cl *ethclient.Client, fun func() (*types.Transaction, error)) error {
	tx, err := fun()
	if err != nil {
		return err
	}
	if service == nil {
		service, err = testNode(*tx.To(), fuzzing.RandCallSubroutine([]common.Address{*tx.To()}))
		if err != nil {
			return err
		}
	}
	tx, err = txToAccessListTx(sk, cl, tx, service)
	if err != nil {
		return err
	}
	if err := cl.SendTransaction(context.Background(), tx); err != nil {
		//panic(err)
	}
	/*
		if _, err := bind.WaitMined(context.Background(), cl, tx); err != nil {
			return err
		}*/
	return nil
}

func getRealBackend() (*ethclient.Client, *ecdsa.PrivateKey) {
	// eth.sendTransaction({from:personal.listAccounts[0], to:"0xb02A2EdA1b317FBd16760128836B0Ac59B560e9D", value: "100000000000000"})

	sk := crypto.ToECDSAUnsafe(common.FromHex(SK))
	if crypto.PubkeyToAddress(sk.PublicKey).Hex() != ADDR {
		panic(fmt.Sprintf("wrong address want %s got %s", crypto.PubkeyToAddress(sk.PublicKey).Hex(), ADDR))
	}
	cl, err := ethclient.Dial("/home/matematik/.yolov3_2/geth.ipc")
	if err != nil {
		panic(err)
	}
	return cl, sk
}

func deploy(sk *ecdsa.PrivateKey, cl *ethclient.Client) (*IO, *DcallProxy, *CallProxy, *CcallProxy, *ScallProxy, error) {
	txopts := bind.NewKeyedTransactor(sk)
	// IO
	_, tx, io, err := DeployIO(txopts, cl)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	// DcallProxy
	_, tx, dcall, err := DeployDcallProxy(txopts, cl)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	// CallProxy
	_, tx, call, err := DeployCallProxy(txopts, cl)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	// CcallProxy
	_, tx, ccall, err := DeployCcallProxy(txopts, cl)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	// ScallProxy
	_, tx, scall, err := DeployScallProxy(txopts, cl)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	if _, err := bind.WaitDeployed(context.Background(), cl, tx); err != nil {
		return nil, nil, nil, nil, nil, err
	}
	return io, dcall, call, ccall, scall, nil
}

func testNode(to common.Address, code []byte) (*eth.Ethereum, error) {
	genesis := generateTestChain(big.NewInt(math.MaxBig63.Int64()), common.HexToAddress(ADDR), to, code)
	// Create node
	n, err := node.New(&node.Config{})
	if err != nil {
		return nil, err
	}
	// Create Ethereum Service
	config := &eth.Config{Genesis: genesis}
	config.Ethash.PowMode = ethash.ModeFake
	return eth.New(n, config)
}

func txToAccessListTx(sk *ecdsa.PrivateKey, cl *ethclient.Client, tx *types.Transaction, ethservice *eth.Ethereum) (*types.Transaction, error) {
	// Create a fake transaction to generate an access list
	var fakeTx *types.Transaction
	if tx.To() != nil {
		fakeTx = types.NewTransaction(0, *tx.To(), tx.Value(), tx.Gas(), tx.GasPrice(), tx.Data())
	} else {
		fakeTx = types.NewContractCreation(0, tx.Value(), tx.Gas(), tx.GasPrice(), tx.Data())
	}
	fakeTx, err := types.SignTx(fakeTx, types.HomesteadSigner{}, sk)
	if err != nil {
		return nil, err
	}

	list, err := ethservice.APIBackend.AccessList(context.Background(), nil, 0, fakeTx)
	if err != nil {
		return nil, err
	}
	// mutate the accesslist
	fmt.Printf("Before: %v\n", list)
	list = mutateAccessList(*list)
	fmt.Printf("After: %v\n", list)
	chainid, err := cl.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	// create proper accesslist transaction
	var newTx *types.Transaction
	if tx.To() == nil {
		newTx = types.NewAccessListContractCreation(chainid, tx.Nonce(), tx.Value(), tx.Gas(), tx.GasPrice(), tx.Data(), list)
	} else {
		newTx = types.NewAccessListTransaction(chainid, tx.Nonce(), *tx.To(), tx.Value(), tx.Gas(), tx.GasPrice(), tx.Data(), list)
	}
	return types.SignTx(newTx, types.NewEIP2718Signer(chainid), sk)
}

func generateTestChain(testBalance *big.Int, testAddr common.Address, dest common.Address, code []byte) *core.Genesis {
	config := params.AllEthashProtocolChanges
	genesis := &core.Genesis{
		Config:    config,
		Alloc:     core.GenesisAlloc{testAddr: {Balance: testBalance}, dest: {Balance: new(big.Int), Code: code}},
		ExtraData: []byte("test genesis"),
		Timestamp: 9000,
	}
	return genesis
}

func mutateAccessList(list types.AccessList) *types.AccessList {
	switch rand.Int31n(5) {
	case 0:
		// Leave the accesslist as is
		return &list
	case 1:
		// delete the access list
		return nil
	case 2:
		// empty the access list
		return &types.AccessList{}
	case 3:
		// add a random entry and random slots to the list
		addr := randomAddress()
		keys := []*common.Hash{}
		for i := 0; i < rand.Intn(10); i++ {
			h := randomHash()
			keys = append(keys, &h)
		}
		tuple := types.AccessTuple{Address: &addr, StorageKeys: keys}
		newList := types.AccessList(append([]types.AccessTuple{tuple}, list...))
		return &newList
	case 4:
		// replace a random entry and random slots of it in the list
		slot := list[rand.Int31n(int32(len(list)))]
		addr := randomAddress()
		keys := []*common.Hash{}
		if len(slot.StorageKeys) == 0 {
			break
		}
		for i := 0; i < rand.Intn(len(slot.StorageKeys)); i++ {
			h := randomHash()
			keys = append(keys, &h)
		}
		tuple := types.AccessTuple{Address: &addr, StorageKeys: keys}
		newList := types.AccessList(append([]types.AccessTuple{tuple}, list...))
		return &newList
	case 5:
		// replace a random slot in an existing entry
		keyIdx := rand.Int31n(int32(len(list)))
		slotIdx := rand.Int31n(int32(len(list[keyIdx].StorageKeys)))
		h := randomHash()
		list[keyIdx].StorageKeys[slotIdx] = &h
	case 6:
		var accesslist []types.AccessTuple
		for i := 0; i < rand.Int(); i++ {
			addr := randomAddress()
			keys := []*common.Hash{}
			// create a fully random access list
			for q := 0; q < rand.Int(); q++ {
				h := randomHash()
				keys = append(keys, &h)
			}
			tuple := types.AccessTuple{Address: &addr, StorageKeys: keys}
			accesslist = append(accesslist, tuple)
		}
		newList := types.AccessList(accesslist)
		return &newList
	}
	return &list
}

func randomHash() common.Hash {
	b := make([]byte, 32)
	rand.Read(b)
	return common.BytesToHash(b)
}

func randomAddress() common.Address {
	switch rand.Int31n(5) {
	case 0, 1, 2:
		b := make([]byte, 20)
		rand.Read(b)
		return common.BytesToAddress(b)
	case 3:
		return common.Address{}
	case 4:
		return common.HexToAddress(ADDR)
	}
	return common.Address{}
}

func createJumpSubTx(addresses []common.Address) []byte {
	p := program.NewProgram()
	code := fuzzing.RandCallSubroutine(addresses)
	p.CreateAndCall(code, true, ops.CALL)
	return p.Bytecode()
}

func jumpSubFuzz(backend *ethclient.Client, sk *ecdsa.PrivateKey) error {
	sender := common.HexToAddress(ADDR)
	code := createJumpSubTx([]common.Address{sender})
	nonce, err := backend.PendingNonceAt(context.Background(), sender)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Nonce: %v\n", nonce)
	gp, _ := backend.SuggestGasPrice(context.Background())
	tx := types.NewContractCreation(nonce, big.NewInt(0), 500000, gp, code)
	eth, err := testNode(common.Address{}, []byte{})
	if err != nil {
		panic(err)
	}
	newTx, err := txToAccessListTx(sk, backend, tx, eth)
	if err != nil {
		panic(err)
	}
	enc, _ := json.Marshal(newTx)
	fmt.Printf("%v\n", string(enc))
	return backend.SendTransaction(context.Background(), newTx)
}
