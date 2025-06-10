package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"pbl3-blockchain/go-client/recarga"
	"pbl3-blockchain/go-client/reserva"
	// Importe aqui se seus bindings estiverem em um pacote diferente de 'main'
	// Ex: "PBL3-Blockchain/go-client"
)

// Importante: A chave privada DEVE ser de uma conta com ETH no Ganache ou testnet
// NUNCA USE CHAVES PRIVADAS REAIS EM CÓDIGO FONTE OU EM PRODUÇÃO!
const privateKeyHex = "SUA_CHAVE_PRIVADA_HEXADECIMAL_COPIADA_DO_GANACHE_AQUI" // <<<<< ATUALIZE AQUI
const ganacheURL = "http://127.0.0.1:7545"                                    // <<<<< ATUALIZE AQUI (era 8545, agora 7545)
const ganacheChainID = 5777                                                   // <<<<< ATUALIZE AQUI (era 1337, agora 5777)

func main() {
	client, err := ethclient.Dial(ganacheURL)
	if err != nil {
		log.Fatalf("Falha ao conectar ao Ethereum: %v", err)
	}
	defer client.Close()

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Falha ao carregar chave privada: %v", err)
	}

	publicKey := privateKey.Public()
	fromAddress := crypto.PubkeyToAddress(publicKey.(ecdsa.PublicKey))

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Falha ao obter nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Falha ao obter gas price: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(ganacheChainID))
	if err != nil {
		log.Fatalf("Falha ao criar transator: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	fmt.Println("Conectado ao nó Ethereum e autenticador pronto.")

	// --- Exemplo de Implantação e Interação com RecargaVeiculo ---
	fmt.Println("\n--- Contrato RecargaVeiculo ---")
	// As funções DeployRecargaVeiculo e DeployReservaPontoRecarga são geradas pelo abigen
	// Se seus contratos Solidity se chamam RecargaVeiculo e ReservaPontoRecarga internamente,
	// e os bindings estão no pacote 'main', o abigen gerou essas funções.
	rechargeAddr, rechargeTx, _, err := recarga.DeployRecarga(auth, client)
	if err != nil {
		log.Fatalf("Falha ao implantar RecargaVeiculo: %v", err)
	}
	fmt.Printf("RecargaVeiculo implantado em: %s\n", rechargeAddr.Hex())
	fmt.Printf("Hash da transação de implantação RecargaVeiculo: %s\n", rechargeTx.Hash().Hex())

	_, err = bind.WaitDeployed(context.Background(), client, rechargeTx)
	if err != nil {
		log.Fatalf("Falha ao esperar pela implantação de RecargaVeiculo: %v", err)
	}
	fmt.Println("RecargaVeiculo implantado e confirmado!")

	// Coloque aqui o código para interagir com RecargaVeiculo (chamar funções, etc.)
	// Lembre-se de incrementar auth.Nonce ou chamar client.PendingNonceAt(context.Background(), fromAddress)
	// antes de cada nova transação de escrita, para evitar erros de nonce.
	auth.Nonce = big.NewInt(int64(nonce + 1)) // Exemplo: incrementando para a próxima transação

	// --- Exemplo de Implantação e Interação com ReservaPontoRecarga ---
	fmt.Println("\n--- Contrato ReservaPontoRecarga ---")
	reserveAddr, reserveTx, _, err := reserva.DeployReserva(auth, client)
	if err != nil {
		log.Fatalf("Falha ao implantar ReservaPontoRecarga: %v", err)
	}
	fmt.Printf("ReservaPontoRecarga implantado em: %s\n", reserveAddr.Hex())
	fmt.Printf("Hash da transação de implantação ReservaPontoRecarga: %s\n", reserveTx.Hash().Hex())

	_, err = bind.WaitDeployed(context.Background(), client, reserveTx)
	if err != nil {
		log.Fatalf("Falha ao esperar pela implantação de ReservaPontoRecarga: %v", err)
	}
	fmt.Println("ReservaPontoRecarga implantado e confirmado!")

	// Coloque aqui o código para interagir com ReservaPontoRecarga (chamar funções, etc.)

	fmt.Println("\nPrograma Go concluído. Verifique o Ganache para as transações!")
}
