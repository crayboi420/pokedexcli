package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type LocationData struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Results `json:"results"`
}
type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func commandMap(cfg *config) error {
	url := cfg.mapURLF

	locobj, err := LocationReader(cfg, url)

	if err != nil {
		return errors.New("no locations after this")
	}

	cfg.mapURLF = locobj.Next
	cfg.mapURLB = locobj.Previous

	fmt.Println("Printing the next 20 locations...")
	LocationPrinter(&locobj)

	return nil
}

func commandMapB(cfg *config) error {
	url := cfg.mapURLB

	locobj, err := LocationReader(cfg, url)

	if err != nil {
		return errors.New("no locations before this")
	}

	cfg.mapURLF = locobj.Next
	cfg.mapURLB = locobj.Previous

	fmt.Println("\nPrinting the last 20 locations...")
	LocationPrinter(&locobj)

	return nil
}

func LocationReader(cfg *config, url string) (LocationData, error) {

	locobj := LocationData{}
	var data []byte

	if val, is := cfg.ch.Get(url); is {
		data = val
		// fmt.Println("Got from map")
	} else {
		res, err := http.Get(url)
		if err != nil {
			return LocationData{}, err
		}
		txt, _ := io.ReadAll(res.Body)
		res.Body.Close()
		data = []byte(txt)

		cfg.ch.Add(url, data)
		// fmt.Println("Added to map")
	}
	json.Unmarshal(data, &locobj)
	return locobj, nil
}

func LocationPrinter(locobj *LocationData) {

	for _, value := range locobj.Results {
		fmt.Println()
		fmt.Print(value.Name)
	}
	fmt.Println()
}
