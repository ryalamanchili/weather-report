package weather

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// Weather represents weather in a location
type Weather struct {
	ID          string    `json:"id"`
	Location    Location  `json:"location"`
	Temperature float64   `json:"temperature"`
	Description string    `json:"description"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

type Location struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CreateReport adds a new weather report to the database.
func (weather *Weather) CreateReport(ctx context.Context, db *mongo.Client) error {
	// Write package level integration tests. Don't write tests according to the API. Write tests oriented towards the weather
	// package functionality. Follow Ardan Labs example
	// Need to spin up a test database
	// Write a test with a new Weather type, and save to DB
	/// Use db, and get the weather collection
	c, cancel := context.WithCancel(ctx)
	defer cancel()

	// if weather missing id, return err
	if weather.ID == "" {
		err := errors.New("Weather is missing ID. Can't save new record")
		cancel()
		return err
	}

	// if weather missing location id, return err
	if weather.Location.ID == "" {
		err := errors.New("Location is missing ID. Can't save new record")
		cancel()
		return err
	}

	// if weather missing temperature, return err
	if weather.Temperature == 0 {
		err := errors.New("Weather is missing temperature. Can't save new record")
		cancel()
		return err
	}

	collection := db.Database("weather").Collection("weather_conditions")
	insertResult, err := collection.InsertOne(c, weather)

	if err != nil {
		cancel()
		return err
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return nil
}
