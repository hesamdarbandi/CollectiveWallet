// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"AllowanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MoneyReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MoneySent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"reduceAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sendMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600e575f80fd5b503380603357604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b603a81603f565b50608e565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6107748061009b5f395ff3fe60806040526004361061009d575f3560e01c8063715018a611610062578063715018a61461017e5780638da5cb5b146101925780638fdf08df146101b85780639b96eece146101d7578063ee4ae2c9146101f6578063f2fde38b14610215575f80fd5b806312065fe0146100dd578063310ec4a7146100ff57806339509351146101205780633e5beab91461013f5780636f9fb98a1461016a575f80fd5b366100d95760405134815233907f27b15ed4cf832749ed39f33a64e4707ed60a761485e41ffec7343ecaddc0c02a9060200160405180910390a2005b5f80fd5b3480156100e8575f80fd5b5033315b6040519081526020015b60405180910390f35b34801561010a575f80fd5b5061011e6101193660046106b2565b610234565b005b34801561012b575f80fd5b5061011e61013a3660046106b2565b6102ac565b34801561014a575f80fd5b506100ec6101593660046106dc565b60016020525f908152604090205481565b348015610175575f80fd5b506100ec610343565b348015610189575f80fd5b5061011e610351565b34801561019d575f80fd5b505f546040516001600160a01b0390911681526020016100f6565b3480156101c3575f80fd5b5061011e6101d23660046106b2565b610364565b3480156101e2575f80fd5b506100ec6101f13660046106dc565b6103f2565b348015610201575f80fd5b5061011e6102103660046106b2565b610409565b348015610220575f80fd5b5061011e61022f3660046106dc565b6105c4565b61023c610601565b6001600160a01b0382165f818152600160209081526040918290205482519081529081018490523392917f3691d1a86d99355e52b689ca70a7bdf6d80763237a6aa06e5fa43964eac7244b910160405180910390a36001600160a01b039091165f90815260016020526040902055565b6102b4610601565b6001600160a01b0382165f818152600160205260409020543391907f3691d1a86d99355e52b689ca70a7bdf6d80763237a6aa06e5fa43964eac7244b906102fb8582610712565b6040805192835260208301919091520160405180910390a36001600160a01b0382165f908152600160205260408120805483929061033a908490610712565b90915550505050565b5f61034c610601565b504790565b610359610601565b6103625f61062d565b565b61036c610601565b6001600160a01b0382165f818152600160205260409020543391907f3691d1a86d99355e52b689ca70a7bdf6d80763237a6aa06e5fa43964eac7244b906103b3858261072b565b6040805192835260208301919091520160405180910390a36001600160a01b0382165f908152600160205260408120805483929061033a90849061072b565b5f6103fb610601565b506001600160a01b03163190565b61041161067c565b6104625760405162461bcd60e51b815260206004820152601a60248201527f556e617574686f72697a656420746f2073656e64206d6f6e657900000000000060448201526064015b60405180910390fd5b804710156104b25760405162461bcd60e51b815260206004820152601b60248201527f546865726520617265206e6f7420656e6f75676820666f756e647300000000006044820152606401610459565b6001600160a01b0382165f9081526001602052604090205481111561053f5760405162461bcd60e51b815260206004820152603c60248201527f41737369676e656420616c6c6f77616e6365206973206e6f7420656e6f75676860448201527f20746f20706572666f726d2074686973207472616e73616374696f6e000000006064820152608401610459565b6105498282610364565b816001600160a01b03167f7f51d406915971d4ac1c91af96be5187ea6ab64753785aad519a533def80a41e8260405161058491815260200190565b60405180910390a26040516001600160a01b0383169082156108fc029083905f818181858888f193505050501580156105bf573d5f803e3d5ffd5b505050565b6105cc610601565b6001600160a01b0381166105f557604051631e4fbdf760e01b81525f6004820152602401610459565b6105fe8161062d565b50565b5f546001600160a01b031633146103625760405163118cdaa760e01b8152336004820152602401610459565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f3361068f5f546001600160a01b031690565b6001600160a01b031614905090565b6001600160a01b03811681146105fe575f80fd5b5f80604083850312156106c3575f80fd5b82356106ce8161069e565b946020939093013593505050565b5f602082840312156106ec575f80fd5b81356106f78161069e565b9392505050565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610725576107256106fe565b92915050565b81810381811115610725576107256106fe56fea2646970667358221220bcac6b9f6c344ba1f855fdf81527e7763cc986aefc9370627de0328c6885c66464736f6c63430008190033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0x3e5beab9.
//
// Solidity: function allowance(address ) view returns(uint256)
func (_Contract *ContractCaller) Allowance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "allowance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0x3e5beab9.
//
// Solidity: function allowance(address ) view returns(uint256)
func (_Contract *ContractSession) Allowance(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Allowance(&_Contract.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0x3e5beab9.
//
// Solidity: function allowance(address ) view returns(uint256)
func (_Contract *ContractCallerSession) Allowance(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Allowance(&_Contract.CallOpts, arg0)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Contract *ContractCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Contract *ContractSession) GetBalance() (*big.Int, error) {
	return _Contract.Contract.GetBalance(&_Contract.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Contract *ContractCallerSession) GetBalance() (*big.Int, error) {
	return _Contract.Contract.GetBalance(&_Contract.CallOpts)
}

// GetBalanceOf is a free data retrieval call binding the contract method 0x9b96eece.
//
// Solidity: function getBalanceOf(address _address) view returns(uint256)
func (_Contract *ContractCaller) GetBalanceOf(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getBalanceOf", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceOf is a free data retrieval call binding the contract method 0x9b96eece.
//
// Solidity: function getBalanceOf(address _address) view returns(uint256)
func (_Contract *ContractSession) GetBalanceOf(_address common.Address) (*big.Int, error) {
	return _Contract.Contract.GetBalanceOf(&_Contract.CallOpts, _address)
}

// GetBalanceOf is a free data retrieval call binding the contract method 0x9b96eece.
//
// Solidity: function getBalanceOf(address _address) view returns(uint256)
func (_Contract *ContractCallerSession) GetBalanceOf(_address common.Address) (*big.Int, error) {
	return _Contract.Contract.GetBalanceOf(&_Contract.CallOpts, _address)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_Contract *ContractCaller) GetContractBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getContractBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_Contract *ContractSession) GetContractBalance() (*big.Int, error) {
	return _Contract.Contract.GetContractBalance(&_Contract.CallOpts)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_Contract *ContractCallerSession) GetContractBalance() (*big.Int, error) {
	return _Contract.Contract.GetContractBalance(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _beneficiary, uint256 _amount) returns()
func (_Contract *ContractTransactor) IncreaseAllowance(opts *bind.TransactOpts, _beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "increaseAllowance", _beneficiary, _amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _beneficiary, uint256 _amount) returns()
func (_Contract *ContractSession) IncreaseAllowance(_beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.IncreaseAllowance(&_Contract.TransactOpts, _beneficiary, _amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _beneficiary, uint256 _amount) returns()
func (_Contract *ContractTransactorSession) IncreaseAllowance(_beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.IncreaseAllowance(&_Contract.TransactOpts, _beneficiary, _amount)
}

// ReduceAllowance is a paid mutator transaction binding the contract method 0x8fdf08df.
//
// Solidity: function reduceAllowance(address _beneficiary, uint256 _amount) returns()
func (_Contract *ContractTransactor) ReduceAllowance(opts *bind.TransactOpts, _beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "reduceAllowance", _beneficiary, _amount)
}

// ReduceAllowance is a paid mutator transaction binding the contract method 0x8fdf08df.
//
// Solidity: function reduceAllowance(address _beneficiary, uint256 _amount) returns()
func (_Contract *ContractSession) ReduceAllowance(_beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ReduceAllowance(&_Contract.TransactOpts, _beneficiary, _amount)
}

// ReduceAllowance is a paid mutator transaction binding the contract method 0x8fdf08df.
//
// Solidity: function reduceAllowance(address _beneficiary, uint256 _amount) returns()
func (_Contract *ContractTransactorSession) ReduceAllowance(_beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ReduceAllowance(&_Contract.TransactOpts, _beneficiary, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// SendMoney is a paid mutator transaction binding the contract method 0xee4ae2c9.
//
// Solidity: function sendMoney(address _to, uint256 _amount) returns()
func (_Contract *ContractTransactor) SendMoney(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sendMoney", _to, _amount)
}

// SendMoney is a paid mutator transaction binding the contract method 0xee4ae2c9.
//
// Solidity: function sendMoney(address _to, uint256 _amount) returns()
func (_Contract *ContractSession) SendMoney(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SendMoney(&_Contract.TransactOpts, _to, _amount)
}

// SendMoney is a paid mutator transaction binding the contract method 0xee4ae2c9.
//
// Solidity: function sendMoney(address _to, uint256 _amount) returns()
func (_Contract *ContractTransactorSession) SendMoney(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SendMoney(&_Contract.TransactOpts, _to, _amount)
}

// SetAllowance is a paid mutator transaction binding the contract method 0x310ec4a7.
//
// Solidity: function setAllowance(address _beneficiary, uint256 _amount) returns()
func (_Contract *ContractTransactor) SetAllowance(opts *bind.TransactOpts, _beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setAllowance", _beneficiary, _amount)
}

// SetAllowance is a paid mutator transaction binding the contract method 0x310ec4a7.
//
// Solidity: function setAllowance(address _beneficiary, uint256 _amount) returns()
func (_Contract *ContractSession) SetAllowance(_beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetAllowance(&_Contract.TransactOpts, _beneficiary, _amount)
}

// SetAllowance is a paid mutator transaction binding the contract method 0x310ec4a7.
//
// Solidity: function setAllowance(address _beneficiary, uint256 _amount) returns()
func (_Contract *ContractTransactorSession) SetAllowance(_beneficiary common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetAllowance(&_Contract.TransactOpts, _beneficiary, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactorSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// ContractAllowanceChangedIterator is returned from FilterAllowanceChanged and is used to iterate over the raw logs and unpacked data for AllowanceChanged events raised by the Contract contract.
type ContractAllowanceChangedIterator struct {
	Event *ContractAllowanceChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAllowanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAllowanceChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAllowanceChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAllowanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAllowanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAllowanceChanged represents a AllowanceChanged event raised by the Contract contract.
type ContractAllowanceChanged struct {
	Beneficiary common.Address
	Sender      common.Address
	PrevAmount  *big.Int
	NewAmount   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllowanceChanged is a free log retrieval operation binding the contract event 0x3691d1a86d99355e52b689ca70a7bdf6d80763237a6aa06e5fa43964eac7244b.
//
// Solidity: event AllowanceChanged(address indexed beneficiary, address indexed sender, uint256 prevAmount, uint256 newAmount)
func (_Contract *ContractFilterer) FilterAllowanceChanged(opts *bind.FilterOpts, beneficiary []common.Address, sender []common.Address) (*ContractAllowanceChangedIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AllowanceChanged", beneficiaryRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractAllowanceChangedIterator{contract: _Contract.contract, event: "AllowanceChanged", logs: logs, sub: sub}, nil
}

// WatchAllowanceChanged is a free log subscription operation binding the contract event 0x3691d1a86d99355e52b689ca70a7bdf6d80763237a6aa06e5fa43964eac7244b.
//
// Solidity: event AllowanceChanged(address indexed beneficiary, address indexed sender, uint256 prevAmount, uint256 newAmount)
func (_Contract *ContractFilterer) WatchAllowanceChanged(opts *bind.WatchOpts, sink chan<- *ContractAllowanceChanged, beneficiary []common.Address, sender []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AllowanceChanged", beneficiaryRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAllowanceChanged)
				if err := _Contract.contract.UnpackLog(event, "AllowanceChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAllowanceChanged is a log parse operation binding the contract event 0x3691d1a86d99355e52b689ca70a7bdf6d80763237a6aa06e5fa43964eac7244b.
//
// Solidity: event AllowanceChanged(address indexed beneficiary, address indexed sender, uint256 prevAmount, uint256 newAmount)
func (_Contract *ContractFilterer) ParseAllowanceChanged(log types.Log) (*ContractAllowanceChanged, error) {
	event := new(ContractAllowanceChanged)
	if err := _Contract.contract.UnpackLog(event, "AllowanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMoneyReceivedIterator is returned from FilterMoneyReceived and is used to iterate over the raw logs and unpacked data for MoneyReceived events raised by the Contract contract.
type ContractMoneyReceivedIterator struct {
	Event *ContractMoneyReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractMoneyReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMoneyReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractMoneyReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractMoneyReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMoneyReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMoneyReceived represents a MoneyReceived event raised by the Contract contract.
type ContractMoneyReceived struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMoneyReceived is a free log retrieval operation binding the contract event 0x27b15ed4cf832749ed39f33a64e4707ed60a761485e41ffec7343ecaddc0c02a.
//
// Solidity: event MoneyReceived(address indexed from, uint256 amount)
func (_Contract *ContractFilterer) FilterMoneyReceived(opts *bind.FilterOpts, from []common.Address) (*ContractMoneyReceivedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "MoneyReceived", fromRule)
	if err != nil {
		return nil, err
	}
	return &ContractMoneyReceivedIterator{contract: _Contract.contract, event: "MoneyReceived", logs: logs, sub: sub}, nil
}

// WatchMoneyReceived is a free log subscription operation binding the contract event 0x27b15ed4cf832749ed39f33a64e4707ed60a761485e41ffec7343ecaddc0c02a.
//
// Solidity: event MoneyReceived(address indexed from, uint256 amount)
func (_Contract *ContractFilterer) WatchMoneyReceived(opts *bind.WatchOpts, sink chan<- *ContractMoneyReceived, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "MoneyReceived", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMoneyReceived)
				if err := _Contract.contract.UnpackLog(event, "MoneyReceived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMoneyReceived is a log parse operation binding the contract event 0x27b15ed4cf832749ed39f33a64e4707ed60a761485e41ffec7343ecaddc0c02a.
//
// Solidity: event MoneyReceived(address indexed from, uint256 amount)
func (_Contract *ContractFilterer) ParseMoneyReceived(log types.Log) (*ContractMoneyReceived, error) {
	event := new(ContractMoneyReceived)
	if err := _Contract.contract.UnpackLog(event, "MoneyReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMoneySentIterator is returned from FilterMoneySent and is used to iterate over the raw logs and unpacked data for MoneySent events raised by the Contract contract.
type ContractMoneySentIterator struct {
	Event *ContractMoneySent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractMoneySentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMoneySent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractMoneySent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractMoneySentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMoneySentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMoneySent represents a MoneySent event raised by the Contract contract.
type ContractMoneySent struct {
	Beneficiary common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMoneySent is a free log retrieval operation binding the contract event 0x7f51d406915971d4ac1c91af96be5187ea6ab64753785aad519a533def80a41e.
//
// Solidity: event MoneySent(address indexed beneficiary, uint256 amount)
func (_Contract *ContractFilterer) FilterMoneySent(opts *bind.FilterOpts, beneficiary []common.Address) (*ContractMoneySentIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "MoneySent", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &ContractMoneySentIterator{contract: _Contract.contract, event: "MoneySent", logs: logs, sub: sub}, nil
}

// WatchMoneySent is a free log subscription operation binding the contract event 0x7f51d406915971d4ac1c91af96be5187ea6ab64753785aad519a533def80a41e.
//
// Solidity: event MoneySent(address indexed beneficiary, uint256 amount)
func (_Contract *ContractFilterer) WatchMoneySent(opts *bind.WatchOpts, sink chan<- *ContractMoneySent, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "MoneySent", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMoneySent)
				if err := _Contract.contract.UnpackLog(event, "MoneySent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMoneySent is a log parse operation binding the contract event 0x7f51d406915971d4ac1c91af96be5187ea6ab64753785aad519a533def80a41e.
//
// Solidity: event MoneySent(address indexed beneficiary, uint256 amount)
func (_Contract *ContractFilterer) ParseMoneySent(log types.Log) (*ContractMoneySent, error) {
	event := new(ContractMoneySent)
	if err := _Contract.contract.UnpackLog(event, "MoneySent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
