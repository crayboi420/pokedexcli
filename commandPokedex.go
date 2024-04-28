package main

import (
	"errors"
	"fmt"
)

func commandPokedex(args []string, cfg *config) error {
	if len(args) > 0 {
		return errors.New("-pokedex doesn't take any arguments")
	}
	if len(cfg.pokedex) == 0 {
		return errors.New("no pokemon in pokedex")
	}
	for key := range cfg.pokedex {
		fmt.Printf("- %s\n", key)
	}
	return nil
}
