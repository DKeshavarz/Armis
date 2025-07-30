package commands

import (
	"context"
	"fmt"

	"github.com/DKeshavarz/armis/internal/servise"
)

type PutCommand struct {
	servise servise.ServiceInterfase
}

func (c *PutCommand) Execute(args []string)(string, error){
	if len(args) != 2 {
		return  "", ErrNotSuitableArgs
	}
	ctx := context.Background()
	err := c.servise.Put(ctx, args[0], args[1])
	return fmt.Sprintf("%s=%s",args[0], args[1]), err
}

//****************************************************************************//

type GetCommand struct {
	servise servise.ServiceInterfase
}

func (c *GetCommand) Execute(args []string)(string, error){
	if len(args) != 1 {
		return  "", ErrNotSuitableArgs
	}
	ctx := context.Background()
	value, err := c.servise.Get(ctx, args[0])
	return fmt.Sprintf("%s=%s",args[0], value), err
}

//****************************************************************************//

type DelCommand struct {
	servise servise.ServiceInterfase
}

func (c *DelCommand) Execute(args []string)(string, error){
	if len(args) != 1 {
		return  "", ErrNotSuitableArgs
	}
	ctx := context.Background()
	err := c.servise.Delete(ctx, args[0])
	return fmt.Sprintf("key [%s] deleted",args[0]), err
}

//****************************************************************************//
