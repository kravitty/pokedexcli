package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kravitty/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := GetCommands()[commandName]
		if exists {
			if commandName == "explore" && len(words) > 1 {
				err := command.callback(cfg, words[1])
				if err != nil {
					fmt.Println(err)
				}
			} else {
				err := command.callback(cfg)
				if err != nil {
					fmt.Println(err)
				}
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"explore": {
			name:        "explore",
			description: "Explore <location> for available Pokemon",
			callback:    func(cfg *config, args ...string) error { return commandExplore(cfg, args[0]) },
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func(cfg *config, args ...string) error { return commandHelp(cfg) },
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    func(cfg *config, args ...string) error { return commandMapf(cfg) },
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    func(cfg *config, args ...string) error { return commandMapb(cfg) },
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    func(cfg *config, args ...string) error { return commandExit(cfg) },
		},
	}
}
