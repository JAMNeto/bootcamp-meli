// Exercício 2 - E-commerce (Parte II)
// Uma grande empresa de vendas na web precisa adicionar funcionalidades para adicionar
// produtos aos usuários. Para fazer isso, eles exigem que usuários e produtos tenham o
// mesmo endereço de memória no main do programa e nas funções.

// Estruturas necessárias:
// - Usuário: Nome, Sobrenome, E-mail, Produtos (array de produtos).
// - Produto: Nome, preço, quantidade.
// Algumas funções necessárias:
// - Novo produto: recebe nome e preço, e retorna um produto.
// - Adicionar produto: recebe usuário, produto e quantidade, não retorna nada, adiciona
// o produto ao usuário.
// - Deletar produtos: recebe um usuário, apaga os produtos do usuário.

package main

import "fmt"

type usuario struct {
	nome, sobrenome, email string
	produtos               []produto
}

type produto struct {
	nome       string
	preco      float64
	quantidade int
}

func novoProduto(nome string, preco float64) produto {
	novoProduto := produto{nome, preco, 0}
	return novoProduto
}

func (u *usuario) adicionaProduto(produto produto, quantidade int) {
	produto.quantidade = quantidade
	u.produtos = append(u.produtos, produto)
}

func (u *usuario) deletarProdutos(usuario usuario) {
	u.produtos = nil
}

func main() {
	u1 := usuario{"Carlos", "Andrade", "carlosandrae@email.com", []produto{}}
	prod1 := novoProduto("Caneta", 1.50)
	prod2 := novoProduto("Mesa", 250.0)
	fmt.Println(u1)
	fmt.Println(prod1)
	u1.adicionaProduto(prod1, 10)
	fmt.Println(u1)
	u1.adicionaProduto(prod2, 5)
	fmt.Println(u1)
	u1.deletarProdutos(u1)
	fmt.Println(u1)
	u1.adicionaProduto(prod2, 5)
	fmt.Println(u1)
}
