package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))

	// for i := 0; i < len(responseObject.Pokemon); i++ {
	// 	fmt.Println(responseObject.Pokemon[i].Species.Name)
	// }

	for _, pokemon := range responseObject.Pokemon {
		fmt.Println(pokemon.Species.Name)
	}
}

type Response struct {
	Name		string 		`json:"name"`
	Pokemon []Pokemon	`json:"pokemon_entries"`
}

type Pokemon struct {
	EntryNo	int 						`json:"entry_number"`
	Species PokemonSpecies	`json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name 	string 	`json:"name"`
}