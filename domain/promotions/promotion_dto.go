package domain_promotions

import (
	"time"
)

type promotion struct {
	Id              int           `json:"id"`
	Title           string        `json:"title"`
	Description     string        `json:"description"`
	Price           int           `json:"price"` //start from price
	StartDate       time.Time     `json:"start_date"`
	EndDate         time.Time     `json:"end_date"`
	PromotionTypeId int           `json:"promotion_type_id"`
	PromotionType   PromotionType `json:"promotion_type"`
	Status          string        `json:"status"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}

// Need RateId          int       `json:"rate_id"`

// type RoomRate struct {
// }
