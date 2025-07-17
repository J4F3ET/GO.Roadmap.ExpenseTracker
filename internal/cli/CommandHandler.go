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
