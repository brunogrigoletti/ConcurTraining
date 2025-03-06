package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fanfathi/CalculadoraFanfa/calculos"
)

func main() {
	var resultado int

	fmt.Println("Insira a operação no formato --> 2+2")
	fmt.Println("Você pode continuar fazendo operações no resultado inserindo somente o operador e número no formato --> - 1")
	fmt.Println("Para sair, digite: DEU!")

	for {
		var input string //tem que ser string por causa do operador
		fmt.Scanln(&input)

		cadaDigitoDaConta := strings.Split(input, " ") //eu tentei usar o espaço para dividir, mas não tava rolando

		if len(cadaDigitoDaConta) == 3 {
			//primeira conta
			num1, _ := strconv.Atoi(cadaDigitoDaConta[0])
			num2, _ := strconv.Atoi(cadaDigitoDaConta[2])

			switch cadaDigitoDaConta[1] {
			case "+":
				resultado = calculos.Soma(num1, num2)
			case "-":
				resultado = calculos.Subtracao(num1, num2)
			case "*":
				resultado = calculos.Multiplicacao(num1, num2)
			case "/":
				resultado = calculos.Divisao(num1, num2)
			default:
				fmt.Println("Operação Inválida. Tenta de novo, zé mané.")
			}
		} else if len(cadaDigitoDaConta) == 2 {
			// segunda+ conta
			numero, _ := strconv.Atoi(cadaDigitoDaConta[1])

			switch cadaDigitoDaConta[0] {
			case "+":
				resultado = calculos.Soma(resultado, numero)
			case "-":
				resultado = calculos.Subtracao(resultado, numero)
			case "*":
				resultado = calculos.Multiplicacao(resultado, numero)
			case "/":
				resultado = calculos.Divisao(resultado, numero)
			default:
				fmt.Println("Operação Inválida. Tenta mais uma vez...")
			}
		} else if input == "DEU!" {
			break
		} else {
			fmt.Println("Formato inválido. Tenta de novo, zé.")
			continue
		}
		fmt.Printf("O resultado é: %d ", resultado)
	}
	fmt.Printf("O resultado final é: %d", resultado)

}
