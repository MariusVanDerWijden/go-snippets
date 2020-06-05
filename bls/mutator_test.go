package main

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/core/vm"
)

func TestMutator(t *testing.T) {
	b := []byte{1, 2, 3}
	config := MutationConfig{
		bin:          true,
		corpus:       make([][]byte, 0),
		MaxInputSize: 4096,
	}
	for i := 0; i < 10; i++ {
		re := Mutate(b, config)
		t.Log(fmt.Sprintf("Mutate %v", re))
		re = sanitizeMutate(re)
		t.Log(fmt.Sprintf("Sanitize %v", re))
	}
	a := NewG1Point(b, config)
	c := NewG1Point(b, config)
	g1Add := new(vm.Bls12381G1Add)
	_, err := g1Add.Run(append(a, c...))
	if err != nil {
		panic(err)
	}

	a = NewG2Point(b, config)
	c = NewG2Point(b, config)
	g2Add := new(vm.Bls12381G2Add)
	_, err = g2Add.Run(append(a, c...))
	if err != nil {
		panic(err)
	}
	panic("adsf")
}
