package bigtest

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// MaxUint256 is the maximum value that can be represented by a uint256
var MaxUint256 = new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 256), common.Big1)

// MaxInt256 is the maximum value that can be represented by a int256
var MaxInt256 = new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 255), common.Big1)

func BenchmarkBigAdd(b *testing.B) {
	start, _ := new(big.Int).SetString("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 0) // -1
	for i := int64(0); i < int64(b.N); i++ {
		ret := new(big.Int).Sub(start, big.NewInt(i))
		ret.Add(MaxUint256, new(big.Int).Neg(ret))
		ret.Add(ret, common.Big1)
		ret.Neg(ret)
	}

}

func BenchmarkBigMod(b *testing.B) {
	start, _ := new(big.Int).SetString("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 0) // -1
	for i := int64(0); i < int64(b.N); i++ {
		ret := new(big.Int).Sub(start, big.NewInt(i))
		ret.Mod(ret, MaxInt256)
		ret.Neg(ret)
	}
}
