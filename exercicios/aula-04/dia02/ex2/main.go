// Exercício 2 - Registrando clientes

// O mesmo escritório do exercício anterior, solicita uma funcionalidade para poder
// cadastrar dados de novos clientes. Os dados necessários para cadastrar um cliente são:
// - Arquivo
// - Nome e sobrenome
// - RG
// - Número de telefone
// - Endereço

// ● Tarefa 1: O número do arquivo deve ser atribuído ou gerado separadamente e antes
// da cobrança das despesas restantes. Desenvolva e implemente uma função para
// gerar um ID que você usará posteriormente para atribuí-lo como um valor a “Arquivo”.
// Se por algum motivo esta função retornar o valor "nil", deve gerar um panic que
// interrompa a execução e aborte

// Tarefa 2: Antes de cadastrar um cliente, você deve verificar se o cliente já existe. Para
// fazer isso, você precisa ler os dados de um arquivo .txt. Em algum lugar do seu
// código, implemente a função para ler um arquivo chamado “customers.txt” (como no
// exercício anterior, este arquivo não existe, então qualquer função que tente lê-lo
// retornará um erro). Você deve lidar adequadamente com esse erro, como vimos até
// agora. Esse erro deve:

// 1. Gerar um panic;
// 2. Imprimir no console a mensagem: “erro: o arquivo indicado não foi encontrado ou
// está danificado”, e continuar com a execução do programa normalmente.

// Requisitos gerais:
// ● Use recover para recuperar o valor dos panics que podem surgir (exceto na tarefa 1).
// ● Lembre-se de realizar as validações necessárias para cada retorno que possa conter
// um valor de erro (por exemplo, aqueles que tentam ler arquivos).
// Crie um erro, personalizando-o ao seu gosto, utilizando qualquer uma das funções
// que o GO disponibiliza para ele (ele também deve realizar a validação relevante para
// o caso de erro retornado).

package main

import (
	"fmt"
	"os"
)

type cliente struct {
	arquivo         int
	nome, sobrenome string
	rg, telefone    int
	endereco        string
}

func idGenerator(clientes []cliente) int {
	newId := len(clientes) + 100
	return newId
}

func cadastrarCliente(nome, sobrenome, endereco string, id, rg, telefone int, clientes []cliente) cliente {
	defer recuperarExecucao()
	if id < 0 {
		panic("Valor inválido de ID")
	}
	cliente := cliente{id, nome, sobrenome, rg, telefone, endereco}
	return cliente
}

func lerArquivo(path string) []byte {
	defer recuperarExecucao()
	arquivo, err := os.ReadFile("customers.txt")
	if err != nil {
		panic("erro: o arquivo indicado não foi encontrado ou está danificado")
	}
	return arquivo
}

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Recuperei a execucao")
	}
}

func main() {
	clientes := make([]cliente, 0)
	id1 := idGenerator(clientes)
	cliente := cadastrarCliente("Carlos", "Souza", "AV", id1, 1233, 3333, clientes)
	clientes = append(clientes, cliente)
	id2 := idGenerator(clientes)
	cliente2 := cadastrarCliente("Augusto", "Mega", "Rua", id2, 4444, 555, clientes)
	clientes = append(clientes, cliente2)
	arquivo := lerArquivo("customers.txt")
	fmt.Println(clientes)
	fmt.Println(cliente)
	fmt.Println(arquivo)
}
