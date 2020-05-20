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
}

func validateUserCredentials(user *User) {
	// lookup the username from DB
}
