package app

import (
	"fmt"
	"pokemon/data"
)

func Run() {
	terminal()
}

func terminal() {
	var p1 string
	for {
		fmt.Print("First Pokémon: ")
		fmt.Scanln(&p1)
		if validatePokemon(p1) {
			break
		}
	}

	var p2 string
	for {
		fmt.Print("Second Pokémon: ")
		fmt.Scanln(&p2)
		if validatePokemon(p2) {
			break
		}
	}

	fmt.Print("Winner: ")
}

func validatePokemon(pokemon string) bool {
	list := data.GetAllPokemon()
	for _, p := range list {
		if p == pokemon {
			return true
		}
	}
	fmt.Println("Invalid Pokémon name. Please try again.")
	return false
}
