package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/ChernakovEgor/pokedexcli/internal/pokeapi"
)

type Config struct {
	client   *pokeapi.Client
	Next     string
	Previous string
	Caught   map[string]pokeapi.PokemonStats
}

func DefaultConfig() *Config {
	client := pokeapi.NewClient(time.Second * 5)
	caught := make(map[string]pokeapi.PokemonStats)
	return &Config{client, pokeapi.GetLocationsURL(), "", caught}
}

func NewConfig(cacheInterval time.Duration) *Config {
	client := pokeapi.NewClient(cacheInterval)
	caught := make(map[string]pokeapi.PokemonStats)
	return &Config{client, pokeapi.GetLocationsURL(), "", caught}
}

func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		if command, ok := commandRegistry[words[0]]; ok {
			command.callback(cfg, words...)
		} else {
			fmt.Println("Unknown command")
		}
	}
}
