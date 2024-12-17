package pokeapi

type PokemonReply struct {
	Encounters []Encounter `json:"pokemon_encounters"`
}

type Encounter struct {
	Pokemon struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"pokemon"`
}

type PokemonStats struct {
	Name       string `json:"name"`
	Experience int    `json:"base_experience"`
}
