package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetExpLocations(areaName string) (ExpLocationResponse, error) {
	url := baseURL + "/location-area/" + areaName

	if val, ok := c.cache.Get(url); ok {
		locationResp := ExpLocationResponse{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return ExpLocationResponse{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ExpLocationResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ExpLocationResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return ExpLocationResponse{}, err
	}

	locationResp := ExpLocationResponse{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return ExpLocationResponse{}, err
	}

	c.cache.Add(url, data)
	return locationResp, nil
}
