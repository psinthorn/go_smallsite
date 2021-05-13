package dbrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/psinthorn/go_smallsite/datasources/drivers"
)

const (
	queryInsertReservation  = "insert into reservations (first_name, last_name, email, phone, room_id, status, start_date, end_date, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) returning id"
	querySelectAllRsvn      = "SELECT * FROM reservations"
	querySearchAvailability = "SELECT count(id) FROM room_allotments WHERE room_no_id = $1 AND $2 < end_date AND $3 > start_date"
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
	PostSearchAvailability(sd string, ed string) (string, error)
	AvailabilityResponse()
}

// PostReservation is reservation page render
func (r *Reservation) Create(rsvn Reservation) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	fmt.Println(rsvn.StartDate)
	fmt.Println(rsvn.EndDate)
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
	// reservation, ok := rp.App.Session.Get(r.Context(), "reservation").(reservations.Reservation)
	// if !ok {
	// 	rp.App.ErrorLog.Println("can't get reservation information from session")
	// 	rp.App.Session.Put(r.Context(), "error", "can't get reservation information from session")
	// 	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	// 	return
	// }

	// rp.App.Session.Remove(r.Context(), "reservation")

	// data := make(map[string]interface{})
	// data["reservation"] = reservation
	// render.Template(w, r, "reservation-summary.page.html", &templates.TemplateData{
	// 	Data: data,
	// })
}

func (r *Reservation) GetByID() {}
func (r *Reservation) Update()  {}
func (r *Reservation) Delete()  {}

// CheckAvailability is check-availability page render
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

// PostSearchAlotment is check-availability page render
func (r *Reservation) PostSearchAvailability(sd string, ed string) (string, error) {
	return "", nil
}

// // AvailabilityResponse is availability response in json
func (r *Reservation) AvailabilityResponse() {
	// 	resp := jsonReponse{
	// 		OK:      true,
	// 		Message: "Hello Json",
	// 	}

	// 	out, err := json.MarshalIndent(resp, "", "     ")
	// 	if err != nil {
	// 		helpers.ServerError(w, err)
	// 		return
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(out)

}

// Reservation is reservation page render
// func Reservation() {
// 	var emptyReservation Reservation
// 	data := make(map[string]interface{})
// 	data["reservation"] = emptyReservation

// }
