package dbrepo

import "time"

type RoomType struct {
	ID        int       `json: "id"`
	Title     string    `json: "title"`
	Desc      string    `json: "desc"`
	Room      string    `json: "status"` // public, disable
	Facility  string    `json: "facility"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}

type Facility struct{}
type Amenity struct{}
type Gallery struct{}
type RoomRate struct{}
