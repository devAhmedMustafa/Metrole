package cli

type CommandRegistry struct {
	commands map[string]Command
}

func NewCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		commands: make(map[string]Command),
	}
}

func (cr *CommandRegistry) RegisterCommand(name string, command Command) {
	cr.commands[name] = command
}

func (cr *CommandRegistry) GetCommand(name string) (Command, bool) {
	command, exists := cr.commands[name]
	return command, exists
}
