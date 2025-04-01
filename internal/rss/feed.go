package rss

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator-cli")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rssFeed := &RSSFeed{}
	err = xml.Unmarshal(body, rssFeed)
	if err != nil {
		return nil, err
	}

	// clean yp the html
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for _, i := range rssFeed.Channel.Item {
		i.Title = html.UnescapeString(i.Title)
		i.Description = html.UnescapeString(i.Description)
	}

	return rssFeed, nil
}
