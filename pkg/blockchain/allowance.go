package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

const (
	SetAction      = "set"
	GetAction      = "get"
	IncreaseAction = "increase"
	ReduceAction   = "reduce"
	SendAction     = "send"
	ReceiveAction  = "receive"
)

type Allowance struct {
	PrivateKey      string
	ContractAddress string
}

func NewAllowance(privateKey string, contractAddress string) *Allowance {
	return &Allowance{
		PrivateKey:      privateKey,
		ContractAddress: contractAddress,
	}
}

func (r *Allowance) GetAllowance(ctx context.Context, client *ethclient.Client, beneficiaryAddress string) (int64, error) {
	contract, err := getContract(ctx, client, r.ContractAddress)
	if err != nil {
		return 0, err
	}

	address := common.HexToAddress(beneficiaryAddress)
	amount, err := contract.Allowance(&bind.CallOpts{Pending: false, Context: ctx}, address)
	return weiToEther(amount).Int64(), nil
}

func (r *Allowance) ChangeAllowance(ctx context.Context, client *ethclient.Client, action string, target string, amount int64) error {
	contract, err := getContract(ctx, client, r.ContractAddress)
	if err != nil {
		return err
	}

	signer, err := getSigner(ctx, client)
	if err != nil {
		return err
	}

	var tx *types.Transaction
	var txErr error
	targetAddress := common.HexToAddress(target)

	var operation string
	switch action {
	case SetAction:
		tx, txErr = contract.SetAllowance(signer, targetAddress, etherToWei(big.NewInt(amount)))
		operation = "set_allowance"
	case IncreaseAction:
		tx, txErr = contract.IncreaseAllowance(signer, targetAddress, etherToWei(big.NewInt(amount)))
		operation = "increase_allowance"
	case ReduceAction:
		tx, txErr = contract.ReduceAllowance(signer, targetAddress, etherToWei(big.NewInt(amount)))
		operation = "reduce_allowance"
	}
	if txErr != nil {
		return txErr
	}
	receipt, err := bind.WaitMined(ctx, client, tx)
	if receipt.Status != types.ReceiptStatusSuccessful || err != nil {
		return err
	}
	processTransaction(ctx, tx, operation)

	return nil
}
