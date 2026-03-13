package cli_commands

import (
	"metrole/src/cli"
	"metrole/src/core/auth"
)

type CheckAuthCommand struct{}

func init() {
	cli.RegisterCommand("is-authed", &CheckAuthCommand{})
}

func (c *CheckAuthCommand) Execute(args []cli.Argument) error {
	authService := auth.NewAuthService()
	if authService.IsAuthenticated() {
		println("Authenticated")
	} else {
		println("Not authenticated - please run 'metrole login' to authenticate")
	}
	return nil
}
