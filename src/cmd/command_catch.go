package main

import (
	"fmt"
	"math/rand"
	"pockedex/src/internal/api"
)


func catch(args ...string) error {
	_, ok := pokedex[args[0]]
	if ok {
		fmt.Print("You alredy have this pokemon!\n")
		return nil
	}
	pokemon, err := api.GetPokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s\n", args[0])
	lvl := pokemon.BaseExperience
	randomNum := rand.Intn(lvl)

	if randomNum > 1 {
		fmt.Printf("\n%s was caught", pokemon.Name)
		pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("\n%s escaped!", pokemon.Name)
	}

	return nil
}