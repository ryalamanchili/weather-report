package handlers

import (
	"log"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func API(log *log.Logger, shutdown chan os.Signal, db *mongo.Client) *mux.Router {
	// Remember to close DB connection
	router := mux.NewRouter()
	return router
}
