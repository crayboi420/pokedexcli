package main

import (
	"bufio"
	"fmt"
	"os"
	"errors"
	"strings"
	// "github.com/mtslzr/pokeapi-go"
	// "net/http"
)

func Mainloop(){
	reader := bufio.NewReader(os.Stdin)
	cfg := config{
		mapURLF: "https://pokeapi.co/api/v2/location-area/",
		mapURLB: "https://pokeapi.co/api/v2/location-area/",
		commandMap: getCommands(),
		}
	for {
		fmt.Printf("pokedex > ")
		text,ok := reader.ReadString('\n')
		if ok!=nil{
			fmt.Println("Reading the string failed")
			break
		}
		if len(text)==1 {continue}
		text = text[:len(text)-1]
 		comm,ok := Parsetext(strings.ToLower(text),&cfg)
		if ok!=nil {
			continue
		}

		err:=comm.callback(&cfg)

		if err!=nil{
			fmt.Println(err)
			continue
		}

		if text == "exit\n" {break}
	}
}

func getCommands() map[string]command{
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
					"map":{
						name: "map",
						description: "Displays the next 20 locations of the map",
						callback: mapF,
					},
					"mapb":{
						name: "mapb",
						description: "Displays the last 20 locations of the map",
						callback: mapB,
					},
				}
}

type config struct{
	mapURLF,mapURLB string
	commandMap map[string]command
}

type command struct{
	name,description string
	callback func(*config) error
}

func Parsetext(s string, cfg *config) (command,error){
	
	comm,ok := cfg.commandMap[s]
	if !ok{
		fmt.Println(s,": not a valid command")
		return command{},errors.New("not in the list of commands")
	}
	return comm,nil
}

func commandHelp(cfg *config) error{
	fmt.Println("\nPokedexCLI is a tool that will function as a pokedex you can use in the command line.")
	fmt.Println("The commands you are allowed to use are:")
	fmt.Println()
	for _,value :=range cfg.commandMap{
		fmt.Printf("- %s : %s\n",value.name,value.description)
	}
	return nil
}
func commandExit(cfg *config) error{
	os.Exit(0)
	return nil
}
