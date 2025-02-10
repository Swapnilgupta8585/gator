package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Swapnilgupta8585/blog_aggregator/internal/config"
	"github.com/Swapnilgupta8585/blog_aggregator/internal/database"

	_ "github.com/lib/pq"
)

// state struct manages the program's state
type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	// Read the JSON file and load the config struct
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	// Open a connection to the database
	cfg.DbURL = "postgres://postgres:postgres@localhost:5432/gator"
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Println("Cannot open a database connection")
		os.Exit(1)
	}
	defer db.Close()

	// Initialize database queries
	dbQueries := database.New(db)

	// Create a pointer instance of the state struct
	programState := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	// Create an instance of the commands struct
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	// Register commands with their respective handler functions
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handleAgg)
	cmds.register("feeds",  handleFeed)

	// Register commands that can only be executed by a logged-in user.  
	// The middleware checks if a user exists in the database before allowing access to these commands.  
	cmds.register("addfeed", middlewareLoggedIn(handleAddFeed))
	cmds.register("follow",  middlewareLoggedIn(handleFollow))
	cmds.register("following",  middlewareLoggedIn(handleFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handleUnfollow))
	cmds.register("browse", middlewareLoggedIn(handleBrowse))

	// Ensure at least one command and one argument is provided
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	// Get command name and arguments
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	// Execute the command
	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
