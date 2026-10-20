package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for cmdName, cmd := range cfg.commands {
		fmt.Printf("%s: %s\n", cmdName, cmd.description)
	}

	return nil
}
