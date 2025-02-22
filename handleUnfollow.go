package main

import (
	"context"
	"fmt"

	"github.com/Swapnilgupta8585/gator/internal/database"
)


// handleUnFollow processes the "Unfollow" command, allowing users to remove a feed from their followed list.
func handleUnfollow(st *state, cmd command, user database.User) error {
	// Ensure the correct number of arguments are provided.
	if len(cmd.Args) != 1 {
		return fmt.Errorf("ERROR | usage: Unfollow <url>")
	}

	// Extract the feed URL from the command arguments.
	url := cmd.Args[0]

	// Retrieve the Feed ID from the database using the given URL.
	feed_id, err := st.db.GetFeedId(context.Background(), url)
	if err != nil {
		return err
	}

	// Define parameters for deleting the follow record.
	params := database.DeleteFeedFollowRecordParams{
		UserID: user.ID,
		FeedID: feed_id,
	}

	// Attempt to delete the follow record from the database.
	err = st.db.DeleteFeedFollowRecord(context.Background(), params)
	if err != nil {
		return err 
	}

	return nil
}
