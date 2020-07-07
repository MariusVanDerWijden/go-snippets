package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/crypto/bls12381"
)

var modulo, _ = big.NewInt(0).SetString("0x1a0111ea397fe69a4b1ba7b6434bacd764774b84f38512bf6730d2a0f6b0f6241eabfffeb153ffffb9feffffffffaaab", 0)

// sanitizeMutate maps arbitrary input to a valid field element on the curve
func sanitizeMutate(data []byte) []byte {
	fe := big.NewInt(0).SetBytes(data)
	if fe.Cmp(modulo) == 1 {
		fe = fe.Mod(fe, modulo)
	}
	buf := make([]byte, 48)
	copy(buf[48-len(fe.Bytes()):], fe.Bytes())
	return buf
}

func NewFieldElement(iv []byte, config MutationConfig) []byte {
	re := Mutate(iv, config)
	return sanitizeMutate(re)
}

func NewG1Point(iv []byte, config MutationConfig) []byte {
	a := NewFieldElement(iv, config)
	// Initialize G1
	g := bls12381.NewG1()
	b, err := g.MapToCurve(a)
	if err != nil {
		panic(err)
	}
	return g.EncodePoint(b)
}

func NewG2Point(iv []byte, config MutationConfig) []byte {
	a := NewFieldElement(iv, config)
	b := NewFieldElement(a, config)
	x := append(a, b...)
	// Initialize G2
	g := bls12381.NewG2()
	// Compute mapping
	res, err := g.MapToCurve(x)
	if err != nil {
		panic(err)
	}
	return g.EncodePoint(res)
}
