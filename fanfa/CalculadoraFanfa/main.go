package main

import (
	"fmt"

	"github.com/fanfathi/CalculadoraFanfa/calculos"
)

func main() {
	var resultado int
	var num1, num2 int
	var operador string

	fmt.Println("Insira um número de cada vez.")
	fmt.Println("Você pode continuar fazendo operações no resultado.")

	fmt.Println("Insira o primeiro número: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Insira um número válido!")
	}

	for {
		fmt.Println("Insira o operador: ")
		fmt.Scan(&operador)

		fmt.Println("Insira o próximo número: ")
		_, err = fmt.Scan(&num2)
		if err != nil {
			fmt.Println("Insira um número válido!")
			continue
		}

		switch operador {
		case "+":
			resultado = calculos.Soma(num1, num2)
		case "-":
			resultado = calculos.Subtracao(num1, num2)
		case "*":
			resultado = calculos.Multiplicacao(num1, num2)
		case "/":
			resultado = calculos.Divisao(num1, num2)
		default:
			fmt.Println("Operação Inválida. Tente novamente.")
		}
		fmt.Printf("O resultado é: %d \n", resultado)

		fmt.Println("Continuar o cálculo usando o resultado? (y/n)")
		var continua string
		fmt.Scan(&continua)
		if continua == "y" || continua == "Y" {
			num1 = resultado
		} else {
			break
		}
	}
	fmt.Printf("O resultado final é: %d\n", resultado)

}
