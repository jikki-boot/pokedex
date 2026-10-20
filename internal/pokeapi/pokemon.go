package pokeapi

import (
	"encoding/json"
	"fmt"
)

const (
	PokemonBaseURL = baseURL + "/pokemon"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := fmt.Sprintf("%s/%s", PokemonBaseURL, name)

	data, err := c.Get(url)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
