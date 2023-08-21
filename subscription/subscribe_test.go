package subscription

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getRealBackend() *ethclient.Client {
	cl, err := ethclient.Dial("http://127.0.0.1:8546")
	if err != nil {
		panic(err)
	}
	return cl
}

func TestSubscribe(t *testing.T) {
	backend := getRealBackend()
	ch := make(chan *types.Header)
	sub, err := backend.SubscribeNewHead(context.Background(), ch)
	if err != nil {
		//panic(err)
	}
	//sub.Unsubscribe()
	_ = sub
	close(ch)
}
