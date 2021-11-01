package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/MariusVanDerWijden/FuzzyVM/fuzzer"
)

func main() {
	path := os.Args[1]
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read %v files\n", len(infos))
	for _, info := range infos {
		filename := fmt.Sprintf("%v/%v", path, info.Name())
		input, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		convertCorpus(input)
	}
}

func convertCorpus(input []byte) {
	fuzzer.Fuzz(input)
}
