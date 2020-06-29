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
