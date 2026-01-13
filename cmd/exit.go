package cmd

import (
	"fmt"
	"os"

	"github.com/hiimchinh/bootdev-pokedex/internal/pokecache"
)

func CommandExit(cfg *Config, cache *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}
