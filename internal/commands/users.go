package commands

import (
	"context"
	"fmt"
	"gator/internal/state"
)

func HandlerUsers(s *state.GatorState, c Command) error {
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error listing users: %v", err)
	}

	current := s.Config.User
	fmt.Printf("Current %v\n", current)

	for _, user := range users {
		if user.Name == current {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
			fmt.Printf("* %v\n", user.Name)
		}
	}
	return nil
}
