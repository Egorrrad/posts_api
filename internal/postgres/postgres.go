package postgres

import (
	"database/sql"
)

type PostgresStorage struct {
	// реализация для PostgreSQL
	DB *sql.DB
}

func OpenDB(dsn string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", dsn) // right or not?
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	store := &PostgresStorage{}
	store.DB = db

	return store, nil
}

func CloseDB(db *sql.DB) {
	db.Close()
}
