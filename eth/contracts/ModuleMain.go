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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"PriceFeed\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subscriptionId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLane\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"callbackGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"vrfCoordinatorV2\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DAILY_REVIEWS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addNewUserToSystem\",\"inputs\":[{\"name\":\"roleType\",\"type\":\"uint8\",\"internalType\":\"enumLibTypeDef.RoleType\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelReservation\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkUpkeep\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"upkeepNeeded\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"performData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contributeRegistration\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initiator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"patient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountInWei\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"contributionLottery\",\"inputs\":[{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"donation\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"finishAppointment\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getProviderAverageScore\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReservationInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserToContract\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"performUpkeep\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rawFulfillRandomWords\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerFundRequest\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountInUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestReservation\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appointTimeSinceEpoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestReview\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startAppointment\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"updateProviderScore\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"provider\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"score\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFund\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"ModuleContribution__ContributionAlreadyRegistered\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"initiator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"patient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountInWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModuleContribution__ContributionRegistered\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"initiator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"patient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountInWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModuleContribution__LotteryCompleted\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModuleContribution__UpkeepPerformed\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"numWords\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModuleJudge__JudgeLotteryDone\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delayTime\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModuleJudge__ProviderScoreUpdated\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"score\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModuleReservation__AppointmentFinished\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountInWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModuleReservation__AppointmentStarted\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountInWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModuleReservation__ReservationCanceled\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountInWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ModuleReservation__ReservationRequested\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"appointTimeSinceEpoch\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"System__DonationLimitReached\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountInWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"System__FundWithdrawn\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountInWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"System__NewDonation\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountInWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"System__NewFundRegistered\",\"inputs\":[{\"name\":\"fundInfoIndex\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amountInUsd\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"roundId\",\"type\":\"uint80\",\"indexed\":true,\"internalType\":\"uint80\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"System__NewUserAdded\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"roleType\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"System__NewUserContractAdded\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"roleType\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ModuleContribution__UpkeepNotNeeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ModuleReservation__NoReservation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ModuleReservation__TransferFail\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoordinatorCanFulfill\",\"inputs\":[{\"name\":\"have\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"want\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// DAILYREVIEWSLIMIT is a free data retrieval call binding the contract method 0x8452b8e2.
//
// Solidity: function DAILY_REVIEWS_LIMIT() view returns(uint32)
func (_Contracts *ContractsCaller) DAILYREVIEWSLIMIT(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "DAILY_REVIEWS_LIMIT")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DAILYREVIEWSLIMIT is a free data retrieval call binding the contract method 0x8452b8e2.
//
// Solidity: function DAILY_REVIEWS_LIMIT() view returns(uint32)
func (_Contracts *ContractsSession) DAILYREVIEWSLIMIT() (uint32, error) {
	return _Contracts.Contract.DAILYREVIEWSLIMIT(&_Contracts.CallOpts)
}

// DAILYREVIEWSLIMIT is a free data retrieval call binding the contract method 0x8452b8e2.
//
// Solidity: function DAILY_REVIEWS_LIMIT() view returns(uint32)
func (_Contracts *ContractsCallerSession) DAILYREVIEWSLIMIT() (uint32, error) {
	return _Contracts.Contract.DAILYREVIEWSLIMIT(&_Contracts.CallOpts)
}

// CheckUpkeep is a free data retrieval call binding the contract method 0x6e04ff0d.
//
// Solidity: function checkUpkeep(bytes ) view returns(bool upkeepNeeded, bytes performData)
func (_Contracts *ContractsCaller) CheckUpkeep(opts *bind.CallOpts, arg0 []byte) (struct {
	UpkeepNeeded bool
	PerformData  []byte
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "checkUpkeep", arg0)

	outstruct := new(struct {
		UpkeepNeeded bool
		PerformData  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UpkeepNeeded = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PerformData = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// CheckUpkeep is a free data retrieval call binding the contract method 0x6e04ff0d.
//
// Solidity: function checkUpkeep(bytes ) view returns(bool upkeepNeeded, bytes performData)
func (_Contracts *ContractsSession) CheckUpkeep(arg0 []byte) (struct {
	UpkeepNeeded bool
	PerformData  []byte
}, error) {
	return _Contracts.Contract.CheckUpkeep(&_Contracts.CallOpts, arg0)
}

// CheckUpkeep is a free data retrieval call binding the contract method 0x6e04ff0d.
//
// Solidity: function checkUpkeep(bytes ) view returns(bool upkeepNeeded, bytes performData)
func (_Contracts *ContractsCallerSession) CheckUpkeep(arg0 []byte) (struct {
	UpkeepNeeded bool
	PerformData  []byte
}, error) {
	return _Contracts.Contract.CheckUpkeep(&_Contracts.CallOpts, arg0)
}

// GetProviderAverageScore is a free data retrieval call binding the contract method 0xb0002779.
//
// Solidity: function getProviderAverageScore(address provider) view returns(uint256, uint256)
func (_Contracts *ContractsCaller) GetProviderAverageScore(opts *bind.CallOpts, provider common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getProviderAverageScore", provider)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetProviderAverageScore is a free data retrieval call binding the contract method 0xb0002779.
//
// Solidity: function getProviderAverageScore(address provider) view returns(uint256, uint256)
func (_Contracts *ContractsSession) GetProviderAverageScore(provider common.Address) (*big.Int, *big.Int, error) {
	return _Contracts.Contract.GetProviderAverageScore(&_Contracts.CallOpts, provider)
}

// GetProviderAverageScore is a free data retrieval call binding the contract method 0xb0002779.
//
// Solidity: function getProviderAverageScore(address provider) view returns(uint256, uint256)
func (_Contracts *ContractsCallerSession) GetProviderAverageScore(provider common.Address) (*big.Int, *big.Int, error) {
	return _Contracts.Contract.GetProviderAverageScore(&_Contracts.CallOpts, provider)
}

// GetReservationInfo is a free data retrieval call binding the contract method 0x8c868adf.
//
// Solidity: function getReservationInfo() view returns(address)
func (_Contracts *ContractsCaller) GetReservationInfo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getReservationInfo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReservationInfo is a free data retrieval call binding the contract method 0x8c868adf.
//
// Solidity: function getReservationInfo() view returns(address)
func (_Contracts *ContractsSession) GetReservationInfo() (common.Address, error) {
	return _Contracts.Contract.GetReservationInfo(&_Contracts.CallOpts)
}

// GetReservationInfo is a free data retrieval call binding the contract method 0x8c868adf.
//
// Solidity: function getReservationInfo() view returns(address)
func (_Contracts *ContractsCallerSession) GetReservationInfo() (common.Address, error) {
	return _Contracts.Contract.GetReservationInfo(&_Contracts.CallOpts)
}

// GetUserToContract is a free data retrieval call binding the contract method 0x2671974c.
//
// Solidity: function getUserToContract(address user) view returns(address)
func (_Contracts *ContractsCaller) GetUserToContract(opts *bind.CallOpts, user common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getUserToContract", user)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUserToContract is a free data retrieval call binding the contract method 0x2671974c.
//
// Solidity: function getUserToContract(address user) view returns(address)
func (_Contracts *ContractsSession) GetUserToContract(user common.Address) (common.Address, error) {
	return _Contracts.Contract.GetUserToContract(&_Contracts.CallOpts, user)
}

// GetUserToContract is a free data retrieval call binding the contract method 0x2671974c.
//
// Solidity: function getUserToContract(address user) view returns(address)
func (_Contracts *ContractsCallerSession) GetUserToContract(user common.Address) (common.Address, error) {
	return _Contracts.Contract.GetUserToContract(&_Contracts.CallOpts, user)
}

// AddNewUserToSystem is a paid mutator transaction binding the contract method 0x5d8f96d4.
//
// Solidity: function addNewUserToSystem(uint8 roleType) returns()
func (_Contracts *ContractsTransactor) AddNewUserToSystem(opts *bind.TransactOpts, roleType uint8) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addNewUserToSystem", roleType)
}

// AddNewUserToSystem is a paid mutator transaction binding the contract method 0x5d8f96d4.
//
// Solidity: function addNewUserToSystem(uint8 roleType) returns()
func (_Contracts *ContractsSession) AddNewUserToSystem(roleType uint8) (*types.Transaction, error) {
	return _Contracts.Contract.AddNewUserToSystem(&_Contracts.TransactOpts, roleType)
}

// AddNewUserToSystem is a paid mutator transaction binding the contract method 0x5d8f96d4.
//
// Solidity: function addNewUserToSystem(uint8 roleType) returns()
func (_Contracts *ContractsTransactorSession) AddNewUserToSystem(roleType uint8) (*types.Transaction, error) {
	return _Contracts.Contract.AddNewUserToSystem(&_Contracts.TransactOpts, roleType)
}

// CancelReservation is a paid mutator transaction binding the contract method 0xdad41717.
//
// Solidity: function cancelReservation(address provider) returns()
func (_Contracts *ContractsTransactor) CancelReservation(opts *bind.TransactOpts, provider common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "cancelReservation", provider)
}

// CancelReservation is a paid mutator transaction binding the contract method 0xdad41717.
//
// Solidity: function cancelReservation(address provider) returns()
func (_Contracts *ContractsSession) CancelReservation(provider common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.CancelReservation(&_Contracts.TransactOpts, provider)
}

// CancelReservation is a paid mutator transaction binding the contract method 0xdad41717.
//
// Solidity: function cancelReservation(address provider) returns()
func (_Contracts *ContractsTransactorSession) CancelReservation(provider common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.CancelReservation(&_Contracts.TransactOpts, provider)
}

// ContributeRegistration is a paid mutator transaction binding the contract method 0xa95c8163.
//
// Solidity: function contributeRegistration(address provider, address initiator, address patient, uint256 amountInWei) payable returns()
func (_Contracts *ContractsTransactor) ContributeRegistration(opts *bind.TransactOpts, provider common.Address, initiator common.Address, patient common.Address, amountInWei *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "contributeRegistration", provider, initiator, patient, amountInWei)
}

// ContributeRegistration is a paid mutator transaction binding the contract method 0xa95c8163.
//
// Solidity: function contributeRegistration(address provider, address initiator, address patient, uint256 amountInWei) payable returns()
func (_Contracts *ContractsSession) ContributeRegistration(provider common.Address, initiator common.Address, patient common.Address, amountInWei *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ContributeRegistration(&_Contracts.TransactOpts, provider, initiator, patient, amountInWei)
}

// ContributeRegistration is a paid mutator transaction binding the contract method 0xa95c8163.
//
// Solidity: function contributeRegistration(address provider, address initiator, address patient, uint256 amountInWei) payable returns()
func (_Contracts *ContractsTransactorSession) ContributeRegistration(provider common.Address, initiator common.Address, patient common.Address, amountInWei *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ContributeRegistration(&_Contracts.TransactOpts, provider, initiator, patient, amountInWei)
}

// ContributionLottery is a paid mutator transaction binding the contract method 0x0e16b989.
//
// Solidity: function contributionLottery(uint256[] randomWords) payable returns()
func (_Contracts *ContractsTransactor) ContributionLottery(opts *bind.TransactOpts, randomWords []*big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "contributionLottery", randomWords)
}

// ContributionLottery is a paid mutator transaction binding the contract method 0x0e16b989.
//
// Solidity: function contributionLottery(uint256[] randomWords) payable returns()
func (_Contracts *ContractsSession) ContributionLottery(randomWords []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ContributionLottery(&_Contracts.TransactOpts, randomWords)
}

// ContributionLottery is a paid mutator transaction binding the contract method 0x0e16b989.
//
// Solidity: function contributionLottery(uint256[] randomWords) payable returns()
func (_Contracts *ContractsTransactorSession) ContributionLottery(randomWords []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ContributionLottery(&_Contracts.TransactOpts, randomWords)
}

// Donation is a paid mutator transaction binding the contract method 0x8b5c9b64.
//
// Solidity: function donation(uint256 index, address user) payable returns()
func (_Contracts *ContractsTransactor) Donation(opts *bind.TransactOpts, index *big.Int, user common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "donation", index, user)
}

// Donation is a paid mutator transaction binding the contract method 0x8b5c9b64.
//
// Solidity: function donation(uint256 index, address user) payable returns()
func (_Contracts *ContractsSession) Donation(index *big.Int, user common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Donation(&_Contracts.TransactOpts, index, user)
}

// Donation is a paid mutator transaction binding the contract method 0x8b5c9b64.
//
// Solidity: function donation(uint256 index, address user) payable returns()
func (_Contracts *ContractsTransactorSession) Donation(index *big.Int, user common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Donation(&_Contracts.TransactOpts, index, user)
}

// FinishAppointment is a paid mutator transaction binding the contract method 0xae270559.
//
// Solidity: function finishAppointment(address provider) payable returns()
func (_Contracts *ContractsTransactor) FinishAppointment(opts *bind.TransactOpts, provider common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "finishAppointment", provider)
}

// FinishAppointment is a paid mutator transaction binding the contract method 0xae270559.
//
// Solidity: function finishAppointment(address provider) payable returns()
func (_Contracts *ContractsSession) FinishAppointment(provider common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.FinishAppointment(&_Contracts.TransactOpts, provider)
}

// FinishAppointment is a paid mutator transaction binding the contract method 0xae270559.
//
// Solidity: function finishAppointment(address provider) payable returns()
func (_Contracts *ContractsTransactorSession) FinishAppointment(provider common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.FinishAppointment(&_Contracts.TransactOpts, provider)
}

// PerformUpkeep is a paid mutator transaction binding the contract method 0x4585e33b.
//
// Solidity: function performUpkeep(bytes ) returns()
func (_Contracts *ContractsTransactor) PerformUpkeep(opts *bind.TransactOpts, arg0 []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "performUpkeep", arg0)
}

// PerformUpkeep is a paid mutator transaction binding the contract method 0x4585e33b.
//
// Solidity: function performUpkeep(bytes ) returns()
func (_Contracts *ContractsSession) PerformUpkeep(arg0 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.PerformUpkeep(&_Contracts.TransactOpts, arg0)
}

// PerformUpkeep is a paid mutator transaction binding the contract method 0x4585e33b.
//
// Solidity: function performUpkeep(bytes ) returns()
func (_Contracts *ContractsTransactorSession) PerformUpkeep(arg0 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.PerformUpkeep(&_Contracts.TransactOpts, arg0)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_Contracts *ContractsTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_Contracts *ContractsSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RawFulfillRandomWords(&_Contracts.TransactOpts, requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_Contracts *ContractsTransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RawFulfillRandomWords(&_Contracts.TransactOpts, requestId, randomWords)
}

// RegisterFundRequest is a paid mutator transaction binding the contract method 0xc6443ecd.
//
// Solidity: function registerFundRequest(address user, address contractAddress, uint256 amountInUsd) returns()
func (_Contracts *ContractsTransactor) RegisterFundRequest(opts *bind.TransactOpts, user common.Address, contractAddress common.Address, amountInUsd *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerFundRequest", user, contractAddress, amountInUsd)
}

// RegisterFundRequest is a paid mutator transaction binding the contract method 0xc6443ecd.
//
// Solidity: function registerFundRequest(address user, address contractAddress, uint256 amountInUsd) returns()
func (_Contracts *ContractsSession) RegisterFundRequest(user common.Address, contractAddress common.Address, amountInUsd *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterFundRequest(&_Contracts.TransactOpts, user, contractAddress, amountInUsd)
}

// RegisterFundRequest is a paid mutator transaction binding the contract method 0xc6443ecd.
//
// Solidity: function registerFundRequest(address user, address contractAddress, uint256 amountInUsd) returns()
func (_Contracts *ContractsTransactorSession) RegisterFundRequest(user common.Address, contractAddress common.Address, amountInUsd *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterFundRequest(&_Contracts.TransactOpts, user, contractAddress, amountInUsd)
}

// RequestReservation is a paid mutator transaction binding the contract method 0x38a6333d.
//
// Solidity: function requestReservation(address provider, uint256 appointTimeSinceEpoch) payable returns()
func (_Contracts *ContractsTransactor) RequestReservation(opts *bind.TransactOpts, provider common.Address, appointTimeSinceEpoch *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "requestReservation", provider, appointTimeSinceEpoch)
}

// RequestReservation is a paid mutator transaction binding the contract method 0x38a6333d.
//
// Solidity: function requestReservation(address provider, uint256 appointTimeSinceEpoch) payable returns()
func (_Contracts *ContractsSession) RequestReservation(provider common.Address, appointTimeSinceEpoch *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RequestReservation(&_Contracts.TransactOpts, provider, appointTimeSinceEpoch)
}

// RequestReservation is a paid mutator transaction binding the contract method 0x38a6333d.
//
// Solidity: function requestReservation(address provider, uint256 appointTimeSinceEpoch) payable returns()
func (_Contracts *ContractsTransactorSession) RequestReservation(provider common.Address, appointTimeSinceEpoch *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RequestReservation(&_Contracts.TransactOpts, provider, appointTimeSinceEpoch)
}

// RequestReview is a paid mutator transaction binding the contract method 0xc71f8345.
//
// Solidity: function requestReview(address provider) returns(uint256 requestId)
func (_Contracts *ContractsTransactor) RequestReview(opts *bind.TransactOpts, provider common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "requestReview", provider)
}

// RequestReview is a paid mutator transaction binding the contract method 0xc71f8345.
//
// Solidity: function requestReview(address provider) returns(uint256 requestId)
func (_Contracts *ContractsSession) RequestReview(provider common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RequestReview(&_Contracts.TransactOpts, provider)
}

// RequestReview is a paid mutator transaction binding the contract method 0xc71f8345.
//
// Solidity: function requestReview(address provider) returns(uint256 requestId)
func (_Contracts *ContractsTransactorSession) RequestReview(provider common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RequestReview(&_Contracts.TransactOpts, provider)
}

// StartAppointment is a paid mutator transaction binding the contract method 0x9dc5cec3.
//
// Solidity: function startAppointment(address provider) payable returns()
func (_Contracts *ContractsTransactor) StartAppointment(opts *bind.TransactOpts, provider common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "startAppointment", provider)
}

// StartAppointment is a paid mutator transaction binding the contract method 0x9dc5cec3.
//
// Solidity: function startAppointment(address provider) payable returns()
func (_Contracts *ContractsSession) StartAppointment(provider common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.StartAppointment(&_Contracts.TransactOpts, provider)
}

// StartAppointment is a paid mutator transaction binding the contract method 0x9dc5cec3.
//
// Solidity: function startAppointment(address provider) payable returns()
func (_Contracts *ContractsTransactorSession) StartAppointment(provider common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.StartAppointment(&_Contracts.TransactOpts, provider)
}

// UpdateProviderScore is a paid mutator transaction binding the contract method 0x24e30dec.
//
// Solidity: function updateProviderScore(address user, address provider, uint256 score) returns()
func (_Contracts *ContractsTransactor) UpdateProviderScore(opts *bind.TransactOpts, user common.Address, provider common.Address, score *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateProviderScore", user, provider, score)
}

// UpdateProviderScore is a paid mutator transaction binding the contract method 0x24e30dec.
//
// Solidity: function updateProviderScore(address user, address provider, uint256 score) returns()
func (_Contracts *ContractsSession) UpdateProviderScore(user common.Address, provider common.Address, score *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateProviderScore(&_Contracts.TransactOpts, user, provider, score)
}

// UpdateProviderScore is a paid mutator transaction binding the contract method 0x24e30dec.
//
// Solidity: function updateProviderScore(address user, address provider, uint256 score) returns()
func (_Contracts *ContractsTransactorSession) UpdateProviderScore(user common.Address, provider common.Address, score *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateProviderScore(&_Contracts.TransactOpts, user, provider, score)
}

// WithdrawFund is a paid mutator transaction binding the contract method 0x0cee1725.
//
// Solidity: function withdrawFund(uint256 index) payable returns()
func (_Contracts *ContractsTransactor) WithdrawFund(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdrawFund", index)
}

// WithdrawFund is a paid mutator transaction binding the contract method 0x0cee1725.
//
// Solidity: function withdrawFund(uint256 index) payable returns()
func (_Contracts *ContractsSession) WithdrawFund(index *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawFund(&_Contracts.TransactOpts, index)
}

// WithdrawFund is a paid mutator transaction binding the contract method 0x0cee1725.
//
// Solidity: function withdrawFund(uint256 index) payable returns()
func (_Contracts *ContractsTransactorSession) WithdrawFund(index *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawFund(&_Contracts.TransactOpts, index)
}

// ContractsModuleContributionContributionAlreadyRegisteredIterator is returned from FilterModuleContributionContributionAlreadyRegistered and is used to iterate over the raw logs and unpacked data for ModuleContributionContributionAlreadyRegistered events raised by the Contracts contract.
type ContractsModuleContributionContributionAlreadyRegisteredIterator struct {
	Event *ContractsModuleContributionContributionAlreadyRegistered // Event containing the contract specifics and raw log

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
func (it *ContractsModuleContributionContributionAlreadyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsModuleContributionContributionAlreadyRegistered)
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
		it.Event = new(ContractsModuleContributionContributionAlreadyRegistered)
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
func (it *ContractsModuleContributionContributionAlreadyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsModuleContributionContributionAlreadyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsModuleContributionContributionAlreadyRegistered represents a ModuleContributionContributionAlreadyRegistered event raised by the Contracts contract.
type ContractsModuleContributionContributionAlreadyRegistered struct {
	Provider    common.Address
	Initiator   common.Address
	Patient     common.Address
	AmountInWei *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterModuleContributionContributionAlreadyRegistered is a free log retrieval operation binding the contract event 0xbaba9b431c740877c581a305d5a1c75f980e5b9078d7dbed59d0f02703a6b5e7.
//
// Solidity: event ModuleContribution__ContributionAlreadyRegistered(address indexed provider, address indexed initiator, address indexed patient, uint256 amountInWei)
func (_Contracts *ContractsFilterer) FilterModuleContributionContributionAlreadyRegistered(opts *bind.FilterOpts, provider []common.Address, initiator []common.Address, patient []common.Address) (*ContractsModuleContributionContributionAlreadyRegisteredIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}
	var patientRule []interface{}
	for _, patientItem := range patient {
		patientRule = append(patientRule, patientItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ModuleContribution__ContributionAlreadyRegistered", providerRule, initiatorRule, patientRule)
	if err != nil {
		return nil, err
	}
	return &ContractsModuleContributionContributionAlreadyRegisteredIterator{contract: _Contracts.contract, event: "ModuleContribution__ContributionAlreadyRegistered", logs: logs, sub: sub}, nil
}

// WatchModuleContributionContributionAlreadyRegistered is a free log subscription operation binding the contract event 0xbaba9b431c740877c581a305d5a1c75f980e5b9078d7dbed59d0f02703a6b5e7.
//
// Solidity: event ModuleContribution__ContributionAlreadyRegistered(address indexed provider, address indexed initiator, address indexed patient, uint256 amountInWei)
func (_Contracts *ContractsFilterer) WatchModuleContributionContributionAlreadyRegistered(opts *bind.WatchOpts, sink chan<- *ContractsModuleContributionContributionAlreadyRegistered, provider []common.Address, initiator []common.Address, patient []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}
	var patientRule []interface{}
	for _, patientItem := range patient {
		patientRule = append(patientRule, patientItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ModuleContribution__ContributionAlreadyRegistered", providerRule, initiatorRule, patientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsModuleContributionContributionAlreadyRegistered)
				if err := _Contracts.contract.UnpackLog(event, "ModuleContribution__ContributionAlreadyRegistered", log); err != nil {
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

// ParseModuleContributionContributionAlreadyRegistered is a log parse operation binding the contract event 0xbaba9b431c740877c581a305d5a1c75f980e5b9078d7dbed59d0f02703a6b5e7.
//
// Solidity: event ModuleContribution__ContributionAlreadyRegistered(address indexed provider, address indexed initiator, address indexed patient, uint256 amountInWei)
func (_Contracts *ContractsFilterer) ParseModuleContributionContributionAlreadyRegistered(log types.Log) (*ContractsModuleContributionContributionAlreadyRegistered, error) {
	event := new(ContractsModuleContributionContributionAlreadyRegistered)
	if err := _Contracts.contract.UnpackLog(event, "ModuleContribution__ContributionAlreadyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsModuleContributionContributionRegisteredIterator is returned from FilterModuleContributionContributionRegistered and is used to iterate over the raw logs and unpacked data for ModuleContributionContributionRegistered events raised by the Contracts contract.
type ContractsModuleContributionContributionRegisteredIterator struct {
	Event *ContractsModuleContributionContributionRegistered // Event containing the contract specifics and raw log

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
func (it *ContractsModuleContributionContributionRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsModuleContributionContributionRegistered)
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
		it.Event = new(ContractsModuleContributionContributionRegistered)
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
func (it *ContractsModuleContributionContributionRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsModuleContributionContributionRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsModuleContributionContributionRegistered represents a ModuleContributionContributionRegistered event raised by the Contracts contract.
type ContractsModuleContributionContributionRegistered struct {
	Provider    common.Address
	Initiator   common.Address
	Patient     common.Address
	AmountInWei *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterModuleContributionContributionRegistered is a free log retrieval operation binding the contract event 0x665b5c6609bc07d181d7afedbd62df710c4b4b748c869eedbefcaf547d4eaec8.
//
// Solidity: event ModuleContribution__ContributionRegistered(address indexed provider, address indexed initiator, address indexed patient, uint256 amountInWei)
func (_Contracts *ContractsFilterer) FilterModuleContributionContributionRegistered(opts *bind.FilterOpts, provider []common.Address, initiator []common.Address, patient []common.Address) (*ContractsModuleContributionContributionRegisteredIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}
	var patientRule []interface{}
	for _, patientItem := range patient {
		patientRule = append(patientRule, patientItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ModuleContribution__ContributionRegistered", providerRule, initiatorRule, patientRule)
	if err != nil {
		return nil, err
	}
	return &ContractsModuleContributionContributionRegisteredIterator{contract: _Contracts.contract, event: "ModuleContribution__ContributionRegistered", logs: logs, sub: sub}, nil
}

// WatchModuleContributionContributionRegistered is a free log subscription operation binding the contract event 0x665b5c6609bc07d181d7afedbd62df710c4b4b748c869eedbefcaf547d4eaec8.
//
// Solidity: event ModuleContribution__ContributionRegistered(address indexed provider, address indexed initiator, address indexed patient, uint256 amountInWei)
func (_Contracts *ContractsFilterer) WatchModuleContributionContributionRegistered(opts *bind.WatchOpts, sink chan<- *ContractsModuleContributionContributionRegistered, provider []common.Address, initiator []common.Address, patient []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}
	var patientRule []interface{}
	for _, patientItem := range patient {
		patientRule = append(patientRule, patientItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ModuleContribution__ContributionRegistered", providerRule, initiatorRule, patientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsModuleContributionContributionRegistered)
				if err := _Contracts.contract.UnpackLog(event, "ModuleContribution__ContributionRegistered", log); err != nil {
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

// ParseModuleContributionContributionRegistered is a log parse operation binding the contract event 0x665b5c6609bc07d181d7afedbd62df710c4b4b748c869eedbefcaf547d4eaec8.
//
// Solidity: event ModuleContribution__ContributionRegistered(address indexed provider, address indexed initiator, address indexed patient, uint256 amountInWei)
func (_Contracts *ContractsFilterer) ParseModuleContributionContributionRegistered(log types.Log) (*ContractsModuleContributionContributionRegistered, error) {
	event := new(ContractsModuleContributionContributionRegistered)
	if err := _Contracts.contract.UnpackLog(event, "ModuleContribution__ContributionRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsModuleContributionLotteryCompletedIterator is returned from FilterModuleContributionLotteryCompleted and is used to iterate over the raw logs and unpacked data for ModuleContributionLotteryCompleted events raised by the Contracts contract.
type ContractsModuleContributionLotteryCompletedIterator struct {
	Event *ContractsModuleContributionLotteryCompleted // Event containing the contract specifics and raw log

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
func (it *ContractsModuleContributionLotteryCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsModuleContributionLotteryCompleted)
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
		it.Event = new(ContractsModuleContributionLotteryCompleted)
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
func (it *ContractsModuleContributionLotteryCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsModuleContributionLotteryCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsModuleContributionLotteryCompleted represents a ModuleContributionLotteryCompleted event raised by the Contracts contract.
type ContractsModuleContributionLotteryCompleted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterModuleContributionLotteryCompleted is a free log retrieval operation binding the contract event 0x561c2ead36037214294bcb7e1beb9e557448ccc2579e4f6a7dc3ed91d026b2f1.
//
// Solidity: event ModuleContribution__LotteryCompleted()
func (_Contracts *ContractsFilterer) FilterModuleContributionLotteryCompleted(opts *bind.FilterOpts) (*ContractsModuleContributionLotteryCompletedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ModuleContribution__LotteryCompleted")
	if err != nil {
		return nil, err
	}
	return &ContractsModuleContributionLotteryCompletedIterator{contract: _Contracts.contract, event: "ModuleContribution__LotteryCompleted", logs: logs, sub: sub}, nil
}

// WatchModuleContributionLotteryCompleted is a free log subscription operation binding the contract event 0x561c2ead36037214294bcb7e1beb9e557448ccc2579e4f6a7dc3ed91d026b2f1.
//
// Solidity: event ModuleContribution__LotteryCompleted()
func (_Contracts *ContractsFilterer) WatchModuleContributionLotteryCompleted(opts *bind.WatchOpts, sink chan<- *ContractsModuleContributionLotteryCompleted) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ModuleContribution__LotteryCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsModuleContributionLotteryCompleted)
				if err := _Contracts.contract.UnpackLog(event, "ModuleContribution__LotteryCompleted", log); err != nil {
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

// ParseModuleContributionLotteryCompleted is a log parse operation binding the contract event 0x561c2ead36037214294bcb7e1beb9e557448ccc2579e4f6a7dc3ed91d026b2f1.
//
// Solidity: event ModuleContribution__LotteryCompleted()
func (_Contracts *ContractsFilterer) ParseModuleContributionLotteryCompleted(log types.Log) (*ContractsModuleContributionLotteryCompleted, error) {
	event := new(ContractsModuleContributionLotteryCompleted)
	if err := _Contracts.contract.UnpackLog(event, "ModuleContribution__LotteryCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsModuleContributionUpkeepPerformedIterator is returned from FilterModuleContributionUpkeepPerformed and is used to iterate over the raw logs and unpacked data for ModuleContributionUpkeepPerformed events raised by the Contracts contract.
type ContractsModuleContributionUpkeepPerformedIterator struct {
	Event *ContractsModuleContributionUpkeepPerformed // Event containing the contract specifics and raw log

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
func (it *ContractsModuleContributionUpkeepPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsModuleContributionUpkeepPerformed)
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
		it.Event = new(ContractsModuleContributionUpkeepPerformed)
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
func (it *ContractsModuleContributionUpkeepPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsModuleContributionUpkeepPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsModuleContributionUpkeepPerformed represents a ModuleContributionUpkeepPerformed event raised by the Contracts contract.
type ContractsModuleContributionUpkeepPerformed struct {
	RequestId *big.Int
	NumWords  uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterModuleContributionUpkeepPerformed is a free log retrieval operation binding the contract event 0xd1cf0d5356bbc998fe713d2ced914605cc1d4b886b6f5440321480f02b532ef4.
//
// Solidity: event ModuleContribution__UpkeepPerformed(uint256 requestId, uint32 numWords)
func (_Contracts *ContractsFilterer) FilterModuleContributionUpkeepPerformed(opts *bind.FilterOpts) (*ContractsModuleContributionUpkeepPerformedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ModuleContribution__UpkeepPerformed")
	if err != nil {
		return nil, err
	}
	return &ContractsModuleContributionUpkeepPerformedIterator{contract: _Contracts.contract, event: "ModuleContribution__UpkeepPerformed", logs: logs, sub: sub}, nil
}

// WatchModuleContributionUpkeepPerformed is a free log subscription operation binding the contract event 0xd1cf0d5356bbc998fe713d2ced914605cc1d4b886b6f5440321480f02b532ef4.
//
// Solidity: event ModuleContribution__UpkeepPerformed(uint256 requestId, uint32 numWords)
func (_Contracts *ContractsFilterer) WatchModuleContributionUpkeepPerformed(opts *bind.WatchOpts, sink chan<- *ContractsModuleContributionUpkeepPerformed) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ModuleContribution__UpkeepPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsModuleContributionUpkeepPerformed)
				if err := _Contracts.contract.UnpackLog(event, "ModuleContribution__UpkeepPerformed", log); err != nil {
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

// ParseModuleContributionUpkeepPerformed is a log parse operation binding the contract event 0xd1cf0d5356bbc998fe713d2ced914605cc1d4b886b6f5440321480f02b532ef4.
//
// Solidity: event ModuleContribution__UpkeepPerformed(uint256 requestId, uint32 numWords)
func (_Contracts *ContractsFilterer) ParseModuleContributionUpkeepPerformed(log types.Log) (*ContractsModuleContributionUpkeepPerformed, error) {
	event := new(ContractsModuleContributionUpkeepPerformed)
	if err := _Contracts.contract.UnpackLog(event, "ModuleContribution__UpkeepPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsModuleJudgeJudgeLotteryDoneIterator is returned from FilterModuleJudgeJudgeLotteryDone and is used to iterate over the raw logs and unpacked data for ModuleJudgeJudgeLotteryDone events raised by the Contracts contract.
type ContractsModuleJudgeJudgeLotteryDoneIterator struct {
	Event *ContractsModuleJudgeJudgeLotteryDone // Event containing the contract specifics and raw log

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
func (it *ContractsModuleJudgeJudgeLotteryDoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsModuleJudgeJudgeLotteryDone)
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
		it.Event = new(ContractsModuleJudgeJudgeLotteryDone)
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
func (it *ContractsModuleJudgeJudgeLotteryDoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsModuleJudgeJudgeLotteryDoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsModuleJudgeJudgeLotteryDone represents a ModuleJudgeJudgeLotteryDone event raised by the Contracts contract.
type ContractsModuleJudgeJudgeLotteryDone struct {
	User      common.Address
	Provider  common.Address
	DelayTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterModuleJudgeJudgeLotteryDone is a free log retrieval operation binding the contract event 0x053c9e01770f245c41b825bb8fb9970669ecbf50fdff48de3d2eb2282397682d.
//
// Solidity: event ModuleJudge__JudgeLotteryDone(address indexed user, address indexed provider, uint256 indexed delayTime)
func (_Contracts *ContractsFilterer) FilterModuleJudgeJudgeLotteryDone(opts *bind.FilterOpts, user []common.Address, provider []common.Address, delayTime []*big.Int) (*ContractsModuleJudgeJudgeLotteryDoneIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var delayTimeRule []interface{}
	for _, delayTimeItem := range delayTime {
		delayTimeRule = append(delayTimeRule, delayTimeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ModuleJudge__JudgeLotteryDone", userRule, providerRule, delayTimeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsModuleJudgeJudgeLotteryDoneIterator{contract: _Contracts.contract, event: "ModuleJudge__JudgeLotteryDone", logs: logs, sub: sub}, nil
}

// WatchModuleJudgeJudgeLotteryDone is a free log subscription operation binding the contract event 0x053c9e01770f245c41b825bb8fb9970669ecbf50fdff48de3d2eb2282397682d.
//
// Solidity: event ModuleJudge__JudgeLotteryDone(address indexed user, address indexed provider, uint256 indexed delayTime)
func (_Contracts *ContractsFilterer) WatchModuleJudgeJudgeLotteryDone(opts *bind.WatchOpts, sink chan<- *ContractsModuleJudgeJudgeLotteryDone, user []common.Address, provider []common.Address, delayTime []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var delayTimeRule []interface{}
	for _, delayTimeItem := range delayTime {
		delayTimeRule = append(delayTimeRule, delayTimeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ModuleJudge__JudgeLotteryDone", userRule, providerRule, delayTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsModuleJudgeJudgeLotteryDone)
				if err := _Contracts.contract.UnpackLog(event, "ModuleJudge__JudgeLotteryDone", log); err != nil {
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

// ParseModuleJudgeJudgeLotteryDone is a log parse operation binding the contract event 0x053c9e01770f245c41b825bb8fb9970669ecbf50fdff48de3d2eb2282397682d.
//
// Solidity: event ModuleJudge__JudgeLotteryDone(address indexed user, address indexed provider, uint256 indexed delayTime)
func (_Contracts *ContractsFilterer) ParseModuleJudgeJudgeLotteryDone(log types.Log) (*ContractsModuleJudgeJudgeLotteryDone, error) {
	event := new(ContractsModuleJudgeJudgeLotteryDone)
	if err := _Contracts.contract.UnpackLog(event, "ModuleJudge__JudgeLotteryDone", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsModuleJudgeProviderScoreUpdatedIterator is returned from FilterModuleJudgeProviderScoreUpdated and is used to iterate over the raw logs and unpacked data for ModuleJudgeProviderScoreUpdated events raised by the Contracts contract.
type ContractsModuleJudgeProviderScoreUpdatedIterator struct {
	Event *ContractsModuleJudgeProviderScoreUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsModuleJudgeProviderScoreUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsModuleJudgeProviderScoreUpdated)
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
		it.Event = new(ContractsModuleJudgeProviderScoreUpdated)
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
func (it *ContractsModuleJudgeProviderScoreUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsModuleJudgeProviderScoreUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsModuleJudgeProviderScoreUpdated represents a ModuleJudgeProviderScoreUpdated event raised by the Contracts contract.
type ContractsModuleJudgeProviderScoreUpdated struct {
	User     common.Address
	Provider common.Address
	Score    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterModuleJudgeProviderScoreUpdated is a free log retrieval operation binding the contract event 0x831b595405735b7b402a60174d869685406f72f4f6ed692683a00a6148c5c894.
//
// Solidity: event ModuleJudge__ProviderScoreUpdated(address indexed user, address indexed provider, uint256 indexed score)
func (_Contracts *ContractsFilterer) FilterModuleJudgeProviderScoreUpdated(opts *bind.FilterOpts, user []common.Address, provider []common.Address, score []*big.Int) (*ContractsModuleJudgeProviderScoreUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var scoreRule []interface{}
	for _, scoreItem := range score {
		scoreRule = append(scoreRule, scoreItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ModuleJudge__ProviderScoreUpdated", userRule, providerRule, scoreRule)
	if err != nil {
		return nil, err
	}
	return &ContractsModuleJudgeProviderScoreUpdatedIterator{contract: _Contracts.contract, event: "ModuleJudge__ProviderScoreUpdated", logs: logs, sub: sub}, nil
}

// WatchModuleJudgeProviderScoreUpdated is a free log subscription operation binding the contract event 0x831b595405735b7b402a60174d869685406f72f4f6ed692683a00a6148c5c894.
//
// Solidity: event ModuleJudge__ProviderScoreUpdated(address indexed user, address indexed provider, uint256 indexed score)
func (_Contracts *ContractsFilterer) WatchModuleJudgeProviderScoreUpdated(opts *bind.WatchOpts, sink chan<- *ContractsModuleJudgeProviderScoreUpdated, user []common.Address, provider []common.Address, score []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var scoreRule []interface{}
	for _, scoreItem := range score {
		scoreRule = append(scoreRule, scoreItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ModuleJudge__ProviderScoreUpdated", userRule, providerRule, scoreRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsModuleJudgeProviderScoreUpdated)
				if err := _Contracts.contract.UnpackLog(event, "ModuleJudge__ProviderScoreUpdated", log); err != nil {
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

// ParseModuleJudgeProviderScoreUpdated is a log parse operation binding the contract event 0x831b595405735b7b402a60174d869685406f72f4f6ed692683a00a6148c5c894.
//
// Solidity: event ModuleJudge__ProviderScoreUpdated(address indexed user, address indexed provider, uint256 indexed score)
func (_Contracts *ContractsFilterer) ParseModuleJudgeProviderScoreUpdated(log types.Log) (*ContractsModuleJudgeProviderScoreUpdated, error) {
	event := new(ContractsModuleJudgeProviderScoreUpdated)
	if err := _Contracts.contract.UnpackLog(event, "ModuleJudge__ProviderScoreUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsModuleReservationAppointmentFinishedIterator is returned from FilterModuleReservationAppointmentFinished and is used to iterate over the raw logs and unpacked data for ModuleReservationAppointmentFinished events raised by the Contracts contract.
type ContractsModuleReservationAppointmentFinishedIterator struct {
	Event *ContractsModuleReservationAppointmentFinished // Event containing the contract specifics and raw log

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
func (it *ContractsModuleReservationAppointmentFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsModuleReservationAppointmentFinished)
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
		it.Event = new(ContractsModuleReservationAppointmentFinished)
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
func (it *ContractsModuleReservationAppointmentFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsModuleReservationAppointmentFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsModuleReservationAppointmentFinished represents a ModuleReservationAppointmentFinished event raised by the Contracts contract.
type ContractsModuleReservationAppointmentFinished struct {
	User        common.Address
	Provider    common.Address
	AmountInWei *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterModuleReservationAppointmentFinished is a free log retrieval operation binding the contract event 0x8cb4a902a72849ff0f35342a3c5d94121489b70117c59b42fdeaaea68f802a3e.
//
// Solidity: event ModuleReservation__AppointmentFinished(address indexed user, address indexed provider, uint256 amountInWei)
func (_Contracts *ContractsFilterer) FilterModuleReservationAppointmentFinished(opts *bind.FilterOpts, user []common.Address, provider []common.Address) (*ContractsModuleReservationAppointmentFinishedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ModuleReservation__AppointmentFinished", userRule, providerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsModuleReservationAppointmentFinishedIterator{contract: _Contracts.contract, event: "ModuleReservation__AppointmentFinished", logs: logs, sub: sub}, nil
}

// WatchModuleReservationAppointmentFinished is a free log subscription operation binding the contract event 0x8cb4a902a72849ff0f35342a3c5d94121489b70117c59b42fdeaaea68f802a3e.
//
// Solidity: event ModuleReservation__AppointmentFinished(address indexed user, address indexed provider, uint256 amountInWei)
func (_Contracts *ContractsFilterer) WatchModuleReservationAppointmentFinished(opts *bind.WatchOpts, sink chan<- *ContractsModuleReservationAppointmentFinished, user []common.Address, provider []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ModuleReservation__AppointmentFinished", userRule, providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsModuleReservationAppointmentFinished)
				if err := _Contracts.contract.UnpackLog(event, "ModuleReservation__AppointmentFinished", log); err != nil {
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

// ParseModuleReservationAppointmentFinished is a log parse operation binding the contract event 0x8cb4a902a72849ff0f35342a3c5d94121489b70117c59b42fdeaaea68f802a3e.
//
// Solidity: event ModuleReservation__AppointmentFinished(address indexed user, address indexed provider, uint256 amountInWei)
func (_Contracts *ContractsFilterer) ParseModuleReservationAppointmentFinished(log types.Log) (*ContractsModuleReservationAppointmentFinished, error) {
	event := new(ContractsModuleReservationAppointmentFinished)
	if err := _Contracts.contract.UnpackLog(event, "ModuleReservation__AppointmentFinished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsModuleReservationAppointmentStartedIterator is returned from FilterModuleReservationAppointmentStarted and is used to iterate over the raw logs and unpacked data for ModuleReservationAppointmentStarted events raised by the Contracts contract.
type ContractsModuleReservationAppointmentStartedIterator struct {
	Event *ContractsModuleReservationAppointmentStarted // Event containing the contract specifics and raw log

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
func (it *ContractsModuleReservationAppointmentStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsModuleReservationAppointmentStarted)
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
		it.Event = new(ContractsModuleReservationAppointmentStarted)
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
func (it *ContractsModuleReservationAppointmentStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsModuleReservationAppointmentStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsModuleReservationAppointmentStarted represents a ModuleReservationAppointmentStarted event raised by the Contracts contract.
type ContractsModuleReservationAppointmentStarted struct {
	User        common.Address
	Provider    common.Address
	AmountInWei *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterModuleReservationAppointmentStarted is a free log retrieval operation binding the contract event 0x91251c1521721d1b9719ef4792104f98acef225ec5fe1a0aaa783e74f10aa360.
//
// Solidity: event ModuleReservation__AppointmentStarted(address indexed user, address indexed provider, uint256 amountInWei)
func (_Contracts *ContractsFilterer) FilterModuleReservationAppointmentStarted(opts *bind.FilterOpts, user []common.Address, provider []common.Address) (*ContractsModuleReservationAppointmentStartedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ModuleReservation__AppointmentStarted", userRule, providerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsModuleReservationAppointmentStartedIterator{contract: _Contracts.contract, event: "ModuleReservation__AppointmentStarted", logs: logs, sub: sub}, nil
}

// WatchModuleReservationAppointmentStarted is a free log subscription operation binding the contract event 0x91251c1521721d1b9719ef4792104f98acef225ec5fe1a0aaa783e74f10aa360.
//
// Solidity: event ModuleReservation__AppointmentStarted(address indexed user, address indexed provider, uint256 amountInWei)
func (_Contracts *ContractsFilterer) WatchModuleReservationAppointmentStarted(opts *bind.WatchOpts, sink chan<- *ContractsModuleReservationAppointmentStarted, user []common.Address, provider []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ModuleReservation__AppointmentStarted", userRule, providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsModuleReservationAppointmentStarted)
				if err := _Contracts.contract.UnpackLog(event, "ModuleReservation__AppointmentStarted", log); err != nil {
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

// ParseModuleReservationAppointmentStarted is a log parse operation binding the contract event 0x91251c1521721d1b9719ef4792104f98acef225ec5fe1a0aaa783e74f10aa360.
//
// Solidity: event ModuleReservation__AppointmentStarted(address indexed user, address indexed provider, uint256 amountInWei)
func (_Contracts *ContractsFilterer) ParseModuleReservationAppointmentStarted(log types.Log) (*ContractsModuleReservationAppointmentStarted, error) {
	event := new(ContractsModuleReservationAppointmentStarted)
	if err := _Contracts.contract.UnpackLog(event, "ModuleReservation__AppointmentStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsModuleReservationReservationCanceledIterator is returned from FilterModuleReservationReservationCanceled and is used to iterate over the raw logs and unpacked data for ModuleReservationReservationCanceled events raised by the Contracts contract.
type ContractsModuleReservationReservationCanceledIterator struct {
	Event *ContractsModuleReservationReservationCanceled // Event containing the contract specifics and raw log

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
func (it *ContractsModuleReservationReservationCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsModuleReservationReservationCanceled)
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
		it.Event = new(ContractsModuleReservationReservationCanceled)
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
func (it *ContractsModuleReservationReservationCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsModuleReservationReservationCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsModuleReservationReservationCanceled represents a ModuleReservationReservationCanceled event raised by the Contracts contract.
type ContractsModuleReservationReservationCanceled struct {
	User        common.Address
	Provider    common.Address
	AmountInWei *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterModuleReservationReservationCanceled is a free log retrieval operation binding the contract event 0x2bb25b59fcc7ab82b00a7016299dcadb30e7212cee0661ed5a233f7f8e691729.
//
// Solidity: event ModuleReservation__ReservationCanceled(address indexed user, address indexed provider, uint256 amountInWei)
func (_Contracts *ContractsFilterer) FilterModuleReservationReservationCanceled(opts *bind.FilterOpts, user []common.Address, provider []common.Address) (*ContractsModuleReservationReservationCanceledIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ModuleReservation__ReservationCanceled", userRule, providerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsModuleReservationReservationCanceledIterator{contract: _Contracts.contract, event: "ModuleReservation__ReservationCanceled", logs: logs, sub: sub}, nil
}

// WatchModuleReservationReservationCanceled is a free log subscription operation binding the contract event 0x2bb25b59fcc7ab82b00a7016299dcadb30e7212cee0661ed5a233f7f8e691729.
//
// Solidity: event ModuleReservation__ReservationCanceled(address indexed user, address indexed provider, uint256 amountInWei)
func (_Contracts *ContractsFilterer) WatchModuleReservationReservationCanceled(opts *bind.WatchOpts, sink chan<- *ContractsModuleReservationReservationCanceled, user []common.Address, provider []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ModuleReservation__ReservationCanceled", userRule, providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsModuleReservationReservationCanceled)
				if err := _Contracts.contract.UnpackLog(event, "ModuleReservation__ReservationCanceled", log); err != nil {
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

// ParseModuleReservationReservationCanceled is a log parse operation binding the contract event 0x2bb25b59fcc7ab82b00a7016299dcadb30e7212cee0661ed5a233f7f8e691729.
//
// Solidity: event ModuleReservation__ReservationCanceled(address indexed user, address indexed provider, uint256 amountInWei)
func (_Contracts *ContractsFilterer) ParseModuleReservationReservationCanceled(log types.Log) (*ContractsModuleReservationReservationCanceled, error) {
	event := new(ContractsModuleReservationReservationCanceled)
	if err := _Contracts.contract.UnpackLog(event, "ModuleReservation__ReservationCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsModuleReservationReservationRequestedIterator is returned from FilterModuleReservationReservationRequested and is used to iterate over the raw logs and unpacked data for ModuleReservationReservationRequested events raised by the Contracts contract.
type ContractsModuleReservationReservationRequestedIterator struct {
	Event *ContractsModuleReservationReservationRequested // Event containing the contract specifics and raw log

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
func (it *ContractsModuleReservationReservationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsModuleReservationReservationRequested)
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
		it.Event = new(ContractsModuleReservationReservationRequested)
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
func (it *ContractsModuleReservationReservationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsModuleReservationReservationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsModuleReservationReservationRequested represents a ModuleReservationReservationRequested event raised by the Contracts contract.
type ContractsModuleReservationReservationRequested struct {
	User                  common.Address
	Provider              common.Address
	AppointTimeSinceEpoch *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterModuleReservationReservationRequested is a free log retrieval operation binding the contract event 0x433f65d46e2328545580fd4c54345c5b0b40677086d5c64c3c0aaa0df13f5dea.
//
// Solidity: event ModuleReservation__ReservationRequested(address indexed user, address indexed provider, uint256 appointTimeSinceEpoch)
func (_Contracts *ContractsFilterer) FilterModuleReservationReservationRequested(opts *bind.FilterOpts, user []common.Address, provider []common.Address) (*ContractsModuleReservationReservationRequestedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ModuleReservation__ReservationRequested", userRule, providerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsModuleReservationReservationRequestedIterator{contract: _Contracts.contract, event: "ModuleReservation__ReservationRequested", logs: logs, sub: sub}, nil
}

// WatchModuleReservationReservationRequested is a free log subscription operation binding the contract event 0x433f65d46e2328545580fd4c54345c5b0b40677086d5c64c3c0aaa0df13f5dea.
//
// Solidity: event ModuleReservation__ReservationRequested(address indexed user, address indexed provider, uint256 appointTimeSinceEpoch)
func (_Contracts *ContractsFilterer) WatchModuleReservationReservationRequested(opts *bind.WatchOpts, sink chan<- *ContractsModuleReservationReservationRequested, user []common.Address, provider []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ModuleReservation__ReservationRequested", userRule, providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsModuleReservationReservationRequested)
				if err := _Contracts.contract.UnpackLog(event, "ModuleReservation__ReservationRequested", log); err != nil {
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

// ParseModuleReservationReservationRequested is a log parse operation binding the contract event 0x433f65d46e2328545580fd4c54345c5b0b40677086d5c64c3c0aaa0df13f5dea.
//
// Solidity: event ModuleReservation__ReservationRequested(address indexed user, address indexed provider, uint256 appointTimeSinceEpoch)
func (_Contracts *ContractsFilterer) ParseModuleReservationReservationRequested(log types.Log) (*ContractsModuleReservationReservationRequested, error) {
	event := new(ContractsModuleReservationReservationRequested)
	if err := _Contracts.contract.UnpackLog(event, "ModuleReservation__ReservationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSystemDonationLimitReachedIterator is returned from FilterSystemDonationLimitReached and is used to iterate over the raw logs and unpacked data for SystemDonationLimitReached events raised by the Contracts contract.
type ContractsSystemDonationLimitReachedIterator struct {
	Event *ContractsSystemDonationLimitReached // Event containing the contract specifics and raw log

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
func (it *ContractsSystemDonationLimitReachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSystemDonationLimitReached)
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
		it.Event = new(ContractsSystemDonationLimitReached)
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
func (it *ContractsSystemDonationLimitReachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSystemDonationLimitReachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSystemDonationLimitReached represents a SystemDonationLimitReached event raised by the Contracts contract.
type ContractsSystemDonationLimitReached struct {
	Index       *big.Int
	User        common.Address
	AmountInWei *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSystemDonationLimitReached is a free log retrieval operation binding the contract event 0xcbf2ccb75f1cf126d7bc4d0447f12dcf04cfa44aacea553bf2b506e4ab080297.
//
// Solidity: event System__DonationLimitReached(uint256 indexed index, address indexed user, uint256 amountInWei)
func (_Contracts *ContractsFilterer) FilterSystemDonationLimitReached(opts *bind.FilterOpts, index []*big.Int, user []common.Address) (*ContractsSystemDonationLimitReachedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "System__DonationLimitReached", indexRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSystemDonationLimitReachedIterator{contract: _Contracts.contract, event: "System__DonationLimitReached", logs: logs, sub: sub}, nil
}

// WatchSystemDonationLimitReached is a free log subscription operation binding the contract event 0xcbf2ccb75f1cf126d7bc4d0447f12dcf04cfa44aacea553bf2b506e4ab080297.
//
// Solidity: event System__DonationLimitReached(uint256 indexed index, address indexed user, uint256 amountInWei)
func (_Contracts *ContractsFilterer) WatchSystemDonationLimitReached(opts *bind.WatchOpts, sink chan<- *ContractsSystemDonationLimitReached, index []*big.Int, user []common.Address) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "System__DonationLimitReached", indexRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSystemDonationLimitReached)
				if err := _Contracts.contract.UnpackLog(event, "System__DonationLimitReached", log); err != nil {
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

// ParseSystemDonationLimitReached is a log parse operation binding the contract event 0xcbf2ccb75f1cf126d7bc4d0447f12dcf04cfa44aacea553bf2b506e4ab080297.
//
// Solidity: event System__DonationLimitReached(uint256 indexed index, address indexed user, uint256 amountInWei)
func (_Contracts *ContractsFilterer) ParseSystemDonationLimitReached(log types.Log) (*ContractsSystemDonationLimitReached, error) {
	event := new(ContractsSystemDonationLimitReached)
	if err := _Contracts.contract.UnpackLog(event, "System__DonationLimitReached", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSystemFundWithdrawnIterator is returned from FilterSystemFundWithdrawn and is used to iterate over the raw logs and unpacked data for SystemFundWithdrawn events raised by the Contracts contract.
type ContractsSystemFundWithdrawnIterator struct {
	Event *ContractsSystemFundWithdrawn // Event containing the contract specifics and raw log

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
func (it *ContractsSystemFundWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSystemFundWithdrawn)
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
		it.Event = new(ContractsSystemFundWithdrawn)
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
func (it *ContractsSystemFundWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSystemFundWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSystemFundWithdrawn represents a SystemFundWithdrawn event raised by the Contracts contract.
type ContractsSystemFundWithdrawn struct {
	Index       *big.Int
	User        common.Address
	AmountInWei *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSystemFundWithdrawn is a free log retrieval operation binding the contract event 0x1f3552f3d60339008e06f8c147fcb620103141caf67d0d147f39a44b6b621ee5.
//
// Solidity: event System__FundWithdrawn(uint256 indexed index, address indexed user, uint256 amountInWei)
func (_Contracts *ContractsFilterer) FilterSystemFundWithdrawn(opts *bind.FilterOpts, index []*big.Int, user []common.Address) (*ContractsSystemFundWithdrawnIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "System__FundWithdrawn", indexRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSystemFundWithdrawnIterator{contract: _Contracts.contract, event: "System__FundWithdrawn", logs: logs, sub: sub}, nil
}

// WatchSystemFundWithdrawn is a free log subscription operation binding the contract event 0x1f3552f3d60339008e06f8c147fcb620103141caf67d0d147f39a44b6b621ee5.
//
// Solidity: event System__FundWithdrawn(uint256 indexed index, address indexed user, uint256 amountInWei)
func (_Contracts *ContractsFilterer) WatchSystemFundWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractsSystemFundWithdrawn, index []*big.Int, user []common.Address) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "System__FundWithdrawn", indexRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSystemFundWithdrawn)
				if err := _Contracts.contract.UnpackLog(event, "System__FundWithdrawn", log); err != nil {
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

// ParseSystemFundWithdrawn is a log parse operation binding the contract event 0x1f3552f3d60339008e06f8c147fcb620103141caf67d0d147f39a44b6b621ee5.
//
// Solidity: event System__FundWithdrawn(uint256 indexed index, address indexed user, uint256 amountInWei)
func (_Contracts *ContractsFilterer) ParseSystemFundWithdrawn(log types.Log) (*ContractsSystemFundWithdrawn, error) {
	event := new(ContractsSystemFundWithdrawn)
	if err := _Contracts.contract.UnpackLog(event, "System__FundWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSystemNewDonationIterator is returned from FilterSystemNewDonation and is used to iterate over the raw logs and unpacked data for SystemNewDonation events raised by the Contracts contract.
type ContractsSystemNewDonationIterator struct {
	Event *ContractsSystemNewDonation // Event containing the contract specifics and raw log

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
func (it *ContractsSystemNewDonationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSystemNewDonation)
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
		it.Event = new(ContractsSystemNewDonation)
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
func (it *ContractsSystemNewDonationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSystemNewDonationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSystemNewDonation represents a SystemNewDonation event raised by the Contracts contract.
type ContractsSystemNewDonation struct {
	Index       *big.Int
	User        common.Address
	Sender      common.Address
	AmountInWei *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSystemNewDonation is a free log retrieval operation binding the contract event 0x44d7f8ea51d29b6caacaeb59dc64025092fc08298718864007d5633bf99ef2e6.
//
// Solidity: event System__NewDonation(uint256 indexed index, address user, address indexed sender, uint256 amountInWei)
func (_Contracts *ContractsFilterer) FilterSystemNewDonation(opts *bind.FilterOpts, index []*big.Int, sender []common.Address) (*ContractsSystemNewDonationIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "System__NewDonation", indexRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSystemNewDonationIterator{contract: _Contracts.contract, event: "System__NewDonation", logs: logs, sub: sub}, nil
}

// WatchSystemNewDonation is a free log subscription operation binding the contract event 0x44d7f8ea51d29b6caacaeb59dc64025092fc08298718864007d5633bf99ef2e6.
//
// Solidity: event System__NewDonation(uint256 indexed index, address user, address indexed sender, uint256 amountInWei)
func (_Contracts *ContractsFilterer) WatchSystemNewDonation(opts *bind.WatchOpts, sink chan<- *ContractsSystemNewDonation, index []*big.Int, sender []common.Address) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "System__NewDonation", indexRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSystemNewDonation)
				if err := _Contracts.contract.UnpackLog(event, "System__NewDonation", log); err != nil {
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

// ParseSystemNewDonation is a log parse operation binding the contract event 0x44d7f8ea51d29b6caacaeb59dc64025092fc08298718864007d5633bf99ef2e6.
//
// Solidity: event System__NewDonation(uint256 indexed index, address user, address indexed sender, uint256 amountInWei)
func (_Contracts *ContractsFilterer) ParseSystemNewDonation(log types.Log) (*ContractsSystemNewDonation, error) {
	event := new(ContractsSystemNewDonation)
	if err := _Contracts.contract.UnpackLog(event, "System__NewDonation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSystemNewFundRegisteredIterator is returned from FilterSystemNewFundRegistered and is used to iterate over the raw logs and unpacked data for SystemNewFundRegistered events raised by the Contracts contract.
type ContractsSystemNewFundRegisteredIterator struct {
	Event *ContractsSystemNewFundRegistered // Event containing the contract specifics and raw log

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
func (it *ContractsSystemNewFundRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSystemNewFundRegistered)
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
		it.Event = new(ContractsSystemNewFundRegistered)
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
func (it *ContractsSystemNewFundRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSystemNewFundRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSystemNewFundRegistered represents a SystemNewFundRegistered event raised by the Contracts contract.
type ContractsSystemNewFundRegistered struct {
	FundInfoIndex *big.Int
	User          common.Address
	AmountInUsd   *big.Int
	RoundId       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSystemNewFundRegistered is a free log retrieval operation binding the contract event 0xcf56491503d63fcce311d3085dbb27c2782c4e9ca64f9dfdb9d25c7f4ccedcb0.
//
// Solidity: event System__NewFundRegistered(uint256 indexed fundInfoIndex, address user, uint256 amountInUsd, uint80 indexed roundId)
func (_Contracts *ContractsFilterer) FilterSystemNewFundRegistered(opts *bind.FilterOpts, fundInfoIndex []*big.Int, roundId []*big.Int) (*ContractsSystemNewFundRegisteredIterator, error) {

	var fundInfoIndexRule []interface{}
	for _, fundInfoIndexItem := range fundInfoIndex {
		fundInfoIndexRule = append(fundInfoIndexRule, fundInfoIndexItem)
	}

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "System__NewFundRegistered", fundInfoIndexRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSystemNewFundRegisteredIterator{contract: _Contracts.contract, event: "System__NewFundRegistered", logs: logs, sub: sub}, nil
}

// WatchSystemNewFundRegistered is a free log subscription operation binding the contract event 0xcf56491503d63fcce311d3085dbb27c2782c4e9ca64f9dfdb9d25c7f4ccedcb0.
//
// Solidity: event System__NewFundRegistered(uint256 indexed fundInfoIndex, address user, uint256 amountInUsd, uint80 indexed roundId)
func (_Contracts *ContractsFilterer) WatchSystemNewFundRegistered(opts *bind.WatchOpts, sink chan<- *ContractsSystemNewFundRegistered, fundInfoIndex []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var fundInfoIndexRule []interface{}
	for _, fundInfoIndexItem := range fundInfoIndex {
		fundInfoIndexRule = append(fundInfoIndexRule, fundInfoIndexItem)
	}

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "System__NewFundRegistered", fundInfoIndexRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSystemNewFundRegistered)
				if err := _Contracts.contract.UnpackLog(event, "System__NewFundRegistered", log); err != nil {
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

// ParseSystemNewFundRegistered is a log parse operation binding the contract event 0xcf56491503d63fcce311d3085dbb27c2782c4e9ca64f9dfdb9d25c7f4ccedcb0.
//
// Solidity: event System__NewFundRegistered(uint256 indexed fundInfoIndex, address user, uint256 amountInUsd, uint80 indexed roundId)
func (_Contracts *ContractsFilterer) ParseSystemNewFundRegistered(log types.Log) (*ContractsSystemNewFundRegistered, error) {
	event := new(ContractsSystemNewFundRegistered)
	if err := _Contracts.contract.UnpackLog(event, "System__NewFundRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSystemNewUserAddedIterator is returned from FilterSystemNewUserAdded and is used to iterate over the raw logs and unpacked data for SystemNewUserAdded events raised by the Contracts contract.
type ContractsSystemNewUserAddedIterator struct {
	Event *ContractsSystemNewUserAdded // Event containing the contract specifics and raw log

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
func (it *ContractsSystemNewUserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSystemNewUserAdded)
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
		it.Event = new(ContractsSystemNewUserAdded)
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
func (it *ContractsSystemNewUserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSystemNewUserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSystemNewUserAdded represents a SystemNewUserAdded event raised by the Contracts contract.
type ContractsSystemNewUserAdded struct {
	User     common.Address
	RoleType uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSystemNewUserAdded is a free log retrieval operation binding the contract event 0xd34811e3a3e8e6940027b2cff6e361bbdf8d93ae261fcd218aae71428c603c3c.
//
// Solidity: event System__NewUserAdded(address indexed user, uint8 indexed roleType)
func (_Contracts *ContractsFilterer) FilterSystemNewUserAdded(opts *bind.FilterOpts, user []common.Address, roleType []uint8) (*ContractsSystemNewUserAddedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var roleTypeRule []interface{}
	for _, roleTypeItem := range roleType {
		roleTypeRule = append(roleTypeRule, roleTypeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "System__NewUserAdded", userRule, roleTypeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSystemNewUserAddedIterator{contract: _Contracts.contract, event: "System__NewUserAdded", logs: logs, sub: sub}, nil
}

// WatchSystemNewUserAdded is a free log subscription operation binding the contract event 0xd34811e3a3e8e6940027b2cff6e361bbdf8d93ae261fcd218aae71428c603c3c.
//
// Solidity: event System__NewUserAdded(address indexed user, uint8 indexed roleType)
func (_Contracts *ContractsFilterer) WatchSystemNewUserAdded(opts *bind.WatchOpts, sink chan<- *ContractsSystemNewUserAdded, user []common.Address, roleType []uint8) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var roleTypeRule []interface{}
	for _, roleTypeItem := range roleType {
		roleTypeRule = append(roleTypeRule, roleTypeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "System__NewUserAdded", userRule, roleTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSystemNewUserAdded)
				if err := _Contracts.contract.UnpackLog(event, "System__NewUserAdded", log); err != nil {
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

// ParseSystemNewUserAdded is a log parse operation binding the contract event 0xd34811e3a3e8e6940027b2cff6e361bbdf8d93ae261fcd218aae71428c603c3c.
//
// Solidity: event System__NewUserAdded(address indexed user, uint8 indexed roleType)
func (_Contracts *ContractsFilterer) ParseSystemNewUserAdded(log types.Log) (*ContractsSystemNewUserAdded, error) {
	event := new(ContractsSystemNewUserAdded)
	if err := _Contracts.contract.UnpackLog(event, "System__NewUserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSystemNewUserContractAddedIterator is returned from FilterSystemNewUserContractAdded and is used to iterate over the raw logs and unpacked data for SystemNewUserContractAdded events raised by the Contracts contract.
type ContractsSystemNewUserContractAddedIterator struct {
	Event *ContractsSystemNewUserContractAdded // Event containing the contract specifics and raw log

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
func (it *ContractsSystemNewUserContractAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSystemNewUserContractAdded)
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
		it.Event = new(ContractsSystemNewUserContractAdded)
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
func (it *ContractsSystemNewUserContractAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSystemNewUserContractAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSystemNewUserContractAdded represents a SystemNewUserContractAdded event raised by the Contracts contract.
type ContractsSystemNewUserContractAdded struct {
	User            common.Address
	ContractAddress common.Address
	RoleType        uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSystemNewUserContractAdded is a free log retrieval operation binding the contract event 0xe34179d898002d6fca95072836c15365846404fe8aaf36c68df01504f4af3a95.
//
// Solidity: event System__NewUserContractAdded(address indexed user, address indexed contractAddress, uint8 indexed roleType)
func (_Contracts *ContractsFilterer) FilterSystemNewUserContractAdded(opts *bind.FilterOpts, user []common.Address, contractAddress []common.Address, roleType []uint8) (*ContractsSystemNewUserContractAddedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var roleTypeRule []interface{}
	for _, roleTypeItem := range roleType {
		roleTypeRule = append(roleTypeRule, roleTypeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "System__NewUserContractAdded", userRule, contractAddressRule, roleTypeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSystemNewUserContractAddedIterator{contract: _Contracts.contract, event: "System__NewUserContractAdded", logs: logs, sub: sub}, nil
}

// WatchSystemNewUserContractAdded is a free log subscription operation binding the contract event 0xe34179d898002d6fca95072836c15365846404fe8aaf36c68df01504f4af3a95.
//
// Solidity: event System__NewUserContractAdded(address indexed user, address indexed contractAddress, uint8 indexed roleType)
func (_Contracts *ContractsFilterer) WatchSystemNewUserContractAdded(opts *bind.WatchOpts, sink chan<- *ContractsSystemNewUserContractAdded, user []common.Address, contractAddress []common.Address, roleType []uint8) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var roleTypeRule []interface{}
	for _, roleTypeItem := range roleType {
		roleTypeRule = append(roleTypeRule, roleTypeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "System__NewUserContractAdded", userRule, contractAddressRule, roleTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSystemNewUserContractAdded)
				if err := _Contracts.contract.UnpackLog(event, "System__NewUserContractAdded", log); err != nil {
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

// ParseSystemNewUserContractAdded is a log parse operation binding the contract event 0xe34179d898002d6fca95072836c15365846404fe8aaf36c68df01504f4af3a95.
//
// Solidity: event System__NewUserContractAdded(address indexed user, address indexed contractAddress, uint8 indexed roleType)
func (_Contracts *ContractsFilterer) ParseSystemNewUserContractAdded(log types.Log) (*ContractsSystemNewUserContractAdded, error) {
	event := new(ContractsSystemNewUserContractAdded)
	if err := _Contracts.contract.UnpackLog(event, "System__NewUserContractAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
