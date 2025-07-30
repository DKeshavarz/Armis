package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/DKeshavarz/armis/internal/servise"
)

type commandsServise struct {
	registery map[string]Command
}

func New(mainServise servise.ServiceInterfase)*commandsServise {
	cmd := &commandsServise{
		map[string]Command{},
	}

	cmd.register("get", &GetCommand{mainServise})
	cmd.register("put", &PutCommand{mainServise})
	cmd.register("del", &DelCommand{mainServise})

	return cmd
}

func(c *commandsServise)Run() error{
	for {
		fmt.Print("$ ")
		commands, err := c.read()
		if err != nil{
			return fmt.Errorf("error in reading: %s", err)
		}

		if len(commands) == 1 && commands[0] == "exit" {
			break
		}

		res, err := c.execute(commands)
		c.show(res, err)
	}
	return nil
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

func (c *commandsServise)read()([]string, error){
	reader := bufio.NewReader(os.Stdin)
	msg, err := reader.ReadString('\n')
	return strings.Fields(msg), err
}

func (c *commandsServise)show(respons string, err error){
	if err != nil {
		fmt.Printf(Red + "err -> %s\n" + Reset, err)
		return
	}
	fmt.Println(respons)
}