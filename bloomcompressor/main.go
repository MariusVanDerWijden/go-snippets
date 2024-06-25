package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	bloom "github.com/holiman/bloomfilter/v2"
)

func main() {
	size := uint64(128 * 1024 * 1024)
	funcs := uint64(8)
	filter, err := bloom.New(size, funcs)
	if err != nil {
		panic(err)
	}

	var (
		lines      int
		duplicates int
		start      = time.Now()

		writer  = bufio.NewWriter(os.Stdout)
		scanner = bufio.NewScanner(os.Stdin)
	)
	for scanner.Scan() {
		lines++
		line := scanner.Bytes()
		hash := lineToHash(line)
		if !filter.ContainsHash(hash) {
			fmt.Println(string(line))
			filter.AddHash(hash)
		} else {
			duplicates++
		}
	}
	writer.Flush()
	fmt.Printf("Found %v duplicates in %v lines in %v\n", duplicates, lines, time.Since(start))
}

var hasher = crypto.NewKeccakState()

// hashes the line using keccak, todo use faster hash
func lineToHash(line []byte) uint64 {
	hash := crypto.HashData(hasher, line)
	return binary.BigEndian.Uint64(hash[:])
}
