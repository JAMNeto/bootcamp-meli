// Exercício 1 - Rede social

// Uma empresa de mídia social precisa implementar uma estrutura de usuários com funções
// que acrescentem informações à estrutura. Para otimizar e economizar memória, eles exigem
// que a estrutura de usuários ocupe o mesmo lugar na memória para o programa principal e
// para as funções:
// - A estrutura deve possuir os seguintes campos: Nome, Sobrenome, idade, email e
// senha
// E devem implementar as seguintes funcionalidades:
// - mudar o nome: me permite mudar o nome e sobrenome
// - mudar a idade: me permite mudar a idade
// - mudar e-mail: me permite mudar o e-mail
// - mudar a senha: me permite mudar a senha

package main

import "fmt"

const (
	nome      = "nome"
	sobrenome = "sobrenome"
	idade     = "idade"
	email     = "email"
	senha     = "senha"
)

type usuario struct {
	nome      string
	sobrenome string
	idade     int
	email     string
	senha     string
}

func (u *usuario) mudaNome(dado string) {
	u.nome = dado
}
func (u *usuario) mudaSobrenome(dado string) {
	u.sobrenome = dado
}
func (u *usuario) mudaIdade(dado int) {
	u.idade = dado
}
func (u *usuario) mudaEmail(dado string) {
	u.email = dado
}
func (u *usuario) mudaSenha(dado string) {
	u.senha = dado
}

func main() {
	u1 := usuario{"Carlos", "Andrade", 30, "carlos@bol.com", "senha123"}
	fmt.Println(u1)
	u1.mudaNome("Eduardo")
	u1.mudaSobrenome("Alves")
	fmt.Println(u1)

}
