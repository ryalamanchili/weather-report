package handlers

import (
	"log"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func API(log *log.Logger, shutdown chan os.Signal, db *mongo.Client) *mux.Router {
	// Remember to close DB connection

	// Create instance of weather handler
	w := weatherHandler{
		db: db,
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/weather/{location}", w.get).Methods("GET")
	router.HandleFunc("/v1/weather", w.create).Methods("POST")
	return router
}
