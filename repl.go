package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/SYEDKHUSHAL/pokedexcli/internal/pokeapi"
)


type config struct {
	pokeapiClient			pokeapi.Client
	nextLocationsURL		*string
	prevLocationsURL 		*string
	coughtPokemon			map[string]pokeapi.Pokemon
}


func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		args := []string{}

		if len(words) > 1 {
			args = words[1:]
		}
		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback(cfg, args...)

			if err != nil {
				fmt.Println(err)
			}
			continue
		}

		fmt.Println("Unknown command")
		continue
	}
}


type cliCommand struct {
	name 			string
	description 	string
	callback		func(*config, ...string) error
}


func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	cleanText := strings.TrimSpace(lower)
	result := strings.Fields(cleanText)
	return result
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: 			"exit",
			description: 	"Exit the Pokedex",
			callback: 		commandExit,
		},
		"help": {
			name: 			"help",
			description:	"Displays a help message",
			callback: 		commandHelp,
		},
		"map": {
			name: 			"mapf",
			description:	"Displays the next 20 pokemon locations",
			callback: 		commandMapF,
		},
		"mapb": {
			name: 			"mapb",
			description:	"Displays the previous 20 pokemon locations",
			callback: 		commandMapB,
		},
		"explore": {
			name: 			"explore <location_name>",
			description:	"Displays the pokemons found in a location",
			callback: 		commandExplore,
		},
		"catch": {
			name: 			"catch <pokemon>",
			description:    "attempts to catch a pokemon",
			callback: 		commandCatch,
		},
		"inspect": {
			name: 			"inspect <pokemon>",
			description:    "inspects an already caught pokemon",
			callback: 		commandInspect,
		},
		"pokedex": {
			name: 			"pokedex",
			description:    "prints the list of cought pokemons",
			callback: 		commandPokedex,
		},
	}
}
