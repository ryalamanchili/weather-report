package main

import (
	"log"
	"os"
	"time"

	"github.com/ardanlabs/service/foundation/database"
	"github.com/gorilla/mux"
)

type Config struct {
	APIHost      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Handler      *mux.Router
	DBUser       string
	DBPassword   string
	DBHost       string
	DBName       string
	DisableTLS   bool
	ReporterURI  string
	ServiceName  string
	Probability  float64
}

// main is the entry point for the weather-report application
func main() {

	// Initialize the config.DB object
	cfg := Config{
		DBHost:     "mongodb://localhost:27017",
		DBName:     "weather",
		DBUser:     "mongo",
		DBPassword: "password1",
	}

	log := log.New(os.Stdout, "WEATHER : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	db, err := database.Open(database.Config{
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		Host:     cfg.DBHost,
		Name:     cfg.DBName,
	})
	if err != nil {
		log.Fatalf("%v connecting to db", err)
		os.Exit(1)
	}
	defer func() {
		log.Printf("main: Database Stopping : %s", cfg.DBHost)
		db.Close()
	}()
}
