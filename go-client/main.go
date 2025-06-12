package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	localclient "pbl3-blockchain/go-client/client"
	"pbl3-blockchain/go-client/recarga"
	"pbl3-blockchain/go-client/reserva"
)

// Importante: A chave privada DEVE ser de uma conta com ETH no Ganache ou testnet
// NUNCA USE CHAVES PRIVADAS REAIS EM CÓDIGO FONTE OU EM PRODUÇÃO!
const privateKeyHex = "0xd0d25dd79bd8fc56cc4be007e07cf114e41430d3b1c108a597883e3c29967d1e" // <<<<< ATUALIZE AQUI
const privateKeyHexWithout0x = "d0d25dd79bd8fc56cc4be007e07cf114e41430d3b1c108a597883e3c29967d1e"
const ganacheURL = "http://127.0.0.1:7545" // <<<<< ATUALIZE AQUI (era 8545, agora 7545)
const ganacheChainID = 1337                // <<<<< ATUALIZE AQUI (era 1337, agora 5777)

func main() {
	client, err := ethclient.Dial(ganacheURL)
	if err != nil {
		log.Fatalf("Falha ao conectar ao Ethereum: %v", err)
	}
	defer client.Close()

	privateKey, err := crypto.HexToECDSA(privateKeyHexWithout0x)
	if err != nil {
		log.Fatalf("Falha ao carregar chave privada: %v", err)
	}

	// publicKey := privateKey.Public()
	// fromAddress := crypto.PubkeyToAddress(publicKey.(ecdsa.PublicKey))

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(ganacheChainID)) // chainId do Ganache
	reservaAddress, txReserva, reservaInstance, err := reserva.DeployReserva(auth, client)
	recargaAddress, txRecarga, recargaInstance, err := recarga.DeployRecarga(auth, client)
	fmt.Println("RESERVA: ", reservaAddress, txReserva)
	fmt.Println("RECARGA: ", recargaAddress, txRecarga)

	c := localclient.Client{
		EthClient:       client,
		RecargaContract: recargaInstance,
		ReservaContract: reservaInstance,
		Auth:            auth,
	}

	clientAddress := common.HexToAddress(privateKeyHex)
	fmt.Println("CLIENTE: ", clientAddress)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== MENU DE TESTE DO SISTEMA DE RECARGA ===")

	for {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1. Reservar ponto")
		fmt.Println("2. Cancelar reserva")
		fmt.Println("3. Listar reservas do usuário")
		fmt.Println("4. Registrar recarga")
		fmt.Println("5. Pagar recarga")
		fmt.Println("6. Listar recargas do usuário")
		fmt.Println("0. Sair")
		fmt.Print("Escolha uma opção: ")

		entrada, _ := reader.ReadString('\n')
		entrada = strings.TrimSpace(entrada)

		switch entrada {
		case "1":
			fmt.Print("Informe o ID do ponto: ")
			pontoStr, _ := reader.ReadString('\n')
			pontoID, _ := strconv.ParseInt(strings.TrimSpace(pontoStr), 10, 64)

			fmt.Print("Informe o horário de início (timestamp): ")
			inicioStr, _ := reader.ReadString('\n')
			inicio, _ := strconv.ParseInt(strings.TrimSpace(inicioStr), 10, 64)

			fmt.Print("Informe o horário de fim (timestamp): ")
			fimStr, _ := reader.ReadString('\n')
			fim, _ := strconv.ParseInt(strings.TrimSpace(fimStr), 10, 64)

			fmt.Printf("-> Reservar ponto %d de %d até %d\n", pontoID, inicio, fim)
			err := c.ReservarPonto(big.NewInt(pontoID), big.NewInt(inicio), big.NewInt(fim))
			if err != nil {
				fmt.Println("Erro ao reservar ponto")
			}

		case "2":
			fmt.Print("Informe o ID da reserva a cancelar: ")
			idStr, _ := reader.ReadString('\n')
			reservaID, _ := strconv.ParseInt(strings.TrimSpace(idStr), 10, 64)

			fmt.Printf("-> Cancelar reserva com ID %d\n", reservaID)
			err := c.CancelarReserva(big.NewInt(reservaID))
			if err != nil {
				fmt.Printf("Erro ao cancelar reserva: %v", err)
			}

		case "3":
			fmt.Printf("-> Listar reservas ativas de %s\n", clientAddress)
			reservasIDs, reservas, err := c.GetReservasDoUsuario(clientAddress, txReserva)
			if err != nil {
				fmt.Printf("Erro ao listar reservas do usuário: %v", err)
			}
			fmt.Println("Reservas: ")
			for idx, reserva := range reservasIDs {
				fmt.Printf("[%d] %d\n", reserva, reservas[idx])
			}

		case "4":
			fmt.Print("Informe o valor da recarga (em ETH): ")
			valorStr, _ := reader.ReadString('\n')
			valorEth, _ := strconv.ParseInt(strings.TrimSpace(valorStr), 10, 64)

			wei := new(big.Int)
			wei.SetString(fmt.Sprintf("%d", valorEth*1e18), 10)

			fmt.Printf("-> Registrar recarga de %s no valor de %s wei\n", clientAddress, wei.String())
			err := c.RegistrarRecarga(clientAddress, big.NewInt(valorEth))
			if err != nil {
				fmt.Printf("Erro ao registrar recarga: %v", err)
			}

		case "5":
			fmt.Print("Informe o ID da recarga a pagar: ")
			idStr, _ := reader.ReadString('\n')
			recargaID, _ := strconv.ParseInt(strings.TrimSpace(idStr), 10, 64)

			fmt.Print("Informe o valor do pagamento (em ETH): ")
			valorStr, _ := reader.ReadString('\n')
			valorEth, _ := strconv.ParseInt(strings.TrimSpace(valorStr), 10, 64)

			wei := new(big.Int)
			wei.SetString(fmt.Sprintf("%d", valorEth*1e18), 10)

			fmt.Printf("-> Pagar recarga %d com %s wei\n", recargaID, wei.String())
			err := c.PagarRecarga(big.NewInt(recargaID), big.NewInt(valorEth))
			if err != nil {
				fmt.Printf("Erro ao realizar pagamento de recarga: %v", err)
			}

		case "6":
			fmt.Printf("-> Listar recargas do usuário %s\n", clientAddress)
			recargasIDs, recargas, err := c.GetRecargasDoUsuario(clientAddress)
			if err != nil {
				fmt.Printf("Erro ao listar reservas do usuário: %v", err)
			}
			fmt.Println("Recargas: ")
			for idx, recarga := range recargasIDs {
				fmt.Printf("[%d] %d\n", recarga, recargas[idx])
			}

		case "0":
			fmt.Println("Saindo...")
			return

		default:
			fmt.Println("Opção inválida!")
		}
	}
}
