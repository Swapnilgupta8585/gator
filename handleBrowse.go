package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Swapnilgupta8585/gator/internal/database"
)



// handleBrowse processes the "browse" command to display the posts(limited) that the user follows.
func handleBrowse(st *state, cmd command, user database.User) error {

	// Prepare parameters to fetch the Post for the given user
	params := database.GetPostsForUserParams{
		UserID : user.ID,
		Limit: 2,  // Default limit
	}

	// If a limit is provided as an argument, update the parameter.
	if len(cmd.Args) == 1 {

		// Convert the argument from string to int.
		limit, err := strconv.Atoi(cmd.Args[0]) 
		if err != nil {
			return fmt.Errorf("invalid limit: must be a number")
		}

		// Ensure the limit is not negative.
		if limit < 0 {
			return fmt.Errorf("invalid limit: must be a positive number")
		}

		// Convert int to int32 and update the parameter.
		params.Limit = int32(limit) 
		
	} else if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: browse <limit(a number)>")
	}
	

	// Fetch the posts from the database.
	posts, err := st.db.GetPostsForUser(context.Background(),params)
	if err != nil {
		return err
	}

	// If no posts are found, inform the user
	if len(posts) == 0 {
		fmt.Println("No posts found for the user.")
		return nil
	}

	// Display the posts.
	for index, post := range posts{
		fmt.Printf("%d. %s\n", index + 1, post.Title)
		fmt.Printf("URL: %s\n",post.Url)
		fmt.Printf("Publish Date: %v\n",post.PublishedAt)
		fmt.Printf("Description: %s\n",post.Description)
	}
	return nil
}
