// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zecreyLegend

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

// StablePriceOracleMetaData contains all meta data concerning the StablePriceOracle contract.
var StablePriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_rentPrices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"RentPriceChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_rentPrices\",\"type\":\"uint256[]\"}],\"name\":\"changeRentPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1Letter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price2Letter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price3Letter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StablePriceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use StablePriceOracleMetaData.ABI instead.
var StablePriceOracleABI = StablePriceOracleMetaData.ABI

// StablePriceOracle is an auto generated Go binding around an Ethereum contract.
type StablePriceOracle struct {
	StablePriceOracleCaller     // Read-only binding to the contract
	StablePriceOracleTransactor // Write-only binding to the contract
	StablePriceOracleFilterer   // Log filterer for contract events
}

// StablePriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type StablePriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StablePriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StablePriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StablePriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StablePriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StablePriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StablePriceOracleSession struct {
	Contract     *StablePriceOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StablePriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StablePriceOracleCallerSession struct {
	Contract *StablePriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// StablePriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StablePriceOracleTransactorSession struct {
	Contract     *StablePriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// StablePriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type StablePriceOracleRaw struct {
	Contract *StablePriceOracle // Generic contract binding to access the raw methods on
}

// StablePriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StablePriceOracleCallerRaw struct {
	Contract *StablePriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// StablePriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StablePriceOracleTransactorRaw struct {
	Contract *StablePriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStablePriceOracle creates a new instance of StablePriceOracle, bound to a specific deployed contract.
func NewStablePriceOracle(address common.Address, backend bind.ContractBackend) (*StablePriceOracle, error) {
	contract, err := bindStablePriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StablePriceOracle{StablePriceOracleCaller: StablePriceOracleCaller{contract: contract}, StablePriceOracleTransactor: StablePriceOracleTransactor{contract: contract}, StablePriceOracleFilterer: StablePriceOracleFilterer{contract: contract}}, nil
}

// NewStablePriceOracleCaller creates a new read-only instance of StablePriceOracle, bound to a specific deployed contract.
func NewStablePriceOracleCaller(address common.Address, caller bind.ContractCaller) (*StablePriceOracleCaller, error) {
	contract, err := bindStablePriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StablePriceOracleCaller{contract: contract}, nil
}

// NewStablePriceOracleTransactor creates a new write-only instance of StablePriceOracle, bound to a specific deployed contract.
func NewStablePriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*StablePriceOracleTransactor, error) {
	contract, err := bindStablePriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StablePriceOracleTransactor{contract: contract}, nil
}

// NewStablePriceOracleFilterer creates a new log filterer instance of StablePriceOracle, bound to a specific deployed contract.
func NewStablePriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*StablePriceOracleFilterer, error) {
	contract, err := bindStablePriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StablePriceOracleFilterer{contract: contract}, nil
}

// bindStablePriceOracle binds a generic wrapper to an already deployed contract.
func bindStablePriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StablePriceOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StablePriceOracle *StablePriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StablePriceOracle.Contract.StablePriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StablePriceOracle *StablePriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StablePriceOracle.Contract.StablePriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StablePriceOracle *StablePriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StablePriceOracle.Contract.StablePriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StablePriceOracle *StablePriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StablePriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StablePriceOracle *StablePriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StablePriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StablePriceOracle *StablePriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StablePriceOracle.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StablePriceOracle *StablePriceOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StablePriceOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StablePriceOracle *StablePriceOracleSession) Owner() (common.Address, error) {
	return _StablePriceOracle.Contract.Owner(&_StablePriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StablePriceOracle *StablePriceOracleCallerSession) Owner() (common.Address, error) {
	return _StablePriceOracle.Contract.Owner(&_StablePriceOracle.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xfe2c6198.
//
// Solidity: function price(string name) view returns(uint256)
func (_StablePriceOracle *StablePriceOracleCaller) Price(opts *bind.CallOpts, name string) (*big.Int, error) {
	var out []interface{}
	err := _StablePriceOracle.contract.Call(opts, &out, "price", name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xfe2c6198.
//
// Solidity: function price(string name) view returns(uint256)
func (_StablePriceOracle *StablePriceOracleSession) Price(name string) (*big.Int, error) {
	return _StablePriceOracle.Contract.Price(&_StablePriceOracle.CallOpts, name)
}

// Price is a free data retrieval call binding the contract method 0xfe2c6198.
//
// Solidity: function price(string name) view returns(uint256)
func (_StablePriceOracle *StablePriceOracleCallerSession) Price(name string) (*big.Int, error) {
	return _StablePriceOracle.Contract.Price(&_StablePriceOracle.CallOpts, name)
}

// Price1Letter is a free data retrieval call binding the contract method 0x2c0fd74c.
//
// Solidity: function price1Letter() view returns(uint256)
func (_StablePriceOracle *StablePriceOracleCaller) Price1Letter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StablePriceOracle.contract.Call(opts, &out, "price1Letter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1Letter is a free data retrieval call binding the contract method 0x2c0fd74c.
//
// Solidity: function price1Letter() view returns(uint256)
func (_StablePriceOracle *StablePriceOracleSession) Price1Letter() (*big.Int, error) {
	return _StablePriceOracle.Contract.Price1Letter(&_StablePriceOracle.CallOpts)
}

// Price1Letter is a free data retrieval call binding the contract method 0x2c0fd74c.
//
// Solidity: function price1Letter() view returns(uint256)
func (_StablePriceOracle *StablePriceOracleCallerSession) Price1Letter() (*big.Int, error) {
	return _StablePriceOracle.Contract.Price1Letter(&_StablePriceOracle.CallOpts)
}

// Price2Letter is a free data retrieval call binding the contract method 0xcd5d2c74.
//
// Solidity: function price2Letter() view returns(uint256)
func (_StablePriceOracle *StablePriceOracleCaller) Price2Letter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StablePriceOracle.contract.Call(opts, &out, "price2Letter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price2Letter is a free data retrieval call binding the contract method 0xcd5d2c74.
//
// Solidity: function price2Letter() view returns(uint256)
func (_StablePriceOracle *StablePriceOracleSession) Price2Letter() (*big.Int, error) {
	return _StablePriceOracle.Contract.Price2Letter(&_StablePriceOracle.CallOpts)
}

// Price2Letter is a free data retrieval call binding the contract method 0xcd5d2c74.
//
// Solidity: function price2Letter() view returns(uint256)
func (_StablePriceOracle *StablePriceOracleCallerSession) Price2Letter() (*big.Int, error) {
	return _StablePriceOracle.Contract.Price2Letter(&_StablePriceOracle.CallOpts)
}

// Price3Letter is a free data retrieval call binding the contract method 0xa200e153.
//
// Solidity: function price3Letter() view returns(uint256)
func (_StablePriceOracle *StablePriceOracleCaller) Price3Letter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StablePriceOracle.contract.Call(opts, &out, "price3Letter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price3Letter is a free data retrieval call binding the contract method 0xa200e153.
//
// Solidity: function price3Letter() view returns(uint256)
func (_StablePriceOracle *StablePriceOracleSession) Price3Letter() (*big.Int, error) {
	return _StablePriceOracle.Contract.Price3Letter(&_StablePriceOracle.CallOpts)
}

// Price3Letter is a free data retrieval call binding the contract method 0xa200e153.
//
// Solidity: function price3Letter() view returns(uint256)
func (_StablePriceOracle *StablePriceOracleCallerSession) Price3Letter() (*big.Int, error) {
	return _StablePriceOracle.Contract.Price3Letter(&_StablePriceOracle.CallOpts)
}

// ChangeRentPrice is a paid mutator transaction binding the contract method 0x50f13aa8.
//
// Solidity: function changeRentPrice(uint256[] _rentPrices) returns()
func (_StablePriceOracle *StablePriceOracleTransactor) ChangeRentPrice(opts *bind.TransactOpts, _rentPrices []*big.Int) (*types.Transaction, error) {
	return _StablePriceOracle.contract.Transact(opts, "changeRentPrice", _rentPrices)
}

// ChangeRentPrice is a paid mutator transaction binding the contract method 0x50f13aa8.
//
// Solidity: function changeRentPrice(uint256[] _rentPrices) returns()
func (_StablePriceOracle *StablePriceOracleSession) ChangeRentPrice(_rentPrices []*big.Int) (*types.Transaction, error) {
	return _StablePriceOracle.Contract.ChangeRentPrice(&_StablePriceOracle.TransactOpts, _rentPrices)
}

// ChangeRentPrice is a paid mutator transaction binding the contract method 0x50f13aa8.
//
// Solidity: function changeRentPrice(uint256[] _rentPrices) returns()
func (_StablePriceOracle *StablePriceOracleTransactorSession) ChangeRentPrice(_rentPrices []*big.Int) (*types.Transaction, error) {
	return _StablePriceOracle.Contract.ChangeRentPrice(&_StablePriceOracle.TransactOpts, _rentPrices)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StablePriceOracle *StablePriceOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StablePriceOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StablePriceOracle *StablePriceOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _StablePriceOracle.Contract.RenounceOwnership(&_StablePriceOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StablePriceOracle *StablePriceOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StablePriceOracle.Contract.RenounceOwnership(&_StablePriceOracle.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StablePriceOracle *StablePriceOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StablePriceOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StablePriceOracle *StablePriceOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StablePriceOracle.Contract.TransferOwnership(&_StablePriceOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StablePriceOracle *StablePriceOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StablePriceOracle.Contract.TransferOwnership(&_StablePriceOracle.TransactOpts, newOwner)
}

// StablePriceOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StablePriceOracle contract.
type StablePriceOracleOwnershipTransferredIterator struct {
	Event *StablePriceOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StablePriceOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablePriceOracleOwnershipTransferred)
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
		it.Event = new(StablePriceOracleOwnershipTransferred)
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
func (it *StablePriceOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablePriceOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablePriceOracleOwnershipTransferred represents a OwnershipTransferred event raised by the StablePriceOracle contract.
type StablePriceOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StablePriceOracle *StablePriceOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StablePriceOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StablePriceOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StablePriceOracleOwnershipTransferredIterator{contract: _StablePriceOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StablePriceOracle *StablePriceOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StablePriceOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StablePriceOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablePriceOracleOwnershipTransferred)
				if err := _StablePriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StablePriceOracle *StablePriceOracleFilterer) ParseOwnershipTransferred(log types.Log) (*StablePriceOracleOwnershipTransferred, error) {
	event := new(StablePriceOracleOwnershipTransferred)
	if err := _StablePriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablePriceOracleRentPriceChangedIterator is returned from FilterRentPriceChanged and is used to iterate over the raw logs and unpacked data for RentPriceChanged events raised by the StablePriceOracle contract.
type StablePriceOracleRentPriceChangedIterator struct {
	Event *StablePriceOracleRentPriceChanged // Event containing the contract specifics and raw log

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
func (it *StablePriceOracleRentPriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablePriceOracleRentPriceChanged)
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
		it.Event = new(StablePriceOracleRentPriceChanged)
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
func (it *StablePriceOracleRentPriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablePriceOracleRentPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablePriceOracleRentPriceChanged represents a RentPriceChanged event raised by the StablePriceOracle contract.
type StablePriceOracleRentPriceChanged struct {
	Prices []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRentPriceChanged is a free log retrieval operation binding the contract event 0x73422d94aedd596c2d4d39f27a01033adc390a9054efaf259afefd95ef7331df.
//
// Solidity: event RentPriceChanged(uint256[] prices)
func (_StablePriceOracle *StablePriceOracleFilterer) FilterRentPriceChanged(opts *bind.FilterOpts) (*StablePriceOracleRentPriceChangedIterator, error) {

	logs, sub, err := _StablePriceOracle.contract.FilterLogs(opts, "RentPriceChanged")
	if err != nil {
		return nil, err
	}
	return &StablePriceOracleRentPriceChangedIterator{contract: _StablePriceOracle.contract, event: "RentPriceChanged", logs: logs, sub: sub}, nil
}

// WatchRentPriceChanged is a free log subscription operation binding the contract event 0x73422d94aedd596c2d4d39f27a01033adc390a9054efaf259afefd95ef7331df.
//
// Solidity: event RentPriceChanged(uint256[] prices)
func (_StablePriceOracle *StablePriceOracleFilterer) WatchRentPriceChanged(opts *bind.WatchOpts, sink chan<- *StablePriceOracleRentPriceChanged) (event.Subscription, error) {

	logs, sub, err := _StablePriceOracle.contract.WatchLogs(opts, "RentPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablePriceOracleRentPriceChanged)
				if err := _StablePriceOracle.contract.UnpackLog(event, "RentPriceChanged", log); err != nil {
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

// ParseRentPriceChanged is a log parse operation binding the contract event 0x73422d94aedd596c2d4d39f27a01033adc390a9054efaf259afefd95ef7331df.
//
// Solidity: event RentPriceChanged(uint256[] prices)
func (_StablePriceOracle *StablePriceOracleFilterer) ParseRentPriceChanged(log types.Log) (*StablePriceOracleRentPriceChanged, error) {
	event := new(StablePriceOracleRentPriceChanged)
	if err := _StablePriceOracle.contract.UnpackLog(event, "RentPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
