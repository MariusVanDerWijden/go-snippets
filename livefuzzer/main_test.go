package main

import (
	"fmt"
	"testing"
)

func TestCreateAddress(t *testing.T) {
	keys, addrs := createAddresses(100)
	fmt.Println(`var (
		keys = []string{`)
	for _, k := range keys {
		fmt.Printf("\"%v\",\n", k)
	}
	fmt.Println("}")
	fmt.Println(`addrs = []string{`)
	for _, a := range addrs {
		fmt.Printf("\"%v\",\n", a)
	}
	fmt.Println("}")
	t.Fail()
}
