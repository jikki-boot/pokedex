package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type config struct {
	commands map[string]cliCommand
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

		commandInput := splitInput[0]
		command, ok := cfg.commands[commandInput]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		command.callback(cfg)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
