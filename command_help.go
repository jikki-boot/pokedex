package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for cmdName, cmd := range cfg.commands {
		fmt.Printf("%s: %s\n", cmdName, cmd.description)
	}

	return nil
}
