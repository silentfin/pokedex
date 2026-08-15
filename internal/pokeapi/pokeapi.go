package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
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

func GetMapAreas(url string) (PokedexLocationArea, error) {
	res, err := http.Get(url)
	if err != nil {
		return PokedexLocationArea{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokedexLocationArea{}, err
	}

	areas := PokedexLocationArea{}
	if err := json.Unmarshal(data, &areas); err != nil {
		return PokedexLocationArea{}, err
	}
	return areas, nil
}
