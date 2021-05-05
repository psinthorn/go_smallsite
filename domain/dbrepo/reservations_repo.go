package dbrepo

import "time"

// Room is the room model
type Room struct {
	ID        int       `json: "id"`
	RoomName  string    `json: "room_name"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
}

// Reservations is the reservation model
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

// Status is the room status  model
type Status struct {
	ID         int       `json: "id"`
	StatusName string    `json: "status_name"`
	CreatedAt  time.Time `json: "created_at"`
	UpdatedAt  time.Time `json: "updated_at"`
}

// RoomRestrictions is the room restriction model
type RoomStatus struct {
	ID            int         `json: "id"`
	RoomID        int         `json: "room_id"`
	ReservationID int         `json: "reservation_id"`
	RestrictionID int         `json: "restriction_id"`
	Room          Room        `json: "room"`
	Reservation   Reservation `json: "reservation"`
	Status        Status      `json: "restriction"`
	StartDate     time.Time   `json: "start_date"`
	EndDate       time.Time   `json: "end_date"`
	CreatedAt     time.Time   `json: "created_at"`
	UpdatedAt     time.Time   `json: "updated_at"`
}

// CheckAvailability is check-availability page render
func (rp *sqlDbRepo) SearchAvailability() {

}

// PostSearchAlotment is check-availability page render
func (rp *sqlDbRepo) PostSearchAvailability() {
	// 	start := r.Form.Get("start")
	// 	end := r.Form.Get("end")
	// 	w.Write([]byte(fmt.Sprintf("Start Date: %s and End date is: %s", start, end)))
	// }

	// type jsonReponse struct {
	// 	OK      bool   `json: "ok"`
	// 	Message string `json: "message"`
}

// // AvailabilityResponse is availability response in json
func (rp *sqlDbRepo) AvailabilityResponse() {
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
func (rp *sqlDbRepo) Reservation() {
	var emptyReservation Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

}

// PostReservation is reservation page render
func (rp *sqlDbRepo) PostReservation() {

	// err := r.ParseForm()
	// if err != nil {
	// 	helpers.ServerError(w, err)
	// 	return
	// }

	// sd := r.Form.Get("start_date")
	// ed := r.Form.Get("end_date")

	// dateLayout := "2006-01-02"
	// startDate, err := time.Parse(dateLayout, sd)
	// if err != nil {
	// 	panic(err)
	// 	// helpers.ServerError(w, err)
	// }
	// endDate, err := time.Parse(dateLayout, ed)
	// if err != nil {
	// 	panic(err)
	// 	// helpers.ServerError(w, err)
	// }

	// roomID, err := strconv.Atoi(r.Form.Get("room_id"))
	// if err != nil {
	// 	panic(err)
	// 	// helpers.ServerError(w, err)
	// }

	// reservation := reservations.Reservation{
	// 	FirstName: r.Form.Get("first_name"),
	// 	LastName:  r.Form.Get("last_name"),
	// 	Email:     r.Form.Get("email"),
	// 	Phone:     r.Form.Get("phone"),
	// 	StartDate: startDate,
	// 	EndDate:   endDate,
	// 	RoomID:    roomID,
	// }

	// form := forms.New(r.PostForm)

	// // form.Has("first_name", r)
	// form.Required("first_name", "last_name", "email")
	// // minimum require on input field
	// form.MinLength("first_name", 3, r)
	// // email validation
	// form.IsEmail("email")

	// if !form.Valid() {
	// 	data := make(map[string]interface{})
	// 	data["reservation"] = reservation
	// 	render.Template(w, r, "make-reservation.page.html", &templates.TemplateData{
	// 		Form: form,
	// 		Data: data,
	// 	})
	// 	return
	// }

	// // _, err = rp.DB.InsertReservation(reservation)
	// // if err != nil {
	// // 	helpers.ServerError(w, err)
	// // }

	// rp.App.Session.Put(r.Context(), "reservation", reservation)
	// rp.App.Session.Put(r.Context(), "success", "Thank you, Please re-check your information for next process :)")
	// http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)

}

// ReservationSummary for customer recheck information before submit
func (rp *sqlDbRepo) ReservationSummary() {
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
