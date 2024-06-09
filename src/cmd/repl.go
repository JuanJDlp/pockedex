package main

import (
	"os"
	"github.com/JuanJDlp/pockedex/src/internal/api"
)

type command struct {
	name        string
	description string
	command     func(...string) error
}

var (
	Commands map[string]command
	pokedex  = make(map[string]api.Pokemon)
)

func init() {
	Commands = map[string]command{
		"help": {
			name:        "help",
			description: "this command will show you all available comands",
			command:     helpcomand,
		},
		"exit": {
			name:        "exit",
			description: "closes the program",
			command: func(...string) error {
				os.Exit(0)
				return nil
			},
		},
		"map": {
			name:        "map",
			description: "The map command displays the names of 20 location areas in the Pokemon world",
			command:     mapFunc,
		},
		"mapb": {
			name:        "map",
			description: "The map command displays the names of 20 location areas in the Pokemon world",
			command:     mapb,
		},
		"explore": {
			name:        "explore",
			description: "it shows you the pokemos from an specific location, ex of usage: explore <area-name>",
			command:     explore,
		},
		"catch": {
			name:        "catch",
			description: "Adds the specified pokemon to the pokedex ex. catch <pokemon-name>",
			command:     catch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspects the stats of an specific pokemon\n  *takes the name of a Pokemon as an argument.",
			command:     inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "shows all pokemos you have caught",
			command:     pokedexCallBack,
		},
	}

}
