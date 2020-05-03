// Package handlers provides functions for dispatching HTTP request paths to handler functions.
package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	. "github.com/derekkenney/weather-report/location"
	. "github.com/gorilla/mux"
)

// Data structure for Encoding Location values to JSON
type Data struct {
	Message string `json: "message"`
	Code    int    `json: "code"`
}

// Routes creates a mux router that maps endpoints like /location to http.HandlerFuncs
func Routes() *Router {
	r := NewRouter()
	r.HandleFunc("/location/longitude/{longitude}/latitude/{latitude}", Location)
	return r
}

//Function Location is the http.HandlerFunc for the request path /location/{long}/{lat}. It expects coordinate arguments.
//If a location is found by using longitude and latitude, a JSON object is returned with location details.
//Also a 200 status code is returned. If the request is missing coordinates, or no location is found, the function
//returns a 404, and empty JSON object.
func Location(w http.ResponseWriter, r *http.Request) {
	data := Data{}
	vars := Vars(r)
	var found = true
	var place = ""

	lat, err := strconv.Atoi(vars["latitude"])
	if err != nil {
		found = false
		log.Printf("\n STACK TRACE\n *****************************")
		log.Printf("%+v\n", err)
		data.Message = "An error occurred getting the location of the latitude."
		data.Code = 500
	}

	long, err := strconv.Atoi(vars["longitude"])
	if err != nil {
		found = false
		log.Printf("\n STACK TRACE\n *****************************")
		log.Printf("%+v\n", err)
		data.Message = "An error occurred getting the location of the longitude."
		data.Code = 500
	}

	if long == 0 && err == nil {
		log.Println("Can't retrieve location. The longitude argument is missing")
		data.Message = "Can't retrieve location. Longitude is missing"
		found = false
		data.Code = 404
	}
	if lat == 0 && err == nil {
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

	if j, err := json.Marshal(data); err != nil {
		log.Printf("\n STACK TRACE\n *******************************")
		fmt.Printf("%+v\n", err)
		http.Error(w, "An error occurred parsing the JSON result of the location", 500)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(data.Code)
		w.Write(j)
	}
}
