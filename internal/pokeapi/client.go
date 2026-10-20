package pokeapi

import (
	"net/http"
	"time"

	"git.jaxon.io/jaxon-boot/pokedex/internal/pokecache"
)

type Client struct {
	HttpClient http.Client
	Cache      pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		HttpClient: http.Client{
			Timeout: timeout,
		},
		Cache: *pokecache.NewCache(cacheInterval),
	}
}
