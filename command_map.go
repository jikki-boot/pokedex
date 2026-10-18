package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	if cfg.nextLocationsURL == nil {
		return errors.New("you're on the last page")
	}

	if err := printAreas(cfg, cfg.nextLocationsURL); err != nil {
		return err
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	if err := printAreas(cfg, cfg.prevLocationsURL); err != nil {
		return err
	}

	return nil
}

func printAreas(cfg *config, url *string) error {
	resource, err := cfg.pokeapiClient.GetLocationAreas(url)
	if err != nil {
		return err
	}

	for _, area := range resource.Results {
		fmt.Println(area.Name)
	}

	cfg.nextLocationsURL = resource.Next
	cfg.prevLocationsURL = resource.Previous

	return nil
}
