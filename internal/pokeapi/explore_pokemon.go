package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(location string) (LocationDetails, error) {

	locationResponse, err := c.ListLocations(nil)
	if err != nil {
		fmt.Println("Error listing locations:", err)
		return LocationDetails{}, err
	}
	//fmt.Println(locationResponse)

	var url string
	for _, result := range locationResponse.Results {
		if result.Name == location {
			url = result.URL
			fmt.Println("Found location: " + location)
		}
		//TODO: Handle error if location not found
	}
	fmt.Println("URL: " + url)

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

	locationsResp := LocationDetails{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationDetails{}, err
	}

	return locationsResp, nil
}
