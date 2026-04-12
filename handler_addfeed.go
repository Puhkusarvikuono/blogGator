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

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 2 {
		return errors.New("the handler expects a two arguments, the name and url of the feed")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		fmt.Println("User is not in the database")
		os.Exit(1)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	})
	if err != nil {
		fmt.Println("Feed link unsuccessful")
		os.Exit(1)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		fmt.Println("Feed follow unsuccessful.")
	} else {
		fmt.Println("Feed follow successful.")
		fmt.Println(feedFollow)
	}

	fmt.Println("Feed linked successfully")
	fmt.Printf("User name: %s\nFeed name:%s\nFeed url:%s", user.Name, feed.Name, feed.Url)

	return nil
}
