package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/psinthorn/go_smallsite/internal/forms"
	"github.com/psinthorn/go_smallsite/internal/helpers"
	"github.com/psinthorn/go_smallsite/internal/render"
	"github.com/psinthorn/go_smallsite/models"
	"github.com/psinthorn/go_smallsite/models/reservations"
)

// CheckAvailability is check-availability page render
func (rp *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Template(w, r, "search-availability.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// PostSearchAlotment is check-availability page render
func (rp *Repository) PostSearchAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("Start Date: %s and End date is: %s", start, end)))
}

type jsonReponse struct {
	OK      bool   `json: "ok"`
	Message string `json: "message"`
}

// AvailabilityResponse is availability response in json
func (rp *Repository) AvailabilityResponse(w http.ResponseWriter, r *http.Request) {
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
func (rp *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation reservations.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.Template(w, r, "make-reservation.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation is reservation page render
func (rp *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	reservation := reservations.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
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

		render.Template(w, r, "make-reservation.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	rp.App.Session.Put(r.Context(), "reservation", reservation)
	rp.App.Session.Put(r.Context(), "success", "Thank you, Please re-check your information for next process :)")
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)

}

// ReservationSummary for customer recheck information before submit
func (rp *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
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
	render.Template(w, r, "reservation-summary.page.html", &models.TemplateData{
		Data: data,
	})
}
