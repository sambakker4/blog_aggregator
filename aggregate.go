package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sambakker4/gator/internal/database"
)

func scrapeFeeds(s *state) error {
	dbFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.db.MarkFetchedFeed(context.Background(), database.MarkFetchedFeedParams{
		ID:        dbFeed.ID,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	feed, err := fetchFeed(context.Background(), dbFeed.Url)
	if err != nil {
		return err
	}

	fmt.Println("---------------------------------------------------")
	fmt.Println(feed.Channel.Title)
	fmt.Println("---------------------------------------------------")
	fmt.Println(feed.Channel.Description)
	fmt.Println("---------------------------------------------------")
	fmt.Println("Link:", feed.Channel.Link)
	fmt.Println("---------------------------------------------------")

	for _, item := range feed.Channel.Item {
		fmt.Println("---------------------------------------------------")
		fmt.Println(item.Title)
		fmt.Println()
		fmt.Println(item.Description)
		fmt.Println()
		fmt.Println("---------------------------------------------------")
	}
	return nil
}
