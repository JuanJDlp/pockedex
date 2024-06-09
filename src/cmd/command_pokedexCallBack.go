package main

import "fmt"

func pokedexCallBack(args ...string) error {
	var result string
	for _, v := range pokedex {
		result += "-" + v.Name + "\n"
	}
	fmt.Print(result)
	return nil
}
