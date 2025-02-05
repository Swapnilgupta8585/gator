package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Swapnilgupta8585/blog_aggregator/internal/database"

	"github.com/google/uuid"
)

// handlerRegister handles user registration.
func handlerRegister(st *state, cmd command) error {
	// Ensure a name is provided; otherwise, return an error.
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	// Extract the name from the command arguments.
	name := cmd.Args[0]

	// Check if the user already exists.
	ctx := context.Background()
	if _, err := st.db.GetUser(ctx, name); err == nil {
		fmt.Printf("%s already exists\n", name)
		os.Exit(1)
	}

	// Prepare user details for creation.
	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	// Create the user in the database.
	if _, err := st.db.CreateUser(ctx, userParams); err != nil {
		return err
	}

	// Set the newly created user as the current user.
	if err := st.cfg.SetUser(name); err != nil {
		return err
	}

	// Confirm successful registration.
	fmt.Printf("User %s registered successfully.\n", name)
	return nil
}

