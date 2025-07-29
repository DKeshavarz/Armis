package commands

type Command interface {
	Execute(args []string) (string, error)
}

type CommandsServise interface {
	Run() error
	execute(args []string) (string, error)
	register(command string, handle Command)
}