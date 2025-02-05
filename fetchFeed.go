package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"
)

// RSSFeed represents the structure of an RSS feed.
// The XML tags map fields to their corresponding elements in the XML feed.
type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`       // Title of the RSS feed
		Link        string    `xml:"link"`        // URL of the RSS feed source
		Description string    `xml:"description"` // Short description of the feed
		Item        []RSSItem `xml:"item"`        // List of articles/items in the feed
	} `xml:"channel"` // The main container for feed data
}

// RSSItem represents a single news article or blog post within an RSS feed.
type RSSItem struct {
	Title       string `xml:"title"`       // Title of the article
	Link        string `xml:"link"`        // URL of the full article
	Description string `xml:"description"` // Summary or excerpt of the article
	PubDate     string `xml:"pubDate"`     // Publication date of the article
}

// fetchFeed retrieves and parses an RSS feed from a given URL.
// It makes an HTTP request, reads the response, and unmarshals the XML data into an RSSFeed struct.
func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// Create an HTTP GET request with the provided context.
	// The context allows for timeout/cancellation if needed.
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err // Return an error if the request couldn't be created
	}

	// Set a custom User-Agent to identify our request (some servers block unknown clients).
	req.Header.Set("User-Agent", "gator")

	// Initialize an HTTP client with a timeout to prevent hanging requests.
	client := http.Client{
		Timeout: 10 * time.Second, // Abort request if it takes more than 10 seconds
	}

	// Send the request and receive a response.
	res, err := client.Do(req)
	if err != nil {
		return nil, err // Return an error if the request fails (e.g., network issue)
	}
	defer res.Body.Close() // Ensure the response body is closed to free resources.

	// Read the response body into a byte slice.
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err // Return an error if reading the response fails
	}

	// Declare a variable to store the parsed RSS feed data.
	var rssFeed RSSFeed

	// Parse the XML response and populate the rssFeed struct.
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return nil, err // Return an error if XML parsing fails
	}

	// Some RSS feeds encode special characters (e.g., &amp; for "&").
	// Unescape them to restore human-readable text.
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)

	// Loop through each RSS item and unescape special characters.
	for i, item := range rssFeed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)             // Fix special characters in the title
		item.Description = html.UnescapeString(item.Description) // Fix special characters in the description
		rssFeed.Channel.Item[i] = item                          // Update the struct with the cleaned data
	}

	// Return the parsed RSS feed data.
	return &rssFeed, nil
}

