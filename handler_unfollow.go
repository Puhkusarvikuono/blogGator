package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Puhkusarvikuono/blogGator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) == 0 {
		return errors.New("the handler expects a single argument, an url")
	}
	targetURL := cmd.Args[0]
	var targetFeed database.Feed

	
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		fmt.Println("No feeds to unfollow.")
		os.Exit(1)
	}

	for _, feed := range feeds {
		if targetURL == feed.Url {
			targetFeed = feed
			break
		}
	}
	
	if targetFeed.Name == "" {
		fmt.Println("No feed found to unfollow")
		os.Exit(1)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		fmt.Println("No follows to unfollow.")
		os.Exit(1)
	}

	for _, follow := range follows {
		if follow.UserID == targetFeed.UserID {
			err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams {
				UserID: follow.UserID,
				FeedID: follow.FeedID,
			})
			if err != nil {
				fmt.Println("Feed unfollow unsuccessful")
				return err
			}
			fmt.Println("Feed unfollow successful")
			return nil 
			}
	}
	
	return errors.New("Feed follow not found")
}
