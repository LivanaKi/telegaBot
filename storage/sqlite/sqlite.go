package sqlite

import (
	"database/sql"
	
)

type Storage struct {
	db *sql.DB
}

func New(path string) (error){
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		return nil
	}

	db.Ping()
	return nil
}