package command

import (
	"context"
	"fmt"
	"github.com/dabates/gator/internal/rss"
	"github.com/dabates/gator/internal/types"
	"time"
)

func AggHandler(s *types.State, cmd Command) error {
	if cmd.Args == nil || len(cmd.Args) == 0 {
		//return fmt.Errorf("invalid arguments: agg <url>")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	//feed, err := rss.FetchFeed(ctx, cmd.Args[0])
	feed, err := rss.FetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println(feed)

	return nil
}
