package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Swapnilgupta8585/blog_aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	st := &state{
		config: cfg,
	}
	commandsInstance := &commands{
		mapToHandlers: make(map[string]func(*state, command) error),
	}
	commandsInstance.register("login", handlerLogin)
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
