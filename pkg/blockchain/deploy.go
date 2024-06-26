package blockchain

import (
	contracts "block-wallet/contracts/implementions"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Deployer struct {
	Address     common.Address
	Transaction *types.Transaction
	Contract    *contracts.Contract
}

func NewDeployer() *Deployer {
	return &Deployer{}
}

func (d *Deployer) Deploy(ctx context.Context, client *ethclient.Client) error {
	signer, err := getSigner(ctx, client)
	if err != nil {
		return err
	}
	address, tx, contract, err := contracts.DeployContract(signer, client)
	if err != nil {
		return err
	}
	_, err = bind.WaitDeployed(ctx, client, tx)
	if err != nil {
		return err
	}
	d.Address = address
	d.Transaction = tx
	d.Contract = contract
	return nil
}

// ContractAddress returns the contract address
func (d *Deployer) ContractAddress() string {
	return d.Address.Hex()
}
