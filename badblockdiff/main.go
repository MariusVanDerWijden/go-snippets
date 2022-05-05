package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
)

var (
	folderA = "blocks"
	folderB = "blocks2"
	prefix  = "block_0x92b354bd-"
)

func main() {
	harmonizeNames(folderA)
	harmonizeNames(folderB)
	infos, err := ioutil.ReadDir(folderA)
	if err != nil {
		panic(err)
	}
	for _, info := range infos {
		diff(info.Name())
	}
}

func harmonizeNames(folder string) {
	infos, err := ioutil.ReadDir(folder)
	if err != nil {
		panic(err)
	}
	for _, info := range infos {
		splits := strings.Split(info.Name(), "-")
		if len(splits) < 4 {
			color.Red("Refusing to harmonize names")
			return
		}
		newName := splits[0]
		for i := 1; i < len(splits)-1; i++ {
			newName = fmt.Sprintf("%v-%v", newName, splits[i])
		}
		old := fmt.Sprintf("%v/%v", folder, info.Name())
		newName = fmt.Sprintf("%v/%v", folder, newName)
		if err := os.Rename(old, newName); err != nil {
			panic(err)
		}
	}
}

func diff(filename string) {
	color.Yellow(filename)
	nA := fmt.Sprintf("%v/%v", folderA, filename)
	nB := fmt.Sprintf("%v/%v", folderB, filename)

	fileA, err := os.Open(nA)
	if err != nil {
		panic(err)
	}

	fileB, err := os.Open(nB)
	if err != nil {
		panic(err)
	}

	scannerA := bufio.NewScanner(fileA)
	scannerB := bufio.NewScanner(fileB)

	bufA := make([]byte, 4*1024*1024)
	scannerA.Buffer(bufA, cap(bufA))

	bufB := make([]byte, 4*1024*1024)
	scannerA.Buffer(bufB, cap(bufB))

	var i int

	for scannerA.Scan() {
		scannerB.Scan()
		dataA := scannerA.Bytes()
		dataB := scannerB.Bytes()
		if !bytes.Equal(dataA, dataB) {
			i++
			if i > 10 {
				// Only print out the first 10 differing lines
				return
			}
			if strings.Contains(string(dataA), "\"time\":") {
				aS := strings.Split(string(dataA), "\"time\":")
				bS := strings.Split(string(dataB), "\"time\":")
				if aS[0] != bS[0] {
					color.Green(aS[0])
					color.Red(bS[0])
				}
			} else {
				color.Green(string(dataA))
				color.Red(string(dataB))
			}
		}
	}
}
