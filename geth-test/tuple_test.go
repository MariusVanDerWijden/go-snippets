package geth

/*
import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestTuple(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	backend, sk := GetSimBackend()
	transactor, _ := bind.NewKeyedTransactorWithChainID(sk, big.NewInt(1337))
	_, _, contract, err := DeployTupleTest(transactor, backend)
	backend.Commit()
	assert.NoError(t, err)
	tupleT := []T{{X: big.NewInt(1), Y: big.NewInt(2)}, {X: big.NewInt(2), Y: big.NewInt(3)}}
	s := TupleTestS{A: big.NewInt(0), B: []*big.Int{big.NewInt(1)}, C: tupleT}
	tx, err := contract.F(transactor, s, tupleT[0], big.NewInt(5))
	backend.Commit()
	assert.NoError(t, err)
	receipt, err := bind.WaitMined(ctx, backend, tx)
	assert.NoError(t, err)
	assert.Equal(t, receipt.Status, types.ReceiptStatusSuccessful)
	t.Log(receipt.Logs)
	opts := bind.FilterOpts{
		Start:   0,
		Context: ctx,
		End:     nil,
	}
	iterF, err := contract.FilterEvF(&opts)
	assert.NoError(t, err)
	if !iterF.Next() {
		t.Fatal("Unsuccessful")
	}
	t.Log(iterF.Event.Arg0)
	t.Log(iterF.Event.Arg1)
	t.Log(iterF.Event.Arg2)
}
*/
