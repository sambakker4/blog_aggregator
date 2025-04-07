package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

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

	fmt.Println(feed.Channel.Title)
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println(feed.Channel.Description)
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println("Link:", feed.Channel.Link, "\n")

	for _, item := range feed.Channel.Item {
		layout, err := detectTimeLayout(item.PubDate)
		if err != nil {
			fmt.Println(item.PubDate)
			return err
		}
		pubDate, err := time.Parse(layout, item.PubDate)
		if err != nil {
			return err
		}
		err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: sql.NullTime{
				Time:  pubDate,
				Valid: true,
			},
			FeedID: dbFeed.ID,
		})

		duplicatePostUrlError := `pq: duplicate key value violates unique constraint "posts_url_key"`

		if err != nil && err.Error() != duplicatePostUrlError {
			return err
		}
	}
	return nil
}

func detectTimeLayout(s string) (string, error) {
	layouts := []string{
		"2006-01-02",
		"02/01/2006",
		"01/02/2006",
		"2006-01-02 15:04:05",
		"02 Jan 2006",
		"January 2, 2006",
		"Mon, 02 Jan 2006 15:04:05 -0700",
	}
	for _, layout := range layouts {
		if _, err := time.Parse(layout, s); err == nil {
			return layout, nil
		}
	}
	return "", errors.New("could not parse time layout")
}
