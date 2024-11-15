package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Swapnilgupta8585/blog_aggregator/internal/database"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the register handler expects a single argument, the username")
	}

	ctx := context.Background()
	name := cmd.args[0]

	//check if user exits
	if _, err := s.db.GetUser(ctx, name); err == nil {
		fmt.Printf("%s already exists\n", name)
		os.Exit(1)
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	_, err := s.db.CreateUser(ctx, userParams)
	if err != nil {
		return err
	}

	err = s.config.SetUser(name)
	if err != nil {
		return err
	}
	s.config.CurrentUserName = name
	fmt.Printf("%s is created in the database successfully\n", name)
	return nil
}
