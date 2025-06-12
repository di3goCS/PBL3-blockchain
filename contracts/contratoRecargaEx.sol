// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract RecargaVeiculo {
    address public owner;
// OWNER DO CONTRATO
    constructor() {
        owner = msg.sender;
    }

    struct Recarga {
        address usuario;
        uint256 valor;
        uint256 timestamp;
        bool pago;
    }
// MAPPING DAS RECARGAS POR USUÁRIO
    mapping(uint256 => Recarga) public recargas;
    mapping(address => uint256[]) public recargasPorUsuario;
    uint256 public contadorRecargas;

// REGISTRA AS RECARGAS E ADD NO CONTADOR E MAPPING
    function registrarRecarga(address usuario, uint256 valor) public {
        contadorRecargas++;
        recargas[contadorRecargas] = Recarga(usuario, valor, block.timestamp, false);
        recargasPorUsuario[usuario].push(contadorRecargas);
    }

// VERIFICA SE A RECARGA JÁ FOI PAGA, SE O USUÁRIO QUE ESTÁ TENTANDO PAGAR É O MESMO DO 
    function pagarRecarga(uint256 recargaId) public payable {
        Recarga storage recarga = recargas[recargaId];
        require(!recarga.pago, "Recarga ja paga");
        require(msg.value >= recarga.valor, "Valor insuficiente");

        recarga.pago = true;
        payable(owner).transfer(msg.value);
    }

// RETORNA OS IDS DAS RECARGAS E OS SEUS VALORES ASSOCIADOS
    function getRecargasDoUsuario(address usuario) public view returns (
        uint256[] memory ids,
        uint256[] memory valores,
        bool[] memory pagos,
        uint256[] memory timestamps
    ) {
        uint256[] memory idsUsuario = recargasPorUsuario[usuario];
        uint256[] memory valoresUsuario = new uint256[](idsUsuario.length);
        bool[] memory pagosUsuario = new bool[](idsUsuario.length);
        uint256[] memory timestampsUsuario = new uint256[](idsUsuario.length);

        for (uint256 i = 0; i < idsUsuario.length; i++) {
            Recarga memory r = recargas[idsUsuario[i]];
            valoresUsuario[i] = r.valor;
            pagosUsuario[i] = r.pago;
            timestampsUsuario[i] = r.timestamp;
        }

        return (idsUsuario, valoresUsuario, pagosUsuario, timestampsUsuario);
    }
}