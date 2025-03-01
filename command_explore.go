package main

import (
	"fmt"
)

func commandExplore(cfg *config, location string) error {
	fmt.Printf("Exploring " + location + "...\n")

	locationDetails, err := cfg.pokeapiClient.GetLocation(location)
	if err != nil {
		fmt.Println("Error getting location:", err)
		return err
	}
	for _, val := range locationDetails.PokemonEncounters {
		fmt.Println(val.Pokemon.Name)
	}
	//TODO: Implement Cache for location details
	return nil
}
