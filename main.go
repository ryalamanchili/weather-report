package main

import (
	"encoding/json"
	"fmt"
	. "github.com/derekkenney/location"
	. "github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      initRoutes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	log.Println("HTTP server is listening ,,,")
	log.Fatal(server.ListenAndServe())
}

func initRoutes() *Router {
	router := NewRouter()

	router.HandleFunc("/api", index).Methods("GET")
	router.HandleFunc("/api/location/{long}/{lat}", location).Methods("GET")
	return router

}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	j, _ := json.Marshal("{'data': 'none'}")
	w.Write(j)
}

func location(w http.ResponseWriter, r *http.Request) {
	vars := Vars(r)
	//validate that we have a user id from the request
	var found = true
	var message = ""
	var place = ""

	if vars["long"] == "" {
		message = "Can't retrieve location. Longitude is missing"
		found = false
	}

	if vars["lat"] == "" {
		message = "Can't retrieve location. Latitude is missing"
		found = false
	}

	if found == true {
		place = GetLocation(vars["long"], vars["lat"])
		message = fmt.Sprintf(`{"data": {"place": %s}}`, place)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)

	//if user id is valid, make call to user service looking for user entity
}
