package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

const (
	dev = "postgresql://postgres:password@localhost:5432/TestTB?sslmode=disable"
)

func Open() (*Database, error) {
	db, err := sql.Open("postgres", dev)

	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
