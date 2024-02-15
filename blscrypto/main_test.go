package main

import (
	"testing"

	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"

	cryptobls12381 "github.com/kilic/bls12-381"
)

/*
cpu: AMD Ryzen 9 5900X 12-Core Processor
BenchmarkG1-24       	   32427	     36491 ns/op	     368 B/op	       6 allocs/op
BenchmarkG2-24       	   30607	     40983 ns/op	     184 B/op	       3 allocs/op
BenchmarkKilic-24    	  293353	      5116 ns/op	    8056 B/op	      49 allocs/op
*/

func BenchmarkG1(b *testing.B) {
	g1 := new(bls12381.G1Affine)
	for i := 0; i < b.N; i++ {
		g1.IsInSubGroup()
	}
}

func BenchmarkG2(b *testing.B) {
	g2 := new(bls12381.G2Affine)
	for i := 0; i < b.N; i++ {
		g2.IsInSubGroup()
	}
}

func BenchmarkKilic(b *testing.B) {
	g1 := new(cryptobls12381.PointG1)
	curve := new(cryptobls12381.G1)
	for i := 0; i < b.N; i++ {
		curve.IsOnCurve(g1)
		curve.InCorrectSubgroup(g1)
	}
}
