package main

import (
	"context"
	"fmt"
	"os"
)

// login handler function
func handlerLogin(st *state, cmd command) error {
	// if length of cmd.Args method if zero return error
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	// Get the Argument for the login function (name of the user)
	name := cmd.Args[0]
	ctx := context.Background()
	_, err := st.db.GetUser(ctx, name)
	if err != nil {
		fmt.Printf("%s doesn't exist\n", name)
		os.Exit(1)
	}

	// if user exists in Database, Set the user in the config file 
	err = st.config.SetUser(name)
	if err != nil {
		return err
	}
	// Set the user in the Config struct in the state struct
	st.config.CurrentUserName = name
	fmt.Printf("%s has been set as the Current User\n", cmd.Args[0])
	return nil
}
