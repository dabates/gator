package command

import (
	"context"
	"fmt"
	"github.com/dabates/gator/internal/database"
	"github.com/dabates/gator/internal/types"
	"github.com/google/uuid"
	"time"
)

func FollowHandler(s *types.State, cmd Command) error {
	if cmd.Args == nil || len(cmd.Args) == 0 {
		return fmt.Errorf("invalid arguments: follow <url>")
	}

	if s.Config.CurrentUserName == "" {
		return fmt.Errorf("not logged in")
	}

	user, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.Db.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	params := database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = s.Db.CreateFeedFollows(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Println("Feed Followed:", feed.Name, "by", user.Name)
	return nil
}
