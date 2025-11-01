package main

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func OpenAndInitDB(dsn string) (*sql.DB, error) {
	db, _ := sql.Open("pgx", dsn)

	db.Exec(`CREATE TABLE IF NOT EXISTS movies (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		year INTEGER NOT NULL
	);
	CREATE TABLE IF NOT EXISTS actors (
		id SERIAL PRIMARY KEY,
		movie_id INTEGER REFERENCES movies(id) ON DELETE CASCADE,
		name TEXT NOT NULL
	);`)

	return db, nil
}
