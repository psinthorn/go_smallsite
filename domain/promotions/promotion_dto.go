package domain

import (
	"time"
)

type promotion struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	// RateId          int       `json:"rate_id"`
	StartDate       time.Time `json: "start_date"`
	EndDate         time.Time `json: "end_date"`
	PromotionTypeId int       `json:"promotion_type_id"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// type RoomRate struct {
// }
