package main

import (
	"errors"
	"fmt"
	"github.com/JuanJDlp/pockedex/src/internal/api"
)

func explore(args ...string) error {
	fmt.Print(args)
	if len(args) == 0 {
		return errors.New("please provide an area name")
	}

	areaName := args[0]
	res, err := api.GetPokemonsFromLocations(areaName)
	if err != nil {
		return err
	}
	var result string
	for _, v := range res.PokemonEncounters {
		result += fmt.Sprintf("-%s\n", v.Pokemon.Name)
	}

	fmt.Println(result)
	return nil

}
