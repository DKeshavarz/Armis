package commands

import (
	"fmt"
	"strings"

	"github.com/DKeshavarz/armis/internal/servise"
)

type commandsServise struct {
	registery map[string]Command
}

func New()*commandsServise {
	cmd := &commandsServise{
		map[string]Command{},
	}
	mainServise := servise.New()
	cmd.register("get", &GetCommand{mainServise})
	cmd.register("put", &PutCommand{mainServise})
	cmd.register("del", &DelCommand{mainServise})
	return cmd
}

func(c *commandsServise)Run() error{
	var str string
	for {
		fmt.Print("$ ")
		fmt.Scan(&str)
		res, err := c.execute(c.extractor(str))
		
		if err != nil {
			fmt.Println("err:", err)
		}else{
			fmt.Println(res)
		}
	}
}

func(c *commandsServise)register(command string, handle Command){
	c.registery[command] = handle
}

func(c *commandsServise) execute(args []string) (string, error){
	cmd, ok := c.registery[args[0]]
	if !ok {
		return "", ErrUnknownCommand
	}
	
	return cmd.Execute(args[1:])
}

func (c *commandsServise)extractor(input string)[]string{
	return strings.Fields(input)
}