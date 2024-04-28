package main

import (
	"encoding/json"
	// "fmt"
	"io"
	"net/http"
)

func urlReader(cfg *config, url string) ([]byte, error) {
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
	return data, nil
}

func ObjReader[T LocationData | ExploreData | CatchData](cfg *config, url string, empty T) (T, error) {
	data, err := urlReader(cfg, url)

	if err != nil {
		return empty, err
	}

	json.Unmarshal(data, &empty)
	return empty, nil
}
