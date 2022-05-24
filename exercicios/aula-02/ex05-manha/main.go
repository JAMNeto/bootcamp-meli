package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func Animal(animal string) (func(qtd int) int, error) {
	switch animal {
	case dog:
		return dogFunction, nil
	case cat:
		return catFunction, nil
	case hamster:
		return hamsterFunction, nil
	case tarantula:
		return tarantulaFunction, nil
	default:
		return nil, errors.New("Animal não cadastrado")
	}
}

func dogFunction(qtd int) int {
	foodNeeds := qtd * 10000
	return foodNeeds
}

func catFunction(qtd int) int {
	foodNeeds := qtd * 5000
	return foodNeeds
}

func hamsterFunction(qtd int) int {
	foodNeeds := qtd * 250
	return foodNeeds
}

func tarantulaFunction(qtd int) int {
	foodNeeds := qtd * 150
	return foodNeeds
}

func main() {
	dogNeeds, err := Animal(dog)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d cachorros precisam de %d gramas de ração", 10, dogNeeds(10))
}
