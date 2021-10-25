package domain_promotions

import "time"

type promotionType struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Facility    string    `json:"facility"`
	Status      string    `json:"status"` // public, disable
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Facility struct{}
type Amenity struct{}
type Gallery struct{}
type RoomRate struct{}
