// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package modules

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

// ADKStakeMetaData contains all meta data concerning the ADKStake contract.
var ADKStakeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EVT_ADKSTAKE\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EVT_ADKUNSTAKE\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GetStakedCurrentAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetWaitCurrentAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"SetMinStakePariod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"SetOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"StakedAmountByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"StakedBlockByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TotalADKStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountToUnstake\",\"type\":\"uint256\"}],\"name\":\"Unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lock_period_blocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"4a3eaab8": "GetStakedCurrentAddress()",
		"140cf07b": "GetWaitCurrentAddress()",
		"62d7ed72": "SetMinStakePariod(uint256)",
		"167d3e9c": "SetOwner(address)",
		"de20bc92": "Stake()",
		"502c3fbf": "StakedAmountByAddress(address)",
		"5f722dde": "StakedBlockByAddress(address)",
		"6944a75f": "TotalADKStaked()",
		"f1f1db1e": "Unstake(uint256)",
		"7f092ad4": "lock_period_blocks()",
		"8da5cb5b": "owner()",
	},
	Bin: "0x60806040526004805461ffff1916905534801561001b57600080fd5b50600280546001600160a01b0319163317905561168060035561079c806100436000396000f3fe6080604052600436106100a05760003560e01c806362d7ed721161006457806362d7ed72146101785780636944a75f146101985780637f092ad4146101ab5780638da5cb5b146101c1578063de20bc92146101f9578063f1f1db1e1461020157600080fd5b8063140cf07b146100b4578063167d3e9c146100dc5780634a3eaab8146100fc578063502c3fbf1461011e5780635f722dde1461014b57600080fd5b366100af576100ad610221565b005b600080fd5b3480156100c057600080fd5b506100c961035b565b6040519081526020015b60405180910390f35b3480156100e857600080fd5b506100ad6100f73660046106d8565b6103b4565b34801561010857600080fd5b50336000908152602081905260409020546100c9565b34801561012a57600080fd5b506100c96101393660046106d8565b60006020819052908152604090205481565b34801561015757600080fd5b506100c96101663660046106d8565b60016020526000908152604090205481565b34801561018457600080fd5b506100ad610193366004610708565b610430565b3480156101a457600080fd5b50476100c9565b3480156101b757600080fd5b506100c960035481565b3480156101cd57600080fd5b506002546101e1906001600160a01b031681565b6040516001600160a01b0390911681526020016100d3565b6100ad610221565b34801561020d57600080fd5b506100ad61021c366004610708565b61048f565b60045460ff16156102855760405162461bcd60e51b815260206004820152602360248201527f726563656976652f5374616b6528293a206e6f207265656e74727920616c6c6f6044820152621dd95960ea1b60648201526084015b60405180910390fd5b6004805460ff19166001179055346102d75760405162461bcd60e51b8152602060048201526015602482015274041444b2076616c75652063616e6e6f74206265203605c1b604482015260640161027c565b33600090815260208190526040812080543492906102f6908490610721565b90915550503360008181526001602090815260409182902043908190558251938452349184019190915282820152517fb7e47cb78d92678b01231ce0c0d4cd490881f5138ae6e315f69b121265de306f9181900360600190a16004805460ff19169055565b600354336000908152600160205260408120549091439161037c9190610721565b116103875750600090565b6003543360009081526001602052604090205443916103a591610721565b6103af9190610739565b905090565b6002546001600160a01b0316331461040e5760405162461bcd60e51b815260206004820152601c60248201527f4e4f542043414c4c454420425920434f4e5452414354204f574e455200000000604482015260640161027c565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b6002546001600160a01b0316331461048a5760405162461bcd60e51b815260206004820152601c60248201527f4e4f542043414c4c454420425920434f4e5452414354204f574e455200000000604482015260640161027c565b600355565b600454610100900460ff16156104e75760405162461bcd60e51b815260206004820152601d60248201527f556e7374616b6528293a206e6f207265656e74727920616c6c6f776564000000604482015260640161027c565b6004805461ff0019166101001790558061054d5760405162461bcd60e51b815260206004820152602160248201527f556e7374616b6528293a20696e76616c696420616d6f756e7420636c61696d656044820152601960fa1b606482015260840161027c565b336000908152602081905260409020548111156105bd5760405162461bcd60e51b815260206004820152602860248201527f556e7374616b6528293a206d73672e73656e64657220686173206e6f7468696e60448201526719c81cdd185ad95960c21b606482015260840161027c565b6003543360009081526001602052604090205443916105db91610721565b11156106395760405162461bcd60e51b815260206004820152602760248201527f556e7374616b6528293a20756e7374616b6520706572696f64206e6f7420636f6044820152661b5c1b195d195960ca1b606482015260840161027c565b3360009081526020819052604081208054839290610658908490610739565b9091555050604051339082156108fc029083906000818181858888f1935050505015801561068a573d6000803e3d6000fd5b506040805133815260208101839052438183015290517fc62930930d8b10f60b293f2b1cd1d55d976c2c5ad273548d546e11d08da99f0a9181900360600190a1506004805461ff0019169055565b6000602082840312156106ea57600080fd5b81356001600160a01b038116811461070157600080fd5b9392505050565b60006020828403121561071a57600080fd5b5035919050565b6000821982111561073457610734610750565b500190565b60008282101561074b5761074b610750565b500390565b634e487b7160e01b600052601160045260246000fdfea26469706673582212207dfc49f1d8c52224757140b99d48f4c1c4d30a65dbfb1b61b6ac68dc14a2ecbb64736f6c63430008070033",
}

// ADKStakeABI is the input ABI used to generate the binding from.
// Deprecated: Use ADKStakeMetaData.ABI instead.
var ADKStakeABI = ADKStakeMetaData.ABI

// Deprecated: Use ADKStakeMetaData.Sigs instead.
// ADKStakeFuncSigs maps the 4-byte function signature to its string representation.
var ADKStakeFuncSigs = ADKStakeMetaData.Sigs

// ADKStakeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ADKStakeMetaData.Bin instead.
var ADKStakeBin = ADKStakeMetaData.Bin

// DeployADKStake deploys a new Ethereum contract, binding an instance of ADKStake to it.
func DeployADKStake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ADKStake, error) {
	parsed, err := ADKStakeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ADKStakeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ADKStake{ADKStakeCaller: ADKStakeCaller{contract: contract}, ADKStakeTransactor: ADKStakeTransactor{contract: contract}, ADKStakeFilterer: ADKStakeFilterer{contract: contract}}, nil
}

// ADKStake is an auto generated Go binding around an Ethereum contract.
type ADKStake struct {
	ADKStakeCaller     // Read-only binding to the contract
	ADKStakeTransactor // Write-only binding to the contract
	ADKStakeFilterer   // Log filterer for contract events
}

// ADKStakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ADKStakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ADKStakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ADKStakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ADKStakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ADKStakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ADKStakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ADKStakeSession struct {
	Contract     *ADKStake         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ADKStakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ADKStakeCallerSession struct {
	Contract *ADKStakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ADKStakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ADKStakeTransactorSession struct {
	Contract     *ADKStakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ADKStakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ADKStakeRaw struct {
	Contract *ADKStake // Generic contract binding to access the raw methods on
}

// ADKStakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ADKStakeCallerRaw struct {
	Contract *ADKStakeCaller // Generic read-only contract binding to access the raw methods on
}

// ADKStakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ADKStakeTransactorRaw struct {
	Contract *ADKStakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewADKStake creates a new instance of ADKStake, bound to a specific deployed contract.
func NewADKStake(address common.Address, backend bind.ContractBackend) (*ADKStake, error) {
	contract, err := bindADKStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ADKStake{ADKStakeCaller: ADKStakeCaller{contract: contract}, ADKStakeTransactor: ADKStakeTransactor{contract: contract}, ADKStakeFilterer: ADKStakeFilterer{contract: contract}}, nil
}

// NewADKStakeCaller creates a new read-only instance of ADKStake, bound to a specific deployed contract.
func NewADKStakeCaller(address common.Address, caller bind.ContractCaller) (*ADKStakeCaller, error) {
	contract, err := bindADKStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ADKStakeCaller{contract: contract}, nil
}

// NewADKStakeTransactor creates a new write-only instance of ADKStake, bound to a specific deployed contract.
func NewADKStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*ADKStakeTransactor, error) {
	contract, err := bindADKStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ADKStakeTransactor{contract: contract}, nil
}

// NewADKStakeFilterer creates a new log filterer instance of ADKStake, bound to a specific deployed contract.
func NewADKStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*ADKStakeFilterer, error) {
	contract, err := bindADKStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ADKStakeFilterer{contract: contract}, nil
}

// bindADKStake binds a generic wrapper to an already deployed contract.
func bindADKStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ADKStakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ADKStake *ADKStakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ADKStake.Contract.ADKStakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ADKStake *ADKStakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ADKStake.Contract.ADKStakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ADKStake *ADKStakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ADKStake.Contract.ADKStakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ADKStake *ADKStakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ADKStake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ADKStake *ADKStakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ADKStake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ADKStake *ADKStakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ADKStake.Contract.contract.Transact(opts, method, params...)
}

// GetStakedCurrentAddress is a free data retrieval call binding the contract method 0x4a3eaab8.
//
// Solidity: function GetStakedCurrentAddress() view returns(uint256)
func (_ADKStake *ADKStakeCaller) GetStakedCurrentAddress(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ADKStake.contract.Call(opts, &out, "GetStakedCurrentAddress")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakedCurrentAddress is a free data retrieval call binding the contract method 0x4a3eaab8.
//
// Solidity: function GetStakedCurrentAddress() view returns(uint256)
func (_ADKStake *ADKStakeSession) GetStakedCurrentAddress() (*big.Int, error) {
	return _ADKStake.Contract.GetStakedCurrentAddress(&_ADKStake.CallOpts)
}

// GetStakedCurrentAddress is a free data retrieval call binding the contract method 0x4a3eaab8.
//
// Solidity: function GetStakedCurrentAddress() view returns(uint256)
func (_ADKStake *ADKStakeCallerSession) GetStakedCurrentAddress() (*big.Int, error) {
	return _ADKStake.Contract.GetStakedCurrentAddress(&_ADKStake.CallOpts)
}

// GetWaitCurrentAddress is a free data retrieval call binding the contract method 0x140cf07b.
//
// Solidity: function GetWaitCurrentAddress() view returns(uint256)
func (_ADKStake *ADKStakeCaller) GetWaitCurrentAddress(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ADKStake.contract.Call(opts, &out, "GetWaitCurrentAddress")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWaitCurrentAddress is a free data retrieval call binding the contract method 0x140cf07b.
//
// Solidity: function GetWaitCurrentAddress() view returns(uint256)
func (_ADKStake *ADKStakeSession) GetWaitCurrentAddress() (*big.Int, error) {
	return _ADKStake.Contract.GetWaitCurrentAddress(&_ADKStake.CallOpts)
}

// GetWaitCurrentAddress is a free data retrieval call binding the contract method 0x140cf07b.
//
// Solidity: function GetWaitCurrentAddress() view returns(uint256)
func (_ADKStake *ADKStakeCallerSession) GetWaitCurrentAddress() (*big.Int, error) {
	return _ADKStake.Contract.GetWaitCurrentAddress(&_ADKStake.CallOpts)
}

// StakedAmountByAddress is a free data retrieval call binding the contract method 0x502c3fbf.
//
// Solidity: function StakedAmountByAddress(address ) view returns(uint256)
func (_ADKStake *ADKStakeCaller) StakedAmountByAddress(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ADKStake.contract.Call(opts, &out, "StakedAmountByAddress", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedAmountByAddress is a free data retrieval call binding the contract method 0x502c3fbf.
//
// Solidity: function StakedAmountByAddress(address ) view returns(uint256)
func (_ADKStake *ADKStakeSession) StakedAmountByAddress(arg0 common.Address) (*big.Int, error) {
	return _ADKStake.Contract.StakedAmountByAddress(&_ADKStake.CallOpts, arg0)
}

// StakedAmountByAddress is a free data retrieval call binding the contract method 0x502c3fbf.
//
// Solidity: function StakedAmountByAddress(address ) view returns(uint256)
func (_ADKStake *ADKStakeCallerSession) StakedAmountByAddress(arg0 common.Address) (*big.Int, error) {
	return _ADKStake.Contract.StakedAmountByAddress(&_ADKStake.CallOpts, arg0)
}

// StakedBlockByAddress is a free data retrieval call binding the contract method 0x5f722dde.
//
// Solidity: function StakedBlockByAddress(address ) view returns(uint256)
func (_ADKStake *ADKStakeCaller) StakedBlockByAddress(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ADKStake.contract.Call(opts, &out, "StakedBlockByAddress", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedBlockByAddress is a free data retrieval call binding the contract method 0x5f722dde.
//
// Solidity: function StakedBlockByAddress(address ) view returns(uint256)
func (_ADKStake *ADKStakeSession) StakedBlockByAddress(arg0 common.Address) (*big.Int, error) {
	return _ADKStake.Contract.StakedBlockByAddress(&_ADKStake.CallOpts, arg0)
}

// StakedBlockByAddress is a free data retrieval call binding the contract method 0x5f722dde.
//
// Solidity: function StakedBlockByAddress(address ) view returns(uint256)
func (_ADKStake *ADKStakeCallerSession) StakedBlockByAddress(arg0 common.Address) (*big.Int, error) {
	return _ADKStake.Contract.StakedBlockByAddress(&_ADKStake.CallOpts, arg0)
}

// TotalADKStaked is a free data retrieval call binding the contract method 0x6944a75f.
//
// Solidity: function TotalADKStaked() view returns(uint256)
func (_ADKStake *ADKStakeCaller) TotalADKStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ADKStake.contract.Call(opts, &out, "TotalADKStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalADKStaked is a free data retrieval call binding the contract method 0x6944a75f.
//
// Solidity: function TotalADKStaked() view returns(uint256)
func (_ADKStake *ADKStakeSession) TotalADKStaked() (*big.Int, error) {
	return _ADKStake.Contract.TotalADKStaked(&_ADKStake.CallOpts)
}

// TotalADKStaked is a free data retrieval call binding the contract method 0x6944a75f.
//
// Solidity: function TotalADKStaked() view returns(uint256)
func (_ADKStake *ADKStakeCallerSession) TotalADKStaked() (*big.Int, error) {
	return _ADKStake.Contract.TotalADKStaked(&_ADKStake.CallOpts)
}

// LockPeriodBlocks is a free data retrieval call binding the contract method 0x7f092ad4.
//
// Solidity: function lock_period_blocks() view returns(uint256)
func (_ADKStake *ADKStakeCaller) LockPeriodBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ADKStake.contract.Call(opts, &out, "lock_period_blocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockPeriodBlocks is a free data retrieval call binding the contract method 0x7f092ad4.
//
// Solidity: function lock_period_blocks() view returns(uint256)
func (_ADKStake *ADKStakeSession) LockPeriodBlocks() (*big.Int, error) {
	return _ADKStake.Contract.LockPeriodBlocks(&_ADKStake.CallOpts)
}

// LockPeriodBlocks is a free data retrieval call binding the contract method 0x7f092ad4.
//
// Solidity: function lock_period_blocks() view returns(uint256)
func (_ADKStake *ADKStakeCallerSession) LockPeriodBlocks() (*big.Int, error) {
	return _ADKStake.Contract.LockPeriodBlocks(&_ADKStake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ADKStake *ADKStakeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ADKStake.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ADKStake *ADKStakeSession) Owner() (common.Address, error) {
	return _ADKStake.Contract.Owner(&_ADKStake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ADKStake *ADKStakeCallerSession) Owner() (common.Address, error) {
	return _ADKStake.Contract.Owner(&_ADKStake.CallOpts)
}

// SetMinStakePariod is a paid mutator transaction binding the contract method 0x62d7ed72.
//
// Solidity: function SetMinStakePariod(uint256 period) returns()
func (_ADKStake *ADKStakeTransactor) SetMinStakePariod(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _ADKStake.contract.Transact(opts, "SetMinStakePariod", period)
}

// SetMinStakePariod is a paid mutator transaction binding the contract method 0x62d7ed72.
//
// Solidity: function SetMinStakePariod(uint256 period) returns()
func (_ADKStake *ADKStakeSession) SetMinStakePariod(period *big.Int) (*types.Transaction, error) {
	return _ADKStake.Contract.SetMinStakePariod(&_ADKStake.TransactOpts, period)
}

// SetMinStakePariod is a paid mutator transaction binding the contract method 0x62d7ed72.
//
// Solidity: function SetMinStakePariod(uint256 period) returns()
func (_ADKStake *ADKStakeTransactorSession) SetMinStakePariod(period *big.Int) (*types.Transaction, error) {
	return _ADKStake.Contract.SetMinStakePariod(&_ADKStake.TransactOpts, period)
}

// SetOwner is a paid mutator transaction binding the contract method 0x167d3e9c.
//
// Solidity: function SetOwner(address _newOwner) returns()
func (_ADKStake *ADKStakeTransactor) SetOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ADKStake.contract.Transact(opts, "SetOwner", _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x167d3e9c.
//
// Solidity: function SetOwner(address _newOwner) returns()
func (_ADKStake *ADKStakeSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _ADKStake.Contract.SetOwner(&_ADKStake.TransactOpts, _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x167d3e9c.
//
// Solidity: function SetOwner(address _newOwner) returns()
func (_ADKStake *ADKStakeTransactorSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _ADKStake.Contract.SetOwner(&_ADKStake.TransactOpts, _newOwner)
}

// Stake is a paid mutator transaction binding the contract method 0xde20bc92.
//
// Solidity: function Stake() payable returns()
func (_ADKStake *ADKStakeTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ADKStake.contract.Transact(opts, "Stake")
}

// Stake is a paid mutator transaction binding the contract method 0xde20bc92.
//
// Solidity: function Stake() payable returns()
func (_ADKStake *ADKStakeSession) Stake() (*types.Transaction, error) {
	return _ADKStake.Contract.Stake(&_ADKStake.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0xde20bc92.
//
// Solidity: function Stake() payable returns()
func (_ADKStake *ADKStakeTransactorSession) Stake() (*types.Transaction, error) {
	return _ADKStake.Contract.Stake(&_ADKStake.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0xf1f1db1e.
//
// Solidity: function Unstake(uint256 _amountToUnstake) returns()
func (_ADKStake *ADKStakeTransactor) Unstake(opts *bind.TransactOpts, _amountToUnstake *big.Int) (*types.Transaction, error) {
	return _ADKStake.contract.Transact(opts, "Unstake", _amountToUnstake)
}

// Unstake is a paid mutator transaction binding the contract method 0xf1f1db1e.
//
// Solidity: function Unstake(uint256 _amountToUnstake) returns()
func (_ADKStake *ADKStakeSession) Unstake(_amountToUnstake *big.Int) (*types.Transaction, error) {
	return _ADKStake.Contract.Unstake(&_ADKStake.TransactOpts, _amountToUnstake)
}

// Unstake is a paid mutator transaction binding the contract method 0xf1f1db1e.
//
// Solidity: function Unstake(uint256 _amountToUnstake) returns()
func (_ADKStake *ADKStakeTransactorSession) Unstake(_amountToUnstake *big.Int) (*types.Transaction, error) {
	return _ADKStake.Contract.Unstake(&_ADKStake.TransactOpts, _amountToUnstake)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ADKStake *ADKStakeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ADKStake.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ADKStake *ADKStakeSession) Receive() (*types.Transaction, error) {
	return _ADKStake.Contract.Receive(&_ADKStake.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ADKStake *ADKStakeTransactorSession) Receive() (*types.Transaction, error) {
	return _ADKStake.Contract.Receive(&_ADKStake.TransactOpts)
}

// ADKStakeEVTADKSTAKEIterator is returned from FilterEVTADKSTAKE and is used to iterate over the raw logs and unpacked data for EVTADKSTAKE events raised by the ADKStake contract.
type ADKStakeEVTADKSTAKEIterator struct {
	Event *ADKStakeEVTADKSTAKE // Event containing the contract specifics and raw log

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
func (it *ADKStakeEVTADKSTAKEIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ADKStakeEVTADKSTAKE)
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
		it.Event = new(ADKStakeEVTADKSTAKE)
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
func (it *ADKStakeEVTADKSTAKEIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ADKStakeEVTADKSTAKEIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ADKStakeEVTADKSTAKE represents a EVTADKSTAKE event raised by the ADKStake contract.
type ADKStakeEVTADKSTAKE struct {
	Arg0 common.Address
	Arg1 *big.Int
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEVTADKSTAKE is a free log retrieval operation binding the contract event 0xb7e47cb78d92678b01231ce0c0d4cd490881f5138ae6e315f69b121265de306f.
//
// Solidity: event EVT_ADKSTAKE(address arg0, uint256 arg1, uint256 arg2)
func (_ADKStake *ADKStakeFilterer) FilterEVTADKSTAKE(opts *bind.FilterOpts) (*ADKStakeEVTADKSTAKEIterator, error) {

	logs, sub, err := _ADKStake.contract.FilterLogs(opts, "EVT_ADKSTAKE")
	if err != nil {
		return nil, err
	}
	return &ADKStakeEVTADKSTAKEIterator{contract: _ADKStake.contract, event: "EVT_ADKSTAKE", logs: logs, sub: sub}, nil
}

// WatchEVTADKSTAKE is a free log subscription operation binding the contract event 0xb7e47cb78d92678b01231ce0c0d4cd490881f5138ae6e315f69b121265de306f.
//
// Solidity: event EVT_ADKSTAKE(address arg0, uint256 arg1, uint256 arg2)
func (_ADKStake *ADKStakeFilterer) WatchEVTADKSTAKE(opts *bind.WatchOpts, sink chan<- *ADKStakeEVTADKSTAKE) (event.Subscription, error) {

	logs, sub, err := _ADKStake.contract.WatchLogs(opts, "EVT_ADKSTAKE")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ADKStakeEVTADKSTAKE)
				if err := _ADKStake.contract.UnpackLog(event, "EVT_ADKSTAKE", log); err != nil {
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

// ParseEVTADKSTAKE is a log parse operation binding the contract event 0xb7e47cb78d92678b01231ce0c0d4cd490881f5138ae6e315f69b121265de306f.
//
// Solidity: event EVT_ADKSTAKE(address arg0, uint256 arg1, uint256 arg2)
func (_ADKStake *ADKStakeFilterer) ParseEVTADKSTAKE(log types.Log) (*ADKStakeEVTADKSTAKE, error) {
	event := new(ADKStakeEVTADKSTAKE)
	if err := _ADKStake.contract.UnpackLog(event, "EVT_ADKSTAKE", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ADKStakeEVTADKUNSTAKEIterator is returned from FilterEVTADKUNSTAKE and is used to iterate over the raw logs and unpacked data for EVTADKUNSTAKE events raised by the ADKStake contract.
type ADKStakeEVTADKUNSTAKEIterator struct {
	Event *ADKStakeEVTADKUNSTAKE // Event containing the contract specifics and raw log

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
func (it *ADKStakeEVTADKUNSTAKEIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ADKStakeEVTADKUNSTAKE)
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
		it.Event = new(ADKStakeEVTADKUNSTAKE)
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
func (it *ADKStakeEVTADKUNSTAKEIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ADKStakeEVTADKUNSTAKEIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ADKStakeEVTADKUNSTAKE represents a EVTADKUNSTAKE event raised by the ADKStake contract.
type ADKStakeEVTADKUNSTAKE struct {
	Arg0 common.Address
	Arg1 *big.Int
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEVTADKUNSTAKE is a free log retrieval operation binding the contract event 0xc62930930d8b10f60b293f2b1cd1d55d976c2c5ad273548d546e11d08da99f0a.
//
// Solidity: event EVT_ADKUNSTAKE(address arg0, uint256 arg1, uint256 arg2)
func (_ADKStake *ADKStakeFilterer) FilterEVTADKUNSTAKE(opts *bind.FilterOpts) (*ADKStakeEVTADKUNSTAKEIterator, error) {

	logs, sub, err := _ADKStake.contract.FilterLogs(opts, "EVT_ADKUNSTAKE")
	if err != nil {
		return nil, err
	}
	return &ADKStakeEVTADKUNSTAKEIterator{contract: _ADKStake.contract, event: "EVT_ADKUNSTAKE", logs: logs, sub: sub}, nil
}

// WatchEVTADKUNSTAKE is a free log subscription operation binding the contract event 0xc62930930d8b10f60b293f2b1cd1d55d976c2c5ad273548d546e11d08da99f0a.
//
// Solidity: event EVT_ADKUNSTAKE(address arg0, uint256 arg1, uint256 arg2)
func (_ADKStake *ADKStakeFilterer) WatchEVTADKUNSTAKE(opts *bind.WatchOpts, sink chan<- *ADKStakeEVTADKUNSTAKE) (event.Subscription, error) {

	logs, sub, err := _ADKStake.contract.WatchLogs(opts, "EVT_ADKUNSTAKE")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ADKStakeEVTADKUNSTAKE)
				if err := _ADKStake.contract.UnpackLog(event, "EVT_ADKUNSTAKE", log); err != nil {
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

// ParseEVTADKUNSTAKE is a log parse operation binding the contract event 0xc62930930d8b10f60b293f2b1cd1d55d976c2c5ad273548d546e11d08da99f0a.
//
// Solidity: event EVT_ADKUNSTAKE(address arg0, uint256 arg1, uint256 arg2)
func (_ADKStake *ADKStakeFilterer) ParseEVTADKUNSTAKE(log types.Log) (*ADKStakeEVTADKUNSTAKE, error) {
	event := new(ADKStakeEVTADKUNSTAKE)
	if err := _ADKStake.contract.UnpackLog(event, "EVT_ADKUNSTAKE", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
