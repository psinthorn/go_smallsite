package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/psinthorn/go_smallsite/configs"
	"github.com/psinthorn/go_smallsite/domain/reservations"
	"github.com/psinthorn/go_smallsite/domain/templates"
	"github.com/psinthorn/go_smallsite/internal/forms"
	"github.com/psinthorn/go_smallsite/internal/helpers"
	"github.com/psinthorn/go_smallsite/internal/render"
)

var (
	ReservationsController reservationsControllerInterface = &reservationsController{}
)

type reservationsControllerInterface interface {
	SearchAvailability(w http.ResponseWriter, r *http.Request)
	PostSearchAvailability(w http.ResponseWriter, r *http.Request)
	AvailabilityResponse(w http.ResponseWriter, r *http.Request)
	Reservation(w http.ResponseWriter, r *http.Request)
	PostReservation(w http.ResponseWriter, r *http.Request)
	ReservationSummary(w http.ResponseWriter, r *http.Request)
}
type reservationsController struct {
	App *configs.AppConfig
	DB  interface{}
}

// CheckAvailability is check-availability page render
func (rp *reservationsController) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Template(w, r, "search-availability.page.html", &templates.TemplateData{
		StringMap: stringMap,
	})

}

// PostSearchAlotment is check-availability page render
func (rp *reservationsController) PostSearchAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("Start Date: %s and End date is: %s", start, end)))
}

type jsonReponse struct {
	OK      bool   `json: "ok"`
	Message string `json: "message"`
}

// AvailabilityResponse is availability response in json
func (rp *reservationsController) AvailabilityResponse(w http.ResponseWriter, r *http.Request) {
	resp := jsonReponse{
		OK:      true,
		Message: "Hello Json",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

// Reservation is reservation page render
func (rp *reservationsController) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation reservations.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.Template(w, r, "make-reservation.page.html", &templates.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation is reservation page render
func (rp *reservationsController) PostReservation(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	sd := r.Form.Get("start_date")
	ed := r.Form.Get("end_date")

	dateLayout := "2006-01-02"
	startDate, err := time.Parse(dateLayout, sd)
	if err != nil {
		panic(err)
		// helpers.ServerError(w, err)
	}
	endDate, err := time.Parse(dateLayout, ed)
	if err != nil {
		panic(err)
		// helpers.ServerError(w, err)
	}

	roomID, err := strconv.Atoi(r.Form.Get("room_id"))
	if err != nil {
		panic(err)
		// helpers.ServerError(w, err)
	}

	reservation := reservations.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
		StartDate: startDate,
		EndDate:   endDate,
		RoomID:    roomID,
	}

	form := forms.New(r.PostForm)

	// form.Has("first_name", r)
	form.Required("first_name", "last_name", "email")
	// minimum require on input field
	form.MinLength("first_name", 3, r)
	// email validation
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.Template(w, r, "make-reservation.page.html", &templates.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	// _, err = rp.DB.InsertReservation(reservation)
	// if err != nil {
	// 	helpers.ServerError(w, err)
	// }

	rp.App.Session.Put(r.Context(), "reservation", reservation)
	rp.App.Session.Put(r.Context(), "success", "Thank you, Please re-check your information for next process :)")
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)

}

// ReservationSummary for customer recheck information before submit
func (rp *reservationsController) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := rp.App.Session.Get(r.Context(), "reservation").(reservations.Reservation)
	if !ok {
		rp.App.ErrorLog.Println("can't get reservation information from session")
		rp.App.Session.Put(r.Context(), "error", "can't get reservation information from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	rp.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation
	render.Template(w, r, "reservation-summary.page.html", &templates.TemplateData{
		Data: data,
	})
}