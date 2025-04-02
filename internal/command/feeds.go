package command

import (
	"context"
	"fmt"
	"github.com/dabates/gator/internal/types"
)

func FeedsHandler(s *types.State, cmd Command) error {
	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		user, err := s.Db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return err
		}

		fmt.Println("Name:", feed.Name)
		fmt.Println("Url:", feed.Url)
		fmt.Println("User:", user.Name)
		fmt.Println("--------------------------------")
	}

	return nil
}
