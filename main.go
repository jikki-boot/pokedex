package main

import (
	"time"

	"git.jaxon.io/jaxon-boot/pokedex/internal/pokeapi"
)

func main() {
	timeout := 5 * time.Second
	cacheInterval := 30 * time.Minute
	pokeapiClient := pokeapi.NewClient(timeout, cacheInterval)
	locationsBaseURL := pokeapi.LocationAreasBaseURL

	cfg := &config{
		commands:         getCommands(),
		pokeapiClient:    pokeapiClient,
		pokedex:          make(map[string]pokeapi.Pokemon),
		nextLocationsURL: &locationsBaseURL,
		prevLocationsURL: nil,
	}

	startRepl(cfg)
}
