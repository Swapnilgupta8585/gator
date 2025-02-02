package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Swapnilgupta8585/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handleAddFeed(st *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}
	username, err := st.db.GetUser(context.Background(), st.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	user_id, err := st.db.GetId(context.Background(), username)
	if err != nil{
		return err
	}

	feedParam := database.CreateFeedParams{
		ID    :    uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name   :   cmd.Args[0],
		Url     :  cmd.Args[1],
		UserID   : user_id,
	}
	_, err = st.db.CreateFeed(context.Background(), feedParam)
	if err != nil{
		return err
	}

	return nil
}