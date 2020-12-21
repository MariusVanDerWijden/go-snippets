// +build gofuzz

package fuzz

import (
	"encoding/binary"

	fuzz "github.com/google/gofuzz"
)

func Fuzz(data []byte) int {
	a := int(binary.BigEndian.Uint64(data))
	if CoolFunction(a, data) == 0 {
		return 0
	}
	return 1
}

func FuzzA(data []byte) int {
	fuzzer := fuzz.NewFromGoFuzz(data)
	var a int
	var b []byte
	fuzzer.Fuzz(&a)
	fuzzer.Fuzz(&b)
	if CoolFunction(a, b) == 0 {
		return 0
	}
	return 1
}
