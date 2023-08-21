package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mariusvanderwijden/threadpool"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"log"
	"math/big"
	"os"
	"strings"
)

var path = "spanish.txt"
var rpc = "http://localhost:8545"

func main() {
	wordlist := readWordlist()
	client, err := ethclient.Dial(rpc)
	if err != nil {
		panic(err)
	}
	pool := threadpool.NewThreadPool(24)
	for i := 0; i < len(wordlist); i++ {
		pool.Get(1)
		go func(i int) {
			var str []string
			for k := 0; k < 12; k++ {
				str = append(str, wordlist[i])
			}
			test(str, client)
			pool.Put(1)
		}(i)

	}
}

func readWordlist() []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func test(str []string, client *ethclient.Client) {
	mnemonic := strings.Join(str, " ")
	//fmt.Println("mnemonic:", mnemonic)
	seed := bip39.NewSeed(mnemonic, "") // Here you can choose to pass in the specified password or empty string , Different passwords generate different mnemonics

	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0") // The last digit is the address of the same mnemonic word id, from 0 Start , The same mnemonic can produce unlimited addresses
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	address := account.Address
	privateKey, _ := wallet.PrivateKeyHex(account)
	//publicKey, _ := wallet.PublicKeyHex(account)

	//fmt.Println("address0:", address)      // id by 0 The address of your wallet
	//	fmt.Println("publicKey:", publicKey)   //  Public key

	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1") // Generate id by 1 The address of your wallet
	account, err = wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	bal, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		panic(err)
	}
	if bal.Cmp(big.NewInt(0)) != 0 {

		fmt.Println("privateKey:", privateKey) //  Private key
		fmt.Printf("%v %v %v", mnemonic, address, bal)
	}
}
