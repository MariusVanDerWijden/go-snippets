package main

import (
	"encoding/binary"
	"fmt"
	"reflect"

	bls "github.com/kilic/bls12-381"
)

func main() {

	eng := bls.NewEngine()
	r := reflect.ValueOf(eng)
	s := binary.TotalSize(r)
	fmt.Println(s)
}
