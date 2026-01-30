package cmd

import (
	"fmt"

	"github.com/hiimchinh/bootdev-pokedex/internal/pokeapi"
	"github.com/hiimchinh/bootdev-pokedex/internal/pokecache"
)

func CommandExplore(cfg *Config, cache *pokecache.Cache, args []string) error {
	baseUrl := "https://pokeapi.co/api/v2/location-area"

	name := args[0]
	url := baseUrl + "/" + name
	res := pokeapi.GetLocationAreaDetail(url, cache)
	fmt.Printf("Exploring %s...\n", name)
	pokemonEncounters := res.PokemonEncounters
	if len(pokemonEncounters) == 0 {
		fmt.Println("Not found any Pokemon here.")
		return nil

	}

	fmt.Println("Found Pokemon:")
	for i := range pokemonEncounters {
		fmt.Printf(" - %s\n", pokemonEncounters[i].Pokemon.Name)
	}
	return nil

}
