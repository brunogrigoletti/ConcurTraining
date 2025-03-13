package calculos

import "fmt"

func Soma(num1 int, num2 int) int {
	return num1 + num2
}

func Subtracao(num1 int, num2 int) int {
	return num1 - num2
}

func Multiplicacao(num1 int, num2 int) int {
	return num1 * num2
}

func Divisao(num1 int, num2 int) int {
	for {
		if num2 == 0 { //a IA sugeriu fazer essa validação nos
			fmt.Println("Não pode dividir por zero!") // switch do main. Eu achei melhor fazer aqui,
			break                                     // pq daí só precisa fazer uma vez.
		} else {
			return num1 / num2
		}
	}
	return 0

}
