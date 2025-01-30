package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(st *state, cmd command) error {
	ctx := context.Background()
	err := st.db.DeleteAllUsers(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	err = st.cfg.SetUser("")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Print("reset successful")
	return nil
}
