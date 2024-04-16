package main

import (
	"block-wallet/pkg/blockchain"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	TimeOut    = "TimeOut"
	RpcAddress = "RPC_ADDRESS"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("error reading .env file %v", err)
	}

	http.HandleFunc("/deploy", handleDeploy)

	err = http.ListenAndServe(":8059", nil)
	if err != nil {
		fmt.Println("server running failed")
	}
}

func handleDeploy(w http.ResponseWriter, r *http.Request) {

	log.Println("deploying contract")
	timeOut, err := time.ParseDuration(os.Getenv(TimeOut))
	if err != nil {
		log.Println(err)
	}

	ctx, cancel := context.WithTimeout(r.Context(), timeOut)
	defer cancel()
	client, err := ethclient.DialContext(ctx, os.Getenv(RpcAddress))
	if err != nil {
		log.Println(err)
	}
	deployer := blockchain.NewDeployer()
	err = deployer.Deploy(ctx, client)
	if err != nil {
		log.Println(err)
	}

	log.Printf("contract deployed at address %s\n", deployer.ContractAddress())
}
