package tysqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	db *sql.DB
}

func NewSQLliteCache() *SQLite {
	return &SQLite{}
}

func (slte *SQLite) Open(db string) error {
	var err error
	slte.db, err = sql.Open("sqlite3", db)
	if err != nil {
		return err
	}
	return nil
}

func (slte *SQLite) DB() *sql.DB {
	return slte.db
}

func (slte *SQLite) Close() error {
	return slte.db.Close()
}
