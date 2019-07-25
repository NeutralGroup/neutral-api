// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package neutral_token

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// NeutralTokenABI is the input ABI used to generate the binding from.
const NeutralTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x06fdde03\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x095ea7b3\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x18160ddd\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x23b872dd\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x313ce567\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x39509351\"},{\"constant\":false,\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x40a141ff\"},{\"constant\":false,\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x4d238c8e\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x70a08231\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x715018a6\"},{\"constant\":true,\"inputs\":[],\"name\":\"isVaultManaged\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x829dd895\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x8da5cb5b\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x8f32d59b\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x95d89b41\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xa457c2d7\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nonce\",\"type\":\"uint256\"},{\"name\":\"_feeDestination\",\"type\":\"address\"},{\"name\":\"_expirationDate\",\"type\":\"uint256\"},{\"name\":\"_expirationBlock\",\"type\":\"uint256\"},{\"name\":\"_source\",\"type\":\"address\"},{\"name\":\"_destination\",\"type\":\"address\"},{\"name\":\"_instrument\",\"type\":\"address\"},{\"name\":\"_instrumentQuantity\",\"type\":\"uint256\"},{\"name\":\"_instrumentOperation\",\"type\":\"uint8\"},{\"name\":\"_neutralQuantity\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_neutralBoundary\",\"type\":\"uint256\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"settle\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xa86f8e66\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xa9059cbb\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xdd62ed3e\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xf2fde38b\"},{\"constant\":false,\"inputs\":[{\"name\":\"_NeutralVault\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xf6d2ee86\"},{\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"signature\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_feeDestination\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_source\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_destination\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_instrument\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_instrumentQuantity\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_instrumentOperation\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"_neutralQuantity\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_fee\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_neutralBoundary\",\"type\":\"uint256\"}],\"name\":\"Settlement\",\"type\":\"event\",\"signature\":\"0x146934e5f2b06d949bf7603893fa5e00cbc5c1dd49bf19d0b8b33c8ac9ab8d99\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\",\"signature\":\"0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\",\"signature\":\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\",\"signature\":\"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925\"}]"

// NeutralToken is an auto generated Go binding around an Ethereum contract.
type NeutralToken struct {
	NeutralTokenCaller     // Read-only binding to the contract
	NeutralTokenTransactor // Write-only binding to the contract
	NeutralTokenFilterer   // Log filterer for contract events
}

// NeutralTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type NeutralTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NeutralTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NeutralTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NeutralTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NeutralTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NeutralTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NeutralTokenSession struct {
	Contract     *NeutralToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NeutralTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NeutralTokenCallerSession struct {
	Contract *NeutralTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// NeutralTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NeutralTokenTransactorSession struct {
	Contract     *NeutralTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NeutralTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type NeutralTokenRaw struct {
	Contract *NeutralToken // Generic contract binding to access the raw methods on
}

// NeutralTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NeutralTokenCallerRaw struct {
	Contract *NeutralTokenCaller // Generic read-only contract binding to access the raw methods on
}

// NeutralTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NeutralTokenTransactorRaw struct {
	Contract *NeutralTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNeutralToken creates a new instance of NeutralToken, bound to a specific deployed contract.
func NewNeutralToken(address common.Address, backend bind.ContractBackend) (*NeutralToken, error) {
	contract, err := bindNeutralToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NeutralToken{NeutralTokenCaller: NeutralTokenCaller{contract: contract}, NeutralTokenTransactor: NeutralTokenTransactor{contract: contract}, NeutralTokenFilterer: NeutralTokenFilterer{contract: contract}}, nil
}

// NewNeutralTokenCaller creates a new read-only instance of NeutralToken, bound to a specific deployed contract.
func NewNeutralTokenCaller(address common.Address, caller bind.ContractCaller) (*NeutralTokenCaller, error) {
	contract, err := bindNeutralToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NeutralTokenCaller{contract: contract}, nil
}

// NewNeutralTokenTransactor creates a new write-only instance of NeutralToken, bound to a specific deployed contract.
func NewNeutralTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*NeutralTokenTransactor, error) {
	contract, err := bindNeutralToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NeutralTokenTransactor{contract: contract}, nil
}

// NewNeutralTokenFilterer creates a new log filterer instance of NeutralToken, bound to a specific deployed contract.
func NewNeutralTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*NeutralTokenFilterer, error) {
	contract, err := bindNeutralToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NeutralTokenFilterer{contract: contract}, nil
}

// bindNeutralToken binds a generic wrapper to an already deployed contract.
func bindNeutralToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NeutralTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NeutralToken *NeutralTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NeutralToken.Contract.NeutralTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NeutralToken *NeutralTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeutralToken.Contract.NeutralTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NeutralToken *NeutralTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NeutralToken.Contract.NeutralTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NeutralToken *NeutralTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NeutralToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NeutralToken *NeutralTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeutralToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NeutralToken *NeutralTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NeutralToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_NeutralToken *NeutralTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NeutralToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_NeutralToken *NeutralTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NeutralToken.Contract.Allowance(&_NeutralToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_NeutralToken *NeutralTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NeutralToken.Contract.Allowance(&_NeutralToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_NeutralToken *NeutralTokenCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NeutralToken.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_NeutralToken *NeutralTokenSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _NeutralToken.Contract.BalanceOf(&_NeutralToken.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) constant returns(uint256)
func (_NeutralToken *NeutralTokenCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _NeutralToken.Contract.BalanceOf(&_NeutralToken.CallOpts, who)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_NeutralToken *NeutralTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _NeutralToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_NeutralToken *NeutralTokenSession) Decimals() (uint8, error) {
	return _NeutralToken.Contract.Decimals(&_NeutralToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_NeutralToken *NeutralTokenCallerSession) Decimals() (uint8, error) {
	return _NeutralToken.Contract.Decimals(&_NeutralToken.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_NeutralToken *NeutralTokenCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NeutralToken.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_NeutralToken *NeutralTokenSession) IsOwner() (bool, error) {
	return _NeutralToken.Contract.IsOwner(&_NeutralToken.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_NeutralToken *NeutralTokenCallerSession) IsOwner() (bool, error) {
	return _NeutralToken.Contract.IsOwner(&_NeutralToken.CallOpts)
}

// IsVaultManaged is a free data retrieval call binding the contract method 0x829dd895.
//
// Solidity: function isVaultManaged() constant returns(uint8)
func (_NeutralToken *NeutralTokenCaller) IsVaultManaged(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _NeutralToken.contract.Call(opts, out, "isVaultManaged")
	return *ret0, err
}

// IsVaultManaged is a free data retrieval call binding the contract method 0x829dd895.
//
// Solidity: function isVaultManaged() constant returns(uint8)
func (_NeutralToken *NeutralTokenSession) IsVaultManaged() (uint8, error) {
	return _NeutralToken.Contract.IsVaultManaged(&_NeutralToken.CallOpts)
}

// IsVaultManaged is a free data retrieval call binding the contract method 0x829dd895.
//
// Solidity: function isVaultManaged() constant returns(uint8)
func (_NeutralToken *NeutralTokenCallerSession) IsVaultManaged() (uint8, error) {
	return _NeutralToken.Contract.IsVaultManaged(&_NeutralToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_NeutralToken *NeutralTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NeutralToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_NeutralToken *NeutralTokenSession) Name() (string, error) {
	return _NeutralToken.Contract.Name(&_NeutralToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_NeutralToken *NeutralTokenCallerSession) Name() (string, error) {
	return _NeutralToken.Contract.Name(&_NeutralToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NeutralToken *NeutralTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NeutralToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NeutralToken *NeutralTokenSession) Owner() (common.Address, error) {
	return _NeutralToken.Contract.Owner(&_NeutralToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NeutralToken *NeutralTokenCallerSession) Owner() (common.Address, error) {
	return _NeutralToken.Contract.Owner(&_NeutralToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_NeutralToken *NeutralTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NeutralToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_NeutralToken *NeutralTokenSession) Symbol() (string, error) {
	return _NeutralToken.Contract.Symbol(&_NeutralToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_NeutralToken *NeutralTokenCallerSession) Symbol() (string, error) {
	return _NeutralToken.Contract.Symbol(&_NeutralToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_NeutralToken *NeutralTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NeutralToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_NeutralToken *NeutralTokenSession) TotalSupply() (*big.Int, error) {
	return _NeutralToken.Contract.TotalSupply(&_NeutralToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_NeutralToken *NeutralTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _NeutralToken.Contract.TotalSupply(&_NeutralToken.CallOpts)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validator) returns(bool)
func (_NeutralToken *NeutralTokenTransactor) AddValidator(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _NeutralToken.contract.Transact(opts, "addValidator", _validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validator) returns(bool)
func (_NeutralToken *NeutralTokenSession) AddValidator(_validator common.Address) (*types.Transaction, error) {
	return _NeutralToken.Contract.AddValidator(&_NeutralToken.TransactOpts, _validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validator) returns(bool)
func (_NeutralToken *NeutralTokenTransactorSession) AddValidator(_validator common.Address) (*types.Transaction, error) {
	return _NeutralToken.Contract.AddValidator(&_NeutralToken.TransactOpts, _validator)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_NeutralToken *NeutralTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _NeutralToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_NeutralToken *NeutralTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _NeutralToken.Contract.Approve(&_NeutralToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_NeutralToken *NeutralTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _NeutralToken.Contract.Approve(&_NeutralToken.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NeutralToken *NeutralTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NeutralToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NeutralToken *NeutralTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NeutralToken.Contract.DecreaseAllowance(&_NeutralToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NeutralToken *NeutralTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NeutralToken.Contract.DecreaseAllowance(&_NeutralToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NeutralToken *NeutralTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NeutralToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NeutralToken *NeutralTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NeutralToken.Contract.IncreaseAllowance(&_NeutralToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NeutralToken *NeutralTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NeutralToken.Contract.IncreaseAllowance(&_NeutralToken.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xf6d2ee86.
//
// Solidity: function initialize(address _NeutralVault, string _name, string _symbol, uint8 _decimals) returns()
func (_NeutralToken *NeutralTokenTransactor) Initialize(opts *bind.TransactOpts, _NeutralVault common.Address, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _NeutralToken.contract.Transact(opts, "initialize", _NeutralVault, _name, _symbol, _decimals)
}

// Initialize is a paid mutator transaction binding the contract method 0xf6d2ee86.
//
// Solidity: function initialize(address _NeutralVault, string _name, string _symbol, uint8 _decimals) returns()
func (_NeutralToken *NeutralTokenSession) Initialize(_NeutralVault common.Address, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _NeutralToken.Contract.Initialize(&_NeutralToken.TransactOpts, _NeutralVault, _name, _symbol, _decimals)
}

// Initialize is a paid mutator transaction binding the contract method 0xf6d2ee86.
//
// Solidity: function initialize(address _NeutralVault, string _name, string _symbol, uint8 _decimals) returns()
func (_NeutralToken *NeutralTokenTransactorSession) Initialize(_NeutralVault common.Address, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _NeutralToken.Contract.Initialize(&_NeutralToken.TransactOpts, _NeutralVault, _name, _symbol, _decimals)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validator) returns(bool)
func (_NeutralToken *NeutralTokenTransactor) RemoveValidator(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _NeutralToken.contract.Transact(opts, "removeValidator", _validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validator) returns(bool)
func (_NeutralToken *NeutralTokenSession) RemoveValidator(_validator common.Address) (*types.Transaction, error) {
	return _NeutralToken.Contract.RemoveValidator(&_NeutralToken.TransactOpts, _validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validator) returns(bool)
func (_NeutralToken *NeutralTokenTransactorSession) RemoveValidator(_validator common.Address) (*types.Transaction, error) {
	return _NeutralToken.Contract.RemoveValidator(&_NeutralToken.TransactOpts, _validator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NeutralToken *NeutralTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeutralToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NeutralToken *NeutralTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _NeutralToken.Contract.RenounceOwnership(&_NeutralToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NeutralToken *NeutralTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NeutralToken.Contract.RenounceOwnership(&_NeutralToken.TransactOpts)
}

// Settle is a paid mutator transaction binding the contract method 0xa86f8e66.
//
// Solidity: function settle(uint256 _nonce, address _feeDestination, uint256 _expirationDate, uint256 _expirationBlock, address _source, address _destination, address _instrument, uint256 _instrumentQuantity, uint8 _instrumentOperation, uint256 _neutralQuantity, uint256 _fee, uint256 _neutralBoundary, bytes _signature) returns(bool)
func (_NeutralToken *NeutralTokenTransactor) Settle(opts *bind.TransactOpts, _nonce *big.Int, _feeDestination common.Address, _expirationDate *big.Int, _expirationBlock *big.Int, _source common.Address, _destination common.Address, _instrument common.Address, _instrumentQuantity *big.Int, _instrumentOperation uint8, _neutralQuantity *big.Int, _fee *big.Int, _neutralBoundary *big.Int, _signature []byte) (*types.Transaction, error) {
	return _NeutralToken.contract.Transact(opts, "settle", _nonce, _feeDestination, _expirationDate, _expirationBlock, _source, _destination, _instrument, _instrumentQuantity, _instrumentOperation, _neutralQuantity, _fee, _neutralBoundary, _signature)
}

// Settle is a paid mutator transaction binding the contract method 0xa86f8e66.
//
// Solidity: function settle(uint256 _nonce, address _feeDestination, uint256 _expirationDate, uint256 _expirationBlock, address _source, address _destination, address _instrument, uint256 _instrumentQuantity, uint8 _instrumentOperation, uint256 _neutralQuantity, uint256 _fee, uint256 _neutralBoundary, bytes _signature) returns(bool)
func (_NeutralToken *NeutralTokenSession) Settle(_nonce *big.Int, _feeDestination common.Address, _expirationDate *big.Int, _expirationBlock *big.Int, _source common.Address, _destination common.Address, _instrument common.Address, _instrumentQuantity *big.Int, _instrumentOperation uint8, _neutralQuantity *big.Int, _fee *big.Int, _neutralBoundary *big.Int, _signature []byte) (*types.Transaction, error) {
	return _NeutralToken.Contract.Settle(&_NeutralToken.TransactOpts, _nonce, _feeDestination, _expirationDate, _expirationBlock, _source, _destination, _instrument, _instrumentQuantity, _instrumentOperation, _neutralQuantity, _fee, _neutralBoundary, _signature)
}

// Settle is a paid mutator transaction binding the contract method 0xa86f8e66.
//
// Solidity: function settle(uint256 _nonce, address _feeDestination, uint256 _expirationDate, uint256 _expirationBlock, address _source, address _destination, address _instrument, uint256 _instrumentQuantity, uint8 _instrumentOperation, uint256 _neutralQuantity, uint256 _fee, uint256 _neutralBoundary, bytes _signature) returns(bool)
func (_NeutralToken *NeutralTokenTransactorSession) Settle(_nonce *big.Int, _feeDestination common.Address, _expirationDate *big.Int, _expirationBlock *big.Int, _source common.Address, _destination common.Address, _instrument common.Address, _instrumentQuantity *big.Int, _instrumentOperation uint8, _neutralQuantity *big.Int, _fee *big.Int, _neutralBoundary *big.Int, _signature []byte) (*types.Transaction, error) {
	return _NeutralToken.Contract.Settle(&_NeutralToken.TransactOpts, _nonce, _feeDestination, _expirationDate, _expirationBlock, _source, _destination, _instrument, _instrumentQuantity, _instrumentOperation, _neutralQuantity, _fee, _neutralBoundary, _signature)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_NeutralToken *NeutralTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NeutralToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_NeutralToken *NeutralTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NeutralToken.Contract.Transfer(&_NeutralToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_NeutralToken *NeutralTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NeutralToken.Contract.Transfer(&_NeutralToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_NeutralToken *NeutralTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NeutralToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_NeutralToken *NeutralTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NeutralToken.Contract.TransferFrom(&_NeutralToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_NeutralToken *NeutralTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _NeutralToken.Contract.TransferFrom(&_NeutralToken.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NeutralToken *NeutralTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NeutralToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NeutralToken *NeutralTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NeutralToken.Contract.TransferOwnership(&_NeutralToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NeutralToken *NeutralTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NeutralToken.Contract.TransferOwnership(&_NeutralToken.TransactOpts, newOwner)
}

// NeutralTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NeutralToken contract.
type NeutralTokenApprovalIterator struct {
	Event *NeutralTokenApproval // Event containing the contract specifics and raw log

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
func (it *NeutralTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeutralTokenApproval)
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
		it.Event = new(NeutralTokenApproval)
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
func (it *NeutralTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeutralTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeutralTokenApproval represents a Approval event raised by the NeutralToken contract.
type NeutralTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NeutralToken *NeutralTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NeutralTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NeutralToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NeutralTokenApprovalIterator{contract: _NeutralToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NeutralToken *NeutralTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NeutralTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NeutralToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeutralTokenApproval)
				if err := _NeutralToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// NeutralTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NeutralToken contract.
type NeutralTokenOwnershipTransferredIterator struct {
	Event *NeutralTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NeutralTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeutralTokenOwnershipTransferred)
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
		it.Event = new(NeutralTokenOwnershipTransferred)
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
func (it *NeutralTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeutralTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeutralTokenOwnershipTransferred represents a OwnershipTransferred event raised by the NeutralToken contract.
type NeutralTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NeutralToken *NeutralTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NeutralTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NeutralToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NeutralTokenOwnershipTransferredIterator{contract: _NeutralToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NeutralToken *NeutralTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NeutralTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NeutralToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeutralTokenOwnershipTransferred)
				if err := _NeutralToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// NeutralTokenSettlementIterator is returned from FilterSettlement and is used to iterate over the raw logs and unpacked data for Settlement events raised by the NeutralToken contract.
type NeutralTokenSettlementIterator struct {
	Event *NeutralTokenSettlement // Event containing the contract specifics and raw log

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
func (it *NeutralTokenSettlementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeutralTokenSettlement)
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
		it.Event = new(NeutralTokenSettlement)
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
func (it *NeutralTokenSettlementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeutralTokenSettlementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeutralTokenSettlement represents a Settlement event raised by the NeutralToken contract.
type NeutralTokenSettlement struct {
	FeeDestination      common.Address
	Source              common.Address
	Destination         common.Address
	Instrument          common.Address
	InstrumentQuantity  *big.Int
	InstrumentOperation uint8
	NeutralQuantity     *big.Int
	Fee                 *big.Int
	NeutralBoundary     *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSettlement is a free log retrieval operation binding the contract event 0x146934e5f2b06d949bf7603893fa5e00cbc5c1dd49bf19d0b8b33c8ac9ab8d99.
//
// Solidity: event Settlement(address _feeDestination, address indexed _source, address indexed _destination, address indexed _instrument, uint256 _instrumentQuantity, uint8 _instrumentOperation, uint256 _neutralQuantity, uint256 _fee, uint256 _neutralBoundary)
func (_NeutralToken *NeutralTokenFilterer) FilterSettlement(opts *bind.FilterOpts, _source []common.Address, _destination []common.Address, _instrument []common.Address) (*NeutralTokenSettlementIterator, error) {

	var _sourceRule []interface{}
	for _, _sourceItem := range _source {
		_sourceRule = append(_sourceRule, _sourceItem)
	}
	var _destinationRule []interface{}
	for _, _destinationItem := range _destination {
		_destinationRule = append(_destinationRule, _destinationItem)
	}
	var _instrumentRule []interface{}
	for _, _instrumentItem := range _instrument {
		_instrumentRule = append(_instrumentRule, _instrumentItem)
	}

	logs, sub, err := _NeutralToken.contract.FilterLogs(opts, "Settlement", _sourceRule, _destinationRule, _instrumentRule)
	if err != nil {
		return nil, err
	}
	return &NeutralTokenSettlementIterator{contract: _NeutralToken.contract, event: "Settlement", logs: logs, sub: sub}, nil
}

// WatchSettlement is a free log subscription operation binding the contract event 0x146934e5f2b06d949bf7603893fa5e00cbc5c1dd49bf19d0b8b33c8ac9ab8d99.
//
// Solidity: event Settlement(address _feeDestination, address indexed _source, address indexed _destination, address indexed _instrument, uint256 _instrumentQuantity, uint8 _instrumentOperation, uint256 _neutralQuantity, uint256 _fee, uint256 _neutralBoundary)
func (_NeutralToken *NeutralTokenFilterer) WatchSettlement(opts *bind.WatchOpts, sink chan<- *NeutralTokenSettlement, _source []common.Address, _destination []common.Address, _instrument []common.Address) (event.Subscription, error) {

	var _sourceRule []interface{}
	for _, _sourceItem := range _source {
		_sourceRule = append(_sourceRule, _sourceItem)
	}
	var _destinationRule []interface{}
	for _, _destinationItem := range _destination {
		_destinationRule = append(_destinationRule, _destinationItem)
	}
	var _instrumentRule []interface{}
	for _, _instrumentItem := range _instrument {
		_instrumentRule = append(_instrumentRule, _instrumentItem)
	}

	logs, sub, err := _NeutralToken.contract.WatchLogs(opts, "Settlement", _sourceRule, _destinationRule, _instrumentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeutralTokenSettlement)
				if err := _NeutralToken.contract.UnpackLog(event, "Settlement", log); err != nil {
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

// NeutralTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NeutralToken contract.
type NeutralTokenTransferIterator struct {
	Event *NeutralTokenTransfer // Event containing the contract specifics and raw log

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
func (it *NeutralTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeutralTokenTransfer)
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
		it.Event = new(NeutralTokenTransfer)
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
func (it *NeutralTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeutralTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeutralTokenTransfer represents a Transfer event raised by the NeutralToken contract.
type NeutralTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NeutralToken *NeutralTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NeutralTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NeutralToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NeutralTokenTransferIterator{contract: _NeutralToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NeutralToken *NeutralTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NeutralTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NeutralToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeutralTokenTransfer)
				if err := _NeutralToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

