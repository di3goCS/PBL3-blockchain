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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"contadorRecargas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usuario\",\"type\":\"address\"}],\"name\":\"getRecargasDoUsuario\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"valores\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"recargaId\",\"type\":\"uint256\"}],\"name\":\"pagarRecarga\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recargas\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"usuario\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pago\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recargasPorUsuario\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usuario\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valor\",\"type\":\"uint256\"}],\"name\":\"registrarRecarga\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610ad98061005b5f395ff3fe60806040526004361061006f575f3560e01c8063b5cd80331161004d578063b5cd803314610101578063c842bccb14610140578063deb9122d1461015c578063f7e04e4d146101995761006f565b806370a0db5e14610073578063733a5173146100af5780638da5cb5b146100d7575b5f5ffd5b34801561007e575f5ffd5b50610099600480360381019061009491906106c8565b6101c3565b6040516100a69190610715565b60405180910390f35b3480156100ba575f5ffd5b506100d560048036038101906100d091906106c8565b6101ee565b005b3480156100e2575f5ffd5b506100eb610332565b6040516100f8919061073d565b60405180910390f35b34801561010c575f5ffd5b5061012760048036038101906101229190610756565b610356565b604051610137949392919061079b565b60405180910390f35b61015a60048036038101906101559190610756565b6103ad565b005b348015610167575f5ffd5b50610182600480360381019061017d91906107de565b6104dd565b6040516101909291906108c0565b60405180910390f35b3480156101a4575f5ffd5b506101ad610631565b6040516101ba9190610715565b60405180910390f35b6002602052815f5260405f2081815481106101dc575f80fd5b905f5260205f20015f91509150505481565b60035f81548092919061020090610922565b919050555060405180608001604052808373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020014281526020015f151581525060015f60035481526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201556060820151816003015f6101000a81548160ff02191690831515021790555090505060025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20600354908060018154018082558091505060019003905f5260205f20015f90919091909150555050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6001602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015490806002015490806003015f9054906101000a900460ff16905084565b5f60015f8381526020019081526020015f209050806003015f9054906101000a900460ff1615610412576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610409906109c3565b60405180910390fd5b8060010154341015610459576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045090610a2b565b60405180910390fd5b6001816003015f6101000a81548160ff0219169083151502179055505f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc3490811502906040515f60405180830381858888f193505050501580156104d8573d5f5f3e3d5ffd5b505050565b6060805f60025f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2080548060200260200160405190810160405280929190818152602001828054801561056657602002820191905f5260205f20905b815481526020019060010190808311610552575b505050505090505f815167ffffffffffffffff81111561058957610588610a49565b5b6040519080825280602002602001820160405280156105b75781602001602082028036833780820191505090505b5090505f5f90505b82518110156106235760015f8483815181106105de576105dd610a76565b5b602002602001015181526020019081526020015f206001015482828151811061060a57610609610a76565b5b60200260200101818152505080806001019150506105bf565b508181935093505050915091565b60035481565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6106648261063b565b9050919050565b6106748161065a565b811461067e575f5ffd5b50565b5f8135905061068f8161066b565b92915050565b5f819050919050565b6106a781610695565b81146106b1575f5ffd5b50565b5f813590506106c28161069e565b92915050565b5f5f604083850312156106de576106dd610637565b5b5f6106eb85828601610681565b92505060206106fc858286016106b4565b9150509250929050565b61070f81610695565b82525050565b5f6020820190506107285f830184610706565b92915050565b6107378161065a565b82525050565b5f6020820190506107505f83018461072e565b92915050565b5f6020828403121561076b5761076a610637565b5b5f610778848285016106b4565b91505092915050565b5f8115159050919050565b61079581610781565b82525050565b5f6080820190506107ae5f83018761072e565b6107bb6020830186610706565b6107c86040830185610706565b6107d5606083018461078c565b95945050505050565b5f602082840312156107f3576107f2610637565b5b5f61080084828501610681565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61083b81610695565b82525050565b5f61084c8383610832565b60208301905092915050565b5f602082019050919050565b5f61086e82610809565b6108788185610813565b935061088383610823565b805f5b838110156108b357815161089a8882610841565b97506108a583610858565b925050600181019050610886565b5085935050505092915050565b5f6040820190508181035f8301526108d88185610864565b905081810360208301526108ec8184610864565b90509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61092c82610695565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361095e5761095d6108f5565b5b600182019050919050565b5f82825260208201905092915050565b7f52656361726761206a61207061676100000000000000000000000000000000005f82015250565b5f6109ad600f83610969565b91506109b882610979565b602082019050919050565b5f6020820190508181035f8301526109da816109a1565b9050919050565b7f56616c6f7220696e737566696369656e746500000000000000000000000000005f82015250565b5f610a15601283610969565b9150610a20826109e1565b602082019050919050565b5f6020820190508181035f830152610a4281610a09565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea2646970667358221220424ad28b72e5c260bd263da04b45613c3e0c85db742a37385fbf7c766a16fb6764736f6c634300081e0033",
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
// Solidity: function getRecargasDoUsuario(address usuario) view returns(uint256[] ids, uint256[] valores)
func (_Recarga *RecargaCaller) GetRecargasDoUsuario(opts *bind.CallOpts, usuario common.Address) (struct {
	Ids     []*big.Int
	Valores []*big.Int
}, error) {
	var out []interface{}
	err := _Recarga.contract.Call(opts, &out, "getRecargasDoUsuario", usuario)

	outstruct := new(struct {
		Ids     []*big.Int
		Valores []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ids = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Valores = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetRecargasDoUsuario is a free data retrieval call binding the contract method 0xdeb9122d.
//
// Solidity: function getRecargasDoUsuario(address usuario) view returns(uint256[] ids, uint256[] valores)
func (_Recarga *RecargaSession) GetRecargasDoUsuario(usuario common.Address) (struct {
	Ids     []*big.Int
	Valores []*big.Int
}, error) {
	return _Recarga.Contract.GetRecargasDoUsuario(&_Recarga.CallOpts, usuario)
}

// GetRecargasDoUsuario is a free data retrieval call binding the contract method 0xdeb9122d.
//
// Solidity: function getRecargasDoUsuario(address usuario) view returns(uint256[] ids, uint256[] valores)
func (_Recarga *RecargaCallerSession) GetRecargasDoUsuario(usuario common.Address) (struct {
	Ids     []*big.Int
	Valores []*big.Int
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
