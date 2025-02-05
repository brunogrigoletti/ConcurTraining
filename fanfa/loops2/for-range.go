package main

import (
	"fmt"
)

func main() {
	coursesInProg := []string{
		"Curso de Go para Iniciantes",
		"Curso de APIs",
		"Curso de AWS",
		"Cursinho de alguma outra coisa"}
	coursesCompleted := []string{
		"Curso de Go para Iniciantes",
		"Cursinho de alguma outra coisa"}

	for _, i := range coursesInProg {
		completed := false
		for _, j := range coursesCompleted {
			if i == j {
				fmt.Println(i, " - COMPLETED.")
				completed = true
				break
			}
		}
		if !completed {
			fmt.Println(i)
		}
	}
}
