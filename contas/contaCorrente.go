package contas

import (
	"banco-go-orientacao-a-objetos-alura/clientes"
	"fmt"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) {
	if valorDoSaque > 0 {
		if c.saldo >= valorDoSaque {
			c.saldo -= valorDoSaque
			fmt.Println("Saque no valor de", valorDoSaque, "realizado com sucesso!")
		} else {
			fmt.Println("Saldo insuficiente! O saque no valor de", valorDoSaque, "não foi realizado")
		}
	} else {
		fmt.Println("O Valor", valorDoSaque, "é inválido para saque!")
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		fmt.Println("O depósito do valor de", valorDoDeposito, "foi realizado com sucesso!")
	} else {
		fmt.Println("O Valor", valorDoDeposito, "é inválido para depósito!")
	}
}

func (c *ContaCorrente) transferir(valorTransferencia float64, contaDestino *ContaCorrente) {
	if valorTransferencia > 0 {
		if c.saldo >= valorTransferencia {
			c.saldo -= valorTransferencia
			contaDestino.Depositar(valorTransferencia)
			fmt.Println("Transferência no valor de", valorTransferencia, "foi realizada com sucesso!")
		} else {
			fmt.Println("Saldo insuficiente! A transferência no valor de", valorTransferencia, "não foi realizado")
		}
	} else {
		fmt.Println("O Valor", valorTransferencia, "é inválido para transferência!")
	}
}

func (c *ContaCorrente) ObterSaldo() (valorDoSaque float64) {
	return c.saldo
}
