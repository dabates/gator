package rss

import (
	"context"
	"fmt"
	"github.com/dabates/gator/internal/types"
)

func ScrapeFeedsHandler(s *types.State) error {
	feed, err := s.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.Db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	data, err := FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for i := range data.Channel.Item {
		fmt.Println("* ", data.Channel.Item[i].Title)
	}

	return nil
}
