package commands

import "errors"

var (
	ErrUnknownCommand = errors.New("command not found")
	ErrNotSuitableArgs = errors.New("the arguments of this command are more or less than needed args")
)