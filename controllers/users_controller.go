package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	domain "github.com/derekkenney/weather-report/domain"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error in request body %v", err)
		http.Error(w, "An error occurred attempting to login", 500)
		return
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	// create a service class, and pass any DB client dependencies to the constructor
}

func validateUserCredentials(user *domain.User) {
	// lookup the username from DB

}
