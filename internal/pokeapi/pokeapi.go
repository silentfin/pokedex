package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/silentfin/pokedex/internal/pokecache"
)

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationAreaData struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

var apiCache = pokecache.NewCache(5 * time.Second)

func GetMapAreas(url string) (LocationAreas, error) {
	var data []byte
	cachedData, found := apiCache.Get(url)
	if found {
		data = cachedData
	} else {
		res, err := http.Get(url)
		if err != nil {
			return LocationAreas{}, err
		}
		defer res.Body.Close()

		if res.StatusCode == http.StatusOK {
			data, err = io.ReadAll(res.Body)
			if err != nil {
				return LocationAreas{}, err
			}
			apiCache.Add(url, data)
		} else {
			return LocationAreas{}, fmt.Errorf("request failed: %d", res.StatusCode)
		}
	}

	areas := LocationAreas{}
	if err := json.Unmarshal(data, &areas); err != nil {
		return LocationAreas{}, err
	}
	return areas, nil
}

func GetPokemonsInLocationArea(url string) (LocationAreaData, error) {
	var data []byte
	cachedData, found := apiCache.Get(url)
	if found {
		data = cachedData
	} else {
		res, err := http.Get(url)
		if err != nil {
			return LocationAreaData{}, err
		}
		defer res.Body.Close()

		if res.StatusCode == http.StatusOK {
			data, err = io.ReadAll(res.Body)
			if err != nil {
				return LocationAreaData{}, err
			}
			apiCache.Add(url, data)
		} else {
			return LocationAreaData{}, fmt.Errorf("request failed: %d", res.StatusCode)
		}
	}
	pokemons := LocationAreaData{}
	if err := json.Unmarshal(data, &pokemons); err != nil {
		return LocationAreaData{}, err
	}
	return pokemons, nil
}

func GetPokemonData(url string) (Pokemon, error) {
	var data []byte
	cachedData, found := apiCache.Get(url)
	if found {
		data = cachedData
	} else {
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		if res.StatusCode == http.StatusOK {
			data, err = io.ReadAll(res.Body)
			if err != nil {
				return Pokemon{}, err
			}
			apiCache.Add(url, data)
		} else {
			return Pokemon{}, fmt.Errorf("request failed: %d", res.StatusCode)
		}
	}
	pokemonInfo := Pokemon{}
	if err := json.Unmarshal(data, &pokemonInfo); err != nil {
		return Pokemon{}, err
	}
	return pokemonInfo, nil
}
