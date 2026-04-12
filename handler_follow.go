package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Puhkusarvikuono/blogGator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) == 0 {
		return errors.New("the handler expects a single argument, an url")
	}
	targetURL := cmd.Args[0]
	var targetFeed database.Feed

	targetUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		fmt.Println("Current user not found in database")
		os.Exit(1)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		fmt.Println("No feeds to follow.")
		os.Exit(1)
	}

	for _, feed := range feeds {
		if targetURL == feed.Url {
			targetFeed = feed
			break
		}
	}

	if targetFeed.Name == "" {
		fmt.Println("No feed with given URL.")
		os.Exit(1)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    targetUser.ID,
		FeedID:    targetFeed.ID,
	})
	if err != nil {
		fmt.Println("Feed follow already exists.")
		os.Exit(1)
	}

	fmt.Println("Feed follow created successfully")
	fmt.Println(feedFollow)
	return nil
}
