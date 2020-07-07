package weather

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

// GetWeatherForLocation returns a weather report for a location. Function takes locationName as an argument
func (weather *Weather) GetWeatherForLocation(ctx context.Context, db *mongo.Client, locationName string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// if location name is missing, return err
	if locationName == "" {
		err := errors.New("Location name is missing. Can't return weather report")
		cancel()
		log.Printf("service: Weather. %v", err)
		return err
	}

	filter := bson.D{{"location.name", locationName}}
	collection := db.Database("weather").Collection("weather_conditions")
	err := collection.FindOne(ctx, filter).Decode(&weather)
	if err != nil {
		cancel()
		log.Printf("service: Weather. %v", err)
		return err
	}

	// want to decode into the weather type, and then return it
	fmt.Printf("Found a single document: %+v\n", weather)
	return nil
}

// CreateReport adds a new weather report to the database.
func (weather *Weather) CreateReport(ctx context.Context, db *mongo.Client) error {
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
