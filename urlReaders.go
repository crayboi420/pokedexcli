package main

import (
	"net/http"
	"encoding/json"
	"io"
)
func urlReader(cfg *config, url string) ([]byte,error){
	var data []byte

	if val, is := cfg.ch.Get(url); is {
		data = val
	} else {
		res, err := http.Get(url)
		if err != nil {
			return data, err
		}
		txt, _ := io.ReadAll(res.Body)
		res.Body.Close()
		data = []byte(txt)

		cfg.ch.Add(url, data)
	}
	return data,nil
}

func LocationReader(cfg *config, url string) (LocationData, error) {

	locobj := LocationData{}
	data,err := urlReader(cfg,url)

	if err!=nil {return locobj,err}
	
	json.Unmarshal(data, &locobj)
	return locobj, nil
}

func PokemonReader(cfg *config, url string) (PokemonData,error){
	pokobj := PokemonData{}
	data,err := urlReader(cfg,url)

	if err!=nil {return pokobj,err}

	json.Unmarshal(data,&pokobj)
	return pokobj,nil
}