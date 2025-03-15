package dtos

import "time"

type Product struct {
	Id               string    `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Price            float64   `json:"price"`
	Quantity         int16     `json:"quantity"`
	Rating           float32   `json:"rating"`
	Review           string    `json:"review"`
	TotalNoOfRatings int64     `json:"total_no_of_ratings"`
	TotalNoOfReviews int64     `json:"total_no_of_reviews"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	DeletedAt        time.Time `json:"deleted_at"`
}

type Products struct {
	Products []*Product `json:"products"`
}
