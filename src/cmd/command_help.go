package main

import(
	"fmt"
)

func helpcomand(args ...string) error {
	var result string
	result = "\n These are the following comands: \n"
	for k , v := range Commands {
		result += "\n" + k + ": " + "\n" + v.description + "\n\n"
	}
	fmt.Println(result)
	return nil
}