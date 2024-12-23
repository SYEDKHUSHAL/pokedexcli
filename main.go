package main

import (
	"time"

	"github.com/SYEDKHUSHAL/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute) 

	cfg := &config{
		pokeapiClient: pokeClient,
		coughtPokemon: map[string]pokeapi.Pokemon{},
	}


	startRepl(cfg)
}