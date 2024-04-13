package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Owner struct {
	PrivateKey      string
	ContractAddress string
}

func NewOwnerRunner(privateKey string, contractAddress string) *Owner {
	return &Owner{
		PrivateKey:      privateKey,
		ContractAddress: contractAddress,
	}
}

func (o *Owner) GetOwner(ctx context.Context, client *ethclient.Client) (string, error) {
	contract, err := getContract(ctx, client, o.ContractAddress)
	if err != nil {
		return "", err
	}

	ownerAddress, err := contract.Owner(&bind.CallOpts{Pending: false, Context: ctx})
	if err != nil {
		return "", err
	}
	return ownerAddress.Hex(), nil
}

func (o *Owner) TransferOwner(ctx context.Context, client *ethclient.Client, targetAddress string) error {
	contract, err := getContract(ctx, client, o.ContractAddress)
	if err != nil {
		return err
	}
	signer, err := getSigner(ctx, client)
	if err != nil {
		return err
	}

	tx, txErr := contract.TransferOwnership(signer, common.HexToAddress(targetAddress))
	if txErr != nil {
		return txErr
	}
	receipt, err := bind.WaitMined(ctx, client, tx)
	if receipt.Status != types.ReceiptStatusSuccessful || err != nil {
		return err
	}
	processTransaction(ctx, tx, "transfer owner")
	return nil
}
