// main.go
package main

import (
	"log"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
)

// App type for weather service
type App struct {
	DB     *mongo.Client
	Router *mux.Router
}

// Initialize creates the database and router objects for the application
func (a *App) Initialize(userName, password, dbName string) {

}

// Run starts the application on the address provided
func (a *App) Run(addr string) {

}

func main() {
	a := App{}
	log.Println("main: Initializing API support")
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)

	log.Println("main: Running application")
	a.Run(":8080")
}
