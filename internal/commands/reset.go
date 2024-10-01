package commands

import (
	"context"
	"gator/internal/state"
	"fmt"
)

func HandlerReset(s *state.GatorState, c Command) error {

	err := s.Db.ResetUserTable(context.Background())
	if err != nil {
		return fmt.Errorf("Error resetting User Table: %v", err)
	}

	err = s.Db.ResetFeedTable(context.Background())
	if err != nil {
		return fmt.Errorf("Error resetting Feed Table: %v", err)
	}

	return nil
}
