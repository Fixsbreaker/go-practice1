package main

import (
	"go_practice2/internal/handlers"
	"go_practice2/internal/middleware"
	"log"
	"net/http"
)

func main() {
	userHandler := http.HandlerFunc(handlers.UserHandler)

	http.Handle("/user", middleware.AuthMiddleware(userHandler))

	port := ":8080"
	log.Printf("Starting server on port %s", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
