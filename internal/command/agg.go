package command

import (
	"fmt"
	"github.com/dabates/gator/internal/rss"
	"github.com/dabates/gator/internal/types"
	"time"
)

func AggHandler(s *types.State, cmd Command) error {
	if cmd.Args == nil || len(cmd.Args) == 0 {
		return fmt.Errorf("invalid arguments: agg <time_between_requests>")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Println("Collecting feeds every ", timeBetweenRequests)

	//feed, err := rss.FetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	//if err != nil {
	//	return err
	//}
	//
	//fmt.Println(feed)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		rss.ScrapeFeedsHandler(s)
	}
}
