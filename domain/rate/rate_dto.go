package domain

import "time"

type Rate struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	RateTypeId      int       `json:"rate_type_id"`
	RoomTypeId      int       `json:"room_type_id"`
	PromotionTypeId int       `json:"promotion_type-id"`
	Rate            float32   `json:"rate"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type RateType struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
