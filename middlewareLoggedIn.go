package main

import (
	"context"

	"github.com/Swapnilgupta8585/gator/internal/database"
)

// middlewareLoggedIn is a middleware that ensures a user is logged in before executing a handler function.
// It abstracts away repetitive login checks and passes the username to the handler.
func middlewareLoggedIn(handler func(st *state, cmd command, user database.User) error) func(st *state, cmd command) error {
	return func (st *state, cmd command) error {
			user, err := st.db.GetUser(context.Background(), st.cfg.CurrentUserName)
			if err != nil {
						return err
			}

			return handler(st, cmd, user)
	}
}


	


