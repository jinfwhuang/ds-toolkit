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
	Keytype   uint8
	Keystatus uint8
	Key       []byte
}

// UserRegistryMetaData contains all meta data concerning the UserRegistry contract.
var UserRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"keytype\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"keystatus\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"addPubkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"computeAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"ethSignedHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllUsers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"keypos\",\"type\":\"uint8\"}],\"name\":\"getImpliedAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"keypos\",\"type\":\"uint8\"}],\"name\":\"getKey\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"keytype\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"keystatus\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"internalType\":\"structPubkey\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getKeys\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"keytype\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"keystatus\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"internalType\":\"structPubkey[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getLenKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getUser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserNonce\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"keytype\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"keykeystatus\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"newUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rndHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"keypos\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"keystatus\",\"type\":\"uint8\"}],\"name\":\"updateKeykeystatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verifyUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// ComputeAddr is a free data retrieval call binding the contract method 0x28da9254.
//
// Solidity: function computeAddr(bytes pubkey) pure returns(address)
func (_UserRegistry *UserRegistryCaller) ComputeAddr(opts *bind.CallOpts, pubkey []byte) (common.Address, error) {
	var out []interface{}
	err := _UserRegistry.contract.Call(opts, &out, "computeAddr", pubkey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeAddr is a free data retrieval call binding the contract method 0x28da9254.
//
// Solidity: function computeAddr(bytes pubkey) pure returns(address)
func (_UserRegistry *UserRegistrySession) ComputeAddr(pubkey []byte) (common.Address, error) {
	return _UserRegistry.Contract.ComputeAddr(&_UserRegistry.CallOpts, pubkey)
}

// ComputeAddr is a free data retrieval call binding the contract method 0x28da9254.
//
// Solidity: function computeAddr(bytes pubkey) pure returns(address)
func (_UserRegistry *UserRegistryCallerSession) ComputeAddr(pubkey []byte) (common.Address, error) {
	return _UserRegistry.Contract.ComputeAddr(&_UserRegistry.CallOpts, pubkey)
}

// EthSignedHash is a free data retrieval call binding the contract method 0xf1520932.
//
// Solidity: function ethSignedHash(bytes32 messageHash) pure returns(bytes32)
func (_UserRegistry *UserRegistryCaller) EthSignedHash(opts *bind.CallOpts, messageHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _UserRegistry.contract.Call(opts, &out, "ethSignedHash", messageHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EthSignedHash is a free data retrieval call binding the contract method 0xf1520932.
//
// Solidity: function ethSignedHash(bytes32 messageHash) pure returns(bytes32)
func (_UserRegistry *UserRegistrySession) EthSignedHash(messageHash [32]byte) ([32]byte, error) {
	return _UserRegistry.Contract.EthSignedHash(&_UserRegistry.CallOpts, messageHash)
}

// EthSignedHash is a free data retrieval call binding the contract method 0xf1520932.
//
// Solidity: function ethSignedHash(bytes32 messageHash) pure returns(bytes32)
func (_UserRegistry *UserRegistryCallerSession) EthSignedHash(messageHash [32]byte) ([32]byte, error) {
	return _UserRegistry.Contract.EthSignedHash(&_UserRegistry.CallOpts, messageHash)
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

// GetImpliedAddr is a free data retrieval call binding the contract method 0x2c0b062c.
//
// Solidity: function getImpliedAddr(address user, uint8 keypos) view returns(address)
func (_UserRegistry *UserRegistryCaller) GetImpliedAddr(opts *bind.CallOpts, user common.Address, keypos uint8) (common.Address, error) {
	var out []interface{}
	err := _UserRegistry.contract.Call(opts, &out, "getImpliedAddr", user, keypos)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImpliedAddr is a free data retrieval call binding the contract method 0x2c0b062c.
//
// Solidity: function getImpliedAddr(address user, uint8 keypos) view returns(address)
func (_UserRegistry *UserRegistrySession) GetImpliedAddr(user common.Address, keypos uint8) (common.Address, error) {
	return _UserRegistry.Contract.GetImpliedAddr(&_UserRegistry.CallOpts, user, keypos)
}

// GetImpliedAddr is a free data retrieval call binding the contract method 0x2c0b062c.
//
// Solidity: function getImpliedAddr(address user, uint8 keypos) view returns(address)
func (_UserRegistry *UserRegistryCallerSession) GetImpliedAddr(user common.Address, keypos uint8) (common.Address, error) {
	return _UserRegistry.Contract.GetImpliedAddr(&_UserRegistry.CallOpts, user, keypos)
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

// GetLenKeys is a free data retrieval call binding the contract method 0x3f11639f.
//
// Solidity: function getLenKeys(address user) view returns(uint256)
func (_UserRegistry *UserRegistryCaller) GetLenKeys(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UserRegistry.contract.Call(opts, &out, "getLenKeys", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLenKeys is a free data retrieval call binding the contract method 0x3f11639f.
//
// Solidity: function getLenKeys(address user) view returns(uint256)
func (_UserRegistry *UserRegistrySession) GetLenKeys(user common.Address) (*big.Int, error) {
	return _UserRegistry.Contract.GetLenKeys(&_UserRegistry.CallOpts, user)
}

// GetLenKeys is a free data retrieval call binding the contract method 0x3f11639f.
//
// Solidity: function getLenKeys(address user) view returns(uint256)
func (_UserRegistry *UserRegistryCallerSession) GetLenKeys(user common.Address) (*big.Int, error) {
	return _UserRegistry.Contract.GetLenKeys(&_UserRegistry.CallOpts, user)
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

// GetUserNonce is a free data retrieval call binding the contract method 0x6834e3a8.
//
// Solidity: function getUserNonce(address user) view returns(uint16)
func (_UserRegistry *UserRegistryCaller) GetUserNonce(opts *bind.CallOpts, user common.Address) (uint16, error) {
	var out []interface{}
	err := _UserRegistry.contract.Call(opts, &out, "getUserNonce", user)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetUserNonce is a free data retrieval call binding the contract method 0x6834e3a8.
//
// Solidity: function getUserNonce(address user) view returns(uint16)
func (_UserRegistry *UserRegistrySession) GetUserNonce(user common.Address) (uint16, error) {
	return _UserRegistry.Contract.GetUserNonce(&_UserRegistry.CallOpts, user)
}

// GetUserNonce is a free data retrieval call binding the contract method 0x6834e3a8.
//
// Solidity: function getUserNonce(address user) view returns(uint16)
func (_UserRegistry *UserRegistryCallerSession) GetUserNonce(user common.Address) (uint16, error) {
	return _UserRegistry.Contract.GetUserNonce(&_UserRegistry.CallOpts, user)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) pure returns(address)
func (_UserRegistry *UserRegistryCaller) Recover(opts *bind.CallOpts, hash [32]byte, signature []byte) (common.Address, error) {
	var out []interface{}
	err := _UserRegistry.contract.Call(opts, &out, "recover", hash, signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) pure returns(address)
func (_UserRegistry *UserRegistrySession) Recover(hash [32]byte, signature []byte) (common.Address, error) {
	return _UserRegistry.Contract.Recover(&_UserRegistry.CallOpts, hash, signature)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(bytes32 hash, bytes signature) pure returns(address)
func (_UserRegistry *UserRegistryCallerSession) Recover(hash [32]byte, signature []byte) (common.Address, error) {
	return _UserRegistry.Contract.Recover(&_UserRegistry.CallOpts, hash, signature)
}

// RndHash is a free data retrieval call binding the contract method 0x2abd5908.
//
// Solidity: function rndHash() view returns(bytes32)
func (_UserRegistry *UserRegistryCaller) RndHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UserRegistry.contract.Call(opts, &out, "rndHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RndHash is a free data retrieval call binding the contract method 0x2abd5908.
//
// Solidity: function rndHash() view returns(bytes32)
func (_UserRegistry *UserRegistrySession) RndHash() ([32]byte, error) {
	return _UserRegistry.Contract.RndHash(&_UserRegistry.CallOpts)
}

// RndHash is a free data retrieval call binding the contract method 0x2abd5908.
//
// Solidity: function rndHash() view returns(bytes32)
func (_UserRegistry *UserRegistryCallerSession) RndHash() ([32]byte, error) {
	return _UserRegistry.Contract.RndHash(&_UserRegistry.CallOpts)
}

// VerifyUser is a free data retrieval call binding the contract method 0x1b5aadf7.
//
// Solidity: function verifyUser(address user, bytes32 msgHash, bytes signature) view returns(bool isValid)
func (_UserRegistry *UserRegistryCaller) VerifyUser(opts *bind.CallOpts, user common.Address, msgHash [32]byte, signature []byte) (bool, error) {
	var out []interface{}
	err := _UserRegistry.contract.Call(opts, &out, "verifyUser", user, msgHash, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyUser is a free data retrieval call binding the contract method 0x1b5aadf7.
//
// Solidity: function verifyUser(address user, bytes32 msgHash, bytes signature) view returns(bool isValid)
func (_UserRegistry *UserRegistrySession) VerifyUser(user common.Address, msgHash [32]byte, signature []byte) (bool, error) {
	return _UserRegistry.Contract.VerifyUser(&_UserRegistry.CallOpts, user, msgHash, signature)
}

// VerifyUser is a free data retrieval call binding the contract method 0x1b5aadf7.
//
// Solidity: function verifyUser(address user, bytes32 msgHash, bytes signature) view returns(bool isValid)
func (_UserRegistry *UserRegistryCallerSession) VerifyUser(user common.Address, msgHash [32]byte, signature []byte) (bool, error) {
	return _UserRegistry.Contract.VerifyUser(&_UserRegistry.CallOpts, user, msgHash, signature)
}

// AddPubkey is a paid mutator transaction binding the contract method 0xdc6f6edb.
//
// Solidity: function addPubkey(address user, uint8 keytype, uint8 keystatus, bytes pubkey, bytes sig) returns()
func (_UserRegistry *UserRegistryTransactor) AddPubkey(opts *bind.TransactOpts, user common.Address, keytype uint8, keystatus uint8, pubkey []byte, sig []byte) (*types.Transaction, error) {
	return _UserRegistry.contract.Transact(opts, "addPubkey", user, keytype, keystatus, pubkey, sig)
}

// AddPubkey is a paid mutator transaction binding the contract method 0xdc6f6edb.
//
// Solidity: function addPubkey(address user, uint8 keytype, uint8 keystatus, bytes pubkey, bytes sig) returns()
func (_UserRegistry *UserRegistrySession) AddPubkey(user common.Address, keytype uint8, keystatus uint8, pubkey []byte, sig []byte) (*types.Transaction, error) {
	return _UserRegistry.Contract.AddPubkey(&_UserRegistry.TransactOpts, user, keytype, keystatus, pubkey, sig)
}

// AddPubkey is a paid mutator transaction binding the contract method 0xdc6f6edb.
//
// Solidity: function addPubkey(address user, uint8 keytype, uint8 keystatus, bytes pubkey, bytes sig) returns()
func (_UserRegistry *UserRegistryTransactorSession) AddPubkey(user common.Address, keytype uint8, keystatus uint8, pubkey []byte, sig []byte) (*types.Transaction, error) {
	return _UserRegistry.Contract.AddPubkey(&_UserRegistry.TransactOpts, user, keytype, keystatus, pubkey, sig)
}

// NewUser is a paid mutator transaction binding the contract method 0xaba0caae.
//
// Solidity: function newUser(address user, string name, uint8 keytype, uint8 keykeystatus, bytes pubkey) returns()
func (_UserRegistry *UserRegistryTransactor) NewUser(opts *bind.TransactOpts, user common.Address, name string, keytype uint8, keykeystatus uint8, pubkey []byte) (*types.Transaction, error) {
	return _UserRegistry.contract.Transact(opts, "newUser", user, name, keytype, keykeystatus, pubkey)
}

// NewUser is a paid mutator transaction binding the contract method 0xaba0caae.
//
// Solidity: function newUser(address user, string name, uint8 keytype, uint8 keykeystatus, bytes pubkey) returns()
func (_UserRegistry *UserRegistrySession) NewUser(user common.Address, name string, keytype uint8, keykeystatus uint8, pubkey []byte) (*types.Transaction, error) {
	return _UserRegistry.Contract.NewUser(&_UserRegistry.TransactOpts, user, name, keytype, keykeystatus, pubkey)
}

// NewUser is a paid mutator transaction binding the contract method 0xaba0caae.
//
// Solidity: function newUser(address user, string name, uint8 keytype, uint8 keykeystatus, bytes pubkey) returns()
func (_UserRegistry *UserRegistryTransactorSession) NewUser(user common.Address, name string, keytype uint8, keykeystatus uint8, pubkey []byte) (*types.Transaction, error) {
	return _UserRegistry.Contract.NewUser(&_UserRegistry.TransactOpts, user, name, keytype, keykeystatus, pubkey)
}

// UpdateKeykeystatus is a paid mutator transaction binding the contract method 0xf37f48a4.
//
// Solidity: function updateKeykeystatus(address user, uint8 keypos, uint8 keystatus) returns()
func (_UserRegistry *UserRegistryTransactor) UpdateKeykeystatus(opts *bind.TransactOpts, user common.Address, keypos uint8, keystatus uint8) (*types.Transaction, error) {
	return _UserRegistry.contract.Transact(opts, "updateKeykeystatus", user, keypos, keystatus)
}

// UpdateKeykeystatus is a paid mutator transaction binding the contract method 0xf37f48a4.
//
// Solidity: function updateKeykeystatus(address user, uint8 keypos, uint8 keystatus) returns()
func (_UserRegistry *UserRegistrySession) UpdateKeykeystatus(user common.Address, keypos uint8, keystatus uint8) (*types.Transaction, error) {
	return _UserRegistry.Contract.UpdateKeykeystatus(&_UserRegistry.TransactOpts, user, keypos, keystatus)
}

// UpdateKeykeystatus is a paid mutator transaction binding the contract method 0xf37f48a4.
//
// Solidity: function updateKeykeystatus(address user, uint8 keypos, uint8 keystatus) returns()
func (_UserRegistry *UserRegistryTransactorSession) UpdateKeykeystatus(user common.Address, keypos uint8, keystatus uint8) (*types.Transaction, error) {
	return _UserRegistry.Contract.UpdateKeykeystatus(&_UserRegistry.TransactOpts, user, keypos, keystatus)
}
