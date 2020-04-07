package main

import (
	"encoding/json"
	"fmt"
	. "github.com/derekkenney/location"
	. "github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
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
}

func location(w http.ResponseWriter, r *http.Request) {
	vars := Vars(r)
	var lat, _ = strconv.Atoi(vars["lat"])
	var long, _ = strconv.Atoi(vars["long"])
	var found = true
	var message = `"data": ""`
	var place = ""
	var code int

	if long == 0 {
		message = "Can't retrieve location. Longitude is missing"
		found = false
	}
	if lat == 0 {
		message = "Can't retrieve location. Latitude is missing"
		found = false
	}
	if found == true {
		place = GetLocation(lat, long)
		message = fmt.Sprintf(`{"data": {"place": %s}}`, place)
	}
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Printf("An error occurred with location handler %v", err.Error())
		code = 500
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(message))
}
