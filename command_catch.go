package main

import (
	"errors"
	"fmt"
	"math/rand"
)


func catchPokemon(baseXP int, scalingFactor float64, minProb float64) bool {
	if baseXP < 0 {
		baseXP = 0
	}

	catchProb := scalingFactor / float64(1 + baseXP)

	if catchProb > 1 {
		catchProb = 1
	}

	if catchProb < minProb {
		catchProb = minProb
	}

	chance := rand.Float64()
	
	return chance <= catchProb
}



func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you need to provide one pokemon name")
	}

	name := args[0]
	pokemonResp , err := cfg.pokeapiClient.GetPokemon(name)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	baseExp := pokemonResp.BaseExperience

	success := catchPokemon(baseExp, 100.00, 0.05)

	if success {
		cfg.coughtPokemon[name] = pokemonResp
		fmt.Printf("%s was caught!\n", name)
	}else {
		fmt.Printf("%s escaped!\n", name)
	}
	
	return nil
}


