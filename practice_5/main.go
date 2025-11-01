package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/toast?sslmode=disable"
	}

	db, err := OpenAndInitDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/movies", MoviesHandler(db))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
