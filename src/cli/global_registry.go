package cli

var defaultRegistry = NewCommandRegistry()

func RegisterCommand(name string, command Command) {
	defaultRegistry.RegisterCommand(name, command)
}

func DefaultRegistry() *CommandRegistry {
	return defaultRegistry
}
