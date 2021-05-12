package dbrepo

import (
	"time"

	"github.com/psinthorn/go_smallsite/domain/rooms"
)

// Reservations is the reservation model
type reservation struct {
	ID        int        `json: "id"`
	FirstName string     `json: "first_name"`
	LastName  string     `json: "last_name"`
	Email     string     `json: "email"`
	Phone     string     `json: "phone"`
	RoomID    int        `json: "room_id"`
	Room      rooms.Room `json: "room"`
	Status    string     `json: "status"` // available, stay, clean, maintenance
	StartDate time.Time  `json: "start_date"`
	EndDate   time.Time  `json: "end_date"`
	CreatedAt time.Time  `json: "created_at"`
	UpdatedAt time.Time  `json: "updated_at"`
}
