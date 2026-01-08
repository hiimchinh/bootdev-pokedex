package cmd

import (
	"fmt"

	"github.com/hiimchinh/bootdev-pokedex/internal/pokeapi"
)

func CommandMap(cfg *Config) error {
	res := pokeapi.GetLocationAreas(*cfg.Next)
	areas := res.Results
	for _, area := range areas {
		fmt.Println(area.Name)
	}

	cfg.Next = res.Next
	cfg.Previous = res.Previous
	return nil
}
