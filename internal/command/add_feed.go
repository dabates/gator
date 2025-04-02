package command

import (
	"context"
	"fmt"
	"github.com/dabates/gator/internal/database"
	"github.com/dabates/gator/internal/types"
	"github.com/google/uuid"
	"time"
)

func AddFeedHandler(s *types.State, cmd Command) error {
	if cmd.Args == nil || len(cmd.Args) < 2 {
		return fmt.Errorf("invalid arguments: addfeed <feedname> <url>")
	}

	user, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(), //uuid.NullUUID{UUID: uuid.New(), Valid: true},
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	feed, err := s.Db.CreateFeed(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}
