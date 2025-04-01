package command

import (
	"context"
	"fmt"
	"github.com/dabates/gator/internal/types"
)

func UsersHandler(s *types.State, cmd Command) error {
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		current := ""
		if user.Name == s.Config.CurrentUserName {
			current = "(current)"
		}

		fmt.Println("*", user.Name, current)
	}

	return nil
}
