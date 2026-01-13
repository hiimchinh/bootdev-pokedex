package cmd

import (
	"fmt"

	"github.com/hiimchinh/bootdev-pokedex/internal/pokeapi"
	"github.com/hiimchinh/bootdev-pokedex/internal/pokecache"
)

func CommandMap(cfg *Config, cache *pokecache.Cache) error {
	res := pokeapi.GetLocationAreas(*cfg.Next, cache)
	areas := res.Results
	for _, area := range areas {
		fmt.Println(area.Name)
	}

	cfg.Next = res.Next
	cfg.Previous = res.Previous
	return nil
}
