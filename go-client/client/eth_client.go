package client

import (
	"context"
	"log"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"pbl3-blockchain/go-client/recarga"
	"pbl3-blockchain/go-client/reserva"
)

type Client struct {
	EthClient       *ethclient.Client
	RecargaContract *recarga.Recarga
	ReservaContract *reserva.Reserva
	Auth            *bind.TransactOpts
}

func NewClient(rpcURL, privateKeyHex, contractRecargaAddress, contractReservaAddress string) (*Client, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337)) // Ganache normalmente usa 1337
	if err != nil {
		return nil, err
	}

	contractRecargaAddr := common.HexToAddress(contractRecargaAddress)
	contractRecargaInstance, err := recarga.NewRecarga(contractRecargaAddr, client)
	if err != nil {
		return nil, err
	}

	contractReservaAddr := common.HexToAddress(contractReservaAddress)
	contractReservaInstance, err := reserva.NewReserva(contractReservaAddr, client)
	if err != nil {
		return nil, err
	}

	return &Client{
		EthClient:       client,
		RecargaContract: contractRecargaInstance,
		ReservaContract: contractReservaInstance,
		Auth:            auth,
	}, nil
}

func (c *Client) RegistrarRecarga(usuario common.Address, valor *big.Int) error {
	tx, err := c.RecargaContract.RegistrarRecarga(c.Auth, usuario, valor)
	if err != nil {
		return err
	}
	log.Printf("Recarga registrada! TX: %s", tx.Hash().Hex())
	return nil
}

func (c *Client) PagarRecarga(recargaID *big.Int, valor *big.Int) error {
	opts := *c.Auth // copia
	opts.Value = valor
	tx, err := c.RecargaContract.PagarRecarga(&opts, recargaID)
	if err != nil {
		return err
	}
	log.Printf("Recarga paga! TX: %s", tx.Hash().Hex())
	return nil
}

func (c *Client) GetRecargasDoUsuario(usuario common.Address) ([]*big.Int, []*big.Int, error) {
	valores, err := c.RecargaContract.GetRecargasDoUsuario(&bind.CallOpts{
		Pending: false,
		From:    usuario,
		Context: context.Background(),
	}, usuario)

	if err != nil {
		return nil, nil, err
	}

	return valores.Ids, valores.Valores, nil
}

// Cria nova instância do cliente de reserva
func NewReservaClient(rpcURL, privateKeyHex, contractAddress string) (*Client, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	log.Fatal("erro ao converter chave pública")
	// }

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337)) // chain ID do Ganache
	if err != nil {
		return nil, err
	}

	contractAddr := common.HexToAddress(contractAddress)
	contractInstance, err := reserva.NewReserva(contractAddr, client)
	if err != nil {
		return nil, err
	}

	return &Client{
		EthClient:       client,
		ReservaContract: contractInstance,
		Auth:            auth,
	}, nil
}

// Função para reservar um ponto
func (c *Client) ReservarPonto(pontoId, inicio, fim *big.Int) error {
	tx, err := c.ReservaContract.ReservarPonto(c.Auth, pontoId, inicio, fim)
	if err != nil {
		return err
	}
	log.Printf("Reserva feita! TX: %s", tx.Hash().Hex())
	return nil
}

// Função para cancelar uma reserva existente
func (c *Client) CancelarReserva(reservaId *big.Int) error {
	tx, err := c.ReservaContract.CancelarReserva(c.Auth, reservaId)
	if err != nil {
		return err
	}
	log.Printf("Reserva cancelada! TX: %s", tx.Hash().Hex())
	return nil
}

// Função para listar as reservas de um usuário
func (c *Client) GetReservasDoUsuario(usuario common.Address, ctx *types.Transaction) ([]*big.Int, []*big.Int, error) {
	values, err := c.ReservaContract.GetReservasDoUsuario(&bind.CallOpts{
		Pending: false,
		From:    usuario,
		Context: context.Background(),
	}, usuario)
	if err != nil {
		slog.Error("Erro executando método no contrato.")
		return nil, nil, err
	}
	return values.Ids, values.Pontos, nil
}
