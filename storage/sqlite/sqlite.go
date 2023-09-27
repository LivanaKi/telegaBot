package sqlite

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func New(path string) (*Storage, error){
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}

	db.Ping()
}