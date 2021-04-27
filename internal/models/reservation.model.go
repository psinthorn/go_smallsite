package models

import "time"

// Reservation holds reservation data
type Reservation struct {
	FirstName string `json: "first_name"`
	LastName  string `json: "last_name"`
	Email     string `json: "email"`
	Phone     string `json: "phone"`
}

// Room is the room model
type Room struct {
	ID        int       `json: "id"`
	RoomName  string    `json: "room_name"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}

// Restrictions is the restriction model
type Restrictions struct {
	ID              int       `json: "id"`
	RestrictionName string    `json: "restriction_name"`
	CreatedAt       time.Time `json: "created_at"`
	UpdatedAt       time.Time `json: "updated_at"`
}

// Reservations is the reservation model
type Reservations struct {
	ID        int       `json: "id"`
	FirstName string    `json: "first_name"`
	LastName  string    `json: "last_name"`
	Email     string    `json: "email"`
	Phone     string    `json: "phone"`
	RoomID    int       `json: "room_id"`
	Room      Room      `json: "room"`
	Processed int       `json: "processed`
	StartDate time.Time `json: "start_date"`
	EndDate   time.Time `json: "end_date"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}

// RoomRestrictions is the room restriction model
type RoomRestrictions struct {
	ID            int          `json: "id"`
	RoomID        int          `json: "room_id"`
	ReservationID int          `json: "reservation_id"`
	RestrictionID int          `json: "restriction_id"`
	Room          Room         `json: "room"`
	Reservation   Reservations `json: "reservation"`
	Restriction   Restrictions `json: "restriction"`
	StartDate     time.Time    `json: "start_date"`
	EndDate       time.Time    `json: "end_date"`
	CreatedAt     time.Time    `json: "created_at"`
	UpdatedAt     time.Time    `json: "updated_at"`
}
