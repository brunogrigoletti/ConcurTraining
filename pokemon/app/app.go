package app

/*
Based on 'How To Consume Data From A REST HTTP API With Go', by Elliot Forbes
https://tutorialedge.net/golang/consuming-restful-api-with-go/
*/

import (
	"fmt"
	"pokemon/data"
	"strings"
)

func Run() {
	terminal()
}

func terminal() {
	var p1 string
	var m1 string
	var pwr1 int
	for {
		fmt.Print("First Pokémon: ")
		fmt.Scanln(&p1)
		if data.ValidatePokemon(p1) {
			fmt.Print(p1 + "'s abilities:\n")
			moves := data.GetMoves(p1)
			for move, power := range moves {
				moveName := strings.ToUpper(string(move[0])) + move[1:]
				if power != nil {
					fmt.Printf("- %s (Power: %d)\n", moveName, *power)
				}
			}
			for {
				fmt.Print("Select a move: ")
				fmt.Scanln(&m1)
				if len(m1) > 0 {
					m1Key := string(m1[0]+32) + m1[1:]
					if power, exists := moves[m1Key]; exists {
						if power != nil {
							pwr1 = *power
						} else {
							pwr1 = 0
						}
						goto firstPokemonSelected
					}
				}
				fmt.Println("Invalid move. Please try again.")
			}
		} else {
			fmt.Println("Invalid Pokémon name. Please try again.")
		}
	}
firstPokemonSelected:

	var p2 string
	var m2 string
	var pwr2 int
	for {
		fmt.Print("Second Pokémon: ")
		fmt.Scanln(&p2)
		if data.ValidatePokemon(p1) {
			fmt.Print(p2 + "'s abilities:\n")
			moves := data.GetMoves(p2)
			for move, power := range moves {
				moveName := strings.ToUpper(string(move[0])) + move[1:]
				if power != nil {
					fmt.Printf("- %s (Power: %d)\n", moveName, *power)
				}
			}
			for {
				fmt.Print("Select a move: ")
				fmt.Scanln(&m2)
				if len(m2) > 0 {
					m2Key := string(m2[0]+32) + m2[1:]
					if power, exists := moves[m2Key]; exists {
						if power != nil {
							pwr2 = *power
						} else {
							pwr2 = 0
						}
						goto secondPokemonSelected
					}
				}
				fmt.Println("Invalid move. Please try again.")
			}
		} else {
			fmt.Println("Invalid Pokémon name. Please try again.")
		}
	}
secondPokemonSelected:

	comparePowers(p1, m1, pwr1, p2, m2, pwr2)
}

func comparePowers(p1, m1 string, pwr1 int, p2, m2 string, pwr2 int) {
	fmt.Printf("%s used %s (Power: %d) and %s used %s (Power: %d).\n",
		strings.Title(p1), strings.Title(m1), pwr1,
		strings.Title(p2), strings.Title(m2), pwr2)
	if pwr1 > pwr2 {
		fmt.Printf("%s won!\n", strings.Title(p1))
	} else if pwr2 > pwr1 {
		fmt.Printf("%s won!\n", strings.Title(p2))
	} else {
		fmt.Println("Draw!")
	}
}
