// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package detailed_mock

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

// DetailedMockABI is the input ABI used to generate the binding from.
const DetailedMockABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x06fdde03\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x095ea7b3\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x18160ddd\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x23b872dd\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x313ce567\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x39509351\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x70a08231\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x95d89b41\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xa457c2d7\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xa9059cbb\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xdd62ed3e\"},{\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"},{\"name\":\"_initialBalance\",\"type\":\"uint256\"},{\"name\":\"_NeutralToken\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"signature\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\",\"signature\":\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\",\"signature\":\"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925\"},{\"constant\":false,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"},{\"name\":\"_quantity\",\"type\":\"uint256\"}],\"name\":\"set_balance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x0776e4fa\"}]"

// DetailedMock is an auto generated Go binding around an Ethereum contract.
type DetailedMock struct {
	DetailedMockCaller     // Read-only binding to the contract
	DetailedMockTransactor // Write-only binding to the contract
	DetailedMockFilterer   // Log filterer for contract events
}

// DetailedMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type DetailedMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DetailedMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DetailedMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DetailedMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DetailedMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DetailedMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DetailedMockSession struct {
	Contract     *DetailedMock     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DetailedMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DetailedMockCallerSession struct {
	Contract *DetailedMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DetailedMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DetailedMockTransactorSession struct {
	Contract     *DetailedMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DetailedMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type DetailedMockRaw struct {
	Contract *DetailedMock // Generic contract binding to access the raw methods on
}

// DetailedMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DetailedMockCallerRaw struct {
	Contract *DetailedMockCaller // Generic read-only contract binding to access the raw methods on
}

// DetailedMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DetailedMockTransactorRaw struct {
	Contract *DetailedMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDetailedMock creates a new instance of DetailedMock, bound to a specific deployed contract.
func NewDetailedMock(address common.Address, backend bind.ContractBackend) (*DetailedMock, error) {
	contract, err := bindDetailedMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DetailedMock{DetailedMockCaller: DetailedMockCaller{contract: contract}, DetailedMockTransactor: DetailedMockTransactor{contract: contract}, DetailedMockFilterer: DetailedMockFilterer{contract: contract}}, nil
}

// NewDetailedMockCaller creates a new read-only instance of DetailedMock, bound to a specific deployed contract.
func NewDetailedMockCaller(address common.Address, caller bind.ContractCaller) (*DetailedMockCaller, error) {
	contract, err := bindDetailedMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DetailedMockCaller{contract: contract}, nil
}

// NewDetailedMockTransactor creates a new write-only instance of DetailedMock, bound to a specific deployed contract.
func NewDetailedMockTransactor(address common.Address, transactor bind.ContractTransactor) (*DetailedMockTransactor, error) {
	contract, err := bindDetailedMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DetailedMockTransactor{contract: contract}, nil
}

// NewDetailedMockFilterer creates a new log filterer instance of DetailedMock, bound to a specific deployed contract.
func NewDetailedMockFilterer(address common.Address, filterer bind.ContractFilterer) (*DetailedMockFilterer, error) {
	contract, err := bindDetailedMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DetailedMockFilterer{contract: contract}, nil
}

// bindDetailedMock binds a generic wrapper to an already deployed contract.
func bindDetailedMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DetailedMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DetailedMock *DetailedMockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DetailedMock.Contract.DetailedMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DetailedMock *DetailedMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DetailedMock.Contract.DetailedMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DetailedMock *DetailedMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DetailedMock.Contract.DetailedMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DetailedMock *DetailedMockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DetailedMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DetailedMock *DetailedMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DetailedMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DetailedMock *DetailedMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DetailedMock.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_DetailedMock *DetailedMockCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DetailedMock.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_DetailedMock *DetailedMockSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DetailedMock.Contract.Allowance(&_DetailedMock.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_DetailedMock *DetailedMockCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DetailedMock.Contract.Allowance(&_DetailedMock.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_DetailedMock *DetailedMockCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DetailedMock.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_DetailedMock *DetailedMockSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DetailedMock.Contract.BalanceOf(&_DetailedMock.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_DetailedMock *DetailedMockCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DetailedMock.Contract.BalanceOf(&_DetailedMock.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_DetailedMock *DetailedMockCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _DetailedMock.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_DetailedMock *DetailedMockSession) Decimals() (uint8, error) {
	return _DetailedMock.Contract.Decimals(&_DetailedMock.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_DetailedMock *DetailedMockCallerSession) Decimals() (uint8, error) {
	return _DetailedMock.Contract.Decimals(&_DetailedMock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DetailedMock *DetailedMockCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DetailedMock.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DetailedMock *DetailedMockSession) Name() (string, error) {
	return _DetailedMock.Contract.Name(&_DetailedMock.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DetailedMock *DetailedMockCallerSession) Name() (string, error) {
	return _DetailedMock.Contract.Name(&_DetailedMock.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DetailedMock *DetailedMockCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DetailedMock.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DetailedMock *DetailedMockSession) Symbol() (string, error) {
	return _DetailedMock.Contract.Symbol(&_DetailedMock.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DetailedMock *DetailedMockCallerSession) Symbol() (string, error) {
	return _DetailedMock.Contract.Symbol(&_DetailedMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DetailedMock *DetailedMockCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DetailedMock.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DetailedMock *DetailedMockSession) TotalSupply() (*big.Int, error) {
	return _DetailedMock.Contract.TotalSupply(&_DetailedMock.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DetailedMock *DetailedMockCallerSession) TotalSupply() (*big.Int, error) {
	return _DetailedMock.Contract.TotalSupply(&_DetailedMock.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_DetailedMock *DetailedMockTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _DetailedMock.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_DetailedMock *DetailedMockSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _DetailedMock.Contract.Approve(&_DetailedMock.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_DetailedMock *DetailedMockTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _DetailedMock.Contract.Approve(&_DetailedMock.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DetailedMock *DetailedMockTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DetailedMock.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DetailedMock *DetailedMockSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DetailedMock.Contract.DecreaseAllowance(&_DetailedMock.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_DetailedMock *DetailedMockTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _DetailedMock.Contract.DecreaseAllowance(&_DetailedMock.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DetailedMock *DetailedMockTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DetailedMock.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DetailedMock *DetailedMockSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DetailedMock.Contract.IncreaseAllowance(&_DetailedMock.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_DetailedMock *DetailedMockTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _DetailedMock.Contract.IncreaseAllowance(&_DetailedMock.TransactOpts, spender, addedValue)
}

// SetBalance is a paid mutator transaction binding the contract method 0x0776e4fa.
//
// Solidity: function set_balance(address _who, uint256 _quantity) returns()
func (_DetailedMock *DetailedMockTransactor) SetBalance(opts *bind.TransactOpts, _who common.Address, _quantity *big.Int) (*types.Transaction, error) {
	return _DetailedMock.contract.Transact(opts, "set_balance", _who, _quantity)
}

// SetBalance is a paid mutator transaction binding the contract method 0x0776e4fa.
//
// Solidity: function set_balance(address _who, uint256 _quantity) returns()
func (_DetailedMock *DetailedMockSession) SetBalance(_who common.Address, _quantity *big.Int) (*types.Transaction, error) {
	return _DetailedMock.Contract.SetBalance(&_DetailedMock.TransactOpts, _who, _quantity)
}

// SetBalance is a paid mutator transaction binding the contract method 0x0776e4fa.
//
// Solidity: function set_balance(address _who, uint256 _quantity) returns()
func (_DetailedMock *DetailedMockTransactorSession) SetBalance(_who common.Address, _quantity *big.Int) (*types.Transaction, error) {
	return _DetailedMock.Contract.SetBalance(&_DetailedMock.TransactOpts, _who, _quantity)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_DetailedMock *DetailedMockTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DetailedMock.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_DetailedMock *DetailedMockSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DetailedMock.Contract.Transfer(&_DetailedMock.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_DetailedMock *DetailedMockTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DetailedMock.Contract.Transfer(&_DetailedMock.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_DetailedMock *DetailedMockTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DetailedMock.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_DetailedMock *DetailedMockSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DetailedMock.Contract.TransferFrom(&_DetailedMock.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_DetailedMock *DetailedMockTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _DetailedMock.Contract.TransferFrom(&_DetailedMock.TransactOpts, from, to, value)
}

// DetailedMockApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DetailedMock contract.
type DetailedMockApprovalIterator struct {
	Event *DetailedMockApproval // Event containing the contract specifics and raw log

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
func (it *DetailedMockApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DetailedMockApproval)
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
		it.Event = new(DetailedMockApproval)
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
func (it *DetailedMockApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DetailedMockApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DetailedMockApproval represents a Approval event raised by the DetailedMock contract.
type DetailedMockApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address spender, uint256 value)
func (_DetailedMock *DetailedMockFilterer) FilterApproval(opts *bind.FilterOpts) (*DetailedMockApprovalIterator, error) {

	logs, sub, err := _DetailedMock.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &DetailedMockApprovalIterator{contract: _DetailedMock.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address spender, uint256 value)
func (_DetailedMock *DetailedMockFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DetailedMockApproval) (event.Subscription, error) {

	logs, sub, err := _DetailedMock.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DetailedMockApproval)
				if err := _DetailedMock.contract.UnpackLog(event, "Approval", log); err != nil {
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

// DetailedMockTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DetailedMock contract.
type DetailedMockTransferIterator struct {
	Event *DetailedMockTransfer // Event containing the contract specifics and raw log

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
func (it *DetailedMockTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DetailedMockTransfer)
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
		it.Event = new(DetailedMockTransfer)
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
func (it *DetailedMockTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DetailedMockTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DetailedMockTransfer represents a Transfer event raised by the DetailedMock contract.
type DetailedMockTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 value)
func (_DetailedMock *DetailedMockFilterer) FilterTransfer(opts *bind.FilterOpts) (*DetailedMockTransferIterator, error) {

	logs, sub, err := _DetailedMock.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &DetailedMockTransferIterator{contract: _DetailedMock.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 value)
func (_DetailedMock *DetailedMockFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DetailedMockTransfer) (event.Subscription, error) {

	logs, sub, err := _DetailedMock.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DetailedMockTransfer)
				if err := _DetailedMock.contract.UnpackLog(event, "Transfer", log); err != nil {
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

