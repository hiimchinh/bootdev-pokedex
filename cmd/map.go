package cmd

import (
	"fmt"

	"github.com/hiimchinh/bootdev-pokedex/internal/pokeapi"
)

func CommandMap() error {
	res := pokeapi.GetLocationAreas()
	fmt.Printf("Res is %v", res)
	return nil
}
