package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {

	ctx := context.Background()
	client, err := ethclient.Dial("https://rpc.ankr.com/eth_goerli")
	if err != nil {
		log.Fatal(err)
	}

	c, err := client.NetworkID(ctx)
	privateKey, err := crypto.HexToECDSA("330134aeeb72e159ef2d53332a0778d7534eea949b137f03a35ee34817002c1b")

	fmt.Println(c)
	fmt.Println(privateKey)
}
