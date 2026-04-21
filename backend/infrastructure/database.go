package infrastructure

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

func NewSQLiteDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id       INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT    NOT NULL UNIQUE,
			password TEXT    NOT NULL
		)
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
