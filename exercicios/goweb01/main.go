package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type Produto struct {
	Id            int
	Nome          string
	Cor           string
	Preco         float64
	Estoque       int
	Codigo        string
	Poblicacao    bool
	DataDeCriacao string
}

type Produtos struct {
	Produtos []Produto
}

func getAll(c *gin.Context) {
	jsonFile, err := ioutil.ReadFile("theme.json")
	if err != nil {
		fmt.Println(err)
	}
	var Produtos Produtos
	err2 := json.Unmarshal(jsonFile, &Produtos)
	if err2 != nil {
		fmt.Println(err2)
	}
	c.JSON(200, Produtos)
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá Neto!",
		})
	})

	router.GET("/products", getAll)

	router.Run()
}
