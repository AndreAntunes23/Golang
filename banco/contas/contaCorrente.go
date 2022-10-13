package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Deposito realizado com sucesso! '\n' Saldo atual: ", c.Saldo
	} else {
		return "O valor do deposito não pode ser menor que zero", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferência float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferência < c.Saldo && valorDaTransferência > 0 {
		c.Saldo -= valorDaTransferência
		contaDestino.Depositar(valorDaTransferência)
		return true
	} else {
		return false
	}
}
