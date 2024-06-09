package main

import (
	"fmt"
	"pockedex/src/internal/api"
)

func mapb(args ...string) error {
	res, err := api.GetLocations(false)

	if err != nil {
		return err
	}

	for _, v := range res.Locations {
		fmt.Println(v.Name)
	}
	return nil

}

func mapFunc(args ...string) error {
	res, err := api.GetLocations(true)

	if err != nil {
		return err
	}

	for _, v := range res.Locations {
		fmt.Println(v.Name)
	}
	return nil
}
