// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fish

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

// FishMetaData contains all meta data concerning the Fish contract.
var FishMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"ConsecutiveTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"airdropBulkTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"airdropTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintAirdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintAirdropTeamMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ownerMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"publicMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_airdropper\",\"type\":\"address\"}],\"name\":\"setAirdropper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setBaseTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"setMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumFukushimaFishNFT.MintStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"setMintStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"name\":\"setPublicMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setReadMeURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"royaltyReceiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"basisPoints\",\"type\":\"uint96\"}],\"name\":\"setRoyaltyInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractSupplyController\",\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setSupplyController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setTermsOfServiceURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setUnrevealedTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setUpdateOnTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"name\":\"setWhitelistMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"}],\"name\":\"whitelistMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"contractSupplyController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getMintTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"minted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintStatus\",\"outputs\":[{\"internalType\":\"enumFukushimaFishNFT.MintStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_FILTER_REGISTRY\",\"outputs\":[{\"internalType\":\"contractIOperatorFilterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBLIC_MINT_COST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"readMeURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"termsOfServiceURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"path\",\"type\":\"uint256\"}],\"name\":\"validate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WHITELIST_MINT_COST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FishABI is the input ABI used to generate the binding from.
// Deprecated: Use FishMetaData.ABI instead.
var FishABI = FishMetaData.ABI

// Fish is an auto generated Go binding around an Ethereum contract.
type Fish struct {
	FishCaller     // Read-only binding to the contract
	FishTransactor // Write-only binding to the contract
	FishFilterer   // Log filterer for contract events
}

// FishCaller is an auto generated read-only Go binding around an Ethereum contract.
type FishCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FishTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FishTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FishFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FishFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FishSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FishSession struct {
	Contract     *Fish             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FishCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FishCallerSession struct {
	Contract *FishCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FishTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FishTransactorSession struct {
	Contract     *FishTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FishRaw is an auto generated low-level Go binding around an Ethereum contract.
type FishRaw struct {
	Contract *Fish // Generic contract binding to access the raw methods on
}

// FishCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FishCallerRaw struct {
	Contract *FishCaller // Generic read-only contract binding to access the raw methods on
}

// FishTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FishTransactorRaw struct {
	Contract *FishTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFish creates a new instance of Fish, bound to a specific deployed contract.
func NewFish(address common.Address, backend bind.ContractBackend) (*Fish, error) {
	contract, err := bindFish(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fish{FishCaller: FishCaller{contract: contract}, FishTransactor: FishTransactor{contract: contract}, FishFilterer: FishFilterer{contract: contract}}, nil
}

// NewFishCaller creates a new read-only instance of Fish, bound to a specific deployed contract.
func NewFishCaller(address common.Address, caller bind.ContractCaller) (*FishCaller, error) {
	contract, err := bindFish(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FishCaller{contract: contract}, nil
}

// NewFishTransactor creates a new write-only instance of Fish, bound to a specific deployed contract.
func NewFishTransactor(address common.Address, transactor bind.ContractTransactor) (*FishTransactor, error) {
	contract, err := bindFish(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FishTransactor{contract: contract}, nil
}

// NewFishFilterer creates a new log filterer instance of Fish, bound to a specific deployed contract.
func NewFishFilterer(address common.Address, filterer bind.ContractFilterer) (*FishFilterer, error) {
	contract, err := bindFish(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FishFilterer{contract: contract}, nil
}

// bindFish binds a generic wrapper to an already deployed contract.
func bindFish(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FishABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fish *FishRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fish.Contract.FishCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fish *FishRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fish.Contract.FishTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fish *FishRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fish.Contract.FishTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fish *FishCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fish.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fish *FishTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fish.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fish *FishTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fish.Contract.contract.Transact(opts, method, params...)
}

// OPERATORFILTERREGISTRY is a free data retrieval call binding the contract method 0x41f43434.
//
// Solidity: function OPERATOR_FILTER_REGISTRY() view returns(address)
func (_Fish *FishCaller) OPERATORFILTERREGISTRY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "OPERATOR_FILTER_REGISTRY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OPERATORFILTERREGISTRY is a free data retrieval call binding the contract method 0x41f43434.
//
// Solidity: function OPERATOR_FILTER_REGISTRY() view returns(address)
func (_Fish *FishSession) OPERATORFILTERREGISTRY() (common.Address, error) {
	return _Fish.Contract.OPERATORFILTERREGISTRY(&_Fish.CallOpts)
}

// OPERATORFILTERREGISTRY is a free data retrieval call binding the contract method 0x41f43434.
//
// Solidity: function OPERATOR_FILTER_REGISTRY() view returns(address)
func (_Fish *FishCallerSession) OPERATORFILTERREGISTRY() (common.Address, error) {
	return _Fish.Contract.OPERATORFILTERREGISTRY(&_Fish.CallOpts)
}

// PUBLICMINTCOST is a free data retrieval call binding the contract method 0xa2080c5b.
//
// Solidity: function PUBLIC_MINT_COST() view returns(uint256)
func (_Fish *FishCaller) PUBLICMINTCOST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "PUBLIC_MINT_COST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBLICMINTCOST is a free data retrieval call binding the contract method 0xa2080c5b.
//
// Solidity: function PUBLIC_MINT_COST() view returns(uint256)
func (_Fish *FishSession) PUBLICMINTCOST() (*big.Int, error) {
	return _Fish.Contract.PUBLICMINTCOST(&_Fish.CallOpts)
}

// PUBLICMINTCOST is a free data retrieval call binding the contract method 0xa2080c5b.
//
// Solidity: function PUBLIC_MINT_COST() view returns(uint256)
func (_Fish *FishCallerSession) PUBLICMINTCOST() (*big.Int, error) {
	return _Fish.Contract.PUBLICMINTCOST(&_Fish.CallOpts)
}

// WHITELISTMINTCOST is a free data retrieval call binding the contract method 0xe3d1a7be.
//
// Solidity: function WHITELIST_MINT_COST() view returns(uint256)
func (_Fish *FishCaller) WHITELISTMINTCOST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "WHITELIST_MINT_COST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WHITELISTMINTCOST is a free data retrieval call binding the contract method 0xe3d1a7be.
//
// Solidity: function WHITELIST_MINT_COST() view returns(uint256)
func (_Fish *FishSession) WHITELISTMINTCOST() (*big.Int, error) {
	return _Fish.Contract.WHITELISTMINTCOST(&_Fish.CallOpts)
}

// WHITELISTMINTCOST is a free data retrieval call binding the contract method 0xe3d1a7be.
//
// Solidity: function WHITELIST_MINT_COST() view returns(uint256)
func (_Fish *FishCallerSession) WHITELISTMINTCOST() (*big.Int, error) {
	return _Fish.Contract.WHITELISTMINTCOST(&_Fish.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Fish *FishCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Fish *FishSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Fish.Contract.BalanceOf(&_Fish.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Fish *FishCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Fish.Contract.BalanceOf(&_Fish.CallOpts, owner)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Fish *FishCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Fish *FishSession) Controller() (common.Address, error) {
	return _Fish.Contract.Controller(&_Fish.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Fish *FishCallerSession) Controller() (common.Address, error) {
	return _Fish.Contract.Controller(&_Fish.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_Fish *FishCaller) Exists(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "exists", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_Fish *FishSession) Exists(id *big.Int) (bool, error) {
	return _Fish.Contract.Exists(&_Fish.CallOpts, id)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_Fish *FishCallerSession) Exists(id *big.Int) (bool, error) {
	return _Fish.Contract.Exists(&_Fish.CallOpts, id)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Fish *FishCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Fish *FishSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Fish.Contract.GetApproved(&_Fish.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Fish *FishCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Fish.Contract.GetApproved(&_Fish.CallOpts, tokenId)
}

// GetMintTime is a free data retrieval call binding the contract method 0xc528cfc4.
//
// Solidity: function getMintTime(uint256 tokenId) view returns(uint256)
func (_Fish *FishCaller) GetMintTime(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "getMintTime", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMintTime is a free data retrieval call binding the contract method 0xc528cfc4.
//
// Solidity: function getMintTime(uint256 tokenId) view returns(uint256)
func (_Fish *FishSession) GetMintTime(tokenId *big.Int) (*big.Int, error) {
	return _Fish.Contract.GetMintTime(&_Fish.CallOpts, tokenId)
}

// GetMintTime is a free data retrieval call binding the contract method 0xc528cfc4.
//
// Solidity: function getMintTime(uint256 tokenId) view returns(uint256)
func (_Fish *FishCallerSession) GetMintTime(tokenId *big.Int) (*big.Int, error) {
	return _Fish.Contract.GetMintTime(&_Fish.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Fish *FishCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Fish *FishSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Fish.Contract.IsApprovedForAll(&_Fish.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Fish *FishCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Fish.Contract.IsApprovedForAll(&_Fish.CallOpts, owner, operator)
}

// MintStatus is a free data retrieval call binding the contract method 0x9da3f8fd.
//
// Solidity: function mintStatus() view returns(uint8)
func (_Fish *FishCaller) MintStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "mintStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MintStatus is a free data retrieval call binding the contract method 0x9da3f8fd.
//
// Solidity: function mintStatus() view returns(uint8)
func (_Fish *FishSession) MintStatus() (uint8, error) {
	return _Fish.Contract.MintStatus(&_Fish.CallOpts)
}

// MintStatus is a free data retrieval call binding the contract method 0x9da3f8fd.
//
// Solidity: function mintStatus() view returns(uint8)
func (_Fish *FishCallerSession) MintStatus() (uint8, error) {
	return _Fish.Contract.MintStatus(&_Fish.CallOpts)
}

// Minted is a free data retrieval call binding the contract method 0x1e7269c5.
//
// Solidity: function minted(address _addr) view returns(uint256)
func (_Fish *FishCaller) Minted(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "minted", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Minted is a free data retrieval call binding the contract method 0x1e7269c5.
//
// Solidity: function minted(address _addr) view returns(uint256)
func (_Fish *FishSession) Minted(_addr common.Address) (*big.Int, error) {
	return _Fish.Contract.Minted(&_Fish.CallOpts, _addr)
}

// Minted is a free data retrieval call binding the contract method 0x1e7269c5.
//
// Solidity: function minted(address _addr) view returns(uint256)
func (_Fish *FishCallerSession) Minted(_addr common.Address) (*big.Int, error) {
	return _Fish.Contract.Minted(&_Fish.CallOpts, _addr)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Fish *FishCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Fish *FishSession) Name() (string, error) {
	return _Fish.Contract.Name(&_Fish.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Fish *FishCallerSession) Name() (string, error) {
	return _Fish.Contract.Name(&_Fish.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fish *FishCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fish *FishSession) Owner() (common.Address, error) {
	return _Fish.Contract.Owner(&_Fish.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fish *FishCallerSession) Owner() (common.Address, error) {
	return _Fish.Contract.Owner(&_Fish.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Fish *FishCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Fish *FishSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Fish.Contract.OwnerOf(&_Fish.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Fish *FishCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Fish.Contract.OwnerOf(&_Fish.CallOpts, tokenId)
}

// ReadMeURI is a free data retrieval call binding the contract method 0x42b41aad.
//
// Solidity: function readMeURI() view returns(string)
func (_Fish *FishCaller) ReadMeURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "readMeURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReadMeURI is a free data retrieval call binding the contract method 0x42b41aad.
//
// Solidity: function readMeURI() view returns(string)
func (_Fish *FishSession) ReadMeURI() (string, error) {
	return _Fish.Contract.ReadMeURI(&_Fish.CallOpts)
}

// ReadMeURI is a free data retrieval call binding the contract method 0x42b41aad.
//
// Solidity: function readMeURI() view returns(string)
func (_Fish *FishCallerSession) ReadMeURI() (string, error) {
	return _Fish.Contract.ReadMeURI(&_Fish.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_Fish *FishCaller) RoyaltyInfo(opts *bind.CallOpts, _tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "royaltyInfo", _tokenId, _salePrice)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_Fish *FishSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _Fish.Contract.RoyaltyInfo(&_Fish.CallOpts, _tokenId, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_Fish *FishCallerSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _Fish.Contract.RoyaltyInfo(&_Fish.CallOpts, _tokenId, _salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Fish *FishCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Fish *FishSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Fish.Contract.SupportsInterface(&_Fish.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Fish *FishCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Fish.Contract.SupportsInterface(&_Fish.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Fish *FishCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Fish *FishSession) Symbol() (string, error) {
	return _Fish.Contract.Symbol(&_Fish.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Fish *FishCallerSession) Symbol() (string, error) {
	return _Fish.Contract.Symbol(&_Fish.CallOpts)
}

// TermsOfServiceURI is a free data retrieval call binding the contract method 0xe49b1a7c.
//
// Solidity: function termsOfServiceURI() view returns(string)
func (_Fish *FishCaller) TermsOfServiceURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "termsOfServiceURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TermsOfServiceURI is a free data retrieval call binding the contract method 0xe49b1a7c.
//
// Solidity: function termsOfServiceURI() view returns(string)
func (_Fish *FishSession) TermsOfServiceURI() (string, error) {
	return _Fish.Contract.TermsOfServiceURI(&_Fish.CallOpts)
}

// TermsOfServiceURI is a free data retrieval call binding the contract method 0xe49b1a7c.
//
// Solidity: function termsOfServiceURI() view returns(string)
func (_Fish *FishCallerSession) TermsOfServiceURI() (string, error) {
	return _Fish.Contract.TermsOfServiceURI(&_Fish.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_Fish *FishCaller) TokenURI(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "tokenURI", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_Fish *FishSession) TokenURI(id *big.Int) (string, error) {
	return _Fish.Contract.TokenURI(&_Fish.CallOpts, id)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_Fish *FishCallerSession) TokenURI(id *big.Int) (string, error) {
	return _Fish.Contract.TokenURI(&_Fish.CallOpts, id)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Fish *FishCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Fish *FishSession) TotalSupply() (*big.Int, error) {
	return _Fish.Contract.TotalSupply(&_Fish.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Fish *FishCallerSession) TotalSupply() (*big.Int, error) {
	return _Fish.Contract.TotalSupply(&_Fish.CallOpts)
}

// Validate is a free data retrieval call binding the contract method 0x7e386b65.
//
// Solidity: function validate(address addr, uint256 limit, bytes32[] proof, uint256 path) view returns(bool)
func (_Fish *FishCaller) Validate(opts *bind.CallOpts, addr common.Address, limit *big.Int, proof [][32]byte, path *big.Int) (bool, error) {
	var out []interface{}
	err := _Fish.contract.Call(opts, &out, "validate", addr, limit, proof, path)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Validate is a free data retrieval call binding the contract method 0x7e386b65.
//
// Solidity: function validate(address addr, uint256 limit, bytes32[] proof, uint256 path) view returns(bool)
func (_Fish *FishSession) Validate(addr common.Address, limit *big.Int, proof [][32]byte, path *big.Int) (bool, error) {
	return _Fish.Contract.Validate(&_Fish.CallOpts, addr, limit, proof, path)
}

// Validate is a free data retrieval call binding the contract method 0x7e386b65.
//
// Solidity: function validate(address addr, uint256 limit, bytes32[] proof, uint256 path) view returns(bool)
func (_Fish *FishCallerSession) Validate(addr common.Address, limit *big.Int, proof [][32]byte, path *big.Int) (bool, error) {
	return _Fish.Contract.Validate(&_Fish.CallOpts, addr, limit, proof, path)
}

// AirdropBulkTransfer is a paid mutator transaction binding the contract method 0x17ba3eff.
//
// Solidity: function airdropBulkTransfer(address to, uint256[] tokenIds) returns()
func (_Fish *FishTransactor) AirdropBulkTransfer(opts *bind.TransactOpts, to common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "airdropBulkTransfer", to, tokenIds)
}

// AirdropBulkTransfer is a paid mutator transaction binding the contract method 0x17ba3eff.
//
// Solidity: function airdropBulkTransfer(address to, uint256[] tokenIds) returns()
func (_Fish *FishSession) AirdropBulkTransfer(to common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Fish.Contract.AirdropBulkTransfer(&_Fish.TransactOpts, to, tokenIds)
}

// AirdropBulkTransfer is a paid mutator transaction binding the contract method 0x17ba3eff.
//
// Solidity: function airdropBulkTransfer(address to, uint256[] tokenIds) returns()
func (_Fish *FishTransactorSession) AirdropBulkTransfer(to common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Fish.Contract.AirdropBulkTransfer(&_Fish.TransactOpts, to, tokenIds)
}

// AirdropTransfer is a paid mutator transaction binding the contract method 0x8c7ec0f1.
//
// Solidity: function airdropTransfer(address to, uint256 tokenId) returns()
func (_Fish *FishTransactor) AirdropTransfer(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "airdropTransfer", to, tokenId)
}

// AirdropTransfer is a paid mutator transaction binding the contract method 0x8c7ec0f1.
//
// Solidity: function airdropTransfer(address to, uint256 tokenId) returns()
func (_Fish *FishSession) AirdropTransfer(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.AirdropTransfer(&_Fish.TransactOpts, to, tokenId)
}

// AirdropTransfer is a paid mutator transaction binding the contract method 0x8c7ec0f1.
//
// Solidity: function airdropTransfer(address to, uint256 tokenId) returns()
func (_Fish *FishTransactorSession) AirdropTransfer(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.AirdropTransfer(&_Fish.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) payable returns()
func (_Fish *FishTransactor) Approve(opts *bind.TransactOpts, operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "approve", operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) payable returns()
func (_Fish *FishSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.Approve(&_Fish.TransactOpts, operator, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address operator, uint256 tokenId) payable returns()
func (_Fish *FishTransactorSession) Approve(operator common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.Approve(&_Fish.TransactOpts, operator, tokenId)
}

// MintAirdrop is a paid mutator transaction binding the contract method 0xcb283fef.
//
// Solidity: function mintAirdrop() returns()
func (_Fish *FishTransactor) MintAirdrop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "mintAirdrop")
}

// MintAirdrop is a paid mutator transaction binding the contract method 0xcb283fef.
//
// Solidity: function mintAirdrop() returns()
func (_Fish *FishSession) MintAirdrop() (*types.Transaction, error) {
	return _Fish.Contract.MintAirdrop(&_Fish.TransactOpts)
}

// MintAirdrop is a paid mutator transaction binding the contract method 0xcb283fef.
//
// Solidity: function mintAirdrop() returns()
func (_Fish *FishTransactorSession) MintAirdrop() (*types.Transaction, error) {
	return _Fish.Contract.MintAirdrop(&_Fish.TransactOpts)
}

// MintAirdropTeamMint is a paid mutator transaction binding the contract method 0xc447e72d.
//
// Solidity: function mintAirdropTeamMint(address to, uint256 amount) returns()
func (_Fish *FishTransactor) MintAirdropTeamMint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "mintAirdropTeamMint", to, amount)
}

// MintAirdropTeamMint is a paid mutator transaction binding the contract method 0xc447e72d.
//
// Solidity: function mintAirdropTeamMint(address to, uint256 amount) returns()
func (_Fish *FishSession) MintAirdropTeamMint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.MintAirdropTeamMint(&_Fish.TransactOpts, to, amount)
}

// MintAirdropTeamMint is a paid mutator transaction binding the contract method 0xc447e72d.
//
// Solidity: function mintAirdropTeamMint(address to, uint256 amount) returns()
func (_Fish *FishTransactorSession) MintAirdropTeamMint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.MintAirdropTeamMint(&_Fish.TransactOpts, to, amount)
}

// OwnerMint is a paid mutator transaction binding the contract method 0x484b973c.
//
// Solidity: function ownerMint(address to, uint256 amount) returns()
func (_Fish *FishTransactor) OwnerMint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "ownerMint", to, amount)
}

// OwnerMint is a paid mutator transaction binding the contract method 0x484b973c.
//
// Solidity: function ownerMint(address to, uint256 amount) returns()
func (_Fish *FishSession) OwnerMint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.OwnerMint(&_Fish.TransactOpts, to, amount)
}

// OwnerMint is a paid mutator transaction binding the contract method 0x484b973c.
//
// Solidity: function ownerMint(address to, uint256 amount) returns()
func (_Fish *FishTransactorSession) OwnerMint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.OwnerMint(&_Fish.TransactOpts, to, amount)
}

// PublicMint is a paid mutator transaction binding the contract method 0x2db11544.
//
// Solidity: function publicMint(uint256 amount) payable returns()
func (_Fish *FishTransactor) PublicMint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "publicMint", amount)
}

// PublicMint is a paid mutator transaction binding the contract method 0x2db11544.
//
// Solidity: function publicMint(uint256 amount) payable returns()
func (_Fish *FishSession) PublicMint(amount *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.PublicMint(&_Fish.TransactOpts, amount)
}

// PublicMint is a paid mutator transaction binding the contract method 0x2db11544.
//
// Solidity: function publicMint(uint256 amount) payable returns()
func (_Fish *FishTransactorSession) PublicMint(amount *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.PublicMint(&_Fish.TransactOpts, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Fish *FishTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Fish *FishSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.SafeTransferFrom(&_Fish.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Fish *FishTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.SafeTransferFrom(&_Fish.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) payable returns()
func (_Fish *FishTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) payable returns()
func (_Fish *FishSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Fish.Contract.SafeTransferFrom0(&_Fish.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) payable returns()
func (_Fish *FishTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Fish.Contract.SafeTransferFrom0(&_Fish.TransactOpts, from, to, tokenId, data)
}

// SetAirdropper is a paid mutator transaction binding the contract method 0xdc489c9b.
//
// Solidity: function setAirdropper(address _airdropper) returns()
func (_Fish *FishTransactor) SetAirdropper(opts *bind.TransactOpts, _airdropper common.Address) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "setAirdropper", _airdropper)
}

// SetAirdropper is a paid mutator transaction binding the contract method 0xdc489c9b.
//
// Solidity: function setAirdropper(address _airdropper) returns()
func (_Fish *FishSession) SetAirdropper(_airdropper common.Address) (*types.Transaction, error) {
	return _Fish.Contract.SetAirdropper(&_Fish.TransactOpts, _airdropper)
}

// SetAirdropper is a paid mutator transaction binding the contract method 0xdc489c9b.
//
// Solidity: function setAirdropper(address _airdropper) returns()
func (_Fish *FishTransactorSession) SetAirdropper(_airdropper common.Address) (*types.Transaction, error) {
	return _Fish.Contract.SetAirdropper(&_Fish.TransactOpts, _airdropper)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Fish *FishTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Fish *FishSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Fish.Contract.SetApprovalForAll(&_Fish.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Fish *FishTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Fish.Contract.SetApprovalForAll(&_Fish.TransactOpts, operator, approved)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string uri) returns()
func (_Fish *FishTransactor) SetBaseTokenURI(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "setBaseTokenURI", uri)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string uri) returns()
func (_Fish *FishSession) SetBaseTokenURI(uri string) (*types.Transaction, error) {
	return _Fish.Contract.SetBaseTokenURI(&_Fish.TransactOpts, uri)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string uri) returns()
func (_Fish *FishTransactorSession) SetBaseTokenURI(uri string) (*types.Transaction, error) {
	return _Fish.Contract.SetBaseTokenURI(&_Fish.TransactOpts, uri)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 merkleRoot) returns()
func (_Fish *FishTransactor) SetMerkleRoot(opts *bind.TransactOpts, merkleRoot [32]byte) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "setMerkleRoot", merkleRoot)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 merkleRoot) returns()
func (_Fish *FishSession) SetMerkleRoot(merkleRoot [32]byte) (*types.Transaction, error) {
	return _Fish.Contract.SetMerkleRoot(&_Fish.TransactOpts, merkleRoot)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 merkleRoot) returns()
func (_Fish *FishTransactorSession) SetMerkleRoot(merkleRoot [32]byte) (*types.Transaction, error) {
	return _Fish.Contract.SetMerkleRoot(&_Fish.TransactOpts, merkleRoot)
}

// SetMintStatus is a paid mutator transaction binding the contract method 0x814c8c55.
//
// Solidity: function setMintStatus(uint8 status) returns()
func (_Fish *FishTransactor) SetMintStatus(opts *bind.TransactOpts, status uint8) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "setMintStatus", status)
}

// SetMintStatus is a paid mutator transaction binding the contract method 0x814c8c55.
//
// Solidity: function setMintStatus(uint8 status) returns()
func (_Fish *FishSession) SetMintStatus(status uint8) (*types.Transaction, error) {
	return _Fish.Contract.SetMintStatus(&_Fish.TransactOpts, status)
}

// SetMintStatus is a paid mutator transaction binding the contract method 0x814c8c55.
//
// Solidity: function setMintStatus(uint8 status) returns()
func (_Fish *FishTransactorSession) SetMintStatus(status uint8) (*types.Transaction, error) {
	return _Fish.Contract.SetMintStatus(&_Fish.TransactOpts, status)
}

// SetPublicMintPrice is a paid mutator transaction binding the contract method 0x5d82cf6e.
//
// Solidity: function setPublicMintPrice(uint256 cost) returns()
func (_Fish *FishTransactor) SetPublicMintPrice(opts *bind.TransactOpts, cost *big.Int) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "setPublicMintPrice", cost)
}

// SetPublicMintPrice is a paid mutator transaction binding the contract method 0x5d82cf6e.
//
// Solidity: function setPublicMintPrice(uint256 cost) returns()
func (_Fish *FishSession) SetPublicMintPrice(cost *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.SetPublicMintPrice(&_Fish.TransactOpts, cost)
}

// SetPublicMintPrice is a paid mutator transaction binding the contract method 0x5d82cf6e.
//
// Solidity: function setPublicMintPrice(uint256 cost) returns()
func (_Fish *FishTransactorSession) SetPublicMintPrice(cost *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.SetPublicMintPrice(&_Fish.TransactOpts, cost)
}

// SetReadMeURI is a paid mutator transaction binding the contract method 0x4ffecb49.
//
// Solidity: function setReadMeURI(string uri) returns()
func (_Fish *FishTransactor) SetReadMeURI(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "setReadMeURI", uri)
}

// SetReadMeURI is a paid mutator transaction binding the contract method 0x4ffecb49.
//
// Solidity: function setReadMeURI(string uri) returns()
func (_Fish *FishSession) SetReadMeURI(uri string) (*types.Transaction, error) {
	return _Fish.Contract.SetReadMeURI(&_Fish.TransactOpts, uri)
}

// SetReadMeURI is a paid mutator transaction binding the contract method 0x4ffecb49.
//
// Solidity: function setReadMeURI(string uri) returns()
func (_Fish *FishTransactorSession) SetReadMeURI(uri string) (*types.Transaction, error) {
	return _Fish.Contract.SetReadMeURI(&_Fish.TransactOpts, uri)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x02fa7c47.
//
// Solidity: function setRoyaltyInfo(address royaltyReceiver, uint96 basisPoints) returns()
func (_Fish *FishTransactor) SetRoyaltyInfo(opts *bind.TransactOpts, royaltyReceiver common.Address, basisPoints *big.Int) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "setRoyaltyInfo", royaltyReceiver, basisPoints)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x02fa7c47.
//
// Solidity: function setRoyaltyInfo(address royaltyReceiver, uint96 basisPoints) returns()
func (_Fish *FishSession) SetRoyaltyInfo(royaltyReceiver common.Address, basisPoints *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.SetRoyaltyInfo(&_Fish.TransactOpts, royaltyReceiver, basisPoints)
}

// SetRoyaltyInfo is a paid mutator transaction binding the contract method 0x02fa7c47.
//
// Solidity: function setRoyaltyInfo(address royaltyReceiver, uint96 basisPoints) returns()
func (_Fish *FishTransactorSession) SetRoyaltyInfo(royaltyReceiver common.Address, basisPoints *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.SetRoyaltyInfo(&_Fish.TransactOpts, royaltyReceiver, basisPoints)
}

// SetSupplyController is a paid mutator transaction binding the contract method 0x52875bc3.
//
// Solidity: function setSupplyController(address _controller) returns()
func (_Fish *FishTransactor) SetSupplyController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "setSupplyController", _controller)
}

// SetSupplyController is a paid mutator transaction binding the contract method 0x52875bc3.
//
// Solidity: function setSupplyController(address _controller) returns()
func (_Fish *FishSession) SetSupplyController(_controller common.Address) (*types.Transaction, error) {
	return _Fish.Contract.SetSupplyController(&_Fish.TransactOpts, _controller)
}

// SetSupplyController is a paid mutator transaction binding the contract method 0x52875bc3.
//
// Solidity: function setSupplyController(address _controller) returns()
func (_Fish *FishTransactorSession) SetSupplyController(_controller common.Address) (*types.Transaction, error) {
	return _Fish.Contract.SetSupplyController(&_Fish.TransactOpts, _controller)
}

// SetTermsOfServiceURI is a paid mutator transaction binding the contract method 0x1babc84a.
//
// Solidity: function setTermsOfServiceURI(string uri) returns()
func (_Fish *FishTransactor) SetTermsOfServiceURI(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "setTermsOfServiceURI", uri)
}

// SetTermsOfServiceURI is a paid mutator transaction binding the contract method 0x1babc84a.
//
// Solidity: function setTermsOfServiceURI(string uri) returns()
func (_Fish *FishSession) SetTermsOfServiceURI(uri string) (*types.Transaction, error) {
	return _Fish.Contract.SetTermsOfServiceURI(&_Fish.TransactOpts, uri)
}

// SetTermsOfServiceURI is a paid mutator transaction binding the contract method 0x1babc84a.
//
// Solidity: function setTermsOfServiceURI(string uri) returns()
func (_Fish *FishTransactorSession) SetTermsOfServiceURI(uri string) (*types.Transaction, error) {
	return _Fish.Contract.SetTermsOfServiceURI(&_Fish.TransactOpts, uri)
}

// SetUnrevealedTokenURI is a paid mutator transaction binding the contract method 0x820de0c5.
//
// Solidity: function setUnrevealedTokenURI(string uri) returns()
func (_Fish *FishTransactor) SetUnrevealedTokenURI(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "setUnrevealedTokenURI", uri)
}

// SetUnrevealedTokenURI is a paid mutator transaction binding the contract method 0x820de0c5.
//
// Solidity: function setUnrevealedTokenURI(string uri) returns()
func (_Fish *FishSession) SetUnrevealedTokenURI(uri string) (*types.Transaction, error) {
	return _Fish.Contract.SetUnrevealedTokenURI(&_Fish.TransactOpts, uri)
}

// SetUnrevealedTokenURI is a paid mutator transaction binding the contract method 0x820de0c5.
//
// Solidity: function setUnrevealedTokenURI(string uri) returns()
func (_Fish *FishTransactorSession) SetUnrevealedTokenURI(uri string) (*types.Transaction, error) {
	return _Fish.Contract.SetUnrevealedTokenURI(&_Fish.TransactOpts, uri)
}

// SetUpdateOnTransfer is a paid mutator transaction binding the contract method 0x8ed6e589.
//
// Solidity: function setUpdateOnTransfer(bool status) returns()
func (_Fish *FishTransactor) SetUpdateOnTransfer(opts *bind.TransactOpts, status bool) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "setUpdateOnTransfer", status)
}

// SetUpdateOnTransfer is a paid mutator transaction binding the contract method 0x8ed6e589.
//
// Solidity: function setUpdateOnTransfer(bool status) returns()
func (_Fish *FishSession) SetUpdateOnTransfer(status bool) (*types.Transaction, error) {
	return _Fish.Contract.SetUpdateOnTransfer(&_Fish.TransactOpts, status)
}

// SetUpdateOnTransfer is a paid mutator transaction binding the contract method 0x8ed6e589.
//
// Solidity: function setUpdateOnTransfer(bool status) returns()
func (_Fish *FishTransactorSession) SetUpdateOnTransfer(status bool) (*types.Transaction, error) {
	return _Fish.Contract.SetUpdateOnTransfer(&_Fish.TransactOpts, status)
}

// SetWhitelistMintPrice is a paid mutator transaction binding the contract method 0xa611708e.
//
// Solidity: function setWhitelistMintPrice(uint256 cost) returns()
func (_Fish *FishTransactor) SetWhitelistMintPrice(opts *bind.TransactOpts, cost *big.Int) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "setWhitelistMintPrice", cost)
}

// SetWhitelistMintPrice is a paid mutator transaction binding the contract method 0xa611708e.
//
// Solidity: function setWhitelistMintPrice(uint256 cost) returns()
func (_Fish *FishSession) SetWhitelistMintPrice(cost *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.SetWhitelistMintPrice(&_Fish.TransactOpts, cost)
}

// SetWhitelistMintPrice is a paid mutator transaction binding the contract method 0xa611708e.
//
// Solidity: function setWhitelistMintPrice(uint256 cost) returns()
func (_Fish *FishTransactorSession) SetWhitelistMintPrice(cost *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.SetWhitelistMintPrice(&_Fish.TransactOpts, cost)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Fish *FishTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Fish *FishSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.TransferFrom(&_Fish.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) payable returns()
func (_Fish *FishTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.TransferFrom(&_Fish.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Fish *FishTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Fish *FishSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Fish.Contract.TransferOwnership(&_Fish.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Fish *FishTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Fish.Contract.TransferOwnership(&_Fish.TransactOpts, newOwner)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0x1c4072e2.
//
// Solidity: function whitelistMint(uint256 amount, uint256 limit, bytes32[] proof, uint256 path) payable returns()
func (_Fish *FishTransactor) WhitelistMint(opts *bind.TransactOpts, amount *big.Int, limit *big.Int, proof [][32]byte, path *big.Int) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "whitelistMint", amount, limit, proof, path)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0x1c4072e2.
//
// Solidity: function whitelistMint(uint256 amount, uint256 limit, bytes32[] proof, uint256 path) payable returns()
func (_Fish *FishSession) WhitelistMint(amount *big.Int, limit *big.Int, proof [][32]byte, path *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.WhitelistMint(&_Fish.TransactOpts, amount, limit, proof, path)
}

// WhitelistMint is a paid mutator transaction binding the contract method 0x1c4072e2.
//
// Solidity: function whitelistMint(uint256 amount, uint256 limit, bytes32[] proof, uint256 path) payable returns()
func (_Fish *FishTransactorSession) WhitelistMint(amount *big.Int, limit *big.Int, proof [][32]byte, path *big.Int) (*types.Transaction, error) {
	return _Fish.Contract.WhitelistMint(&_Fish.TransactOpts, amount, limit, proof, path)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_Fish *FishTransactor) Withdraw(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Fish.contract.Transact(opts, "withdraw", to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_Fish *FishSession) Withdraw(to common.Address) (*types.Transaction, error) {
	return _Fish.Contract.Withdraw(&_Fish.TransactOpts, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_Fish *FishTransactorSession) Withdraw(to common.Address) (*types.Transaction, error) {
	return _Fish.Contract.Withdraw(&_Fish.TransactOpts, to)
}

// FishApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Fish contract.
type FishApprovalIterator struct {
	Event *FishApproval // Event containing the contract specifics and raw log

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
func (it *FishApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FishApproval)
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
		it.Event = new(FishApproval)
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
func (it *FishApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FishApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FishApproval represents a Approval event raised by the Fish contract.
type FishApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Fish *FishFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*FishApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Fish.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FishApprovalIterator{contract: _Fish.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Fish *FishFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FishApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Fish.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FishApproval)
				if err := _Fish.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Fish *FishFilterer) ParseApproval(log types.Log) (*FishApproval, error) {
	event := new(FishApproval)
	if err := _Fish.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FishApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Fish contract.
type FishApprovalForAllIterator struct {
	Event *FishApprovalForAll // Event containing the contract specifics and raw log

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
func (it *FishApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FishApprovalForAll)
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
		it.Event = new(FishApprovalForAll)
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
func (it *FishApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FishApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FishApprovalForAll represents a ApprovalForAll event raised by the Fish contract.
type FishApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Fish *FishFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*FishApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Fish.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &FishApprovalForAllIterator{contract: _Fish.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Fish *FishFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *FishApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Fish.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FishApprovalForAll)
				if err := _Fish.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Fish *FishFilterer) ParseApprovalForAll(log types.Log) (*FishApprovalForAll, error) {
	event := new(FishApprovalForAll)
	if err := _Fish.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FishConsecutiveTransferIterator is returned from FilterConsecutiveTransfer and is used to iterate over the raw logs and unpacked data for ConsecutiveTransfer events raised by the Fish contract.
type FishConsecutiveTransferIterator struct {
	Event *FishConsecutiveTransfer // Event containing the contract specifics and raw log

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
func (it *FishConsecutiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FishConsecutiveTransfer)
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
		it.Event = new(FishConsecutiveTransfer)
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
func (it *FishConsecutiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FishConsecutiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FishConsecutiveTransfer represents a ConsecutiveTransfer event raised by the Fish contract.
type FishConsecutiveTransfer struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	From        common.Address
	To          common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsecutiveTransfer is a free log retrieval operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Fish *FishFilterer) FilterConsecutiveTransfer(opts *bind.FilterOpts, fromTokenId []*big.Int, from []common.Address, to []common.Address) (*FishConsecutiveTransferIterator, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Fish.contract.FilterLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FishConsecutiveTransferIterator{contract: _Fish.contract, event: "ConsecutiveTransfer", logs: logs, sub: sub}, nil
}

// WatchConsecutiveTransfer is a free log subscription operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Fish *FishFilterer) WatchConsecutiveTransfer(opts *bind.WatchOpts, sink chan<- *FishConsecutiveTransfer, fromTokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromTokenIdRule []interface{}
	for _, fromTokenIdItem := range fromTokenId {
		fromTokenIdRule = append(fromTokenIdRule, fromTokenIdItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Fish.contract.WatchLogs(opts, "ConsecutiveTransfer", fromTokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FishConsecutiveTransfer)
				if err := _Fish.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
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

// ParseConsecutiveTransfer is a log parse operation binding the contract event 0xdeaa91b6123d068f5821d0fb0678463d1a8a6079fe8af5de3ce5e896dcf9133d.
//
// Solidity: event ConsecutiveTransfer(uint256 indexed fromTokenId, uint256 toTokenId, address indexed from, address indexed to)
func (_Fish *FishFilterer) ParseConsecutiveTransfer(log types.Log) (*FishConsecutiveTransfer, error) {
	event := new(FishConsecutiveTransfer)
	if err := _Fish.contract.UnpackLog(event, "ConsecutiveTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FishOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Fish contract.
type FishOwnershipTransferredIterator struct {
	Event *FishOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FishOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FishOwnershipTransferred)
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
		it.Event = new(FishOwnershipTransferred)
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
func (it *FishOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FishOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FishOwnershipTransferred represents a OwnershipTransferred event raised by the Fish contract.
type FishOwnershipTransferred struct {
	User     common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_Fish *FishFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, user []common.Address, newOwner []common.Address) (*FishOwnershipTransferredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Fish.contract.FilterLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FishOwnershipTransferredIterator{contract: _Fish.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_Fish *FishFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FishOwnershipTransferred, user []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Fish.contract.WatchLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FishOwnershipTransferred)
				if err := _Fish.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_Fish *FishFilterer) ParseOwnershipTransferred(log types.Log) (*FishOwnershipTransferred, error) {
	event := new(FishOwnershipTransferred)
	if err := _Fish.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FishTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Fish contract.
type FishTransferIterator struct {
	Event *FishTransfer // Event containing the contract specifics and raw log

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
func (it *FishTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FishTransfer)
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
		it.Event = new(FishTransfer)
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
func (it *FishTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FishTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FishTransfer represents a Transfer event raised by the Fish contract.
type FishTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Fish *FishFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*FishTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Fish.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &FishTransferIterator{contract: _Fish.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Fish *FishFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FishTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Fish.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FishTransfer)
				if err := _Fish.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Fish *FishFilterer) ParseTransfer(log types.Log) (*FishTransfer, error) {
	event := new(FishTransfer)
	if err := _Fish.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
