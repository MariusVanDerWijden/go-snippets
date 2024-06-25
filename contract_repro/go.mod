module github.com/MariusVanDerWijden/go-snippets/contract_repro

go 1.16

require (
	github.com/MariusVanDerWijden/go-ethereum v1.8.22
	github.com/ethereum/go-ethereum v1.10.6-0.20210715235240-f05419f0fb8c
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d
)

replace github.com/ethereum/go-ethereum => ../../../ethereum/go-ethereum
