package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/derekkenney/weather-report/business/data/weather"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel/api/global"
)

var msg string

type weatherHandler struct {
	db *mongo.Client
}

func (wh *weatherHandler) get(w http.ResponseWriter, r *http.Request) {
	// Get data from request needed for service call
	vars := mux.Vars(r)
	location := vars["location"]

	ctx, span := global.Tracer("service").Start(r.Context(), "handlers.weather.get")
	defer span.End()

	// look up weather by location from business package
	forecast, err := weather.Get(ctx, wh.db, location, time.Now())

	// encode and return results by writing to the response
	//Encoder will write the JSON user object to the response stream
	err = json.NewEncoder(w).Encode(forecast)
	if err != nil {
		msg = fmt.Sprintf("An error occurred marshaling response data. %v\n", err)
		http.Error(w, msg, 500)
		return
	}
}

func (wh *weatherHandler) create(w http.ResponseWriter, r *http.Request) {
	// Read the request body. Close the body
	ctx, span := global.Tracer("service").Start(r.Context(), "handlers.weather.get")
	defer span.End()

	var newWeather weather.Weather
	err := json.NewDecoder(r.Body).Decode(&newWeather)

	if err != nil {
		msg = fmt.Sprintf("Failed to add weather condition %v\n", err)
		http.Error(w, msg, 500)
		return
	}

	newWeather, err = weather.Create(ctx, wh.db, newWeather, time.Now())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Encode the weather as JSON to return
	j, err := json.Marshal(newWeather)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}
