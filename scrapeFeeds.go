package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Swapnilgupta8585/blog_aggregator/internal/database"
)


// scrapeFeeds fetches the next feed from the database, marks it as fetched, retrieves its RSS data, and prints the titles.
func scrapeFeeds(st *state) error {
	// Retrieve the next feed that needs to be fetched.
	feed, err := st.db.GetNextFeedToFetch(context.Background())
	if err != nil{
		return err
	}

	// Prepare parameters to update the feed's fetch timestamp.
	params := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
							Time: time.Now(),
							Valid: true,
						},
		UpdatedAt  :   time.Now(),
		ID       :    feed.ID,
	}

	// Mark the feed as fetched in the database.
	err = st.db.MarkFeedFetched(context.Background(), params)
	if err != nil {
		return err
	}

	// Fetch the RSS feed data from the given URL.
	RssFeed, err := fetchFeed(context.Background(), feed.Url) 
	if err != nil{
		return err
	}

	// Loop through the RSS feed items and print their titles.
	for index, value := range RssFeed.Channel.Item {
		fmt.Printf("%d.  %s\n",index + 1, value.Title)
	}

	return nil
}