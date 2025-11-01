package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func MoviesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		yearMin, _ := strconv.Atoi(r.URL.Query().Get("year_min"))
		yearMax, _ := strconv.Atoi(r.URL.Query().Get("year_max"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

		if limit == 0 {
			limit = 50
		}

		sqlStr := `
			SELECT m.id, m.title, m.year, COUNT(a.id) AS actor_count
			FROM movies m
			LEFT JOIN actors a ON a.movie_id = m.id
			WHERE ($1 = 0 OR m.year >= $1)
			  AND ($2 = 0 OR m.year <= $2)
			GROUP BY m.id
			ORDER BY m.year DESC
			LIMIT $3 OFFSET $4;`

		start := time.Now()
		rows, _ := db.Query(sqlStr, yearMin, yearMax, limit, offset)
		dur := time.Since(start)

		var movies []Movie
		for rows.Next() {
			var m Movie
			rows.Scan(&m.ID, &m.Title, &m.Year, &m.ActorCount)
			movies = append(movies, m)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Query-Time", dur.String())
		json.NewEncoder(w).Encode(movies)
	}
}
