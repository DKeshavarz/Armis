package commands

import "errors"

var (
	ErrUnknownCommand = errors.New("command not found")
)