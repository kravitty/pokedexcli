package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	fmt.Println("Check cache")
	//Check cache first
	if cacheData, exists := c.cache.Get(url); exists {
		fmt.Println("Cache used start ListLocations")
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(cacheData, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResp, nil
	}

	// Execute request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	fmt.Println("Do request")
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// Add to cache
	c.cache.Add(url, dat)
	fmt.Println("Cache added: ", url)

	return locationsResp, nil
}
