package weather

import (
	"context"
	"database/sql"
	"time"
)

// Weather represents weather in a location
type Weather struct {
	ID          string    `json:"id"`
	LocationID  string    `json:"location_id"`
	Description string    `json:"description"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}

// CreateReport adds a new weather report to the database.
func (w *Weather) CreateReport(context context.Context, db *sql.DB) {

}
