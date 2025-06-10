// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reservaId\",\"type\":\"uint256\"}],\"name\":\"cancelarReserva\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contadorReservas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usuario\",\"type\":\"address\"}],\"name\":\"getReservasDoUsuario\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"pontos\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pontoId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inicio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fim\",\"type\":\"uint256\"}],\"name\":\"reservarPonto\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reservas\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"usuario\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pontoId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inicio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fim\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ativa\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reservasPorUsuario\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506112bc8061005b5f395ff3fe608060405234801561000f575f5ffd5b506004361061007b575f3560e01c80639c33e0af116100595780639c33e0af146100d55780639f43685314610105578063aa34dd4214610139578063ad852310146101575761007b565b80632a05ba701461007f5780637b4ae18c1461009b5780638da5cb5b146100b7575b5f5ffd5b61009960048036038101906100949190610c35565b610188565b005b6100b560048036038101906100b09190610c60565b610391565b005b6100bf610839565b6040516100cc9190610cef565b60405180910390f35b6100ef60048036038101906100ea9190610d32565b61085d565b6040516100fc9190610d7f565b60405180910390f35b61011f600480360381019061011a9190610c35565b610888565b604051610130959493929190610db2565b60405180910390f35b6101416108e5565b60405161014e9190610d7f565b60405180910390f35b610171600480360381019061016c9190610e03565b6108eb565b60405161017f929190610ee5565b60405180910390f35b5f60015f8381526020019081526020015f2090503373ffffffffffffffffffffffffffffffffffffffff16815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461022c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161022390610f74565b60405180910390fd5b806004015f9054906101000a900460ff1661027c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027390610fdc565b60405180910390fd5b5f816004015f6101000a81548160ff0219169083151502179055505f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090505f5f90505b818054905081101561038b57838282815481106102fc576102fb610ffa565b5b905f5260205f2001540361037e57816001838054905061031c9190611054565b8154811061032d5761032c610ffa565b5b905f5260205f20015482828154811061034957610348610ffa565b5b905f5260205f2001819055508180548061036657610365611087565b5b600190038181905f5260205f20015f9055905561038b565b80806001019150506102dc565b50505050565b8082106103d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ca906110fe565b60405180910390fd5b5f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2080548060200260200160405190810160405280929190818152602001828054801561045957602002820191905f5260205f20905b815481526020019060010190808311610445575b505050505090505f5f90505b81518110156105a2575f60015f84848151811061048557610484610ffa565b5b602002602001015181526020019081526020015f206040518060a00160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152602001600482015f9054906101000a900460ff1615151515815250509050806080015180156105545750610553816040015182606001518787610be5565b5b15610594576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058b9061118c565b60405180910390fd5b508080600101915050610465565b505f600190505b60035481116106e1575f60015f8381526020019081526020015f206040518060a00160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152602001600482015f9054906101000a900460ff1615151515815250509050806080015180156106715750858160200151145b801561068d575061068c816040015182606001518787610be5565b5b156106cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106c4906111f4565b60405180910390fd5b5080806106d990611212565b9150506105a9565b5060035f8154809291906106f490611212565b91905055506040518060a001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020018381526020016001151581525060015f60035481526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155606082015181600301556080820151816004015f6101000a81548160ff02191690831515021790555090505060025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600354908060018154018082558091505060019003905f5260205f20015f909190919091505550505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6002602052815f5260405f208181548110610876575f80fd5b905f5260205f20015f91509150505481565b6001602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015490806002015490806003015490806004015f9054906101000a900460ff16905085565b60035481565b6060805f60025f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2080548060200260200160405190810160405280929190818152602001828054801561097457602002820191905f5260205f20905b815481526020019060010190808311610960575b505050505090505f5f90505f5f90505b82518110156109e95760015f8483815181106109a3576109a2610ffa565b5b602002602001015181526020019081526020015f206004015f9054906101000a900460ff16156109dc5781806109d890611212565b9250505b8080600101915050610984565b508067ffffffffffffffff811115610a0457610a03611259565b5b604051908082528060200260200182016040528015610a325781602001602082028036833780820191505090505b5093508067ffffffffffffffff811115610a4f57610a4e611259565b5b604051908082528060200260200182016040528015610a7d5781602001602082028036833780820191505090505b5092505f5f90505f5f90505b8351811015610bdc575f60015f868481518110610aa957610aa8610ffa565b5b602002602001015181526020019081526020015f206040518060a00160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820154815260200160038201548152602001600482015f9054906101000a900460ff1615151515815250509050806080015115610bce57848281518110610b7457610b73610ffa565b5b6020026020010151878481518110610b8f57610b8e610ffa565b5b6020026020010181815250508060200151868481518110610bb357610bb2610ffa565b5b6020026020010181815250508280610bca90611212565b9350505b508080600101915050610a89565b50505050915091565b5f8185108015610bf457508383105b9050949350505050565b5f5ffd5b5f819050919050565b610c1481610c02565b8114610c1e575f5ffd5b50565b5f81359050610c2f81610c0b565b92915050565b5f60208284031215610c4a57610c49610bfe565b5b5f610c5784828501610c21565b91505092915050565b5f5f5f60608486031215610c7757610c76610bfe565b5b5f610c8486828701610c21565b9350506020610c9586828701610c21565b9250506040610ca686828701610c21565b9150509250925092565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610cd982610cb0565b9050919050565b610ce981610ccf565b82525050565b5f602082019050610d025f830184610ce0565b92915050565b610d1181610ccf565b8114610d1b575f5ffd5b50565b5f81359050610d2c81610d08565b92915050565b5f5f60408385031215610d4857610d47610bfe565b5b5f610d5585828601610d1e565b9250506020610d6685828601610c21565b9150509250929050565b610d7981610c02565b82525050565b5f602082019050610d925f830184610d70565b92915050565b5f8115159050919050565b610dac81610d98565b82525050565b5f60a082019050610dc55f830188610ce0565b610dd26020830187610d70565b610ddf6040830186610d70565b610dec6060830185610d70565b610df96080830184610da3565b9695505050505050565b5f60208284031215610e1857610e17610bfe565b5b5f610e2584828501610d1e565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610e6081610c02565b82525050565b5f610e718383610e57565b60208301905092915050565b5f602082019050919050565b5f610e9382610e2e565b610e9d8185610e38565b9350610ea883610e48565b805f5b83811015610ed8578151610ebf8882610e66565b9750610eca83610e7d565b925050600181019050610eab565b5085935050505092915050565b5f6040820190508181035f830152610efd8185610e89565b90508181036020830152610f118184610e89565b90509392505050565b5f82825260208201905092915050565b7f4e616f206175746f72697a61646f0000000000000000000000000000000000005f82015250565b5f610f5e600e83610f1a565b9150610f6982610f2a565b602082019050919050565b5f6020820190508181035f830152610f8b81610f52565b9050919050565b7f52657365727661206a612063616e63656c6164610000000000000000000000005f82015250565b5f610fc6601483610f1a565b9150610fd182610f92565b602082019050919050565b5f6020820190508181035f830152610ff381610fba565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61105e82610c02565b915061106983610c02565b925082820390508181111561108157611080611027565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f496e74657276616c6f20696e76616c69646f00000000000000000000000000005f82015250565b5f6110e8601283610f1a565b91506110f3826110b4565b602082019050919050565b5f6020820190508181035f830152611115816110dc565b9050919050565b7f566f6365206a6120706f737375692072657365727661206e6573736520686f725f8201527f6172696f00000000000000000000000000000000000000000000000000000000602082015250565b5f611176602483610f1a565b91506111818261111c565b604082019050919050565b5f6020820190508181035f8301526111a38161116a565b9050919050565b7f506f6e746f206a612072657365727661646f206e6573736520686f726172696f5f82015250565b5f6111de602083610f1a565b91506111e9826111aa565b602082019050919050565b5f6020820190508181035f83015261120b816111d2565b9050919050565b5f61121c82610c02565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361124e5761124d611027565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffdfea26469706673582212207c894513f085bcc8d350ac554f30b641036cd9dacc90a2fef52a8ee0726a568564736f6c634300081e0033",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// MainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MainMetaData.Bin instead.
var MainBin = MainMetaData.Bin

// DeployMain deploys a new Ethereum contract, binding an instance of Main to it.
func DeployMain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Main, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// ContadorReservas is a free data retrieval call binding the contract method 0xaa34dd42.
//
// Solidity: function contadorReservas() view returns(uint256)
func (_Main *MainCaller) ContadorReservas(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "contadorReservas")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContadorReservas is a free data retrieval call binding the contract method 0xaa34dd42.
//
// Solidity: function contadorReservas() view returns(uint256)
func (_Main *MainSession) ContadorReservas() (*big.Int, error) {
	return _Main.Contract.ContadorReservas(&_Main.CallOpts)
}

// ContadorReservas is a free data retrieval call binding the contract method 0xaa34dd42.
//
// Solidity: function contadorReservas() view returns(uint256)
func (_Main *MainCallerSession) ContadorReservas() (*big.Int, error) {
	return _Main.Contract.ContadorReservas(&_Main.CallOpts)
}

// GetReservasDoUsuario is a free data retrieval call binding the contract method 0xad852310.
//
// Solidity: function getReservasDoUsuario(address usuario) view returns(uint256[] ids, uint256[] pontos)
func (_Main *MainCaller) GetReservasDoUsuario(opts *bind.CallOpts, usuario common.Address) (struct {
	Ids    []*big.Int
	Pontos []*big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getReservasDoUsuario", usuario)

	outstruct := new(struct {
		Ids    []*big.Int
		Pontos []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ids = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Pontos = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetReservasDoUsuario is a free data retrieval call binding the contract method 0xad852310.
//
// Solidity: function getReservasDoUsuario(address usuario) view returns(uint256[] ids, uint256[] pontos)
func (_Main *MainSession) GetReservasDoUsuario(usuario common.Address) (struct {
	Ids    []*big.Int
	Pontos []*big.Int
}, error) {
	return _Main.Contract.GetReservasDoUsuario(&_Main.CallOpts, usuario)
}

// GetReservasDoUsuario is a free data retrieval call binding the contract method 0xad852310.
//
// Solidity: function getReservasDoUsuario(address usuario) view returns(uint256[] ids, uint256[] pontos)
func (_Main *MainCallerSession) GetReservasDoUsuario(usuario common.Address) (struct {
	Ids    []*big.Int
	Pontos []*big.Int
}, error) {
	return _Main.Contract.GetReservasDoUsuario(&_Main.CallOpts, usuario)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainSession) Owner() (common.Address, error) {
	return _Main.Contract.Owner(&_Main.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Main *MainCallerSession) Owner() (common.Address, error) {
	return _Main.Contract.Owner(&_Main.CallOpts)
}

// Reservas is a free data retrieval call binding the contract method 0x9f436853.
//
// Solidity: function reservas(uint256 ) view returns(address usuario, uint256 pontoId, uint256 inicio, uint256 fim, bool ativa)
func (_Main *MainCaller) Reservas(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Usuario common.Address
	PontoId *big.Int
	Inicio  *big.Int
	Fim     *big.Int
	Ativa   bool
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "reservas", arg0)

	outstruct := new(struct {
		Usuario common.Address
		PontoId *big.Int
		Inicio  *big.Int
		Fim     *big.Int
		Ativa   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Usuario = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PontoId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Inicio = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Fim = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Ativa = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Reservas is a free data retrieval call binding the contract method 0x9f436853.
//
// Solidity: function reservas(uint256 ) view returns(address usuario, uint256 pontoId, uint256 inicio, uint256 fim, bool ativa)
func (_Main *MainSession) Reservas(arg0 *big.Int) (struct {
	Usuario common.Address
	PontoId *big.Int
	Inicio  *big.Int
	Fim     *big.Int
	Ativa   bool
}, error) {
	return _Main.Contract.Reservas(&_Main.CallOpts, arg0)
}

// Reservas is a free data retrieval call binding the contract method 0x9f436853.
//
// Solidity: function reservas(uint256 ) view returns(address usuario, uint256 pontoId, uint256 inicio, uint256 fim, bool ativa)
func (_Main *MainCallerSession) Reservas(arg0 *big.Int) (struct {
	Usuario common.Address
	PontoId *big.Int
	Inicio  *big.Int
	Fim     *big.Int
	Ativa   bool
}, error) {
	return _Main.Contract.Reservas(&_Main.CallOpts, arg0)
}

// ReservasPorUsuario is a free data retrieval call binding the contract method 0x9c33e0af.
//
// Solidity: function reservasPorUsuario(address , uint256 ) view returns(uint256)
func (_Main *MainCaller) ReservasPorUsuario(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "reservasPorUsuario", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservasPorUsuario is a free data retrieval call binding the contract method 0x9c33e0af.
//
// Solidity: function reservasPorUsuario(address , uint256 ) view returns(uint256)
func (_Main *MainSession) ReservasPorUsuario(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Main.Contract.ReservasPorUsuario(&_Main.CallOpts, arg0, arg1)
}

// ReservasPorUsuario is a free data retrieval call binding the contract method 0x9c33e0af.
//
// Solidity: function reservasPorUsuario(address , uint256 ) view returns(uint256)
func (_Main *MainCallerSession) ReservasPorUsuario(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Main.Contract.ReservasPorUsuario(&_Main.CallOpts, arg0, arg1)
}

// CancelarReserva is a paid mutator transaction binding the contract method 0x2a05ba70.
//
// Solidity: function cancelarReserva(uint256 reservaId) returns()
func (_Main *MainTransactor) CancelarReserva(opts *bind.TransactOpts, reservaId *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "cancelarReserva", reservaId)
}

// CancelarReserva is a paid mutator transaction binding the contract method 0x2a05ba70.
//
// Solidity: function cancelarReserva(uint256 reservaId) returns()
func (_Main *MainSession) CancelarReserva(reservaId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.CancelarReserva(&_Main.TransactOpts, reservaId)
}

// CancelarReserva is a paid mutator transaction binding the contract method 0x2a05ba70.
//
// Solidity: function cancelarReserva(uint256 reservaId) returns()
func (_Main *MainTransactorSession) CancelarReserva(reservaId *big.Int) (*types.Transaction, error) {
	return _Main.Contract.CancelarReserva(&_Main.TransactOpts, reservaId)
}

// ReservarPonto is a paid mutator transaction binding the contract method 0x7b4ae18c.
//
// Solidity: function reservarPonto(uint256 pontoId, uint256 inicio, uint256 fim) returns()
func (_Main *MainTransactor) ReservarPonto(opts *bind.TransactOpts, pontoId *big.Int, inicio *big.Int, fim *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "reservarPonto", pontoId, inicio, fim)
}

// ReservarPonto is a paid mutator transaction binding the contract method 0x7b4ae18c.
//
// Solidity: function reservarPonto(uint256 pontoId, uint256 inicio, uint256 fim) returns()
func (_Main *MainSession) ReservarPonto(pontoId *big.Int, inicio *big.Int, fim *big.Int) (*types.Transaction, error) {
	return _Main.Contract.ReservarPonto(&_Main.TransactOpts, pontoId, inicio, fim)
}

// ReservarPonto is a paid mutator transaction binding the contract method 0x7b4ae18c.
//
// Solidity: function reservarPonto(uint256 pontoId, uint256 inicio, uint256 fim) returns()
func (_Main *MainTransactorSession) ReservarPonto(pontoId *big.Int, inicio *big.Int, fim *big.Int) (*types.Transaction, error) {
	return _Main.Contract.ReservarPonto(&_Main.TransactOpts, pontoId, inicio, fim)
}
