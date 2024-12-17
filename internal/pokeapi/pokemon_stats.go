package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (PokemonStats, error) {
	url := baseURL + "/pokemon/" + name

	if data, ok := c.cache.Get(url); ok {
		var pokestats PokemonStats
		err := json.Unmarshal(data, &pokestats)
		if err != nil {
			return PokemonStats{}, fmt.Errorf("error unmarshaling data from cache: %v", err)
		}

		return pokestats, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return PokemonStats{}, fmt.Errorf("error getting pokemon '%s': %v", name, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonStats{}, fmt.Errorf("error reading body: %v", err)
	}

	var pokestats PokemonStats
	err = json.Unmarshal(data, &pokestats)
	if err != nil {
		return PokemonStats{}, fmt.Errorf("error unmarshaling pokemon data: %v", err)
	}

	return pokestats, nil
}
