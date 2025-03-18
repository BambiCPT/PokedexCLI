package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
}

var config = Config{
	NextURL: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
	PrevURL: "",
}

func getLocations(URL string) (LocationResponse, error) {
	var locationData LocationResponse
	res, err := http.Get(URL)
	if err != nil {
		return locationData, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationData); err != nil {
		return locationData, err
	}

	return locationData, nil
}

func mapCommand(cfg *Config) error {
	locationData, err := getLocations(cfg.NextURL)
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

func mapBackCommand(cfg *Config) error {
	if cfg.PrevURL == "" {
		fmt.Println("You're on the first page.")
		return nil
	}

	locationData, err := getLocations(cfg.PrevURL)
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
