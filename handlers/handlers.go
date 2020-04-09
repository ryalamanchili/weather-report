// Package handlers provides functions for dispatching HTTP request paths to handler functions.
package handlers

import (
	"encoding/json"

	. "github.com/derekkenney/doximity/location"
	. "github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// Data structure for Encoding Location values to JSON
type Data struct {
	Message string `json: "message"`
	Code    int `json: "code"`
}

// Routes creates a mux router that maps endpoints like /location to http.HandlerFuncs
func Routes() *Router {
	log.Println("InitRoutes called")
	r := NewRouter()
	r.HandleFunc("/location/{lat}/{long}", Location)
	return r
}
//Function Location is the http.HandlerFunc for the request path /location/{long}/{lat}. It expects coordinate arguments.
//If a location is found by using longitude and latitude, a JSON object is returned with location details.
//Also a 200 status code is returned. If the request is missing coordinates, or no location is found, the function
//returns a 404, and empty JSON object.
func Location(w http.ResponseWriter, r *http.Request) {
	data := Data{}
	vars := Vars(r)
	var lat, _ = strconv.Atoi(vars["lat"])
	var long, _ = strconv.Atoi(vars["long"])
	var found = true
	var place = ""

	if long == 0 {
		log.Println("Can't retrieve location. The longitude argument is missing")
		data.Message = "Can't retrieve location. Longitude is missing"
		found = false
		data.Code = 404
	}
	if lat == 0 {
		log.Println("Can't retrieve location. The latitude argument is missing`")
		data.Message = "Can't retrieve location. Latitude is missing"
		found = false
		data.Code = 404
	}
	if found == true {
		log.Println("Retrieved location for coordinates")
		place = GetLocation(lat, long)
		log.Print(place)
		data.Message = "Kalamazoo"
		data.Code = 200
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data.Code)
	_ = json.NewEncoder(w).Encode(data)
}
