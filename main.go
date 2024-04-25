package main

import (
	"block-wallet/pkg/blockchain"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"
)

const (
	TimeOut    = "TimeOut"
	RpcAddress = "RPC_ADDRESS"
)

type AllowanceRequest struct {
	Address string  `json:"address"`
	Action  string  `json:"action"`
	Amount  big.Int `json:"amount"`
}

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

var transferActions = map[string]struct{}{
	"send":    {},
	"receive": {},
}

var ownershipActions = map[string]struct{}{
	"get":      {},
	"transfer": {},
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("error reading .env file %v", err)
	}

	http.HandleFunc("/deploy", handleDeploy)
	http.HandleFunc("/allowance", handleAllowance)
	http.HandleFunc("/transfer", HandleTransfer)
	http.HandleFunc("/ownership", HandleOwnerShip)

	err = http.ListenAndServe(":8059", nil)
	if err != nil {
		fmt.Println("server running failed")
	}
}

func HandleOwnerShip(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var request map[string]interface{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Println(err)
	}

	action := request["action"].(string)
	targetAddress := request["address"].(string)

	if _, ok := ownershipActions[action]; !ok {
		log.Println("invalid action")
	}

	timeOut, err := time.ParseDuration(os.Getenv(TimeOut))
	if err != nil {
		log.Println(err)
	}

	ctxCall, cancel := context.WithTimeout(r.Context(), timeOut)
	defer cancel()
	client, err := ethclient.DialContext(ctxCall, os.Getenv("WEBSOCKET_ADDRESS"))
	if err != nil {
		log.Println(err)
	}
	runner := blockchain.NewOwnerRunner(os.Getenv("OWNER_PRIVATE_KEY"), os.Getenv("RPC_ADDRESS"))

	switch action {
	case blockchain.GetAction:
		owner, err := runner.GetOwner(r.Context(), client)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Current owner address is %s\n", owner)
	default:
		if targetAddress == "" {
			log.Println("target address is empty")
		}
		err = runner.TransferOwner(r.Context(), client, targetAddress)
		if err != nil {
			log.Println(err)
		}
	}

	log.Println("onwership is done")
}

func HandleTransfer(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var request map[string]interface{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Println(err)
	}

	action := request["action"].(string)
	amount := request["amount"].(int64)
	targetAddress := request["address"].(string)

	if _, ok := transferActions[action]; !ok {
		log.Println(ErrInvalidAllowanceAction)
	}

	timeOut, err := time.ParseDuration(os.Getenv(TimeOut))
	if err != nil {
		log.Println(err)
	}

	ctxCall, cancel := context.WithTimeout(r.Context(), timeOut)
	defer cancel()
	client, err := ethclient.DialContext(ctxCall, os.Getenv("WEBSOCKET_ADDRESS"))
	if err != nil {
		log.Println(err)
	}
	runner := blockchain.NewTransfersRunner(os.Getenv("OWNER_PRIVATE_KEY"), os.Getenv("RPC_ADDRESS"))

	if amount <= 0 {
		log.Println(ErrInvalidAllowanceAmount)
	}

	switch action {
	case blockchain.SendAction:
		if targetAddress == "" {
			log.Println("target address is empty")
		}
		err = runner.Send(r.Context(), client, targetAddress, amount)
	case blockchain.ReceiveAction:
		err = runner.Receive(r.Context(), client, amount)
	}
	if err != nil {
		log.Println(err)
	}

	log.Println("transfer done")
}

func handleAllowance(w http.ResponseWriter, r *http.Request) {

	log.Println("handle allowance request")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	var request map[string]interface{}
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Println(err)
	}

	action := request["action"].(string)
	targetAddress := request["targetAddress"].(string)
	amount, _ := request["amount"].(int64)

	timeOut, err := time.ParseDuration(os.Getenv(TimeOut))
	if err != nil {
		log.Println(err)
	}
	if _, ok := allowanceActions[action]; !ok {
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
