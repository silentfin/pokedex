package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		actualInput := cleanInput(input)
		commandName := actualInput[0]
		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}
}
