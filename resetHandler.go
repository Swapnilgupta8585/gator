package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()
	err := s.db.DeleteAllUsers(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	err = s.config.SetUser("")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	s.config.CurrentUserName = ""
	fmt.Print("reset successful")
	return nil
}
