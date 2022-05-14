// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
)

// Pubkey is an auto generated low-level Go binding around an user-defined struct.
type Pubkey struct {
	Keytype uint8
	Status  uint8
	Key     []byte
}

// UserRegistryMetaData contains all meta data concerning the UserRegistry contract.
var UserRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"KeyType\",\"name\":\"keytype\",\"type\":\"uint8\"},{\"internalType\":\"KeyStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"addPubkey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllUsers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"keypos\",\"type\":\"uint8\"}],\"name\":\"getKey\",\"outputs\":[{\"components\":[{\"internalType\":\"KeyType\",\"name\":\"keytype\",\"type\":\"uint8\"},{\"internalType\":\"KeyStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"internalType\":\"structPubkey\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getKeyLen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getKeys\",\"outputs\":[{\"components\":[{\"internalType\":\"KeyType\",\"name\":\"keytype\",\"type\":\"uint8\"},{\"internalType\":\"KeyStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"internalType\":\"structPubkey[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getUser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"KeyType\",\"name\":\"keytype\",\"type\":\"uint8\"},{\"internalType\":\"KeyStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"newUser\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"keypos\",\"type\":\"uint8\"},{\"internalType\":\"KeyStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"updateKeyStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UserRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use UserRegistryMetaData.ABI instead.
var UserRegistryABI = UserRegistryMetaData.ABI

// UserRegistry is an auto generated Go binding around an Ethereum contract.
type UserRegistry struct {
	UserRegistryCaller     // Read-only binding to the contract
	UserRegistryTransactor // Write-only binding to the contract
	UserRegistryFilterer   // Log filterer for contract events
}

// UserRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserRegistrySession struct {
	Contract     *UserRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserRegistryCallerSession struct {
	Contract *UserRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// UserRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserRegistryTransactorSession struct {
	Contract     *UserRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// UserRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserRegistryRaw struct {
	Contract *UserRegistry // Generic contract binding to access the raw methods on
}

// UserRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserRegistryCallerRaw struct {
	Contract *UserRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// UserRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserRegistryTransactorRaw struct {
	Contract *UserRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserRegistry creates a new instance of UserRegistry, bound to a specific deployed contract.
func NewUserRegistry(address common.Address, backend bind.ContractBackend) (*UserRegistry, error) {
	contract, err := bindUserRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserRegistry{UserRegistryCaller: UserRegistryCaller{contract: contract}, UserRegistryTransactor: UserRegistryTransactor{contract: contract}, UserRegistryFilterer: UserRegistryFilterer{contract: contract}}, nil
}

// NewUserRegistryCaller creates a new read-only instance of UserRegistry, bound to a specific deployed contract.
func NewUserRegistryCaller(address common.Address, caller bind.ContractCaller) (*UserRegistryCaller, error) {
	contract, err := bindUserRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserRegistryCaller{contract: contract}, nil
}

// NewUserRegistryTransactor creates a new write-only instance of UserRegistry, bound to a specific deployed contract.
func NewUserRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*UserRegistryTransactor, error) {
	contract, err := bindUserRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserRegistryTransactor{contract: contract}, nil
}

// NewUserRegistryFilterer creates a new log filterer instance of UserRegistry, bound to a specific deployed contract.
func NewUserRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*UserRegistryFilterer, error) {
	contract, err := bindUserRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserRegistryFilterer{contract: contract}, nil
}

// bindUserRegistry binds a generic wrapper to an already deployed contract.
func bindUserRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserRegistry *UserRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserRegistry.Contract.UserRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserRegistry *UserRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserRegistry.Contract.UserRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserRegistry *UserRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserRegistry.Contract.UserRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserRegistry *UserRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserRegistry *UserRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserRegistry *UserRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetAllUsers is a free data retrieval call binding the contract method 0xe2842d79.
//
// Solidity: function getAllUsers() view returns(address[])
func (_UserRegistry *UserRegistryCaller) GetAllUsers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _UserRegistry.contract.Call(opts, &out, "getAllUsers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllUsers is a free data retrieval call binding the contract method 0xe2842d79.
//
// Solidity: function getAllUsers() view returns(address[])
func (_UserRegistry *UserRegistrySession) GetAllUsers() ([]common.Address, error) {
	return _UserRegistry.Contract.GetAllUsers(&_UserRegistry.CallOpts)
}

// GetAllUsers is a free data retrieval call binding the contract method 0xe2842d79.
//
// Solidity: function getAllUsers() view returns(address[])
func (_UserRegistry *UserRegistryCallerSession) GetAllUsers() ([]common.Address, error) {
	return _UserRegistry.Contract.GetAllUsers(&_UserRegistry.CallOpts)
}

// GetKey is a free data retrieval call binding the contract method 0xb6e1a1e2.
//
// Solidity: function getKey(address user, uint8 keypos) view returns((uint8,uint8,bytes))
func (_UserRegistry *UserRegistryCaller) GetKey(opts *bind.CallOpts, user common.Address, keypos uint8) (Pubkey, error) {
	var out []interface{}
	err := _UserRegistry.contract.Call(opts, &out, "getKey", user, keypos)

	if err != nil {
		return *new(Pubkey), err
	}

	out0 := *abi.ConvertType(out[0], new(Pubkey)).(*Pubkey)

	return out0, err

}

// GetKey is a free data retrieval call binding the contract method 0xb6e1a1e2.
//
// Solidity: function getKey(address user, uint8 keypos) view returns((uint8,uint8,bytes))
func (_UserRegistry *UserRegistrySession) GetKey(user common.Address, keypos uint8) (Pubkey, error) {
	return _UserRegistry.Contract.GetKey(&_UserRegistry.CallOpts, user, keypos)
}

// GetKey is a free data retrieval call binding the contract method 0xb6e1a1e2.
//
// Solidity: function getKey(address user, uint8 keypos) view returns((uint8,uint8,bytes))
func (_UserRegistry *UserRegistryCallerSession) GetKey(user common.Address, keypos uint8) (Pubkey, error) {
	return _UserRegistry.Contract.GetKey(&_UserRegistry.CallOpts, user, keypos)
}

// GetKeyLen is a free data retrieval call binding the contract method 0x5c666e54.
//
// Solidity: function getKeyLen(address user) view returns(uint256)
func (_UserRegistry *UserRegistryCaller) GetKeyLen(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UserRegistry.contract.Call(opts, &out, "getKeyLen", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKeyLen is a free data retrieval call binding the contract method 0x5c666e54.
//
// Solidity: function getKeyLen(address user) view returns(uint256)
func (_UserRegistry *UserRegistrySession) GetKeyLen(user common.Address) (*big.Int, error) {
	return _UserRegistry.Contract.GetKeyLen(&_UserRegistry.CallOpts, user)
}

// GetKeyLen is a free data retrieval call binding the contract method 0x5c666e54.
//
// Solidity: function getKeyLen(address user) view returns(uint256)
func (_UserRegistry *UserRegistryCallerSession) GetKeyLen(user common.Address) (*big.Int, error) {
	return _UserRegistry.Contract.GetKeyLen(&_UserRegistry.CallOpts, user)
}

// GetKeys is a free data retrieval call binding the contract method 0x34e80c34.
//
// Solidity: function getKeys(address user) view returns((uint8,uint8,bytes)[])
func (_UserRegistry *UserRegistryCaller) GetKeys(opts *bind.CallOpts, user common.Address) ([]Pubkey, error) {
	var out []interface{}
	err := _UserRegistry.contract.Call(opts, &out, "getKeys", user)

	if err != nil {
		return *new([]Pubkey), err
	}

	out0 := *abi.ConvertType(out[0], new([]Pubkey)).(*[]Pubkey)

	return out0, err

}

// GetKeys is a free data retrieval call binding the contract method 0x34e80c34.
//
// Solidity: function getKeys(address user) view returns((uint8,uint8,bytes)[])
func (_UserRegistry *UserRegistrySession) GetKeys(user common.Address) ([]Pubkey, error) {
	return _UserRegistry.Contract.GetKeys(&_UserRegistry.CallOpts, user)
}

// GetKeys is a free data retrieval call binding the contract method 0x34e80c34.
//
// Solidity: function getKeys(address user) view returns((uint8,uint8,bytes)[])
func (_UserRegistry *UserRegistryCallerSession) GetKeys(user common.Address) ([]Pubkey, error) {
	return _UserRegistry.Contract.GetKeys(&_UserRegistry.CallOpts, user)
}

// GetName is a free data retrieval call binding the contract method 0x5fd4b08a.
//
// Solidity: function getName(address user) view returns(string)
func (_UserRegistry *UserRegistryCaller) GetName(opts *bind.CallOpts, user common.Address) (string, error) {
	var out []interface{}
	err := _UserRegistry.contract.Call(opts, &out, "getName", user)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x5fd4b08a.
//
// Solidity: function getName(address user) view returns(string)
func (_UserRegistry *UserRegistrySession) GetName(user common.Address) (string, error) {
	return _UserRegistry.Contract.GetName(&_UserRegistry.CallOpts, user)
}

// GetName is a free data retrieval call binding the contract method 0x5fd4b08a.
//
// Solidity: function getName(address user) view returns(string)
func (_UserRegistry *UserRegistryCallerSession) GetName(user common.Address) (string, error) {
	return _UserRegistry.Contract.GetName(&_UserRegistry.CallOpts, user)
}

// GetUser is a free data retrieval call binding the contract method 0x31feb671.
//
// Solidity: function getUser(string name) view returns(address)
func (_UserRegistry *UserRegistryCaller) GetUser(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _UserRegistry.contract.Call(opts, &out, "getUser", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUser is a free data retrieval call binding the contract method 0x31feb671.
//
// Solidity: function getUser(string name) view returns(address)
func (_UserRegistry *UserRegistrySession) GetUser(name string) (common.Address, error) {
	return _UserRegistry.Contract.GetUser(&_UserRegistry.CallOpts, name)
}

// GetUser is a free data retrieval call binding the contract method 0x31feb671.
//
// Solidity: function getUser(string name) view returns(address)
func (_UserRegistry *UserRegistryCallerSession) GetUser(name string) (common.Address, error) {
	return _UserRegistry.Contract.GetUser(&_UserRegistry.CallOpts, name)
}

// AddPubkey is a paid mutator transaction binding the contract method 0xbf5fc61d.
//
// Solidity: function addPubkey(address user, uint8 keytype, uint8 status, bytes key) returns(string)
func (_UserRegistry *UserRegistryTransactor) AddPubkey(opts *bind.TransactOpts, user common.Address, keytype uint8, status uint8, key []byte) (*types.Transaction, error) {
	return _UserRegistry.contract.Transact(opts, "addPubkey", user, keytype, status, key)
}

// AddPubkey is a paid mutator transaction binding the contract method 0xbf5fc61d.
//
// Solidity: function addPubkey(address user, uint8 keytype, uint8 status, bytes key) returns(string)
func (_UserRegistry *UserRegistrySession) AddPubkey(user common.Address, keytype uint8, status uint8, key []byte) (*types.Transaction, error) {
	return _UserRegistry.Contract.AddPubkey(&_UserRegistry.TransactOpts, user, keytype, status, key)
}

// AddPubkey is a paid mutator transaction binding the contract method 0xbf5fc61d.
//
// Solidity: function addPubkey(address user, uint8 keytype, uint8 status, bytes key) returns(string)
func (_UserRegistry *UserRegistryTransactorSession) AddPubkey(user common.Address, keytype uint8, status uint8, key []byte) (*types.Transaction, error) {
	return _UserRegistry.Contract.AddPubkey(&_UserRegistry.TransactOpts, user, keytype, status, key)
}

// NewUser is a paid mutator transaction binding the contract method 0xaba0caae.
//
// Solidity: function newUser(address user, string name, uint8 keytype, uint8 status, bytes key) returns(string)
func (_UserRegistry *UserRegistryTransactor) NewUser(opts *bind.TransactOpts, user common.Address, name string, keytype uint8, status uint8, key []byte) (*types.Transaction, error) {
	return _UserRegistry.contract.Transact(opts, "newUser", user, name, keytype, status, key)
}

// NewUser is a paid mutator transaction binding the contract method 0xaba0caae.
//
// Solidity: function newUser(address user, string name, uint8 keytype, uint8 status, bytes key) returns(string)
func (_UserRegistry *UserRegistrySession) NewUser(user common.Address, name string, keytype uint8, status uint8, key []byte) (*types.Transaction, error) {
	return _UserRegistry.Contract.NewUser(&_UserRegistry.TransactOpts, user, name, keytype, status, key)
}

// NewUser is a paid mutator transaction binding the contract method 0xaba0caae.
//
// Solidity: function newUser(address user, string name, uint8 keytype, uint8 status, bytes key) returns(string)
func (_UserRegistry *UserRegistryTransactorSession) NewUser(user common.Address, name string, keytype uint8, status uint8, key []byte) (*types.Transaction, error) {
	return _UserRegistry.Contract.NewUser(&_UserRegistry.TransactOpts, user, name, keytype, status, key)
}

// UpdateKeyStatus is a paid mutator transaction binding the contract method 0x1dd98477.
//
// Solidity: function updateKeyStatus(address user, uint8 keypos, uint8 status) returns()
func (_UserRegistry *UserRegistryTransactor) UpdateKeyStatus(opts *bind.TransactOpts, user common.Address, keypos uint8, status uint8) (*types.Transaction, error) {
	return _UserRegistry.contract.Transact(opts, "updateKeyStatus", user, keypos, status)
}

// UpdateKeyStatus is a paid mutator transaction binding the contract method 0x1dd98477.
//
// Solidity: function updateKeyStatus(address user, uint8 keypos, uint8 status) returns()
func (_UserRegistry *UserRegistrySession) UpdateKeyStatus(user common.Address, keypos uint8, status uint8) (*types.Transaction, error) {
	return _UserRegistry.Contract.UpdateKeyStatus(&_UserRegistry.TransactOpts, user, keypos, status)
}

// UpdateKeyStatus is a paid mutator transaction binding the contract method 0x1dd98477.
//
// Solidity: function updateKeyStatus(address user, uint8 keypos, uint8 status) returns()
func (_UserRegistry *UserRegistryTransactorSession) UpdateKeyStatus(user common.Address, keypos uint8, status uint8) (*types.Transaction, error) {
	return _UserRegistry.Contract.UpdateKeyStatus(&_UserRegistry.TransactOpts, user, keypos, status)
}
