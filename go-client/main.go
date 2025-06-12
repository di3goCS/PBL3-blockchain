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
	// Importe aqui se seus bindings estiverem em um pacote diferente de 'main'
	// Ex: "PBL3-Blockchain/go-client"
)

// Importante: A chave privada DEVE ser de uma conta com ETH no Ganache ou testnet
// NUNCA USE CHAVES PRIVADAS REAIS EM CÓDIGO FONTE OU EM PRODUÇÃO!
const privateKeyHex = "0xd0d25dd79bd8fc56cc4be007e07cf114e41430d3b1c108a597883e3c29967d1e" // <<<<< ATUALIZE AQUI
const privateKeyHexWithout0x = "d0d25dd79bd8fc56cc4be007e07cf114e41430d3b1c108a597883e3c29967d1e"
const ganacheURL = "http://127.0.0.1:7545" // <<<<< ATUALIZE AQUI (era 8545, agora 7545)
const ganacheChainID = 1337                // <<<<< ATUALIZE AQUI (era 1337, agora 5777)

func main() {

	// clientAddress = "0x1A9de9b0C321442FFd0752Fa8EbdE58E8899F008"
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

	// c, err := localclient.NewClient(ganacheURL, privateKeyHex, reservaAddress.Hex(), recargaAddress.Hex())
	// if err != nil {
	// 	fmt.Println("Erro instanciando client.", err)
	// }

	clientAddress := common.HexToAddress(privateKeyHex)

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
				fmt.Println("Erro ao cancelar reserva")
			}

		case "3":
			fmt.Printf("-> Listar reservas ativas de %s\n", clientAddress)
			reservasIDs, reservas, err := c.GetReservasDoUsuario(clientAddress)
			if err != nil {
				fmt.Println("Erro ao listar reservas do usuário")
			}
			fmt.Println("Reservas: ")
			for idx, reserva := range reservasIDs {
				fmt.Printf("[%d] %d\n", reserva, reservas[idx])
			}

		case "4":
			fmt.Print("Informe o endereço do usuário: ")
			endereco, _ := reader.ReadString('\n')
			endereco = strings.TrimSpace(endereco)

			fmt.Print("Informe o valor da recarga (em ETH): ")
			valorStr, _ := reader.ReadString('\n')
			valorEth, _ := strconv.ParseInt(strings.TrimSpace(valorStr), 10, 64)

			wei := new(big.Int)
			wei.SetString(fmt.Sprintf("%.0f", valorEth*1e18), 10)

			fmt.Printf("-> Registrar recarga de %s no valor de %s wei\n", endereco, wei.String())
			err := c.RegistrarRecarga(clientAddress, big.NewInt(valorEth))
			if err != nil {
				fmt.Println("Erro ao registrar recarga")
			}

		case "5":
			fmt.Print("Informe o ID da recarga a pagar: ")
			idStr, _ := reader.ReadString('\n')
			recargaID, _ := strconv.ParseInt(strings.TrimSpace(idStr), 10, 64)

			fmt.Print("Informe o valor do pagamento (em ETH): ")
			valorStr, _ := reader.ReadString('\n')
			valorEth, _ := strconv.ParseInt(strings.TrimSpace(valorStr), 10, 64)

			wei := new(big.Int)
			wei.SetString(fmt.Sprintf("%.0f", valorEth*1e18), 10)

			fmt.Printf("-> Pagar recarga %d com %s wei\n", recargaID, wei.String())
			err := c.PagarRecarga(big.NewInt(recargaID), big.NewInt(valorEth))
			if err != nil {
				fmt.Println("Erro ao realizar pagamento de recarga")
			}

		case "6":
			fmt.Print("Informe o endereço do usuário: ")
			endereco, _ := reader.ReadString('\n')
			endereco = strings.TrimSpace(endereco)

			fmt.Printf("-> Listar recargas do usuário %s\n", endereco)
			recargasIDs, recargas, err := c.GetRecargasDoUsuario(clientAddress)
			if err != nil {
				fmt.Println("Erro ao listar reservas do usuário")
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

	// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
	// 	log.Fatalf("Falha ao obter nonce: %v", err)
	// }

	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatalf("Falha ao obter gas price: %v", err)
	// }

	// auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(ganacheChainID))
	// if err != nil {
	// 	log.Fatalf("Falha ao criar transator: %v", err)
	// }
	// auth.Nonce = big.NewInt(int64(nonce))
	// auth.Value = big.NewInt(0)
	// auth.GasLimit = uint64(3000000)
	// auth.GasPrice = gasPrice

	// fmt.Println("Conectado ao nó Ethereum e autenticador pronto.")

	// // --- Exemplo de Implantação e Interação com RecargaVeiculo ---
	// fmt.Println("\n--- Contrato RecargaVeiculo ---")
	// // As funções DeployRecargaVeiculo e DeployReservaPontoRecarga são geradas pelo abigen
	// // Se seus contratos Solidity se chamam RecargaVeiculo e ReservaPontoRecarga internamente,
	// // e os bindings estão no pacote 'main', o abigen gerou essas funções.
	// rechargeAddr, rechargeTx, _, err := recarga.DeployRecarga(auth, client)
	// if err != nil {
	// 	log.Fatalf("Falha ao implantar RecargaVeiculo: %v", err)
	// }
	// fmt.Printf("RecargaVeiculo implantado em: %s\n", rechargeAddr.Hex())
	// fmt.Printf("Hash da transação de implantação RecargaVeiculo: %s\n", rechargeTx.Hash().Hex())

	// _, err = bind.WaitDeployed(context.Background(), client, rechargeTx)
	// if err != nil {
	// 	log.Fatalf("Falha ao esperar pela implantação de RecargaVeiculo: %v", err)
	// }
	// fmt.Println("RecargaVeiculo implantado e confirmado!")

	// // Coloque aqui o código para interagir com RecargaVeiculo (chamar funções, etc.)
	// // Lembre-se de incrementar auth.Nonce ou chamar client.PendingNonceAt(context.Background(), fromAddress)
	// // antes de cada nova transação de escrita, para evitar erros de nonce.
	// auth.Nonce = big.NewInt(int64(nonce + 1)) // Exemplo: incrementando para a próxima transação

	// // --- Exemplo de Implantação e Interação com ReservaPontoRecarga ---
	// fmt.Println("\n--- Contrato ReservaPontoRecarga ---")
	// reserveAddr, reserveTx, _, err := reserva.DeployReserva(auth, client)
	// if err != nil {
	// 	log.Fatalf("Falha ao implantar ReservaPontoRecarga: %v", err)
	// }
	// fmt.Printf("ReservaPontoRecarga implantado em: %s\n", reserveAddr.Hex())
	// fmt.Printf("Hash da transação de implantação ReservaPontoRecarga: %s\n", reserveTx.Hash().Hex())

	// _, err = bind.WaitDeployed(context.Background(), client, reserveTx)
	// if err != nil {
	// 	log.Fatalf("Falha ao esperar pela implantação de ReservaPontoRecarga: %v", err)
	// }
	// fmt.Println("ReservaPontoRecarga implantado e confirmado!")

	// // Coloque aqui o código para interagir com ReservaPontoRecarga (chamar funções, etc.)

	// fmt.Println("\nPrograma Go concluído. Verifique o Ganache para as transações!")
}
