package main

import (
	"context"
	"io/ioutil"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

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

			var in []byte
			k := rand.Int31n(100)
			switch i {
			case 10, 11:
				in = make([]byte, 256)
			case 12:
				in = make([]byte, 160*k)
			case 13, 14:
				in = make([]byte, 512)
			case 15, 16:
				in = make([]byte, 288*k)
			case 17:
				in = make([]byte, 384*k)
			}
			copy(in, input)
			/*
				if i == 10 || i == 11 || i == 13 || i == 14 {
					name := rand.Int63()
					half := make([]byte, len(in)/2)
					copy(half, in)
					ioutil.WriteFile(fmt.Sprintf("corpus/%d", name), half, 0644)
					name = rand.Int63()
					double := make([]byte, len(in)*2)
					copy(double, in)
					copy(double[len(in):], in)
					ioutil.WriteFile(fmt.Sprintf("corpus/%d", name), double, 0644)
					t.Log("Writing doubles/halfes")
				}*/
			transactor := bind.NewKeyedTransactor(sk)
			tx, err := contract.CallPrec(transactor, big.NewInt(i), input)
			if err == nil {
				_, err = bind.WaitMined(ctx, backend, tx)
				t.Log("Successful")
			}
			//t.Log(i, len(in), in, common.Bytes2Hex(in), err)
			//time.Sleep(3 * time.Second)
		}
		//time.Sleep(30 * time.Second)

	}

}
