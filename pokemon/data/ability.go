package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type PokemonResponse struct {
	Moves []string `json:"moves"`
}

type Moves struct {
	Moves []struct {
		Move Move `json:"move"`
	} `json:"moves"`
}

type Move struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Power struct {
	Power *int `json:"power"`
}

func GetMoves(pkmn string) map[string]*int {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pkmn)
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Moves
	json.Unmarshal(responseData, &responseObject)

	moves := make(map[string]*int)
	for _, moveEntry := range responseObject.Moves {
		moveName := moveEntry.Move.Name
		moveURL := moveEntry.Move.URL

		moveResp, err := http.Get(moveURL)
		if err != nil {
			moves[moveName] = nil
			continue
		}
		defer moveResp.Body.Close()

		moveData, err := ioutil.ReadAll(moveResp.Body)
		if err != nil {
			moves[moveName] = nil
			continue
		}

		var pwr Power
		json.Unmarshal(moveData, &pwr)

		moves[moveName] = pwr.Power
	}

	return moves
}
