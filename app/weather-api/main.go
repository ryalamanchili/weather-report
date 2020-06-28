package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/derekkenney/weather-report/app/foundation/database"
	"github.com/derekkenney/weather-report/app/weather-api/handlers"
	"github.com/gorilla/mux"
)

type Config struct {
	APIHost         string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	Handler         *mux.Router
	DBUser          string
	DBPassword      string
	DBHost          string
	DBName          string
	DisableTLS      bool
	ReporterURI     string
	ServiceName     string
	Probability     float64
	ShutdownTimeout time.Duration
}

// main is the entry point for the weather-report application
func main() {

	// Initialize the config.DB object
	cfg := Config{
		DBHost:          "mongodb://localhost:27017",
		DBName:          "weather",
		DBUser:          "mongo",
		DBPassword:      "password1",
		APIHost:         ":8080",
		ShutdownTimeout: 10 * time.Second,
	}

	log := log.New(os.Stdout, "Weather : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	db, err := database.Open(database.Config{
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		Host:     cfg.DBHost,
		Name:     cfg.DBName,
	})
	if err != nil {
		log.Fatalf("%v connecting to db", err)
	}
	// =========================================================================
	// Start API Service
	log.Println("main: Initializing API support")

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// create and start web server
	api := http.Server{
		Addr:         cfg.APIHost,
		Handler:      handlers.API(log, shutdown, db),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Start the service listening for requests.
	go func() {
		log.Printf("main: API listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	// Shutdown
	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		log.Fatalf("main: %v : Server error", err)

	case sig := <-shutdown:
		log.Printf("main: %v : Start shutdown", sig)

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()

		// Asking listener to shutdown and shed load.
		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			log.Fatalf("main: %v : Shutdown error", err)
		}
	}
}
