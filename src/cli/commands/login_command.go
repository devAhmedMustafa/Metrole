package cli_commands

import (
	"fmt"
	"metrole/src/cli"
	"metrole/src/core/auth"
)

type LoginCommand struct {
}

func init() {
	cli.RegisterCommand("login", &LoginCommand{})
}

func (c *LoginCommand) Execute(args []cli.Argument) error {

	codeChan := auth.StartCallbackListener()

	fmt.Println("Please login via browser...")

	auth.InitiateLogin()

	code := <-codeChan

	idToken, err := auth.ExchangeCodeForIdToken(code)
	if err != nil {
		fmt.Println("Token exchange failed:", err)
		return err
	}

	authService := auth.NewAuthService()
	err = authService.Login(idToken)

	if err != nil {
		fmt.Println("Login failed:", err)
		return err
	}

	fmt.Println("Login successful!")

	return nil
}
