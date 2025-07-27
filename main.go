package main

import (
	"fmt"
	"os"
	"ExpenseTracker/internal/cli"
)

func main() {
	
	commandHandler := cli.NewCommandHandler()
	err := commandHandler.Execute(os.Args[1:])
	if err != nil {
		fmt.Println("Error executing command:", err.Error())
	}
}