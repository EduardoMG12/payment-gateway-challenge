package connection

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDatabase(databaseUrl string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}
	return db, nil
}
