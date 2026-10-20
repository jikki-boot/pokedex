package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"git.jaxon.io/jaxon-boot/pokedex/internal/pokeapi"
)

type config struct {
	commands         map[string]cliCommand
	pokeapiClient    pokeapi.Client
	pokedex          map[string]pokeapi.Pokemon
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); !ok {
			log.Fatal("Error reading input")
		}

		input := scanner.Text()
		splitInput := cleanInput(input)
		if len(splitInput) == 0 {
			continue
		}

		commandInput, args := splitInput[0], splitInput[1:]
		if command, exists := cfg.commands[commandInput]; exists {
			if err := command.callback(cfg, args...); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch <pokemon_name>",
			description: "attempts to catch a specified pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore <location-area>",
			description: "Explores a given location on the map",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View Pokedex entry for a specified pokemon",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
	}
}
