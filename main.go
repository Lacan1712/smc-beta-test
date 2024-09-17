package main

import (
	Apresentation "SpringManagerCLI/src/ASCII/Apresentation"
	"SpringManagerCLI/src/commands"
)

func main() {
	Apresentation.Apresentation()
	commands.InitCommand()
}
