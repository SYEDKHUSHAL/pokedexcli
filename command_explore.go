package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	pokemonLocationResp, err := cfg.pokeapiClient.ListLocationsPokemon(args[0])

	if err != nil {
		return err
	}

	for i, enc := range pokemonLocationResp.PokemonEncounters {
		fmt.Printf("%d : %s \n", i, enc.Pokemon.Name)
	}

	return nil
}

// func commandExplore(cfg *config, args ...string) error {
// 	if len(args) != 1 {
// 		return errors.New("you must provide a location name")
// 	}

// 	name := args[0]
// 	location, err := cfg.pokeapiClient.ListLocationsPokemon(name)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("Exploring %s...\n", location.Name)
// 	fmt.Println("Found Pokemon: ")
// 	for _, enc := range location.PokemonEncounters {
// 		fmt.Printf(" - %s\n", enc.Pokemon.Name)
// 	}
// 	return nil
// }