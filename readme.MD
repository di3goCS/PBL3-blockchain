## Ativos
### Recarga (Charging Session)
Representa o ato de carregar um veículo em um ponto de recarga.

É o registro real da operação, incluindo:
- Identificação do ponto de recarga (referência, não tokenizado)
- Usuário que realizou a recarga
- Tempo de início e fim da recarga
- Quantidade de energia consumida (kWh)
- Status da recarga (iniciada, concluída, cancelada)
- Pode ser representado como uma struct dentro do contrato ou até um NFT que documenta a sessão.

### Reserva
- Registro da reserva para um ponto e horário específico.
- Relaciona usuário, ponto e intervalo de tempo reservado.
- Serve para garantir o direito de uso do ponto para aquela janela.

### Pagamento / Crédito
- Gerencia os pagamentos e créditos dos usuários.
- Vinculado a recargas e reservas para controle financeiro.

### Usuário
- Representado por endereço Ethereum, com histórico de reservas, recargas e pagamentos.

## Contratos
### Contrato de Recargas
- Cria e registra sessões de recarga.
- Armazena informações detalhadas da recarga (ponto, usuário, energia consumida, tempo).
- Garante que só recargas válidas sejam registradas.
- Pode emitir um NFT que representa a recarga para fins de auditoria.

### Contrato de Reserva
- Continua responsável por gerenciar reservas dos pontos.
- Garante que o ponto estará disponível para o usuário na janela reservada.

### Contrato de Pagamento / Créditos
- Lida com pagamentos e créditos relacionados a reservas e recargas.
- Garante que a recarga só ocorra se houver saldo disponível ou pagamento confirmado.

### Contrato de Histórico / Auditoria
- Mantém logs imutáveis das recargas e pagamentos para transparência.

## Fluxo básico
- Usuário reserva um ponto para um horário (Contrato de Reserva).
- Na hora da recarga, a sessão de recarga é criada e registrada (Contrato de Recargas).
- Usuário paga a reserva ou cria um débito (Contrato de Pagamento).
- Eventos são emitidos para sistemas externos (frontends, monitoramento).

### Como executar:

