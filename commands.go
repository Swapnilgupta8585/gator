package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	mapToHandlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.mapToHandlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {

	handlerFunc, exists := c.mapToHandlers[cmd.name]
	if !exists {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	err := handlerFunc(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
