package bigfuzz

import (
	"fmt"
	"math/big"
)

func Fuzz(input []byte) {
	base := new(big.Int)
	mod := new(big.Int)
	exp := new(big.Int)

	res := base.Exp(base, exp, mod)
	res2 := new(big.Int).Set(base)
	for i := 0; i < int(exp.Int64()); i++ {
		res2 = res2.Mul(res2, base)
		res2 = res2.Mod(res2, mod)
	}

	if res2.Cmp(res) != 0 {
		fmt.Printf("%v %v %v %v %v\n", base, exp, mod, res, res2)
		panic("crash")
	}
}
