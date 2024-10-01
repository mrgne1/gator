package commands

import (
	"context"
	"database/sql"
	"fmt"
	"gator/internal/database"
	"gator/internal/state"
	"log"
	"time"

	"github.com/google/uuid"
)

func HandlerRegister(s *state.GatorState, c Command) error {
	if len(c.Args) < 1 {
		return ErrNoUserName
	}

	name := c.Args[0]
	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Name:      name,
	}

	user, err := s.Db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Register User faied: %v", err)
	}

	err = s.Config.SetUser(name)
	if err != nil {
		return fmt.Errorf("Register user failed to set user: %v", err)
	}

	log.Printf("Register User: %v\n", user)
	fmt.Println("Created User")
	return nil
}
