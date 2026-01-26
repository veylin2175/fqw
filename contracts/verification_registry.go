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

// VerificationRegistryMetaData contains all meta data concerning the VerificationRegistry contract.
var VerificationRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_issuerRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NFTMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"vcHash\",\"type\":\"bytes32\"}],\"name\":\"VerificationRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"VerificationRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"issuerRegistry\",\"outputs\":[{\"internalType\":\"contractIssuerRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vcHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"registerVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"revokeVerification\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"}],\"name\":\"setNFTContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"vcHash\",\"type\":\"bytes32\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"revoked\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"issuedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161138b38038061138b83398181016040528101906100319190610143565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361009f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610096906101ee565b60405180910390fd5b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505061020c565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610112826100e9565b9050919050565b61012281610108565b811461012c575f5ffd5b50565b5f8151905061013d81610119565b92915050565b5f60208284031215610158576101576100e5565b5b5f6101658482850161012f565b91505092915050565b5f82825260208201905092915050565b7f566572696669636174696f6e52656769737472793a207a65726f2061646472655f8201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b5f6101d860228361016e565b91506101e38261017e565b604082019050919050565b5f6020820190508181035f830152610205816101cc565b9050919050565b611172806102195f395ff3fe608060405234801561000f575f5ffd5b506004361061007b575f3560e01c8063a7ccabdf11610059578063a7ccabdf146100ec578063d082e38114610108578063d56d229d14610126578063ddfccf46146101445761007b565b80631571f0c61461007f5780631d97ce311461009b5780638bfc1851146100ce575b5f5ffd5b61009960048036038101906100949190610970565b610174565b005b6100b560048036038101906100b091906109ce565b6102b6565b6040516100c59493929190610a74565b60405180910390f35b6100d661038f565b6040516100e39190610b12565b60405180910390f35b61010660048036038101906101019190610b55565b6103b4565b005b6101106104f4565b60405161011d9190610b80565b60405180910390f35b61012e6104fa565b60405161013b9190610b99565b60405180910390f35b61015e60048036038101906101599190610bb2565b61051f565b60405161016b9190610b80565b60405180910390f35b5f5f5f8381526020019081526020015f2090503373ffffffffffffffffffffffffffffffffffffffff16816001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610218576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161020f90610c4a565b60405180910390fd5b806003015f9054906101000a900460ff1615610269576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026090610cd8565b60405180910390fd5b6001816003015f6101000a81548160ff021916908315150217905550817fea6232c80090e591dc995461c6943381a4246f004efce58e311f6dae40e3f83e60405160405180910390a25050565b5f5f5f5f5f5f5f8881526020019081526020015f206040518060800160405290815f8201548152602001600182015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201548152602001600382015f9054906101000a900460ff161515151581525050905085815f015114801561036e57508060600151155b94508060200151935080606001519250806040015191505092959194509250565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f73ffffffffffffffffffffffffffffffffffffffff1660035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610443576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161043a90610d66565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036104b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a890610df4565b60405180910390fd5b8060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60015481565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ef2ed1a4336040518263ffffffff1660e01b815260040161057a9190610b99565b602060405180830381865afa158015610595573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105b99190610e3c565b6105f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ef90610ed7565b60405180910390fd5b5f5f1b830361063c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063390610f3f565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036106aa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a190610fcd565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff1660035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610739576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107309061105b565b60405180910390fd5b5f60015f8154610748906110a6565b919050819055905060405180608001604052808581526020013373ffffffffffffffffffffffffffffffffffffffff1681526020014281526020015f15158152505f5f8381526020019081526020015f205f820151815f01556020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201556060820151816003015f6101000a81548160ff0219169083151502179055509050503373ffffffffffffffffffffffffffffffffffffffff16817f0e5026488e8f9bb8126eae97bb91a9b730c9e194938b6f06a2c1309e19671c868660405161085a91906110fc565b60405180910390a360035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1984836040518363ffffffff1660e01b81526004016108be929190611115565b5f604051808303815f87803b1580156108d5575f5ffd5b505af11580156108e7573d5f5f3e3d5ffd5b505050508273ffffffffffffffffffffffffffffffffffffffff16817f0176f203df400d7bd5f1b1c9ef36c16709bf3b5d9fd35f000a6bae32393f66c360405160405180910390a38091505092915050565b5f5ffd5b5f819050919050565b61094f8161093d565b8114610959575f5ffd5b50565b5f8135905061096a81610946565b92915050565b5f6020828403121561098557610984610939565b5b5f6109928482850161095c565b91505092915050565b5f819050919050565b6109ad8161099b565b81146109b7575f5ffd5b50565b5f813590506109c8816109a4565b92915050565b5f5f604083850312156109e4576109e3610939565b5b5f6109f18582860161095c565b9250506020610a02858286016109ba565b9150509250929050565b5f8115159050919050565b610a2081610a0c565b82525050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610a4f82610a26565b9050919050565b610a5f81610a45565b82525050565b610a6e8161093d565b82525050565b5f608082019050610a875f830187610a17565b610a946020830186610a56565b610aa16040830185610a17565b610aae6060830184610a65565b95945050505050565b5f819050919050565b5f610ada610ad5610ad084610a26565b610ab7565b610a26565b9050919050565b5f610aeb82610ac0565b9050919050565b5f610afc82610ae1565b9050919050565b610b0c81610af2565b82525050565b5f602082019050610b255f830184610b03565b92915050565b610b3481610a45565b8114610b3e575f5ffd5b50565b5f81359050610b4f81610b2b565b92915050565b5f60208284031215610b6a57610b69610939565b5b5f610b7784828501610b41565b91505092915050565b5f602082019050610b935f830184610a65565b92915050565b5f602082019050610bac5f830184610a56565b92915050565b5f5f60408385031215610bc857610bc7610939565b5b5f610bd5858286016109ba565b9250506020610be685828601610b41565b9150509250929050565b5f82825260208201905092915050565b7f566572696669636174696f6e52656769737472793a206e6f74206973737565725f82015250565b5f610c34602083610bf0565b9150610c3f82610c00565b602082019050919050565b5f6020820190508181035f830152610c6181610c28565b9050919050565b7f566572696669636174696f6e52656769737472793a20616c72656164792072655f8201527f766f6b6564000000000000000000000000000000000000000000000000000000602082015250565b5f610cc2602583610bf0565b9150610ccd82610c68565b604082019050919050565b5f6020820190508181035f830152610cef81610cb6565b9050919050565b7f566572696669636174696f6e52656769737472793a204e465420616c726561645f8201527f7920736574000000000000000000000000000000000000000000000000000000602082015250565b5f610d50602583610bf0565b9150610d5b82610cf6565b604082019050919050565b5f6020820190508181035f830152610d7d81610d44565b9050919050565b7f566572696669636174696f6e52656769737472793a207a65726f2061646472655f8201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b5f610dde602283610bf0565b9150610de982610d84565b604082019050919050565b5f6020820190508181035f830152610e0b81610dd2565b9050919050565b610e1b81610a0c565b8114610e25575f5ffd5b50565b5f81519050610e3681610e12565b92915050565b5f60208284031215610e5157610e50610939565b5b5f610e5e84828501610e28565b91505092915050565b7f566572696669636174696f6e52656769737472793a206e6f74207472757374655f8201527f6420697373756572000000000000000000000000000000000000000000000000602082015250565b5f610ec1602883610bf0565b9150610ecc82610e67565b604082019050919050565b5f6020820190508181035f830152610eee81610eb5565b9050919050565b7f566572696669636174696f6e52656769737472793a20656d70747920686173685f82015250565b5f610f29602083610bf0565b9150610f3482610ef5565b602082019050919050565b5f6020820190508181035f830152610f5681610f1d565b9050919050565b7f566572696669636174696f6e52656769737472793a207a65726f207375626a655f8201527f6374000000000000000000000000000000000000000000000000000000000000602082015250565b5f610fb7602283610bf0565b9150610fc282610f5d565b604082019050919050565b5f6020820190508181035f830152610fe481610fab565b9050919050565b7f566572696669636174696f6e52656769737472793a204e4654206e6f742073655f8201527f7400000000000000000000000000000000000000000000000000000000000000602082015250565b5f611045602183610bf0565b915061105082610feb565b604082019050919050565b5f6020820190508181035f83015261107281611039565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6110b08261093d565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036110e2576110e1611079565b5b600182019050919050565b6110f68161099b565b82525050565b5f60208201905061110f5f8301846110ed565b92915050565b5f6040820190506111285f830185610a56565b6111356020830184610a65565b939250505056fea2646970667358221220de48614e2e83c6b6dc630b729df53bb05e0080eaa58576a2757bf853648df9aa64736f6c63430008210033",
}

// VerificationRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use VerificationRegistryMetaData.ABI instead.
var VerificationRegistryABI = VerificationRegistryMetaData.ABI

// VerificationRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerificationRegistryMetaData.Bin instead.
var VerificationRegistryBin = VerificationRegistryMetaData.Bin

// DeployVerificationRegistry deploys a new Ethereum contract, binding an instance of VerificationRegistry to it.
func DeployVerificationRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _issuerRegistry common.Address) (common.Address, *types.Transaction, *VerificationRegistry, error) {
	parsed, err := VerificationRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerificationRegistryBin), backend, _issuerRegistry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerificationRegistry{VerificationRegistryCaller: VerificationRegistryCaller{contract: contract}, VerificationRegistryTransactor: VerificationRegistryTransactor{contract: contract}, VerificationRegistryFilterer: VerificationRegistryFilterer{contract: contract}}, nil
}

// VerificationRegistry is an auto generated Go binding around an Ethereum contract.
type VerificationRegistry struct {
	VerificationRegistryCaller     // Read-only binding to the contract
	VerificationRegistryTransactor // Write-only binding to the contract
	VerificationRegistryFilterer   // Log filterer for contract events
}

// VerificationRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerificationRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerificationRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerificationRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerificationRegistrySession struct {
	Contract     *VerificationRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// VerificationRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerificationRegistryCallerSession struct {
	Contract *VerificationRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// VerificationRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerificationRegistryTransactorSession struct {
	Contract     *VerificationRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// VerificationRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerificationRegistryRaw struct {
	Contract *VerificationRegistry // Generic contract binding to access the raw methods on
}

// VerificationRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerificationRegistryCallerRaw struct {
	Contract *VerificationRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// VerificationRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerificationRegistryTransactorRaw struct {
	Contract *VerificationRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerificationRegistry creates a new instance of VerificationRegistry, bound to a specific deployed contract.
func NewVerificationRegistry(address common.Address, backend bind.ContractBackend) (*VerificationRegistry, error) {
	contract, err := bindVerificationRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerificationRegistry{VerificationRegistryCaller: VerificationRegistryCaller{contract: contract}, VerificationRegistryTransactor: VerificationRegistryTransactor{contract: contract}, VerificationRegistryFilterer: VerificationRegistryFilterer{contract: contract}}, nil
}

// NewVerificationRegistryCaller creates a new read-only instance of VerificationRegistry, bound to a specific deployed contract.
func NewVerificationRegistryCaller(address common.Address, caller bind.ContractCaller) (*VerificationRegistryCaller, error) {
	contract, err := bindVerificationRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerificationRegistryCaller{contract: contract}, nil
}

// NewVerificationRegistryTransactor creates a new write-only instance of VerificationRegistry, bound to a specific deployed contract.
func NewVerificationRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*VerificationRegistryTransactor, error) {
	contract, err := bindVerificationRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerificationRegistryTransactor{contract: contract}, nil
}

// NewVerificationRegistryFilterer creates a new log filterer instance of VerificationRegistry, bound to a specific deployed contract.
func NewVerificationRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*VerificationRegistryFilterer, error) {
	contract, err := bindVerificationRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerificationRegistryFilterer{contract: contract}, nil
}

// bindVerificationRegistry binds a generic wrapper to an already deployed contract.
func bindVerificationRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerificationRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerificationRegistry *VerificationRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerificationRegistry.Contract.VerificationRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerificationRegistry *VerificationRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerificationRegistry.Contract.VerificationRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerificationRegistry *VerificationRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerificationRegistry.Contract.VerificationRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerificationRegistry *VerificationRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerificationRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerificationRegistry *VerificationRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerificationRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerificationRegistry *VerificationRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerificationRegistry.Contract.contract.Transact(opts, method, params...)
}

// IssuerRegistry is a free data retrieval call binding the contract method 0x8bfc1851.
//
// Solidity: function issuerRegistry() view returns(address)
func (_VerificationRegistry *VerificationRegistryCaller) IssuerRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerificationRegistry.contract.Call(opts, &out, "issuerRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IssuerRegistry is a free data retrieval call binding the contract method 0x8bfc1851.
//
// Solidity: function issuerRegistry() view returns(address)
func (_VerificationRegistry *VerificationRegistrySession) IssuerRegistry() (common.Address, error) {
	return _VerificationRegistry.Contract.IssuerRegistry(&_VerificationRegistry.CallOpts)
}

// IssuerRegistry is a free data retrieval call binding the contract method 0x8bfc1851.
//
// Solidity: function issuerRegistry() view returns(address)
func (_VerificationRegistry *VerificationRegistryCallerSession) IssuerRegistry() (common.Address, error) {
	return _VerificationRegistry.Contract.IssuerRegistry(&_VerificationRegistry.CallOpts)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_VerificationRegistry *VerificationRegistryCaller) NftContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerificationRegistry.contract.Call(opts, &out, "nftContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_VerificationRegistry *VerificationRegistrySession) NftContract() (common.Address, error) {
	return _VerificationRegistry.Contract.NftContract(&_VerificationRegistry.CallOpts)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_VerificationRegistry *VerificationRegistryCallerSession) NftContract() (common.Address, error) {
	return _VerificationRegistry.Contract.NftContract(&_VerificationRegistry.CallOpts)
}

// TokenCounter is a free data retrieval call binding the contract method 0xd082e381.
//
// Solidity: function tokenCounter() view returns(uint256)
func (_VerificationRegistry *VerificationRegistryCaller) TokenCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VerificationRegistry.contract.Call(opts, &out, "tokenCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCounter is a free data retrieval call binding the contract method 0xd082e381.
//
// Solidity: function tokenCounter() view returns(uint256)
func (_VerificationRegistry *VerificationRegistrySession) TokenCounter() (*big.Int, error) {
	return _VerificationRegistry.Contract.TokenCounter(&_VerificationRegistry.CallOpts)
}

// TokenCounter is a free data retrieval call binding the contract method 0xd082e381.
//
// Solidity: function tokenCounter() view returns(uint256)
func (_VerificationRegistry *VerificationRegistryCallerSession) TokenCounter() (*big.Int, error) {
	return _VerificationRegistry.Contract.TokenCounter(&_VerificationRegistry.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0x1d97ce31.
//
// Solidity: function verify(uint256 tokenId, bytes32 vcHash) view returns(bool valid, address issuer, bool revoked, uint256 issuedAt)
func (_VerificationRegistry *VerificationRegistryCaller) Verify(opts *bind.CallOpts, tokenId *big.Int, vcHash [32]byte) (struct {
	Valid    bool
	Issuer   common.Address
	Revoked  bool
	IssuedAt *big.Int
}, error) {
	var out []interface{}
	err := _VerificationRegistry.contract.Call(opts, &out, "verify", tokenId, vcHash)

	outstruct := new(struct {
		Valid    bool
		Issuer   common.Address
		Revoked  bool
		IssuedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Valid = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Issuer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Revoked = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.IssuedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Verify is a free data retrieval call binding the contract method 0x1d97ce31.
//
// Solidity: function verify(uint256 tokenId, bytes32 vcHash) view returns(bool valid, address issuer, bool revoked, uint256 issuedAt)
func (_VerificationRegistry *VerificationRegistrySession) Verify(tokenId *big.Int, vcHash [32]byte) (struct {
	Valid    bool
	Issuer   common.Address
	Revoked  bool
	IssuedAt *big.Int
}, error) {
	return _VerificationRegistry.Contract.Verify(&_VerificationRegistry.CallOpts, tokenId, vcHash)
}

// Verify is a free data retrieval call binding the contract method 0x1d97ce31.
//
// Solidity: function verify(uint256 tokenId, bytes32 vcHash) view returns(bool valid, address issuer, bool revoked, uint256 issuedAt)
func (_VerificationRegistry *VerificationRegistryCallerSession) Verify(tokenId *big.Int, vcHash [32]byte) (struct {
	Valid    bool
	Issuer   common.Address
	Revoked  bool
	IssuedAt *big.Int
}, error) {
	return _VerificationRegistry.Contract.Verify(&_VerificationRegistry.CallOpts, tokenId, vcHash)
}

// RegisterVerification is a paid mutator transaction binding the contract method 0xddfccf46.
//
// Solidity: function registerVerification(bytes32 vcHash, address subject) returns(uint256)
func (_VerificationRegistry *VerificationRegistryTransactor) RegisterVerification(opts *bind.TransactOpts, vcHash [32]byte, subject common.Address) (*types.Transaction, error) {
	return _VerificationRegistry.contract.Transact(opts, "registerVerification", vcHash, subject)
}

// RegisterVerification is a paid mutator transaction binding the contract method 0xddfccf46.
//
// Solidity: function registerVerification(bytes32 vcHash, address subject) returns(uint256)
func (_VerificationRegistry *VerificationRegistrySession) RegisterVerification(vcHash [32]byte, subject common.Address) (*types.Transaction, error) {
	return _VerificationRegistry.Contract.RegisterVerification(&_VerificationRegistry.TransactOpts, vcHash, subject)
}

// RegisterVerification is a paid mutator transaction binding the contract method 0xddfccf46.
//
// Solidity: function registerVerification(bytes32 vcHash, address subject) returns(uint256)
func (_VerificationRegistry *VerificationRegistryTransactorSession) RegisterVerification(vcHash [32]byte, subject common.Address) (*types.Transaction, error) {
	return _VerificationRegistry.Contract.RegisterVerification(&_VerificationRegistry.TransactOpts, vcHash, subject)
}

// RevokeVerification is a paid mutator transaction binding the contract method 0x1571f0c6.
//
// Solidity: function revokeVerification(uint256 tokenId) returns()
func (_VerificationRegistry *VerificationRegistryTransactor) RevokeVerification(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _VerificationRegistry.contract.Transact(opts, "revokeVerification", tokenId)
}

// RevokeVerification is a paid mutator transaction binding the contract method 0x1571f0c6.
//
// Solidity: function revokeVerification(uint256 tokenId) returns()
func (_VerificationRegistry *VerificationRegistrySession) RevokeVerification(tokenId *big.Int) (*types.Transaction, error) {
	return _VerificationRegistry.Contract.RevokeVerification(&_VerificationRegistry.TransactOpts, tokenId)
}

// RevokeVerification is a paid mutator transaction binding the contract method 0x1571f0c6.
//
// Solidity: function revokeVerification(uint256 tokenId) returns()
func (_VerificationRegistry *VerificationRegistryTransactorSession) RevokeVerification(tokenId *big.Int) (*types.Transaction, error) {
	return _VerificationRegistry.Contract.RevokeVerification(&_VerificationRegistry.TransactOpts, tokenId)
}

// SetNFTContract is a paid mutator transaction binding the contract method 0xa7ccabdf.
//
// Solidity: function setNFTContract(address _nftContract) returns()
func (_VerificationRegistry *VerificationRegistryTransactor) SetNFTContract(opts *bind.TransactOpts, _nftContract common.Address) (*types.Transaction, error) {
	return _VerificationRegistry.contract.Transact(opts, "setNFTContract", _nftContract)
}

// SetNFTContract is a paid mutator transaction binding the contract method 0xa7ccabdf.
//
// Solidity: function setNFTContract(address _nftContract) returns()
func (_VerificationRegistry *VerificationRegistrySession) SetNFTContract(_nftContract common.Address) (*types.Transaction, error) {
	return _VerificationRegistry.Contract.SetNFTContract(&_VerificationRegistry.TransactOpts, _nftContract)
}

// SetNFTContract is a paid mutator transaction binding the contract method 0xa7ccabdf.
//
// Solidity: function setNFTContract(address _nftContract) returns()
func (_VerificationRegistry *VerificationRegistryTransactorSession) SetNFTContract(_nftContract common.Address) (*types.Transaction, error) {
	return _VerificationRegistry.Contract.SetNFTContract(&_VerificationRegistry.TransactOpts, _nftContract)
}

// VerificationRegistryNFTMintedIterator is returned from FilterNFTMinted and is used to iterate over the raw logs and unpacked data for NFTMinted events raised by the VerificationRegistry contract.
type VerificationRegistryNFTMintedIterator struct {
	Event *VerificationRegistryNFTMinted // Event containing the contract specifics and raw log

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
func (it *VerificationRegistryNFTMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerificationRegistryNFTMinted)
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
		it.Event = new(VerificationRegistryNFTMinted)
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
func (it *VerificationRegistryNFTMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerificationRegistryNFTMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerificationRegistryNFTMinted represents a NFTMinted event raised by the VerificationRegistry contract.
type VerificationRegistryNFTMinted struct {
	TokenId *big.Int
	Owner   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNFTMinted is a free log retrieval operation binding the contract event 0x0176f203df400d7bd5f1b1c9ef36c16709bf3b5d9fd35f000a6bae32393f66c3.
//
// Solidity: event NFTMinted(uint256 indexed tokenId, address indexed owner)
func (_VerificationRegistry *VerificationRegistryFilterer) FilterNFTMinted(opts *bind.FilterOpts, tokenId []*big.Int, owner []common.Address) (*VerificationRegistryNFTMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VerificationRegistry.contract.FilterLogs(opts, "NFTMinted", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &VerificationRegistryNFTMintedIterator{contract: _VerificationRegistry.contract, event: "NFTMinted", logs: logs, sub: sub}, nil
}

// WatchNFTMinted is a free log subscription operation binding the contract event 0x0176f203df400d7bd5f1b1c9ef36c16709bf3b5d9fd35f000a6bae32393f66c3.
//
// Solidity: event NFTMinted(uint256 indexed tokenId, address indexed owner)
func (_VerificationRegistry *VerificationRegistryFilterer) WatchNFTMinted(opts *bind.WatchOpts, sink chan<- *VerificationRegistryNFTMinted, tokenId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VerificationRegistry.contract.WatchLogs(opts, "NFTMinted", tokenIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerificationRegistryNFTMinted)
				if err := _VerificationRegistry.contract.UnpackLog(event, "NFTMinted", log); err != nil {
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

// ParseNFTMinted is a log parse operation binding the contract event 0x0176f203df400d7bd5f1b1c9ef36c16709bf3b5d9fd35f000a6bae32393f66c3.
//
// Solidity: event NFTMinted(uint256 indexed tokenId, address indexed owner)
func (_VerificationRegistry *VerificationRegistryFilterer) ParseNFTMinted(log types.Log) (*VerificationRegistryNFTMinted, error) {
	event := new(VerificationRegistryNFTMinted)
	if err := _VerificationRegistry.contract.UnpackLog(event, "NFTMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerificationRegistryVerificationRegisteredIterator is returned from FilterVerificationRegistered and is used to iterate over the raw logs and unpacked data for VerificationRegistered events raised by the VerificationRegistry contract.
type VerificationRegistryVerificationRegisteredIterator struct {
	Event *VerificationRegistryVerificationRegistered // Event containing the contract specifics and raw log

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
func (it *VerificationRegistryVerificationRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerificationRegistryVerificationRegistered)
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
		it.Event = new(VerificationRegistryVerificationRegistered)
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
func (it *VerificationRegistryVerificationRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerificationRegistryVerificationRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerificationRegistryVerificationRegistered represents a VerificationRegistered event raised by the VerificationRegistry contract.
type VerificationRegistryVerificationRegistered struct {
	TokenId *big.Int
	Issuer  common.Address
	VcHash  [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVerificationRegistered is a free log retrieval operation binding the contract event 0x0e5026488e8f9bb8126eae97bb91a9b730c9e194938b6f06a2c1309e19671c86.
//
// Solidity: event VerificationRegistered(uint256 indexed tokenId, address indexed issuer, bytes32 vcHash)
func (_VerificationRegistry *VerificationRegistryFilterer) FilterVerificationRegistered(opts *bind.FilterOpts, tokenId []*big.Int, issuer []common.Address) (*VerificationRegistryVerificationRegisteredIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _VerificationRegistry.contract.FilterLogs(opts, "VerificationRegistered", tokenIdRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &VerificationRegistryVerificationRegisteredIterator{contract: _VerificationRegistry.contract, event: "VerificationRegistered", logs: logs, sub: sub}, nil
}

// WatchVerificationRegistered is a free log subscription operation binding the contract event 0x0e5026488e8f9bb8126eae97bb91a9b730c9e194938b6f06a2c1309e19671c86.
//
// Solidity: event VerificationRegistered(uint256 indexed tokenId, address indexed issuer, bytes32 vcHash)
func (_VerificationRegistry *VerificationRegistryFilterer) WatchVerificationRegistered(opts *bind.WatchOpts, sink chan<- *VerificationRegistryVerificationRegistered, tokenId []*big.Int, issuer []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _VerificationRegistry.contract.WatchLogs(opts, "VerificationRegistered", tokenIdRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerificationRegistryVerificationRegistered)
				if err := _VerificationRegistry.contract.UnpackLog(event, "VerificationRegistered", log); err != nil {
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

// ParseVerificationRegistered is a log parse operation binding the contract event 0x0e5026488e8f9bb8126eae97bb91a9b730c9e194938b6f06a2c1309e19671c86.
//
// Solidity: event VerificationRegistered(uint256 indexed tokenId, address indexed issuer, bytes32 vcHash)
func (_VerificationRegistry *VerificationRegistryFilterer) ParseVerificationRegistered(log types.Log) (*VerificationRegistryVerificationRegistered, error) {
	event := new(VerificationRegistryVerificationRegistered)
	if err := _VerificationRegistry.contract.UnpackLog(event, "VerificationRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerificationRegistryVerificationRevokedIterator is returned from FilterVerificationRevoked and is used to iterate over the raw logs and unpacked data for VerificationRevoked events raised by the VerificationRegistry contract.
type VerificationRegistryVerificationRevokedIterator struct {
	Event *VerificationRegistryVerificationRevoked // Event containing the contract specifics and raw log

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
func (it *VerificationRegistryVerificationRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerificationRegistryVerificationRevoked)
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
		it.Event = new(VerificationRegistryVerificationRevoked)
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
func (it *VerificationRegistryVerificationRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerificationRegistryVerificationRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerificationRegistryVerificationRevoked represents a VerificationRevoked event raised by the VerificationRegistry contract.
type VerificationRegistryVerificationRevoked struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVerificationRevoked is a free log retrieval operation binding the contract event 0xea6232c80090e591dc995461c6943381a4246f004efce58e311f6dae40e3f83e.
//
// Solidity: event VerificationRevoked(uint256 indexed tokenId)
func (_VerificationRegistry *VerificationRegistryFilterer) FilterVerificationRevoked(opts *bind.FilterOpts, tokenId []*big.Int) (*VerificationRegistryVerificationRevokedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VerificationRegistry.contract.FilterLogs(opts, "VerificationRevoked", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &VerificationRegistryVerificationRevokedIterator{contract: _VerificationRegistry.contract, event: "VerificationRevoked", logs: logs, sub: sub}, nil
}

// WatchVerificationRevoked is a free log subscription operation binding the contract event 0xea6232c80090e591dc995461c6943381a4246f004efce58e311f6dae40e3f83e.
//
// Solidity: event VerificationRevoked(uint256 indexed tokenId)
func (_VerificationRegistry *VerificationRegistryFilterer) WatchVerificationRevoked(opts *bind.WatchOpts, sink chan<- *VerificationRegistryVerificationRevoked, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _VerificationRegistry.contract.WatchLogs(opts, "VerificationRevoked", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerificationRegistryVerificationRevoked)
				if err := _VerificationRegistry.contract.UnpackLog(event, "VerificationRevoked", log); err != nil {
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

// ParseVerificationRevoked is a log parse operation binding the contract event 0xea6232c80090e591dc995461c6943381a4246f004efce58e311f6dae40e3f83e.
//
// Solidity: event VerificationRevoked(uint256 indexed tokenId)
func (_VerificationRegistry *VerificationRegistryFilterer) ParseVerificationRevoked(log types.Log) (*VerificationRegistryVerificationRevoked, error) {
	event := new(VerificationRegistryVerificationRevoked)
	if err := _VerificationRegistry.contract.UnpackLog(event, "VerificationRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
