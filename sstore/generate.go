package main

// wget -nc "https://github.com/ethereum/solidity/releases/download/v0.8.20/solc-static-linux"
// chmod +x ./solc-static-linux

//go:generate ./solc-static-linux --overwrite --bin --abi -o ./ contract.sol
//go:generate abigen --pkg main --type WarmSStore --abi WarmSStore.abi --bin WarmSStore.bin --out WarmSStore.go
