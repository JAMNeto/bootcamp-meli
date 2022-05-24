package main

import "fmt"

func calculaImposto(funcionario string, salario float64) {
	totalDesconto := float64(0)
	if salario < 150000 {
		totalDesconto = salario * 0.17
	} else {
		totalDesconto = salario * 0.27
	}
	fmt.Printf("O funcionário %s recebeu um salário de %.2f, será descontado o valor de %.2f em impostos\n", funcionario, salario, totalDesconto)
}

func main() {
	calculaImposto("Carlos", 15000)
	calculaImposto("Augusto", 150000)
}
