package commands

import (
	"context"

	"github.com/DKeshavarz/armis/internal/servise"
)

type PutCommand struct {
	servise servise.ServiceInterfase
}

func (c *PutCommand) Execute(args []string)(string, error){
	ctx := context.Background()
	err := c.servise.Put(ctx, args[0], args[1])
	return "Done", err
}

//****************************************************************************//

type GetCommand struct {
	servise servise.ServiceInterfase
}

func (c *GetCommand) Execute(args []string)(string, error){
	return "exe get", nil
}

//****************************************************************************//

type DelCommand struct {
	servise servise.ServiceInterfase
}

func (c *DelCommand) Execute(args []string)(string, error){
	return "exe del", nil
}