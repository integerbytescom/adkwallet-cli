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

// ADKTokenMetaData contains all meta data concerning the ADKToken contract.
var ADKTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_meshaddr\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"_value\",\"type\":\"int256\"}],\"name\":\"AZ9Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_meshaddr\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"_value\",\"type\":\"int256\"}],\"name\":\"GENESISTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethAddr\",\"type\":\"address\"}],\"name\":\"ADDR_TO_AZ9\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ADKTransactionsContract\",\"outputs\":[{\"internalType\":\"contractADKTransactions\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AGSClaimContract\",\"outputs\":[{\"internalType\":\"contractAGSClaim\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"adkString\",\"type\":\"string\"}],\"name\":\"AZ9_TO_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_adkAddr\",\"type\":\"string\"}],\"name\":\"AZ9balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"ethAddr\",\"type\":\"bytes\"}],\"name\":\"BADDR_TO_AZ9\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adkgo_genesis_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_meshaddr\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"_value\",\"type\":\"int256\"}],\"name\":\"genesisTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"known_addresses\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"linked_list_all_balances\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_meshaddr\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"_value\",\"type\":\"int256\"}],\"name\":\"meshTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_meshaddr\",\"type\":\"string\"}],\"name\":\"registerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b79c88ec": "ADDR_TO_AZ9(address)",
		"8ab04a78": "ADKTransactionsContract()",
		"41ef2686": "AGSClaimContract()",
		"a205a8a9": "AZ9_TO_ADDR(string)",
		"0e7fdbb1": "AZ9balanceOf(string)",
		"78f7c041": "BADDR_TO_AZ9(bytes)",
		"c4c7e151": "adkgo_genesis_address()",
		"dd62ed3e": "allowance(address,address)",
		"5c658165": "allowed(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"27e235e3": "balances(address)",
		"313ce567": "decimals()",
		"a85cfb66": "genesisTransaction(string,int256)",
		"3f43347a": "known_addresses(address)",
		"2a5d2cf7": "linked_list_all_balances(address)",
		"8756ee19": "meshTransaction(string,int256)",
		"06fdde03": "name()",
		"bb1bd28f": "registerAddress(string)",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
	Bin: "0x00",
}

// ADKTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ADKTokenMetaData.ABI instead.
var ADKTokenABI = ADKTokenMetaData.ABI

// Deprecated: Use ADKTokenMetaData.Sigs instead.
// ADKTokenFuncSigs maps the 4-byte function signature to its string representation.
var ADKTokenFuncSigs = ADKTokenMetaData.Sigs

// ADKTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ADKTokenMetaData.Bin instead.
var ADKTokenBin = ADKTokenMetaData.Bin

// DeployADKToken deploys a new Ethereum contract, binding an instance of ADKToken to it.
func DeployADKToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ADKToken, error) {
	parsed, err := ADKTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ADKTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ADKToken{ADKTokenCaller: ADKTokenCaller{contract: contract}, ADKTokenTransactor: ADKTokenTransactor{contract: contract}, ADKTokenFilterer: ADKTokenFilterer{contract: contract}}, nil
}

// ADKToken is an auto generated Go binding around an Ethereum contract.
type ADKToken struct {
	ADKTokenCaller     // Read-only binding to the contract
	ADKTokenTransactor // Write-only binding to the contract
	ADKTokenFilterer   // Log filterer for contract events
}

// ADKTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ADKTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ADKTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ADKTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ADKTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ADKTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ADKTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ADKTokenSession struct {
	Contract     *ADKToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ADKTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ADKTokenCallerSession struct {
	Contract *ADKTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ADKTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ADKTokenTransactorSession struct {
	Contract     *ADKTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ADKTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ADKTokenRaw struct {
	Contract *ADKToken // Generic contract binding to access the raw methods on
}

// ADKTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ADKTokenCallerRaw struct {
	Contract *ADKTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ADKTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ADKTokenTransactorRaw struct {
	Contract *ADKTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewADKToken creates a new instance of ADKToken, bound to a specific deployed contract.
func NewADKToken(address common.Address, backend bind.ContractBackend) (*ADKToken, error) {
	contract, err := bindADKToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ADKToken{ADKTokenCaller: ADKTokenCaller{contract: contract}, ADKTokenTransactor: ADKTokenTransactor{contract: contract}, ADKTokenFilterer: ADKTokenFilterer{contract: contract}}, nil
}

// NewADKTokenCaller creates a new read-only instance of ADKToken, bound to a specific deployed contract.
func NewADKTokenCaller(address common.Address, caller bind.ContractCaller) (*ADKTokenCaller, error) {
	contract, err := bindADKToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ADKTokenCaller{contract: contract}, nil
}

// NewADKTokenTransactor creates a new write-only instance of ADKToken, bound to a specific deployed contract.
func NewADKTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ADKTokenTransactor, error) {
	contract, err := bindADKToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ADKTokenTransactor{contract: contract}, nil
}

// NewADKTokenFilterer creates a new log filterer instance of ADKToken, bound to a specific deployed contract.
func NewADKTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ADKTokenFilterer, error) {
	contract, err := bindADKToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ADKTokenFilterer{contract: contract}, nil
}

// bindADKToken binds a generic wrapper to an already deployed contract.
func bindADKToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ADKTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ADKToken *ADKTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ADKToken.Contract.ADKTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ADKToken *ADKTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ADKToken.Contract.ADKTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ADKToken *ADKTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ADKToken.Contract.ADKTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ADKToken *ADKTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ADKToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ADKToken *ADKTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ADKToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ADKToken *ADKTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ADKToken.Contract.contract.Transact(opts, method, params...)
}

// ADDRTOAZ9 is a free data retrieval call binding the contract method 0xb79c88ec.
//
// Solidity: function ADDR_TO_AZ9(address ethAddr) pure returns(string)
func (_ADKToken *ADKTokenCaller) ADDRTOAZ9(opts *bind.CallOpts, ethAddr common.Address) (string, error) {
	var out []interface{}
	err := _ADKToken.contract.Call(opts, &out, "ADDR_TO_AZ9", ethAddr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ADDRTOAZ9 is a free data retrieval call binding the contract method 0xb79c88ec.
//
// Solidity: function ADDR_TO_AZ9(address ethAddr) pure returns(string)
func (_ADKToken *ADKTokenSession) ADDRTOAZ9(ethAddr common.Address) (string, error) {
	return _ADKToken.Contract.ADDRTOAZ9(&_ADKToken.CallOpts, ethAddr)
}

// ADDRTOAZ9 is a free data retrieval call binding the contract method 0xb79c88ec.
//
// Solidity: function ADDR_TO_AZ9(address ethAddr) pure returns(string)
func (_ADKToken *ADKTokenCallerSession) ADDRTOAZ9(ethAddr common.Address) (string, error) {
	return _ADKToken.Contract.ADDRTOAZ9(&_ADKToken.CallOpts, ethAddr)
}

// ADKTransactionsContract is a free data retrieval call binding the contract method 0x8ab04a78.
//
// Solidity: function ADKTransactionsContract() view returns(address)
func (_ADKToken *ADKTokenCaller) ADKTransactionsContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ADKToken.contract.Call(opts, &out, "ADKTransactionsContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADKTransactionsContract is a free data retrieval call binding the contract method 0x8ab04a78.
//
// Solidity: function ADKTransactionsContract() view returns(address)
func (_ADKToken *ADKTokenSession) ADKTransactionsContract() (common.Address, error) {
	return _ADKToken.Contract.ADKTransactionsContract(&_ADKToken.CallOpts)
}

// ADKTransactionsContract is a free data retrieval call binding the contract method 0x8ab04a78.
//
// Solidity: function ADKTransactionsContract() view returns(address)
func (_ADKToken *ADKTokenCallerSession) ADKTransactionsContract() (common.Address, error) {
	return _ADKToken.Contract.ADKTransactionsContract(&_ADKToken.CallOpts)
}

// AGSClaimContract is a free data retrieval call binding the contract method 0x41ef2686.
//
// Solidity: function AGSClaimContract() view returns(address)
func (_ADKToken *ADKTokenCaller) AGSClaimContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ADKToken.contract.Call(opts, &out, "AGSClaimContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AGSClaimContract is a free data retrieval call binding the contract method 0x41ef2686.
//
// Solidity: function AGSClaimContract() view returns(address)
func (_ADKToken *ADKTokenSession) AGSClaimContract() (common.Address, error) {
	return _ADKToken.Contract.AGSClaimContract(&_ADKToken.CallOpts)
}

// AGSClaimContract is a free data retrieval call binding the contract method 0x41ef2686.
//
// Solidity: function AGSClaimContract() view returns(address)
func (_ADKToken *ADKTokenCallerSession) AGSClaimContract() (common.Address, error) {
	return _ADKToken.Contract.AGSClaimContract(&_ADKToken.CallOpts)
}

// AZ9TOADDR is a free data retrieval call binding the contract method 0xa205a8a9.
//
// Solidity: function AZ9_TO_ADDR(string adkString) pure returns(address)
func (_ADKToken *ADKTokenCaller) AZ9TOADDR(opts *bind.CallOpts, adkString string) (common.Address, error) {
	var out []interface{}
	err := _ADKToken.contract.Call(opts, &out, "AZ9_TO_ADDR", adkString)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AZ9TOADDR is a free data retrieval call binding the contract method 0xa205a8a9.
//
// Solidity: function AZ9_TO_ADDR(string adkString) pure returns(address)
func (_ADKToken *ADKTokenSession) AZ9TOADDR(adkString string) (common.Address, error) {
	return _ADKToken.Contract.AZ9TOADDR(&_ADKToken.CallOpts, adkString)
}

// AZ9TOADDR is a free data retrieval call binding the contract method 0xa205a8a9.
//
// Solidity: function AZ9_TO_ADDR(string adkString) pure returns(address)
func (_ADKToken *ADKTokenCallerSession) AZ9TOADDR(adkString string) (common.Address, error) {
	return _ADKToken.Contract.AZ9TOADDR(&_ADKToken.CallOpts, adkString)
}

// AZ9balanceOf is a free data retrieval call binding the contract method 0x0e7fdbb1.
//
// Solidity: function AZ9balanceOf(string _adkAddr) view returns(uint256 balance)
func (_ADKToken *ADKTokenCaller) AZ9balanceOf(opts *bind.CallOpts, _adkAddr string) (*big.Int, error) {
	var out []interface{}
	err := _ADKToken.contract.Call(opts, &out, "AZ9balanceOf", _adkAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AZ9balanceOf is a free data retrieval call binding the contract method 0x0e7fdbb1.
//
// Solidity: function AZ9balanceOf(string _adkAddr) view returns(uint256 balance)
func (_ADKToken *ADKTokenSession) AZ9balanceOf(_adkAddr string) (*big.Int, error) {
	return _ADKToken.Contract.AZ9balanceOf(&_ADKToken.CallOpts, _adkAddr)
}

// AZ9balanceOf is a free data retrieval call binding the contract method 0x0e7fdbb1.
//
// Solidity: function AZ9balanceOf(string _adkAddr) view returns(uint256 balance)
func (_ADKToken *ADKTokenCallerSession) AZ9balanceOf(_adkAddr string) (*big.Int, error) {
	return _ADKToken.Contract.AZ9balanceOf(&_ADKToken.CallOpts, _adkAddr)
}

// BADDRTOAZ9 is a free data retrieval call binding the contract method 0x78f7c041.
//
// Solidity: function BADDR_TO_AZ9(bytes ethAddr) pure returns(string)
func (_ADKToken *ADKTokenCaller) BADDRTOAZ9(opts *bind.CallOpts, ethAddr []byte) (string, error) {
	var out []interface{}
	err := _ADKToken.contract.Call(opts, &out, "BADDR_TO_AZ9", ethAddr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BADDRTOAZ9 is a free data retrieval call binding the contract method 0x78f7c041.
//
// Solidity: function BADDR_TO_AZ9(bytes ethAddr) pure returns(string)
func (_ADKToken *ADKTokenSession) BADDRTOAZ9(ethAddr []byte) (string, error) {
	return _ADKToken.Contract.BADDRTOAZ9(&_ADKToken.CallOpts, ethAddr)
}

// BADDRTOAZ9 is a free data retrieval call binding the contract method 0x78f7c041.
//
// Solidity: function BADDR_TO_AZ9(bytes ethAddr) pure returns(string)
func (_ADKToken *ADKTokenCallerSession) BADDRTOAZ9(ethAddr []byte) (string, error) {
	return _ADKToken.Contract.BADDRTOAZ9(&_ADKToken.CallOpts, ethAddr)
}

// AdkgoGenesisAddress is a free data retrieval call binding the contract method 0xc4c7e151.
//
// Solidity: function adkgo_genesis_address() view returns(address)
func (_ADKToken *ADKTokenCaller) AdkgoGenesisAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ADKToken.contract.Call(opts, &out, "adkgo_genesis_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdkgoGenesisAddress is a free data retrieval call binding the contract method 0xc4c7e151.
//
// Solidity: function adkgo_genesis_address() view returns(address)
func (_ADKToken *ADKTokenSession) AdkgoGenesisAddress() (common.Address, error) {
	return _ADKToken.Contract.AdkgoGenesisAddress(&_ADKToken.CallOpts)
}

// AdkgoGenesisAddress is a free data retrieval call binding the contract method 0xc4c7e151.
//
// Solidity: function adkgo_genesis_address() view returns(address)
func (_ADKToken *ADKTokenCallerSession) AdkgoGenesisAddress() (common.Address, error) {
	return _ADKToken.Contract.AdkgoGenesisAddress(&_ADKToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_ADKToken *ADKTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ADKToken.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_ADKToken *ADKTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ADKToken.Contract.Allowance(&_ADKToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_ADKToken *ADKTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ADKToken.Contract.Allowance(&_ADKToken.CallOpts, _owner, _spender)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_ADKToken *ADKTokenCaller) Allowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ADKToken.contract.Call(opts, &out, "allowed", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_ADKToken *ADKTokenSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ADKToken.Contract.Allowed(&_ADKToken.CallOpts, arg0, arg1)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_ADKToken *ADKTokenCallerSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ADKToken.Contract.Allowed(&_ADKToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_ADKToken *ADKTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ADKToken.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_ADKToken *ADKTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ADKToken.Contract.BalanceOf(&_ADKToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_ADKToken *ADKTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ADKToken.Contract.BalanceOf(&_ADKToken.CallOpts, _owner)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_ADKToken *ADKTokenCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ADKToken.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_ADKToken *ADKTokenSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ADKToken.Contract.Balances(&_ADKToken.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_ADKToken *ADKTokenCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _ADKToken.Contract.Balances(&_ADKToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ADKToken *ADKTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ADKToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ADKToken *ADKTokenSession) Decimals() (uint8, error) {
	return _ADKToken.Contract.Decimals(&_ADKToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ADKToken *ADKTokenCallerSession) Decimals() (uint8, error) {
	return _ADKToken.Contract.Decimals(&_ADKToken.CallOpts)
}

// KnownAddresses is a free data retrieval call binding the contract method 0x3f43347a.
//
// Solidity: function known_addresses(address ) view returns(string)
func (_ADKToken *ADKTokenCaller) KnownAddresses(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _ADKToken.contract.Call(opts, &out, "known_addresses", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// KnownAddresses is a free data retrieval call binding the contract method 0x3f43347a.
//
// Solidity: function known_addresses(address ) view returns(string)
func (_ADKToken *ADKTokenSession) KnownAddresses(arg0 common.Address) (string, error) {
	return _ADKToken.Contract.KnownAddresses(&_ADKToken.CallOpts, arg0)
}

// KnownAddresses is a free data retrieval call binding the contract method 0x3f43347a.
//
// Solidity: function known_addresses(address ) view returns(string)
func (_ADKToken *ADKTokenCallerSession) KnownAddresses(arg0 common.Address) (string, error) {
	return _ADKToken.Contract.KnownAddresses(&_ADKToken.CallOpts, arg0)
}

// LinkedListAllBalances is a free data retrieval call binding the contract method 0x2a5d2cf7.
//
// Solidity: function linked_list_all_balances(address ) view returns(address)
func (_ADKToken *ADKTokenCaller) LinkedListAllBalances(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ADKToken.contract.Call(opts, &out, "linked_list_all_balances", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LinkedListAllBalances is a free data retrieval call binding the contract method 0x2a5d2cf7.
//
// Solidity: function linked_list_all_balances(address ) view returns(address)
func (_ADKToken *ADKTokenSession) LinkedListAllBalances(arg0 common.Address) (common.Address, error) {
	return _ADKToken.Contract.LinkedListAllBalances(&_ADKToken.CallOpts, arg0)
}

// LinkedListAllBalances is a free data retrieval call binding the contract method 0x2a5d2cf7.
//
// Solidity: function linked_list_all_balances(address ) view returns(address)
func (_ADKToken *ADKTokenCallerSession) LinkedListAllBalances(arg0 common.Address) (common.Address, error) {
	return _ADKToken.Contract.LinkedListAllBalances(&_ADKToken.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ADKToken *ADKTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ADKToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ADKToken *ADKTokenSession) Name() (string, error) {
	return _ADKToken.Contract.Name(&_ADKToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ADKToken *ADKTokenCallerSession) Name() (string, error) {
	return _ADKToken.Contract.Name(&_ADKToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ADKToken *ADKTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ADKToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ADKToken *ADKTokenSession) Symbol() (string, error) {
	return _ADKToken.Contract.Symbol(&_ADKToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ADKToken *ADKTokenCallerSession) Symbol() (string, error) {
	return _ADKToken.Contract.Symbol(&_ADKToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ADKToken *ADKTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ADKToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ADKToken *ADKTokenSession) TotalSupply() (*big.Int, error) {
	return _ADKToken.Contract.TotalSupply(&_ADKToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ADKToken *ADKTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ADKToken.Contract.TotalSupply(&_ADKToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_ADKToken *ADKTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ADKToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_ADKToken *ADKTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ADKToken.Contract.Approve(&_ADKToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_ADKToken *ADKTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ADKToken.Contract.Approve(&_ADKToken.TransactOpts, _spender, _value)
}

// GenesisTransaction is a paid mutator transaction binding the contract method 0xa85cfb66.
//
// Solidity: function genesisTransaction(string _meshaddr, int256 _value) returns()
func (_ADKToken *ADKTokenTransactor) GenesisTransaction(opts *bind.TransactOpts, _meshaddr string, _value *big.Int) (*types.Transaction, error) {
	return _ADKToken.contract.Transact(opts, "genesisTransaction", _meshaddr, _value)
}

// GenesisTransaction is a paid mutator transaction binding the contract method 0xa85cfb66.
//
// Solidity: function genesisTransaction(string _meshaddr, int256 _value) returns()
func (_ADKToken *ADKTokenSession) GenesisTransaction(_meshaddr string, _value *big.Int) (*types.Transaction, error) {
	return _ADKToken.Contract.GenesisTransaction(&_ADKToken.TransactOpts, _meshaddr, _value)
}

// GenesisTransaction is a paid mutator transaction binding the contract method 0xa85cfb66.
//
// Solidity: function genesisTransaction(string _meshaddr, int256 _value) returns()
func (_ADKToken *ADKTokenTransactorSession) GenesisTransaction(_meshaddr string, _value *big.Int) (*types.Transaction, error) {
	return _ADKToken.Contract.GenesisTransaction(&_ADKToken.TransactOpts, _meshaddr, _value)
}

// MeshTransaction is a paid mutator transaction binding the contract method 0x8756ee19.
//
// Solidity: function meshTransaction(string _meshaddr, int256 _value) returns()
func (_ADKToken *ADKTokenTransactor) MeshTransaction(opts *bind.TransactOpts, _meshaddr string, _value *big.Int) (*types.Transaction, error) {
	return _ADKToken.contract.Transact(opts, "meshTransaction", _meshaddr, _value)
}

// MeshTransaction is a paid mutator transaction binding the contract method 0x8756ee19.
//
// Solidity: function meshTransaction(string _meshaddr, int256 _value) returns()
func (_ADKToken *ADKTokenSession) MeshTransaction(_meshaddr string, _value *big.Int) (*types.Transaction, error) {
	return _ADKToken.Contract.MeshTransaction(&_ADKToken.TransactOpts, _meshaddr, _value)
}

// MeshTransaction is a paid mutator transaction binding the contract method 0x8756ee19.
//
// Solidity: function meshTransaction(string _meshaddr, int256 _value) returns()
func (_ADKToken *ADKTokenTransactorSession) MeshTransaction(_meshaddr string, _value *big.Int) (*types.Transaction, error) {
	return _ADKToken.Contract.MeshTransaction(&_ADKToken.TransactOpts, _meshaddr, _value)
}

// RegisterAddress is a paid mutator transaction binding the contract method 0xbb1bd28f.
//
// Solidity: function registerAddress(string _meshaddr) returns()
func (_ADKToken *ADKTokenTransactor) RegisterAddress(opts *bind.TransactOpts, _meshaddr string) (*types.Transaction, error) {
	return _ADKToken.contract.Transact(opts, "registerAddress", _meshaddr)
}

// RegisterAddress is a paid mutator transaction binding the contract method 0xbb1bd28f.
//
// Solidity: function registerAddress(string _meshaddr) returns()
func (_ADKToken *ADKTokenSession) RegisterAddress(_meshaddr string) (*types.Transaction, error) {
	return _ADKToken.Contract.RegisterAddress(&_ADKToken.TransactOpts, _meshaddr)
}

// RegisterAddress is a paid mutator transaction binding the contract method 0xbb1bd28f.
//
// Solidity: function registerAddress(string _meshaddr) returns()
func (_ADKToken *ADKTokenTransactorSession) RegisterAddress(_meshaddr string) (*types.Transaction, error) {
	return _ADKToken.Contract.RegisterAddress(&_ADKToken.TransactOpts, _meshaddr)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_ADKToken *ADKTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ADKToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_ADKToken *ADKTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ADKToken.Contract.Transfer(&_ADKToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_ADKToken *ADKTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ADKToken.Contract.Transfer(&_ADKToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_ADKToken *ADKTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ADKToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_ADKToken *ADKTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ADKToken.Contract.TransferFrom(&_ADKToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_ADKToken *ADKTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ADKToken.Contract.TransferFrom(&_ADKToken.TransactOpts, _from, _to, _value)
}

// ADKTokenAZ9TransferIterator is returned from FilterAZ9Transfer and is used to iterate over the raw logs and unpacked data for AZ9Transfer events raised by the ADKToken contract.
type ADKTokenAZ9TransferIterator struct {
	Event *ADKTokenAZ9Transfer // Event containing the contract specifics and raw log

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
func (it *ADKTokenAZ9TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ADKTokenAZ9Transfer)
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
		it.Event = new(ADKTokenAZ9Transfer)
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
func (it *ADKTokenAZ9TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ADKTokenAZ9TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ADKTokenAZ9Transfer represents a AZ9Transfer event raised by the ADKToken contract.
type ADKTokenAZ9Transfer struct {
	Meshaddr common.Hash
	Addr     common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAZ9Transfer is a free log retrieval operation binding the contract event 0xa93369924490bc7698b9d289925c432d4f16b9465c1590324916dba7ed166246.
//
// Solidity: event AZ9Transfer(string indexed _meshaddr, address indexed _addr, int256 _value)
func (_ADKToken *ADKTokenFilterer) FilterAZ9Transfer(opts *bind.FilterOpts, _meshaddr []string, _addr []common.Address) (*ADKTokenAZ9TransferIterator, error) {

	var _meshaddrRule []interface{}
	for _, _meshaddrItem := range _meshaddr {
		_meshaddrRule = append(_meshaddrRule, _meshaddrItem)
	}
	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _ADKToken.contract.FilterLogs(opts, "AZ9Transfer", _meshaddrRule, _addrRule)
	if err != nil {
		return nil, err
	}
	return &ADKTokenAZ9TransferIterator{contract: _ADKToken.contract, event: "AZ9Transfer", logs: logs, sub: sub}, nil
}

// WatchAZ9Transfer is a free log subscription operation binding the contract event 0xa93369924490bc7698b9d289925c432d4f16b9465c1590324916dba7ed166246.
//
// Solidity: event AZ9Transfer(string indexed _meshaddr, address indexed _addr, int256 _value)
func (_ADKToken *ADKTokenFilterer) WatchAZ9Transfer(opts *bind.WatchOpts, sink chan<- *ADKTokenAZ9Transfer, _meshaddr []string, _addr []common.Address) (event.Subscription, error) {

	var _meshaddrRule []interface{}
	for _, _meshaddrItem := range _meshaddr {
		_meshaddrRule = append(_meshaddrRule, _meshaddrItem)
	}
	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _ADKToken.contract.WatchLogs(opts, "AZ9Transfer", _meshaddrRule, _addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ADKTokenAZ9Transfer)
				if err := _ADKToken.contract.UnpackLog(event, "AZ9Transfer", log); err != nil {
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

// ParseAZ9Transfer is a log parse operation binding the contract event 0xa93369924490bc7698b9d289925c432d4f16b9465c1590324916dba7ed166246.
//
// Solidity: event AZ9Transfer(string indexed _meshaddr, address indexed _addr, int256 _value)
func (_ADKToken *ADKTokenFilterer) ParseAZ9Transfer(log types.Log) (*ADKTokenAZ9Transfer, error) {
	event := new(ADKTokenAZ9Transfer)
	if err := _ADKToken.contract.UnpackLog(event, "AZ9Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ADKTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ADKToken contract.
type ADKTokenApprovalIterator struct {
	Event *ADKTokenApproval // Event containing the contract specifics and raw log

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
func (it *ADKTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ADKTokenApproval)
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
		it.Event = new(ADKTokenApproval)
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
func (it *ADKTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ADKTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ADKTokenApproval represents a Approval event raised by the ADKToken contract.
type ADKTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_ADKToken *ADKTokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*ADKTokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ADKToken.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &ADKTokenApprovalIterator{contract: _ADKToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_ADKToken *ADKTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ADKTokenApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ADKToken.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ADKTokenApproval)
				if err := _ADKToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_ADKToken *ADKTokenFilterer) ParseApproval(log types.Log) (*ADKTokenApproval, error) {
	event := new(ADKTokenApproval)
	if err := _ADKToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ADKTokenGENESISTransactionIterator is returned from FilterGENESISTransaction and is used to iterate over the raw logs and unpacked data for GENESISTransaction events raised by the ADKToken contract.
type ADKTokenGENESISTransactionIterator struct {
	Event *ADKTokenGENESISTransaction // Event containing the contract specifics and raw log

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
func (it *ADKTokenGENESISTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ADKTokenGENESISTransaction)
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
		it.Event = new(ADKTokenGENESISTransaction)
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
func (it *ADKTokenGENESISTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ADKTokenGENESISTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ADKTokenGENESISTransaction represents a GENESISTransaction event raised by the ADKToken contract.
type ADKTokenGENESISTransaction struct {
	Meshaddr common.Hash
	Addr     common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGENESISTransaction is a free log retrieval operation binding the contract event 0x32ac6262ad9a118aca2d9d65c3cfb1c88b8d120f8a750a635168ba94fcf6e164.
//
// Solidity: event GENESISTransaction(string indexed _meshaddr, address indexed _addr, int256 _value)
func (_ADKToken *ADKTokenFilterer) FilterGENESISTransaction(opts *bind.FilterOpts, _meshaddr []string, _addr []common.Address) (*ADKTokenGENESISTransactionIterator, error) {

	var _meshaddrRule []interface{}
	for _, _meshaddrItem := range _meshaddr {
		_meshaddrRule = append(_meshaddrRule, _meshaddrItem)
	}
	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _ADKToken.contract.FilterLogs(opts, "GENESISTransaction", _meshaddrRule, _addrRule)
	if err != nil {
		return nil, err
	}
	return &ADKTokenGENESISTransactionIterator{contract: _ADKToken.contract, event: "GENESISTransaction", logs: logs, sub: sub}, nil
}

// WatchGENESISTransaction is a free log subscription operation binding the contract event 0x32ac6262ad9a118aca2d9d65c3cfb1c88b8d120f8a750a635168ba94fcf6e164.
//
// Solidity: event GENESISTransaction(string indexed _meshaddr, address indexed _addr, int256 _value)
func (_ADKToken *ADKTokenFilterer) WatchGENESISTransaction(opts *bind.WatchOpts, sink chan<- *ADKTokenGENESISTransaction, _meshaddr []string, _addr []common.Address) (event.Subscription, error) {

	var _meshaddrRule []interface{}
	for _, _meshaddrItem := range _meshaddr {
		_meshaddrRule = append(_meshaddrRule, _meshaddrItem)
	}
	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}

	logs, sub, err := _ADKToken.contract.WatchLogs(opts, "GENESISTransaction", _meshaddrRule, _addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ADKTokenGENESISTransaction)
				if err := _ADKToken.contract.UnpackLog(event, "GENESISTransaction", log); err != nil {
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

// ParseGENESISTransaction is a log parse operation binding the contract event 0x32ac6262ad9a118aca2d9d65c3cfb1c88b8d120f8a750a635168ba94fcf6e164.
//
// Solidity: event GENESISTransaction(string indexed _meshaddr, address indexed _addr, int256 _value)
func (_ADKToken *ADKTokenFilterer) ParseGENESISTransaction(log types.Log) (*ADKTokenGENESISTransaction, error) {
	event := new(ADKTokenGENESISTransaction)
	if err := _ADKToken.contract.UnpackLog(event, "GENESISTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ADKTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ADKToken contract.
type ADKTokenTransferIterator struct {
	Event *ADKTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ADKTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ADKTokenTransfer)
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
		it.Event = new(ADKTokenTransfer)
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
func (it *ADKTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ADKTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ADKTokenTransfer represents a Transfer event raised by the ADKToken contract.
type ADKTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_ADKToken *ADKTokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ADKTokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ADKToken.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ADKTokenTransferIterator{contract: _ADKToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_ADKToken *ADKTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ADKTokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ADKToken.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ADKTokenTransfer)
				if err := _ADKToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_ADKToken *ADKTokenFilterer) ParseTransfer(log types.Log) (*ADKTokenTransfer, error) {
	event := new(ADKTokenTransfer)
	if err := _ADKToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ADKTokenInterfaceMetaData contains all meta data concerning the ADKTokenInterface contract.
var ADKTokenInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"adkString\",\"type\":\"string\"}],\"name\":\"AZ9_TO_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_adkAddr\",\"type\":\"string\"}],\"name\":\"AZ9balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_meshaddr\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"_value\",\"type\":\"int256\"}],\"name\":\"meshTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a205a8a9": "AZ9_TO_ADDR(string)",
		"0e7fdbb1": "AZ9balanceOf(string)",
		"8756ee19": "meshTransaction(string,int256)",
	},
}

// ADKTokenInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use ADKTokenInterfaceMetaData.ABI instead.
var ADKTokenInterfaceABI = ADKTokenInterfaceMetaData.ABI

// Deprecated: Use ADKTokenInterfaceMetaData.Sigs instead.
// ADKTokenInterfaceFuncSigs maps the 4-byte function signature to its string representation.
var ADKTokenInterfaceFuncSigs = ADKTokenInterfaceMetaData.Sigs

// ADKTokenInterface is an auto generated Go binding around an Ethereum contract.
type ADKTokenInterface struct {
	ADKTokenInterfaceCaller     // Read-only binding to the contract
	ADKTokenInterfaceTransactor // Write-only binding to the contract
	ADKTokenInterfaceFilterer   // Log filterer for contract events
}

// ADKTokenInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ADKTokenInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ADKTokenInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ADKTokenInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ADKTokenInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ADKTokenInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ADKTokenInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ADKTokenInterfaceSession struct {
	Contract     *ADKTokenInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ADKTokenInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ADKTokenInterfaceCallerSession struct {
	Contract *ADKTokenInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ADKTokenInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ADKTokenInterfaceTransactorSession struct {
	Contract     *ADKTokenInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ADKTokenInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ADKTokenInterfaceRaw struct {
	Contract *ADKTokenInterface // Generic contract binding to access the raw methods on
}

// ADKTokenInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ADKTokenInterfaceCallerRaw struct {
	Contract *ADKTokenInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// ADKTokenInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ADKTokenInterfaceTransactorRaw struct {
	Contract *ADKTokenInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewADKTokenInterface creates a new instance of ADKTokenInterface, bound to a specific deployed contract.
func NewADKTokenInterface(address common.Address, backend bind.ContractBackend) (*ADKTokenInterface, error) {
	contract, err := bindADKTokenInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ADKTokenInterface{ADKTokenInterfaceCaller: ADKTokenInterfaceCaller{contract: contract}, ADKTokenInterfaceTransactor: ADKTokenInterfaceTransactor{contract: contract}, ADKTokenInterfaceFilterer: ADKTokenInterfaceFilterer{contract: contract}}, nil
}

// NewADKTokenInterfaceCaller creates a new read-only instance of ADKTokenInterface, bound to a specific deployed contract.
func NewADKTokenInterfaceCaller(address common.Address, caller bind.ContractCaller) (*ADKTokenInterfaceCaller, error) {
	contract, err := bindADKTokenInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ADKTokenInterfaceCaller{contract: contract}, nil
}

// NewADKTokenInterfaceTransactor creates a new write-only instance of ADKTokenInterface, bound to a specific deployed contract.
func NewADKTokenInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ADKTokenInterfaceTransactor, error) {
	contract, err := bindADKTokenInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ADKTokenInterfaceTransactor{contract: contract}, nil
}

// NewADKTokenInterfaceFilterer creates a new log filterer instance of ADKTokenInterface, bound to a specific deployed contract.
func NewADKTokenInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ADKTokenInterfaceFilterer, error) {
	contract, err := bindADKTokenInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ADKTokenInterfaceFilterer{contract: contract}, nil
}

// bindADKTokenInterface binds a generic wrapper to an already deployed contract.
func bindADKTokenInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ADKTokenInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ADKTokenInterface *ADKTokenInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ADKTokenInterface.Contract.ADKTokenInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ADKTokenInterface *ADKTokenInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ADKTokenInterface.Contract.ADKTokenInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ADKTokenInterface *ADKTokenInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ADKTokenInterface.Contract.ADKTokenInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ADKTokenInterface *ADKTokenInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ADKTokenInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ADKTokenInterface *ADKTokenInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ADKTokenInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ADKTokenInterface *ADKTokenInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ADKTokenInterface.Contract.contract.Transact(opts, method, params...)
}

// AZ9TOADDR is a free data retrieval call binding the contract method 0xa205a8a9.
//
// Solidity: function AZ9_TO_ADDR(string adkString) pure returns(address)
func (_ADKTokenInterface *ADKTokenInterfaceCaller) AZ9TOADDR(opts *bind.CallOpts, adkString string) (common.Address, error) {
	var out []interface{}
	err := _ADKTokenInterface.contract.Call(opts, &out, "AZ9_TO_ADDR", adkString)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AZ9TOADDR is a free data retrieval call binding the contract method 0xa205a8a9.
//
// Solidity: function AZ9_TO_ADDR(string adkString) pure returns(address)
func (_ADKTokenInterface *ADKTokenInterfaceSession) AZ9TOADDR(adkString string) (common.Address, error) {
	return _ADKTokenInterface.Contract.AZ9TOADDR(&_ADKTokenInterface.CallOpts, adkString)
}

// AZ9TOADDR is a free data retrieval call binding the contract method 0xa205a8a9.
//
// Solidity: function AZ9_TO_ADDR(string adkString) pure returns(address)
func (_ADKTokenInterface *ADKTokenInterfaceCallerSession) AZ9TOADDR(adkString string) (common.Address, error) {
	return _ADKTokenInterface.Contract.AZ9TOADDR(&_ADKTokenInterface.CallOpts, adkString)
}

// AZ9balanceOf is a free data retrieval call binding the contract method 0x0e7fdbb1.
//
// Solidity: function AZ9balanceOf(string _adkAddr) view returns(uint256 balance)
func (_ADKTokenInterface *ADKTokenInterfaceCaller) AZ9balanceOf(opts *bind.CallOpts, _adkAddr string) (*big.Int, error) {
	var out []interface{}
	err := _ADKTokenInterface.contract.Call(opts, &out, "AZ9balanceOf", _adkAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AZ9balanceOf is a free data retrieval call binding the contract method 0x0e7fdbb1.
//
// Solidity: function AZ9balanceOf(string _adkAddr) view returns(uint256 balance)
func (_ADKTokenInterface *ADKTokenInterfaceSession) AZ9balanceOf(_adkAddr string) (*big.Int, error) {
	return _ADKTokenInterface.Contract.AZ9balanceOf(&_ADKTokenInterface.CallOpts, _adkAddr)
}

// AZ9balanceOf is a free data retrieval call binding the contract method 0x0e7fdbb1.
//
// Solidity: function AZ9balanceOf(string _adkAddr) view returns(uint256 balance)
func (_ADKTokenInterface *ADKTokenInterfaceCallerSession) AZ9balanceOf(_adkAddr string) (*big.Int, error) {
	return _ADKTokenInterface.Contract.AZ9balanceOf(&_ADKTokenInterface.CallOpts, _adkAddr)
}

// MeshTransaction is a paid mutator transaction binding the contract method 0x8756ee19.
//
// Solidity: function meshTransaction(string _meshaddr, int256 _value) returns()
func (_ADKTokenInterface *ADKTokenInterfaceTransactor) MeshTransaction(opts *bind.TransactOpts, _meshaddr string, _value *big.Int) (*types.Transaction, error) {
	return _ADKTokenInterface.contract.Transact(opts, "meshTransaction", _meshaddr, _value)
}

// MeshTransaction is a paid mutator transaction binding the contract method 0x8756ee19.
//
// Solidity: function meshTransaction(string _meshaddr, int256 _value) returns()
func (_ADKTokenInterface *ADKTokenInterfaceSession) MeshTransaction(_meshaddr string, _value *big.Int) (*types.Transaction, error) {
	return _ADKTokenInterface.Contract.MeshTransaction(&_ADKTokenInterface.TransactOpts, _meshaddr, _value)
}

// MeshTransaction is a paid mutator transaction binding the contract method 0x8756ee19.
//
// Solidity: function meshTransaction(string _meshaddr, int256 _value) returns()
func (_ADKTokenInterface *ADKTokenInterfaceTransactorSession) MeshTransaction(_meshaddr string, _value *big.Int) (*types.Transaction, error) {
	return _ADKTokenInterface.Contract.MeshTransaction(&_ADKTokenInterface.TransactOpts, _meshaddr, _value)
}

// ADKTokenInterface2MetaData contains all meta data concerning the ADKTokenInterface2 contract.
var ADKTokenInterface2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"adkString\",\"type\":\"string\"}],\"name\":\"AZ9_TO_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_adkAddr\",\"type\":\"string\"}],\"name\":\"AZ9balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_meshaddr\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"_value\",\"type\":\"int256\"}],\"name\":\"meshTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a205a8a9": "AZ9_TO_ADDR(string)",
		"0e7fdbb1": "AZ9balanceOf(string)",
		"8756ee19": "meshTransaction(string,int256)",
	},
}

// ADKTokenInterface2ABI is the input ABI used to generate the binding from.
// Deprecated: Use ADKTokenInterface2MetaData.ABI instead.
var ADKTokenInterface2ABI = ADKTokenInterface2MetaData.ABI

// Deprecated: Use ADKTokenInterface2MetaData.Sigs instead.
// ADKTokenInterface2FuncSigs maps the 4-byte function signature to its string representation.
var ADKTokenInterface2FuncSigs = ADKTokenInterface2MetaData.Sigs

// ADKTokenInterface2 is an auto generated Go binding around an Ethereum contract.
type ADKTokenInterface2 struct {
	ADKTokenInterface2Caller     // Read-only binding to the contract
	ADKTokenInterface2Transactor // Write-only binding to the contract
	ADKTokenInterface2Filterer   // Log filterer for contract events
}

// ADKTokenInterface2Caller is an auto generated read-only Go binding around an Ethereum contract.
type ADKTokenInterface2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ADKTokenInterface2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ADKTokenInterface2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ADKTokenInterface2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ADKTokenInterface2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ADKTokenInterface2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ADKTokenInterface2Session struct {
	Contract     *ADKTokenInterface2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ADKTokenInterface2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ADKTokenInterface2CallerSession struct {
	Contract *ADKTokenInterface2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ADKTokenInterface2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ADKTokenInterface2TransactorSession struct {
	Contract     *ADKTokenInterface2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ADKTokenInterface2Raw is an auto generated low-level Go binding around an Ethereum contract.
type ADKTokenInterface2Raw struct {
	Contract *ADKTokenInterface2 // Generic contract binding to access the raw methods on
}

// ADKTokenInterface2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ADKTokenInterface2CallerRaw struct {
	Contract *ADKTokenInterface2Caller // Generic read-only contract binding to access the raw methods on
}

// ADKTokenInterface2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ADKTokenInterface2TransactorRaw struct {
	Contract *ADKTokenInterface2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewADKTokenInterface2 creates a new instance of ADKTokenInterface2, bound to a specific deployed contract.
func NewADKTokenInterface2(address common.Address, backend bind.ContractBackend) (*ADKTokenInterface2, error) {
	contract, err := bindADKTokenInterface2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ADKTokenInterface2{ADKTokenInterface2Caller: ADKTokenInterface2Caller{contract: contract}, ADKTokenInterface2Transactor: ADKTokenInterface2Transactor{contract: contract}, ADKTokenInterface2Filterer: ADKTokenInterface2Filterer{contract: contract}}, nil
}

// NewADKTokenInterface2Caller creates a new read-only instance of ADKTokenInterface2, bound to a specific deployed contract.
func NewADKTokenInterface2Caller(address common.Address, caller bind.ContractCaller) (*ADKTokenInterface2Caller, error) {
	contract, err := bindADKTokenInterface2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ADKTokenInterface2Caller{contract: contract}, nil
}

// NewADKTokenInterface2Transactor creates a new write-only instance of ADKTokenInterface2, bound to a specific deployed contract.
func NewADKTokenInterface2Transactor(address common.Address, transactor bind.ContractTransactor) (*ADKTokenInterface2Transactor, error) {
	contract, err := bindADKTokenInterface2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ADKTokenInterface2Transactor{contract: contract}, nil
}

// NewADKTokenInterface2Filterer creates a new log filterer instance of ADKTokenInterface2, bound to a specific deployed contract.
func NewADKTokenInterface2Filterer(address common.Address, filterer bind.ContractFilterer) (*ADKTokenInterface2Filterer, error) {
	contract, err := bindADKTokenInterface2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ADKTokenInterface2Filterer{contract: contract}, nil
}

// bindADKTokenInterface2 binds a generic wrapper to an already deployed contract.
func bindADKTokenInterface2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ADKTokenInterface2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ADKTokenInterface2 *ADKTokenInterface2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ADKTokenInterface2.Contract.ADKTokenInterface2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ADKTokenInterface2 *ADKTokenInterface2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ADKTokenInterface2.Contract.ADKTokenInterface2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ADKTokenInterface2 *ADKTokenInterface2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ADKTokenInterface2.Contract.ADKTokenInterface2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ADKTokenInterface2 *ADKTokenInterface2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ADKTokenInterface2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ADKTokenInterface2 *ADKTokenInterface2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ADKTokenInterface2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ADKTokenInterface2 *ADKTokenInterface2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ADKTokenInterface2.Contract.contract.Transact(opts, method, params...)
}

// AZ9TOADDR is a free data retrieval call binding the contract method 0xa205a8a9.
//
// Solidity: function AZ9_TO_ADDR(string adkString) pure returns(address)
func (_ADKTokenInterface2 *ADKTokenInterface2Caller) AZ9TOADDR(opts *bind.CallOpts, adkString string) (common.Address, error) {
	var out []interface{}
	err := _ADKTokenInterface2.contract.Call(opts, &out, "AZ9_TO_ADDR", adkString)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AZ9TOADDR is a free data retrieval call binding the contract method 0xa205a8a9.
//
// Solidity: function AZ9_TO_ADDR(string adkString) pure returns(address)
func (_ADKTokenInterface2 *ADKTokenInterface2Session) AZ9TOADDR(adkString string) (common.Address, error) {
	return _ADKTokenInterface2.Contract.AZ9TOADDR(&_ADKTokenInterface2.CallOpts, adkString)
}

// AZ9TOADDR is a free data retrieval call binding the contract method 0xa205a8a9.
//
// Solidity: function AZ9_TO_ADDR(string adkString) pure returns(address)
func (_ADKTokenInterface2 *ADKTokenInterface2CallerSession) AZ9TOADDR(adkString string) (common.Address, error) {
	return _ADKTokenInterface2.Contract.AZ9TOADDR(&_ADKTokenInterface2.CallOpts, adkString)
}

// AZ9balanceOf is a free data retrieval call binding the contract method 0x0e7fdbb1.
//
// Solidity: function AZ9balanceOf(string _adkAddr) view returns(uint256 balance)
func (_ADKTokenInterface2 *ADKTokenInterface2Caller) AZ9balanceOf(opts *bind.CallOpts, _adkAddr string) (*big.Int, error) {
	var out []interface{}
	err := _ADKTokenInterface2.contract.Call(opts, &out, "AZ9balanceOf", _adkAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AZ9balanceOf is a free data retrieval call binding the contract method 0x0e7fdbb1.
//
// Solidity: function AZ9balanceOf(string _adkAddr) view returns(uint256 balance)
func (_ADKTokenInterface2 *ADKTokenInterface2Session) AZ9balanceOf(_adkAddr string) (*big.Int, error) {
	return _ADKTokenInterface2.Contract.AZ9balanceOf(&_ADKTokenInterface2.CallOpts, _adkAddr)
}

// AZ9balanceOf is a free data retrieval call binding the contract method 0x0e7fdbb1.
//
// Solidity: function AZ9balanceOf(string _adkAddr) view returns(uint256 balance)
func (_ADKTokenInterface2 *ADKTokenInterface2CallerSession) AZ9balanceOf(_adkAddr string) (*big.Int, error) {
	return _ADKTokenInterface2.Contract.AZ9balanceOf(&_ADKTokenInterface2.CallOpts, _adkAddr)
}

// MeshTransaction is a paid mutator transaction binding the contract method 0x8756ee19.
//
// Solidity: function meshTransaction(string _meshaddr, int256 _value) returns()
func (_ADKTokenInterface2 *ADKTokenInterface2Transactor) MeshTransaction(opts *bind.TransactOpts, _meshaddr string, _value *big.Int) (*types.Transaction, error) {
	return _ADKTokenInterface2.contract.Transact(opts, "meshTransaction", _meshaddr, _value)
}

// MeshTransaction is a paid mutator transaction binding the contract method 0x8756ee19.
//
// Solidity: function meshTransaction(string _meshaddr, int256 _value) returns()
func (_ADKTokenInterface2 *ADKTokenInterface2Session) MeshTransaction(_meshaddr string, _value *big.Int) (*types.Transaction, error) {
	return _ADKTokenInterface2.Contract.MeshTransaction(&_ADKTokenInterface2.TransactOpts, _meshaddr, _value)
}

// MeshTransaction is a paid mutator transaction binding the contract method 0x8756ee19.
//
// Solidity: function meshTransaction(string _meshaddr, int256 _value) returns()
func (_ADKTokenInterface2 *ADKTokenInterface2TransactorSession) MeshTransaction(_meshaddr string, _value *big.Int) (*types.Transaction, error) {
	return _ADKTokenInterface2.Contract.MeshTransaction(&_ADKTokenInterface2.TransactOpts, _meshaddr, _value)
}

// ADKTransactionsMetaData contains all meta data concerning the ADKTransactions contract.
var ADKTransactionsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_genesis_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ADKTokenContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"intdata\",\"type\":\"int256\"}],\"name\":\"EventLogInt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"strdata\",\"type\":\"string\"}],\"name\":\"EventLogString\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"uintdata\",\"type\":\"uint256\"}],\"name\":\"EventLogUInt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"addressSHA3\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionSHA3\",\"type\":\"bytes32\"}],\"name\":\"transactions_by_address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bundleSHA3\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionSHA3\",\"type\":\"bytes32\"}],\"name\":\"transactions_by_bundle\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADKTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_address\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"_value\",\"type\":\"int256\"}],\"name\":\"ADM_LoadADKBalances\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_addresses\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"_value1\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_value2\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_value3\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_value4\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_value5\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_value6\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_value7\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_value8\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_value9\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_value10\",\"type\":\"int256\"}],\"name\":\"ADM_LoadADKBalancesBulk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_transaction\",\"type\":\"string\"}],\"name\":\"ADM_LoadTransactionsUnchecked\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newContract\",\"type\":\"address\"}],\"name\":\"ADM_SetADKTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ClientProofOfWorkRequirement\",\"type\":\"uint256\"}],\"name\":\"ADM_SetDifficulty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_genesisAddress\",\"type\":\"address\"}],\"name\":\"ADM_SetGenesisAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_newTip\",\"type\":\"bytes32\"}],\"name\":\"ADM_SetTip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ClientProofOfWorkRequirement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"CurlHashOP\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_adkAddr\",\"type\":\"string\"}],\"name\":\"GetAZ9balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"addressString\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numTx\",\"type\":\"uint256\"}],\"name\":\"GetTxByAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bundleString\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"numTx\",\"type\":\"uint256\"}],\"name\":\"GetTxByBundle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"transactiondata\",\"type\":\"string\"}],\"name\":\"PostTransactions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"TryteToIntValue\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"test\",\"type\":\"uint16\"}],\"name\":\"TryteToTrits\",\"outputs\":[{\"internalType\":\"int16[3]\",\"name\":\"\",\"type\":\"int16[3]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adkgo_genesis_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"meshTip\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"spent_addresses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transaction_branch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transaction_hashes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transaction_index\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transaction_indexed_by_seq\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transaction_trunk\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transactionhash_by_address\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transactionhash_by_address_count\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transactionhash_by_bundle\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transactionhash_by_bundle_count\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tx_count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a91fa7cb": "ADKTokenAddress()",
		"64180204": "ADM_LoadADKBalances(string,int256)",
		"6a82351e": "ADM_LoadADKBalancesBulk(string,int256,int256,int256,int256,int256,int256,int256,int256,int256,int256)",
		"50e6b384": "ADM_LoadTransactionsUnchecked(string)",
		"ab815b02": "ADM_SetADKTokenContract(address)",
		"d45c0ebc": "ADM_SetDifficulty(uint256)",
		"d48a1ce5": "ADM_SetGenesisAddress(address)",
		"018e3998": "ADM_SetTip(bytes32)",
		"2076e796": "ClientProofOfWorkRequirement()",
		"016d1c68": "CurlHashOP(string)",
		"55b6193c": "GetAZ9balanceOf(string)",
		"d9fc5448": "GetTxByAddress(string,uint256)",
		"49607e7d": "GetTxByBundle(string,uint256)",
		"7fa5e242": "PostTransactions(string)",
		"510bcf5f": "TryteToIntValue(bytes)",
		"f3067872": "TryteToTrits(uint16)",
		"c4c7e151": "adkgo_genesis_address()",
		"11cf02ec": "meshTip()",
		"de5ed777": "spent_addresses(bytes32)",
		"4893b64e": "transaction_branch(bytes32)",
		"f7a8ffd7": "transaction_hashes(bytes32)",
		"428a4e9c": "transaction_index(bytes32)",
		"c1cba762": "transaction_indexed_by_seq(uint256)",
		"ad69e6a0": "transaction_trunk(bytes32)",
		"b1406a64": "transactionhash_by_address(bytes32)",
		"dcb711ff": "transactionhash_by_address_count(bytes32)",
		"91da69af": "transactionhash_by_bundle(bytes32)",
		"fc74fa69": "transactionhash_by_bundle_count(bytes32)",
		"642f2eaf": "transactions(bytes32)",
		"aa21064f": "tx_count()",
	},
	Bin: "0x60806040526001805460ff60a01b191690556014805461ffff191690553480156200002957600080fd5b5060405162003d5938038062003d598339810160408190526200004c9162000367565b600080546001600160a01b038084166001600160a01b0319928316179092556001805492851692909116919091179055600f6011556200008b62000093565b5050620004c0565b600154600160a01b900460ff1615620000ab57600080fd5b6001805460ff60a01b1916600160a01b1790556040805160808101909152605180825262003d0860208301398051620000ed91600d9160209091019062000220565b5060408051610a71808252610aa0820190925260009160208201818036833701905050905060005b610a718110156200016657603960f81b8282815181106200013a576200013a620004aa565b60200101906001600160f81b031916908160001a905350806200015d8162000480565b91505062000115565b5080516200017c90600c90602084019062000220565b506000600d6040516200019091906200039f565b60405180910390209050600c60026000838152602001908152602001600020908054620001bd9062000443565b620001ca929190620002af565b506000818152600360205260409020600d8054620001e89062000443565b620001f5929190620002af565b5060108190556000818152600e60209081526040808320849055600f909152812091909155600b5550565b8280546200022e9062000443565b90600052602060002090601f0160209004810192826200025257600085556200029d565b82601f106200026d57805160ff19168380011785556200029d565b828001600101855582156200029d579182015b828111156200029d57825182559160200191906001019062000280565b50620002ab92915062000333565b5090565b828054620002bd9062000443565b90600052602060002090601f016020900481019282620002e157600085556200029d565b82601f10620002f457805485556200029d565b828001600101855582156200029d57600052602060002091601f016020900482015b828111156200029d57825482559160010191906001019062000316565b5b80821115620002ab576000815560010162000334565b80516001600160a01b03811681146200036257600080fd5b919050565b600080604083850312156200037b57600080fd5b62000386836200034a565b915062000396602084016200034a565b90509250929050565b600080835481600182811c915080831680620003bc57607f831692505b6020808410821415620003dd57634e487b7160e01b86526022600452602486fd5b818015620003f45760018114620004065762000435565b60ff1986168952848901965062000435565b60008a81526020902060005b868110156200042d5781548b82015290850190830162000412565b505084890196505b509498975050505050505050565b600181811c908216806200045857607f821691505b602082108114156200047a57634e487b7160e01b600052602260045260246000fd5b50919050565b6000600019821415620004a357634e487b7160e01b600052601160045260246000fd5b5060010190565b634e487b7160e01b600052603260045260246000fd5b61383880620004d06000396000f3fe608060405234801561001057600080fd5b50600436106101da5760003560e01c8063a91fa7cb11610104578063d45c0ebc116100a2578063de5ed77711610071578063de5ed77714610463578063f306787214610496578063f7a8ffd7146104b6578063fc74fa69146104c957600080fd5b8063d45c0ebc146103ef578063d48a1ce514610402578063d9fc544814610415578063dcb711ff1461042857600080fd5b8063ad69e6a0116100de578063ad69e6a01461037c578063b1406a641461039c578063c1cba762146103bc578063c4c7e151146103dc57600080fd5b8063a91fa7cb14610335578063aa21064f14610360578063ab815b021461036957600080fd5b806350e6b3841161017c578063642f2eaf1161014b578063642f2eaf146102dc5780636a82351e146102ef5780637fa5e2421461030257806391da69af1461031557600080fd5b806350e6b38414610290578063510bcf5f146102a357806355b6193c146102b657806364180204146102c957600080fd5b80632076e796116101b85780632076e79614610234578063428a4e9c1461023d5780634893b64e1461025d57806349607e7d1461027d57600080fd5b8063016d1c68146101df578063018e39981461020857806311cf02ec1461021d575b600080fd5b6101f26101ed366004613122565b6104ef565b6040516101ff91906133d8565b60405180910390f35b61021b6102163660046130b8565b6106d2565b005b61022660105481565b6040519081526020016101ff565b61022660115481565b61022661024b3660046130b8565b600a6020526000908152604090205481565b61022661026b3660046130b8565b600f6020526000908152604090205481565b6101f261028b366004613157565b61070a565b61021b61029e366004613122565b6107f0565b6102266102b13660046130d1565b610b22565b6102266102c4366004613122565b610bd8565b61021b6102d7366004613157565b610c59565b6101f26102ea3660046130b8565b610ceb565b61021b6102fd36600461319c565b610d85565b6101f2610310366004613122565b611344565b6102266103233660046130b8565b60066020526000908152604090205481565b600054610348906001600160a01b031681565b6040516001600160a01b0390911681526020016101ff565b610226600b5481565b61021b61037736600461307e565b612574565b61022661038a3660046130b8565b600e6020526000908152604090205481565b6102266103aa3660046130b8565b60056020526000908152604090205481565b6102266103ca3660046130b8565b60096020526000908152604090205481565b600154610348906001600160a01b031681565b61021b6103fd3660046130b8565b6125c0565b61021b61041036600461307e565b6125ef565b6101f2610423366004613157565b61263b565b61044e6104363660046130b8565b60076020526000908152604090205463ffffffff1681565b60405163ffffffff90911681526020016101ff565b6104866104713660046130b8565b60046020526000908152604090205460ff1681565b60405190151581526020016101ff565b6104a96104a4366004613230565b61267f565b6040516101ff91906133a4565b6101f26104c43660046130b8565b6127af565b61044e6104d73660046130b8565b60086020526000908152604090205463ffffffff1681565b60405160609060009061050a908290819086906020016132e4565b60408051601f19818403018152919052805160208201819020919250600160fd1b908390600a90811061053f5761053f6137be565b60200101906001600160f81b031916908160001a90535081516020830181902090600160fe1b908490600a908110610579576105796137be565b60200101906001600160f81b031916908160001a905350825160208401206040805160518082526080820190925260009160208201818036833701905050905060005b602081101561065f578481602081106105d7576105d76137be565b1a60f81b8282815181106105ed576105ed6137be565b60200101906001600160f81b031916908160001a905350838160208110610616576106166137be565b1a60f81b82610626836020613493565b81518110610636576106366137be565b60200101906001600160f81b031916908160001a9053508061065781613712565b9150506105bc565b5060005b60118110156106c75782816020811061067e5761067e6137be565b1a60f81b8261068e836040613493565b8151811061069e5761069e6137be565b60200101906001600160f81b031916908160001a905350806106bf81613712565b915050610663565b509695505050505050565b6001546001600160a01b031633146107055760405162461bcd60e51b81526004016106fc9061340d565b60405180910390fd5b601055565b8151602083012060609061071e8382613493565b6000818152600660205260409020549091508061074e5760405180602001604052806000815250925050506107ea565b60008181526003602052604090208054610767906136b5565b80601f0160208091040260200160405190810160405280929190818152602001828054610793906136b5565b80156107e05780601f106107b5576101008083540402835291602001916107e0565b820191906000526020600020905b8154815290600101906020018083116107c357829003601f168201915b5050505050925050505b92915050565b80600081905060018151101561080557600080fd5b600160005b825181101561089657828181518110610825576108256137be565b60209101015160f81c6039148061087b5750604183828151811061084b5761084b6137be565b016020015160f81c1080159061087b5750605a838281518110610870576108706137be565b016020015160f81c11155b61088457600091505b8061088e81613712565b91505061080a565b50806108d55760405162461bcd60e51b815260206004820152600e60248201526d494e56414c49442054525954455360901b60448201526064016106fc565b6001546001600160a01b031633146108ff5760405162461bcd60e51b81526004016106fc9061340d565b83518490610a71146109535760405162461bcd60e51b815260206004820152601a60248201527f494e56414c4944205452414e53414354494f4e204c454e47544800000000000060448201526064016106fc565b61095b612e9e565b6020810186905261096b866104ef565b6101a082018190528051602091820120610180830181905281830151600091825260028352604090912081516109a693919290910190612f2a565b506101a0810151610180820151600090815260036020908152604090912082516109d593919290910190612f2a565b506109e58261097e6109cf6127c8565b80516020909101206101c0820152610a02826109cf610a206127c8565b80516020918201206101e083019081526101c0830151610180840180516000908152600e8552604080822093909355925190518352600f8452912055810151610a509061088b6108dc612893565b8152610a618261092d61097e6127c8565b6101008201819052815180516020918201208251919092012061018083015160405182917fa551994b821690876343950f8de610eb0796baaded797266283c842d257dea9f91610ab391815260200190565b60405180910390a2817f3a3fefec9be0b15f11bf30dc166021af2910538b9322c12e44963bbbe72f7d7b846101800151604051610af291815260200190565b60405180910390a2610b098184610180015161297e565b610b18828461018001516129e8565b5050505050505050565b805160009081905b8015610bd1576000610b5e85610b4160018561363a565b81518110610b5157610b516137be565b016020015160f81c61267f565b604081015190915060010b610b748460036134ff565b610b7e9190613435565b602082015190935060010b610b948460036134ff565b610b9e9190613435565b815190935060010b610bb18460036134ff565b610bbb9190613435565b9250508080610bc99061369e565b915050610b2a565b5092915050565b60008054604051630e7fdbb160e01b81526001600160a01b0390911690630e7fdbb190610c099085906004016133d8565b60206040518083038186803b158015610c2157600080fd5b505afa158015610c35573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107ea9190613254565b6001546001600160a01b03163314610c835760405162461bcd60e51b81526004016106fc9061340d565b600054604051638756ee1960e01b81526001600160a01b0390911690638756ee1990610cb590859085906004016133eb565b600060405180830381600087803b158015610ccf57600080fd5b505af1158015610ce3573d6000803e3d6000fd5b505050505050565b60026020526000908152604090208054610d04906136b5565b80601f0160208091040260200160405190810160405280929190818152602001828054610d30906136b5565b8015610d7d5780601f10610d5257610100808354040283529160200191610d7d565b820191906000526020600020905b815481529060010190602001808311610d6057829003601f168201915b505050505081565b6001546001600160a01b03163314610daf5760405162461bcd60e51b81526004016106fc9061340d565b60008b5161032a14610e375760405162461bcd60e51b815260206004820152604560248201527f537472696e67206d75737420636f6e7461696e2031302061646472657373657360448201527f20776974686f757420636865636b73756d20666f722062756c6b2070726f63656064820152647373696e6760d81b608482015260a4016106fc565b6000546001600160a01b0316638756ee19610e5d8e84610e588160516134ab565b612893565b8d6040518363ffffffff1660e01b8152600401610e7b9291906133eb565b600060405180830381600087803b158015610e9557600080fd5b505af1158015610ea9573d6000803e3d6000fd5b50505050605181610eba91906134ab565b6000549091506001600160a01b0316638756ee19610ede8e84610e588160516134ab565b8c6040518363ffffffff1660e01b8152600401610efc9291906133eb565b600060405180830381600087803b158015610f1657600080fd5b505af1158015610f2a573d6000803e3d6000fd5b50505050605181610f3b91906134ab565b6000549091506001600160a01b0316638756ee19610f5f8e84610e588160516134ab565b8b6040518363ffffffff1660e01b8152600401610f7d9291906133eb565b600060405180830381600087803b158015610f9757600080fd5b505af1158015610fab573d6000803e3d6000fd5b50505050605181610fbc91906134ab565b6000549091506001600160a01b0316638756ee19610fe08e84610e588160516134ab565b8a6040518363ffffffff1660e01b8152600401610ffe9291906133eb565b600060405180830381600087803b15801561101857600080fd5b505af115801561102c573d6000803e3d6000fd5b5050505060518161103d91906134ab565b6000549091506001600160a01b0316638756ee196110618e84610e588160516134ab565b896040518363ffffffff1660e01b815260040161107f9291906133eb565b600060405180830381600087803b15801561109957600080fd5b505af11580156110ad573d6000803e3d6000fd5b505050506051816110be91906134ab565b6000549091506001600160a01b0316638756ee196110e28e84610e588160516134ab565b886040518363ffffffff1660e01b81526004016111009291906133eb565b600060405180830381600087803b15801561111a57600080fd5b505af115801561112e573d6000803e3d6000fd5b5050505060518161113f91906134ab565b6000549091506001600160a01b0316638756ee196111638e84610e588160516134ab565b876040518363ffffffff1660e01b81526004016111819291906133eb565b600060405180830381600087803b15801561119b57600080fd5b505af11580156111af573d6000803e3d6000fd5b505050506051816111c091906134ab565b6000549091506001600160a01b0316638756ee196111e48e84610e588160516134ab565b866040518363ffffffff1660e01b81526004016112029291906133eb565b600060405180830381600087803b15801561121c57600080fd5b505af1158015611230573d6000803e3d6000fd5b5050505060518161124191906134ab565b6000549091506001600160a01b0316638756ee196112658e84610e588160516134ab565b856040518363ffffffff1660e01b81526004016112839291906133eb565b600060405180830381600087803b15801561129d57600080fd5b505af11580156112b1573d6000803e3d6000fd5b505050506051816112c291906134ab565b6000549091506001600160a01b0316638756ee196112e68e84610e588160516134ab565b846040518363ffffffff1660e01b81526004016113049291906133eb565b600060405180830381600087803b15801561131e57600080fd5b505af1158015611332573d6000803e3d6000fd5b50505050505050505050505050505050565b606081600081905060018151101561135b57600080fd5b600160005b82518110156113ec5782818151811061137b5761137b6137be565b60209101015160f81c603914806113d1575060418382815181106113a1576113a16137be565b016020015160f81c108015906113d15750605a8382815181106113c6576113c66137be565b016020015160f81c11155b6113da57600091505b806113e481613712565b915050611360565b508061142b5760405162461bcd60e51b815260206004820152600e60248201526d494e56414c49442054525954455360901b60448201526064016106fc565b60145460ff16156114735760405162461bcd60e51b81526020600482015260126024820152717265656e7472792070726576656e7465642160701b60448201526064016106fc565b6014805460ff191660011790558451859061149190610a7190613768565b15801561149f575060008151115b6114eb5760405162461bcd60e51b815260206004820152601d60248201527f496e76616c6964207472616e73616374696f6e287329206c656e67746800000060448201526064016106fc565b6000610a7182516114fc91906134eb565b9050600061150b8260a2613584565b61ffff1667ffffffffffffffff811115611527576115276137d4565b6040519080825280601f01601f191660200182016040528015611551576020820181803683370190505b506040805160208101909152600080825291925060001990825b8561ffff168163ffffffff16101561213b57611585612e9e565b61159182610a716135ae565b63ffffffff1661016082018190526115b1908e90610e5881610a716134ab565b602082018190526115c1906104ef565b6101a082015260005b60036011546115d991906134eb565b81101561167b5760006115ed82605061363a565b9050826101a001518181518110611606576116066137be565b6020910101516001600160f81b031916603960f81b146116685760405162461bcd60e51b815260206004820152601d60248201527f5452414e53414354494f4e20504f57204e4f5420434f4d504c4554454400000060448201526064016106fc565b508061167381613712565b9150506115ca565b506101a081015180516020918201206101808301819052600090815260029091526040902080546116ab906136b5565b1590506116fa5760405162461bcd60e51b815260206004820152601d60248201527f5452414e53414354494f4e20414c52454144592050524f43455353454400000060448201526064016106fc565b60208082015161018083015160009081526002835260409020815161172493919290910190612f2a565b506101a08101516101808201516000908152600360209081526040909120825161175393919290910190612f2a565b50600b805490600061176483613712565b909155505061018081018051600b8054600090815260096020908152604080832094909455915493518152600a90915220556101608101516117dd9089906117ae9061097e6134ab565b63ffffffff1683610160015161097e6117c791906134ab565b6117d29060516134ab565b63ffffffff166127c8565b80516020918201206101808301516000908152600e90925260409091205561016081015161182c908990611813906109cf6134ab565b63ffffffff168361016001516109cf6117c791906134ab565b8051602091820120610180830180516000908152600f8452604080822093909355905163ffffffff8616808352601390945291902055611870576101808101516010555b60005b60a28161ffff16101561192257888161ffff1683610160015161088b61189991906134ab565b6118a391906134ab565b63ffffffff16815181106118b9576118b96137be565b01602001516001600160f81b0319168761ffff83166118d98660a26135ae565b6118e391906134ab565b63ffffffff16815181106118f9576118f96137be565b60200101906001600160f81b031916908160001a9053508061191a816136f0565b915050611873565b50611936816020015161088b6108dc612893565b80825260405160009160129161194c9190613299565b9081526020016040518091039020819055506119916102b1898361016001516108dc61197891906134ab565b63ffffffff168461016001516108e76117d291906134ab565b604082018190526119a29086613435565b94506608e1bc9bf03fff1981604001511215611a005760405162461bcd60e51b815260206004820152601960248201527f7472616e73616374696f6e2076616c756520746f6f206c6f770000000000000060448201526064016106fc565b6608e1bc9bf0400081604001511315611a5b5760405162461bcd60e51b815260206004820152601a60248201527f7472616e73616374696f6e2076616c756520746f6f206869676800000000000060448201526064016106fc565b6608e1bc9bf03fff19851215611ac85760405162461bcd60e51b815260206004820152602c60248201527f62756e646c652063756d6d756c6174697665207472616e73616374696f6e207660448201526b616c756520746f6f206c6f7760a01b60648201526084016106fc565b6608e1bc9bf04000851315611b355760405162461bcd60e51b815260206004820152602d60248201527f62756e646c652063756d6d756c6174697665207472616e73616374696f6e207660448201526c0c2d8eaca40e8dede40d0d2ced609b1b60648201526084016106fc565b611b6588826101600151610924611b4c91906134ab565b63ffffffff1683610160015161092d6117d291906134ab565b60e0820152610160810151611b9b908990611b829061092d6134ab565b63ffffffff1683610160015161092d6117c791906134ab565b61010082015263ffffffff8216611bc8578061010001519250611bc18160e00151610b22565b9350611c76565b83611bd68260e00151610b22565b14611c235760405162461bcd60e51b815260206004820152601860248201527f6c617374496e646578206e6f7420636f6e73697374656e74000000000000000060448201526064016106fc565b611c3283826101000151612a52565b611c765760405162461bcd60e51b8152602060048201526015602482015274189d5b991b19481b9bdd0818dbdb9cda5cdd195b9d605a1b60448201526064016106fc565b611ca68882610160015161091b611c8d91906134ab565b63ffffffff168361016001516109246117d291906134ab565b6101208201819052611cb790610b22565b610140820181905263ffffffff831614611d135760405162461bcd60e51b815260206004820152601c60248201527f7472616e73616374696f6e2073657175656e636520696e76616c69640000000060448201526064016106fc565b611d1e600188613617565b61ffff168263ffffffff161415611d8457838263ffffffff1614611d845760405162461bcd60e51b815260206004820152601d60248201527f6c617374207472616e73616374696f6e20213d206c617374496e64657800000060448201526064016106fc565b60008160400151121561201b57838263ffffffff1610611dde5760405162461bcd60e51b81526020600482015260156024820152746d697373696e6720326e64207369676e617475726560581b60448201526064016106fc565b611e1f8d826101600151610a71611df591906134ab565b611e019061088b6134ab565b610160840151611e1390610a716134ab565b610e58906108dc6134ab565b6060820152610160810151611e71906102b1908a90611e4090610a716134ab565b611e4c906108dc6134ab565b63ffffffff16846101600151610a71611e6591906134ab565b6117d2906108e76134ab565b608082015280516060820151611e879190612a52565b611edd5760405162461bcd60e51b815260206004820152602160248201527f326e64207369676e61747572652068617320696e76616c6964206164647265736044820152607360f81b60648201526084016106fc565b608081015115611f2f5760405162461bcd60e51b815260206004820152601c60248201527f326e64207369676e6174757265206973206e6f7420302076616c75650000000060448201526064016106fc565b611f578d8261016001516000611f4591906134ab565b610160840151610e589061088b6134ab565b60a0820152610160810151611f92908e90611f7490610a716134ab565b610160840151611f8690610a716134ab565b610e589061088b6134ab565b60c08201819052815160a0830151611fb492611fae9190612aab565b85612ad7565b611ff45760405162461bcd60e51b8152602060048201526011602482015270494e56414c4944205349474e415455524560781b60448201526064016106fc565b606081015180516020918201206000908152600490915260409020805460ff191660011790555b805160405161202a9190613299565b60405180910390207f6308696fb14e1bcd3fcde7330f7607d979f745988c8b9ffaaa407c856e64b8df826040015160405161206791815260200190565b60405180910390a261010081015180516020918201208251805192019190912061018083015160405183917fa551994b821690876343950f8de610eb0796baaded797266283c842d257dea9f916120c091815260200190565b60405180910390a2807f3a3fefec9be0b15f11bf30dc166021af2910538b9322c12e44963bbbe72f7d7b8461018001516040516120ff91815260200190565b60405180910390a2612116818461018001516129e8565b6121258284610180015161297e565b50505080806121339061372d565b91505061156b565b5060005b8561ffff168163ffffffff16101561229057612159612e9e565b63ffffffff82166000908152601360209081526040808320546101808501818152908452600e8352818420546101c0860190815290518452600f8352818420546101e08601525183526002909152902080546121b4906136b5565b151590506122045760405162461bcd60e51b815260206004820181905260248201527f5452554e4b205452414e53414354494f4e20444f4553204e4f5420455849535460448201526064016106fc565b6101e081015160009081526002602052604090208054612223906136b5565b1515905061227d5760405162461bcd60e51b815260206004820152602160248201527f4252414e4348205452414e53414354494f4e20444f4553204e4f5420455849536044820152601560fa1b60648201526084016106fc565b50806122888161372d565b91505061213f565b50821561229f5761229f61377c565b60006122aa856104ef565b90506122b68282612a52565b6123025760405162461bcd60e51b815260206004820152601960248201527f43414c43554c415445442042554e444c4520444946464552530000000000000060448201526064016106fc565b61230b8c612b11565b5060005b8661ffff168163ffffffff16101561255a57612329612e9e565b61233582610a716135ae565b63ffffffff166101608201819052612359906102b1908b90611978906108dc6134ab565b6040820181905260001215612547576123918e82610160015161088b61237f91906134ab565b610160840151610e58906108dc6134ab565b81526101608101516123aa908f90611f459060006134ab565b60a082015260008054825160405163a205a8a960e01b81526001600160a01b039092169163a205a8a9916123e0916004016133d8565b60206040518083038186803b1580156123f857600080fd5b505afa15801561240c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612430919061309b565b9050600080826001600160a01b03166000612710908660a0015187604001516040516024016124609291906133eb565b60408051601f198184030181529181526020820180516001600160e01b031663167ba96d60e11b179052516124959190613299565b600060405180830381858888f193505050503d80600081146124d3576040519150601f19603f3d011682016040523d82523d6000602084013e6124d8565b606091505b5091509150816125435760405162461bcd60e51b815260206004820152603060248201527f43414c4c454420434f4e5452414354202841444b5472616e73616374696f6e4e60448201526f37ba34b33c94902922ab22a92a22a21760811b60648201526084016106fc565b5050505b50806125528161372d565b91505061230f565b506014805460ff191690559b9a5050505050505050505050565b6001546001600160a01b0316331461259e5760405162461bcd60e51b81526004016106fc9061340d565b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b031633146125ea5760405162461bcd60e51b81526004016106fc9061340d565b601155565b6001546001600160a01b031633146126195760405162461bcd60e51b81526004016106fc9061340d565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b8151602083012060609061264f8382613493565b6000818152600560205260409020549091508061074e5760405180602001604052806000815250925050506107ea565b612687612fae565b8161ffff16603914156126b6575050604080516060810182526000808252602082018190529181019190915290565b60418261ffff16101580156126d05750605a8261ffff1611155b61270c5760405162461bcd60e51b815260206004820152600d60248201526c494e56414c494420545259544560981b60448201526064016106fc565b6000612719836003613476565b9050604051806060016040528060016003846127359190613747565b61273f91906135d1565b60010b60010b815260200160016003808561275a91906134ca565b6127649190613747565b61276e91906135d1565b60010b60010b81526020016001600360098561278a91906134ca565b6127949190613747565b61279e91906135d1565b60010b60010b815250915050919050565b60036020526000908152604090208054610d04906136b5565b606060006127d6848461363a565b67ffffffffffffffff8111156127ee576127ee6137d4565b6040519080825280601f01601f191660200182016040528015612818576020820181803683370190505b509050835b8381101561288a57858181518110612837576128376137be565b01602001516001600160f81b03191682612851878461363a565b81518110612861576128616137be565b60200101906001600160f81b031916908160001a9053508061288281613712565b91505061281d565b50949350505050565b60608360006128a28585613651565b63ffffffff1667ffffffffffffffff8111156128c0576128c06137d4565b6040519080825280601f01601f1916602001820160405280156128ea576020820181803683370190505b509050845b8463ffffffff168163ffffffff16101561297457828163ffffffff168151811061291b5761291b6137be565b01602001516001600160f81b031916826129358884613651565b63ffffffff168151811061294b5761294b6137be565b60200101906001600160f81b031916908160001a9053508061296c8161372d565b9150506128ef565b5095945050505050565b60008281526008602052604081205463ffffffff169061299e8285613493565b60008181526006602052604090208490559050816129bb8161372d565b600095865260086020526040909520805463ffffffff191663ffffffff9096169590951790945550505050565b60008281526007602052604081205463ffffffff1690612a088285613493565b6000818152600560205260409020849055905081612a258161372d565b600095865260076020526040909520805463ffffffff191663ffffffff9096169590951790945550505050565b600081604051602001612a659190613299565b6040516020818303038152906040528051906020012083604051602001612a8c9190613299565b6040516020818303038152906040528051906020012014905092915050565b60608282604051602001612ac09291906132b5565b604051602081830303815290604052905092915050565b600080848484604051602001612aef9392919061333a565b60408051808303601f1901815291905280516020909101201595945050505050565b601454600090610100900460ff1615612b865760405162461bcd60e51b815260206004820152603160248201527f6d7574657820636865636b206661696c6564206f6e2056616c696461746542616044820152701b185b98d95cd05b99151c985b9cd858dd607a1b60648201526084016106fc565b6014805461ff00191661010017905581518290612ba690610a7190613768565b158015612bb4575060008151115b612c005760405162461bcd60e51b815260206004820152601d60248201527f496e76616c6964207472616e73616374696f6e287329206c656e67746800000060448201526064016106fc565b6000610a718251612c1191906134eb565b90506000805b8261ffff168163ffffffff161015612cd9576000612c3782610a716135ae565b90506000612c5788612c4b8461088b6134ab565b610e58856108dc6134ab565b90506000612c806102b188612c6e866108dc6134ab565b63ffffffff166117d2876108e76134ab565b9050612c8c8186613435565b945080601283604051612c9f9190613299565b90815260200160405180910390206000828254612cbc9190613435565b925050819055505050508080612cd19061372d565b915050612c17565b508015612ce857612ce861377c565b60005b8261ffff168163ffffffff161015612e87576000612d0b82610a716135ae565b90506000612d1f88612c4b8461088b6134ab565b90506000612d2c82610bd8565b90506000601283604051612d409190613299565b90815260200160405180910390205482612d5a9190613435565b1215612d9f5760405162461bcd60e51b8152602060048201526014602482015273494e53554646494349454e542042414c414e434560601b60448201526064016106fc565b601282604051612daf9190613299565b908152602001604051809103902054600014612e4e576000546040516001600160a01b0390911690638756ee19908490601290612ded908390613299565b908152604051908190036020018120546001600160e01b031960e085901b168252612e1b92916004016133eb565b600060405180830381600087803b158015612e3557600080fd5b505af1158015612e49573d6000803e3d6000fd5b505050505b6000601283604051612e609190613299565b9081526040519081900360200190205550829150612e7f90508161372d565b915050612ceb565b50506014805461ff00191690555060019392505050565b6040518061020001604052806060815260200160608152602001600081526020016060815260200160008152602001606081526020016060815260200160608152602001606081526020016060815260200160008152602001600063ffffffff168152602001600080191681526020016060815260200160008019168152602001600080191681525090565b828054612f36906136b5565b90600052602060002090601f016020900481019282612f585760008555612f9e565b82601f10612f7157805160ff1916838001178555612f9e565b82800160010185558215612f9e579182015b82811115612f9e578251825591602001919060010190612f83565b50612faa929150612fcc565b5090565b60405180606001604052806003906020820280368337509192915050565b5b80821115612faa5760008155600101612fcd565b600067ffffffffffffffff80841115612ffc57612ffc6137d4565b604051601f8501601f19908116603f01168101908282118183101715613024576130246137d4565b8160405280935085815286868601111561303d57600080fd5b858560208301376000602087830101525050509392505050565b600082601f83011261306857600080fd5b61307783833560208501612fe1565b9392505050565b60006020828403121561309057600080fd5b8135613077816137ea565b6000602082840312156130ad57600080fd5b8151613077816137ea565b6000602082840312156130ca57600080fd5b5035919050565b6000602082840312156130e357600080fd5b813567ffffffffffffffff8111156130fa57600080fd5b8201601f8101841361310b57600080fd5b61311a84823560208401612fe1565b949350505050565b60006020828403121561313457600080fd5b813567ffffffffffffffff81111561314b57600080fd5b61311a84828501613057565b6000806040838503121561316a57600080fd5b823567ffffffffffffffff81111561318157600080fd5b61318d85828601613057565b95602094909401359450505050565b60008060008060008060008060008060006101608c8e0312156131be57600080fd5b8b3567ffffffffffffffff8111156131d557600080fd5b6131e18e828f01613057565b9e60208e01359e5060408e01359d60608101359d5060808101359c5060a08101359b5060c08101359a5060e0810135995061010081013598506101208101359750610140013595509350505050565b60006020828403121561324257600080fd5b813561ffff8116811461307757600080fd5b60006020828403121561326657600080fd5b5051919050565b6000815180845261328581602086016020860161366e565b601f01601f19169290920160200192915050565b600082516132ab81846020870161366e565b9190910192915050565b600083516132c781846020880161366e565b8351908301906132db81836020880161366e565b01949350505050565b657b4355524c7d60d01b81526309082a6960e31b60068201526001600160f81b0319848116600a8301528316600b820152815160009061332b81600c85016020870161366e565b91909101600c01949350505050565b657b4355524c7d60d01b81526556414c53494760d01b60068201526000845161336a81600c85016020890161366e565b84519083019061338181600c84016020890161366e565b845191019061339781600c84016020880161366e565b01600c0195945050505050565b60608101818360005b60038110156133cf578151600190810b845260209384019390920191016133ad565b50505092915050565b602081526000613077602083018461326d565b6040815260006133fe604083018561326d565b90508260208301529392505050565b6020808252600e908201526d1393d5081055551213d49256915160921b604082015260600190565b600080821280156001600160ff1b038490038513161561345757613457613792565b600160ff1b839003841281161561347057613470613792565b50500190565b600061ffff8083168185168083038211156132db576132db613792565b600082198211156134a6576134a6613792565b500190565b600063ffffffff8083168185168083038211156132db576132db613792565b600061ffff808416806134df576134df6137a8565b92169190910492915050565b6000826134fa576134fa6137a8565b500490565b60006001600160ff1b038184138284138082168684048611161561352557613525613792565b600160ff1b600087128281168783058912161561354457613544613792565b6000871292508782058712848416161561356057613560613792565b8785058712818416161561357657613576613792565b505050929093029392505050565b600061ffff808316818516818304811182151516156135a5576135a5613792565b02949350505050565b600063ffffffff808316818516818304811182151516156135a5576135a5613792565b60008160010b8360010b6000811281617fff19018312811516156135f7576135f7613792565b81617fff01831381161561360d5761360d613792565b5090039392505050565b600061ffff8381169083168181101561363257613632613792565b039392505050565b60008282101561364c5761364c613792565b500390565b600063ffffffff8381169083168181101561363257613632613792565b60005b83811015613689578181015183820152602001613671565b83811115613698576000848401525b50505050565b6000816136ad576136ad613792565b506000190190565b600181811c908216806136c957607f821691505b602082108114156136ea57634e487b7160e01b600052602260045260246000fd5b50919050565b600061ffff8083168181141561370857613708613792565b6001019392505050565b600060001982141561372657613726613792565b5060010190565b600063ffffffff8083168181141561370857613708613792565b600061ffff8084168061375c5761375c6137a8565b92169190910692915050565b600082613777576137776137a8565b500690565b634e487b7160e01b600052600160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146137ff57600080fd5b5056fea264697066735822122021b5fd776dd2819ea1ff3e69b5cb5b5d4195dd841efbafd0db803fa6fa1457ad64736f6c63430008070033393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939",
}

// ADKTransactionsABI is the input ABI used to generate the binding from.
// Deprecated: Use ADKTransactionsMetaData.ABI instead.
var ADKTransactionsABI = ADKTransactionsMetaData.ABI

// Deprecated: Use ADKTransactionsMetaData.Sigs instead.
// ADKTransactionsFuncSigs maps the 4-byte function signature to its string representation.
var ADKTransactionsFuncSigs = ADKTransactionsMetaData.Sigs

// ADKTransactionsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ADKTransactionsMetaData.Bin instead.
var ADKTransactionsBin = ADKTransactionsMetaData.Bin

// DeployADKTransactions deploys a new Ethereum contract, binding an instance of ADKTransactions to it.
func DeployADKTransactions(auth *bind.TransactOpts, backend bind.ContractBackend, _genesis_account common.Address, _ADKTokenContract common.Address) (common.Address, *types.Transaction, *ADKTransactions, error) {
	parsed, err := ADKTransactionsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ADKTransactionsBin), backend, _genesis_account, _ADKTokenContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ADKTransactions{ADKTransactionsCaller: ADKTransactionsCaller{contract: contract}, ADKTransactionsTransactor: ADKTransactionsTransactor{contract: contract}, ADKTransactionsFilterer: ADKTransactionsFilterer{contract: contract}}, nil
}

// ADKTransactions is an auto generated Go binding around an Ethereum contract.
type ADKTransactions struct {
	ADKTransactionsCaller     // Read-only binding to the contract
	ADKTransactionsTransactor // Write-only binding to the contract
	ADKTransactionsFilterer   // Log filterer for contract events
}

// ADKTransactionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ADKTransactionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ADKTransactionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ADKTransactionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ADKTransactionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ADKTransactionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ADKTransactionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ADKTransactionsSession struct {
	Contract     *ADKTransactions  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ADKTransactionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ADKTransactionsCallerSession struct {
	Contract *ADKTransactionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ADKTransactionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ADKTransactionsTransactorSession struct {
	Contract     *ADKTransactionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ADKTransactionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ADKTransactionsRaw struct {
	Contract *ADKTransactions // Generic contract binding to access the raw methods on
}

// ADKTransactionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ADKTransactionsCallerRaw struct {
	Contract *ADKTransactionsCaller // Generic read-only contract binding to access the raw methods on
}

// ADKTransactionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ADKTransactionsTransactorRaw struct {
	Contract *ADKTransactionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewADKTransactions creates a new instance of ADKTransactions, bound to a specific deployed contract.
func NewADKTransactions(address common.Address, backend bind.ContractBackend) (*ADKTransactions, error) {
	contract, err := bindADKTransactions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ADKTransactions{ADKTransactionsCaller: ADKTransactionsCaller{contract: contract}, ADKTransactionsTransactor: ADKTransactionsTransactor{contract: contract}, ADKTransactionsFilterer: ADKTransactionsFilterer{contract: contract}}, nil
}

// NewADKTransactionsCaller creates a new read-only instance of ADKTransactions, bound to a specific deployed contract.
func NewADKTransactionsCaller(address common.Address, caller bind.ContractCaller) (*ADKTransactionsCaller, error) {
	contract, err := bindADKTransactions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ADKTransactionsCaller{contract: contract}, nil
}

// NewADKTransactionsTransactor creates a new write-only instance of ADKTransactions, bound to a specific deployed contract.
func NewADKTransactionsTransactor(address common.Address, transactor bind.ContractTransactor) (*ADKTransactionsTransactor, error) {
	contract, err := bindADKTransactions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ADKTransactionsTransactor{contract: contract}, nil
}

// NewADKTransactionsFilterer creates a new log filterer instance of ADKTransactions, bound to a specific deployed contract.
func NewADKTransactionsFilterer(address common.Address, filterer bind.ContractFilterer) (*ADKTransactionsFilterer, error) {
	contract, err := bindADKTransactions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ADKTransactionsFilterer{contract: contract}, nil
}

// bindADKTransactions binds a generic wrapper to an already deployed contract.
func bindADKTransactions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ADKTransactionsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ADKTransactions *ADKTransactionsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ADKTransactions.Contract.ADKTransactionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ADKTransactions *ADKTransactionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ADKTransactions.Contract.ADKTransactionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ADKTransactions *ADKTransactionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ADKTransactions.Contract.ADKTransactionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ADKTransactions *ADKTransactionsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ADKTransactions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ADKTransactions *ADKTransactionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ADKTransactions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ADKTransactions *ADKTransactionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ADKTransactions.Contract.contract.Transact(opts, method, params...)
}

// ADKTokenAddress is a free data retrieval call binding the contract method 0xa91fa7cb.
//
// Solidity: function ADKTokenAddress() view returns(address)
func (_ADKTransactions *ADKTransactionsCaller) ADKTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "ADKTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADKTokenAddress is a free data retrieval call binding the contract method 0xa91fa7cb.
//
// Solidity: function ADKTokenAddress() view returns(address)
func (_ADKTransactions *ADKTransactionsSession) ADKTokenAddress() (common.Address, error) {
	return _ADKTransactions.Contract.ADKTokenAddress(&_ADKTransactions.CallOpts)
}

// ADKTokenAddress is a free data retrieval call binding the contract method 0xa91fa7cb.
//
// Solidity: function ADKTokenAddress() view returns(address)
func (_ADKTransactions *ADKTransactionsCallerSession) ADKTokenAddress() (common.Address, error) {
	return _ADKTransactions.Contract.ADKTokenAddress(&_ADKTransactions.CallOpts)
}

// ClientProofOfWorkRequirement is a free data retrieval call binding the contract method 0x2076e796.
//
// Solidity: function ClientProofOfWorkRequirement() view returns(uint256)
func (_ADKTransactions *ADKTransactionsCaller) ClientProofOfWorkRequirement(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "ClientProofOfWorkRequirement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClientProofOfWorkRequirement is a free data retrieval call binding the contract method 0x2076e796.
//
// Solidity: function ClientProofOfWorkRequirement() view returns(uint256)
func (_ADKTransactions *ADKTransactionsSession) ClientProofOfWorkRequirement() (*big.Int, error) {
	return _ADKTransactions.Contract.ClientProofOfWorkRequirement(&_ADKTransactions.CallOpts)
}

// ClientProofOfWorkRequirement is a free data retrieval call binding the contract method 0x2076e796.
//
// Solidity: function ClientProofOfWorkRequirement() view returns(uint256)
func (_ADKTransactions *ADKTransactionsCallerSession) ClientProofOfWorkRequirement() (*big.Int, error) {
	return _ADKTransactions.Contract.ClientProofOfWorkRequirement(&_ADKTransactions.CallOpts)
}

// CurlHashOP is a free data retrieval call binding the contract method 0x016d1c68.
//
// Solidity: function CurlHashOP(string str) pure returns(string)
func (_ADKTransactions *ADKTransactionsCaller) CurlHashOP(opts *bind.CallOpts, str string) (string, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "CurlHashOP", str)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CurlHashOP is a free data retrieval call binding the contract method 0x016d1c68.
//
// Solidity: function CurlHashOP(string str) pure returns(string)
func (_ADKTransactions *ADKTransactionsSession) CurlHashOP(str string) (string, error) {
	return _ADKTransactions.Contract.CurlHashOP(&_ADKTransactions.CallOpts, str)
}

// CurlHashOP is a free data retrieval call binding the contract method 0x016d1c68.
//
// Solidity: function CurlHashOP(string str) pure returns(string)
func (_ADKTransactions *ADKTransactionsCallerSession) CurlHashOP(str string) (string, error) {
	return _ADKTransactions.Contract.CurlHashOP(&_ADKTransactions.CallOpts, str)
}

// GetAZ9balanceOf is a free data retrieval call binding the contract method 0x55b6193c.
//
// Solidity: function GetAZ9balanceOf(string _adkAddr) view returns(uint256 balance)
func (_ADKTransactions *ADKTransactionsCaller) GetAZ9balanceOf(opts *bind.CallOpts, _adkAddr string) (*big.Int, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "GetAZ9balanceOf", _adkAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAZ9balanceOf is a free data retrieval call binding the contract method 0x55b6193c.
//
// Solidity: function GetAZ9balanceOf(string _adkAddr) view returns(uint256 balance)
func (_ADKTransactions *ADKTransactionsSession) GetAZ9balanceOf(_adkAddr string) (*big.Int, error) {
	return _ADKTransactions.Contract.GetAZ9balanceOf(&_ADKTransactions.CallOpts, _adkAddr)
}

// GetAZ9balanceOf is a free data retrieval call binding the contract method 0x55b6193c.
//
// Solidity: function GetAZ9balanceOf(string _adkAddr) view returns(uint256 balance)
func (_ADKTransactions *ADKTransactionsCallerSession) GetAZ9balanceOf(_adkAddr string) (*big.Int, error) {
	return _ADKTransactions.Contract.GetAZ9balanceOf(&_ADKTransactions.CallOpts, _adkAddr)
}

// GetTxByAddress is a free data retrieval call binding the contract method 0xd9fc5448.
//
// Solidity: function GetTxByAddress(string addressString, uint256 numTx) view returns(string)
func (_ADKTransactions *ADKTransactionsCaller) GetTxByAddress(opts *bind.CallOpts, addressString string, numTx *big.Int) (string, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "GetTxByAddress", addressString, numTx)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTxByAddress is a free data retrieval call binding the contract method 0xd9fc5448.
//
// Solidity: function GetTxByAddress(string addressString, uint256 numTx) view returns(string)
func (_ADKTransactions *ADKTransactionsSession) GetTxByAddress(addressString string, numTx *big.Int) (string, error) {
	return _ADKTransactions.Contract.GetTxByAddress(&_ADKTransactions.CallOpts, addressString, numTx)
}

// GetTxByAddress is a free data retrieval call binding the contract method 0xd9fc5448.
//
// Solidity: function GetTxByAddress(string addressString, uint256 numTx) view returns(string)
func (_ADKTransactions *ADKTransactionsCallerSession) GetTxByAddress(addressString string, numTx *big.Int) (string, error) {
	return _ADKTransactions.Contract.GetTxByAddress(&_ADKTransactions.CallOpts, addressString, numTx)
}

// GetTxByBundle is a free data retrieval call binding the contract method 0x49607e7d.
//
// Solidity: function GetTxByBundle(string bundleString, uint256 numTx) view returns(string)
func (_ADKTransactions *ADKTransactionsCaller) GetTxByBundle(opts *bind.CallOpts, bundleString string, numTx *big.Int) (string, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "GetTxByBundle", bundleString, numTx)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetTxByBundle is a free data retrieval call binding the contract method 0x49607e7d.
//
// Solidity: function GetTxByBundle(string bundleString, uint256 numTx) view returns(string)
func (_ADKTransactions *ADKTransactionsSession) GetTxByBundle(bundleString string, numTx *big.Int) (string, error) {
	return _ADKTransactions.Contract.GetTxByBundle(&_ADKTransactions.CallOpts, bundleString, numTx)
}

// GetTxByBundle is a free data retrieval call binding the contract method 0x49607e7d.
//
// Solidity: function GetTxByBundle(string bundleString, uint256 numTx) view returns(string)
func (_ADKTransactions *ADKTransactionsCallerSession) GetTxByBundle(bundleString string, numTx *big.Int) (string, error) {
	return _ADKTransactions.Contract.GetTxByBundle(&_ADKTransactions.CallOpts, bundleString, numTx)
}

// TryteToIntValue is a free data retrieval call binding the contract method 0x510bcf5f.
//
// Solidity: function TryteToIntValue(bytes data) pure returns(int256)
func (_ADKTransactions *ADKTransactionsCaller) TryteToIntValue(opts *bind.CallOpts, data []byte) (*big.Int, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "TryteToIntValue", data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TryteToIntValue is a free data retrieval call binding the contract method 0x510bcf5f.
//
// Solidity: function TryteToIntValue(bytes data) pure returns(int256)
func (_ADKTransactions *ADKTransactionsSession) TryteToIntValue(data []byte) (*big.Int, error) {
	return _ADKTransactions.Contract.TryteToIntValue(&_ADKTransactions.CallOpts, data)
}

// TryteToIntValue is a free data retrieval call binding the contract method 0x510bcf5f.
//
// Solidity: function TryteToIntValue(bytes data) pure returns(int256)
func (_ADKTransactions *ADKTransactionsCallerSession) TryteToIntValue(data []byte) (*big.Int, error) {
	return _ADKTransactions.Contract.TryteToIntValue(&_ADKTransactions.CallOpts, data)
}

// TryteToTrits is a free data retrieval call binding the contract method 0xf3067872.
//
// Solidity: function TryteToTrits(uint16 test) pure returns(int16[3])
func (_ADKTransactions *ADKTransactionsCaller) TryteToTrits(opts *bind.CallOpts, test uint16) ([3]int16, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "TryteToTrits", test)

	if err != nil {
		return *new([3]int16), err
	}

	out0 := *abi.ConvertType(out[0], new([3]int16)).(*[3]int16)

	return out0, err

}

// TryteToTrits is a free data retrieval call binding the contract method 0xf3067872.
//
// Solidity: function TryteToTrits(uint16 test) pure returns(int16[3])
func (_ADKTransactions *ADKTransactionsSession) TryteToTrits(test uint16) ([3]int16, error) {
	return _ADKTransactions.Contract.TryteToTrits(&_ADKTransactions.CallOpts, test)
}

// TryteToTrits is a free data retrieval call binding the contract method 0xf3067872.
//
// Solidity: function TryteToTrits(uint16 test) pure returns(int16[3])
func (_ADKTransactions *ADKTransactionsCallerSession) TryteToTrits(test uint16) ([3]int16, error) {
	return _ADKTransactions.Contract.TryteToTrits(&_ADKTransactions.CallOpts, test)
}

// AdkgoGenesisAddress is a free data retrieval call binding the contract method 0xc4c7e151.
//
// Solidity: function adkgo_genesis_address() view returns(address)
func (_ADKTransactions *ADKTransactionsCaller) AdkgoGenesisAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "adkgo_genesis_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdkgoGenesisAddress is a free data retrieval call binding the contract method 0xc4c7e151.
//
// Solidity: function adkgo_genesis_address() view returns(address)
func (_ADKTransactions *ADKTransactionsSession) AdkgoGenesisAddress() (common.Address, error) {
	return _ADKTransactions.Contract.AdkgoGenesisAddress(&_ADKTransactions.CallOpts)
}

// AdkgoGenesisAddress is a free data retrieval call binding the contract method 0xc4c7e151.
//
// Solidity: function adkgo_genesis_address() view returns(address)
func (_ADKTransactions *ADKTransactionsCallerSession) AdkgoGenesisAddress() (common.Address, error) {
	return _ADKTransactions.Contract.AdkgoGenesisAddress(&_ADKTransactions.CallOpts)
}

// MeshTip is a free data retrieval call binding the contract method 0x11cf02ec.
//
// Solidity: function meshTip() view returns(bytes32)
func (_ADKTransactions *ADKTransactionsCaller) MeshTip(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "meshTip")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MeshTip is a free data retrieval call binding the contract method 0x11cf02ec.
//
// Solidity: function meshTip() view returns(bytes32)
func (_ADKTransactions *ADKTransactionsSession) MeshTip() ([32]byte, error) {
	return _ADKTransactions.Contract.MeshTip(&_ADKTransactions.CallOpts)
}

// MeshTip is a free data retrieval call binding the contract method 0x11cf02ec.
//
// Solidity: function meshTip() view returns(bytes32)
func (_ADKTransactions *ADKTransactionsCallerSession) MeshTip() ([32]byte, error) {
	return _ADKTransactions.Contract.MeshTip(&_ADKTransactions.CallOpts)
}

// SpentAddresses is a free data retrieval call binding the contract method 0xde5ed777.
//
// Solidity: function spent_addresses(bytes32 ) view returns(bool)
func (_ADKTransactions *ADKTransactionsCaller) SpentAddresses(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "spent_addresses", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SpentAddresses is a free data retrieval call binding the contract method 0xde5ed777.
//
// Solidity: function spent_addresses(bytes32 ) view returns(bool)
func (_ADKTransactions *ADKTransactionsSession) SpentAddresses(arg0 [32]byte) (bool, error) {
	return _ADKTransactions.Contract.SpentAddresses(&_ADKTransactions.CallOpts, arg0)
}

// SpentAddresses is a free data retrieval call binding the contract method 0xde5ed777.
//
// Solidity: function spent_addresses(bytes32 ) view returns(bool)
func (_ADKTransactions *ADKTransactionsCallerSession) SpentAddresses(arg0 [32]byte) (bool, error) {
	return _ADKTransactions.Contract.SpentAddresses(&_ADKTransactions.CallOpts, arg0)
}

// TransactionBranch is a free data retrieval call binding the contract method 0x4893b64e.
//
// Solidity: function transaction_branch(bytes32 ) view returns(bytes32)
func (_ADKTransactions *ADKTransactionsCaller) TransactionBranch(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "transaction_branch", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TransactionBranch is a free data retrieval call binding the contract method 0x4893b64e.
//
// Solidity: function transaction_branch(bytes32 ) view returns(bytes32)
func (_ADKTransactions *ADKTransactionsSession) TransactionBranch(arg0 [32]byte) ([32]byte, error) {
	return _ADKTransactions.Contract.TransactionBranch(&_ADKTransactions.CallOpts, arg0)
}

// TransactionBranch is a free data retrieval call binding the contract method 0x4893b64e.
//
// Solidity: function transaction_branch(bytes32 ) view returns(bytes32)
func (_ADKTransactions *ADKTransactionsCallerSession) TransactionBranch(arg0 [32]byte) ([32]byte, error) {
	return _ADKTransactions.Contract.TransactionBranch(&_ADKTransactions.CallOpts, arg0)
}

// TransactionHashes is a free data retrieval call binding the contract method 0xf7a8ffd7.
//
// Solidity: function transaction_hashes(bytes32 ) view returns(string)
func (_ADKTransactions *ADKTransactionsCaller) TransactionHashes(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "transaction_hashes", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TransactionHashes is a free data retrieval call binding the contract method 0xf7a8ffd7.
//
// Solidity: function transaction_hashes(bytes32 ) view returns(string)
func (_ADKTransactions *ADKTransactionsSession) TransactionHashes(arg0 [32]byte) (string, error) {
	return _ADKTransactions.Contract.TransactionHashes(&_ADKTransactions.CallOpts, arg0)
}

// TransactionHashes is a free data retrieval call binding the contract method 0xf7a8ffd7.
//
// Solidity: function transaction_hashes(bytes32 ) view returns(string)
func (_ADKTransactions *ADKTransactionsCallerSession) TransactionHashes(arg0 [32]byte) (string, error) {
	return _ADKTransactions.Contract.TransactionHashes(&_ADKTransactions.CallOpts, arg0)
}

// TransactionIndex is a free data retrieval call binding the contract method 0x428a4e9c.
//
// Solidity: function transaction_index(bytes32 ) view returns(uint256)
func (_ADKTransactions *ADKTransactionsCaller) TransactionIndex(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "transaction_index", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransactionIndex is a free data retrieval call binding the contract method 0x428a4e9c.
//
// Solidity: function transaction_index(bytes32 ) view returns(uint256)
func (_ADKTransactions *ADKTransactionsSession) TransactionIndex(arg0 [32]byte) (*big.Int, error) {
	return _ADKTransactions.Contract.TransactionIndex(&_ADKTransactions.CallOpts, arg0)
}

// TransactionIndex is a free data retrieval call binding the contract method 0x428a4e9c.
//
// Solidity: function transaction_index(bytes32 ) view returns(uint256)
func (_ADKTransactions *ADKTransactionsCallerSession) TransactionIndex(arg0 [32]byte) (*big.Int, error) {
	return _ADKTransactions.Contract.TransactionIndex(&_ADKTransactions.CallOpts, arg0)
}

// TransactionIndexedBySeq is a free data retrieval call binding the contract method 0xc1cba762.
//
// Solidity: function transaction_indexed_by_seq(uint256 ) view returns(bytes32)
func (_ADKTransactions *ADKTransactionsCaller) TransactionIndexedBySeq(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "transaction_indexed_by_seq", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TransactionIndexedBySeq is a free data retrieval call binding the contract method 0xc1cba762.
//
// Solidity: function transaction_indexed_by_seq(uint256 ) view returns(bytes32)
func (_ADKTransactions *ADKTransactionsSession) TransactionIndexedBySeq(arg0 *big.Int) ([32]byte, error) {
	return _ADKTransactions.Contract.TransactionIndexedBySeq(&_ADKTransactions.CallOpts, arg0)
}

// TransactionIndexedBySeq is a free data retrieval call binding the contract method 0xc1cba762.
//
// Solidity: function transaction_indexed_by_seq(uint256 ) view returns(bytes32)
func (_ADKTransactions *ADKTransactionsCallerSession) TransactionIndexedBySeq(arg0 *big.Int) ([32]byte, error) {
	return _ADKTransactions.Contract.TransactionIndexedBySeq(&_ADKTransactions.CallOpts, arg0)
}

// TransactionTrunk is a free data retrieval call binding the contract method 0xad69e6a0.
//
// Solidity: function transaction_trunk(bytes32 ) view returns(bytes32)
func (_ADKTransactions *ADKTransactionsCaller) TransactionTrunk(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "transaction_trunk", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TransactionTrunk is a free data retrieval call binding the contract method 0xad69e6a0.
//
// Solidity: function transaction_trunk(bytes32 ) view returns(bytes32)
func (_ADKTransactions *ADKTransactionsSession) TransactionTrunk(arg0 [32]byte) ([32]byte, error) {
	return _ADKTransactions.Contract.TransactionTrunk(&_ADKTransactions.CallOpts, arg0)
}

// TransactionTrunk is a free data retrieval call binding the contract method 0xad69e6a0.
//
// Solidity: function transaction_trunk(bytes32 ) view returns(bytes32)
func (_ADKTransactions *ADKTransactionsCallerSession) TransactionTrunk(arg0 [32]byte) ([32]byte, error) {
	return _ADKTransactions.Contract.TransactionTrunk(&_ADKTransactions.CallOpts, arg0)
}

// TransactionhashByAddress is a free data retrieval call binding the contract method 0xb1406a64.
//
// Solidity: function transactionhash_by_address(bytes32 ) view returns(bytes32)
func (_ADKTransactions *ADKTransactionsCaller) TransactionhashByAddress(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "transactionhash_by_address", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TransactionhashByAddress is a free data retrieval call binding the contract method 0xb1406a64.
//
// Solidity: function transactionhash_by_address(bytes32 ) view returns(bytes32)
func (_ADKTransactions *ADKTransactionsSession) TransactionhashByAddress(arg0 [32]byte) ([32]byte, error) {
	return _ADKTransactions.Contract.TransactionhashByAddress(&_ADKTransactions.CallOpts, arg0)
}

// TransactionhashByAddress is a free data retrieval call binding the contract method 0xb1406a64.
//
// Solidity: function transactionhash_by_address(bytes32 ) view returns(bytes32)
func (_ADKTransactions *ADKTransactionsCallerSession) TransactionhashByAddress(arg0 [32]byte) ([32]byte, error) {
	return _ADKTransactions.Contract.TransactionhashByAddress(&_ADKTransactions.CallOpts, arg0)
}

// TransactionhashByAddressCount is a free data retrieval call binding the contract method 0xdcb711ff.
//
// Solidity: function transactionhash_by_address_count(bytes32 ) view returns(uint32)
func (_ADKTransactions *ADKTransactionsCaller) TransactionhashByAddressCount(opts *bind.CallOpts, arg0 [32]byte) (uint32, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "transactionhash_by_address_count", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TransactionhashByAddressCount is a free data retrieval call binding the contract method 0xdcb711ff.
//
// Solidity: function transactionhash_by_address_count(bytes32 ) view returns(uint32)
func (_ADKTransactions *ADKTransactionsSession) TransactionhashByAddressCount(arg0 [32]byte) (uint32, error) {
	return _ADKTransactions.Contract.TransactionhashByAddressCount(&_ADKTransactions.CallOpts, arg0)
}

// TransactionhashByAddressCount is a free data retrieval call binding the contract method 0xdcb711ff.
//
// Solidity: function transactionhash_by_address_count(bytes32 ) view returns(uint32)
func (_ADKTransactions *ADKTransactionsCallerSession) TransactionhashByAddressCount(arg0 [32]byte) (uint32, error) {
	return _ADKTransactions.Contract.TransactionhashByAddressCount(&_ADKTransactions.CallOpts, arg0)
}

// TransactionhashByBundle is a free data retrieval call binding the contract method 0x91da69af.
//
// Solidity: function transactionhash_by_bundle(bytes32 ) view returns(bytes32)
func (_ADKTransactions *ADKTransactionsCaller) TransactionhashByBundle(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "transactionhash_by_bundle", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TransactionhashByBundle is a free data retrieval call binding the contract method 0x91da69af.
//
// Solidity: function transactionhash_by_bundle(bytes32 ) view returns(bytes32)
func (_ADKTransactions *ADKTransactionsSession) TransactionhashByBundle(arg0 [32]byte) ([32]byte, error) {
	return _ADKTransactions.Contract.TransactionhashByBundle(&_ADKTransactions.CallOpts, arg0)
}

// TransactionhashByBundle is a free data retrieval call binding the contract method 0x91da69af.
//
// Solidity: function transactionhash_by_bundle(bytes32 ) view returns(bytes32)
func (_ADKTransactions *ADKTransactionsCallerSession) TransactionhashByBundle(arg0 [32]byte) ([32]byte, error) {
	return _ADKTransactions.Contract.TransactionhashByBundle(&_ADKTransactions.CallOpts, arg0)
}

// TransactionhashByBundleCount is a free data retrieval call binding the contract method 0xfc74fa69.
//
// Solidity: function transactionhash_by_bundle_count(bytes32 ) view returns(uint32)
func (_ADKTransactions *ADKTransactionsCaller) TransactionhashByBundleCount(opts *bind.CallOpts, arg0 [32]byte) (uint32, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "transactionhash_by_bundle_count", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TransactionhashByBundleCount is a free data retrieval call binding the contract method 0xfc74fa69.
//
// Solidity: function transactionhash_by_bundle_count(bytes32 ) view returns(uint32)
func (_ADKTransactions *ADKTransactionsSession) TransactionhashByBundleCount(arg0 [32]byte) (uint32, error) {
	return _ADKTransactions.Contract.TransactionhashByBundleCount(&_ADKTransactions.CallOpts, arg0)
}

// TransactionhashByBundleCount is a free data retrieval call binding the contract method 0xfc74fa69.
//
// Solidity: function transactionhash_by_bundle_count(bytes32 ) view returns(uint32)
func (_ADKTransactions *ADKTransactionsCallerSession) TransactionhashByBundleCount(arg0 [32]byte) (uint32, error) {
	return _ADKTransactions.Contract.TransactionhashByBundleCount(&_ADKTransactions.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions(bytes32 ) view returns(string)
func (_ADKTransactions *ADKTransactionsCaller) Transactions(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "transactions", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions(bytes32 ) view returns(string)
func (_ADKTransactions *ADKTransactionsSession) Transactions(arg0 [32]byte) (string, error) {
	return _ADKTransactions.Contract.Transactions(&_ADKTransactions.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions(bytes32 ) view returns(string)
func (_ADKTransactions *ADKTransactionsCallerSession) Transactions(arg0 [32]byte) (string, error) {
	return _ADKTransactions.Contract.Transactions(&_ADKTransactions.CallOpts, arg0)
}

// TxCount is a free data retrieval call binding the contract method 0xaa21064f.
//
// Solidity: function tx_count() view returns(uint256)
func (_ADKTransactions *ADKTransactionsCaller) TxCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ADKTransactions.contract.Call(opts, &out, "tx_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TxCount is a free data retrieval call binding the contract method 0xaa21064f.
//
// Solidity: function tx_count() view returns(uint256)
func (_ADKTransactions *ADKTransactionsSession) TxCount() (*big.Int, error) {
	return _ADKTransactions.Contract.TxCount(&_ADKTransactions.CallOpts)
}

// TxCount is a free data retrieval call binding the contract method 0xaa21064f.
//
// Solidity: function tx_count() view returns(uint256)
func (_ADKTransactions *ADKTransactionsCallerSession) TxCount() (*big.Int, error) {
	return _ADKTransactions.Contract.TxCount(&_ADKTransactions.CallOpts)
}

// ADMLoadADKBalances is a paid mutator transaction binding the contract method 0x64180204.
//
// Solidity: function ADM_LoadADKBalances(string _address, int256 _value) returns()
func (_ADKTransactions *ADKTransactionsTransactor) ADMLoadADKBalances(opts *bind.TransactOpts, _address string, _value *big.Int) (*types.Transaction, error) {
	return _ADKTransactions.contract.Transact(opts, "ADM_LoadADKBalances", _address, _value)
}

// ADMLoadADKBalances is a paid mutator transaction binding the contract method 0x64180204.
//
// Solidity: function ADM_LoadADKBalances(string _address, int256 _value) returns()
func (_ADKTransactions *ADKTransactionsSession) ADMLoadADKBalances(_address string, _value *big.Int) (*types.Transaction, error) {
	return _ADKTransactions.Contract.ADMLoadADKBalances(&_ADKTransactions.TransactOpts, _address, _value)
}

// ADMLoadADKBalances is a paid mutator transaction binding the contract method 0x64180204.
//
// Solidity: function ADM_LoadADKBalances(string _address, int256 _value) returns()
func (_ADKTransactions *ADKTransactionsTransactorSession) ADMLoadADKBalances(_address string, _value *big.Int) (*types.Transaction, error) {
	return _ADKTransactions.Contract.ADMLoadADKBalances(&_ADKTransactions.TransactOpts, _address, _value)
}

// ADMLoadADKBalancesBulk is a paid mutator transaction binding the contract method 0x6a82351e.
//
// Solidity: function ADM_LoadADKBalancesBulk(string _addresses, int256 _value1, int256 _value2, int256 _value3, int256 _value4, int256 _value5, int256 _value6, int256 _value7, int256 _value8, int256 _value9, int256 _value10) returns()
func (_ADKTransactions *ADKTransactionsTransactor) ADMLoadADKBalancesBulk(opts *bind.TransactOpts, _addresses string, _value1 *big.Int, _value2 *big.Int, _value3 *big.Int, _value4 *big.Int, _value5 *big.Int, _value6 *big.Int, _value7 *big.Int, _value8 *big.Int, _value9 *big.Int, _value10 *big.Int) (*types.Transaction, error) {
	return _ADKTransactions.contract.Transact(opts, "ADM_LoadADKBalancesBulk", _addresses, _value1, _value2, _value3, _value4, _value5, _value6, _value7, _value8, _value9, _value10)
}

// ADMLoadADKBalancesBulk is a paid mutator transaction binding the contract method 0x6a82351e.
//
// Solidity: function ADM_LoadADKBalancesBulk(string _addresses, int256 _value1, int256 _value2, int256 _value3, int256 _value4, int256 _value5, int256 _value6, int256 _value7, int256 _value8, int256 _value9, int256 _value10) returns()
func (_ADKTransactions *ADKTransactionsSession) ADMLoadADKBalancesBulk(_addresses string, _value1 *big.Int, _value2 *big.Int, _value3 *big.Int, _value4 *big.Int, _value5 *big.Int, _value6 *big.Int, _value7 *big.Int, _value8 *big.Int, _value9 *big.Int, _value10 *big.Int) (*types.Transaction, error) {
	return _ADKTransactions.Contract.ADMLoadADKBalancesBulk(&_ADKTransactions.TransactOpts, _addresses, _value1, _value2, _value3, _value4, _value5, _value6, _value7, _value8, _value9, _value10)
}

// ADMLoadADKBalancesBulk is a paid mutator transaction binding the contract method 0x6a82351e.
//
// Solidity: function ADM_LoadADKBalancesBulk(string _addresses, int256 _value1, int256 _value2, int256 _value3, int256 _value4, int256 _value5, int256 _value6, int256 _value7, int256 _value8, int256 _value9, int256 _value10) returns()
func (_ADKTransactions *ADKTransactionsTransactorSession) ADMLoadADKBalancesBulk(_addresses string, _value1 *big.Int, _value2 *big.Int, _value3 *big.Int, _value4 *big.Int, _value5 *big.Int, _value6 *big.Int, _value7 *big.Int, _value8 *big.Int, _value9 *big.Int, _value10 *big.Int) (*types.Transaction, error) {
	return _ADKTransactions.Contract.ADMLoadADKBalancesBulk(&_ADKTransactions.TransactOpts, _addresses, _value1, _value2, _value3, _value4, _value5, _value6, _value7, _value8, _value9, _value10)
}

// ADMLoadTransactionsUnchecked is a paid mutator transaction binding the contract method 0x50e6b384.
//
// Solidity: function ADM_LoadTransactionsUnchecked(string _transaction) returns()
func (_ADKTransactions *ADKTransactionsTransactor) ADMLoadTransactionsUnchecked(opts *bind.TransactOpts, _transaction string) (*types.Transaction, error) {
	return _ADKTransactions.contract.Transact(opts, "ADM_LoadTransactionsUnchecked", _transaction)
}

// ADMLoadTransactionsUnchecked is a paid mutator transaction binding the contract method 0x50e6b384.
//
// Solidity: function ADM_LoadTransactionsUnchecked(string _transaction) returns()
func (_ADKTransactions *ADKTransactionsSession) ADMLoadTransactionsUnchecked(_transaction string) (*types.Transaction, error) {
	return _ADKTransactions.Contract.ADMLoadTransactionsUnchecked(&_ADKTransactions.TransactOpts, _transaction)
}

// ADMLoadTransactionsUnchecked is a paid mutator transaction binding the contract method 0x50e6b384.
//
// Solidity: function ADM_LoadTransactionsUnchecked(string _transaction) returns()
func (_ADKTransactions *ADKTransactionsTransactorSession) ADMLoadTransactionsUnchecked(_transaction string) (*types.Transaction, error) {
	return _ADKTransactions.Contract.ADMLoadTransactionsUnchecked(&_ADKTransactions.TransactOpts, _transaction)
}

// ADMSetADKTokenContract is a paid mutator transaction binding the contract method 0xab815b02.
//
// Solidity: function ADM_SetADKTokenContract(address _newContract) returns()
func (_ADKTransactions *ADKTransactionsTransactor) ADMSetADKTokenContract(opts *bind.TransactOpts, _newContract common.Address) (*types.Transaction, error) {
	return _ADKTransactions.contract.Transact(opts, "ADM_SetADKTokenContract", _newContract)
}

// ADMSetADKTokenContract is a paid mutator transaction binding the contract method 0xab815b02.
//
// Solidity: function ADM_SetADKTokenContract(address _newContract) returns()
func (_ADKTransactions *ADKTransactionsSession) ADMSetADKTokenContract(_newContract common.Address) (*types.Transaction, error) {
	return _ADKTransactions.Contract.ADMSetADKTokenContract(&_ADKTransactions.TransactOpts, _newContract)
}

// ADMSetADKTokenContract is a paid mutator transaction binding the contract method 0xab815b02.
//
// Solidity: function ADM_SetADKTokenContract(address _newContract) returns()
func (_ADKTransactions *ADKTransactionsTransactorSession) ADMSetADKTokenContract(_newContract common.Address) (*types.Transaction, error) {
	return _ADKTransactions.Contract.ADMSetADKTokenContract(&_ADKTransactions.TransactOpts, _newContract)
}

// ADMSetDifficulty is a paid mutator transaction binding the contract method 0xd45c0ebc.
//
// Solidity: function ADM_SetDifficulty(uint256 _ClientProofOfWorkRequirement) returns()
func (_ADKTransactions *ADKTransactionsTransactor) ADMSetDifficulty(opts *bind.TransactOpts, _ClientProofOfWorkRequirement *big.Int) (*types.Transaction, error) {
	return _ADKTransactions.contract.Transact(opts, "ADM_SetDifficulty", _ClientProofOfWorkRequirement)
}

// ADMSetDifficulty is a paid mutator transaction binding the contract method 0xd45c0ebc.
//
// Solidity: function ADM_SetDifficulty(uint256 _ClientProofOfWorkRequirement) returns()
func (_ADKTransactions *ADKTransactionsSession) ADMSetDifficulty(_ClientProofOfWorkRequirement *big.Int) (*types.Transaction, error) {
	return _ADKTransactions.Contract.ADMSetDifficulty(&_ADKTransactions.TransactOpts, _ClientProofOfWorkRequirement)
}

// ADMSetDifficulty is a paid mutator transaction binding the contract method 0xd45c0ebc.
//
// Solidity: function ADM_SetDifficulty(uint256 _ClientProofOfWorkRequirement) returns()
func (_ADKTransactions *ADKTransactionsTransactorSession) ADMSetDifficulty(_ClientProofOfWorkRequirement *big.Int) (*types.Transaction, error) {
	return _ADKTransactions.Contract.ADMSetDifficulty(&_ADKTransactions.TransactOpts, _ClientProofOfWorkRequirement)
}

// ADMSetGenesisAddress is a paid mutator transaction binding the contract method 0xd48a1ce5.
//
// Solidity: function ADM_SetGenesisAddress(address _genesisAddress) returns()
func (_ADKTransactions *ADKTransactionsTransactor) ADMSetGenesisAddress(opts *bind.TransactOpts, _genesisAddress common.Address) (*types.Transaction, error) {
	return _ADKTransactions.contract.Transact(opts, "ADM_SetGenesisAddress", _genesisAddress)
}

// ADMSetGenesisAddress is a paid mutator transaction binding the contract method 0xd48a1ce5.
//
// Solidity: function ADM_SetGenesisAddress(address _genesisAddress) returns()
func (_ADKTransactions *ADKTransactionsSession) ADMSetGenesisAddress(_genesisAddress common.Address) (*types.Transaction, error) {
	return _ADKTransactions.Contract.ADMSetGenesisAddress(&_ADKTransactions.TransactOpts, _genesisAddress)
}

// ADMSetGenesisAddress is a paid mutator transaction binding the contract method 0xd48a1ce5.
//
// Solidity: function ADM_SetGenesisAddress(address _genesisAddress) returns()
func (_ADKTransactions *ADKTransactionsTransactorSession) ADMSetGenesisAddress(_genesisAddress common.Address) (*types.Transaction, error) {
	return _ADKTransactions.Contract.ADMSetGenesisAddress(&_ADKTransactions.TransactOpts, _genesisAddress)
}

// ADMSetTip is a paid mutator transaction binding the contract method 0x018e3998.
//
// Solidity: function ADM_SetTip(bytes32 _newTip) returns()
func (_ADKTransactions *ADKTransactionsTransactor) ADMSetTip(opts *bind.TransactOpts, _newTip [32]byte) (*types.Transaction, error) {
	return _ADKTransactions.contract.Transact(opts, "ADM_SetTip", _newTip)
}

// ADMSetTip is a paid mutator transaction binding the contract method 0x018e3998.
//
// Solidity: function ADM_SetTip(bytes32 _newTip) returns()
func (_ADKTransactions *ADKTransactionsSession) ADMSetTip(_newTip [32]byte) (*types.Transaction, error) {
	return _ADKTransactions.Contract.ADMSetTip(&_ADKTransactions.TransactOpts, _newTip)
}

// ADMSetTip is a paid mutator transaction binding the contract method 0x018e3998.
//
// Solidity: function ADM_SetTip(bytes32 _newTip) returns()
func (_ADKTransactions *ADKTransactionsTransactorSession) ADMSetTip(_newTip [32]byte) (*types.Transaction, error) {
	return _ADKTransactions.Contract.ADMSetTip(&_ADKTransactions.TransactOpts, _newTip)
}

// PostTransactions is a paid mutator transaction binding the contract method 0x7fa5e242.
//
// Solidity: function PostTransactions(string transactiondata) returns(string)
func (_ADKTransactions *ADKTransactionsTransactor) PostTransactions(opts *bind.TransactOpts, transactiondata string) (*types.Transaction, error) {
	return _ADKTransactions.contract.Transact(opts, "PostTransactions", transactiondata)
}

// PostTransactions is a paid mutator transaction binding the contract method 0x7fa5e242.
//
// Solidity: function PostTransactions(string transactiondata) returns(string)
func (_ADKTransactions *ADKTransactionsSession) PostTransactions(transactiondata string) (*types.Transaction, error) {
	return _ADKTransactions.Contract.PostTransactions(&_ADKTransactions.TransactOpts, transactiondata)
}

// PostTransactions is a paid mutator transaction binding the contract method 0x7fa5e242.
//
// Solidity: function PostTransactions(string transactiondata) returns(string)
func (_ADKTransactions *ADKTransactionsTransactorSession) PostTransactions(transactiondata string) (*types.Transaction, error) {
	return _ADKTransactions.Contract.PostTransactions(&_ADKTransactions.TransactOpts, transactiondata)
}

// ADKTransactionsEventLogIntIterator is returned from FilterEventLogInt and is used to iterate over the raw logs and unpacked data for EventLogInt events raised by the ADKTransactions contract.
type ADKTransactionsEventLogIntIterator struct {
	Event *ADKTransactionsEventLogInt // Event containing the contract specifics and raw log

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
func (it *ADKTransactionsEventLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ADKTransactionsEventLogInt)
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
		it.Event = new(ADKTransactionsEventLogInt)
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
func (it *ADKTransactionsEventLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ADKTransactionsEventLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ADKTransactionsEventLogInt represents a EventLogInt event raised by the ADKTransactions contract.
type ADKTransactionsEventLogInt struct {
	Info    common.Hash
	Intdata *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEventLogInt is a free log retrieval operation binding the contract event 0x6308696fb14e1bcd3fcde7330f7607d979f745988c8b9ffaaa407c856e64b8df.
//
// Solidity: event EventLogInt(string indexed info, int256 intdata)
func (_ADKTransactions *ADKTransactionsFilterer) FilterEventLogInt(opts *bind.FilterOpts, info []string) (*ADKTransactionsEventLogIntIterator, error) {

	var infoRule []interface{}
	for _, infoItem := range info {
		infoRule = append(infoRule, infoItem)
	}

	logs, sub, err := _ADKTransactions.contract.FilterLogs(opts, "EventLogInt", infoRule)
	if err != nil {
		return nil, err
	}
	return &ADKTransactionsEventLogIntIterator{contract: _ADKTransactions.contract, event: "EventLogInt", logs: logs, sub: sub}, nil
}

// WatchEventLogInt is a free log subscription operation binding the contract event 0x6308696fb14e1bcd3fcde7330f7607d979f745988c8b9ffaaa407c856e64b8df.
//
// Solidity: event EventLogInt(string indexed info, int256 intdata)
func (_ADKTransactions *ADKTransactionsFilterer) WatchEventLogInt(opts *bind.WatchOpts, sink chan<- *ADKTransactionsEventLogInt, info []string) (event.Subscription, error) {

	var infoRule []interface{}
	for _, infoItem := range info {
		infoRule = append(infoRule, infoItem)
	}

	logs, sub, err := _ADKTransactions.contract.WatchLogs(opts, "EventLogInt", infoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ADKTransactionsEventLogInt)
				if err := _ADKTransactions.contract.UnpackLog(event, "EventLogInt", log); err != nil {
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

// ParseEventLogInt is a log parse operation binding the contract event 0x6308696fb14e1bcd3fcde7330f7607d979f745988c8b9ffaaa407c856e64b8df.
//
// Solidity: event EventLogInt(string indexed info, int256 intdata)
func (_ADKTransactions *ADKTransactionsFilterer) ParseEventLogInt(log types.Log) (*ADKTransactionsEventLogInt, error) {
	event := new(ADKTransactionsEventLogInt)
	if err := _ADKTransactions.contract.UnpackLog(event, "EventLogInt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ADKTransactionsEventLogStringIterator is returned from FilterEventLogString and is used to iterate over the raw logs and unpacked data for EventLogString events raised by the ADKTransactions contract.
type ADKTransactionsEventLogStringIterator struct {
	Event *ADKTransactionsEventLogString // Event containing the contract specifics and raw log

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
func (it *ADKTransactionsEventLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ADKTransactionsEventLogString)
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
		it.Event = new(ADKTransactionsEventLogString)
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
func (it *ADKTransactionsEventLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ADKTransactionsEventLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ADKTransactionsEventLogString represents a EventLogString event raised by the ADKTransactions contract.
type ADKTransactionsEventLogString struct {
	Info    common.Hash
	Strdata string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEventLogString is a free log retrieval operation binding the contract event 0xccd963281651b70f9dc4bf799196583260466d9300a1c2431c37e93f40faa616.
//
// Solidity: event EventLogString(string indexed info, string strdata)
func (_ADKTransactions *ADKTransactionsFilterer) FilterEventLogString(opts *bind.FilterOpts, info []string) (*ADKTransactionsEventLogStringIterator, error) {

	var infoRule []interface{}
	for _, infoItem := range info {
		infoRule = append(infoRule, infoItem)
	}

	logs, sub, err := _ADKTransactions.contract.FilterLogs(opts, "EventLogString", infoRule)
	if err != nil {
		return nil, err
	}
	return &ADKTransactionsEventLogStringIterator{contract: _ADKTransactions.contract, event: "EventLogString", logs: logs, sub: sub}, nil
}

// WatchEventLogString is a free log subscription operation binding the contract event 0xccd963281651b70f9dc4bf799196583260466d9300a1c2431c37e93f40faa616.
//
// Solidity: event EventLogString(string indexed info, string strdata)
func (_ADKTransactions *ADKTransactionsFilterer) WatchEventLogString(opts *bind.WatchOpts, sink chan<- *ADKTransactionsEventLogString, info []string) (event.Subscription, error) {

	var infoRule []interface{}
	for _, infoItem := range info {
		infoRule = append(infoRule, infoItem)
	}

	logs, sub, err := _ADKTransactions.contract.WatchLogs(opts, "EventLogString", infoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ADKTransactionsEventLogString)
				if err := _ADKTransactions.contract.UnpackLog(event, "EventLogString", log); err != nil {
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

// ParseEventLogString is a log parse operation binding the contract event 0xccd963281651b70f9dc4bf799196583260466d9300a1c2431c37e93f40faa616.
//
// Solidity: event EventLogString(string indexed info, string strdata)
func (_ADKTransactions *ADKTransactionsFilterer) ParseEventLogString(log types.Log) (*ADKTransactionsEventLogString, error) {
	event := new(ADKTransactionsEventLogString)
	if err := _ADKTransactions.contract.UnpackLog(event, "EventLogString", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ADKTransactionsEventLogUIntIterator is returned from FilterEventLogUInt and is used to iterate over the raw logs and unpacked data for EventLogUInt events raised by the ADKTransactions contract.
type ADKTransactionsEventLogUIntIterator struct {
	Event *ADKTransactionsEventLogUInt // Event containing the contract specifics and raw log

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
func (it *ADKTransactionsEventLogUIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ADKTransactionsEventLogUInt)
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
		it.Event = new(ADKTransactionsEventLogUInt)
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
func (it *ADKTransactionsEventLogUIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ADKTransactionsEventLogUIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ADKTransactionsEventLogUInt represents a EventLogUInt event raised by the ADKTransactions contract.
type ADKTransactionsEventLogUInt struct {
	Info     common.Hash
	Uintdata *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEventLogUInt is a free log retrieval operation binding the contract event 0xa579b94ccb4ee3bbc6ce74784fcad38eb022c0e5a7e4cefe1d869dae14f9a0c3.
//
// Solidity: event EventLogUInt(string indexed info, uint256 uintdata)
func (_ADKTransactions *ADKTransactionsFilterer) FilterEventLogUInt(opts *bind.FilterOpts, info []string) (*ADKTransactionsEventLogUIntIterator, error) {

	var infoRule []interface{}
	for _, infoItem := range info {
		infoRule = append(infoRule, infoItem)
	}

	logs, sub, err := _ADKTransactions.contract.FilterLogs(opts, "EventLogUInt", infoRule)
	if err != nil {
		return nil, err
	}
	return &ADKTransactionsEventLogUIntIterator{contract: _ADKTransactions.contract, event: "EventLogUInt", logs: logs, sub: sub}, nil
}

// WatchEventLogUInt is a free log subscription operation binding the contract event 0xa579b94ccb4ee3bbc6ce74784fcad38eb022c0e5a7e4cefe1d869dae14f9a0c3.
//
// Solidity: event EventLogUInt(string indexed info, uint256 uintdata)
func (_ADKTransactions *ADKTransactionsFilterer) WatchEventLogUInt(opts *bind.WatchOpts, sink chan<- *ADKTransactionsEventLogUInt, info []string) (event.Subscription, error) {

	var infoRule []interface{}
	for _, infoItem := range info {
		infoRule = append(infoRule, infoItem)
	}

	logs, sub, err := _ADKTransactions.contract.WatchLogs(opts, "EventLogUInt", infoRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ADKTransactionsEventLogUInt)
				if err := _ADKTransactions.contract.UnpackLog(event, "EventLogUInt", log); err != nil {
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

// ParseEventLogUInt is a log parse operation binding the contract event 0xa579b94ccb4ee3bbc6ce74784fcad38eb022c0e5a7e4cefe1d869dae14f9a0c3.
//
// Solidity: event EventLogUInt(string indexed info, uint256 uintdata)
func (_ADKTransactions *ADKTransactionsFilterer) ParseEventLogUInt(log types.Log) (*ADKTransactionsEventLogUInt, error) {
	event := new(ADKTransactionsEventLogUInt)
	if err := _ADKTransactions.contract.UnpackLog(event, "EventLogUInt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ADKTransactionsTransactionsByAddressIterator is returned from FilterTransactionsByAddress and is used to iterate over the raw logs and unpacked data for TransactionsByAddress events raised by the ADKTransactions contract.
type ADKTransactionsTransactionsByAddressIterator struct {
	Event *ADKTransactionsTransactionsByAddress // Event containing the contract specifics and raw log

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
func (it *ADKTransactionsTransactionsByAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ADKTransactionsTransactionsByAddress)
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
		it.Event = new(ADKTransactionsTransactionsByAddress)
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
func (it *ADKTransactionsTransactionsByAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ADKTransactionsTransactionsByAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ADKTransactionsTransactionsByAddress represents a TransactionsByAddress event raised by the ADKTransactions contract.
type ADKTransactionsTransactionsByAddress struct {
	AddressSHA3     [32]byte
	TransactionSHA3 [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransactionsByAddress is a free log retrieval operation binding the contract event 0x3a3fefec9be0b15f11bf30dc166021af2910538b9322c12e44963bbbe72f7d7b.
//
// Solidity: event transactions_by_address(bytes32 indexed addressSHA3, bytes32 transactionSHA3)
func (_ADKTransactions *ADKTransactionsFilterer) FilterTransactionsByAddress(opts *bind.FilterOpts, addressSHA3 [][32]byte) (*ADKTransactionsTransactionsByAddressIterator, error) {

	var addressSHA3Rule []interface{}
	for _, addressSHA3Item := range addressSHA3 {
		addressSHA3Rule = append(addressSHA3Rule, addressSHA3Item)
	}

	logs, sub, err := _ADKTransactions.contract.FilterLogs(opts, "transactions_by_address", addressSHA3Rule)
	if err != nil {
		return nil, err
	}
	return &ADKTransactionsTransactionsByAddressIterator{contract: _ADKTransactions.contract, event: "transactions_by_address", logs: logs, sub: sub}, nil
}

// WatchTransactionsByAddress is a free log subscription operation binding the contract event 0x3a3fefec9be0b15f11bf30dc166021af2910538b9322c12e44963bbbe72f7d7b.
//
// Solidity: event transactions_by_address(bytes32 indexed addressSHA3, bytes32 transactionSHA3)
func (_ADKTransactions *ADKTransactionsFilterer) WatchTransactionsByAddress(opts *bind.WatchOpts, sink chan<- *ADKTransactionsTransactionsByAddress, addressSHA3 [][32]byte) (event.Subscription, error) {

	var addressSHA3Rule []interface{}
	for _, addressSHA3Item := range addressSHA3 {
		addressSHA3Rule = append(addressSHA3Rule, addressSHA3Item)
	}

	logs, sub, err := _ADKTransactions.contract.WatchLogs(opts, "transactions_by_address", addressSHA3Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ADKTransactionsTransactionsByAddress)
				if err := _ADKTransactions.contract.UnpackLog(event, "transactions_by_address", log); err != nil {
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

// ParseTransactionsByAddress is a log parse operation binding the contract event 0x3a3fefec9be0b15f11bf30dc166021af2910538b9322c12e44963bbbe72f7d7b.
//
// Solidity: event transactions_by_address(bytes32 indexed addressSHA3, bytes32 transactionSHA3)
func (_ADKTransactions *ADKTransactionsFilterer) ParseTransactionsByAddress(log types.Log) (*ADKTransactionsTransactionsByAddress, error) {
	event := new(ADKTransactionsTransactionsByAddress)
	if err := _ADKTransactions.contract.UnpackLog(event, "transactions_by_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ADKTransactionsTransactionsByBundleIterator is returned from FilterTransactionsByBundle and is used to iterate over the raw logs and unpacked data for TransactionsByBundle events raised by the ADKTransactions contract.
type ADKTransactionsTransactionsByBundleIterator struct {
	Event *ADKTransactionsTransactionsByBundle // Event containing the contract specifics and raw log

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
func (it *ADKTransactionsTransactionsByBundleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ADKTransactionsTransactionsByBundle)
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
		it.Event = new(ADKTransactionsTransactionsByBundle)
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
func (it *ADKTransactionsTransactionsByBundleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ADKTransactionsTransactionsByBundleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ADKTransactionsTransactionsByBundle represents a TransactionsByBundle event raised by the ADKTransactions contract.
type ADKTransactionsTransactionsByBundle struct {
	BundleSHA3      [32]byte
	TransactionSHA3 [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransactionsByBundle is a free log retrieval operation binding the contract event 0xa551994b821690876343950f8de610eb0796baaded797266283c842d257dea9f.
//
// Solidity: event transactions_by_bundle(bytes32 indexed bundleSHA3, bytes32 transactionSHA3)
func (_ADKTransactions *ADKTransactionsFilterer) FilterTransactionsByBundle(opts *bind.FilterOpts, bundleSHA3 [][32]byte) (*ADKTransactionsTransactionsByBundleIterator, error) {

	var bundleSHA3Rule []interface{}
	for _, bundleSHA3Item := range bundleSHA3 {
		bundleSHA3Rule = append(bundleSHA3Rule, bundleSHA3Item)
	}

	logs, sub, err := _ADKTransactions.contract.FilterLogs(opts, "transactions_by_bundle", bundleSHA3Rule)
	if err != nil {
		return nil, err
	}
	return &ADKTransactionsTransactionsByBundleIterator{contract: _ADKTransactions.contract, event: "transactions_by_bundle", logs: logs, sub: sub}, nil
}

// WatchTransactionsByBundle is a free log subscription operation binding the contract event 0xa551994b821690876343950f8de610eb0796baaded797266283c842d257dea9f.
//
// Solidity: event transactions_by_bundle(bytes32 indexed bundleSHA3, bytes32 transactionSHA3)
func (_ADKTransactions *ADKTransactionsFilterer) WatchTransactionsByBundle(opts *bind.WatchOpts, sink chan<- *ADKTransactionsTransactionsByBundle, bundleSHA3 [][32]byte) (event.Subscription, error) {

	var bundleSHA3Rule []interface{}
	for _, bundleSHA3Item := range bundleSHA3 {
		bundleSHA3Rule = append(bundleSHA3Rule, bundleSHA3Item)
	}

	logs, sub, err := _ADKTransactions.contract.WatchLogs(opts, "transactions_by_bundle", bundleSHA3Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ADKTransactionsTransactionsByBundle)
				if err := _ADKTransactions.contract.UnpackLog(event, "transactions_by_bundle", log); err != nil {
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

// ParseTransactionsByBundle is a log parse operation binding the contract event 0xa551994b821690876343950f8de610eb0796baaded797266283c842d257dea9f.
//
// Solidity: event transactions_by_bundle(bytes32 indexed bundleSHA3, bytes32 transactionSHA3)
func (_ADKTransactions *ADKTransactionsFilterer) ParseTransactionsByBundle(log types.Log) (*ADKTransactionsTransactionsByBundle, error) {
	event := new(ADKTransactionsTransactionsByBundle)
	if err := _ADKTransactions.contract.UnpackLog(event, "transactions_by_bundle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AGSClaimMetaData contains all meta data concerning the AGSClaim contract.
var AGSClaimMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_genesis_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ADKTokenContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimedAmount\",\"type\":\"uint256\"}],\"name\":\"EventClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"addrAZ9\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimedAmount\",\"type\":\"uint256\"}],\"name\":\"EventClaimedAZ9\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"ADKTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_genesisAddress\",\"type\":\"address\"}],\"name\":\"ADM_SetGenesisAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_AZ9addr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_claimableAmount\",\"type\":\"uint256\"}],\"name\":\"ADM_setClaimableAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_addresses\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value3\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value4\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value5\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value6\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value7\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value8\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value9\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value10\",\"type\":\"uint256\"}],\"name\":\"ADM_setClaimableAmountBulk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ClientProofOfWorkRequirement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"CurlHashOP\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"InitialAGSAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"transactiondata\",\"type\":\"string\"}],\"name\":\"PostTransactions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adkgo_genesis_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"claimableAZ9\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bal\",\"type\":\"uint256\"}],\"name\":\"recoverAGS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a91fa7cb": "ADKTokenAddress()",
		"d48a1ce5": "ADM_SetGenesisAddress(address)",
		"f5702e4e": "ADM_setClaimableAmount(string,uint256)",
		"f1e13cfe": "ADM_setClaimableAmountBulk(string,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)",
		"2076e796": "ClientProofOfWorkRequirement()",
		"016d1c68": "CurlHashOP(string)",
		"f7509c52": "InitialAGSAmount()",
		"7fa5e242": "PostTransactions(string)",
		"c4c7e151": "adkgo_genesis_address()",
		"402914f5": "claimable(address)",
		"af7f893b": "claimableAZ9(string)",
		"a9c02d49": "recoverAGS(uint256)",
	},
	Bin: "0x60806040526006805460ff191690553480156200001b57600080fd5b5060405162001b6538038062001b658339810160408190526200003e9162000095565b600080546001600160a01b039283166001600160a01b03199182161790915560018054939092169216919091179055600f600255620000cd565b80516001600160a01b03811681146200009057600080fd5b919050565b60008060408385031215620000a957600080fd5b620000b48362000078565b9150620000c46020840162000078565b90509250929050565b611a8880620000dd6000396000f3fe6080604052600436106100a45760003560e01c8063af7f893b11610061578063af7f893b146101a5578063c4c7e151146101dd578063d48a1ce5146101fd578063f1e13cfe1461021d578063f5702e4e1461023d578063f7509c521461025d57005b8063016d1c68146100a65780632076e796146100dc578063402914f5146101005780637fa5e2421461012d578063a91fa7cb1461014d578063a9c02d4914610185575b005b3480156100b257600080fd5b506100c66100c13660046115de565b610273565b6040516100d39190611818565b60405180910390f35b3480156100e857600080fd5b506100f260025481565b6040519081526020016100d3565b34801561010c57600080fd5b506100f261011b36600461159d565b60036020526000908152604090205481565b34801561013957600080fd5b506100c66101483660046115de565b610456565b34801561015957600080fd5b5060005461016d906001600160a01b031681565b6040516001600160a01b0390911681526020016100d3565b34801561019157600080fd5b506100a46101a03660046116f4565b610e78565b3480156101b157600080fd5b506100f26101c03660046115de565b805160208183018101805160048252928201919093012091525481565b3480156101e957600080fd5b5060015461016d906001600160a01b031681565b34801561020957600080fd5b506100a461021836600461159d565b610ed3565b34801561022957600080fd5b506100a4610238366004611660565b610f1f565b34801561024957600080fd5b506100a461025836600461161b565b61114e565b34801561026957600080fd5b506100f260055481565b60405160609060009061028e90829081908690602001611758565b60408051601f19818403018152919052805160208201819020919250600160fd1b908390600a9081106102c3576102c3611a0e565b60200101906001600160f81b031916908160001a90535081516020830181902090600160fe1b908490600a9081106102fd576102fd611a0e565b60200101906001600160f81b031916908160001a905350825160208401206040805160518082526080820190925260009160208201818036833701905050905060005b60208110156103e35784816020811061035b5761035b611a0e565b1a60f81b82828151811061037157610371611a0e565b60200101906001600160f81b031916908160001a90535083816020811061039a5761039a611a0e565b1a60f81b826103aa836020611873565b815181106103ba576103ba611a0e565b60200101906001600160f81b031916908160001a905350806103db81611999565b915050610340565b5060005b601181101561044b5782816020811061040257610402611a0e565b1a60f81b82610412836040611873565b8151811061042257610422611a0e565b60200101906001600160f81b031916908160001a9053508061044381611999565b9150506103e7565b509695505050505050565b606081600081905060018151101561046d57600080fd5b600160005b82518110156104fe5782818151811061048d5761048d611a0e565b60209101015160f81c603914806104e3575060418382815181106104b3576104b3611a0e565b016020015160f81c108015906104e35750605a8382815181106104d8576104d8611a0e565b016020015160f81c11155b6104ec57600091505b806104f681611999565b915050610472565b50806105425760405162461bcd60e51b815260206004820152600e60248201526d494e56414c49442054525954455360901b60448201526064015b60405180910390fd5b60065460ff161561058a5760405162461bcd60e51b81526020600482015260126024820152717265656e7472792070726576656e7465642160701b6044820152606401610539565b6006805460ff19166001179055845185906105a890610a71906119ce565b1580156105b6575060008151115b6106025760405162461bcd60e51b815260206004820152601d60248201527f496e76616c6964207472616e73616374696f6e287329206c656e6774680000006044820152606401610539565b6000610a71825161061391906118aa565b90508061ffff166003146106695760405162461bcd60e51b815260206004820152601860248201527f6d75737420686176652033207472616e73616374696f6e7300000000000000006044820152606401610539565b60006106768260a26118be565b61ffff1667ffffffffffffffff81111561069257610692611a24565b6040519080825280601f01601f1916602001820160405280156106bc576020820181803683370190505b506040805160208101909152600080825291925090606081805b8661ffff168163ffffffff161015610cb557610740604051806101200160405280606081526020016060815260200160608152602001606081526020016060815260200160608152602001600063ffffffff16815260200160008019168152602001606081525090565b61074c82610a716118e8565b63ffffffff1660c08201819052610770908f9061076b81610a7161188b565b61129b565b6020820181905261078090610273565b61010082015260005b600360025461079891906118aa565b81101561083a5760006107ac82605061190b565b905082610100015181815181106107c5576107c5611a0e565b6020910101516001600160f81b031916603960f81b146108275760405162461bcd60e51b815260206004820152601d60248201527f5452414e53414354494f4e20504f57204e4f5420434f4d504c455445440000006044820152606401610539565b508061083281611999565b915050610789565b50610100810151805160209091012060e082015260005b60a28161ffff1610156108ff57898161ffff168360c0015161088b610876919061188b565b610880919061188b565b63ffffffff168151811061089657610896611a0e565b01602001516001600160f81b0319168861ffff83166108b68660a26118e8565b6108c0919061188b565b63ffffffff16815181106108d6576108d6611a0e565b60200101906001600160f81b031916908160001a905350806108f781611977565b915050610851565b50610913816020015161088b6108dc61129b565b815260c081015161095a908a9061092c9061092d61188b565b63ffffffff168360c0015161092d610944919061188b565b61094f90605161188b565b63ffffffff16611386565b60a082015263ffffffff8216610a00578060a0015195506109af6109818f6000601061129b565b6040518060400160405280601081526020016f21a620a4a6aa2920a729a0a1aa24a7a760811b815250611451565b6109fb5760405162461bcd60e51b815260206004820152601760248201527f6e6f74206120636c61696d207472616e73636174696f6e0000000000000000006044820152606401610539565b610a52565b610a0e868260a00151611451565b610a525760405162461bcd60e51b8152602060048201526015602482015274189d5b991b19481b9bdd0818dbdb9cda5cdd195b9d605a1b6044820152606401610539565b63ffffffff8216610ae557600054815160405163a205a8a960e01b81526001600160a01b039092169163a205a8a991610a8d91600401611818565b60206040518083038186803b158015610aa557600080fd5b505afa158015610ab9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610add91906115c1565b815190955093505b8163ffffffff1660011415610ca257600054815160405163a205a8a960e01b81526001600160a01b039092169163a205a8a991610b2491600401611818565b60206040518083038186803b158015610b3c57600080fd5b505afa158015610b50573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b7491906115c1565b9250610bb58e8260c00151610a71610b8c919061188b565b610b989061088b61188b565b60c0840151610ba990610a7161188b565b61076b906108dc61188b565b604082018190528151610bc791611451565b610c1d5760405162461bcd60e51b815260206004820152602160248201527f326e64207369676e61747572652068617320696e76616c6964206164647265736044820152607360f81b6064820152608401610539565b610c2c8e610a716112fc61129b565b6060820152610c408e6114e2611d6d61129b565b6080820181905281516060830151610c6292610c5c91906114aa565b886114d6565b610ca25760405162461bcd60e51b8152602060048201526011602482015270494e56414c4944205349474e415455524560781b6044820152606401610539565b5080610cad816119b4565b9150506106d6565b506000610cc186610273565b9050610ccd8582611451565b610d195760405162461bcd60e51b815260206004820152601960248201527f43414c43554c415445442042554e444c452044494646455253000000000000006044820152606401610539565b6001600160a01b0382166000908152600360205260408082208054908390559051909190600490610d4b90879061170d565b9081526040519081900360200190205580610d9b5760405162461bcd60e51b815260206004820152601060248201526f4e6f7468696e6720746f20636c61696d60801b6044820152606401610539565b6040516001600160a01b0386169082156108fc029083906000818181858888f19350505050158015610dd1573d6000803e3d6000fd5b50826001600160a01b03167f92c59d33103d56fd232bed949f5c3a793e67ab8b3b6de334a349fba4d901d4bc82604051610e0d91815260200190565b60405180910390a283604051610e23919061170d565b604051908190038120828252907ff09d7862a749844d0be7c0c15c0bf70d2dac7127deb9486545be11987affb0409060200160405180910390a2506006805460ff191690559c9b505050505050505050505050565b6001546001600160a01b03163314610ea25760405162461bcd60e51b81526004016105399061184b565b604051339082156108fc029083906000818181858888f19350505050158015610ecf573d6000803e3d6000fd5b5050565b6001546001600160a01b03163314610efd5760405162461bcd60e51b81526004016105399061184b565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b03163314610f495760405162461bcd60e51b81526004016105399061184b565b60008b5161032a14610fd15760405162461bcd60e51b815260206004820152604560248201527f537472696e67206d75737420636f6e7461696e2031302061646472657373657360448201527f20776974686f757420636865636b73756d20666f722062756c6b2070726f63656064820152647373696e6760d81b608482015260a401610539565b610fea610fe48d8361076b81605161188b565b8c61114e565b610ff560518261188b565b905061101061100a8d8361076b81605161188b565b8b61114e565b61101b60518261188b565b90506110366110308d8361076b81605161188b565b8a61114e565b61104160518261188b565b905061105c6110568d8361076b81605161188b565b8961114e565b61106760518261188b565b905061108261107c8d8361076b81605161188b565b8861114e565b61108d60518261188b565b90506110a86110a28d8361076b81605161188b565b8761114e565b6110b360518261188b565b90506110ce6110c88d8361076b81605161188b565b8661114e565b6110d960518261188b565b90506110f46110ee8d8361076b81605161188b565b8561114e565b6110ff60518261188b565b905061111a6111148d8361076b81605161188b565b8461114e565b61112560518261188b565b905061114061113a8d8361076b81605161188b565b8361114e565b505050505050505050505050565b6001546001600160a01b031633146111785760405162461bcd60e51b81526004016105399061184b565b81516051146111bb5760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b6044820152606401610539565b6000805460405163a205a8a960e01b81526001600160a01b039091169063a205a8a9906111ec908690600401611818565b60206040518083038186803b15801561120457600080fd5b505afa158015611218573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061123c91906115c1565b6001600160a01b0381166000908152600360205260409081902084905551909150829060049061126d90869061170d565b90815260200160405180910390208190555081600560008282546112919190611873565b9091555050505050565b60608360006112aa8585611922565b63ffffffff1667ffffffffffffffff8111156112c8576112c8611a24565b6040519080825280601f01601f1916602001820160405280156112f2576020820181803683370190505b509050845b8463ffffffff168163ffffffff16101561137c57828163ffffffff168151811061132357611323611a0e565b01602001516001600160f81b0319168261133d8884611922565b63ffffffff168151811061135357611353611a0e565b60200101906001600160f81b031916908160001a90535080611374816119b4565b9150506112f7565b5095945050505050565b60606000611394848461190b565b67ffffffffffffffff8111156113ac576113ac611a24565b6040519080825280601f01601f1916602001820160405280156113d6576020820181803683370190505b509050835b83811015611448578581815181106113f5576113f5611a0e565b01602001516001600160f81b0319168261140f878461190b565b8151811061141f5761141f611a0e565b60200101906001600160f81b031916908160001a9053508061144081611999565b9150506113db565b50949350505050565b600081604051602001611464919061170d565b604051602081830303815290604052805190602001208360405160200161148b919061170d565b6040516020818303038152906040528051906020012014905092915050565b606082826040516020016114bf929190611729565b604051602081830303815290604052905092915050565b6000808484846040516020016114ee939291906117ae565b60408051808303601f1901815291905280516020909101201595945050505050565b600082601f83011261152157600080fd5b813567ffffffffffffffff8082111561153c5761153c611a24565b604051601f8301601f19908116603f0116810190828211818310171561156457611564611a24565b8160405283815286602085880101111561157d57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000602082840312156115af57600080fd5b81356115ba81611a3a565b9392505050565b6000602082840312156115d357600080fd5b81516115ba81611a3a565b6000602082840312156115f057600080fd5b813567ffffffffffffffff81111561160757600080fd5b61161384828501611510565b949350505050565b6000806040838503121561162e57600080fd5b823567ffffffffffffffff81111561164557600080fd5b61165185828601611510565b95602094909401359450505050565b60008060008060008060008060008060006101608c8e03121561168257600080fd5b8b3567ffffffffffffffff81111561169957600080fd5b6116a58e828f01611510565b9e60208e01359e5060408e01359d60608101359d5060808101359c5060a08101359b5060c08101359a5060e0810135995061010081013598506101208101359750610140013595509350505050565b60006020828403121561170657600080fd5b5035919050565b6000825161171f818460208701611947565b9190910192915050565b6000835161173b818460208801611947565b83519083019061174f818360208801611947565b01949350505050565b657b4355524c7d60d01b81526309082a6960e31b60068201526001600160f81b0319848116600a8301528316600b820152815160009061179f81600c850160208701611947565b91909101600c01949350505050565b657b4355524c7d60d01b81526556414c53494760d01b6006820152600084516117de81600c850160208901611947565b8451908301906117f581600c840160208901611947565b845191019061180b81600c840160208801611947565b01600c0195945050505050565b6020815260008251806020840152611837816040850160208701611947565b601f01601f19169190910160400192915050565b6020808252600e908201526d1393d5081055551213d49256915160921b604082015260600190565b60008219821115611886576118866119e2565b500190565b600063ffffffff80831681851680830382111561174f5761174f6119e2565b6000826118b9576118b96119f8565b500490565b600061ffff808316818516818304811182151516156118df576118df6119e2565b02949350505050565b600063ffffffff808316818516818304811182151516156118df576118df6119e2565b60008282101561191d5761191d6119e2565b500390565b600063ffffffff8381169083168181101561193f5761193f6119e2565b039392505050565b60005b8381101561196257818101518382015260200161194a565b83811115611971576000848401525b50505050565b600061ffff8083168181141561198f5761198f6119e2565b6001019392505050565b60006000198214156119ad576119ad6119e2565b5060010190565b600063ffffffff8083168181141561198f5761198f6119e2565b6000826119dd576119dd6119f8565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114611a4f57600080fd5b5056fea264697066735822122034957912cb8ff2de5dcaa253add769abcec4f89ce9d9a1a615846a910556130364736f6c63430008070033",
}

// AGSClaimABI is the input ABI used to generate the binding from.
// Deprecated: Use AGSClaimMetaData.ABI instead.
var AGSClaimABI = AGSClaimMetaData.ABI

// Deprecated: Use AGSClaimMetaData.Sigs instead.
// AGSClaimFuncSigs maps the 4-byte function signature to its string representation.
var AGSClaimFuncSigs = AGSClaimMetaData.Sigs

// AGSClaimBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AGSClaimMetaData.Bin instead.
var AGSClaimBin = AGSClaimMetaData.Bin

// DeployAGSClaim deploys a new Ethereum contract, binding an instance of AGSClaim to it.
func DeployAGSClaim(auth *bind.TransactOpts, backend bind.ContractBackend, _genesis_account common.Address, _ADKTokenContract common.Address) (common.Address, *types.Transaction, *AGSClaim, error) {
	parsed, err := AGSClaimMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AGSClaimBin), backend, _genesis_account, _ADKTokenContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AGSClaim{AGSClaimCaller: AGSClaimCaller{contract: contract}, AGSClaimTransactor: AGSClaimTransactor{contract: contract}, AGSClaimFilterer: AGSClaimFilterer{contract: contract}}, nil
}

// AGSClaim is an auto generated Go binding around an Ethereum contract.
type AGSClaim struct {
	AGSClaimCaller     // Read-only binding to the contract
	AGSClaimTransactor // Write-only binding to the contract
	AGSClaimFilterer   // Log filterer for contract events
}

// AGSClaimCaller is an auto generated read-only Go binding around an Ethereum contract.
type AGSClaimCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AGSClaimTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AGSClaimTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AGSClaimFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AGSClaimFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AGSClaimSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AGSClaimSession struct {
	Contract     *AGSClaim         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AGSClaimCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AGSClaimCallerSession struct {
	Contract *AGSClaimCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AGSClaimTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AGSClaimTransactorSession struct {
	Contract     *AGSClaimTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AGSClaimRaw is an auto generated low-level Go binding around an Ethereum contract.
type AGSClaimRaw struct {
	Contract *AGSClaim // Generic contract binding to access the raw methods on
}

// AGSClaimCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AGSClaimCallerRaw struct {
	Contract *AGSClaimCaller // Generic read-only contract binding to access the raw methods on
}

// AGSClaimTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AGSClaimTransactorRaw struct {
	Contract *AGSClaimTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAGSClaim creates a new instance of AGSClaim, bound to a specific deployed contract.
func NewAGSClaim(address common.Address, backend bind.ContractBackend) (*AGSClaim, error) {
	contract, err := bindAGSClaim(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AGSClaim{AGSClaimCaller: AGSClaimCaller{contract: contract}, AGSClaimTransactor: AGSClaimTransactor{contract: contract}, AGSClaimFilterer: AGSClaimFilterer{contract: contract}}, nil
}

// NewAGSClaimCaller creates a new read-only instance of AGSClaim, bound to a specific deployed contract.
func NewAGSClaimCaller(address common.Address, caller bind.ContractCaller) (*AGSClaimCaller, error) {
	contract, err := bindAGSClaim(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AGSClaimCaller{contract: contract}, nil
}

// NewAGSClaimTransactor creates a new write-only instance of AGSClaim, bound to a specific deployed contract.
func NewAGSClaimTransactor(address common.Address, transactor bind.ContractTransactor) (*AGSClaimTransactor, error) {
	contract, err := bindAGSClaim(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AGSClaimTransactor{contract: contract}, nil
}

// NewAGSClaimFilterer creates a new log filterer instance of AGSClaim, bound to a specific deployed contract.
func NewAGSClaimFilterer(address common.Address, filterer bind.ContractFilterer) (*AGSClaimFilterer, error) {
	contract, err := bindAGSClaim(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AGSClaimFilterer{contract: contract}, nil
}

// bindAGSClaim binds a generic wrapper to an already deployed contract.
func bindAGSClaim(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AGSClaimABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AGSClaim *AGSClaimRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AGSClaim.Contract.AGSClaimCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AGSClaim *AGSClaimRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AGSClaim.Contract.AGSClaimTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AGSClaim *AGSClaimRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AGSClaim.Contract.AGSClaimTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AGSClaim *AGSClaimCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AGSClaim.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AGSClaim *AGSClaimTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AGSClaim.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AGSClaim *AGSClaimTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AGSClaim.Contract.contract.Transact(opts, method, params...)
}

// ADKTokenAddress is a free data retrieval call binding the contract method 0xa91fa7cb.
//
// Solidity: function ADKTokenAddress() view returns(address)
func (_AGSClaim *AGSClaimCaller) ADKTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AGSClaim.contract.Call(opts, &out, "ADKTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADKTokenAddress is a free data retrieval call binding the contract method 0xa91fa7cb.
//
// Solidity: function ADKTokenAddress() view returns(address)
func (_AGSClaim *AGSClaimSession) ADKTokenAddress() (common.Address, error) {
	return _AGSClaim.Contract.ADKTokenAddress(&_AGSClaim.CallOpts)
}

// ADKTokenAddress is a free data retrieval call binding the contract method 0xa91fa7cb.
//
// Solidity: function ADKTokenAddress() view returns(address)
func (_AGSClaim *AGSClaimCallerSession) ADKTokenAddress() (common.Address, error) {
	return _AGSClaim.Contract.ADKTokenAddress(&_AGSClaim.CallOpts)
}

// ClientProofOfWorkRequirement is a free data retrieval call binding the contract method 0x2076e796.
//
// Solidity: function ClientProofOfWorkRequirement() view returns(uint256)
func (_AGSClaim *AGSClaimCaller) ClientProofOfWorkRequirement(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AGSClaim.contract.Call(opts, &out, "ClientProofOfWorkRequirement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClientProofOfWorkRequirement is a free data retrieval call binding the contract method 0x2076e796.
//
// Solidity: function ClientProofOfWorkRequirement() view returns(uint256)
func (_AGSClaim *AGSClaimSession) ClientProofOfWorkRequirement() (*big.Int, error) {
	return _AGSClaim.Contract.ClientProofOfWorkRequirement(&_AGSClaim.CallOpts)
}

// ClientProofOfWorkRequirement is a free data retrieval call binding the contract method 0x2076e796.
//
// Solidity: function ClientProofOfWorkRequirement() view returns(uint256)
func (_AGSClaim *AGSClaimCallerSession) ClientProofOfWorkRequirement() (*big.Int, error) {
	return _AGSClaim.Contract.ClientProofOfWorkRequirement(&_AGSClaim.CallOpts)
}

// CurlHashOP is a free data retrieval call binding the contract method 0x016d1c68.
//
// Solidity: function CurlHashOP(string str) pure returns(string)
func (_AGSClaim *AGSClaimCaller) CurlHashOP(opts *bind.CallOpts, str string) (string, error) {
	var out []interface{}
	err := _AGSClaim.contract.Call(opts, &out, "CurlHashOP", str)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CurlHashOP is a free data retrieval call binding the contract method 0x016d1c68.
//
// Solidity: function CurlHashOP(string str) pure returns(string)
func (_AGSClaim *AGSClaimSession) CurlHashOP(str string) (string, error) {
	return _AGSClaim.Contract.CurlHashOP(&_AGSClaim.CallOpts, str)
}

// CurlHashOP is a free data retrieval call binding the contract method 0x016d1c68.
//
// Solidity: function CurlHashOP(string str) pure returns(string)
func (_AGSClaim *AGSClaimCallerSession) CurlHashOP(str string) (string, error) {
	return _AGSClaim.Contract.CurlHashOP(&_AGSClaim.CallOpts, str)
}

// InitialAGSAmount is a free data retrieval call binding the contract method 0xf7509c52.
//
// Solidity: function InitialAGSAmount() view returns(uint256)
func (_AGSClaim *AGSClaimCaller) InitialAGSAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AGSClaim.contract.Call(opts, &out, "InitialAGSAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialAGSAmount is a free data retrieval call binding the contract method 0xf7509c52.
//
// Solidity: function InitialAGSAmount() view returns(uint256)
func (_AGSClaim *AGSClaimSession) InitialAGSAmount() (*big.Int, error) {
	return _AGSClaim.Contract.InitialAGSAmount(&_AGSClaim.CallOpts)
}

// InitialAGSAmount is a free data retrieval call binding the contract method 0xf7509c52.
//
// Solidity: function InitialAGSAmount() view returns(uint256)
func (_AGSClaim *AGSClaimCallerSession) InitialAGSAmount() (*big.Int, error) {
	return _AGSClaim.Contract.InitialAGSAmount(&_AGSClaim.CallOpts)
}

// AdkgoGenesisAddress is a free data retrieval call binding the contract method 0xc4c7e151.
//
// Solidity: function adkgo_genesis_address() view returns(address)
func (_AGSClaim *AGSClaimCaller) AdkgoGenesisAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AGSClaim.contract.Call(opts, &out, "adkgo_genesis_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdkgoGenesisAddress is a free data retrieval call binding the contract method 0xc4c7e151.
//
// Solidity: function adkgo_genesis_address() view returns(address)
func (_AGSClaim *AGSClaimSession) AdkgoGenesisAddress() (common.Address, error) {
	return _AGSClaim.Contract.AdkgoGenesisAddress(&_AGSClaim.CallOpts)
}

// AdkgoGenesisAddress is a free data retrieval call binding the contract method 0xc4c7e151.
//
// Solidity: function adkgo_genesis_address() view returns(address)
func (_AGSClaim *AGSClaimCallerSession) AdkgoGenesisAddress() (common.Address, error) {
	return _AGSClaim.Contract.AdkgoGenesisAddress(&_AGSClaim.CallOpts)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_AGSClaim *AGSClaimCaller) Claimable(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AGSClaim.contract.Call(opts, &out, "claimable", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_AGSClaim *AGSClaimSession) Claimable(arg0 common.Address) (*big.Int, error) {
	return _AGSClaim.Contract.Claimable(&_AGSClaim.CallOpts, arg0)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_AGSClaim *AGSClaimCallerSession) Claimable(arg0 common.Address) (*big.Int, error) {
	return _AGSClaim.Contract.Claimable(&_AGSClaim.CallOpts, arg0)
}

// ClaimableAZ9 is a free data retrieval call binding the contract method 0xaf7f893b.
//
// Solidity: function claimableAZ9(string ) view returns(uint256)
func (_AGSClaim *AGSClaimCaller) ClaimableAZ9(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _AGSClaim.contract.Call(opts, &out, "claimableAZ9", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableAZ9 is a free data retrieval call binding the contract method 0xaf7f893b.
//
// Solidity: function claimableAZ9(string ) view returns(uint256)
func (_AGSClaim *AGSClaimSession) ClaimableAZ9(arg0 string) (*big.Int, error) {
	return _AGSClaim.Contract.ClaimableAZ9(&_AGSClaim.CallOpts, arg0)
}

// ClaimableAZ9 is a free data retrieval call binding the contract method 0xaf7f893b.
//
// Solidity: function claimableAZ9(string ) view returns(uint256)
func (_AGSClaim *AGSClaimCallerSession) ClaimableAZ9(arg0 string) (*big.Int, error) {
	return _AGSClaim.Contract.ClaimableAZ9(&_AGSClaim.CallOpts, arg0)
}

// ADMSetGenesisAddress is a paid mutator transaction binding the contract method 0xd48a1ce5.
//
// Solidity: function ADM_SetGenesisAddress(address _genesisAddress) returns()
func (_AGSClaim *AGSClaimTransactor) ADMSetGenesisAddress(opts *bind.TransactOpts, _genesisAddress common.Address) (*types.Transaction, error) {
	return _AGSClaim.contract.Transact(opts, "ADM_SetGenesisAddress", _genesisAddress)
}

// ADMSetGenesisAddress is a paid mutator transaction binding the contract method 0xd48a1ce5.
//
// Solidity: function ADM_SetGenesisAddress(address _genesisAddress) returns()
func (_AGSClaim *AGSClaimSession) ADMSetGenesisAddress(_genesisAddress common.Address) (*types.Transaction, error) {
	return _AGSClaim.Contract.ADMSetGenesisAddress(&_AGSClaim.TransactOpts, _genesisAddress)
}

// ADMSetGenesisAddress is a paid mutator transaction binding the contract method 0xd48a1ce5.
//
// Solidity: function ADM_SetGenesisAddress(address _genesisAddress) returns()
func (_AGSClaim *AGSClaimTransactorSession) ADMSetGenesisAddress(_genesisAddress common.Address) (*types.Transaction, error) {
	return _AGSClaim.Contract.ADMSetGenesisAddress(&_AGSClaim.TransactOpts, _genesisAddress)
}

// ADMSetClaimableAmount is a paid mutator transaction binding the contract method 0xf5702e4e.
//
// Solidity: function ADM_setClaimableAmount(string _AZ9addr, uint256 _claimableAmount) returns()
func (_AGSClaim *AGSClaimTransactor) ADMSetClaimableAmount(opts *bind.TransactOpts, _AZ9addr string, _claimableAmount *big.Int) (*types.Transaction, error) {
	return _AGSClaim.contract.Transact(opts, "ADM_setClaimableAmount", _AZ9addr, _claimableAmount)
}

// ADMSetClaimableAmount is a paid mutator transaction binding the contract method 0xf5702e4e.
//
// Solidity: function ADM_setClaimableAmount(string _AZ9addr, uint256 _claimableAmount) returns()
func (_AGSClaim *AGSClaimSession) ADMSetClaimableAmount(_AZ9addr string, _claimableAmount *big.Int) (*types.Transaction, error) {
	return _AGSClaim.Contract.ADMSetClaimableAmount(&_AGSClaim.TransactOpts, _AZ9addr, _claimableAmount)
}

// ADMSetClaimableAmount is a paid mutator transaction binding the contract method 0xf5702e4e.
//
// Solidity: function ADM_setClaimableAmount(string _AZ9addr, uint256 _claimableAmount) returns()
func (_AGSClaim *AGSClaimTransactorSession) ADMSetClaimableAmount(_AZ9addr string, _claimableAmount *big.Int) (*types.Transaction, error) {
	return _AGSClaim.Contract.ADMSetClaimableAmount(&_AGSClaim.TransactOpts, _AZ9addr, _claimableAmount)
}

// ADMSetClaimableAmountBulk is a paid mutator transaction binding the contract method 0xf1e13cfe.
//
// Solidity: function ADM_setClaimableAmountBulk(string _addresses, uint256 _value1, uint256 _value2, uint256 _value3, uint256 _value4, uint256 _value5, uint256 _value6, uint256 _value7, uint256 _value8, uint256 _value9, uint256 _value10) returns()
func (_AGSClaim *AGSClaimTransactor) ADMSetClaimableAmountBulk(opts *bind.TransactOpts, _addresses string, _value1 *big.Int, _value2 *big.Int, _value3 *big.Int, _value4 *big.Int, _value5 *big.Int, _value6 *big.Int, _value7 *big.Int, _value8 *big.Int, _value9 *big.Int, _value10 *big.Int) (*types.Transaction, error) {
	return _AGSClaim.contract.Transact(opts, "ADM_setClaimableAmountBulk", _addresses, _value1, _value2, _value3, _value4, _value5, _value6, _value7, _value8, _value9, _value10)
}

// ADMSetClaimableAmountBulk is a paid mutator transaction binding the contract method 0xf1e13cfe.
//
// Solidity: function ADM_setClaimableAmountBulk(string _addresses, uint256 _value1, uint256 _value2, uint256 _value3, uint256 _value4, uint256 _value5, uint256 _value6, uint256 _value7, uint256 _value8, uint256 _value9, uint256 _value10) returns()
func (_AGSClaim *AGSClaimSession) ADMSetClaimableAmountBulk(_addresses string, _value1 *big.Int, _value2 *big.Int, _value3 *big.Int, _value4 *big.Int, _value5 *big.Int, _value6 *big.Int, _value7 *big.Int, _value8 *big.Int, _value9 *big.Int, _value10 *big.Int) (*types.Transaction, error) {
	return _AGSClaim.Contract.ADMSetClaimableAmountBulk(&_AGSClaim.TransactOpts, _addresses, _value1, _value2, _value3, _value4, _value5, _value6, _value7, _value8, _value9, _value10)
}

// ADMSetClaimableAmountBulk is a paid mutator transaction binding the contract method 0xf1e13cfe.
//
// Solidity: function ADM_setClaimableAmountBulk(string _addresses, uint256 _value1, uint256 _value2, uint256 _value3, uint256 _value4, uint256 _value5, uint256 _value6, uint256 _value7, uint256 _value8, uint256 _value9, uint256 _value10) returns()
func (_AGSClaim *AGSClaimTransactorSession) ADMSetClaimableAmountBulk(_addresses string, _value1 *big.Int, _value2 *big.Int, _value3 *big.Int, _value4 *big.Int, _value5 *big.Int, _value6 *big.Int, _value7 *big.Int, _value8 *big.Int, _value9 *big.Int, _value10 *big.Int) (*types.Transaction, error) {
	return _AGSClaim.Contract.ADMSetClaimableAmountBulk(&_AGSClaim.TransactOpts, _addresses, _value1, _value2, _value3, _value4, _value5, _value6, _value7, _value8, _value9, _value10)
}

// PostTransactions is a paid mutator transaction binding the contract method 0x7fa5e242.
//
// Solidity: function PostTransactions(string transactiondata) returns(string)
func (_AGSClaim *AGSClaimTransactor) PostTransactions(opts *bind.TransactOpts, transactiondata string) (*types.Transaction, error) {
	return _AGSClaim.contract.Transact(opts, "PostTransactions", transactiondata)
}

// PostTransactions is a paid mutator transaction binding the contract method 0x7fa5e242.
//
// Solidity: function PostTransactions(string transactiondata) returns(string)
func (_AGSClaim *AGSClaimSession) PostTransactions(transactiondata string) (*types.Transaction, error) {
	return _AGSClaim.Contract.PostTransactions(&_AGSClaim.TransactOpts, transactiondata)
}

// PostTransactions is a paid mutator transaction binding the contract method 0x7fa5e242.
//
// Solidity: function PostTransactions(string transactiondata) returns(string)
func (_AGSClaim *AGSClaimTransactorSession) PostTransactions(transactiondata string) (*types.Transaction, error) {
	return _AGSClaim.Contract.PostTransactions(&_AGSClaim.TransactOpts, transactiondata)
}

// RecoverAGS is a paid mutator transaction binding the contract method 0xa9c02d49.
//
// Solidity: function recoverAGS(uint256 _bal) returns()
func (_AGSClaim *AGSClaimTransactor) RecoverAGS(opts *bind.TransactOpts, _bal *big.Int) (*types.Transaction, error) {
	return _AGSClaim.contract.Transact(opts, "recoverAGS", _bal)
}

// RecoverAGS is a paid mutator transaction binding the contract method 0xa9c02d49.
//
// Solidity: function recoverAGS(uint256 _bal) returns()
func (_AGSClaim *AGSClaimSession) RecoverAGS(_bal *big.Int) (*types.Transaction, error) {
	return _AGSClaim.Contract.RecoverAGS(&_AGSClaim.TransactOpts, _bal)
}

// RecoverAGS is a paid mutator transaction binding the contract method 0xa9c02d49.
//
// Solidity: function recoverAGS(uint256 _bal) returns()
func (_AGSClaim *AGSClaimTransactorSession) RecoverAGS(_bal *big.Int) (*types.Transaction, error) {
	return _AGSClaim.Contract.RecoverAGS(&_AGSClaim.TransactOpts, _bal)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AGSClaim *AGSClaimTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _AGSClaim.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AGSClaim *AGSClaimSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AGSClaim.Contract.Fallback(&_AGSClaim.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AGSClaim *AGSClaimTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AGSClaim.Contract.Fallback(&_AGSClaim.TransactOpts, calldata)
}

// AGSClaimEventClaimedIterator is returned from FilterEventClaimed and is used to iterate over the raw logs and unpacked data for EventClaimed events raised by the AGSClaim contract.
type AGSClaimEventClaimedIterator struct {
	Event *AGSClaimEventClaimed // Event containing the contract specifics and raw log

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
func (it *AGSClaimEventClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AGSClaimEventClaimed)
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
		it.Event = new(AGSClaimEventClaimed)
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
func (it *AGSClaimEventClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AGSClaimEventClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AGSClaimEventClaimed represents a EventClaimed event raised by the AGSClaim contract.
type AGSClaimEventClaimed struct {
	Addr          common.Address
	ClaimedAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEventClaimed is a free log retrieval operation binding the contract event 0x92c59d33103d56fd232bed949f5c3a793e67ab8b3b6de334a349fba4d901d4bc.
//
// Solidity: event EventClaimed(address indexed addr, uint256 claimedAmount)
func (_AGSClaim *AGSClaimFilterer) FilterEventClaimed(opts *bind.FilterOpts, addr []common.Address) (*AGSClaimEventClaimedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _AGSClaim.contract.FilterLogs(opts, "EventClaimed", addrRule)
	if err != nil {
		return nil, err
	}
	return &AGSClaimEventClaimedIterator{contract: _AGSClaim.contract, event: "EventClaimed", logs: logs, sub: sub}, nil
}

// WatchEventClaimed is a free log subscription operation binding the contract event 0x92c59d33103d56fd232bed949f5c3a793e67ab8b3b6de334a349fba4d901d4bc.
//
// Solidity: event EventClaimed(address indexed addr, uint256 claimedAmount)
func (_AGSClaim *AGSClaimFilterer) WatchEventClaimed(opts *bind.WatchOpts, sink chan<- *AGSClaimEventClaimed, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _AGSClaim.contract.WatchLogs(opts, "EventClaimed", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AGSClaimEventClaimed)
				if err := _AGSClaim.contract.UnpackLog(event, "EventClaimed", log); err != nil {
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

// ParseEventClaimed is a log parse operation binding the contract event 0x92c59d33103d56fd232bed949f5c3a793e67ab8b3b6de334a349fba4d901d4bc.
//
// Solidity: event EventClaimed(address indexed addr, uint256 claimedAmount)
func (_AGSClaim *AGSClaimFilterer) ParseEventClaimed(log types.Log) (*AGSClaimEventClaimed, error) {
	event := new(AGSClaimEventClaimed)
	if err := _AGSClaim.contract.UnpackLog(event, "EventClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AGSClaimEventClaimedAZ9Iterator is returned from FilterEventClaimedAZ9 and is used to iterate over the raw logs and unpacked data for EventClaimedAZ9 events raised by the AGSClaim contract.
type AGSClaimEventClaimedAZ9Iterator struct {
	Event *AGSClaimEventClaimedAZ9 // Event containing the contract specifics and raw log

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
func (it *AGSClaimEventClaimedAZ9Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AGSClaimEventClaimedAZ9)
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
		it.Event = new(AGSClaimEventClaimedAZ9)
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
func (it *AGSClaimEventClaimedAZ9Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AGSClaimEventClaimedAZ9Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AGSClaimEventClaimedAZ9 represents a EventClaimedAZ9 event raised by the AGSClaim contract.
type AGSClaimEventClaimedAZ9 struct {
	AddrAZ9       common.Hash
	ClaimedAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEventClaimedAZ9 is a free log retrieval operation binding the contract event 0xf09d7862a749844d0be7c0c15c0bf70d2dac7127deb9486545be11987affb040.
//
// Solidity: event EventClaimedAZ9(string indexed addrAZ9, uint256 claimedAmount)
func (_AGSClaim *AGSClaimFilterer) FilterEventClaimedAZ9(opts *bind.FilterOpts, addrAZ9 []string) (*AGSClaimEventClaimedAZ9Iterator, error) {

	var addrAZ9Rule []interface{}
	for _, addrAZ9Item := range addrAZ9 {
		addrAZ9Rule = append(addrAZ9Rule, addrAZ9Item)
	}

	logs, sub, err := _AGSClaim.contract.FilterLogs(opts, "EventClaimedAZ9", addrAZ9Rule)
	if err != nil {
		return nil, err
	}
	return &AGSClaimEventClaimedAZ9Iterator{contract: _AGSClaim.contract, event: "EventClaimedAZ9", logs: logs, sub: sub}, nil
}

// WatchEventClaimedAZ9 is a free log subscription operation binding the contract event 0xf09d7862a749844d0be7c0c15c0bf70d2dac7127deb9486545be11987affb040.
//
// Solidity: event EventClaimedAZ9(string indexed addrAZ9, uint256 claimedAmount)
func (_AGSClaim *AGSClaimFilterer) WatchEventClaimedAZ9(opts *bind.WatchOpts, sink chan<- *AGSClaimEventClaimedAZ9, addrAZ9 []string) (event.Subscription, error) {

	var addrAZ9Rule []interface{}
	for _, addrAZ9Item := range addrAZ9 {
		addrAZ9Rule = append(addrAZ9Rule, addrAZ9Item)
	}

	logs, sub, err := _AGSClaim.contract.WatchLogs(opts, "EventClaimedAZ9", addrAZ9Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AGSClaimEventClaimedAZ9)
				if err := _AGSClaim.contract.UnpackLog(event, "EventClaimedAZ9", log); err != nil {
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

// ParseEventClaimedAZ9 is a log parse operation binding the contract event 0xf09d7862a749844d0be7c0c15c0bf70d2dac7127deb9486545be11987affb040.
//
// Solidity: event EventClaimedAZ9(string indexed addrAZ9, uint256 claimedAmount)
func (_AGSClaim *AGSClaimFilterer) ParseEventClaimedAZ9(log types.Log) (*AGSClaimEventClaimedAZ9, error) {
	event := new(AGSClaimEventClaimedAZ9)
	if err := _AGSClaim.contract.UnpackLog(event, "EventClaimedAZ9", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
