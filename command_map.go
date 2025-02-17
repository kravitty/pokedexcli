package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type results struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var MapUrl = "https://pokeapi.co/api/v2/location-area"
var MapbUrl = ""

func commandMap() error {
	res, err := http.Get(MapUrl)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%v", body)
	results := results{}
	err = json.Unmarshal([]byte(body), &results)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println("Next:", results.Next)
	MapUrl = results.Next
	MapbUrl = results.Previous
	for _, value := range results.Results {
		fmt.Println(value.Name)
	}
	return nil
}
