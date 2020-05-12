// Package handlers provides functions for dispatching HTTP request paths to handler functions.
package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Data structure for Encoding Location values to JSON
type Data struct {
	Message string `json: "message"`
	Code    int    `json: "code"`
}

//GetLocation is the http.HandlerFunc for the request path /location/{long}/{lat}. It expects coordinate arguments.
//If a location is found by using longitude and latitude, a JSON object is returned with location details.
//Also a 200 status code is returned. If the request is missing coordinates, or no location is found, the function
//returns a 404, and empty JSON object.
func GetLocation(w http.ResponseWriter, r *http.Request) {
	data := Data{}
	//	vars := mux.Vars(r)
	var found = false
	var place = ""

	// call weather service passing in location argument

	if found == true {
		log.Println("Retrieved location for coordinates")
		log.Print(place)
		data.Message = "Kalamazoo"
		data.Code = 200
	}

	j, err := json.Marshal(data)
	if err != nil {
		log.Printf("\n STACK TRACE\n *******************************")
		fmt.Printf("%+v\n", err)
		http.Error(w, "An error occurred parsing the JSON result of the location", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)

}
