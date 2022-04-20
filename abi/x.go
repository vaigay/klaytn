// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"math/big"
	"strings"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = klaytn.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AbiABI is the input ABI used to generate the binding from.
const AbiABI = "[{\"inputs\":[],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"updateAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_object\",\"type\":\"address\"}],\"name\":\"blockObject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"getBlock\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_object\",\"type\":\"address\"}],\"name\":\"unblockObject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AbiBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const AbiBinRuntime = ``

// Abi is an auto generated Go binding around a Klaytn contract.
type Abi struct {
	AbiCaller     // Read-only binding to the contract
	AbiTransactor // Write-only binding to the contract
	AbiFilterer   // Log filterer for contract events
}

// AbiCaller is an auto generated read-only Go binding around a Klaytn contract.
type AbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiTransactor is an auto generated write-only Go binding around a Klaytn contract.
type AbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type AbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type AbiSession struct {
	Contract     *Abi              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type AbiCallerSession struct {
	Contract *AbiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbiTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type AbiTransactorSession struct {
	Contract     *AbiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiRaw is an auto generated low-level Go binding around a Klaytn contract.
type AbiRaw struct {
	Contract *Abi // Generic contract binding to access the raw methods on
}

// AbiCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type AbiCallerRaw struct {
	Contract *AbiCaller // Generic read-only contract binding to access the raw methods on
}

// AbiTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type AbiTransactorRaw struct {
	Contract *AbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbi creates a new instance of Abi, bound to a specific deployed contract.
func NewAbi(address common.Address, backend bind.ContractBackend) (*Abi, error) {
	contract, err := bindAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abi{AbiCaller: AbiCaller{contract: contract}, AbiTransactor: AbiTransactor{contract: contract}, AbiFilterer: AbiFilterer{contract: contract}}, nil
}

// NewAbiCaller creates a new read-only instance of Abi, bound to a specific deployed contract.
func NewAbiCaller(address common.Address, caller bind.ContractCaller) (*AbiCaller, error) {
	contract, err := bindAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbiCaller{contract: contract}, nil
}

// NewAbiTransactor creates a new write-only instance of Abi, bound to a specific deployed contract.
func NewAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*AbiTransactor, error) {
	contract, err := bindAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbiTransactor{contract: contract}, nil
}

// NewAbiFilterer creates a new log filterer instance of Abi, bound to a specific deployed contract.
func NewAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*AbiFilterer, error) {
	contract, err := bindAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbiFilterer{contract: contract}, nil
}

// bindAbi binds a generic wrapper to an already deployed contract.
func bindAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.AbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transact(opts, method, params...)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Abi *AbiCaller) GetAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Abi.contract.Call(opts, out, "getAdmin")
	return *ret0, err
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Abi *AbiSession) GetAdmin() (common.Address, error) {
	return _Abi.Contract.GetAdmin(&_Abi.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Abi *AbiCallerSession) GetAdmin() (common.Address, error) {
	return _Abi.Contract.GetAdmin(&_Abi.CallOpts)
}

// GetBlock is a free data retrieval call binding the contract method 0x0a169068.
//
// Solidity: function getBlock(address[] _addresses) view returns(address)
func (_Abi *AbiCaller) GetBlock(opts *bind.CallOpts, _addresses []common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Abi.contract.Call(opts, out, "getBlock", _addresses)
	return *ret0, err
}

// GetBlock is a free data retrieval call binding the contract method 0x0a169068.
//
// Solidity: function getBlock(address[] _addresses) view returns(address)
func (_Abi *AbiSession) GetBlock(_addresses []common.Address) (common.Address, error) {
	return _Abi.Contract.GetBlock(&_Abi.CallOpts, _addresses)
}

// GetBlock is a free data retrieval call binding the contract method 0x0a169068.
//
// Solidity: function getBlock(address[] _addresses) view returns(address)
func (_Abi *AbiCallerSession) GetBlock(_addresses []common.Address) (common.Address, error) {
	return _Abi.Contract.GetBlock(&_Abi.CallOpts, _addresses)
}

// BlockObject is a paid mutator transaction binding the contract method 0xbc847732.
//
// Solidity: function blockObject(address _object) returns()
func (_Abi *AbiTransactor) BlockObject(opts *bind.TransactOpts, _object common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "blockObject", _object)
}

// BlockObject is a paid mutator transaction binding the contract method 0xbc847732.
//
// Solidity: function blockObject(address _object) returns()
func (_Abi *AbiSession) BlockObject(_object common.Address) (*types.Transaction, error) {
	return _Abi.Contract.BlockObject(&_Abi.TransactOpts, _object)
}

// BlockObject is a paid mutator transaction binding the contract method 0xbc847732.
//
// Solidity: function blockObject(address _object) returns()
func (_Abi *AbiTransactorSession) BlockObject(_object common.Address) (*types.Transaction, error) {
	return _Abi.Contract.BlockObject(&_Abi.TransactOpts, _object)
}

// UnblockObject is a paid mutator transaction binding the contract method 0x2f12b6c3.
//
// Solidity: function unblockObject(address _object) returns()
func (_Abi *AbiTransactor) UnblockObject(opts *bind.TransactOpts, _object common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "unblockObject", _object)
}

// UnblockObject is a paid mutator transaction binding the contract method 0x2f12b6c3.
//
// Solidity: function unblockObject(address _object) returns()
func (_Abi *AbiSession) UnblockObject(_object common.Address) (*types.Transaction, error) {
	return _Abi.Contract.UnblockObject(&_Abi.TransactOpts, _object)
}

// UnblockObject is a paid mutator transaction binding the contract method 0x2f12b6c3.
//
// Solidity: function unblockObject(address _object) returns()
func (_Abi *AbiTransactorSession) UnblockObject(_object common.Address) (*types.Transaction, error) {
	return _Abi.Contract.UnblockObject(&_Abi.TransactOpts, _object)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xe2f273bd.
//
// Solidity: function updateAdmin(address _newAdmin) returns()
func (_Abi *AbiTransactor) UpdateAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "updateAdmin", _newAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xe2f273bd.
//
// Solidity: function updateAdmin(address _newAdmin) returns()
func (_Abi *AbiSession) UpdateAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Abi.Contract.UpdateAdmin(&_Abi.TransactOpts, _newAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xe2f273bd.
//
// Solidity: function updateAdmin(address _newAdmin) returns()
func (_Abi *AbiTransactorSession) UpdateAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Abi.Contract.UpdateAdmin(&_Abi.TransactOpts, _newAdmin)
}
