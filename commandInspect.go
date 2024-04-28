package main

import (
	"errors"
	"fmt"
)

func commandInspect(args []string, cfg *config) error {
	if len(args) != 1 {
		return errors.New("-inspect requires one argument")
	}
	pok, ok := cfg.pokedex[args[0]]
	if !ok {
		return errors.New("you haven't caught that pokemon yet")
	}
	fmt.Printf("Name : %s\n", pok.Name)
	fmt.Printf("Height: %d\n", pok.Height)
	fmt.Printf("Weight: %d\n", pok.Weight)
	fmt.Println("Stats:")
	for _, stat := range pok.Stats {
		fmt.Printf("   - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pok.Types {
		fmt.Printf("   - %s\n", typ.Type.Name)
	}
	return nil
}
