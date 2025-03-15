package models

import (
	"time"
)

type Product struct {
	Id               int
	Name             string
	Description      string
	Price            float64
	Quantity         int16
	Rating           float32
	Review           string
	TotalNoOfRatings int64
	TotalNoOfReviews int64
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}
