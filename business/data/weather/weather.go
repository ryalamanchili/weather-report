package weather

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// Create inserts a new user into the database.
func Get(ctx context.Context, db *mongo.Client, location string, now time.Time) (Weather, error) {
	// Mocking a Weather entity to return
	w := Weather{
		LocationID:  "wea-10",
		Temperature: 76.5,
		Description: "Mild, and breezy",
	}

	return w, nil
}
