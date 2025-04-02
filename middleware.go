package main

import (
	"context"
	"github.com/dabates/gator/internal/command"
	"github.com/dabates/gator/internal/database"
	"github.com/dabates/gator/internal/types"
)

func middlewareLoggedIn(handler func(s *types.State, cmd command.Command, user database.User) error) func(*types.State, command.Command) error {
	return func(s *types.State, cmd command.Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
