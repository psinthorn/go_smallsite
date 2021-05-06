package dbrepo

var ReservationInterace reservationDomainInterface = &Reservation{}

type reservationDomainInterface interface {
	Create()
	GetAll()
	GetByID()
	Update()
	Delete()

	SearchAvailability()
	PostSearchAvailability(string, string) (string, error)
	AvailabilityResponse()
}

// PostReservation is reservation page render
func (r *Reservation) Create() {

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

func (r *Reservation) GetByID() {

}

func (r *Reservation) Update() {}

func (r *Reservation) Delete() {}

// CheckAvailability is check-availability page render
func (r *Reservation) SearchAvailability() {

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
