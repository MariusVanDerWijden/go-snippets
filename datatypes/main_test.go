package main

import (
	"testing"

	"github.com/go-snippets/big"
)

func BenchmarkBigCmp(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		a := big.NewInt(int64(i % 1000))
		if a.BitLen() == 0 {
			t++
		}
	}
}

func BenchmarkBigCmp2(b *testing.B) {
	t := 0
	for i := 0; i < b.N; i++ {
		a := big.NewInt(int64(i % 1000))
		if a.Cmp(big.NewInt(0)) == 0 {
			t++
		}
	}
}
