package cli

import (
	"ExpenseTracker/internal/usecase"
)

type CommandHandler struct {
	commands map[string]usecase.Command
}

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{
		commands: map[string]usecase.Command{
			"help": &usecase.HelpCommand{},
		},
	}
}


func (ch *CommandHandler) Execute(args []string) error {
	if len(args) == 0 {
		return NewErrorNoCommandProvided()
	}

	command, exists := ch.commands[args[0]]
	if !exists {
		return NewErrorCommandNotFound(args[0])
	}

	err := command.Execute(args[1:])
	if err != nil {
		return err
	}
	return nil
}
