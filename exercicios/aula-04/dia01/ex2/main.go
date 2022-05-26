package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	var salario int
	if salario < 15000 {
		log.Println(errors.New("erro: O salário digitado não está dentro do valor mínimo"))
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}
}
