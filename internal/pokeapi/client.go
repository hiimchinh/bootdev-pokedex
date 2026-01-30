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

type ResLocationAreaDetail struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
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

func GetLocationAreaDetail(url string, cache *pokecache.Cache) *ResLocationAreaDetail {
	cacheEntry, ok := cache.Get(url)
	if ok {
		locationAreaDetail := &ResLocationAreaDetail{}
		err := json.Unmarshal(cacheEntry, locationAreaDetail)
		if err != nil {
			log.Fatal(err)
		}
		return locationAreaDetail

	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code %d and body: %s", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	cache.Add(url, body)
	detail := &ResLocationAreaDetail{}
	err = json.Unmarshal(body, detail)
	if err != nil {
		log.Fatal(err)
	}

	return detail

}
