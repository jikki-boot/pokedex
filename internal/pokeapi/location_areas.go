package pokeapi

import (
	"encoding/json"
	"net/http"
)

const (
	LocationAreasBaseURL = baseURL + "/location-area"
)

// func (c *Client) LocationAreasBaseURL() string {
// 	return LocationAreasBaseURL
// }

func (c *Client) GetLocationAreas(pageURL *string) (Resource, error) {
	url := LocationAreasBaseURL
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Resource{}, err
	}

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return Resource{}, err
	}
	defer res.Body.Close()

	var resource Resource
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&resource); err != nil {
		return Resource{}, err
	}

	return resource, nil
}
