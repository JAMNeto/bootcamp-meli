package main

import "fmt"

func operation(oper string) (func(valores ...int) int, error) {
	switch oper {
	case "minimum":
		return calculaMinimo, nil
	case "average":
		return calculaMedia, nil
	case "maximum":
		return calculaMaximo, nil
	default:
		return nil, nil
	}

}

func calculaMedia(numeros ...int) int {
	somaNumeros := 0
	totalNumeros := len(numeros)
	for _, numero := range numeros {
		somaNumeros += numero
	}
	media := somaNumeros / totalNumeros
	return media
}

func calculaMinimo(numeros ...int) int {
	menorNumero := numeros[0]
	for _, numero := range numeros {
		if numero < menorNumero {
			menorNumero = numero
		}
	}
	return menorNumero
}

func calculaMaximo(numeros ...int) int {
	maiorNumero := 0
	for _, numero := range numeros {
		if numero > maiorNumero {
			maiorNumero = numero
		}
	}
	return maiorNumero
}

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	minFunc, _ := operation(minimum)
	averageFunc, _ := operation(average)
	maxFunc, _ := operation(maximum)

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 1, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5, 20)
	fmt.Println(minValue)
	fmt.Println(averageValue)
	fmt.Println(maxValue)

}
