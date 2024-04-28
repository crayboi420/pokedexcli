package main

import (
	"errors"
	"fmt"
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

func commandMap(args []string, cfg *config) error {
	if len(args) > 0 {
		return errors.New("-map doesn't accept arguments")
	}
	url := cfg.mapURLF

	locobj, err := ObjReader(cfg, url, LocationData{})

	if err != nil {
		return errors.New("no locations after this")
	}

	cfg.mapURLF = locobj.Next
	cfg.mapURLB = locobj.Previous

	fmt.Println("Printing the next 20 locations...")
	LocationPrinter(&locobj)

	return nil
}

func commandMapB(args []string, cfg *config) error {

	if len(args) > 0 {
		return errors.New("-mapb doesn't accept arguments")
	}

	url := cfg.mapURLB

	locobj, err := ObjReader(cfg, url, LocationData{})

	if err != nil {
		return errors.New("no locations before this")
	}

	cfg.mapURLF = locobj.Next
	cfg.mapURLB = locobj.Previous

	fmt.Println("\nPrinting the last 20 locations...")
	LocationPrinter(&locobj)

	return nil
}

func LocationPrinter(locobj *LocationData) {

	for _, value := range locobj.Results {
		fmt.Println()
		fmt.Print(value.Name)
	}
	fmt.Println()
}
