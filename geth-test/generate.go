package geth

//go:generate ./solc-static-linux-8-7 --overwrite --bin --abi -o ./ Eventer.sol
//go:generate ./solc-static-linux-8-7 --overwrite --bin --abi -o ./ keyinfo.sol
//go:generate abigen --pkg geth --type Eventer --abi Eventer.abi --bin Eventer.bin --out eventer.go
//go:generate abigen --pkg geth --type keyinfo --abi KeyInfo.abi --bin KeyInfo.bin --out keyinfo.go
//go:generate abigen --pkg geth --type array --abi Array.abi --bin Array.bin --out array.go
//go:generate abigen --pkg geth --type ReceiveFallback --abi ReceiveFallback.abi --bin ReceiveFallback.bin --out receivefallback.go
//go:generate abigen --pkg geth --type TupleTest --abi TupleTest.abi --bin TupleTest.bin --out tupletest.go
