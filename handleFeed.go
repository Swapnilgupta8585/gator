package main

import (
	"context"
	"fmt"
)

// handleFeed retrieves and prints all the feeds stored in the database.
func handleFeed(st *state, cmd command) error {
	// Retrieve all feeds from the database.
	feeds, err := st.db.GetFeeds(context.Background())
	if err != nil {
		return nil 
	}

	// Loop through each feed and print its details.
	for index, feed := range feeds {
		fmt.Printf("FEED NUMBER %d\n", index+1) // Display feed index (1-based).
		fmt.Printf("NAME: %s\n", feed.Name)     // Feed title or name.
		fmt.Printf("URL: %s\n", feed.Url)       // RSS feed URL.
		fmt.Printf("USERNAME: %s\n", feed.Name_2) // User associated with the feed.
		fmt.Println() // Print an empty line for better readability.
	}

	return nil
}
