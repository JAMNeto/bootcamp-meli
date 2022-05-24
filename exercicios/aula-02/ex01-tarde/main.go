package main

import "fmt"

type aluno struct {
	nome         string
	sobrenome    string
	rg           int
	dataAdmissao string
}

func main() {
	a1 := aluno{"Carlos", "Silva", 32211223, "22/05/22"}

	fmt.Println(a1)

	a1.nome = "Pedro"

	fmt.Println(a1)
}
