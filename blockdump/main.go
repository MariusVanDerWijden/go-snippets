package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/goevmlab/ops"
)

type line struct {
	Root []byte `json:"root"`
	Code string `json:"code"`
}

func main() {
	input, err := os.Open("/home/matematik/go/src/github.com/ethereum/go-ethereum/build/bin/blockdump")
	if err != nil {
		fmt.Print(err)
	}
	scanner := bufio.NewScanner(input)
	var countMap [256]int
	var k int
	for scanner.Scan() {
		data := scanner.Bytes()
		var elem line
		err := json.Unmarshal(data, &elem)
		if err != nil {
			continue
		}
		fmt.Println(k)
		k++
		if len(elem.Code) > 0 {
			code := common.FromHex(elem.Code)
			for _, el := range code {
				countMap[el]++
			}
		}
	}
	for i, el := range countMap {
		op := ops.OpCode(i)
		fmt.Printf("%d: %d: %v\n", i, el, op.String())
	}
}
