package state

import (
	"database/sql"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"

	_ "github.com/lib/pq"
)

type GatorState struct {
	Config config.Config
	Db     *database.Queries
}

func NewGatorState() (GatorState, error) {
	cfg, err := config.Read()
	if err != nil {
		return GatorState{}, err
	}

	db, err := sql.Open("postgres", cfg.ConnectionString)
	if err != nil {
		return GatorState{}, fmt.Errorf("Can't open DB for GatorState: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return GatorState{}, fmt.Errorf("DB Ping failed: %v", err)
	}

	return GatorState{
		Config: cfg,
		Db:     database.New(db),
	}, nil
}
