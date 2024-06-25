package main

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type initS struct {
	connext      common.Address
	updateDomain uint32
	updater      common.Address
	splitmain    common.Address
	split        common.Address
	owner        common.Address
}

func TestMain(t *testing.T) {
	i := []interface{}{
		common.Address{1},
		uint32(2),
		common.Address{0},
		common.Address{4},
		common.Address{5},
		common.Address{0},
	}
	res, err := rlp.EncodeToBytes(i)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x", res)
	panic("adsf")
}
