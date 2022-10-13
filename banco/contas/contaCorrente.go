package contas

import (
	"banco/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso! '\n' saldo atual: ", c.saldo
	} else {
		return "O valor do deposito não pode ser menor que zero", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferência float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferência < c.saldo && valorDaTransferência > 0 {
		c.saldo -= valorDaTransferência
		contaDestino.Depositar(valorDaTransferência)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) Obtersaldo() float64 {
	return c.saldo
}
