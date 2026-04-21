package main

import (
	"fmt"
	"time"
	"os"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		fmt.Println("agg takes a single argument: time between requests")
		os.Exit(1)
	}
	
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}
	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			fmt.Println("error fetching feed:", err)
		}
		fmt.Printf("next fetch in %v \n", timeBetweenRequests)
	}
	return nil
}
