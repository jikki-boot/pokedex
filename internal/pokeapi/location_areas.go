package pokeapi

import (
	"encoding/json"
	"fmt"
)

const (
	LocationAreasBaseURL = baseURL + "/location-area"
)

func (c *Client) GetLocationAreas(pageURL *string) (Resource, error) {
	url := LocationAreasBaseURL
	if pageURL != nil {
		url = *pageURL
	}

	data, err := c.Get(url)
	if err != nil {
		return Resource{}, err
	}

	var resource Resource
	if err := json.Unmarshal(data, &resource); err != nil {
		return Resource{}, err
	}

	return resource, nil
}

func (c *Client) GetPokemon(area string) ([]Pokemon, error) {
	url := fmt.Sprintf("%s/%s", LocationAreasBaseURL, area)

	data, err := c.Get(url)
	if err != nil {
		return []Pokemon{}, nil
	}

	var locationArea LocationArea
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return []Pokemon{}, err
	}

	pokemon := make([]Pokemon, len(locationArea.PokemonEncounters))
	for i, encounter := range locationArea.PokemonEncounters {
		pokemon[i] = encounter.Pokemon
	}

	return pokemon, nil
}
