package handlers

import (
	"encoding/json"
	"fmt"
	. "github.com/derekkenney/location"
	. "github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var (
	message = `"data": ""`
	code    int
)

func Routes() *Router {
	log.Println("InitRoutes called")
	r := NewRouter()
	r.HandleFunc("/location/{lat}/{long}", Location)
	return r
}
func Location(w http.ResponseWriter, r *http.Request) {
	vars := Vars(r)
	var lat, _ = strconv.Atoi(vars["lat"])
	var long, _ = strconv.Atoi(vars["long"])
	var found = true
	var place = ""

	if long == 0 {
		log.Println("Can't retrieve location. The longitude argument is missing")
		message = fmt.Sprintf(`{"data": {"message": "Can't retrieve location. Longitude is missing"}}`)
		found = false
		code = 404
	}
	if lat == 0 {
		log.Println("Can't retrieve location. The latitude argument is missing`")
		message = fmt.Sprintf(`{"data": {"message": "Can't retrieve location. Latitude is missing"}}`)
		found = false
		code = 404
	}
	if found == true {
		log.Println("Retrieved location for coordinates")
		place = GetLocation(lat, long)
		message = fmt.Sprintf(`{"data": {"place": %s}}`, place)
		code = 200
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(message)
}
