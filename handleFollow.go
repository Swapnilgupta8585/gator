package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Swapnilgupta8585/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

// handleFollow processes the "follow" command, allowing users to follow a feed by its URL.
func handleFollow(st *state, cmd command) error {
	// Ensure the correct number of arguments are provided.
	if len(cmd.Args) != 1 {
		return fmt.Errorf("ERROR | usage: follow <url>")
	}

	// Retrieve user ID from the database using the current username.
	user_id, err := st.db.GetId(context.Background(), st.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	// Get the feed URL from the command arguments.
	url := cmd.Args[0]

	// Retrieve the feed ID from the database using the feed URL.
	feed_id, err := st.db.GetFeedId(context.Background(), url)
	if err != nil {
		return nil // Return early if feed is not found.
	}

	// Create parameters to establish a relationship between user and feed.
	feedFollowParam := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user_id,
		FeedID:    feed_id,
	}

	// Insert the feed-follow relationship into the database.
	feedFollow, err := st.db.CreateFeedFollow(context.Background(), feedFollowParam)
	if err != nil {
		return err
	}

	// Print details about the newly followed feed.
	fmt.Printf("FEED NAME : %s\n", feedFollow[0].FeedName)
	fmt.Printf("USERNAME : %s\n", feedFollow[0].UserName)
	return nil
}
