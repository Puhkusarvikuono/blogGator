package main

import (
	"context"
	"fmt"
	"log"
)

func handlerAgg(_ *state, cmd command) error {
	if len(cmd.Args) > 0 {
		fmt.Println("agg takes no arguments")
	}
	feedURL := "https://www.wagslane.dev/index.xml"
	RSSFeed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	println(&RSSFeed)
	return nil
}
