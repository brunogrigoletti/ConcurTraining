package conditionals

import (
	"fmt"
)

func ExercicioIf() {
	//pesoThiago := 95
	//pesoMaeTiano := 300

	if pesoThiago, pesoMaeTiano := 95, 300; pesoThiago < pesoMaeTiano { //da pra declarar as variáveis na linha do if mesmo
		fmt.Println("Thiagão tá sarado hein, bãi") // mas aí elas só valem dentro do if, depois elas vão
	} else if pesoThiago == pesoMaeTiano { // pra banha.
		fmt.Println("A mãe do Tiano tá sarada que nem o Thiago hein bãããã")
	} else {
		fmt.Println("Thiago tá gordão, vai malhar essa raba, gordo!")
	}
}
