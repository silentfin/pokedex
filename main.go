package main

func main() {
	configs := config{commands: getCommands()}
	repl(&configs)
}
