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

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	cfg.DbURL = "postgres://postgres:postgres@localhost:5432/gator"
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Printf("can not open a database connection")
		os.Exit(1)
	}

	dbQueries := database.New(db)
	st := &state{
		config: cfg,
		db:     dbQueries,
	}
	commandsInstance := &commands{
		mapToHandlers: make(map[string]func(*state, command) error),
	}
	commandsInstance.register("login", handlerLogin)
	commandsInstance.register("register", handlerRegister)
	commandsInstance.register("reset", handlerReset)
	argument := os.Args
	if len(argument) < 2 {
		fmt.Printf("not enough argumnets provided\n")
		os.Exit(1)
	}

	commandInstance := &command{
		name: argument[1],
		args: argument[2:],
	}

	err = commandsInstance.run(st, *commandInstance)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

}
