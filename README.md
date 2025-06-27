# gator

A Go-based Blog Aggregator that collects, organizes, and serves blog posts from multiple sources. This project helps centralize various blogs into one place, making content consumption efficient and seamless.

## Features
- Fetch blog posts from multiple sources via RSS feeds
- Store and manage blog data efficiently using PostgreSQL
- Support multiple users
- Built with Golang for high performance

## Installation

### Prerequisites
Ensure you have the following installed:
- **Go** (1.21+ recommended)
- **PostgreSQL** (installed & running)

### Install `gator` CLI
Run the following command to install the `gator` CLI tool:

```sh
go install github.com/Swapnilgupta8585/gator@latest
```

### Ensure `gator` is Accessible
Make sure your Go binary directory is in your `PATH` so you can run `gator` globally:

```bash
export PATH=$(go env GOPATH)/bin:$PATH
```
Add this to your `~/.bashrc` or `~/.zshrc` to make it permanent.

### Set Up Configuration File
Create a configuration file (`.gatorconfig.json`) in your home directory with the necessary database details:

```bash
echo '{
  "db_url": "postgres://<user>:<password>@localhost:5432/gator",
  "current_user_name": ""
}' > ~/.gatorconfig.json
```

Replace `db_url` with your actual PostgreSQL connection string if needed.

## Usage
Run `gator <command>` to interact with the aggregator.

### Available Commands

| Command                                | Description |
|----------------------------------------|-------------|
| `gator register <your_name>`           | Register as a new user(first time) |
| `gator login <your_name>`              | Login as a user |
| `gator users`                          | List all users (highlighting the current user) |
| `gator addfeed <name> <url>`           | Add a new feed |
| `gator feeds`                          | List all available feeds |
| `gator follow <url>`                   | Follow a feed by its URL |
| `gator browse`                         | Display posts from followed feeds (recent entries only) |
| `gator agg <time>`                     | Periodically fetch feed data (see time formats below) |
| `gator unfollow <url>`                 | Unfollow a feed by its URL |
| `gator following`                      | See all the feeds you are followingI

### Time Format for Auto-Fetch
When using the `agg` command, specify time intervals using these formats:
- `1s` - one second
- `1m` - one minute
- `1h` - one hour
Example: `gator agg 30m` will fetch new posts every 30 minutes

## Getting Started
Here's a quick walkthrough to set up and use `gator`:

```sh
# Install gator
go install github.com/Swapnilgupta8585/gator@latest

# Create a config file
nano ~/.gatorconfig.json  # Or use echo command from above

# Register a user
gator register Swapnil

# Add an RSS feed (example)
gator addfeed "golang" "https://blog.golang.org/feed.atom"

# Follow the feed
gator follow "https://blog.golang.org/feed.atom"

# Start the aggregator to collect posts (required before browsing)
gator agg 1m  # Fetches new posts every minute

# Browse your collected posts (run this in a separate terminal)
gator browse  # Shows posts that have been collected
```

### Important Note
- The agg command must be running to collect posts
- Run browse in a separate terminal while agg is running to view posts

Enjoy using `gator` to simplify your blog reading experience! 
