package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Swapnilgupta8585/blog_aggregator/internal/database"

	"github.com/google/uuid"
)

func handlerRegister(st *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]
	
	//check if user exits
	ctx := context.Background()
	if _, err := st.db.GetUser(ctx, name); err == nil {
		fmt.Printf("%s already exists\n", name)
		os.Exit(1)
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	_, err := st.db.CreateUser(ctx, userParams)
	if err != nil {
		return err
	}

	err = st.cfg.SetUser(name)
	if err != nil {
		return err
	}
	fmt.Printf("%s is created in the database successfully\n", name)
	return nil
}
