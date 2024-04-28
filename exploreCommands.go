package main

import (
	"errors"
	"fmt"
) 

type PokemonData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandExplore(args []string, cfg *config) error {
	
	if len(args)!=1 {
		return errors.New("-explore needs one argument")
	}

	url := "https://pokeapi.co/api/v2/location-area/"
	url += args[0]

	pokobj,err := PokemonReader(cfg,url)

	if err!=nil {
		return errors.New("there was an error obtaining the pokemon")
	}
	if len(pokobj.PokemonEncounters) == 0{
		return errors.New("no pokemon in area/not a valid area")
	}
	fmt.Println("Exploring ",args[0]," ...")

	PokemonPrinter(&pokobj)

	return nil
}

func PokemonPrinter(pocobj *PokemonData){
	for _,val := range pocobj.PokemonEncounters{
		fmt.Println("-",val.Pokemon.Name)
	}
}