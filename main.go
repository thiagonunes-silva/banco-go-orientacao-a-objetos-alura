package main

import (
	"banco-go-orientacao-a-objetos-alura/clientes"
	"banco-go-orientacao-a-objetos-alura/contas"
	"fmt"
)

func main() {
	titular := clientes.Titular{
		Nome:      "Titular",
		CPF:       "123",
		Profissao: "Desenv Go",
	}

	conta1 := contas.ContaCorrente{
		Titular:       titular,
		NumeroAgencia: 1234,
		NumeroConta:   123456,
	}

	fmt.Println(conta1)

	conta1.Sacar(30.78)

	fmt.Println(conta1)

	conta1.Sacar(230.78)

	fmt.Println(conta1)

	conta1.Sacar(-230.78)

	fmt.Println(conta1)
}
