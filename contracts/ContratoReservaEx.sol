// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ReservaPontoRecarga {
    address public owner;

    constructor() {
        owner = msg.sender;
    }

    struct Reserva {
        address usuario;
        uint256 pontoId;
        uint256 inicio;
        uint256 fim;
        bool ativa;
    }

    mapping(uint256 => Reserva) public reservas;
    mapping(address => uint256[]) public reservasPorUsuario;
    uint256 public contadorReservas;

    /// Cria uma nova reserva, validando conflitos por usuário e ponto
    function reservarPonto(uint256 pontoId, uint256 inicio, uint256 fim) public {
        require(inicio < fim, "Intervalo invalido");

        // Verificar conflitos do usuário (não pode ter 2 reservas simultâneas em qualquer ponto)
        uint256[] memory idsUsuario = reservasPorUsuario[msg.sender];
        for (uint256 i = 0; i < idsUsuario.length; i++) {
            Reserva memory r = reservas[idsUsuario[i]];
            if (r.ativa && intervalosConflitantes(r.inicio, r.fim, inicio, fim)) {
                revert("Voce ja possui reserva nesse horario");
            }
        }

        // Verificar se outro usuário já reservou o mesmo ponto nesse horário
        for (uint256 i = 1; i <= contadorReservas; i++) {
            Reserva memory r = reservas[i];
            if (r.ativa && r.pontoId == pontoId && intervalosConflitantes(r.inicio, r.fim, inicio, fim)) {
                revert("Ponto ja reservado nesse horario");
            }
        }

        // Criar reserva
        contadorReservas++;
        reservas[contadorReservas] = Reserva(msg.sender, pontoId, inicio, fim, true);
        reservasPorUsuario[msg.sender].push(contadorReservas);
    }

    /// Cancela uma reserva e remove do histórico do usuário
    function cancelarReserva(uint256 reservaId) public {
        Reserva storage r = reservas[reservaId];
        require(r.usuario == msg.sender, "Nao autorizado");
        require(r.ativa, "Reserva ja cancelada");

        r.ativa = false;

        // Remover reserva da lista do usuário
        uint256[] storage lista = reservasPorUsuario[msg.sender];
        for (uint256 i = 0; i < lista.length; i++) {
            if (lista[i] == reservaId) {
                // Substitui pelo último elemento e faz pop
                lista[i] = lista[lista.length - 1];
                lista.pop();
                break;
            }
        }
    }

    /// Retorna os IDs e pontoIds ativos do usuário
    function getReservasDoUsuario(address usuario) public view returns (uint256[] memory ids, uint256[] memory pontos) {
        uint256[] memory lista = reservasPorUsuario[usuario];
        uint256 count = 0;

        // Contar reservas ativas
        for (uint256 i = 0; i < lista.length; i++) {
            if (reservas[lista[i]].ativa) {
                count++;
            }
        }

        ids = new uint256[](count);
        pontos = new uint256[](count);
        uint256 j = 0;

        for (uint256 i = 0; i < lista.length; i++) {
            Reserva memory r = reservas[lista[i]];
            if (r.ativa) {
                ids[j] = lista[i];
                pontos[j] = r.pontoId;
                j++;
            }
        }
    }

    /// Função auxiliar para verificar se dois intervalos conflitam
    function intervalosConflitantes(uint256 inicio1, uint256 fim1, uint256 inicio2, uint256 fim2) internal pure returns (bool) {
        return (inicio1 < fim2 && inicio2 < fim1); // intervalo1 começa antes do fim2 E intervalo2 começa antes do fim1
    }
}
