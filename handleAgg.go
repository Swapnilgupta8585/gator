package main

import (
	"fmt"
	"time"
)

// handleAgg executes the "agg" command by periodically fetching and displaying feed data.
func handleAgg(st *state, cmd command) error {
	// Ensure the correct number of arguments are provided.
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: agg <time_between_reqs> like 1s, 1m, 1h etc")
	}

	// Parse the provided duration string into a time.Duration value.
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return nil
	}

	fmt.Printf("Collecting feeds every %s\n", cmd.Args[0])

	// Create a ticker that triggers at the specified interval.
	ticker := time.NewTicker(timeBetweenRequests)

	// Continuously scrape feeds each time the ticker signals.
	for ; ; <-ticker.C {
		scrapeFeeds(st)
	}
	
}
