package main

import (
	"time"

	"git.jaxon.io/jaxon-boot/pokedex/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5 * time.Second)
	locationsBaseURL := pokeapi.LocationAreasBaseURL

	cfg := &config{
		commands:         getCommands(),
		pokeapiClient:    pokeapiClient,
		nextLocationsURL: &locationsBaseURL,
		prevLocationsURL: nil,
	}

	startRepl(cfg)
}
