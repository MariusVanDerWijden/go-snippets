module github.com/go-snippets

go 1.15

replace github.com/ethereum/go-ethereum => /home/matematik/go/src/github.com/ethereum/go-ethereum

replace github.com/holiman/goevmlab => /home/matematik/ethereum/goevmlab

require (
	github.com/MariusVanDerWijden/FuzzyVM v0.0.0-20230329161016-f7146b6f1dca
	github.com/ethereum/go-ethereum v1.11.5
	github.com/golang/snappy v0.0.5-0.20220116011046-fa5810519dcb
	github.com/google/gofuzz v1.1.1-0.20200604201612-c04b05f3adfa
	github.com/holiman/goevmlab v0.0.0-20230316064510-98c61355fce0
	github.com/mariusvanderwijden/go-snippets v0.0.0-20220929094507-5288846e1a92
	github.com/stretchr/testify v1.8.0
	github.com/supranational/blst v0.3.8-0.20220526154634-513d2456b344
	golang.org/x/net v0.8.0
	golang.org/x/sys v0.6.0
)
