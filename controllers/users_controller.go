package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	domain "github.com/derekkenney/weather-report/domain"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Error in request body %v", err)
		return
	}

	// create a service class, and pass any DB client dependencies to the constructor
}

func validateUserCredentials(user *domain.User) {
	// lookup the username from DB

}
