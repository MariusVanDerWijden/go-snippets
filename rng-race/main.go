package main

import (
	"math/rand"
)

func main() {
	source := rand.New(rand.NewSource(123))
	go func() {
		source.Intn(1234)
	}()
	source.Intn(1234)
}
