package rss

import (
	"context"
	"fmt"
	"github.com/dabates/gator/internal/database"
	"github.com/dabates/gator/internal/types"
	"github.com/google/uuid"
	"strings"
	"time"
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

	for _, i := range data.Channel.Item {
		fmt.Println("* ", i.Title)

		savePost(i, feed.ID, s.Db)
	}

	return nil
}

func savePost(i RSSItem, feedID uuid.UUID, db *database.Queries) error {
	// deal with already exists, but nothing else
	_, err := db.CreatePost(context.Background(), database.CreatePostParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       i.Title,
		Url:         i.Link,
		Description: i.Description,
		PublishedAt: time.Now(),
		FeedID:      feedID,
	})

	if err != nil && !strings.Contains(err.Error(), " duplicate key value violates unique constraint \"posts_url_key\"") {
		return err
	}

	return nil
}
