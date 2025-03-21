package commands

import (
	"errors"
	"fmt"
)

func MapCommand(cfg *Config) error {
	locationResp, err := cfg.Client.GetLocations(cfg.NextURL)
	if err != nil {
		return err
	}

	cfg.NextURL = safeStringPtr(&locationResp.Next)
	cfg.PrevURL = safeStringPtr(&locationResp.Previous)

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

	cfg.NextURL = safeStringPtr(&locationResp.Next)
	cfg.PrevURL = safeStringPtr(&locationResp.Previous)

	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func safeStringPtr(s *string) *string {
	if s == nil {
		return nil
	}
	str := *s
	return &str
}
