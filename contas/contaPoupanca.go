package contas

import (
	"banco-go-orientacao-a-objetos-alura/clientes"
	"fmt"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) {
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

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		fmt.Println("O depósito do valor de", valorDoDeposito, "foi realizado com sucesso!")
	} else {
		fmt.Println("O Valor", valorDoDeposito, "é inválido para depósito!")
	}
}

func (c *ContaPoupanca) ObterSaldo() (valorDoSaque float64) {
	return c.saldo
}
