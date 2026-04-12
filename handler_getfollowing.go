package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Puhkusarvikuono/blogGator/internal/database"
)

func handlerFollowing(s *state, _ command, user database.User) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		fmt.Println("No follows to list.")
		os.Exit(1)
	}

	sumFollows := len(follows)
	fmt.Printf("User has %d follows.\n", sumFollows)
	for i, follow := range follows {
		fmt.Printf("Feed %d:\nName: %s\n", i+1, follow.FeedName)
	}

	return nil
}
