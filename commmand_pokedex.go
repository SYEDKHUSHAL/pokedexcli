package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.coughtPokemon) == 0 {
		return errors.New("you have not cought any pokemons yet")	
	}
	fmt.Println("Your Pokedex:")
	for _, value := range cfg.coughtPokemon {
		fmt.Printf("  - %s\n",value.Name)
	}
	
	return nil
}