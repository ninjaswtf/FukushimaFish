// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bulktransfer

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

// BulktransferMetaData contains all meta data concerning the Bulktransfer contract.
var BulktransferMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokens\",\"type\":\"uint256[]\"}],\"name\":\"bulkTransferTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BulktransferABI is the input ABI used to generate the binding from.
// Deprecated: Use BulktransferMetaData.ABI instead.
var BulktransferABI = BulktransferMetaData.ABI

// Bulktransfer is an auto generated Go binding around an Ethereum contract.
type Bulktransfer struct {
	BulktransferCaller     // Read-only binding to the contract
	BulktransferTransactor // Write-only binding to the contract
	BulktransferFilterer   // Log filterer for contract events
}

// BulktransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type BulktransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BulktransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BulktransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BulktransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BulktransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BulktransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BulktransferSession struct {
	Contract     *Bulktransfer     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BulktransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BulktransferCallerSession struct {
	Contract *BulktransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BulktransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BulktransferTransactorSession struct {
	Contract     *BulktransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BulktransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type BulktransferRaw struct {
	Contract *Bulktransfer // Generic contract binding to access the raw methods on
}

// BulktransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BulktransferCallerRaw struct {
	Contract *BulktransferCaller // Generic read-only contract binding to access the raw methods on
}

// BulktransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BulktransferTransactorRaw struct {
	Contract *BulktransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBulktransfer creates a new instance of Bulktransfer, bound to a specific deployed contract.
func NewBulktransfer(address common.Address, backend bind.ContractBackend) (*Bulktransfer, error) {
	contract, err := bindBulktransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bulktransfer{BulktransferCaller: BulktransferCaller{contract: contract}, BulktransferTransactor: BulktransferTransactor{contract: contract}, BulktransferFilterer: BulktransferFilterer{contract: contract}}, nil
}

// NewBulktransferCaller creates a new read-only instance of Bulktransfer, bound to a specific deployed contract.
func NewBulktransferCaller(address common.Address, caller bind.ContractCaller) (*BulktransferCaller, error) {
	contract, err := bindBulktransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BulktransferCaller{contract: contract}, nil
}

// NewBulktransferTransactor creates a new write-only instance of Bulktransfer, bound to a specific deployed contract.
func NewBulktransferTransactor(address common.Address, transactor bind.ContractTransactor) (*BulktransferTransactor, error) {
	contract, err := bindBulktransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BulktransferTransactor{contract: contract}, nil
}

// NewBulktransferFilterer creates a new log filterer instance of Bulktransfer, bound to a specific deployed contract.
func NewBulktransferFilterer(address common.Address, filterer bind.ContractFilterer) (*BulktransferFilterer, error) {
	contract, err := bindBulktransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BulktransferFilterer{contract: contract}, nil
}

// bindBulktransfer binds a generic wrapper to an already deployed contract.
func bindBulktransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BulktransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bulktransfer *BulktransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bulktransfer.Contract.BulktransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bulktransfer *BulktransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bulktransfer.Contract.BulktransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bulktransfer *BulktransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bulktransfer.Contract.BulktransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bulktransfer *BulktransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bulktransfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bulktransfer *BulktransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bulktransfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bulktransfer *BulktransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bulktransfer.Contract.contract.Transact(opts, method, params...)
}

// BulkTransferTo is a paid mutator transaction binding the contract method 0x11305cec.
//
// Solidity: function bulkTransferTo(address nft, address to, uint256[] tokens) returns()
func (_Bulktransfer *BulktransferTransactor) BulkTransferTo(opts *bind.TransactOpts, nft common.Address, to common.Address, tokens []*big.Int) (*types.Transaction, error) {
	return _Bulktransfer.contract.Transact(opts, "bulkTransferTo", nft, to, tokens)
}

// BulkTransferTo is a paid mutator transaction binding the contract method 0x11305cec.
//
// Solidity: function bulkTransferTo(address nft, address to, uint256[] tokens) returns()
func (_Bulktransfer *BulktransferSession) BulkTransferTo(nft common.Address, to common.Address, tokens []*big.Int) (*types.Transaction, error) {
	return _Bulktransfer.Contract.BulkTransferTo(&_Bulktransfer.TransactOpts, nft, to, tokens)
}

// BulkTransferTo is a paid mutator transaction binding the contract method 0x11305cec.
//
// Solidity: function bulkTransferTo(address nft, address to, uint256[] tokens) returns()
func (_Bulktransfer *BulktransferTransactorSession) BulkTransferTo(nft common.Address, to common.Address, tokens []*big.Int) (*types.Transaction, error) {
	return _Bulktransfer.Contract.BulkTransferTo(&_Bulktransfer.TransactOpts, nft, to, tokens)
}
