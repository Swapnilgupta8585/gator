package main

import (
	"context"
	"fmt"
)

func handlerUsers(st *state, cmd command) error {
	ctx := context.Background()
	userNames, err := st.db.GetUsers(ctx)
	if err != nil {
		return err
	}
	for _, val := range userNames {
		if st.cfg.CurrentUserName == val {
			fmt.Printf("* %s (current)\n", val)
			continue
		}
		fmt.Printf("* %s\n", val)
	}
	return nil
}
