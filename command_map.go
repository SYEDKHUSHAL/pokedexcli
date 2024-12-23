package main

import (
	"fmt"
	// "errors"
)



func commandMapF(cfg *config, args ...string) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil

}

func commandMapB(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
	}else{
		locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
		if err != nil {
			return err
		}
	
		cfg.nextLocationsURL = locationResp.Next
		cfg.prevLocationsURL = locationResp.Previous
	
		for _, loc := range locationResp.Results {
			fmt.Println(loc.Name)
		}
	}

	return nil
}