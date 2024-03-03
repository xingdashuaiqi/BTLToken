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

// ERCMetaData contains all meta data concerning the ERC contract.
var ERCMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERCABI is the input ABI used to generate the binding from.
// Deprecated: Use ERCMetaData.ABI instead.
var ERCABI = ERCMetaData.ABI

// ERC is an auto generated Go binding around an Ethereum contract.
type ERC struct {
	ERCCaller     // Read-only binding to the contract
	ERCTransactor // Write-only binding to the contract
	ERCFilterer   // Log filterer for contract events
}

// ERCCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERCSession struct {
	Contract     *ERC              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERCCallerSession struct {
	Contract *ERCCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERCTransactorSession struct {
	Contract     *ERCTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERCRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERCRaw struct {
	Contract *ERC // Generic contract binding to access the raw methods on
}

// ERCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERCCallerRaw struct {
	Contract *ERCCaller // Generic read-only contract binding to access the raw methods on
}

// ERCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERCTransactorRaw struct {
	Contract *ERCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC creates a new instance of ERC, bound to a specific deployed contract.
func NewERC(address common.Address, backend bind.ContractBackend) (*ERC, error) {
	contract, err := bindERC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC{ERCCaller: ERCCaller{contract: contract}, ERCTransactor: ERCTransactor{contract: contract}, ERCFilterer: ERCFilterer{contract: contract}}, nil
}

// NewERCCaller creates a new read-only instance of ERC, bound to a specific deployed contract.
func NewERCCaller(address common.Address, caller bind.ContractCaller) (*ERCCaller, error) {
	contract, err := bindERC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERCCaller{contract: contract}, nil
}

// NewERCTransactor creates a new write-only instance of ERC, bound to a specific deployed contract.
func NewERCTransactor(address common.Address, transactor bind.ContractTransactor) (*ERCTransactor, error) {
	contract, err := bindERC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERCTransactor{contract: contract}, nil
}

// NewERCFilterer creates a new log filterer instance of ERC, bound to a specific deployed contract.
func NewERCFilterer(address common.Address, filterer bind.ContractFilterer) (*ERCFilterer, error) {
	contract, err := bindERC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERCFilterer{contract: contract}, nil
}

// bindERC binds a generic wrapper to an already deployed contract.
func bindERC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC *ERCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC.Contract.ERCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC *ERCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC.Contract.ERCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC *ERCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC.Contract.ERCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC *ERCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC *ERCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC *ERCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC *ERCCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC *ERCSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC.Contract.Allowance(&_ERC.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC *ERCCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC.Contract.Allowance(&_ERC.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC *ERCCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC *ERCSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC.Contract.BalanceOf(&_ERC.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC *ERCCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC.Contract.BalanceOf(&_ERC.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC *ERCCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC *ERCSession) Decimals() (uint8, error) {
	return _ERC.Contract.Decimals(&_ERC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC *ERCCallerSession) Decimals() (uint8, error) {
	return _ERC.Contract.Decimals(&_ERC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC *ERCCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC *ERCSession) Name() (string, error) {
	return _ERC.Contract.Name(&_ERC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC *ERCCallerSession) Name() (string, error) {
	return _ERC.Contract.Name(&_ERC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC *ERCCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC *ERCSession) Symbol() (string, error) {
	return _ERC.Contract.Symbol(&_ERC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC *ERCCallerSession) Symbol() (string, error) {
	return _ERC.Contract.Symbol(&_ERC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC *ERCCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC *ERCSession) TotalSupply() (*big.Int, error) {
	return _ERC.Contract.TotalSupply(&_ERC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC *ERCCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC.Contract.TotalSupply(&_ERC.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC *ERCTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC *ERCSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC.Contract.Approve(&_ERC.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC *ERCTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC.Contract.Approve(&_ERC.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC *ERCTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC *ERCSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC.Contract.Transfer(&_ERC.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC *ERCTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC.Contract.Transfer(&_ERC.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC *ERCTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC *ERCSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC.Contract.TransferFrom(&_ERC.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC *ERCTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC.Contract.TransferFrom(&_ERC.TransactOpts, from, to, value)
}

// ERCApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC contract.
type ERCApprovalIterator struct {
	Event *ERCApproval // Event containing the contract specifics and raw log

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
func (it *ERCApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERCApproval)
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
		it.Event = new(ERCApproval)
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
func (it *ERCApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERCApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERCApproval represents a Approval event raised by the ERC contract.
type ERCApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC *ERCFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERCApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERCApprovalIterator{contract: _ERC.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC *ERCFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERCApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERCApproval)
				if err := _ERC.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC *ERCFilterer) ParseApproval(log types.Log) (*ERCApproval, error) {
	event := new(ERCApproval)
	if err := _ERC.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERCTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC contract.
type ERCTransferIterator struct {
	Event *ERCTransfer // Event containing the contract specifics and raw log

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
func (it *ERCTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERCTransfer)
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
		it.Event = new(ERCTransfer)
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
func (it *ERCTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERCTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERCTransfer represents a Transfer event raised by the ERC contract.
type ERCTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC *ERCFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERCTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERCTransferIterator{contract: _ERC.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC *ERCFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERCTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERCTransfer)
				if err := _ERC.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC *ERCFilterer) ParseTransfer(log types.Log) (*ERCTransfer, error) {
	event := new(ERCTransfer)
	if err := _ERC.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
