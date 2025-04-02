package command

import (
	"context"
	"fmt"
	"github.com/dabates/gator/internal/database"
	"github.com/dabates/gator/internal/types"
	"github.com/google/uuid"
	"time"
)

func RegisterHandler(s *types.State, cmd Command) error {
	if cmd.Args == nil || len(cmd.Args) == 0 {
		return fmt.Errorf("invalid arguments: register <username>")
	}

	u, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err == nil || u.Name != "" {
		return fmt.Errorf("user already exists")
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		Name:      cmd.Args[0],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err := s.Db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	s.Config.SetUser(user.Name)
	fmt.Println("User", user.Name, "created")

	return nil
}
