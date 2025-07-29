package commands

import "github.com/DKeshavarz/armis/internal/storage"

type PutCommand struct {
	Store *storage.StorageInterface
}

func (c *PutCommand) Execute(args []string)(string, error){
	return "exe put", nil
}

//****************************************************************************//

type GetCommand struct {
	Store *storage.StorageInterface
}

func (c *GetCommand) Execute(args []string)(string, error){
	return "exe get", nil
}

//****************************************************************************//

type DelCommand struct {
	Store *storage.StorageInterface
}

func (c *DelCommand) Execute(args []string)(string, error){
	return "exe del", nil
}