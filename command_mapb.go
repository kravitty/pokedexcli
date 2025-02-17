package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMapb() error {
	res, err := http.Get(MapbUrl)
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
	MapUrl = results.Next
	MapbUrl = results.Previous
	for _, value := range results.Results {
		fmt.Println(value.Name)
	}
	return nil
}
