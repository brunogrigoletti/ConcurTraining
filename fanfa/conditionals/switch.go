package main

import (
	"fmt"
)

func main() {
	name := "Thiago"

	switch name {
	case "thiago":
		fmt.Println("Case 1. thiago com 't' minúsculo.")
	case "Thiago":
		fmt.Println("Case 2. Acertou! Thiago com 'T' maiúsculo!")
		fallthrough //fallthrough faz com que o próximo "case" seja exectuado também. O default tb.
	case "Thiagão":
		fmt.Println("Case 3. Thiagão é tri, mas o nome é 'Thiago' mesmo.")
	case "Ricardo":
		fmt.Println("Case 4. Quem raios é Ricardo?!")
	default:
		fmt.Println("Chegou no Default.")
	}
}
