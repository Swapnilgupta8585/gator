package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}

	s.config.CurrentUserName = cmd.args[0]
	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("%s has been set as the Current User\n", cmd.args[0])
	return nil
}
