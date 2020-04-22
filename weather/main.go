package main

import (
	"github.com/derekkenney/weather-report/handlers"
	"log"
	"net/http"
	"time"
)

// main is the entry point for the weather-report application
func main() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      handlers.Routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	log.Println("HTTP server has started. Listening on :8080")
	log.Fatal(server.ListenAndServe())
}