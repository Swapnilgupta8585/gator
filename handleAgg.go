package main

import (
	"context"
	"fmt"
)




func handleAgg(st *state, cmd command) error{
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil{
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	// +%v adds struct field names to the output as well with thier values
	fmt.Printf("Feed: %+v\n",feed)
	return nil
}