package command

import (
	"fmt"
	. "github.com/dabates/gator/internal/types"
)

type Commands struct {
	Commands map[string]func(state *State, command Command) error
}

func (c *Commands) Register(name string, f func(state *State, command Command) error) {
	c.Commands[name] = f
}

func (c *Commands) Run(s *State, command Command) error {
	toRun, ok := c.Commands[command.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", command.Name)
	}

	err := toRun(s, command)
	if err != nil {
		return err
	}

	return nil
}
