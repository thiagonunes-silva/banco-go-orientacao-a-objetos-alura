package contas

type pagaBoleto interface {
	Sacar(valorDoSaque float64)
}

func PagarBoleto(conta pagaBoleto, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}
