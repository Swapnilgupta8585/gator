package main

import (
	"context"
	"fmt"
	"os"
)

// handlerLogin processes the login command.
func handlerLogin(st *state, cmd command) error {
	// Ensure a username is provided.
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	// Extract the username from the command arguments.
	name := cmd.Args[0]

	// Create a new context.
	ctx := context.Background()

	// Check if the user exists in the database.
	if _, err := st.db.GetUser(ctx, name); err != nil {
		fmt.Printf("%s doesn't exist\n", name)
		os.Exit(1)
	}

	// Set the existing user as the current user in the config.
	if err := st.cfg.SetUser(name); err != nil {
		return err
	}

	// Confirm successful login.
	fmt.Printf("%s is now the active user.\n", name)
	return nil
}


