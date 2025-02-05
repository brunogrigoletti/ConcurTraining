package main

import (
	"fmt"
)

func main() {
	kgDosGuri := map[string]int{
		"Thiago":       95,
		"Dani":         65,
		"Tiano":        120,
		"Mãe do Tiano": 300,
	}

	fmt.Println(kgDosGuri)

	kgDosGuri["Thiago"] = 85
	kgDosGuri["Mãe do Tiano"] = 500

	fmt.Println(kgDosGuri)

	delete(kgDosGuri, "Mãe do Tiano")

	fmt.Println(kgDosGuri)
}
