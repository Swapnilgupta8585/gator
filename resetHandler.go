package main

import (
	"context"
	"fmt"
	"os"
)

// handlerReset processes the reset command by deleting all users and resetting the config.
func handlerReset(st *state, cmd command) error {
	// Create a new context.
	ctx := context.Background()

	// Delete all users from the database.
	if err := st.db.DeleteAllUsers(ctx); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Clear the current user from the config file.
	if err := st.cfg.SetUser(""); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Confirm reset completion.
	fmt.Println("Reset successful.")
	return nil
}

