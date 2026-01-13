package cmd

import (
	"fmt"

	"github.com/hiimchinh/bootdev-pokedex/internal/pokeapi"
	"github.com/hiimchinh/bootdev-pokedex/internal/pokecache"
)

func CommandMapb(cfg *Config, cache *pokecache.Cache) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	res := pokeapi.GetLocationAreas(*cfg.Previous, cache)
	cfg.Next = res.Next
	cfg.Previous = res.Previous

	areas := res.Results
	for _, area := range areas {
		fmt.Println(area.Name)
	}

	return nil
}
