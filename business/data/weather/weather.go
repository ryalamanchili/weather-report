package weather

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// Gets weather by location name
func Get(ctx context.Context, db *mongo.Client, location string, now time.Time) (Weather, error) {
	//TODO connect to mongodb database weather, and collection locations`

	// Mocking a Weather entity to return
	w := Weather{
		LocationID:  "wea-10",
		Temperature: 76.5,
		Description: "Mild, and breezy",
	}

	return w, nil
}

// Creates a new weather status for a location
func Create(ctx context.Context, db *mongo.Client, weather Weather, now time.Time) (Weather, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	return Weather{
		LocationID:  "wea-101",
		Temperature: 71.04,
		Description: "mild and breezy",
	}, nil
}
