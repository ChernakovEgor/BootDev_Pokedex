package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	// "github.com/ChernakovEgor/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(conf *Config, args ...string) error
}

var commandRegistry map[string]cliCommand

func init() {
	commandRegistry = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List all Pokemon in area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "explore",
			description: "Try to catch the Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught Pokemon",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandExit(conf *Config, _ ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *Config, _ ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for k, v := range commandRegistry {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}

func commandMap(conf *Config, _ ...string) error {
	if conf.Next == "" {
		fmt.Println("you're on the last page")
		return nil
	}
	reply, err := conf.client.GetLocations(conf.Next)
	if err != nil {
		log.Fatalln(err)
	}

	for _, location := range reply.Results {
		fmt.Println(location.Name)
	}

	conf.Next = reply.Next
	conf.Previous = reply.Previous
	return nil
}

func commandMapb(conf *Config, _ ...string) error {
	if conf.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	reply, err := conf.client.GetLocations(conf.Previous)
	if err != nil {
		log.Fatalln(err)
	}

	for _, location := range reply.Results {
		fmt.Println(location.Name)
	}

	conf.Next = reply.Next
	conf.Previous = reply.Previous
	return nil
}

func commandExplore(ctx *Config, args ...string) error {
	log.Println(args)
	res, err := ctx.client.GetPokemons(args[1])
	if err != nil {
		log.Fatalf("failed to explore: %v", err)
	}
	var pokemons []string
	for _, entry := range res.Encounters {
		pokemons = append(pokemons, "- "+entry.Pokemon.Name)
	}

	foundPokemons := strings.Join(pokemons, "\n")
	fmt.Printf("Exploring %s...\n", args[1])
	fmt.Println("Found Pokemon:")
	fmt.Println(foundPokemons)
	return nil
}

func commandCatch(ctx *Config, args ...string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", args[1])
	stats, err := ctx.client.GetPokemon(args[1])
	if err != nil {
		log.Fatalln("error getting pokemon:", err)
	}

	r := rand.Intn(300)
	if r >= stats.Experience {
		fmt.Printf("%s was caught!\n", args[1])
		ctx.Caught[args[1]] = stats
	} else {
		fmt.Printf("%s escaped!\n", args[1])
	}

	return nil
}

func commandInspect(ctx *Config, args ...string) error {
	value, ok := ctx.Caught[args[1]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", value.Name)
	fmt.Printf("Experience: %d\n", value.Experience)

	return nil
}

func commandPokedex(ctx *Config, _ ...string) error {
	fmt.Println("Your Pokedex:")
	for key := range ctx.Caught {
		fmt.Printf(" - %s\n", key)
	}

	return nil
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	res := strings.Fields(lower)
	return res
}
