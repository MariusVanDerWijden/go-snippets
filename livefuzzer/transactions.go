package main

import (
	"github.com/MariusVanDerWijden/FuzzyVM/filler"
	"github.com/MariusVanDerWijden/FuzzyVM/generator"
)

func randomCode(f *filler.Filler) []byte {
	_, code := generator.GenerateProgram(f)
	return code
}
