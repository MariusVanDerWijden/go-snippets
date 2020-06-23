package main

import (
	"context"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

/*
func TestFallback(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	backend, sk := getRealBackend()
	contract, err := NewCallBLS(common.HexToAddress(Contract), backend)
	assert.NoError(t, err)

	dir, err := ioutil.ReadDir("corpus")
	if err != nil {
		t.Error(err)
	}
	for _, info := range dir {
		input, err := ioutil.ReadFile("corpus/" + info.Name())
		assert.NoError(t, err)

		for i := int64(10); i < 18; i++ {
			transactor := bind.NewKeyedTransactor(sk)
			tx, err := contract.CallPrec(transactor, big.NewInt(i), input)
			if err == nil {
				_, err = bind.WaitMined(ctx, backend, tx)
				t.Log("Successful")
				t.Log(i)
				t.Log(input)
			}
		}
	}

}
*/

/*
func TestFallbackOffline(t *testing.T) {
	dir, err := ioutil.ReadDir("corpus")
	if err != nil {
		t.Error(err)
	}
	for _, info := range dir {
		input, err := ioutil.ReadFile("corpus/" + info.Name())
		assert.NoError(t, err)

		for i := int64(10); i < 18; i++ {
			bls := new(vm.Bls12381G1Add)
			_, err = bls.Run(input)
			if err == nil {
				t.Log("Successful")
				t.Log(i)
				t.Log(common.ToHex(input))
			}
			//t.Log(i, len(in), in, common.Bytes2Hex(in), err)
			//time.Sleep(3 * time.Second)
		}
		//time.Sleep(30 * time.Second)

	}
	panic("aadsf")
}
*/

/*
func TestMulG2Live(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 240*time.Minute)
	defer cancel()
	backend, sk := getRealBackend()
	auth := bind.NewKeyedTransactor(sk)
	_, tx, contract, err := DeployEIP2537Caller(auth, backend)
	assert.NoError(t, err)
	_, err = bind.WaitDeployed(ctx, backend, tx)
	assert.NoError(t, err)
	transactor := bind.NewKeyedTransactor(sk)
	config := MutationConfig{
		bin:          true,
		corpus:       makeCorpus(),
		MaxInputSize: 4096,
	}

	iv := []byte{1, 2, 3}
	for i := 0; i < 1000; i++ {
		for k := 0; k < 100; k++ {
			input := NewG2Point(iv, config)
			mul := make([]byte, 32)
			rand.Read(mul)
			tx, err = contract.CallWithMutation(transactor, common.BigToAddress(big.NewInt(0x0e)), append(input, mul...))
			iv = append(input, mul...)
		}
		if err == nil {
			_, err = bind.WaitMined(ctx, backend, tx)
			t.Logf("Call successful: %d", i)
			b, err := contract.LastSuccess(nil)
			assert.NoError(t, err)
			t.Log(b)
		}
	}

	panic(err)
}
*/
func TestMulExpG1Live(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 240*time.Minute)
	defer cancel()
	backend, sk := getRealBackend()
	auth := bind.NewKeyedTransactor(sk)
	_, tx, contract, err := DeployEIP2537Caller(auth, backend)
	assert.NoError(t, err)
	_, err = bind.WaitDeployed(ctx, backend, tx)
	assert.NoError(t, err)
	transactor := bind.NewKeyedTransactor(sk)
	config := MutationConfig{
		bin:          true,
		corpus:       makeCorpus(),
		MaxInputSize: 4096,
	}
	_ = rand.Int()

	iv := []byte{1, 2, 3, 4}
	for i := 0; i < 1000; i++ {
		for j := 0; j < 130; j++ {
			var in []byte
			for k := 0; k < j; k++ {
				input := NewG1Point(iv, config)
				mul := make([]byte, 32)
				rand.Read(mul)
				in = append(in, append(input, mul...)...)
			}
			tx, err = contract.CallWithMutation(transactor, common.BigToAddress(big.NewInt(0x0c)), in)
			iv = append(in)
		}
		if err == nil {
			_, err = bind.WaitMined(ctx, backend, tx)
			t.Logf("Call successful: %d", i)
			b, err := contract.LastSuccess(nil)
			assert.NoError(t, err)
			t.Log(b)
		}
	}

	panic(err)
}

func TestMulG1Live(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 240*time.Minute)
	defer cancel()
	backend, sk := getRealBackend()
	auth := bind.NewKeyedTransactor(sk)
	_, tx, contract, err := DeployEIP2537Caller(auth, backend)
	assert.NoError(t, err)
	_, err = bind.WaitDeployed(ctx, backend, tx)
	assert.NoError(t, err)
	transactor := bind.NewKeyedTransactor(sk)
	config := MutationConfig{
		bin:          true,
		corpus:       makeCorpus(),
		MaxInputSize: 4096,
	}

	iv := []byte{1, 2, 3}
	for i := 0; i < 1000; i++ {
		for k := 0; k < 100; k++ {
			input := NewG1Point(iv, config)
			mul := make([]byte, 32)
			rand.Read(mul)
			tx, err = contract.CallWithMutation(transactor, common.BigToAddress(big.NewInt(0x0b)), append(input, mul...))
			iv = append(input, mul...)
		}
		if err == nil {
			_, err = bind.WaitMined(ctx, backend, tx)
			t.Logf("Call successful: %d", i)
			b, err := contract.LastSuccess(nil)
			assert.NoError(t, err)
			t.Log(b)
		}
	}

	panic(err)
}

func TestAddG2Live(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 240*time.Minute)
	defer cancel()
	backend, sk := getRealBackend()
	auth := bind.NewKeyedTransactor(sk)
	_, tx, contract, err := DeployEIP2537Caller(auth, backend)
	assert.NoError(t, err)
	_, err = bind.WaitDeployed(ctx, backend, tx)
	assert.NoError(t, err)
	transactor := bind.NewKeyedTransactor(sk)
	config := MutationConfig{
		bin:          true,
		corpus:       make([][]byte, 0),
		MaxInputSize: 4096,
	}

	iv := []byte{1, 2, 3}
	for i := 0; i < 100; i++ {
		for k := 0; k < 100; k++ {
			input := NewG2Point(iv, config)
			input2 := NewG2Point(iv, config)
			tx, err = contract.CallWithMutation(transactor, common.BigToAddress(big.NewInt(0x0d)), append(input, input2...))
			iv = append(input, input2...)
		}
		if err == nil {
			_, err = bind.WaitMined(ctx, backend, tx)
			t.Log("Call successful")
			b, err := contract.LastSuccess(nil)
			assert.NoError(t, err)
			t.Log(b)
		}
	}

	panic(err)
}

func TestAddLive(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 240*time.Minute)
	defer cancel()
	backend, sk := getRealBackend()
	auth := bind.NewKeyedTransactor(sk)
	_, tx, contract, err := DeployEIP2537Caller(auth, backend)
	assert.NoError(t, err)
	_, err = bind.WaitDeployed(ctx, backend, tx)
	assert.NoError(t, err)
	transactor := bind.NewKeyedTransactor(sk)
	config := MutationConfig{
		bin:          true,
		corpus:       make([][]byte, 0),
		MaxInputSize: 4096,
	}

	iv := []byte{1, 2, 3}
	for i := 0; i < 10000; i++ {
		for k := 0; k < 100; k++ {
			input := NewG1Point(iv, config)
			input2 := NewG1Point(iv, config)
			tx, err = contract.CallWithMutation(transactor, common.BigToAddress(big.NewInt(0x0a)), append(input, input2...))
			iv = append(input, input2...)
		}
		if err == nil {
			_, err = bind.WaitMined(ctx, backend, tx)
			t.Log("Call successful")
			b, err := contract.LastSuccess(nil)
			assert.NoError(t, err)
			t.Log(b)
		}
	}

	panic(err)
}
