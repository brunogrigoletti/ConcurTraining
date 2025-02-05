package main

import (
	"fmt"
)

func main() {
	type pessoa struct {
		nome        string
		altura      float64
		corDoOlho   string
		corDoCabelo string
	}
	//da pra declarar uma variável do tipo "pessoa" das seguintes maneiras:
	//jeito 1: var thiagaoDaMassa pessoa
	//jeito 2: thiagaoDaMassa := new(pessoa)
	//nesses dois jeitos, a variável é inicializada com todos os valores 0 ou nulos

	thiagaoDaMassa := pessoa{
		nome:        "Thiago",
		altura:      1.75,
		corDoOlho:   "Azul",
		corDoCabelo: "Sei lá, preto",
	}
	fmt.Println(thiagaoDaMassa)
	fmt.Println(thiagaoDaMassa.corDoCabelo)
}
