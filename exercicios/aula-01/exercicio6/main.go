package main

import "fmt"

func main() {

	idade := 23
	trabalhoCom1ano := true
	salario := 10000

	if idade < 22 || trabalhoCom1ano == false {
		fmt.Println("Você não preenche os requisitos para conseguir um empréstimo")
	} else if salario >= 100000 {
		fmt.Println("Parabéns! Você conseguiu o empréstimo sem juros!")
	} else {
		fmt.Println("Parabéns! Você conseguiu o empréstimo!")
	}
}
