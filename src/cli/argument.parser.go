package cli

import (
	"strconv"
	"strings"
)

type ArgumentParser struct{}

func NewArgumentParser() *ArgumentParser {
	return &ArgumentParser{}
}

func (ap *ArgumentParser) Parse(args []string) ([]Argument, error) {
	if len(args) < 2 {
		return nil, nil
	}

	var parsedArgs []Argument
	for i, arg := range args[1:] {
		if len(arg) > 0 && arg[0] == '-' {
			parts := strings.Split(arg, "=")

			if len(parts) == 2 {
				parsedArgs = append(parsedArgs, Argument{Name: parts[0], Value: parts[1]})
			} else {
				parsedArgs = append(parsedArgs, Argument{Name: strconv.Itoa(i), Value: arg})
			}

		} else {
			parsedArgs = append(parsedArgs, Argument{Value: arg})
		}
	}

	return parsedArgs, nil
}
