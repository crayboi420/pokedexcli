package main

import (
	"bufio"
	"fmt"
	"os"
	"errors"
	// "github.com/mtslzr/pokeapi-go"
	// "net/http"
)

func Mainloop(){
	reader := bufio.NewReader(os.Stdin)
	cfg := config{
		mapApiID: 1,
		mapApiIDNext: 20,
		commandMap: map[string]command{
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
			},
		}
	for {
		fmt.Printf("pokedex > ")
		text,ok := reader.ReadString('\n')
		if ok!=nil{
			fmt.Println("Reading the string failed")
			break
		}

		comm,ok := Parsetext(text[:len(text)-1],&cfg)
		if ok!=nil {
			continue
		}
		
		if comm.callback(&cfg)!=nil{
			continue
		}

		if text == "exit\n" {break}
	}
}

type config struct{
	mapApiID, mapApiIDNext int
	commandMap map[string]command
}

type command struct{
	name,description string
	callback func(*config) error
}

func Parsetext(s string, cfg *config) (command,error){
	
	comm,ok := cfg.commandMap[s]
	if !ok{
		fmt.Println("Not a valid command")
		return command{},errors.New("not in the list of commands")
	}
	return comm,nil
}

func commandHelp(cfg *config) error{
	fmt.Println("\nPokedexCLI is a tool that will function as a pokedex you can use in the command line.")
	fmt.Println("The commands you are allowed to use are:\n")
	for key,value :=range cfg.commandMap{
		fmt.Println(key,": ",value.description)
		// fmt.Println(value.description)
	}
	return nil
}
func commandExit(cfg *config) error{
	os.Exit(0)
	return nil
}
func mapF(cfg *config) error {
	return nil
}
func mapB(cfg *config) error{
	return nil
}