package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/Swapnilgupta8585/gator/internal/database"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// scrapeFeeds fetches the next feed from the database, marks it as fetched, retrieves its RSS data, and stores it in the database.
func scrapeFeeds(st *state) error {
	// Retrieve the next feed that needs to be fetched.
	feed, err := st.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	// Prepare parameters to update the feed's fetch timestamp.
	params := database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: time.Now(),
		ID:        feed.ID,
	}

	// Mark the feed as fetched in the database.
	err = st.db.MarkFeedFetched(context.Background(), params)
	if err != nil {
		return err
	}

	// Fetch the RSS feed data from the given URL.
	RssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	// Loop through the RSS feed items and print their titles.
	for _, post := range RssFeed.Channel.Item {
		feed_id, err := st.db.GetFeedId(context.Background(), post.Link)
		if err != nil {
			return err
		}

		published_at, err := dateparse.ParseAny(post.PubDate)
		if err != nil {
			// If parsing fails, use current time or feed fetch time as default
			published_at = time.Now() // or feed.LastFetchedAt.Time
			log.Printf("Could not parse date %q, using default: %v", post.PubDate, published_at)
		}
		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       post.Title,
			Url:         post.Link,
			Description: post.Description,
			PublishedAt: published_at,
			FeedID:      feed_id,
		}
		_, err = st.db.CreatePost(context.Background(), params)
		if err != nil {
			_, err = st.db.CreatePost(context.Background(), params)
			if err != nil {
				// Check if it's a duplicate key error
				if pqErr, ok := err.(*pq.Error); ok {
					// 23505 is the PostgreSQL error code for unique_violation
					if pqErr.Code == "23505" {
						// Duplicate URL - just log and continue
						log.Printf("Skipping duplicate post: %s", post.Title)
						continue
					}
				}
				// If it's any other type of error, log it and return
				log.Printf("Error creating post: %v", err)
				return err
			}
		}

	}
	return nil

}
