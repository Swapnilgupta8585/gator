package main

import "fmt"

// command struct stores the command name and its arguments
type command struct {
	Name string
	Args []string
}

// commands struct maintains a map linking command names to handler functions
type commands struct {
	registeredCommands map[string]func(*state, command) error
}

// register adds a new command and its corresponding handler function
func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

// run executes the handler function if the command exists
func (c *commands) run(st *state, cmd command) error {
	handlerFunc, exists := c.registeredCommands[cmd.Name]
	// Return an error if the command is not registered
	if !exists {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	// Execute the handler function if it exists
	return handlerFunc(st, cmd)
}

