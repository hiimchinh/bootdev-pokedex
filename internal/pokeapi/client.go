package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hiimchinh/bootdev-pokedex/internal/pokecache"
)

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ResLocationAreas struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

func GetLocationAreas(url string, cache *pokecache.Cache) ResLocationAreas {
	cacheEntry, ok := cache.Get(url)
	if ok {
		locationAreas := ResLocationAreas{}
		err := json.Unmarshal(cacheEntry, &locationAreas)
		if err != nil {
			fmt.Println(err)
		}
		return locationAreas
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	cache.Add(url, body)
	locationAreas := ResLocationAreas{}
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		fmt.Println(err)

	}
	return locationAreas

}
