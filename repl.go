package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/silentfin/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	previous string
	next     string
	commands map[string]cliCommand
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
		"map": {
			name:        "map",
			description: "Show current 20 map areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous 20 map areas",
			callback:    commandMapPrevious,
		},
	}
}

func commandHelp(configs *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, command := range configs.commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandExit(configs *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(configs *config) error {
	var url string
	if configs.next != "" {
		url = configs.next
	} else {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	areas, err := pokeapi.GetMapAreas(url)
	if err != nil {
		return err
	}

	configs.next = areas.Next
	configs.previous = areas.Previous

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMapPrevious(configs *config) error {
	var url string
	if configs.previous != "" {
		url = configs.previous
	} else {
		fmt.Println("No previous pages")
		return nil
	}

	areas, err := pokeapi.GetMapAreas(url)
	if err != nil {
		return err
	}

	configs.next = areas.Next
	configs.previous = areas.Previous

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func repl(configs *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		actualInput := cleanInput(input)
		if len(actualInput) == 0 {
			continue
		}
		commandName := actualInput[0]
		command, ok := configs.commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(configs)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
