package commands

import (
	"errors"
	"gator/internal/state"
)

var ErrNoUserName error = errors.New("No Username")
var ErrNoURL error = errors.New("No URL")
var ErrUnknownUser error = errors.New("Unknown user")
var ErrUnknownCommand  error = errors.New("Unknown command")

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	dispatch map[string]func(*state.GatorState, Command) error
}

func NewCommands() Commands {
	cmd := Commands{
		dispatch: make(map[string]func(*state.GatorState, Command) error),
	}
	return cmd
}

func (c *Commands) Register(name string, f func(*state.GatorState, Command) error) {
	c.dispatch[name] = f
}

func (c *Commands) Run(s *state.GatorState, cmd Command) error {
	fn, ok := c.dispatch[cmd.Name]
	if !ok {
		return ErrUnknownCommand
	}
	return fn(s, cmd)
}
