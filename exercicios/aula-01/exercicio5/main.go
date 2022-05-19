package main

import "fmt"

func main() {

	var word string

	word = "Testando"

	fmt.Printf("A palavra \"%s\" tem %d letras\n", word, len(word))

	for i := 0; i < len(word); i++ {
		fmt.Printf("%c\n", word[i])
	}
}
