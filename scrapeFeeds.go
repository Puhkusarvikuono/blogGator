package main

import (
	"fmt"
	"context"
	"time"
	"errors"
	"github.com/Puhkusarvikuono/blogGator/internal/database"
	"database/sql"
	"github.com/google/uuid"
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
  
	scrapeCount := 0

	for _, item := range rss.Channel.Item {
		pubDate, _ := parseTime(item.PubDate)
		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams {
			ID:					uuid.New(),
			CreatedAt:	time.Now(),
			UpdatedAt:	time.Now(),
			Title:			item.Title,
			Url:				item.Link,
			Description:	sql.NullString{String: item.Description, Valid: true},
			PublishedAt:	sql.NullTime{Time:pubDate, Valid: true},
			FeedID:			feed.ID,
		})
		if err != nil {
			continue
		}
		scrapeCount += 1 
	}

	fmt.Printf("Successfully created %d posts\n", scrapeCount)

	return nil
}

func parseTime(s string) (time.Time, error) {
    formats := []string{
        time.RFC1123Z,
        time.RFC1123,
        time.RFC3339,
        // add more as you encounter them
    }
    for _, format := range formats {
        t, err := time.Parse(format, s)
        if err == nil {
            return t, nil
        }
    }
    return time.Time{}, fmt.Errorf("could not parse time: %s", s)
}
