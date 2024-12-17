package pokeapi

import (
	"github.com/ChernakovEgor/pokedexcli/internal/pokecache"
	"time"
)

const baseURL = "https://pokeapi.co/api/v2/"

type Client struct {
	cache *pokecache.Cache
}

func NewClient(cacheInterval time.Duration) *Client {
	cache := pokecache.NewCache(cacheInterval)
	return &Client{cache}
}
