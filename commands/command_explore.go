package commands

import (
	"errors"
	"fmt"
)

func ExploreCommand(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide a location area name")
	}

	areaName := args[0]
	locationResp, err := cfg.Client.GetExpLocations(areaName)
	if err != nil {
		return err
	}

	if len(locationResp.PokemonEncounters) == 0 {
		fmt.Println("No Pokémon found in this area.")
		return nil
	}

	fmt.Printf("Exploring %v...", areaName)
	fmt.Println("Found Pokémon: ")
	for _, encounter := range locationResp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
