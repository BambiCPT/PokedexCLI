package commands

import (
	"errors"
	"fmt"

	"github.com/BambiCPT/pokedexcli/internal/pokeapi"
)

type Config struct {
	NextURL *string
	PrevURL *string
	Client  pokeapi.Client
}

func MapCommand(cfg *Config) error {
	locationResp, err := cfg.Client.GetLocations(cfg.NextURL)
	if err != nil {
		return err
	}

	*cfg.NextURL = locationResp.Next
	*cfg.PrevURL = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func MapBackCommand(cfg *Config) error {
	if *cfg.PrevURL == "" {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.Client.GetLocations(cfg.PrevURL)
	if err != nil {
		return err
	}

	*cfg.NextURL = locationResp.Next
	*cfg.PrevURL = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}
