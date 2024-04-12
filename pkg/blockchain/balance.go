package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	AddressBalance  = "address"
	ContractBalance = "contract"
)

type Balance struct {
	privateKey      string
	contractAddress string
}

func NewBalanceRunner(privateKey string, contractAddress string) *Balance {
	return &Balance{
		privateKey:      privateKey,
		contractAddress: contractAddress,
	}
}

func (b *Balance) GetContractBalance(ctx context.Context, client *ethclient.Client) (int64, error) {
	value, err := client.BalanceAt(ctx, common.HexToAddress(b.contractAddress), nil)
	if err != nil {
		return 0, err
	}
	return weiToEther(value).Int64(), nil
}

// GetAddressBalance returns the balance of a given address
func (b *Balance) GetAddressBalance(ctx context.Context, client *ethclient.Client, address string) (int64, error) {
	value, err := client.BalanceAt(ctx, common.HexToAddress(address), nil)
	if err != nil {
		return 0, err
	}
	return weiToEther(value).Int64(), nil
}
