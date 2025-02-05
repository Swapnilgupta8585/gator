package main

import (
	"context"
	"fmt"
)

// handlerUsers processes the "users" command by listing all users.
func handlerUsers(st *state, cmd command) error {
	// Create a new context.
	ctx := context.Background()

	// Retrieve all users from the database.
	userNames, err := st.db.GetUsers(ctx)
	if err != nil {
		return err
	}

	// Display the list of users, marking the current user.
	for _, val := range userNames {
		if st.cfg.CurrentUserName == val {
			fmt.Printf("* %s (current)\n", val)
			continue
		}
		fmt.Printf("* %s\n", val)
	}

	return nil
}
