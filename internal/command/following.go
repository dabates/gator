package command

import (
	"context"
	"fmt"
	"github.com/dabates/gator/internal/types"
)

func FollowingHandler(s *types.State, cmd Command) error {
	user, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return err
	}

	following, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Println("Following by", user.Name+":")
	for _, feed := range following {
		fmt.Println("* ", feed.FeedName)
	}

	return nil
}
