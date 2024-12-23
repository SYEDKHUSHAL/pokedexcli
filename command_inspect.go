package main

import (
	"errors"
	"fmt"
)


func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	
	name := args[0]

	if pokemon, ok := cfg.coughtPokemon[name]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, item := range pokemon.Stats{
			fmt.Printf("  - %s: %d\n", item.Stat.Name, item.Effort)
		}

		fmt.Println("Types:")
		for _, item := range pokemon.Types{
			fmt.Printf("  - %s\n", item.Type.Name)
		}
	}else{
		return errors.New("you must have cought that pokemon")
	}

	return nil
}