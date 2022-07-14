package sqlite

import (
	"errors"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewSQLite3(dbPath string) (*sqlx.DB, error) {
	if _, err := os.Stat(dbPath); err != nil {
		return nil, errors.New("the database on the route `"+dbPath+"` does not exist")
	}

	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
