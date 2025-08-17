// O desafio era ler e printar o JSON como uma lista.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Struct com as informações de cada hotel.
type Hotel struct {
	Name      string  `json:"name"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	Country   string  `json:"country"`
	Amenities string  `json:"amenities"`
	Rating    float64 `json:"rating"`
	Reviews   string  `json:"reviews"`
}

func main() {

	jsonFile, err := os.Open("Input.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	// Lê o jsonFile aberto como um array de bytes
	byteValue, _ := io.ReadAll(jsonFile)

	//inicializa o array de hotéis
	var hoteis []Hotel

	// agora a variável hoteis contém o byteValue(array de bytes com todos os hotéis)
	json.Unmarshal(byteValue, &hoteis)

	for _, h := range hoteis {
		fmt.Println("Hotel Name: " + h.Name)
		fmt.Println("Hotel City: " + h.City)
		fmt.Println("Hotel State: " + h.State)
		fmt.Println("Hotel Country: " + h.Country)
		fmt.Println("Hotel Amenities: " + h.Amenities)
		fmt.Println("Hotel Rating: " + strconv.Itoa(int(h.Rating)))
		//teve que converter pq não da pra concatenar string com float64
		fmt.Println("Hotel Reviews: " + h.Reviews)
		fmt.Println("--------------------------------")
		fmt.Println()
	}
	//se quiser printar a lista feia, só usar:
	// fmt.Println(hoteis)
}
