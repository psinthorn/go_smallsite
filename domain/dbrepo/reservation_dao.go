package dbrepo

import (
	"context"
	"time"

	"github.com/psinthorn/go_smallsite/datasources/drivers"
	"github.com/psinthorn/go_smallsite/domain/rooms"
)

const (
	queryInsertReservation         = "insert into reservations (first_name, last_name, email, phone, room_id, status, start_date, end_date, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) returning id"
	querySelectAllRsvn             = "SELECT * FROM reservations"
	querySearchAvailability        = "SELECT count(id) FROM room_allotments WHERE room_no_id = $1 AND $2 < end_date AND $3 > start_date"
	querySearchAvailabilityAllRoom = `SELECT r.id, r.roomtype_id, r.room_no FROM rooms r WHERE r.id not in (SELECT ra.room_no_id FROM room_allotments ra WHERE $1 < ra.end_date AND $2 > ra.start_date)`
)

var ReservationService reservationDomainInterface = &Reservation{}

type Reservation reservation
type reservationDomainInterface interface {
	Create(Reservation) (int, error)
	GetAll()
	GetByID()
	Update()
	Delete()

	SearchAvailabilityByRoomId(roomID int, start, end time.Time) (bool, error)
	SearchAvailabilityAllRoom(start, end time.Time) ([]rooms.Room, error)
}

// PostReservation is reservation page render
func (r *Reservation) Create(rsvn Reservation) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		panic(err)
	}
	var newReservationId int
	err = dbConn.SQL.QueryRowContext(ctx, queryInsertReservation, rsvn.FirstName, rsvn.LastName, rsvn.Email, rsvn.Phone, rsvn.RoomID, rsvn.Status, rsvn.StartDate, rsvn.EndDate, rsvn.CreatedAt, rsvn.UpdatedAt).Scan(&newReservationId)
	if err != nil {
		return 0, err
	}
	defer dbConn.SQL.Close()

	return newReservationId, nil

}

// ReservationSummary for customer recheck information before submit
func (r *Reservation) GetAll() {

}

func (r *Reservation) GetByID() {}
func (r *Reservation) Update()  {}
func (r *Reservation) Delete()  {}

// CheckAvailabilityByRoomId return available room of give room id with given dates
func (r *Reservation) SearchAvailabilityByRoomId(roomID int, start, end time.Time) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return false, err
	}

	var numRows int
	row := dbConn.SQL.QueryRowContext(ctx, querySearchAvailability, roomID, start, end)
	err = row.Scan(&numRows)
	if err != nil {
		return false, err
	}

	if numRows == 0 {
		return true, nil
	}

	defer dbConn.SQL.Close()
	return false, nil

}

// SearchCheckAvailabilityAllRoom return a slice of available rooms for given date range
func (r *Reservation) SearchAvailabilityAllRoom(startDate, endDate time.Time) ([]rooms.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		return nil, err
	}

	var availableRooms []rooms.Room
	rows, err := dbConn.SQL.QueryContext(ctx, querySearchAvailabilityAllRoom, startDate, endDate)

	for rows.Next() {
		var room rooms.Room
		err := rows.Scan(
			&room.ID,
			&room.RoomTypeId,
			&room.RoomNo,
		)
		if err != nil {
			return nil, err
		}
		availableRooms = append(availableRooms, room)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	defer dbConn.SQL.Close()
	return availableRooms, nil

}
