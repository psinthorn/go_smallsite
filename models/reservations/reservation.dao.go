package reservations

import (
	"context"
	"time"
)

const (
	InsertReservation  = `insert into reservations (first_name,last_name,email,phone,start_date,end_date,room_id,created_at,updated_at,	) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
	GetAllReservation  = ``
	GetReservationById = ``
	UpdateReservation  = ``
	DeleteReservation  = ``
)

func (m *SQLDBRepo) InsertReservation(rsvn reservations.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var rsvnID int
	stmt := InsertReservation

	_, err := m.DB.ExecContext(ctx, stmt,
		rsvn.FirstName,
		rsvn.LastName,
		rsvn.Email,
		rsvn.Phone,
		rsvn.StartDate,
		rsvn.EndDate,
		rsvn.RoomID,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return 0, err
	}

	return rsvnID, nil
}
