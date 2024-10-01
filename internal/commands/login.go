package commands

import (
	"context"
	"fmt"

	"gator/internal/state"
)

func HandlerLogin(s *state.GatorState, cmd Command) error {
	if len(cmd.Args) == 0 {
		return ErrNoUserName
	}

	name := cmd.Args[0]

	_, err := s.Db.GetUser(context.Background(), name)
	if err != nil {
		return ErrUnknownUser
	}

	err = s.Config.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Println("Logged in")
	return nil
}
