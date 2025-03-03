package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetLocation -
func (c *Client) GetLocation(locationName string) (LocationDetails, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationResp := LocationDetails{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationDetails{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetails{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationDetails{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationDetails{}, err
	}

	locationResp := LocationDetails{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return LocationDetails{}, err
	}

	c.cache.Add(url, dat)

	return locationResp, nil
}
