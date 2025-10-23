package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitSQLite(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	schema := `CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		body TEXT,
		user_id INTEGER
	);`

	_, err = db.Exec(schema)
	return db, err
}
