package cmd

import (
	"fmt"

	"github.com/hiimchinh/bootdev-pokedex/internal/pokeapi"
)

func CommandMapb(cfg *Config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	res := pokeapi.GetLocationAreas(*cfg.Previous)
	cfg.Next = res.Next
	cfg.Previous = res.Previous

	areas := res.Results
	for _, area := range areas {
		fmt.Println(area.Name)
	}

	return nil
}
