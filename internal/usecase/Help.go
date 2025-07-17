package usecase

import (
	"fmt"
)

type HelpCommand struct{}

func (c *HelpCommand) Execute(args []string) error {
	fmt.Println("Inicializando algo con:", args)
	return nil
}
