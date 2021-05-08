package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/psinthorn/go_smallsite/domain/dbrepo"
	"github.com/psinthorn/go_smallsite/domain/templates"
	"github.com/psinthorn/go_smallsite/internal/forms"
	"github.com/psinthorn/go_smallsite/internal/helpers"
	"github.com/psinthorn/go_smallsite/internal/render"
)

//Room
func (rp *Repository) Room() error {
	return nil
}

// PostReservation is reservation page render
func (rp *Repository) CreatRoom(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Hey room create")
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	room := dbrepo.Room{
		RoomTypeId: 1,
		RoomName:   "101",
		RoomNo:     "101",
		Desc:       "superior room",
		Status:     "available",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		// CreatedAt:  time.Now(),
		// UpdatedAt:  time.Now(),
		// RoomTypeId: r.Form.Get("room_type_id"),
		// RoomName:   r.Form.Get("room_name"),
		// RoomNo:     r.Form.Get("room_no"),
		// Desc:       r.Form.Get("desc"),
		// Status:     r.Form.Get("status"),
		// CreatedAt:  time.Now(),
		// UpdatedAt:  time.Now(),
	}

	form := forms.New(r.PostForm)

	// // form.Has("first_name", r)
	// form.Required("first_name", "last_name", "email")
	// // minimum require on input field
	// form.MinLength("first_name", 3, r)
	// // email validation
	// form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["room"] = room
		render.Template(w, r, "room-create.html", &templates.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	_, err = dbrepo.RoomService.Create(room)
	if err != nil {
		panic(err)
	}

	rp.App.Session.Put(r.Context(), "room", room)
	rp.App.Session.Put(r.Context(), "success", "New room is added :)")
	http.Redirect(w, r, "/rooms", http.StatusSeeOther)

}

// ReservationSummary for customer recheck information before submit
func (rp *Repository) RoomSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := rp.App.Session.Get(r.Context(), "reservation").(dbrepo.Reservation)
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
