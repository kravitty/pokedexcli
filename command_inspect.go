package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	if caughtPokemon, exists := cfg.caughtPokemon[pokemonName]; exists {
		fmt.Println("Name:", caughtPokemon.Name)
		fmt.Println("Height:", caughtPokemon.Height)
		fmt.Println("Weight:", caughtPokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range caughtPokemon.Stats {
			fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typeInfo := range caughtPokemon.Types {
			fmt.Println("  -", typeInfo.Type.Name)
		}
		return nil
	} else {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

}
