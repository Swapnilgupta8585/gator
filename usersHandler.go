package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	ctx := context.Background()
	userNames, err := s.db.GetUsers(ctx)
	if err != nil {
		return err
	}
	for _, val := range userNames {
		if s.config.CurrentUserName == val {
			fmt.Printf("* %s (current)\n", val)
			continue
		}
		fmt.Printf("* %s\n", val)
	}
	return nil
}
