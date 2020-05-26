package app

import (
	"log"
	"net/http"
	"time"

	"github.com/derekkenney/weather-report/controllers"
	"github.com/gorilla/mux"
)

var (
	err error
)

func StartApp() {
	log.Println("StartApp()")
	server := &http.Server{
		Addr:         ":8080",
		Handler:      initRoutes(),
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	log.Println("HTTP server is listening ...")
	log.Fatal(server.ListenAndServe())
}

func initRoutes() *mux.Router {
	log.Println("initRoutes")
	router := mux.NewRouter()
	//router.HandleFunc("user/register", controllers.Register).Methods("POST")
	router.HandleFunc("/user/login", controllers.Login).Methods("POST")
	return router
}
