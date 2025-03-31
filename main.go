package main

import (
	"fmt"
	"github.com/dabates/gator/internal/command"
	"github.com/dabates/gator/internal/config"
	"github.com/dabates/gator/internal/types"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments")
		// print help
		os.Exit(1)
	}

	state := types.State{
		Config: config.Read(),
	}

	commands := &command.Commands{
		Commands: make(map[string]func(state *types.State, command command.Command) error),
	}

	setupHandlers(commands)
	err := commands.Run(&state, command.Command{Name: os.Args[1], Args: os.Args[2:]})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func setupHandlers(commands *command.Commands) {
	commands.Register("login", handlerLogin)
}

func handlerLogin(s *types.State, cmd command.Command) error {
	if cmd.Args == nil || len(cmd.Args) == 0 {
		return fmt.Errorf("invalid arguments: login <username>")
	}

	s.Config.SetUser(cmd.Args[0])

	fmt.Println("Logged in")

	return nil
}
