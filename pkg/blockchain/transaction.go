package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type TransactionInfo struct {
	Operation string   `json:"operation"`
	Gas       uint64   `json:"gas"`
	GasPrice  *big.Int `json:"gasPrice"`
	Cost      *big.Int `json:"cost"`
}

func processTransaction(ctx context.Context, tx *types.Transaction, operation string) {
	if tx == nil {
		return
	}
	_ = TransactionInfo{
		Operation: operation,
		Gas:       tx.Gas(),
		GasPrice:  tx.GasPrice(),
		Cost:      tx.Cost(),
	}
	// TODO: process the whole transaction and use if for stats
}
