package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/crayboi420/pokedexcli/internal/pokecache"
)

type config struct {
	mapURLF, mapURLB string
	commandMap       map[string]command
	ch               pokecache.Cache
	pokedex          map[string]CatchData
}

type command struct {
	name, description string
	callback          func([]string, *config) error
}

func getCommands() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations of the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the last 20 locations of the map",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Prints the list of pokemon for the given area name",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to catch the pokemon passed as argument",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect pokemon already caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Prints out the list of pokemons in the pokedex",
			callback:    commandPokedex,
		},
	}
}

func Mainloop() {
	reader := bufio.NewReader(os.Stdin)
	cfg := config{
		mapURLF:    "https://pokeapi.co/api/v2/location-area/",
		mapURLB:    "https://pokeapi.co/api/v2/location-area/",
		commandMap: getCommands(),
		ch:         pokecache.NewCache(10 * time.Second),
		pokedex:    map[string]CatchData{},
	}

	for {
		fmt.Printf("pokedex > ")

		text, ok := reader.ReadString('\n')
		if ok != nil {
			fmt.Println("Reading the string failed")
			break
		}
		text = text[:len(text)-1] //Remove last \n
		if len(text) == 0 {
			continue
		}

		comm, args, err := Parsetext(text, &cfg)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = comm.callback(args, &cfg)

		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func Parsetext(text string, cfg *config) (command, []string, error) {

	text = strings.ToLower(text)
	words := strings.Split(text, " ")

	comm, ok := cfg.commandMap[words[0]]
	args := []string{}
	if len(words)-1 > 0 {
		args = words[1:]
	}
	if !ok {
		return command{}, []string{}, errors.New(words[0] + ": not a valid command")
	}
	return comm, args, nil
}

func commandHelp(args []string, cfg *config) error {
	if len(args) > 0 {
		return errors.New("-help doesn't accept any arguments")
	}
	fmt.Println("\nPokedexCLI is a tool that will function as a pokedex you can use in the command line.")
	fmt.Println("The commands you are allowed to use are:")
	fmt.Println()
	for _, value := range cfg.commandMap {
		fmt.Printf("- %s : %s\n", value.name, value.description)
	}
	return nil
}
func commandExit(args []string, cfg *config) error {
	if len(args) > 0 {
		return errors.New("-exit doesn't accept any arguments")
	}
	os.Exit(0)
	return nil
}
