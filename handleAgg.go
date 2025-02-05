package main

import (
	"context"
	"fmt"
)

// handleAgg processes the "agg" command by fetching and displaying a feed.
func handleAgg(st *state, cmd command) error {
	// Fetch the RSS feed from the given URL.
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	// %+v formats the struct with field names and values.
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}
