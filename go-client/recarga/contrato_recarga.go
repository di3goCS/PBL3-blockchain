// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package recarga

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

// RecargaMetaData contains all meta data concerning the Recarga contract.
var RecargaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"contadorRecargas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usuario\",\"type\":\"address\"}],\"name\":\"getRecargasDoUsuario\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"valores\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"pagos\",\"type\":\"bool[]\"},{\"internalType\":\"uint256[]\",\"name\":\"timestamps\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"recargaId\",\"type\":\"uint256\"}],\"name\":\"pagarRecarga\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recargas\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"usuario\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pago\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recargasPorUsuario\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usuario\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valor\",\"type\":\"uint256\"}],\"name\":\"registrarRecarga\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610d438061005b5f395ff3fe60806040526004361061006f575f3560e01c8063b5cd80331161004d578063b5cd803314610101578063c842bccb14610140578063deb9122d1461015c578063f7e04e4d1461019b5761006f565b806370a0db5e14610073578063733a5173146100af5780638da5cb5b146100d7575b5f5ffd5b34801561007e575f5ffd5b5061009960048036038101906100949190610851565b6101c5565b6040516100a6919061089e565b60405180910390f35b3480156100ba575f5ffd5b506100d560048036038101906100d09190610851565b6101f0565b005b3480156100e2575f5ffd5b506100eb610334565b6040516100f891906108c6565b60405180910390f35b34801561010c575f5ffd5b50610127600480360381019061012291906108df565b610358565b6040516101379493929190610924565b60405180910390f35b61015a600480360381019061015591906108df565b6103af565b005b348015610167575f5ffd5b50610182600480360381019061017d9190610967565b6104df565b6040516101929493929190610b00565b60405180910390f35b3480156101a6575f5ffd5b506101af6107ba565b6040516101bc919061089e565b60405180910390f35b6002602052815f5260405f2081815481106101de575f80fd5b905f5260205f20015f91509150505481565b60035f81548092919061020290610b8c565b919050555060405180608001604052808373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020014281526020015f151581525060015f60035481526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555090505060025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600354908060018154018082558091505060019003905f5260205f20015f90919091909150555050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6001602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015490806002015490806003015f9054906101000a900460ff16905084565b5f60015f8381526020019081526020015f209050806003015f9054906101000a900460ff1615610414576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040b90610c2d565b60405180910390fd5b806001015434101561045b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045290610c95565b60405180910390fd5b6001816003015f6101000a81548160ff0219169083151502179055505f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc3490811502906040515f60405180830381858888f193505050501580156104da573d5f5f3e3d5ffd5b505050565b6060806060805f60025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2080548060200260200160405190810160405280929190818152602001828054801561056b57602002820191905f5260205f20905b815481526020019060010190808311610557575b505050505090505f815167ffffffffffffffff81111561058e5761058d610cb3565b5b6040519080825280602002602001820160405280156105bc5781602001602082028036833780820191505090505b5090505f825167ffffffffffffffff8111156105db576105da610cb3565b5b6040519080825280602002602001820160405280156106095781602001602082028036833780820191505090505b5090505f835167ffffffffffffffff81111561062857610627610cb3565b5b6040519080825280602002602001820160405280156106565781602001602082028036833780820191505090505b5090505f5f90505b84518110156107a2575f60015f87848151811061067e5761067d610ce0565b5b602002602001015181526020019081526020015f206040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201548152602001600382015f9054906101000a900460ff1615151515815250509050806020015185838151811061073a57610739610ce0565b5b602002602001018181525050806060015184838151811061075e5761075d610ce0565b5b602002602001019015159081151581525050806040015183838151811061078857610787610ce0565b5b60200260200101818152505050808060010191505061065e565b50838383839750975097509750505050509193509193565b60035481565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6107ed826107c4565b9050919050565b6107fd816107e3565b8114610807575f5ffd5b50565b5f81359050610818816107f4565b92915050565b5f819050919050565b6108308161081e565b811461083a575f5ffd5b50565b5f8135905061084b81610827565b92915050565b5f5f60408385031215610867576108666107c0565b5b5f6108748582860161080a565b92505060206108858582860161083d565b9150509250929050565b6108988161081e565b82525050565b5f6020820190506108b15f83018461088f565b92915050565b6108c0816107e3565b82525050565b5f6020820190506108d95f8301846108b7565b92915050565b5f602082840312156108f4576108f36107c0565b5b5f6109018482850161083d565b91505092915050565b5f8115159050919050565b61091e8161090a565b82525050565b5f6080820190506109375f8301876108b7565b610944602083018661088f565b610951604083018561088f565b61095e6060830184610915565b95945050505050565b5f6020828403121561097c5761097b6107c0565b5b5f6109898482850161080a565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6109c48161081e565b82525050565b5f6109d583836109bb565b60208301905092915050565b5f602082019050919050565b5f6109f782610992565b610a01818561099c565b9350610a0c836109ac565b805f5b83811015610a3c578151610a2388826109ca565b9750610a2e836109e1565b925050600181019050610a0f565b5085935050505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610a7b8161090a565b82525050565b5f610a8c8383610a72565b60208301905092915050565b5f602082019050919050565b5f610aae82610a49565b610ab88185610a53565b9350610ac383610a63565b805f5b83811015610af3578151610ada8882610a81565b9750610ae583610a98565b925050600181019050610ac6565b5085935050505092915050565b5f6080820190508181035f830152610b1881876109ed565b90508181036020830152610b2c81866109ed565b90508181036040830152610b408185610aa4565b90508181036060830152610b5481846109ed565b905095945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610b968261081e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610bc857610bc7610b5f565b5b600182019050919050565b5f82825260208201905092915050565b7f52656361726761206a61207061676100000000000000000000000000000000005f82015250565b5f610c17600f83610bd3565b9150610c2282610be3565b602082019050919050565b5f6020820190508181035f830152610c4481610c0b565b9050919050565b7f56616c6f7220696e737566696369656e746500000000000000000000000000005f82015250565b5f610c7f601283610bd3565b9150610c8a82610c4b565b602082019050919050565b5f6020820190508181035f830152610cac81610c73565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea2646970667358221220beace602c2d06793a74fc01c56c5675ebb0784aa230074005184c2da8d8e459864736f6c634300081e0033",
}

// RecargaABI is the input ABI used to generate the binding from.
// Deprecated: Use RecargaMetaData.ABI instead.
var RecargaABI = RecargaMetaData.ABI

// RecargaBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RecargaMetaData.Bin instead.
var RecargaBin = RecargaMetaData.Bin

// DeployRecarga deploys a new Ethereum contract, binding an instance of Recarga to it.
func DeployRecarga(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Recarga, error) {
	parsed, err := RecargaMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RecargaBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Recarga{RecargaCaller: RecargaCaller{contract: contract}, RecargaTransactor: RecargaTransactor{contract: contract}, RecargaFilterer: RecargaFilterer{contract: contract}}, nil
}

// Recarga is an auto generated Go binding around an Ethereum contract.
type Recarga struct {
	RecargaCaller     // Read-only binding to the contract
	RecargaTransactor // Write-only binding to the contract
	RecargaFilterer   // Log filterer for contract events
}

// RecargaCaller is an auto generated read-only Go binding around an Ethereum contract.
type RecargaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecargaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RecargaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecargaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RecargaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecargaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RecargaSession struct {
	Contract     *Recarga          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RecargaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RecargaCallerSession struct {
	Contract *RecargaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RecargaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RecargaTransactorSession struct {
	Contract     *RecargaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RecargaRaw is an auto generated low-level Go binding around an Ethereum contract.
type RecargaRaw struct {
	Contract *Recarga // Generic contract binding to access the raw methods on
}

// RecargaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RecargaCallerRaw struct {
	Contract *RecargaCaller // Generic read-only contract binding to access the raw methods on
}

// RecargaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RecargaTransactorRaw struct {
	Contract *RecargaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRecarga creates a new instance of Recarga, bound to a specific deployed contract.
func NewRecarga(address common.Address, backend bind.ContractBackend) (*Recarga, error) {
	contract, err := bindRecarga(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Recarga{RecargaCaller: RecargaCaller{contract: contract}, RecargaTransactor: RecargaTransactor{contract: contract}, RecargaFilterer: RecargaFilterer{contract: contract}}, nil
}

// NewRecargaCaller creates a new read-only instance of Recarga, bound to a specific deployed contract.
func NewRecargaCaller(address common.Address, caller bind.ContractCaller) (*RecargaCaller, error) {
	contract, err := bindRecarga(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecargaCaller{contract: contract}, nil
}

// NewRecargaTransactor creates a new write-only instance of Recarga, bound to a specific deployed contract.
func NewRecargaTransactor(address common.Address, transactor bind.ContractTransactor) (*RecargaTransactor, error) {
	contract, err := bindRecarga(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecargaTransactor{contract: contract}, nil
}

// NewRecargaFilterer creates a new log filterer instance of Recarga, bound to a specific deployed contract.
func NewRecargaFilterer(address common.Address, filterer bind.ContractFilterer) (*RecargaFilterer, error) {
	contract, err := bindRecarga(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecargaFilterer{contract: contract}, nil
}

// bindRecarga binds a generic wrapper to an already deployed contract.
func bindRecarga(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RecargaMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Recarga *RecargaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Recarga.Contract.RecargaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Recarga *RecargaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recarga.Contract.RecargaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Recarga *RecargaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Recarga.Contract.RecargaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Recarga *RecargaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Recarga.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Recarga *RecargaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Recarga.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Recarga *RecargaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Recarga.Contract.contract.Transact(opts, method, params...)
}

// ContadorRecargas is a free data retrieval call binding the contract method 0xf7e04e4d.
//
// Solidity: function contadorRecargas() view returns(uint256)
func (_Recarga *RecargaCaller) ContadorRecargas(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Recarga.contract.Call(opts, &out, "contadorRecargas")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContadorRecargas is a free data retrieval call binding the contract method 0xf7e04e4d.
//
// Solidity: function contadorRecargas() view returns(uint256)
func (_Recarga *RecargaSession) ContadorRecargas() (*big.Int, error) {
	return _Recarga.Contract.ContadorRecargas(&_Recarga.CallOpts)
}

// ContadorRecargas is a free data retrieval call binding the contract method 0xf7e04e4d.
//
// Solidity: function contadorRecargas() view returns(uint256)
func (_Recarga *RecargaCallerSession) ContadorRecargas() (*big.Int, error) {
	return _Recarga.Contract.ContadorRecargas(&_Recarga.CallOpts)
}

// GetRecargasDoUsuario is a free data retrieval call binding the contract method 0xdeb9122d.
//
// Solidity: function getRecargasDoUsuario(address usuario) view returns(uint256[] ids, uint256[] valores, bool[] pagos, uint256[] timestamps)
func (_Recarga *RecargaCaller) GetRecargasDoUsuario(opts *bind.CallOpts, usuario common.Address) (struct {
	Ids        []*big.Int
	Valores    []*big.Int
	Pagos      []bool
	Timestamps []*big.Int
}, error) {
	var out []interface{}
	err := _Recarga.contract.Call(opts, &out, "getRecargasDoUsuario", usuario)

	outstruct := new(struct {
		Ids        []*big.Int
		Valores    []*big.Int
		Pagos      []bool
		Timestamps []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ids = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Valores = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Pagos = *abi.ConvertType(out[2], new([]bool)).(*[]bool)
	outstruct.Timestamps = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetRecargasDoUsuario is a free data retrieval call binding the contract method 0xdeb9122d.
//
// Solidity: function getRecargasDoUsuario(address usuario) view returns(uint256[] ids, uint256[] valores, bool[] pagos, uint256[] timestamps)
func (_Recarga *RecargaSession) GetRecargasDoUsuario(usuario common.Address) (struct {
	Ids        []*big.Int
	Valores    []*big.Int
	Pagos      []bool
	Timestamps []*big.Int
}, error) {
	return _Recarga.Contract.GetRecargasDoUsuario(&_Recarga.CallOpts, usuario)
}

// GetRecargasDoUsuario is a free data retrieval call binding the contract method 0xdeb9122d.
//
// Solidity: function getRecargasDoUsuario(address usuario) view returns(uint256[] ids, uint256[] valores, bool[] pagos, uint256[] timestamps)
func (_Recarga *RecargaCallerSession) GetRecargasDoUsuario(usuario common.Address) (struct {
	Ids        []*big.Int
	Valores    []*big.Int
	Pagos      []bool
	Timestamps []*big.Int
}, error) {
	return _Recarga.Contract.GetRecargasDoUsuario(&_Recarga.CallOpts, usuario)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Recarga *RecargaCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Recarga.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Recarga *RecargaSession) Owner() (common.Address, error) {
	return _Recarga.Contract.Owner(&_Recarga.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Recarga *RecargaCallerSession) Owner() (common.Address, error) {
	return _Recarga.Contract.Owner(&_Recarga.CallOpts)
}

// Recargas is a free data retrieval call binding the contract method 0xb5cd8033.
//
// Solidity: function recargas(uint256 ) view returns(address usuario, uint256 valor, uint256 timestamp, bool pago)
func (_Recarga *RecargaCaller) Recargas(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Usuario   common.Address
	Valor     *big.Int
	Timestamp *big.Int
	Pago      bool
}, error) {
	var out []interface{}
	err := _Recarga.contract.Call(opts, &out, "recargas", arg0)

	outstruct := new(struct {
		Usuario   common.Address
		Valor     *big.Int
		Timestamp *big.Int
		Pago      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Usuario = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Valor = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Pago = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Recargas is a free data retrieval call binding the contract method 0xb5cd8033.
//
// Solidity: function recargas(uint256 ) view returns(address usuario, uint256 valor, uint256 timestamp, bool pago)
func (_Recarga *RecargaSession) Recargas(arg0 *big.Int) (struct {
	Usuario   common.Address
	Valor     *big.Int
	Timestamp *big.Int
	Pago      bool
}, error) {
	return _Recarga.Contract.Recargas(&_Recarga.CallOpts, arg0)
}

// Recargas is a free data retrieval call binding the contract method 0xb5cd8033.
//
// Solidity: function recargas(uint256 ) view returns(address usuario, uint256 valor, uint256 timestamp, bool pago)
func (_Recarga *RecargaCallerSession) Recargas(arg0 *big.Int) (struct {
	Usuario   common.Address
	Valor     *big.Int
	Timestamp *big.Int
	Pago      bool
}, error) {
	return _Recarga.Contract.Recargas(&_Recarga.CallOpts, arg0)
}

// RecargasPorUsuario is a free data retrieval call binding the contract method 0x70a0db5e.
//
// Solidity: function recargasPorUsuario(address , uint256 ) view returns(uint256)
func (_Recarga *RecargaCaller) RecargasPorUsuario(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Recarga.contract.Call(opts, &out, "recargasPorUsuario", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RecargasPorUsuario is a free data retrieval call binding the contract method 0x70a0db5e.
//
// Solidity: function recargasPorUsuario(address , uint256 ) view returns(uint256)
func (_Recarga *RecargaSession) RecargasPorUsuario(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Recarga.Contract.RecargasPorUsuario(&_Recarga.CallOpts, arg0, arg1)
}

// RecargasPorUsuario is a free data retrieval call binding the contract method 0x70a0db5e.
//
// Solidity: function recargasPorUsuario(address , uint256 ) view returns(uint256)
func (_Recarga *RecargaCallerSession) RecargasPorUsuario(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Recarga.Contract.RecargasPorUsuario(&_Recarga.CallOpts, arg0, arg1)
}

// PagarRecarga is a paid mutator transaction binding the contract method 0xc842bccb.
//
// Solidity: function pagarRecarga(uint256 recargaId) payable returns()
func (_Recarga *RecargaTransactor) PagarRecarga(opts *bind.TransactOpts, recargaId *big.Int) (*types.Transaction, error) {
	return _Recarga.contract.Transact(opts, "pagarRecarga", recargaId)
}

// PagarRecarga is a paid mutator transaction binding the contract method 0xc842bccb.
//
// Solidity: function pagarRecarga(uint256 recargaId) payable returns()
func (_Recarga *RecargaSession) PagarRecarga(recargaId *big.Int) (*types.Transaction, error) {
	return _Recarga.Contract.PagarRecarga(&_Recarga.TransactOpts, recargaId)
}

// PagarRecarga is a paid mutator transaction binding the contract method 0xc842bccb.
//
// Solidity: function pagarRecarga(uint256 recargaId) payable returns()
func (_Recarga *RecargaTransactorSession) PagarRecarga(recargaId *big.Int) (*types.Transaction, error) {
	return _Recarga.Contract.PagarRecarga(&_Recarga.TransactOpts, recargaId)
}

// RegistrarRecarga is a paid mutator transaction binding the contract method 0x733a5173.
//
// Solidity: function registrarRecarga(address usuario, uint256 valor) returns()
func (_Recarga *RecargaTransactor) RegistrarRecarga(opts *bind.TransactOpts, usuario common.Address, valor *big.Int) (*types.Transaction, error) {
	return _Recarga.contract.Transact(opts, "registrarRecarga", usuario, valor)
}

// RegistrarRecarga is a paid mutator transaction binding the contract method 0x733a5173.
//
// Solidity: function registrarRecarga(address usuario, uint256 valor) returns()
func (_Recarga *RecargaSession) RegistrarRecarga(usuario common.Address, valor *big.Int) (*types.Transaction, error) {
	return _Recarga.Contract.RegistrarRecarga(&_Recarga.TransactOpts, usuario, valor)
}

// RegistrarRecarga is a paid mutator transaction binding the contract method 0x733a5173.
//
// Solidity: function registrarRecarga(address usuario, uint256 valor) returns()
func (_Recarga *RecargaTransactorSession) RegistrarRecarga(usuario common.Address, valor *big.Int) (*types.Transaction, error) {
	return _Recarga.Contract.RegistrarRecarga(&_Recarga.TransactOpts, usuario, valor)
}
