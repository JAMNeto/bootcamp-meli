package main

import (
	"errors"
	"fmt"
)

const (
	pequeno = "pequeno"
	medio   = "medio"
	grande  = "grande"
)

type loja struct {
	produtos []produto
}

type produto struct {
	tipo  string
	nome  string
	preco float64
}

type Produto interface {
	CalcularCusto()
}

type Ecommerce interface {
	Total()
	Adicionar()
}

func (p produto) CalcularCusto() (float64, error) {
	switch p.tipo {
	case "pequeno":
		return p.preco, nil
	case "medio":
		return p.preco * 1.03, nil
	case "grande":
		return (p.preco * 1.06) + 2500, nil
	default:
		return 0, errors.New("Tipo inválido de produto")
	}
}

func (l loja) Total() float64 {
	total := 0.0

	for _, prd := range l.produtos {
		if preco, err := prd.CalcularCusto(); err != nil {
			fmt.Println(err)
		} else {
			total += preco
		}
	}
	return total
}

func (l *loja) Adicionar(p produto) {
	l.produtos = append(l.produtos, p)
}

func novoProduto(tipo, nome string, preco float64) produto {
	novoProd := produto{tipo, nome, preco}
	return novoProd
}

func novaLoja() loja {
	novaLo := loja{}
	return novaLo
}

func main() {
	caneta := novoProduto(pequeno, "caneta", 1.50)
	fmt.Println(caneta)
	cadeira := novoProduto(medio, "cadeira", 250.00)
	fmt.Println(cadeira)
	mesa := novoProduto(grande, "mesa", 550.00)
	fmt.Println(mesa)
	lojaX := novaLoja()
	lojaX.Adicionar(caneta)
	lojaX.Adicionar(cadeira)
	lojaX.Adicionar(mesa)
	total := lojaX.Total()
	fmt.Println(lojaX)
	fmt.Println(total)
}
