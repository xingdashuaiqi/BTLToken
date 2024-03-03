// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// BTLTokenMetaData contains all meta data concerning the BTLToken contract.
var BTLTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_lpwallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_NFTwallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdt\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NFTwallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Open\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_swapRouter\",\"outputs\":[{\"internalType\":\"contractISwapRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addwhiteList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"blackList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyburn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buylpfee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buynftfee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buytotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpwallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nowstatus\",\"outputs\":[{\"internalType\":\"enumBLT.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingburn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setadmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_blackList\",\"type\":\"address\"}],\"name\":\"setblacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteListLPamount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BTLTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use BTLTokenMetaData.ABI instead.
var BTLTokenABI = BTLTokenMetaData.ABI

// BTLToken is an auto generated Go binding around an Ethereum contract.
type BTLToken struct {
	BTLTokenCaller     // Read-only binding to the contract
	BTLTokenTransactor // Write-only binding to the contract
	BTLTokenFilterer   // Log filterer for contract events
}

// BTLTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BTLTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTLTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BTLTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTLTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BTLTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTLTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BTLTokenSession struct {
	Contract     *BTLToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BTLTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BTLTokenCallerSession struct {
	Contract *BTLTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BTLTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BTLTokenTransactorSession struct {
	Contract     *BTLTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BTLTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BTLTokenRaw struct {
	Contract *BTLToken // Generic contract binding to access the raw methods on
}

// BTLTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BTLTokenCallerRaw struct {
	Contract *BTLTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BTLTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BTLTokenTransactorRaw struct {
	Contract *BTLTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBTLToken creates a new instance of BTLToken, bound to a specific deployed contract.
func NewBTLToken(address common.Address, backend bind.ContractBackend) (*BTLToken, error) {
	contract, err := bindBTLToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BTLToken{BTLTokenCaller: BTLTokenCaller{contract: contract}, BTLTokenTransactor: BTLTokenTransactor{contract: contract}, BTLTokenFilterer: BTLTokenFilterer{contract: contract}}, nil
}

// NewBTLTokenCaller creates a new read-only instance of BTLToken, bound to a specific deployed contract.
func NewBTLTokenCaller(address common.Address, caller bind.ContractCaller) (*BTLTokenCaller, error) {
	contract, err := bindBTLToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BTLTokenCaller{contract: contract}, nil
}

// NewBTLTokenTransactor creates a new write-only instance of BTLToken, bound to a specific deployed contract.
func NewBTLTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BTLTokenTransactor, error) {
	contract, err := bindBTLToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BTLTokenTransactor{contract: contract}, nil
}

// NewBTLTokenFilterer creates a new log filterer instance of BTLToken, bound to a specific deployed contract.
func NewBTLTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BTLTokenFilterer, error) {
	contract, err := bindBTLToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BTLTokenFilterer{contract: contract}, nil
}

// bindBTLToken binds a generic wrapper to an already deployed contract.
func bindBTLToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BTLTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BTLToken *BTLTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BTLToken.Contract.BTLTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BTLToken *BTLTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTLToken.Contract.BTLTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BTLToken *BTLTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BTLToken.Contract.BTLTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BTLToken *BTLTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BTLToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BTLToken *BTLTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTLToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BTLToken *BTLTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BTLToken.Contract.contract.Transact(opts, method, params...)
}

// NFTwallet is a free data retrieval call binding the contract method 0x0c26877d.
//
// Solidity: function NFTwallet() view returns(address)
func (_BTLToken *BTLTokenCaller) NFTwallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "NFTwallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NFTwallet is a free data retrieval call binding the contract method 0x0c26877d.
//
// Solidity: function NFTwallet() view returns(address)
func (_BTLToken *BTLTokenSession) NFTwallet() (common.Address, error) {
	return _BTLToken.Contract.NFTwallet(&_BTLToken.CallOpts)
}

// NFTwallet is a free data retrieval call binding the contract method 0x0c26877d.
//
// Solidity: function NFTwallet() view returns(address)
func (_BTLToken *BTLTokenCallerSession) NFTwallet() (common.Address, error) {
	return _BTLToken.Contract.NFTwallet(&_BTLToken.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0x206c17bb.
//
// Solidity: function _swapRouter() view returns(address)
func (_BTLToken *BTLTokenCaller) SwapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "_swapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapRouter is a free data retrieval call binding the contract method 0x206c17bb.
//
// Solidity: function _swapRouter() view returns(address)
func (_BTLToken *BTLTokenSession) SwapRouter() (common.Address, error) {
	return _BTLToken.Contract.SwapRouter(&_BTLToken.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0x206c17bb.
//
// Solidity: function _swapRouter() view returns(address)
func (_BTLToken *BTLTokenCallerSession) SwapRouter() (common.Address, error) {
	return _BTLToken.Contract.SwapRouter(&_BTLToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_BTLToken *BTLTokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_BTLToken *BTLTokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _BTLToken.Contract.Allowance(&_BTLToken.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_BTLToken *BTLTokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _BTLToken.Contract.Allowance(&_BTLToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _address) view returns(uint256 amount)
func (_BTLToken *BTLTokenCaller) BalanceOf(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "balanceOf", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _address) view returns(uint256 amount)
func (_BTLToken *BTLTokenSession) BalanceOf(_address common.Address) (*big.Int, error) {
	return _BTLToken.Contract.BalanceOf(&_BTLToken.CallOpts, _address)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _address) view returns(uint256 amount)
func (_BTLToken *BTLTokenCallerSession) BalanceOf(_address common.Address) (*big.Int, error) {
	return _BTLToken.Contract.BalanceOf(&_BTLToken.CallOpts, _address)
}

// BlackList is a free data retrieval call binding the contract method 0x4838d165.
//
// Solidity: function blackList(address ) view returns(bool)
func (_BTLToken *BTLTokenCaller) BlackList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "blackList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlackList is a free data retrieval call binding the contract method 0x4838d165.
//
// Solidity: function blackList(address ) view returns(bool)
func (_BTLToken *BTLTokenSession) BlackList(arg0 common.Address) (bool, error) {
	return _BTLToken.Contract.BlackList(&_BTLToken.CallOpts, arg0)
}

// BlackList is a free data retrieval call binding the contract method 0x4838d165.
//
// Solidity: function blackList(address ) view returns(bool)
func (_BTLToken *BTLTokenCallerSession) BlackList(arg0 common.Address) (bool, error) {
	return _BTLToken.Contract.BlackList(&_BTLToken.CallOpts, arg0)
}

// Buyburn is a free data retrieval call binding the contract method 0x0490b9ea.
//
// Solidity: function buyburn() view returns(uint256)
func (_BTLToken *BTLTokenCaller) Buyburn(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "buyburn")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Buyburn is a free data retrieval call binding the contract method 0x0490b9ea.
//
// Solidity: function buyburn() view returns(uint256)
func (_BTLToken *BTLTokenSession) Buyburn() (*big.Int, error) {
	return _BTLToken.Contract.Buyburn(&_BTLToken.CallOpts)
}

// Buyburn is a free data retrieval call binding the contract method 0x0490b9ea.
//
// Solidity: function buyburn() view returns(uint256)
func (_BTLToken *BTLTokenCallerSession) Buyburn() (*big.Int, error) {
	return _BTLToken.Contract.Buyburn(&_BTLToken.CallOpts)
}

// Buylpfee is a free data retrieval call binding the contract method 0xf94b52ae.
//
// Solidity: function buylpfee() view returns(uint256)
func (_BTLToken *BTLTokenCaller) Buylpfee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "buylpfee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Buylpfee is a free data retrieval call binding the contract method 0xf94b52ae.
//
// Solidity: function buylpfee() view returns(uint256)
func (_BTLToken *BTLTokenSession) Buylpfee() (*big.Int, error) {
	return _BTLToken.Contract.Buylpfee(&_BTLToken.CallOpts)
}

// Buylpfee is a free data retrieval call binding the contract method 0xf94b52ae.
//
// Solidity: function buylpfee() view returns(uint256)
func (_BTLToken *BTLTokenCallerSession) Buylpfee() (*big.Int, error) {
	return _BTLToken.Contract.Buylpfee(&_BTLToken.CallOpts)
}

// Buynftfee is a free data retrieval call binding the contract method 0xf6e9436a.
//
// Solidity: function buynftfee() view returns(uint256)
func (_BTLToken *BTLTokenCaller) Buynftfee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "buynftfee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Buynftfee is a free data retrieval call binding the contract method 0xf6e9436a.
//
// Solidity: function buynftfee() view returns(uint256)
func (_BTLToken *BTLTokenSession) Buynftfee() (*big.Int, error) {
	return _BTLToken.Contract.Buynftfee(&_BTLToken.CallOpts)
}

// Buynftfee is a free data retrieval call binding the contract method 0xf6e9436a.
//
// Solidity: function buynftfee() view returns(uint256)
func (_BTLToken *BTLTokenCallerSession) Buynftfee() (*big.Int, error) {
	return _BTLToken.Contract.Buynftfee(&_BTLToken.CallOpts)
}

// Buytotal is a free data retrieval call binding the contract method 0x347cd290.
//
// Solidity: function buytotal() view returns(uint256)
func (_BTLToken *BTLTokenCaller) Buytotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "buytotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Buytotal is a free data retrieval call binding the contract method 0x347cd290.
//
// Solidity: function buytotal() view returns(uint256)
func (_BTLToken *BTLTokenSession) Buytotal() (*big.Int, error) {
	return _BTLToken.Contract.Buytotal(&_BTLToken.CallOpts)
}

// Buytotal is a free data retrieval call binding the contract method 0x347cd290.
//
// Solidity: function buytotal() view returns(uint256)
func (_BTLToken *BTLTokenCallerSession) Buytotal() (*big.Int, error) {
	return _BTLToken.Contract.Buytotal(&_BTLToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BTLToken *BTLTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BTLToken *BTLTokenSession) Decimals() (uint8, error) {
	return _BTLToken.Contract.Decimals(&_BTLToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BTLToken *BTLTokenCallerSession) Decimals() (uint8, error) {
	return _BTLToken.Contract.Decimals(&_BTLToken.CallOpts)
}

// Lpwallet is a free data retrieval call binding the contract method 0x204ee5f7.
//
// Solidity: function lpwallet() view returns(address)
func (_BTLToken *BTLTokenCaller) Lpwallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "lpwallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lpwallet is a free data retrieval call binding the contract method 0x204ee5f7.
//
// Solidity: function lpwallet() view returns(address)
func (_BTLToken *BTLTokenSession) Lpwallet() (common.Address, error) {
	return _BTLToken.Contract.Lpwallet(&_BTLToken.CallOpts)
}

// Lpwallet is a free data retrieval call binding the contract method 0x204ee5f7.
//
// Solidity: function lpwallet() view returns(address)
func (_BTLToken *BTLTokenCallerSession) Lpwallet() (common.Address, error) {
	return _BTLToken.Contract.Lpwallet(&_BTLToken.CallOpts)
}

// MainPair is a free data retrieval call binding the contract method 0x85af30c5.
//
// Solidity: function mainPair() view returns(address)
func (_BTLToken *BTLTokenCaller) MainPair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "mainPair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MainPair is a free data retrieval call binding the contract method 0x85af30c5.
//
// Solidity: function mainPair() view returns(address)
func (_BTLToken *BTLTokenSession) MainPair() (common.Address, error) {
	return _BTLToken.Contract.MainPair(&_BTLToken.CallOpts)
}

// MainPair is a free data retrieval call binding the contract method 0x85af30c5.
//
// Solidity: function mainPair() view returns(address)
func (_BTLToken *BTLTokenCallerSession) MainPair() (common.Address, error) {
	return _BTLToken.Contract.MainPair(&_BTLToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BTLToken *BTLTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BTLToken *BTLTokenSession) Name() (string, error) {
	return _BTLToken.Contract.Name(&_BTLToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BTLToken *BTLTokenCallerSession) Name() (string, error) {
	return _BTLToken.Contract.Name(&_BTLToken.CallOpts)
}

// Nowstatus is a free data retrieval call binding the contract method 0x83f57264.
//
// Solidity: function nowstatus() view returns(uint8)
func (_BTLToken *BTLTokenCaller) Nowstatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "nowstatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Nowstatus is a free data retrieval call binding the contract method 0x83f57264.
//
// Solidity: function nowstatus() view returns(uint8)
func (_BTLToken *BTLTokenSession) Nowstatus() (uint8, error) {
	return _BTLToken.Contract.Nowstatus(&_BTLToken.CallOpts)
}

// Nowstatus is a free data retrieval call binding the contract method 0x83f57264.
//
// Solidity: function nowstatus() view returns(uint8)
func (_BTLToken *BTLTokenCallerSession) Nowstatus() (uint8, error) {
	return _BTLToken.Contract.Nowstatus(&_BTLToken.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address)
func (_BTLToken *BTLTokenCaller) Owners(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "owners")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address)
func (_BTLToken *BTLTokenSession) Owners() (common.Address, error) {
	return _BTLToken.Contract.Owners(&_BTLToken.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address)
func (_BTLToken *BTLTokenCallerSession) Owners() (common.Address, error) {
	return _BTLToken.Contract.Owners(&_BTLToken.CallOpts)
}

// Pendingburn is a free data retrieval call binding the contract method 0xc70fa495.
//
// Solidity: function pendingburn(address ) view returns(uint256)
func (_BTLToken *BTLTokenCaller) Pendingburn(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "pendingburn", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pendingburn is a free data retrieval call binding the contract method 0xc70fa495.
//
// Solidity: function pendingburn(address ) view returns(uint256)
func (_BTLToken *BTLTokenSession) Pendingburn(arg0 common.Address) (*big.Int, error) {
	return _BTLToken.Contract.Pendingburn(&_BTLToken.CallOpts, arg0)
}

// Pendingburn is a free data retrieval call binding the contract method 0xc70fa495.
//
// Solidity: function pendingburn(address ) view returns(uint256)
func (_BTLToken *BTLTokenCallerSession) Pendingburn(arg0 common.Address) (*big.Int, error) {
	return _BTLToken.Contract.Pendingburn(&_BTLToken.CallOpts, arg0)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_BTLToken *BTLTokenCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_BTLToken *BTLTokenSession) Router() (common.Address, error) {
	return _BTLToken.Contract.Router(&_BTLToken.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_BTLToken *BTLTokenCallerSession) Router() (common.Address, error) {
	return _BTLToken.Contract.Router(&_BTLToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BTLToken *BTLTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BTLToken *BTLTokenSession) Symbol() (string, error) {
	return _BTLToken.Contract.Symbol(&_BTLToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BTLToken *BTLTokenCallerSession) Symbol() (string, error) {
	return _BTLToken.Contract.Symbol(&_BTLToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BTLToken *BTLTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BTLToken *BTLTokenSession) TotalSupply() (*big.Int, error) {
	return _BTLToken.Contract.TotalSupply(&_BTLToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BTLToken *BTLTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BTLToken.Contract.TotalSupply(&_BTLToken.CallOpts)
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_BTLToken *BTLTokenCaller) WhiteList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "whiteList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_BTLToken *BTLTokenSession) WhiteList(arg0 common.Address) (bool, error) {
	return _BTLToken.Contract.WhiteList(&_BTLToken.CallOpts, arg0)
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_BTLToken *BTLTokenCallerSession) WhiteList(arg0 common.Address) (bool, error) {
	return _BTLToken.Contract.WhiteList(&_BTLToken.CallOpts, arg0)
}

// WhiteListLPamount is a free data retrieval call binding the contract method 0xa89ed221.
//
// Solidity: function whiteListLPamount(address ) view returns(uint256)
func (_BTLToken *BTLTokenCaller) WhiteListLPamount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BTLToken.contract.Call(opts, &out, "whiteListLPamount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhiteListLPamount is a free data retrieval call binding the contract method 0xa89ed221.
//
// Solidity: function whiteListLPamount(address ) view returns(uint256)
func (_BTLToken *BTLTokenSession) WhiteListLPamount(arg0 common.Address) (*big.Int, error) {
	return _BTLToken.Contract.WhiteListLPamount(&_BTLToken.CallOpts, arg0)
}

// WhiteListLPamount is a free data retrieval call binding the contract method 0xa89ed221.
//
// Solidity: function whiteListLPamount(address ) view returns(uint256)
func (_BTLToken *BTLTokenCallerSession) WhiteListLPamount(arg0 common.Address) (*big.Int, error) {
	return _BTLToken.Contract.WhiteListLPamount(&_BTLToken.CallOpts, arg0)
}

// Close is a paid mutator transaction binding the contract method 0xc35789cc.
//
// Solidity: function Close() returns()
func (_BTLToken *BTLTokenTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTLToken.contract.Transact(opts, "Close")
}

// Close is a paid mutator transaction binding the contract method 0xc35789cc.
//
// Solidity: function Close() returns()
func (_BTLToken *BTLTokenSession) Close() (*types.Transaction, error) {
	return _BTLToken.Contract.Close(&_BTLToken.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0xc35789cc.
//
// Solidity: function Close() returns()
func (_BTLToken *BTLTokenTransactorSession) Close() (*types.Transaction, error) {
	return _BTLToken.Contract.Close(&_BTLToken.TransactOpts)
}

// Open is a paid mutator transaction binding the contract method 0x59ebeb90.
//
// Solidity: function Open() returns()
func (_BTLToken *BTLTokenTransactor) Open(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTLToken.contract.Transact(opts, "Open")
}

// Open is a paid mutator transaction binding the contract method 0x59ebeb90.
//
// Solidity: function Open() returns()
func (_BTLToken *BTLTokenSession) Open() (*types.Transaction, error) {
	return _BTLToken.Contract.Open(&_BTLToken.TransactOpts)
}

// Open is a paid mutator transaction binding the contract method 0x59ebeb90.
//
// Solidity: function Open() returns()
func (_BTLToken *BTLTokenTransactorSession) Open() (*types.Transaction, error) {
	return _BTLToken.Contract.Open(&_BTLToken.TransactOpts)
}

// AddwhiteList is a paid mutator transaction binding the contract method 0x17c5ab32.
//
// Solidity: function addwhiteList(address _address) returns()
func (_BTLToken *BTLTokenTransactor) AddwhiteList(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _BTLToken.contract.Transact(opts, "addwhiteList", _address)
}

// AddwhiteList is a paid mutator transaction binding the contract method 0x17c5ab32.
//
// Solidity: function addwhiteList(address _address) returns()
func (_BTLToken *BTLTokenSession) AddwhiteList(_address common.Address) (*types.Transaction, error) {
	return _BTLToken.Contract.AddwhiteList(&_BTLToken.TransactOpts, _address)
}

// AddwhiteList is a paid mutator transaction binding the contract method 0x17c5ab32.
//
// Solidity: function addwhiteList(address _address) returns()
func (_BTLToken *BTLTokenTransactorSession) AddwhiteList(_address common.Address) (*types.Transaction, error) {
	return _BTLToken.Contract.AddwhiteList(&_BTLToken.TransactOpts, _address)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_BTLToken *BTLTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BTLToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_BTLToken *BTLTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BTLToken.Contract.Approve(&_BTLToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_BTLToken *BTLTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BTLToken.Contract.Approve(&_BTLToken.TransactOpts, _spender, _value)
}

// Setadmin is a paid mutator transaction binding the contract method 0xdded49cb.
//
// Solidity: function setadmin(address _address) returns()
func (_BTLToken *BTLTokenTransactor) Setadmin(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _BTLToken.contract.Transact(opts, "setadmin", _address)
}

// Setadmin is a paid mutator transaction binding the contract method 0xdded49cb.
//
// Solidity: function setadmin(address _address) returns()
func (_BTLToken *BTLTokenSession) Setadmin(_address common.Address) (*types.Transaction, error) {
	return _BTLToken.Contract.Setadmin(&_BTLToken.TransactOpts, _address)
}

// Setadmin is a paid mutator transaction binding the contract method 0xdded49cb.
//
// Solidity: function setadmin(address _address) returns()
func (_BTLToken *BTLTokenTransactorSession) Setadmin(_address common.Address) (*types.Transaction, error) {
	return _BTLToken.Contract.Setadmin(&_BTLToken.TransactOpts, _address)
}

// Setblacklist is a paid mutator transaction binding the contract method 0x5c5ba448.
//
// Solidity: function setblacklist(address _blackList) returns()
func (_BTLToken *BTLTokenTransactor) Setblacklist(opts *bind.TransactOpts, _blackList common.Address) (*types.Transaction, error) {
	return _BTLToken.contract.Transact(opts, "setblacklist", _blackList)
}

// Setblacklist is a paid mutator transaction binding the contract method 0x5c5ba448.
//
// Solidity: function setblacklist(address _blackList) returns()
func (_BTLToken *BTLTokenSession) Setblacklist(_blackList common.Address) (*types.Transaction, error) {
	return _BTLToken.Contract.Setblacklist(&_BTLToken.TransactOpts, _blackList)
}

// Setblacklist is a paid mutator transaction binding the contract method 0x5c5ba448.
//
// Solidity: function setblacklist(address _blackList) returns()
func (_BTLToken *BTLTokenTransactorSession) Setblacklist(_blackList common.Address) (*types.Transaction, error) {
	return _BTLToken.Contract.Setblacklist(&_BTLToken.TransactOpts, _blackList)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_BTLToken *BTLTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BTLToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_BTLToken *BTLTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BTLToken.Contract.Transfer(&_BTLToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_BTLToken *BTLTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BTLToken.Contract.Transfer(&_BTLToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_BTLToken *BTLTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BTLToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_BTLToken *BTLTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BTLToken.Contract.TransferFrom(&_BTLToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_BTLToken *BTLTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BTLToken.Contract.TransferFrom(&_BTLToken.TransactOpts, _from, _to, _value)
}

// BTLTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BTLToken contract.
type BTLTokenApprovalIterator struct {
	Event *BTLTokenApproval // Event containing the contract specifics and raw log

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
func (it *BTLTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BTLTokenApproval)
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
		it.Event = new(BTLTokenApproval)
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
func (it *BTLTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BTLTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BTLTokenApproval represents a Approval event raised by the BTLToken contract.
type BTLTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BTLToken *BTLTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BTLTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BTLToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BTLTokenApprovalIterator{contract: _BTLToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BTLToken *BTLTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BTLTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BTLToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BTLTokenApproval)
				if err := _BTLToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BTLToken *BTLTokenFilterer) ParseApproval(log types.Log) (*BTLTokenApproval, error) {
	event := new(BTLTokenApproval)
	if err := _BTLToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BTLTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BTLToken contract.
type BTLTokenTransferIterator struct {
	Event *BTLTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BTLTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BTLTokenTransfer)
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
		it.Event = new(BTLTokenTransfer)
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
func (it *BTLTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BTLTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BTLTokenTransfer represents a Transfer event raised by the BTLToken contract.
type BTLTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BTLToken *BTLTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BTLTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BTLToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BTLTokenTransferIterator{contract: _BTLToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BTLToken *BTLTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BTLTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BTLToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BTLTokenTransfer)
				if err := _BTLToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BTLToken *BTLTokenFilterer) ParseTransfer(log types.Log) (*BTLTokenTransfer, error) {
	event := new(BTLTokenTransfer)
	if err := _BTLToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
