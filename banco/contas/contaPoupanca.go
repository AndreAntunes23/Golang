package contas

import "banco/clientes"

type ContaPoupança struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operação int
	saldo                                float64
}

func (c *ContaPoupança) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupança) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso! '\n' saldo atual: ", c.saldo
	} else {
		return "O valor do deposito não pode ser menor que zero", c.saldo
	}
}

func (c *ContaPoupança) ObterSaldo() float64 {
	return c.saldo
}
