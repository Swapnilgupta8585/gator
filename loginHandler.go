package main

import (
	"context"
	"fmt"
	"os"
)

// login handler function
func handlerLogin(st *state, cmd command) error {
	// Return an error if no arguments are provided
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	// Get the username from the command arguments
	name := cmd.Args[0]
	ctx := context.Background()
	_, err := st.db.GetUser(ctx, name)
	if err != nil {
		fmt.Printf("%s doesn't exist\n", name)
		os.Exit(1)
	}

	// If user exists in the database, update the config file
	err = st.cfg.SetUser(name)
	if err != nil {
		return err
	}
	
	fmt.Printf("%s has been set as the current user\n", name)
	return nil
}

