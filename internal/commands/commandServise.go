package commands

import "fmt"

type commandsServise struct {
	registery map[string]Command
}

func New()*commandsServise {
	cmd := &commandsServise{
		map[string]Command{},
	}
	cmd.register("get", &GetCommand{})
	cmd.register("put", &PutCommand{})
	cmd.register("del", &DelCommand{})
	return cmd
}

func(c *commandsServise)Run() error{
	var str string
	for {
		fmt.Print("$ ")
		fmt.Scan(&str)
		res, _ := c.execute([]string{str})
		fmt.Println(res)
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