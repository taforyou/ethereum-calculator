package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	store "./contracts" // for demo
)

func main() {

	client, err := ethclient.Dial("https://goerli.infura.io")

	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x447ff6F6db540De0eEa497F340984cBFCE02fe0A")

	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	callOpts := &bind.CallOpts{}

	var a, b uint64
	a = 10
	b = 2

	result, err := instance.Divide(callOpts, big.NewInt(int64(a)), big.NewInt(int64(b)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
