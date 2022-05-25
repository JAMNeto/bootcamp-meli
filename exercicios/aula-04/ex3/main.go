package main

import (
	"fmt"
	"log"
)

type myError struct {
	message string
}

func (m *myError) Error(salario int) error {
	msg := fmt.Errorf("erro: o mínimo tributável é 15.000 e o salário informado é %d", salario)
	return msg
}

func main() {
	var salario int
	var novoErro myError
	if salario < 15000 {
		log.Println(novoErro.Error(salario))
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}
}
