package domain

type Location struct {
	LocationId   string `json:"location_id"`
	LocationName string `json:"location_name"`
	Longitude    int64  `json:"longitude, omitempty"`
	Latitude     int64  `json:"latitude, omitempty"`
}
