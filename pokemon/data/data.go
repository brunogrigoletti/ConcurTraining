package data

import (
	"fmt"
	"strings"

	"github.com/JoshGuarino/PokeGo/pkg/resources/pokemon"
)

func GetAllPokemon() []string {
	pokemonGroup := pokemon.NewPokemonGroup()

	pokemonList, err := pokemonGroup.GetPokemonList(151, 0)
	if err != nil {
		fmt.Println("Error fetching Pokemon list:", err)
		return nil
	}

	var allPokemon []string

	for _, name := range pokemonList.Results {
		p, err := pokemonGroup.GetPokemon(name.Name)
		if err != nil {
			fmt.Println("Error fetching Pokemon:", err)
			continue
		}
		allPokemon = append(allPokemon, strings.ToUpper(string(p.Name[0]))+p.Name[1:])
	}
	return allPokemon
}
