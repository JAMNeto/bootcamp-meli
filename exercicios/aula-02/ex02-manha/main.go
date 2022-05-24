package main

import (
	"errors"
	"fmt"
)

func calculaMedia(notas ...float64) (float64, error) {
	somaNotas := float64(0)
	totalNotas := float64(len(notas))
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("A nota não pode ser negativa\n")
		} else {
			somaNotas += nota
		}
	}
	media := somaNotas / totalNotas
	return media, nil
}

func main() {
	media, err := calculaMedia(6, -7, 5, 6)

	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("Sua média é %.2f\n", media)
	}

}
