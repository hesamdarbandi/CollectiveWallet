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
	address     common.Address
	transaction *types.Transaction
	contract    *contracts.Contract
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
	d.address = address
	d.transaction = tx
	d.contract = contract
	return nil
}

// ContractAddress returns the contract address
func (d *Deployer) ContractAddress() string {
	return d.address.Hex()
}
