package command

import (
	"context"
	"fmt"
	"github.com/dabates/gator/internal/types"
)

func LoginHandler(s *types.State, cmd Command) error {
	if cmd.Args == nil || len(cmd.Args) == 0 {
		return fmt.Errorf("invalid arguments: login <username>")
	}

	u, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err != nil || u.Name == "" {
		return fmt.Errorf("user does not exist")
	}

	s.Config.SetUser(u.Name)

	fmt.Println("Logged in", u.Name)

	return nil
}
