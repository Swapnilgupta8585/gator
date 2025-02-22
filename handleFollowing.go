package main

import (
	"context"
	"fmt"

	"github.com/Swapnilgupta8585/gator/internal/database"
)

// handleFollowing handles the "following" command and displays all feeds the user is following.
func handleFollowing(st *state, cmd command, user database.User) error {
	// Retrieve the current user's ID from the database using their username.
	currentUserId, err := st.db.GetId(context.Background(), user.Name)
	if err != nil {
		return err 
	}

	// Fetch all the feeds that the current user is following from the database.
	feedFollowsForUser, err := st.db.GetFeedFollowsForUser(context.Background(), currentUserId)
	if err != nil {
		return err 
	}

	// Loop through each followed feed and print its details.
	for index, item := range feedFollowsForUser {
		fmt.Printf("FEED NUMBER %d\n", index+1) // Display the feed's index (1-based).
		fmt.Printf("FEED NAME : %s\n", item.FeedName) // Display the name of the followed feed.
	}
	return nil
}
