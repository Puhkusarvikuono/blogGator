package main

import (
	"context"
	"fmt"
	"os"
)

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		fmt.Println("No feeds to list.")
		os.Exit(1)
	}

	for _, feed := range feeds {
		user, err := s.db.GetIDName(context.Background(), feed.UserID)
		if err != nil {
			fmt.Println("Invalid user ID.")
			os.Exit(1)
		}
		fmt.Printf("Feed: %s\n URL: %s\nUser name: %s\n", feed.Name, feed.Url, user.Name)
	}

	return nil
}
