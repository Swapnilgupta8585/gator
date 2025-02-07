package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Swapnilgupta8585/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

// handleAddFeed handles the "addfeed" command to add a feed to the database and follow it.
func handleAddFeed(st *state, cmd command, user database.User) error {
	// Ensure exactly two arguments are provided: feed name and feed URL.
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	// Retrieve user ID from the database using the username.
	user_id, err := st.db.GetId(context.Background(), user.Name)
	if err != nil {
		return err // Return error if user can't be found.
	}

	// Prepare parameters for creating a new feed in the database.
	feedParam := database.CreateFeedParams{
		ID:        uuid.New(),            // Generate a unique ID for the feed.
		CreatedAt: time.Now(),            // Set the creation time.
		UpdatedAt: time.Now(),            // Set the update time (initially same as creation).
		Name:      cmd.Args[0],           // Feed name from the command argument.
		Url:       cmd.Args[1],           // Feed URL from the command argument.
		UserID:    user_id,               // Link the feed to the user who adds it.
	}

	// Add the new feed to the database.
	_, err = st.db.CreateFeed(context.Background(), feedParam)
	if err != nil {
		return err // Return error if feed creation fails.
	}

	// Retrieve the feed ID using the feed URL.
	feed_id, err := st.db.GetFeedId(context.Background(), feedParam.Url)
	if err != nil {
		return err // Return error if feed ID retrieval fails.
	}

	// Prepare parameters for linking the user to the newly created feed (i.e., follow the feed).
	createFeedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),            // Generate a unique ID for the follow relationship.
		CreatedAt: time.Now(),            // Set the creation timestamp.
		UpdatedAt: time.Now(),            // Set the update timestamp (initially same as creation).
		UserID:    user_id,               // The user following the feed.
		FeedID:    feed_id,               // The feed being followed.
	}

	// Add the feed-follow relationship to the database.
	if _, err = st.db.CreateFeedFollow(context.Background(), createFeedFollowParams); err != nil {
		return err // Return error if creating the follow relationship fails.
	}

	return nil
}
