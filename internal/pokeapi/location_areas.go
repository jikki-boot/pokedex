package pokeapi

import (
	"encoding/json"
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
