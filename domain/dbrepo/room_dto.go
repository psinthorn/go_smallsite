package dbrepo

import "time"

// Room is the room model
type room struct {
	ID          int       `json: "id"`
	RoomTypeId  int       `json: "roomtype_id"`
	RoomName    string    `json: "room_name"`
	RoomNo      string    `json: "room_no"`
	Description string    `json: "description"`
	Status      string    `json: "status"`
	CreatedAt   time.Time `json: "created_at"`
	UpdatedAt   time.Time `json: "updated_at"`
}

// Status is the room status  model
type RoomStatus struct {
	ID          int       `json: "id"`
	Title       string    `json: "status_name"`
	Description string    `json: "description"`
	CreatedAt   time.Time `json: "created_at"`
	UpdatedAt   time.Time `json: "updated_at"`
}

// RoomRestrictions is the room restriction model
type RoomAllotmentStatus struct {
	ID            int         `json: "id"`
	RoomID        int         `json: "room_id"`
	ReservationID int         `json: "reservation_id"`
	RestrictionID int         `json: "restriction_id"`
	Room          Room        `json: "room"`
	Reservation   Reservation `json: "reservation"`
	RoomStatus    RoomStatus  `json: "restriction"`
	StartDate     time.Time   `json: "start_date"`
	EndDate       time.Time   `json: "end_date"`
	CreatedAt     time.Time   `json: "created_at"`
	UpdatedAt     time.Time   `json: "updated_at"`
}
