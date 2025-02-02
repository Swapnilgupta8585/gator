package main

import (
	"context"
	"fmt"
)


func handleFeed(st *state, cmd command) error {
	feeds, err := st.db.GetFeeds(context.Background())
	if err != nil {
		return nil
	}
	for index, feed := range feeds{
		fmt.Printf("FEED NUMBER %d\n", index + 1)
		fmt.Printf("NAME: %s\n", feed.Name)
		fmt.Printf("URL: %s\n", feed.Url)
		fmt.Printf("USERNAME: %s\n", feed.Name_2)
		fmt.Println()
	}
	return nil
}