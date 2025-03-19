package commands

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BambiCPT/pokedexcli/internal/pokecache"
)

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

type Config struct {
	NextURL string
	PrevURL string
	Cache   *pokecache.Cache
}

func getLocations(cfg *Config, URL string) (LocationResponse, error) {
	var locationData LocationResponse

	if cachedData, found := cfg.Cache.Get(URL); found {
		if err := json.Unmarshal(cachedData, &locationData); err == nil {
			return locationData, nil
		}
	}

	res, err := http.Get(URL)
	if err != nil {
		return locationData, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationData); err != nil {
		return locationData, err
	}

	if cfg.Cache != nil {
		cachedJSON, err := json.Marshal(locationData)
		if err == nil {
			cfg.Cache.Add(URL, cachedJSON)
		}
	}

	return locationData, nil
}

func MapCommand(cfg *Config) error {
	locationData, err := getLocations(cfg, cfg.NextURL)
	if err != nil {
		return err
	}

	cfg.NextURL = locationData.Next
	cfg.PrevURL = locationData.Previous

	for _, location := range locationData.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func MapBackCommand(cfg *Config) error {
	if cfg.PrevURL == "" {
		fmt.Println("You're on the first page.")
		return nil
	}

	locationData, err := getLocations(cfg, cfg.PrevURL)
	if err != nil {
		return err
	}

	cfg.NextURL = locationData.Next
	cfg.PrevURL = locationData.Previous

	for _, location := range locationData.Results {
		fmt.Println(location.Name)
	}
	return nil
}
