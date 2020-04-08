package geth

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNegativeEvent(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	backend, sk := getSimBackend()
	transactor := bind.NewKeyedTransactor(sk)
	_, _, eventer, err := DeployEventer(transactor, backend)
	assert.NoError(t, err)

	tx, err := eventer.GetEvent(transactor)
	backend.Commit()
	assert.NoError(t, err)
	receipt, err := bind.WaitMined(ctx, backend, tx)
	assert.NoError(t, err)
	assert.True(t, receipt.Status != types.ReceiptStatusFailed)

	// filter for events
	opts := bind.FilterOpts{
		Start:   0,
		Context: ctx,
		End:     nil,
	}
	// Set to 2,3 for functioning filter
	iter, err := eventer.FilterTestInt8(&opts, []int8{-2}, []int8{-3})
	assert.NoError(t, err)
	if iter.Next() {
		t.Log(iter.Event.Out1)
		t.Log(iter.Event.Out2)
		t.Log(iter.Event.Raw)
		t.Log(iter.Next())
		t.Error("Successful if it hits this error")
	}
	t.Error("Unsuccessful")
}

func TestAnonymousEvent(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	backend, sk := getSimBackend()
	transactor := bind.NewKeyedTransactor(sk)
	_, _, eventer, err := DeployEventer(transactor, backend)
	assert.NoError(t, err)

	tx, err := eventer.AnonEvent(transactor)
	backend.Commit()
	assert.NoError(t, err)
	receipt, err := bind.WaitMined(ctx, backend, tx)
	assert.NoError(t, err)
	assert.True(t, receipt.Status != types.ReceiptStatusFailed)

	// filter for events
	opts := bind.FilterOpts{
		Start:   0,
		Context: ctx,
		End:     nil,
	}
	// Set to 2,3 for functioning filter
	iter, err := eventer.FilterAnonEvent(&opts)
	assert.NoError(t, err)
	if iter.Next() {
		t.Log(iter.Event.Arg0)
		t.Log(iter.Event.Arg1)
		t.Log(iter.Event.Raw)
		t.Log(iter.Next())
		t.Error("Successful if it hits this error")
	}
	t.Error(iter.Error())
	t.Error("Unsuccessful")
}

func TestFail(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	backend, sk := getSimBackend()
	transactor := bind.NewKeyedTransactor(sk)
	_, _, eventer, err := DeployEventer(transactor, backend)
	assert.NoError(t, err)

	tx, err := eventer.Fail(transactor)
	require.NoError(t, err)
	backend.Commit()
	receipt, err := bind.WaitMined(ctx, backend, tx)
	assert.NoError(t, err)
	assert.True(t, receipt.Status != types.ReceiptStatusFailed)
}
