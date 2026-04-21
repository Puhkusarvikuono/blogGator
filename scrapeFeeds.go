package main

import (
	"fmt"
	"context"
	"time"
	"errors"
	"github.com/Puhkusarvikuono/blogGator/internal/database"
	"database/sql"
)

func scrapeFeeds(s *state) error {
	// get the next feed
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return errors.New("No feeds to fetch")
	}
	// mark it as fetched
	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams {
		ID: feed.ID,
		UpdatedAt: time.Now(),
		LastFetchedAt: sql.NullTime{Time:time.Now(), Valid: true},
	})
	if err != nil {
		return errors.New("Feed marked unsuccessfully")
	}
	fmt.Println("Feed marked successfully")
	
	// fetch the feed using the URL 
	rss, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	
	// iterate and print the titles
	fmt.Println("fetching feed at", time.Now())
	for _, item := range rss.Channel.Item {
		fmt.Printf("Feed Title: %s\n", item.Title)
	}
	fmt.Printf("found %d posts. ", len(rss.Channel.Item))
	return nil
}
