package cli

type Command interface {
	Execute(args []Argument) error
}
