module github.com/go-snippets/geth-test

go 1.13

require (
	github.com/ethereum/go-ethereum v1.9.12
	github.com/stretchr/testify v1.5.1
)

replace github.com/ethereum/go-ethereum => ./../../ethereum/go-ethereum/
