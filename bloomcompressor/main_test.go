package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	bloom "github.com/holiman/bloomfilter/v2"
	"golang.org/x/crypto/sha3"
)

func BenchmarkBloom(b *testing.B) {
	size := uint64(128 * 1024 * 1024 * 8)
	funcs := uint64(8)
	filter, err := bloom.New(size, funcs)
	if err != nil {
		panic(err)
	}
	h := hash([]byte{127, 0, 0, 1, 120, 0, 0, 0})
	for i := 0; i < b.N; i++ {
		if filter.ContainsHash(h) {

		}
		filter.AddHash(h)
	}
}

func TestOtherBloom(t *testing.T) {
	size := uint64(128 * 1024 * 1024 * 8)
	funcs := uint64(8)
	filter, err := bloom.New(size, funcs)
	if err != nil {
		panic(err)
	}
	filter.AddHash(hash([]byte{127, 0, 0, 1, 120, 0, 0, 0}))
}

func hash(f []byte) uint64 {
	return binary.BigEndian.Uint64(f)
}

func TestBloomFilter(t *testing.T) {
	Init()
	filter := new(types.Bloom)
	fmt.Printf("%v of %v \n", countOnes(filter.Bytes()), len(filter)*8)
	filter.Add([]byte{1})
	fmt.Printf("%v of %v \n", countOnes(filter.Bytes()), len(filter)*8)
	for i := 0; i < len(filter); i++ {
		slice := make([]byte, 8)
		binary.PutUvarint(slice, uint64(i))
		hash := sha3.Sum256(slice)
		filter.Add(hash[:])
	}
	fmt.Printf("%v of %v \n", countOnes(filter.Bytes()), len(filter)*8)
	for i := 0; i < len(filter); i++ {
		slice := make([]byte, 8)
		binary.PutUvarint(slice, uint64(i+len(filter)))
		hash := sha3.Sum256(slice)
		filter.Add(hash[:])
	}
	fmt.Printf("%v of %v \n", countOnes(filter.Bytes()), len(filter)*8)
	panic("adf")
}

func TestCount(t *testing.T) {
	Init()
	tests := []struct {
		x    int
		want int
	}{
		{1, 1},
		{2, 1},
		{0, 0},
		{255, 8},
	}
	for _, test := range tests {
		if test.want != table[test.x] {
			t.Fatalf("want %v got %v", test.want, table[test.x])
		}
	}
}

var table []int

func Init() {
	for i := 0; i < 256; i++ {
		cnt := i & 1
		cnt += (i >> 7 & 1)
		cnt += (i >> 6 & 1)
		cnt += (i >> 5 & 1)
		cnt += (i >> 4 & 1)
		cnt += (i >> 3 & 1)
		cnt += (i >> 2 & 1)
		cnt += (i >> 1 & 1)
		table = append(table, cnt)
	}
}

func countOnes(buffer []byte) int {
	res := 0
	for _, b := range buffer {
		res += table[b]
	}
	return res
}

func TestMakeTestFile(t *testing.T) {
	file, err := os.Create("testdata")
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	for i := 0; i < 40_000_000; i++ {
		a := rand.Int31n(256)
		b := rand.Int31n(256)
		datacenter := rand.Int31n(256)
		writer.WriteString(fmt.Sprintf("%v.%v.%v.%v %v\n", a, a, b, b, datacenter))
	}
	if err := writer.Flush(); err != nil {
		panic(err)
	}
	if err := file.Close(); err != nil {
		panic(err)
	}
}
