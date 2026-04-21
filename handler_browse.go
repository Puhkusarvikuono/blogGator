package main

import (
	"fmt"
	"context"
	"github.com/Puhkusarvikuono/blogGator/internal/database"
	"strconv"
	"os"
)

 
func handlerBrowser(s *state, cmd command, user database.User) error {
	var limit int32

	limit = 2

	if len(cmd.Args) > 0 {
		n, err := strconv.Atoi(cmd.Args[0])
		if err == nil {
			limit = int32(n)
		}
	}
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		fmt.Println("No feeds found for user")
		os.Exit(1)
	}
	

	for _, follow := range feedFollows {
		posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams {
			FeedID:	follow.FeedID,
			Limit: limit, 
		})
		if err != nil {
			fmt.Printf("Unable to get user posts for feed: %s", follow.FeedName)
			continue
		}
		for i, post := range posts {
			fmt.Printf("----\nPost number: %d\n----\nCreated at: %v\nUpdated at: %v\nTitle: %s\nUrl: %s\nFeed ID: %v\n", 
				i, 
				post.CreatedAt, 
				post.UpdatedAt, 
				post.Title, 
				post.Url, 
				post.FeedID)
			if post.Description.Valid {
				fmt.Printf("Description: %s\n", post.Description.String)
			}
			if post.PublishedAt.Valid {
				fmt.Printf("Published at: %v\n", post.PublishedAt.Time)
			}
		}
	}

	return nil
}
