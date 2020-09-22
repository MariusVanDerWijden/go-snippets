package geth

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestBoard(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	backend, sk := getRealBackend()
	transactor := bind.NewKeyedTransactor(sk)
	_, txDep, board, err := DeployBigBoard(transactor, backend)
	bind.WaitDeployed(ctx, backend, txDep)
	assert.NoError(t, err)
	tx, err := board.Fill1(transactor)
	//_, err = board.Fill2(transactor)
	//_, err = board.Fill3(transactor)
	assert.NoError(t, err)
	receipt, err := bind.WaitMined(ctx, backend, tx)
	assert.NoError(t, err)
	assert.Equal(t, receipt.Status, types.ReceiptStatusSuccessful)
	for i := 0; i < 1000; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			state, err := board.GetBoardState(&bind.CallOpts{Pending: false})
			assert.NoError(t, err)
			for i := 0; i < 3; i++ {
				for k := 0; k < 10; k++ {
					for l := 0; l < 9; l++ {
						assert.Equal(t, uint8(i*(k+l)), state[i][k][l])
					}
				}
			}
		})
	}
}

func TestArray(t *testing.T) {
	backend, sk := getSimBackend()
	transactor := bind.NewKeyedTransactor(sk)
	_, _, board, err := DeployArray(transactor, backend)
	backend.Commit()
	assert.NoError(t, err)
	state, err := board.GetArray(nil)
	assert.NoError(t, err)
	fmt.Printf("%v", state)
	assert.Equal(t, state[0][0], uint8(0))
	assert.Equal(t, state[1][0], uint8(1))
	assert.Equal(t, state[2][0], uint8(2))
	assert.Equal(t, state[3][0], uint8(3))
	assert.Equal(t, state[0][1], uint8(4))
	assert.Equal(t, state[1][1], uint8(5))
	assert.Equal(t, state[2][1], uint8(6))
	assert.Equal(t, state[3][1], uint8(7))
}

/*
func TestIsArray(t *testing.T) {
	backend, sk := getSimBackend()
	transactor := bind.NewKeyedTransactor(sk)
	addr, _, _, _ := DeployArray(transactor, backend)
	checkopts := bind.CheckOpts{
		Address:     addr,
		BlockNumber: nil,
		Context:     context.Background(),
	}
	is, err := IsArray(&checkopts, backend)
	assert.NoError(t, err, "IsArray should not fail")
	assert.False(t, is, "IsArray should return false if contract not deployed")
	backend.Commit()
	is, err = IsArray(&checkopts, backend)
	assert.NoError(t, err, "IsArray should not fail")
	assert.True(t, is, "IsArray should return true once contract is deployed")
}*/
