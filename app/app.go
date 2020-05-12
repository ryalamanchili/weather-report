package app

import (
	"log"
	"net/http"
	"time"

	"github.com/derekkenney/weather-report/controllers"
	"github.com/gorilla/mux"
)

func StartApp() error {

	log.Println("StartApp()")
	server := &http.Server{
		Addr:         ":8080",
		Handler:      initRoutes(),
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	log.Println("HTTP server is listening ...")
	return nil
}

func initRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/weather/location/{location/}", controllers.GetLocation).Methods("GET")
	return router
}
