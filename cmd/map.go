package cmd

import (
	"fmt"

	"github.com/hiimchinh/bootdev-pokedex/internal/pokeapi"
)

func CommandMap() error {
	res := pokeapi.GetLocationAreas()
	areas := res.Results
	for _, area := range areas {
		fmt.Println(area.Name)
	}
	return nil
}
