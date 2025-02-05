package main

import (
	"fmt"
)

func main() {
	courses := []string{"Go Fundamentals",
		"Cursinho de Docker",
		"Cursinho de Kubernetes",
		"Outros cursinho de coisa"}

	for _, i := range courses {
		fmt.Println(i)
	}

	sliceOfInt := []int{1, 2, 3, 4, 5}

	for _, i := range sliceOfInt {
		fmt.Println(i)
	}
}
