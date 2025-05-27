package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type AllResponse struct {
	Results []Pokemon `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetPkmns() []string {
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/?offset=0&limit=151")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject AllResponse
	json.Unmarshal(responseData, &responseObject)

	var allPokemon []string
	for i := 0; i < len(responseObject.Results); i++ {
		p := responseObject.Results[i].Name
		allPokemon = append(allPokemon, strings.Title(p))
	}

	return allPokemon
}

func ValidatePokemon(pokemon string) bool {
	list := GetPkmns()
	for _, p := range list {
		if p == pokemon {
			return true
		}
	}
	return false
}
