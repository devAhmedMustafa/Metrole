package bootstrap

import (
	"metrole/src/cli"
	_ "metrole/src/cli/commands"
)

func CommandRegistry() *cli.CommandRegistry {
	return cli.DefaultRegistry()
}
