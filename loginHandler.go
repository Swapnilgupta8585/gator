package main

import (
	"context"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}

	name := cmd.args[0]
	ctx := context.Background()
	_, err := s.db.GetUser(ctx, name)
	if err != nil {
		fmt.Printf("%s doesn't exist\n", name)
		os.Exit(1)
	}
	err = s.config.SetUser(name)
	if err != nil {
		return err
	}
	s.config.CurrentUserName = name
	fmt.Printf("%s has been set as the Current User\n", cmd.args[0])
	return nil
}
