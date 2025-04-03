package command

import (
	"context"
	"fmt"
	"github.com/dabates/gator/internal/types"
	"strconv"
)

func BrowseHandler(s *types.State, cmd Command) error {
	limit := 2

	if len(cmd.Args) == 1 {
		num, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return err
		}
		limit = num
	}

	posts, err := s.Db.GetPosts(context.Background(), int32(limit))
	if err != nil {
		return err
	}

	for _, i := range posts {
		fmt.Println("Title:", i.Title)
		fmt.Println("Date Published:", i.PublishedAt)
		fmt.Println(i.Description)
	}

	return nil
}
