package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nWelcome!")
		fmt.Println(menu())
		fmt.Print(">")
		reader.Scan()
		command := reader.Text()
		args := cleanInput(command)
		mainLogic(args...)
	}
}
//cleans the input when received from the user
func cleanInput(input string) []string {
	lower := strings.ToLower(input)
	trimmed := strings.TrimSpace(lower)
	return strings.Fields(trimmed)
}

func menu() string {
	var result string
	result = "\n These are the following comands: \n"
	count := 1
	for k := range Commands {
		result += strconv.Itoa(count) + ". " + k + ": " + Commands[k].description + "\n"
		count++
	}

	return result
}

func mainLogic(command ...string) {
	if len(command) == 0 {
		fmt.Println("No command provided")
		return
	}

	// Extract the command key and the arguments
	cmdKey := command[0]
	args := command[1:]

	// Search for the command in the Commands map
	value, ok := Commands[cmdKey]
	if !ok {
		fmt.Println("The command does not exist")
		return
	}

	// Call the command function with the remaining arguments
	err := value.command(args...)
	if err != nil {
		fmt.Println(err)
	}
}
