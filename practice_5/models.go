package main

type Movie struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	Year       int    `json:"year"`
	ActorCount int    `json:"actor_count"`
}
