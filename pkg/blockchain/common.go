package blockchain

import (
	contracts "block-wallet/contracts/implementions"
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"log"
	"math/big"
	"os"
	"regexp"
	"strconv"
)

var (
	ErrInvalidKey             = errors.New("invalid key")
	ErrInvalidAddress         = errors.New("invalid address")
	ErrInvalidContractAddress = errors.New("invalid contract address")
)

func getSigner(ctx context.Context, client *ethclient.Client) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(os.Getenv("PrivateKey"))
	if err != nil {
		return nil, err
	}
	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, ErrInvalidKey
	}

	address := crypto.PubkeyToAddress(*publicKey)
	nonce, err := client.PendingNonceAt(ctx, address)
	if err != nil {
		return nil, err
	}
	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	signer, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	signer.Nonce = big.NewInt(int64(nonce))
	signer.Value = big.NewInt(getEnvInt("WeiFounds"))
	signer.GasLimit = uint64(getEnvInt("GasLimit"))
	signer.GasPrice = gasPrice

	return signer, nil
}

func getContract(ctx context.Context, client *ethclient.Client, contractAddress string) (*contracts.Contract, error) {
	err := validateContractAddress(ctx, client, contractAddress)
	if err != nil {
		return nil, err
	}
	contract, err := contracts.NewContract(common.HexToAddress(contractAddress), client)
	if err != nil {
		return nil, err
	}
	return contract, nil
}

func validateContractAddress(ctx context.Context, client *ethclient.Client, address string) error {
	if err := validateAddress(address); err != nil {
		return err
	}
	contractAddress := common.HexToAddress(address)
	bytecode, err := client.CodeAt(ctx, contractAddress, nil)
	if err != nil {
		return err
	}
	if len(bytecode) > 0 {
		return nil
	}
	return ErrInvalidContractAddress
}

func validateAddress(address string) error {
	regex := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if ok := regex.MatchString(address); !ok {
		return ErrInvalidAddress
	}
	return nil
}

func getEnvInt(key string) int64 {

	value := os.Getenv(key)
	res, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return res
}

func etherToWei(eth *big.Int) *big.Int {
	return new(big.Int).Mul(eth, big.NewInt(params.Ether))
}

func weiToEther(wei *big.Int) *big.Int {
	return new(big.Int).Div(wei, big.NewInt(params.Ether))
}
