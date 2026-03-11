package main

import (
	"metrole/src/bootstrap"
	"metrole/src/cli"
	"metrole/src/config"
	"os"
)

func main() {

	config.Load()

	args := os.Args[1:]

	command_registry := bootstrap.CommandRegistry()
	argument_parser := cli.NewArgumentParser()
	parser := cli.NewCommandParser(argument_parser, command_registry)

	command, execArgs, err := parser.Parse(args)
	if err != nil {
		panic(err)
	}

	if err := command.Execute(execArgs); err != nil {
		panic(err)
	}
}
