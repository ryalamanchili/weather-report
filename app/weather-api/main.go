package main

import (
	"log"
	"os"
	"time"

	"github.com/ardanlabs/service/foundation/database"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// main is the entry point for the weather-report application
func main() {
	log := log.New(os.Stdout, "WEATHER : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	if err := run(log); err != nil {
		log.Println("main: error:", err)
		os.Exit(1)
	}
}

func run(log *log.Logger) error {
	// =========================================================================
	// Configuration
	var cfg struct {
		Web struct {
			APIHost      string
			ReadTimeout  time.Duration
			WriteTimeout time.Duration
			Handler      *mux.Router
		}
		DB struct {
			User     string
			Password string
			Host     string
			Name     string
		}
		Zipkin struct {
			ReporterURI string
			ServiceName string
			Probability float64
		}
	}
	db, err := database.Open(database.Config{
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Host:     cfg.DB.Host,
		Name:     cfg.DB.Name,
	})
	if err != nil {
		return errors.Wrap(err, "connecting to db")
	}
	defer func() {
		log.Printf("main: Database Stopping : %s", cfg.DB.Host)
		db.Close()
	}()
	return nil
}
