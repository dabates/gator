package command

import (
	"context"
	"fmt"
	"github.com/dabates/gator/internal/database"
	"github.com/dabates/gator/internal/types"
	"github.com/google/uuid"
	"time"
)

func AddFeedHandler(s *types.State, cmd Command, user database.User) error {
	if cmd.Args == nil || len(cmd.Args) < 2 {
		return fmt.Errorf("invalid arguments: addfeed <feedname> <url>")
	}

	feed, err := s.Db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	// create follow record
	_, err = s.Db.CreateFeedFollows(context.Background(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}
