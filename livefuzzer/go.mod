module github.com/MariusVanDerWijden/go-snippets/livefuzzer

go 1.16

require (
	github.com/MariusVanDerWijden/FuzzyVM v0.0.0-20210428104145-c6788db9a2eb
	github.com/ethereum/go-ethereum v1.10.3
)

replace github.com/ethereum/go-ethereum => ./../../../ethereum/go-ethereum/
