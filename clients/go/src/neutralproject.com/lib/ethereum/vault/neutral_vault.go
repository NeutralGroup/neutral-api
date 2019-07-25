// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package neutral_vault

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

// NeutralVaultABI is the input ABI used to generate the binding from.
const NeutralVaultABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"permissions\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x1f9838b5\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x8f601f66\"},{\"constant\":true,\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xa3f4df7e\"},{\"constant\":true,\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xffa1ad74\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_forBenefitOf\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_quantity\",\"type\":\"uint256\"}],\"name\":\"EventDeposit\",\"type\":\"event\",\"signature\":\"0xc1688a9c6f6cbeae1273c097de87e31f0ead87a885c837b465128530bf4f0a6b\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_quantity\",\"type\":\"uint256\"}],\"name\":\"EventWithdraw\",\"type\":\"event\",\"signature\":\"0xfbd82180d75f2c932d80f8507177f579a217c3a17512ca663e5268e6b0f5e618\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_caller\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_quantity\",\"type\":\"uint256\"}],\"name\":\"EventTransfer\",\"type\":\"event\",\"signature\":\"0x95b51c2e04b16e13b8a19dc7fb13f8b67524965c055a397ce2103599e60d5145\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_who\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_grant\",\"type\":\"address\"}],\"name\":\"EventGrant\",\"type\":\"event\",\"signature\":\"0x65827c2b1ba78e5ec9e2564752e42a49acb6e22daf74362a82607d4ce895ea9e\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_who\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_revokedCaller\",\"type\":\"address\"}],\"name\":\"EventRevoke\",\"type\":\"event\",\"signature\":\"0xbec7db8863156b5a6fe40aea301cf11e61f98a4346764b67de4f1014962e5f9b\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"EventOptionalManagement\",\"type\":\"event\",\"signature\":\"0x8cd81c49a14bd177077f12d9b0c412f01894dda43cd0754e86deaa2fccefe35b\"},{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"readAvailableBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x5d973e47\"},{\"constant\":true,\"inputs\":[{\"name\":\"_vaultid\",\"type\":\"address\"},{\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"readPermission\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xb1383fdc\"},{\"constant\":false,\"inputs\":[{\"name\":\"_forBenefitOf\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_quantity\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x8340f549\"},{\"constant\":false,\"inputs\":[{\"name\":\"_where\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_quantity\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xd9caed12\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_quantity\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xf18d03cc\"},{\"constant\":false,\"inputs\":[{\"name\":\"_grant\",\"type\":\"address\"}],\"name\":\"grant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x70284d19\"},{\"constant\":false,\"inputs\":[{\"name\":\"_revokedCaller\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x74a8f103\"},{\"constant\":false,\"inputs\":[],\"name\":\"optionalManagement\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xb5f5e722\"}]"

// NeutralVault is an auto generated Go binding around an Ethereum contract.
type NeutralVault struct {
	NeutralVaultCaller     // Read-only binding to the contract
	NeutralVaultTransactor // Write-only binding to the contract
	NeutralVaultFilterer   // Log filterer for contract events
}

// NeutralVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type NeutralVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NeutralVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NeutralVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NeutralVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NeutralVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NeutralVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NeutralVaultSession struct {
	Contract     *NeutralVault     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NeutralVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NeutralVaultCallerSession struct {
	Contract *NeutralVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// NeutralVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NeutralVaultTransactorSession struct {
	Contract     *NeutralVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NeutralVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type NeutralVaultRaw struct {
	Contract *NeutralVault // Generic contract binding to access the raw methods on
}

// NeutralVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NeutralVaultCallerRaw struct {
	Contract *NeutralVaultCaller // Generic read-only contract binding to access the raw methods on
}

// NeutralVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NeutralVaultTransactorRaw struct {
	Contract *NeutralVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNeutralVault creates a new instance of NeutralVault, bound to a specific deployed contract.
func NewNeutralVault(address common.Address, backend bind.ContractBackend) (*NeutralVault, error) {
	contract, err := bindNeutralVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NeutralVault{NeutralVaultCaller: NeutralVaultCaller{contract: contract}, NeutralVaultTransactor: NeutralVaultTransactor{contract: contract}, NeutralVaultFilterer: NeutralVaultFilterer{contract: contract}}, nil
}

// NewNeutralVaultCaller creates a new read-only instance of NeutralVault, bound to a specific deployed contract.
func NewNeutralVaultCaller(address common.Address, caller bind.ContractCaller) (*NeutralVaultCaller, error) {
	contract, err := bindNeutralVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NeutralVaultCaller{contract: contract}, nil
}

// NewNeutralVaultTransactor creates a new write-only instance of NeutralVault, bound to a specific deployed contract.
func NewNeutralVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*NeutralVaultTransactor, error) {
	contract, err := bindNeutralVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NeutralVaultTransactor{contract: contract}, nil
}

// NewNeutralVaultFilterer creates a new log filterer instance of NeutralVault, bound to a specific deployed contract.
func NewNeutralVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*NeutralVaultFilterer, error) {
	contract, err := bindNeutralVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NeutralVaultFilterer{contract: contract}, nil
}

// bindNeutralVault binds a generic wrapper to an already deployed contract.
func bindNeutralVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NeutralVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NeutralVault *NeutralVaultRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NeutralVault.Contract.NeutralVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NeutralVault *NeutralVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeutralVault.Contract.NeutralVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NeutralVault *NeutralVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NeutralVault.Contract.NeutralVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NeutralVault *NeutralVaultCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NeutralVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NeutralVault *NeutralVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeutralVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NeutralVault *NeutralVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NeutralVault.Contract.contract.Transact(opts, method, params...)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() constant returns(string)
func (_NeutralVault *NeutralVaultCaller) NAME(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NeutralVault.contract.Call(opts, out, "NAME")
	return *ret0, err
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() constant returns(string)
func (_NeutralVault *NeutralVaultSession) NAME() (string, error) {
	return _NeutralVault.Contract.NAME(&_NeutralVault.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() constant returns(string)
func (_NeutralVault *NeutralVaultCallerSession) NAME() (string, error) {
	return _NeutralVault.Contract.NAME(&_NeutralVault.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_NeutralVault *NeutralVaultCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NeutralVault.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_NeutralVault *NeutralVaultSession) VERSION() (string, error) {
	return _NeutralVault.Contract.VERSION(&_NeutralVault.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_NeutralVault *NeutralVaultCallerSession) VERSION() (string, error) {
	return _NeutralVault.Contract.VERSION(&_NeutralVault.CallOpts)
}

// Deposits is a free data retrieval call binding the contract method 0x8f601f66.
//
// Solidity: function deposits(address , address ) constant returns(uint256)
func (_NeutralVault *NeutralVaultCaller) Deposits(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NeutralVault.contract.Call(opts, out, "deposits", arg0, arg1)
	return *ret0, err
}

// Deposits is a free data retrieval call binding the contract method 0x8f601f66.
//
// Solidity: function deposits(address , address ) constant returns(uint256)
func (_NeutralVault *NeutralVaultSession) Deposits(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _NeutralVault.Contract.Deposits(&_NeutralVault.CallOpts, arg0, arg1)
}

// Deposits is a free data retrieval call binding the contract method 0x8f601f66.
//
// Solidity: function deposits(address , address ) constant returns(uint256)
func (_NeutralVault *NeutralVaultCallerSession) Deposits(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _NeutralVault.Contract.Deposits(&_NeutralVault.CallOpts, arg0, arg1)
}

// Permissions is a free data retrieval call binding the contract method 0x1f9838b5.
//
// Solidity: function permissions(address , address ) constant returns(bool)
func (_NeutralVault *NeutralVaultCaller) Permissions(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NeutralVault.contract.Call(opts, out, "permissions", arg0, arg1)
	return *ret0, err
}

// Permissions is a free data retrieval call binding the contract method 0x1f9838b5.
//
// Solidity: function permissions(address , address ) constant returns(bool)
func (_NeutralVault *NeutralVaultSession) Permissions(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _NeutralVault.Contract.Permissions(&_NeutralVault.CallOpts, arg0, arg1)
}

// Permissions is a free data retrieval call binding the contract method 0x1f9838b5.
//
// Solidity: function permissions(address , address ) constant returns(bool)
func (_NeutralVault *NeutralVaultCallerSession) Permissions(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _NeutralVault.Contract.Permissions(&_NeutralVault.CallOpts, arg0, arg1)
}

// ReadAvailableBalance is a free data retrieval call binding the contract method 0x5d973e47.
//
// Solidity: function readAvailableBalance(address _who, address _token) constant returns(uint256)
func (_NeutralVault *NeutralVaultCaller) ReadAvailableBalance(opts *bind.CallOpts, _who common.Address, _token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NeutralVault.contract.Call(opts, out, "readAvailableBalance", _who, _token)
	return *ret0, err
}

// ReadAvailableBalance is a free data retrieval call binding the contract method 0x5d973e47.
//
// Solidity: function readAvailableBalance(address _who, address _token) constant returns(uint256)
func (_NeutralVault *NeutralVaultSession) ReadAvailableBalance(_who common.Address, _token common.Address) (*big.Int, error) {
	return _NeutralVault.Contract.ReadAvailableBalance(&_NeutralVault.CallOpts, _who, _token)
}

// ReadAvailableBalance is a free data retrieval call binding the contract method 0x5d973e47.
//
// Solidity: function readAvailableBalance(address _who, address _token) constant returns(uint256)
func (_NeutralVault *NeutralVaultCallerSession) ReadAvailableBalance(_who common.Address, _token common.Address) (*big.Int, error) {
	return _NeutralVault.Contract.ReadAvailableBalance(&_NeutralVault.CallOpts, _who, _token)
}

// ReadPermission is a free data retrieval call binding the contract method 0xb1383fdc.
//
// Solidity: function readPermission(address _vaultid, address _caller) constant returns(bool)
func (_NeutralVault *NeutralVaultCaller) ReadPermission(opts *bind.CallOpts, _vaultid common.Address, _caller common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NeutralVault.contract.Call(opts, out, "readPermission", _vaultid, _caller)
	return *ret0, err
}

// ReadPermission is a free data retrieval call binding the contract method 0xb1383fdc.
//
// Solidity: function readPermission(address _vaultid, address _caller) constant returns(bool)
func (_NeutralVault *NeutralVaultSession) ReadPermission(_vaultid common.Address, _caller common.Address) (bool, error) {
	return _NeutralVault.Contract.ReadPermission(&_NeutralVault.CallOpts, _vaultid, _caller)
}

// ReadPermission is a free data retrieval call binding the contract method 0xb1383fdc.
//
// Solidity: function readPermission(address _vaultid, address _caller) constant returns(bool)
func (_NeutralVault *NeutralVaultCallerSession) ReadPermission(_vaultid common.Address, _caller common.Address) (bool, error) {
	return _NeutralVault.Contract.ReadPermission(&_NeutralVault.CallOpts, _vaultid, _caller)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address _forBenefitOf, address _token, uint256 _quantity) returns(bool)
func (_NeutralVault *NeutralVaultTransactor) Deposit(opts *bind.TransactOpts, _forBenefitOf common.Address, _token common.Address, _quantity *big.Int) (*types.Transaction, error) {
	return _NeutralVault.contract.Transact(opts, "deposit", _forBenefitOf, _token, _quantity)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address _forBenefitOf, address _token, uint256 _quantity) returns(bool)
func (_NeutralVault *NeutralVaultSession) Deposit(_forBenefitOf common.Address, _token common.Address, _quantity *big.Int) (*types.Transaction, error) {
	return _NeutralVault.Contract.Deposit(&_NeutralVault.TransactOpts, _forBenefitOf, _token, _quantity)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address _forBenefitOf, address _token, uint256 _quantity) returns(bool)
func (_NeutralVault *NeutralVaultTransactorSession) Deposit(_forBenefitOf common.Address, _token common.Address, _quantity *big.Int) (*types.Transaction, error) {
	return _NeutralVault.Contract.Deposit(&_NeutralVault.TransactOpts, _forBenefitOf, _token, _quantity)
}

// Grant is a paid mutator transaction binding the contract method 0x70284d19.
//
// Solidity: function grant(address _grant) returns(bool)
func (_NeutralVault *NeutralVaultTransactor) Grant(opts *bind.TransactOpts, _grant common.Address) (*types.Transaction, error) {
	return _NeutralVault.contract.Transact(opts, "grant", _grant)
}

// Grant is a paid mutator transaction binding the contract method 0x70284d19.
//
// Solidity: function grant(address _grant) returns(bool)
func (_NeutralVault *NeutralVaultSession) Grant(_grant common.Address) (*types.Transaction, error) {
	return _NeutralVault.Contract.Grant(&_NeutralVault.TransactOpts, _grant)
}

// Grant is a paid mutator transaction binding the contract method 0x70284d19.
//
// Solidity: function grant(address _grant) returns(bool)
func (_NeutralVault *NeutralVaultTransactorSession) Grant(_grant common.Address) (*types.Transaction, error) {
	return _NeutralVault.Contract.Grant(&_NeutralVault.TransactOpts, _grant)
}

// OptionalManagement is a paid mutator transaction binding the contract method 0xb5f5e722.
//
// Solidity: function optionalManagement() returns(bool)
func (_NeutralVault *NeutralVaultTransactor) OptionalManagement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NeutralVault.contract.Transact(opts, "optionalManagement")
}

// OptionalManagement is a paid mutator transaction binding the contract method 0xb5f5e722.
//
// Solidity: function optionalManagement() returns(bool)
func (_NeutralVault *NeutralVaultSession) OptionalManagement() (*types.Transaction, error) {
	return _NeutralVault.Contract.OptionalManagement(&_NeutralVault.TransactOpts)
}

// OptionalManagement is a paid mutator transaction binding the contract method 0xb5f5e722.
//
// Solidity: function optionalManagement() returns(bool)
func (_NeutralVault *NeutralVaultTransactorSession) OptionalManagement() (*types.Transaction, error) {
	return _NeutralVault.Contract.OptionalManagement(&_NeutralVault.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address _revokedCaller) returns(bool)
func (_NeutralVault *NeutralVaultTransactor) Revoke(opts *bind.TransactOpts, _revokedCaller common.Address) (*types.Transaction, error) {
	return _NeutralVault.contract.Transact(opts, "revoke", _revokedCaller)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address _revokedCaller) returns(bool)
func (_NeutralVault *NeutralVaultSession) Revoke(_revokedCaller common.Address) (*types.Transaction, error) {
	return _NeutralVault.Contract.Revoke(&_NeutralVault.TransactOpts, _revokedCaller)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address _revokedCaller) returns(bool)
func (_NeutralVault *NeutralVaultTransactorSession) Revoke(_revokedCaller common.Address) (*types.Transaction, error) {
	return _NeutralVault.Contract.Revoke(&_NeutralVault.TransactOpts, _revokedCaller)
}

// Transfer is a paid mutator transaction binding the contract method 0xf18d03cc.
//
// Solidity: function transfer(address _from, address _to, address _token, uint256 _quantity) returns(bool)
func (_NeutralVault *NeutralVaultTransactor) Transfer(opts *bind.TransactOpts, _from common.Address, _to common.Address, _token common.Address, _quantity *big.Int) (*types.Transaction, error) {
	return _NeutralVault.contract.Transact(opts, "transfer", _from, _to, _token, _quantity)
}

// Transfer is a paid mutator transaction binding the contract method 0xf18d03cc.
//
// Solidity: function transfer(address _from, address _to, address _token, uint256 _quantity) returns(bool)
func (_NeutralVault *NeutralVaultSession) Transfer(_from common.Address, _to common.Address, _token common.Address, _quantity *big.Int) (*types.Transaction, error) {
	return _NeutralVault.Contract.Transfer(&_NeutralVault.TransactOpts, _from, _to, _token, _quantity)
}

// Transfer is a paid mutator transaction binding the contract method 0xf18d03cc.
//
// Solidity: function transfer(address _from, address _to, address _token, uint256 _quantity) returns(bool)
func (_NeutralVault *NeutralVaultTransactorSession) Transfer(_from common.Address, _to common.Address, _token common.Address, _quantity *big.Int) (*types.Transaction, error) {
	return _NeutralVault.Contract.Transfer(&_NeutralVault.TransactOpts, _from, _to, _token, _quantity)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _where, address _token, uint256 _quantity) returns(bool)
func (_NeutralVault *NeutralVaultTransactor) Withdraw(opts *bind.TransactOpts, _where common.Address, _token common.Address, _quantity *big.Int) (*types.Transaction, error) {
	return _NeutralVault.contract.Transact(opts, "withdraw", _where, _token, _quantity)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _where, address _token, uint256 _quantity) returns(bool)
func (_NeutralVault *NeutralVaultSession) Withdraw(_where common.Address, _token common.Address, _quantity *big.Int) (*types.Transaction, error) {
	return _NeutralVault.Contract.Withdraw(&_NeutralVault.TransactOpts, _where, _token, _quantity)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address _where, address _token, uint256 _quantity) returns(bool)
func (_NeutralVault *NeutralVaultTransactorSession) Withdraw(_where common.Address, _token common.Address, _quantity *big.Int) (*types.Transaction, error) {
	return _NeutralVault.Contract.Withdraw(&_NeutralVault.TransactOpts, _where, _token, _quantity)
}

// NeutralVaultEventDepositIterator is returned from FilterEventDeposit and is used to iterate over the raw logs and unpacked data for EventDeposit events raised by the NeutralVault contract.
type NeutralVaultEventDepositIterator struct {
	Event *NeutralVaultEventDeposit // Event containing the contract specifics and raw log

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
func (it *NeutralVaultEventDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeutralVaultEventDeposit)
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
		it.Event = new(NeutralVaultEventDeposit)
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
func (it *NeutralVaultEventDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeutralVaultEventDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeutralVaultEventDeposit represents a EventDeposit event raised by the NeutralVault contract.
type NeutralVaultEventDeposit struct {
	From         common.Address
	ForBenefitOf common.Address
	Token        common.Address
	Quantity     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEventDeposit is a free log retrieval operation binding the contract event 0xc1688a9c6f6cbeae1273c097de87e31f0ead87a885c837b465128530bf4f0a6b.
//
// Solidity: event EventDeposit(address indexed _from, address indexed _forBenefitOf, address indexed _token, uint256 _quantity)
func (_NeutralVault *NeutralVaultFilterer) FilterEventDeposit(opts *bind.FilterOpts, _from []common.Address, _forBenefitOf []common.Address, _token []common.Address) (*NeutralVaultEventDepositIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _forBenefitOfRule []interface{}
	for _, _forBenefitOfItem := range _forBenefitOf {
		_forBenefitOfRule = append(_forBenefitOfRule, _forBenefitOfItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _NeutralVault.contract.FilterLogs(opts, "EventDeposit", _fromRule, _forBenefitOfRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return &NeutralVaultEventDepositIterator{contract: _NeutralVault.contract, event: "EventDeposit", logs: logs, sub: sub}, nil
}

// WatchEventDeposit is a free log subscription operation binding the contract event 0xc1688a9c6f6cbeae1273c097de87e31f0ead87a885c837b465128530bf4f0a6b.
//
// Solidity: event EventDeposit(address indexed _from, address indexed _forBenefitOf, address indexed _token, uint256 _quantity)
func (_NeutralVault *NeutralVaultFilterer) WatchEventDeposit(opts *bind.WatchOpts, sink chan<- *NeutralVaultEventDeposit, _from []common.Address, _forBenefitOf []common.Address, _token []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _forBenefitOfRule []interface{}
	for _, _forBenefitOfItem := range _forBenefitOf {
		_forBenefitOfRule = append(_forBenefitOfRule, _forBenefitOfItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _NeutralVault.contract.WatchLogs(opts, "EventDeposit", _fromRule, _forBenefitOfRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeutralVaultEventDeposit)
				if err := _NeutralVault.contract.UnpackLog(event, "EventDeposit", log); err != nil {
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

// NeutralVaultEventGrantIterator is returned from FilterEventGrant and is used to iterate over the raw logs and unpacked data for EventGrant events raised by the NeutralVault contract.
type NeutralVaultEventGrantIterator struct {
	Event *NeutralVaultEventGrant // Event containing the contract specifics and raw log

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
func (it *NeutralVaultEventGrantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeutralVaultEventGrant)
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
		it.Event = new(NeutralVaultEventGrant)
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
func (it *NeutralVaultEventGrantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeutralVaultEventGrantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeutralVaultEventGrant represents a EventGrant event raised by the NeutralVault contract.
type NeutralVaultEventGrant struct {
	Who   common.Address
	Grant common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventGrant is a free log retrieval operation binding the contract event 0x65827c2b1ba78e5ec9e2564752e42a49acb6e22daf74362a82607d4ce895ea9e.
//
// Solidity: event EventGrant(address indexed _who, address indexed _grant)
func (_NeutralVault *NeutralVaultFilterer) FilterEventGrant(opts *bind.FilterOpts, _who []common.Address, _grant []common.Address) (*NeutralVaultEventGrantIterator, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}
	var _grantRule []interface{}
	for _, _grantItem := range _grant {
		_grantRule = append(_grantRule, _grantItem)
	}

	logs, sub, err := _NeutralVault.contract.FilterLogs(opts, "EventGrant", _whoRule, _grantRule)
	if err != nil {
		return nil, err
	}
	return &NeutralVaultEventGrantIterator{contract: _NeutralVault.contract, event: "EventGrant", logs: logs, sub: sub}, nil
}

// WatchEventGrant is a free log subscription operation binding the contract event 0x65827c2b1ba78e5ec9e2564752e42a49acb6e22daf74362a82607d4ce895ea9e.
//
// Solidity: event EventGrant(address indexed _who, address indexed _grant)
func (_NeutralVault *NeutralVaultFilterer) WatchEventGrant(opts *bind.WatchOpts, sink chan<- *NeutralVaultEventGrant, _who []common.Address, _grant []common.Address) (event.Subscription, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}
	var _grantRule []interface{}
	for _, _grantItem := range _grant {
		_grantRule = append(_grantRule, _grantItem)
	}

	logs, sub, err := _NeutralVault.contract.WatchLogs(opts, "EventGrant", _whoRule, _grantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeutralVaultEventGrant)
				if err := _NeutralVault.contract.UnpackLog(event, "EventGrant", log); err != nil {
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

// NeutralVaultEventOptionalManagementIterator is returned from FilterEventOptionalManagement and is used to iterate over the raw logs and unpacked data for EventOptionalManagement events raised by the NeutralVault contract.
type NeutralVaultEventOptionalManagementIterator struct {
	Event *NeutralVaultEventOptionalManagement // Event containing the contract specifics and raw log

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
func (it *NeutralVaultEventOptionalManagementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeutralVaultEventOptionalManagement)
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
		it.Event = new(NeutralVaultEventOptionalManagement)
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
func (it *NeutralVaultEventOptionalManagementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeutralVaultEventOptionalManagementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeutralVaultEventOptionalManagement represents a EventOptionalManagement event raised by the NeutralVault contract.
type NeutralVaultEventOptionalManagement struct {
	Token  common.Address
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEventOptionalManagement is a free log retrieval operation binding the contract event 0x8cd81c49a14bd177077f12d9b0c412f01894dda43cd0754e86deaa2fccefe35b.
//
// Solidity: event EventOptionalManagement(address indexed _token, bool _status)
func (_NeutralVault *NeutralVaultFilterer) FilterEventOptionalManagement(opts *bind.FilterOpts, _token []common.Address) (*NeutralVaultEventOptionalManagementIterator, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _NeutralVault.contract.FilterLogs(opts, "EventOptionalManagement", _tokenRule)
	if err != nil {
		return nil, err
	}
	return &NeutralVaultEventOptionalManagementIterator{contract: _NeutralVault.contract, event: "EventOptionalManagement", logs: logs, sub: sub}, nil
}

// WatchEventOptionalManagement is a free log subscription operation binding the contract event 0x8cd81c49a14bd177077f12d9b0c412f01894dda43cd0754e86deaa2fccefe35b.
//
// Solidity: event EventOptionalManagement(address indexed _token, bool _status)
func (_NeutralVault *NeutralVaultFilterer) WatchEventOptionalManagement(opts *bind.WatchOpts, sink chan<- *NeutralVaultEventOptionalManagement, _token []common.Address) (event.Subscription, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _NeutralVault.contract.WatchLogs(opts, "EventOptionalManagement", _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeutralVaultEventOptionalManagement)
				if err := _NeutralVault.contract.UnpackLog(event, "EventOptionalManagement", log); err != nil {
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

// NeutralVaultEventRevokeIterator is returned from FilterEventRevoke and is used to iterate over the raw logs and unpacked data for EventRevoke events raised by the NeutralVault contract.
type NeutralVaultEventRevokeIterator struct {
	Event *NeutralVaultEventRevoke // Event containing the contract specifics and raw log

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
func (it *NeutralVaultEventRevokeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeutralVaultEventRevoke)
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
		it.Event = new(NeutralVaultEventRevoke)
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
func (it *NeutralVaultEventRevokeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeutralVaultEventRevokeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeutralVaultEventRevoke represents a EventRevoke event raised by the NeutralVault contract.
type NeutralVaultEventRevoke struct {
	Who           common.Address
	RevokedCaller common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEventRevoke is a free log retrieval operation binding the contract event 0xbec7db8863156b5a6fe40aea301cf11e61f98a4346764b67de4f1014962e5f9b.
//
// Solidity: event EventRevoke(address indexed _who, address indexed _revokedCaller)
func (_NeutralVault *NeutralVaultFilterer) FilterEventRevoke(opts *bind.FilterOpts, _who []common.Address, _revokedCaller []common.Address) (*NeutralVaultEventRevokeIterator, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}
	var _revokedCallerRule []interface{}
	for _, _revokedCallerItem := range _revokedCaller {
		_revokedCallerRule = append(_revokedCallerRule, _revokedCallerItem)
	}

	logs, sub, err := _NeutralVault.contract.FilterLogs(opts, "EventRevoke", _whoRule, _revokedCallerRule)
	if err != nil {
		return nil, err
	}
	return &NeutralVaultEventRevokeIterator{contract: _NeutralVault.contract, event: "EventRevoke", logs: logs, sub: sub}, nil
}

// WatchEventRevoke is a free log subscription operation binding the contract event 0xbec7db8863156b5a6fe40aea301cf11e61f98a4346764b67de4f1014962e5f9b.
//
// Solidity: event EventRevoke(address indexed _who, address indexed _revokedCaller)
func (_NeutralVault *NeutralVaultFilterer) WatchEventRevoke(opts *bind.WatchOpts, sink chan<- *NeutralVaultEventRevoke, _who []common.Address, _revokedCaller []common.Address) (event.Subscription, error) {

	var _whoRule []interface{}
	for _, _whoItem := range _who {
		_whoRule = append(_whoRule, _whoItem)
	}
	var _revokedCallerRule []interface{}
	for _, _revokedCallerItem := range _revokedCaller {
		_revokedCallerRule = append(_revokedCallerRule, _revokedCallerItem)
	}

	logs, sub, err := _NeutralVault.contract.WatchLogs(opts, "EventRevoke", _whoRule, _revokedCallerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeutralVaultEventRevoke)
				if err := _NeutralVault.contract.UnpackLog(event, "EventRevoke", log); err != nil {
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

// NeutralVaultEventTransferIterator is returned from FilterEventTransfer and is used to iterate over the raw logs and unpacked data for EventTransfer events raised by the NeutralVault contract.
type NeutralVaultEventTransferIterator struct {
	Event *NeutralVaultEventTransfer // Event containing the contract specifics and raw log

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
func (it *NeutralVaultEventTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeutralVaultEventTransfer)
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
		it.Event = new(NeutralVaultEventTransfer)
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
func (it *NeutralVaultEventTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeutralVaultEventTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeutralVaultEventTransfer represents a EventTransfer event raised by the NeutralVault contract.
type NeutralVaultEventTransfer struct {
	From     common.Address
	To       common.Address
	Caller   common.Address
	Token    common.Address
	Quantity *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEventTransfer is a free log retrieval operation binding the contract event 0x95b51c2e04b16e13b8a19dc7fb13f8b67524965c055a397ce2103599e60d5145.
//
// Solidity: event EventTransfer(address indexed _from, address indexed _to, address _caller, address indexed _token, uint256 _quantity)
func (_NeutralVault *NeutralVaultFilterer) FilterEventTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _token []common.Address) (*NeutralVaultEventTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _NeutralVault.contract.FilterLogs(opts, "EventTransfer", _fromRule, _toRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return &NeutralVaultEventTransferIterator{contract: _NeutralVault.contract, event: "EventTransfer", logs: logs, sub: sub}, nil
}

// WatchEventTransfer is a free log subscription operation binding the contract event 0x95b51c2e04b16e13b8a19dc7fb13f8b67524965c055a397ce2103599e60d5145.
//
// Solidity: event EventTransfer(address indexed _from, address indexed _to, address _caller, address indexed _token, uint256 _quantity)
func (_NeutralVault *NeutralVaultFilterer) WatchEventTransfer(opts *bind.WatchOpts, sink chan<- *NeutralVaultEventTransfer, _from []common.Address, _to []common.Address, _token []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _NeutralVault.contract.WatchLogs(opts, "EventTransfer", _fromRule, _toRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeutralVaultEventTransfer)
				if err := _NeutralVault.contract.UnpackLog(event, "EventTransfer", log); err != nil {
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

// NeutralVaultEventWithdrawIterator is returned from FilterEventWithdraw and is used to iterate over the raw logs and unpacked data for EventWithdraw events raised by the NeutralVault contract.
type NeutralVaultEventWithdrawIterator struct {
	Event *NeutralVaultEventWithdraw // Event containing the contract specifics and raw log

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
func (it *NeutralVaultEventWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NeutralVaultEventWithdraw)
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
		it.Event = new(NeutralVaultEventWithdraw)
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
func (it *NeutralVaultEventWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NeutralVaultEventWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NeutralVaultEventWithdraw represents a EventWithdraw event raised by the NeutralVault contract.
type NeutralVaultEventWithdraw struct {
	From     common.Address
	To       common.Address
	Token    common.Address
	Quantity *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEventWithdraw is a free log retrieval operation binding the contract event 0xfbd82180d75f2c932d80f8507177f579a217c3a17512ca663e5268e6b0f5e618.
//
// Solidity: event EventWithdraw(address indexed _from, address indexed _to, address indexed _token, uint256 _quantity)
func (_NeutralVault *NeutralVaultFilterer) FilterEventWithdraw(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _token []common.Address) (*NeutralVaultEventWithdrawIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _NeutralVault.contract.FilterLogs(opts, "EventWithdraw", _fromRule, _toRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return &NeutralVaultEventWithdrawIterator{contract: _NeutralVault.contract, event: "EventWithdraw", logs: logs, sub: sub}, nil
}

// WatchEventWithdraw is a free log subscription operation binding the contract event 0xfbd82180d75f2c932d80f8507177f579a217c3a17512ca663e5268e6b0f5e618.
//
// Solidity: event EventWithdraw(address indexed _from, address indexed _to, address indexed _token, uint256 _quantity)
func (_NeutralVault *NeutralVaultFilterer) WatchEventWithdraw(opts *bind.WatchOpts, sink chan<- *NeutralVaultEventWithdraw, _from []common.Address, _to []common.Address, _token []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _NeutralVault.contract.WatchLogs(opts, "EventWithdraw", _fromRule, _toRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NeutralVaultEventWithdraw)
				if err := _NeutralVault.contract.UnpackLog(event, "EventWithdraw", log); err != nil {
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

