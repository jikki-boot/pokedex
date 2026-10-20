package pokeapi

import (
	"io"
	"net/http"
)

func (c *Client) Get(url string) ([]byte, error) {
	cachedData, exists := c.Cache.Get(url)
	if exists {
		return cachedData, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	c.Cache.Add(url, data)

	return data, nil
}
