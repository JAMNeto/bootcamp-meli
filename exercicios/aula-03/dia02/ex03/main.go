// Uma empresa nacional é responsável pela venda de produtos, serviços e manutenção.
// Para isso, eles precisam realizar um programa que seja responsável por calcular o preço total
// dos Produtos, Serviços e Manutenção. Devido à forte demanda e para otimizar a velocidade,
// eles exigem que o cálculo da soma seja realizado em paralelo por meio de 3 go routines.

// Precisamos de 3 estruturas:
// - Produtos: nome, preço, quantidade.
// - Serviços: nome, preço, minutos trabalhados.
// - Manutenção: nome, preço.
// Precisamos de 3 funções:
// - Somar Produtos: recebe um array de produto e devolve o preço total (preço *
// quantidade).
// - Somar Serviços: recebe uma array de serviço e devolve o preço total (preço * média
// hora trabalhada, se ele não vier trabalhar por 30 minutos, ele será cobrado como se
// tivesse trabalhado meia hora).
// - Somar Manutenção: recebe um array de manutenção e devolve o preço total.

// Os 3 devem ser executados concomitantemente e ao final o valor final deve ser mostrado na
// tela (somando o total dos 3).

package main

import (
	"fmt"
	"time"
)

type produto struct {
	nome       string
	preco      float64
	quantidade int
}

type servico struct {
	nome               string
	preco              float64
	minutosTrabalhados int
}

type manutencao struct {
	nome  string
	preco float64
}

func somarProdutos(produtos []produto) float64 {
	precoProd := 0.0
	for _, prod := range produtos {
		precoProdutos := prod.preco * float64(prod.quantidade)
		precoProd += precoProdutos
	}
	fmt.Println("1 - Da funcao soma produtos", precoProd)
	return precoProd
}

func somarServicos(servicos []servico) float64 {
	precoServ := 0.0
	for _, serv := range servicos {
		precoServicos := serv.preco * float64(serv.minutosTrabalhados)
		precoServ += precoServicos
	}
	fmt.Println("2 - Da funçao soma serviços", precoServ)
	return precoServ
}

func somarManutencao(manutencoes []manutencao) float64 {
	precoManut := 0.0
	for _, manut := range manutencoes {
		precoManut += manut.preco
	}
	fmt.Println("3 - Da funcao soma manutencoes", precoManut)
	return precoManut
}

func retornarTotal(produtos []produto, servicos []servico, manutencoes []manutencao) float64 {
	somaProdutos := somarProdutos(produtos)
	// fmt.Println(somaProdutos)
	somaServicos := somarServicos(servicos)
	// fmt.Println(somaServicos)
	somaManutencao := somarManutencao(manutencoes)
	// fmt.Println(somaManutencao)
	resultado := (somaProdutos + somaServicos) + somaManutencao
	fmt.Println("4 - Da funcao total da soma", resultado)
	return resultado
}

func main() {
	p1 := produto{"Mesa", 250.0, 10}
	p2 := produto{"Cadeira", 150.50, 30}
	p3 := produto{"Pratos", 25.0, 100}
	p4 := produto{"Sofá", 700.0, 6}
	prods := []produto{p1, p2, p3, p4}
	// fmt.Println(prods)
	s1 := servico{"Consertar pia", 30.0, 30}
	s2 := servico{"Apertar portas", 50.0, 90}
	servs := []servico{s1, s2}
	// fmt.Println(servs)
	m1 := manutencao{"Trocar óleo", 180.0}
	m2 := manutencao{"Balanceamento", 80.0}
	manuts := []manutencao{m1, m2}
	// fmt.Println(manuts)
	for {
		go somarManutencao(manuts)
		go somarProdutos(prods)
		go somarServicos(servs)
		go retornarTotal(prods, servs, manuts)
		fmt.Println("---------------------------")
		time.Sleep(5 * time.Second)

	}
}
