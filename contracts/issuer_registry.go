// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = abi.ConvertType
)

// IssuerRegistryMetaData contains all meta data concerning the IssuerRegistry contract.
var IssuerRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"IssuerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"IssuerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"addIssuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"isTrustedIssuer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"removeIssuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506109c98061005b5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c806320694db01461005957806347bc7093146100755780638da5cb5b14610091578063ef2ed1a4146100af578063f2fde38b146100df575b5f5ffd5b610073600480360381019061006e919061075d565b6100fb565b005b61008f600480360381019061008a919061075d565b61031b565b005b6100996104cc565b6040516100a69190610797565b60405180910390f35b6100c960048036038101906100c4919061075d565b6104f0565b6040516100d691906107ca565b60405180910390f35b6100f960048036038101906100f4919061075d565b610542565b005b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610189576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101809061083d565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036101f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ee906108a5565b60405180910390fd5b60015f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610281576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102789061090d565b60405180910390fd5b6001805f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167f05e7c881d716bee8cb7ed92293133ba156704252439e5c502c277448f04e20c260405160405180910390a250565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103a09061083d565b60405180910390fd5b60015f8273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16610432576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042990610975565b60405180910390fd5b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167faf66545c919a3be306ee446d8f42a9558b5b022620df880517bc9593ec0f2d5260405160405180910390a250565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f60015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105c79061083d565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361063e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610635906108a5565b60405180910390fd5b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61072c82610703565b9050919050565b61073c81610722565b8114610746575f5ffd5b50565b5f8135905061075781610733565b92915050565b5f60208284031215610772576107716106ff565b5b5f61077f84828501610749565b91505092915050565b61079181610722565b82525050565b5f6020820190506107aa5f830184610788565b92915050565b5f8115159050919050565b6107c4816107b0565b82525050565b5f6020820190506107dd5f8301846107bb565b92915050565b5f82825260208201905092915050565b7f49737375657252656769737472793a206e6f74206f776e6572000000000000005f82015250565b5f6108276019836107e3565b9150610832826107f3565b602082019050919050565b5f6020820190508181035f8301526108548161081b565b9050919050565b7f49737375657252656769737472793a207a65726f2061646472657373000000005f82015250565b5f61088f601c836107e3565b915061089a8261085b565b602082019050919050565b5f6020820190508181035f8301526108bc81610883565b9050919050565b7f49737375657252656769737472793a20616c72656164792074727573746564005f82015250565b5f6108f7601f836107e3565b9150610902826108c3565b602082019050919050565b5f6020820190508181035f830152610924816108eb565b9050919050565b7f49737375657252656769737472793a206e6f74207472757374656400000000005f82015250565b5f61095f601b836107e3565b915061096a8261092b565b602082019050919050565b5f6020820190508181035f83015261098c81610953565b905091905056fea264697066735822122013f6227cd90c8286965924e00abe48cb5f5538f1d54071debcae2ec3e09ba79b64736f6c63430008210033",
}

// IssuerRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IssuerRegistryMetaData.ABI instead.
var IssuerRegistryABI = IssuerRegistryMetaData.ABI

// IssuerRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IssuerRegistryMetaData.Bin instead.
var IssuerRegistryBin = IssuerRegistryMetaData.Bin

// DeployIssuerRegistry deploys a new Ethereum contract, binding an instance of IssuerRegistry to it.
func DeployIssuerRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IssuerRegistry, error) {
	parsed, err := IssuerRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IssuerRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IssuerRegistry{IssuerRegistryCaller: IssuerRegistryCaller{contract: contract}, IssuerRegistryTransactor: IssuerRegistryTransactor{contract: contract}, IssuerRegistryFilterer: IssuerRegistryFilterer{contract: contract}}, nil
}

// IssuerRegistry is an auto generated Go binding around an Ethereum contract.
type IssuerRegistry struct {
	IssuerRegistryCaller     // Read-only binding to the contract
	IssuerRegistryTransactor // Write-only binding to the contract
	IssuerRegistryFilterer   // Log filterer for contract events
}

// IssuerRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IssuerRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IssuerRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IssuerRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IssuerRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IssuerRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IssuerRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IssuerRegistrySession struct {
	Contract     *IssuerRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IssuerRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IssuerRegistryCallerSession struct {
	Contract *IssuerRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IssuerRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IssuerRegistryTransactorSession struct {
	Contract     *IssuerRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IssuerRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IssuerRegistryRaw struct {
	Contract *IssuerRegistry // Generic contract binding to access the raw methods on
}

// IssuerRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IssuerRegistryCallerRaw struct {
	Contract *IssuerRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IssuerRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IssuerRegistryTransactorRaw struct {
	Contract *IssuerRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIssuerRegistry creates a new instance of IssuerRegistry, bound to a specific deployed contract.
func NewIssuerRegistry(address common.Address, backend bind.ContractBackend) (*IssuerRegistry, error) {
	contract, err := bindIssuerRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IssuerRegistry{IssuerRegistryCaller: IssuerRegistryCaller{contract: contract}, IssuerRegistryTransactor: IssuerRegistryTransactor{contract: contract}, IssuerRegistryFilterer: IssuerRegistryFilterer{contract: contract}}, nil
}

// NewIssuerRegistryCaller creates a new read-only instance of IssuerRegistry, bound to a specific deployed contract.
func NewIssuerRegistryCaller(address common.Address, caller bind.ContractCaller) (*IssuerRegistryCaller, error) {
	contract, err := bindIssuerRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IssuerRegistryCaller{contract: contract}, nil
}

// NewIssuerRegistryTransactor creates a new write-only instance of IssuerRegistry, bound to a specific deployed contract.
func NewIssuerRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IssuerRegistryTransactor, error) {
	contract, err := bindIssuerRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IssuerRegistryTransactor{contract: contract}, nil
}

// NewIssuerRegistryFilterer creates a new log filterer instance of IssuerRegistry, bound to a specific deployed contract.
func NewIssuerRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IssuerRegistryFilterer, error) {
	contract, err := bindIssuerRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IssuerRegistryFilterer{contract: contract}, nil
}

// bindIssuerRegistry binds a generic wrapper to an already deployed contract.
func bindIssuerRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IssuerRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IssuerRegistry *IssuerRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IssuerRegistry.Contract.IssuerRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IssuerRegistry *IssuerRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IssuerRegistry.Contract.IssuerRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IssuerRegistry *IssuerRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IssuerRegistry.Contract.IssuerRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IssuerRegistry *IssuerRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IssuerRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IssuerRegistry *IssuerRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IssuerRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IssuerRegistry *IssuerRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IssuerRegistry.Contract.contract.Transact(opts, method, params...)
}

// IsTrustedIssuer is a free data retrieval call binding the contract method 0xef2ed1a4.
//
// Solidity: function isTrustedIssuer(address issuer) view returns(bool)
func (_IssuerRegistry *IssuerRegistryCaller) IsTrustedIssuer(opts *bind.CallOpts, issuer common.Address) (bool, error) {
	var out []interface{}
	err := _IssuerRegistry.contract.Call(opts, &out, "isTrustedIssuer", issuer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedIssuer is a free data retrieval call binding the contract method 0xef2ed1a4.
//
// Solidity: function isTrustedIssuer(address issuer) view returns(bool)
func (_IssuerRegistry *IssuerRegistrySession) IsTrustedIssuer(issuer common.Address) (bool, error) {
	return _IssuerRegistry.Contract.IsTrustedIssuer(&_IssuerRegistry.CallOpts, issuer)
}

// IsTrustedIssuer is a free data retrieval call binding the contract method 0xef2ed1a4.
//
// Solidity: function isTrustedIssuer(address issuer) view returns(bool)
func (_IssuerRegistry *IssuerRegistryCallerSession) IsTrustedIssuer(issuer common.Address) (bool, error) {
	return _IssuerRegistry.Contract.IsTrustedIssuer(&_IssuerRegistry.CallOpts, issuer)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IssuerRegistry *IssuerRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IssuerRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IssuerRegistry *IssuerRegistrySession) Owner() (common.Address, error) {
	return _IssuerRegistry.Contract.Owner(&_IssuerRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IssuerRegistry *IssuerRegistryCallerSession) Owner() (common.Address, error) {
	return _IssuerRegistry.Contract.Owner(&_IssuerRegistry.CallOpts)
}

// AddIssuer is a paid mutator transaction binding the contract method 0x20694db0.
//
// Solidity: function addIssuer(address issuer) returns()
func (_IssuerRegistry *IssuerRegistryTransactor) AddIssuer(opts *bind.TransactOpts, issuer common.Address) (*types.Transaction, error) {
	return _IssuerRegistry.contract.Transact(opts, "addIssuer", issuer)
}

// AddIssuer is a paid mutator transaction binding the contract method 0x20694db0.
//
// Solidity: function addIssuer(address issuer) returns()
func (_IssuerRegistry *IssuerRegistrySession) AddIssuer(issuer common.Address) (*types.Transaction, error) {
	return _IssuerRegistry.Contract.AddIssuer(&_IssuerRegistry.TransactOpts, issuer)
}

// AddIssuer is a paid mutator transaction binding the contract method 0x20694db0.
//
// Solidity: function addIssuer(address issuer) returns()
func (_IssuerRegistry *IssuerRegistryTransactorSession) AddIssuer(issuer common.Address) (*types.Transaction, error) {
	return _IssuerRegistry.Contract.AddIssuer(&_IssuerRegistry.TransactOpts, issuer)
}

// RemoveIssuer is a paid mutator transaction binding the contract method 0x47bc7093.
//
// Solidity: function removeIssuer(address issuer) returns()
func (_IssuerRegistry *IssuerRegistryTransactor) RemoveIssuer(opts *bind.TransactOpts, issuer common.Address) (*types.Transaction, error) {
	return _IssuerRegistry.contract.Transact(opts, "removeIssuer", issuer)
}

// RemoveIssuer is a paid mutator transaction binding the contract method 0x47bc7093.
//
// Solidity: function removeIssuer(address issuer) returns()
func (_IssuerRegistry *IssuerRegistrySession) RemoveIssuer(issuer common.Address) (*types.Transaction, error) {
	return _IssuerRegistry.Contract.RemoveIssuer(&_IssuerRegistry.TransactOpts, issuer)
}

// RemoveIssuer is a paid mutator transaction binding the contract method 0x47bc7093.
//
// Solidity: function removeIssuer(address issuer) returns()
func (_IssuerRegistry *IssuerRegistryTransactorSession) RemoveIssuer(issuer common.Address) (*types.Transaction, error) {
	return _IssuerRegistry.Contract.RemoveIssuer(&_IssuerRegistry.TransactOpts, issuer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IssuerRegistry *IssuerRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IssuerRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IssuerRegistry *IssuerRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IssuerRegistry.Contract.TransferOwnership(&_IssuerRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IssuerRegistry *IssuerRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IssuerRegistry.Contract.TransferOwnership(&_IssuerRegistry.TransactOpts, newOwner)
}

// IssuerRegistryIssuerAddedIterator is returned from FilterIssuerAdded and is used to iterate over the raw logs and unpacked data for IssuerAdded events raised by the IssuerRegistry contract.
type IssuerRegistryIssuerAddedIterator struct {
	Event *IssuerRegistryIssuerAdded // Event containing the contract specifics and raw log

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
func (it *IssuerRegistryIssuerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IssuerRegistryIssuerAdded)
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
		it.Event = new(IssuerRegistryIssuerAdded)
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
func (it *IssuerRegistryIssuerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IssuerRegistryIssuerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IssuerRegistryIssuerAdded represents a IssuerAdded event raised by the IssuerRegistry contract.
type IssuerRegistryIssuerAdded struct {
	Issuer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssuerAdded is a free log retrieval operation binding the contract event 0x05e7c881d716bee8cb7ed92293133ba156704252439e5c502c277448f04e20c2.
//
// Solidity: event IssuerAdded(address indexed issuer)
func (_IssuerRegistry *IssuerRegistryFilterer) FilterIssuerAdded(opts *bind.FilterOpts, issuer []common.Address) (*IssuerRegistryIssuerAddedIterator, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _IssuerRegistry.contract.FilterLogs(opts, "IssuerAdded", issuerRule)
	if err != nil {
		return nil, err
	}
	return &IssuerRegistryIssuerAddedIterator{contract: _IssuerRegistry.contract, event: "IssuerAdded", logs: logs, sub: sub}, nil
}

// WatchIssuerAdded is a free log subscription operation binding the contract event 0x05e7c881d716bee8cb7ed92293133ba156704252439e5c502c277448f04e20c2.
//
// Solidity: event IssuerAdded(address indexed issuer)
func (_IssuerRegistry *IssuerRegistryFilterer) WatchIssuerAdded(opts *bind.WatchOpts, sink chan<- *IssuerRegistryIssuerAdded, issuer []common.Address) (event.Subscription, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _IssuerRegistry.contract.WatchLogs(opts, "IssuerAdded", issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IssuerRegistryIssuerAdded)
				if err := _IssuerRegistry.contract.UnpackLog(event, "IssuerAdded", log); err != nil {
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

// ParseIssuerAdded is a log parse operation binding the contract event 0x05e7c881d716bee8cb7ed92293133ba156704252439e5c502c277448f04e20c2.
//
// Solidity: event IssuerAdded(address indexed issuer)
func (_IssuerRegistry *IssuerRegistryFilterer) ParseIssuerAdded(log types.Log) (*IssuerRegistryIssuerAdded, error) {
	event := new(IssuerRegistryIssuerAdded)
	if err := _IssuerRegistry.contract.UnpackLog(event, "IssuerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IssuerRegistryIssuerRemovedIterator is returned from FilterIssuerRemoved and is used to iterate over the raw logs and unpacked data for IssuerRemoved events raised by the IssuerRegistry contract.
type IssuerRegistryIssuerRemovedIterator struct {
	Event *IssuerRegistryIssuerRemoved // Event containing the contract specifics and raw log

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
func (it *IssuerRegistryIssuerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IssuerRegistryIssuerRemoved)
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
		it.Event = new(IssuerRegistryIssuerRemoved)
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
func (it *IssuerRegistryIssuerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IssuerRegistryIssuerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IssuerRegistryIssuerRemoved represents a IssuerRemoved event raised by the IssuerRegistry contract.
type IssuerRegistryIssuerRemoved struct {
	Issuer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssuerRemoved is a free log retrieval operation binding the contract event 0xaf66545c919a3be306ee446d8f42a9558b5b022620df880517bc9593ec0f2d52.
//
// Solidity: event IssuerRemoved(address indexed issuer)
func (_IssuerRegistry *IssuerRegistryFilterer) FilterIssuerRemoved(opts *bind.FilterOpts, issuer []common.Address) (*IssuerRegistryIssuerRemovedIterator, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _IssuerRegistry.contract.FilterLogs(opts, "IssuerRemoved", issuerRule)
	if err != nil {
		return nil, err
	}
	return &IssuerRegistryIssuerRemovedIterator{contract: _IssuerRegistry.contract, event: "IssuerRemoved", logs: logs, sub: sub}, nil
}

// WatchIssuerRemoved is a free log subscription operation binding the contract event 0xaf66545c919a3be306ee446d8f42a9558b5b022620df880517bc9593ec0f2d52.
//
// Solidity: event IssuerRemoved(address indexed issuer)
func (_IssuerRegistry *IssuerRegistryFilterer) WatchIssuerRemoved(opts *bind.WatchOpts, sink chan<- *IssuerRegistryIssuerRemoved, issuer []common.Address) (event.Subscription, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _IssuerRegistry.contract.WatchLogs(opts, "IssuerRemoved", issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IssuerRegistryIssuerRemoved)
				if err := _IssuerRegistry.contract.UnpackLog(event, "IssuerRemoved", log); err != nil {
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

// ParseIssuerRemoved is a log parse operation binding the contract event 0xaf66545c919a3be306ee446d8f42a9558b5b022620df880517bc9593ec0f2d52.
//
// Solidity: event IssuerRemoved(address indexed issuer)
func (_IssuerRegistry *IssuerRegistryFilterer) ParseIssuerRemoved(log types.Log) (*IssuerRegistryIssuerRemoved, error) {
	event := new(IssuerRegistryIssuerRemoved)
	if err := _IssuerRegistry.contract.UnpackLog(event, "IssuerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IssuerRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IssuerRegistry contract.
type IssuerRegistryOwnershipTransferredIterator struct {
	Event *IssuerRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IssuerRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IssuerRegistryOwnershipTransferred)
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
		it.Event = new(IssuerRegistryOwnershipTransferred)
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
func (it *IssuerRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IssuerRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IssuerRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the IssuerRegistry contract.
type IssuerRegistryOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_IssuerRegistry *IssuerRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*IssuerRegistryOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IssuerRegistry.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IssuerRegistryOwnershipTransferredIterator{contract: _IssuerRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_IssuerRegistry *IssuerRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IssuerRegistryOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IssuerRegistry.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IssuerRegistryOwnershipTransferred)
				if err := _IssuerRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_IssuerRegistry *IssuerRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*IssuerRegistryOwnershipTransferred, error) {
	event := new(IssuerRegistryOwnershipTransferred)
	if err := _IssuerRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
