package domain

import "time"

type Weather struct {
	Temperature float32   `json:"temperature, omitempty"`
	Description string    `json:"description, omitempty"`
	CreateDate  time.Time `json:"create_date, omitempty"`
}
