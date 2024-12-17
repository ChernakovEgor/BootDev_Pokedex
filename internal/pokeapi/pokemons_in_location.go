package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemons(location string) (PokemonReply, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + location
	if data, ok := c.cache.Get(url); ok {
		var reply PokemonReply
		err := json.Unmarshal(data, &reply)
		if err != nil {
			return PokemonReply{}, err
		}
		return reply, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return PokemonReply{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonReply{}, fmt.Errorf("error reading body: %v", err)
	}
	c.cache.Add(url, data)

	var reply PokemonReply
	// fmt.Println(string(data))
	err = json.Unmarshal(data, &reply)
	if err != nil {
		return PokemonReply{}, fmt.Errorf("error unmarshaling: %v", err)
	}

	return reply, nil
}
