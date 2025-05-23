package main

import (
	"fmt"
	commands "github.com/neel004/github-user-activity/commands"
)

func main() {
	command := commands.NewRootCommand()
	if err := command.Execute(); err != nil {
		fmt.Errorf("error encountered while executing command, %w", err)
	}
}
