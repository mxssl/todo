package db

import (
	"github.com/jmoiron/sqlx"
)

// DB ...
type DB struct {
	*sqlx.DB
}

// NewDB ...
func NewDB(dbConnString string) (*DB, error) {
	db, err := sqlx.Connect("postgres", dbConnString)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

// Close ...
func (d *DB) Close() error {
	return d.DB.Close()
}
