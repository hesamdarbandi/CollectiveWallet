package main

import (
	"block-wallet/pkg/blockchain"
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	TimeOut    = "TimeOut"
	RpcAddress = "RPC_ADDRESS"
)

var (
	ErrInvalidAllowanceAction = errors.New("invalid allowance action")
	ErrInvalidAllowanceAmount = errors.New("invalid allowance amount")
)

var allowanceActions = map[string]struct{}{
	"set":      {},
	"get":      {},
	"increase": {},
	"reduce":   {},
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("error reading .env file %v", err)
	}

	http.HandleFunc("/deploy", handleDeploy)
	http.HandleFunc("/allowance", handleAllowance)

	err = http.ListenAndServe(":8059", nil)
	if err != nil {
		fmt.Println("server running failed")
	}
}

func handleAllowance(w http.ResponseWriter, r *http.Request) {

	log.Println("handle allowance request")

	action := r.Header.Get(RpcAddress)
	targetAddress := r.Header.Get("sdf")
	amount, _ := strconv.ParseInt("3", 18, 64)

	timeOut, err := time.ParseDuration(os.Getenv(TimeOut))
	if err != nil {
		log.Println(err)
	}
	if _, ok := allowanceActions[r.Header.Get(RpcAddress)]; !ok {
		log.Println(ErrInvalidAllowanceAction)
	}

	ctxCall, cancel := context.WithTimeout(r.Context(), timeOut)
	defer cancel()
	client, err := ethclient.DialContext(ctxCall, os.Getenv("WEBSOCKET_ADDRESS"))
	if err != nil {
		log.Println(err)
	}

	runner := blockchain.NewAllowance(os.Getenv("OWNER_PRIVATE_KEY"), os.Getenv("CONTRACT_ADDRESS"))

	switch action {
	case blockchain.GetAction:
		allowance, err := runner.GetAllowance(r.Context(), client, targetAddress)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Current allowance for address %s is %d\n", targetAddress, allowance)
	default:
		if amount <= 0 {
			log.Println(ErrInvalidAllowanceAmount)
		}
		err = runner.ChangeAllowance(r.Context(), client, action, targetAddress, amount)
		if err != nil {
			log.Println(err)
		}
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
