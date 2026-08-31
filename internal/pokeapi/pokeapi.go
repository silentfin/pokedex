package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/silentfin/pokedex/internal/pokecache"
)

type PokedexLocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var apiCache = pokecache.NewCache(5 * time.Second)

func GetMapAreas(url string) (PokedexLocationArea, error) {
	var data []byte
	cachedData, found := apiCache.Get(url)
	if found {
		data = cachedData
	} else {
		res, err := http.Get(url)
		if err != nil {
			return PokedexLocationArea{}, err
		}
		defer res.Body.Close()

		if res.StatusCode == http.StatusOK {
			data, err = io.ReadAll(res.Body)
			if err != nil {
				return PokedexLocationArea{}, err
			}
			apiCache.Add(url, data)
		} else {
			return PokedexLocationArea{}, fmt.Errorf("request failed: %d", res.StatusCode)
		}
	}

	areas := PokedexLocationArea{}
	if err := json.Unmarshal(data, &areas); err != nil {
		return PokedexLocationArea{}, err
	}
	return areas, nil
}
