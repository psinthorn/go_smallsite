package reservations

import "time"

// // Reservation holds reservation data
// type Reservation struct {
// 	FirstName string `json: "first_name"`
// 	LastName  string `json: "last_name"`
// 	Email     string `json: "email"`
// 	Phone     string `json: "phone"`
// }

// Room is the room model
// Room infomation
type Room struct {
	ID              int       `json: "id"`
	RoomCatID       int       `json: "room_cat_id`
	RoomName        string    `json: "room_name"`
	RoomNo          int       `json:"room_no"`
	RoomDescription string    `json: "room_description"`
	RoomFacility    string    `json: ""room_facility"`
	RoomAmenity     string    `json:"room_amenity"`
	RoomGallery     string    `json: "room_gallery`
	RoomVideo       string    `json:"room_video"`
	CreatedAt       time.Time `json: "created_at"`
	UpdatedAt       time.Time `json: "updated_at"`
}

// Reservations is the reservation model
// Room resevetion
type Reservation struct {
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

// RoomStatus is list of status model
// Available, Reserved, Stay, Checkout, Cleanup, Maintenance, Closed
type RoomStatus struct {
	ID         int       `json: "id"`
	StatusIcon string    `json: "status_icon"`
	StatusName string    `json: "status_name"`
	StatusDesc string    `json: "status_desc"`
	CreatedAt  time.Time `json: "created_at"`
	UpdatedAt  time.Time `json: "updated_at"`
}

// RoomAlotmentStatus room reservarion restrictions model
type RoomAlotmentStatus struct {
	ID            int         `json: "id"`
	StatusNote    string      `json: "status_note"`
	RoomID        int         `json: "room_id"`
	Room          Room        `json: "room"`
	RoomStatusID  int         `json: "restriction_id"`
	RoomStatus    RoomStatus  `json: "restriction"`
	ReservationID int         `json: "reservation_id"`
	Reservation   Reservation `json: "reservation"`
	StartDate     time.Time   `json: "start_date"`
	EndDate       time.Time   `json: "end_date"`
	CreatedAt     time.Time   `json: "created_at"`
	UpdatedAt     time.Time   `json: "updated_at"`
}

type RoomGallery struct {
}

type RoomFacility struct {
}

type RoomAmenity struct {
}

type RoomCategory struct {
}
