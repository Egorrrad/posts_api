package postgres

import (
	"database/sql"
)

type PostgresStorage struct {
	// реализация для PostgreSQL
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn) // right or not?
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
