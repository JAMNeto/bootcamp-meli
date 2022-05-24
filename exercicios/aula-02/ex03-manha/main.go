package main

import "fmt"

func calculaSalario(categoria string, horas float64) (salario float64) {
	switch categoria {
	case "c":
		salario = horas * 1000
	case "b":
		salario = horas * 1500
		if horas >= 160 {
			salario += (salario * 0.2)
		}
	case "a":
		salario = horas * 3000
		if horas >= 160 {
			salario += (salario * 0.5)
		}
	}
	return salario
}

func main() {
	fmt.Println("O salário de chico é:", calculaSalario("a", 160))
}
