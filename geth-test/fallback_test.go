package geth

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestFallback(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	backend, sk := getSimBackend()
	transactor := bind.NewKeyedTransactor(sk)
	_, _, contract, err := DeployReceiveFallback(transactor, backend)
	assert.NoError(t, err)
	raw := ReceiveFallbackRaw{contract}
	// filter for events
	opts := bind.FilterOpts{
		Start:   0,
		Context: ctx,
		End:     nil,
	}
	// Receive
	tx, err := raw.Transfer(transactor)
	assert.NoError(t, err)
	backend.Commit()
	assert.NoError(t, err)
	receipt, err := bind.WaitMined(ctx, backend, tx)
	assert.NoError(t, err)
	assert.True(t, receipt.Status != types.ReceiptStatusFailed)

	iter, err := contract.FilterReceive(&opts)
	assert.NoError(t, err)
	if !iter.Next() {
		t.Fatal("Unsuccessful")
	}
	t.Log(iter.Event.Sender)

	// Fallback
	tx2, err := raw.Transact(transactor, "()")
	assert.NoError(t, err)
	backend.Commit()
	assert.NoError(t, err)
	receipt2, err := bind.WaitMined(ctx, backend, tx2)
	assert.NoError(t, err)
	assert.True(t, receipt2.Status != types.ReceiptStatusFailed)
	iterF, err := contract.FilterFallback(&opts)
	assert.NoError(t, err)
	if !iterF.Next() {
		t.Fatal("Unsuccessful")
	}
	t.Log(iterF.Event.Sender)

}
