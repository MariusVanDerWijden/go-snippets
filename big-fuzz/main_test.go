package bigfuzz

import (
	"fmt"
	"math/big"
	"testing"
)

func FuzzMod(f *testing.F) {
	/*
		f.Add("2938462938472983472983659726349017249287491026512746239764525612965293865296239471239874193284792387498274256129746192347",
			"298472983472983471903246121093472394872319615612417471234712061",
			"29834729834729834729347290846729561262544958723956495615629569234729836259263598127342374289365912465901365498236492183464")
	*/
	f.Add("375", "249", "388")
	//f.Add("79228162514264337593543950319", "298472983472983471903246121093472394872319615612417471234712061", "38844611825139420086633594242")

	fuzz := func(t *testing.T, baseS, expS, modS string) {

		base, ok := new(big.Int).SetString(baseS, 0)
		if !ok {
			return
		}
		exp, ok := new(big.Int).SetString(expS, 0)
		if !ok {
			return
		}
		mod, _ := new(big.Int).SetString(modS, 0)
		if !ok {
			return
		}
		res := new(big.Int).Exp(base, exp, mod)
		res2 := new(big.Int).Set(base)
		for i := 1; i < int(exp.Int64()); i++ {
			res2 = res2.Mul(res2, base)
			res2 = res2.Mod(res2, mod)
		}
		if exp.Cmp(big.NewInt(0)) == 0 {
			res2 = big.NewInt(1)
		}

		if res2.Cmp(res) != 0 {
			panic(fmt.Sprintf("crash: %v %v %v %v %v\n", base, exp, mod, res, res2))
		}
	}

	f.Fuzz(fuzz)
}
