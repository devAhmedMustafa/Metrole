package cli

import "fmt"

type CommandParser struct {
	argumentParser  *ArgumentParser
	commandRegistry *CommandRegistry
}

func NewCommandParser(argumentParser *ArgumentParser, commandRegistry *CommandRegistry) *CommandParser {
	return &CommandParser{
		argumentParser:  argumentParser,
		commandRegistry: commandRegistry,
	}
}

func (cp *CommandParser) Parse(args []string) (Command, []Argument, error) {
	if len(args) == 0 {
		return nil, nil, fmt.Errorf("no command provided")
	}

	commandArgs, err := cp.argumentParser.Parse(args)
	if err != nil {
		return nil, nil, err
	}

	commandName := args[0]

	command, exists := cp.commandRegistry.GetCommand(commandName)
	if !exists {
		return nil, nil, fmt.Errorf("command not found")
	}

	return command, commandArgs, nil
}
