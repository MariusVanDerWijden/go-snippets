package geth

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestBinding(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	backend, sk := GetSimBackend()
	transactor, _ := bind.NewKeyedTransactorWithChainID(sk, big.NewInt(1337))
	// Deploy
	addr, tx, err := bind.DeployContract2(transactor, common.FromHex(EventerMetaData.Bin), nil, backend)
	backend.Commit()
	assert.NoError(t, err)
	// Call AnonEvent()
	eventer, err := NewEventer()
	assert.NoError(t, err)
	call, err := eventer.PackAnonEvent()
	assert.NoError(t, err)
	tx, err = bind.Transact2(transactor, addr, call, backend)
	backend.Commit()
	assert.NoError(t, err)
	receipt, err := bind.WaitMined(ctx, backend, tx)
	assert.NoError(t, err)
	assert.True(t, receipt.Status != types.ReceiptStatusFailed)
	// Unpack events
	pk := crypto.PubkeyToAddress(sk.PublicKey)
	for _, log := range receipt.Logs {
		ev, err := eventer.UnpackAnonEventEvent(*log)
		assert.NoError(t, err)
		assert.Equal(t, ev.Arg0, pk)
		assert.Equal(t, ev.Arg1, pk)
	}
}
