// routes -> controllers -> services -> domain

package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/psinthorn/go_smallsite/domain/dbrepo"
	"github.com/psinthorn/go_smallsite/domain/rooms"
	"github.com/psinthorn/go_smallsite/domain/templates"
	"github.com/psinthorn/go_smallsite/internal/forms"
	"github.com/psinthorn/go_smallsite/internal/helpers"
	"github.com/psinthorn/go_smallsite/internal/render"
	"github.com/psinthorn/go_smallsite/internal/utils"
)

// var (
// 	ReservationsController reservationsController
// )

// type reservationsControllerInterface interface {
// 	SearchAvailability(w http.ResponseWriter, r *http.Request)
// 	PostSearchAvailability(w http.ResponseWriter, r *http.Request)
// 	AvailabilityResponse(w http.ResponseWriter, r *http.Request)
// 	Reservation(w http.ResponseWriter, r *http.Request)
// 	PostReservation(w http.ResponseWriter, r *http.Request)
// 	ReservationSummary(w http.ResponseWriter, r *http.Request)
// }

// type reservationsController struct {
// 	// App *configs.AppConfig
// 	// DB  interface{}
// }

// CheckAvailability is check-availability page render
func (rp *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Template(w, r, "search-availability.page.html", &templates.TemplateData{
		StringMap: stringMap,
	})

}

// PostSearchAlotment is check-availability page render
func (rp *Repository) PostSearchAvailability(w http.ResponseWriter, r *http.Request) {

	startDate, err := utils.UtilsService.StringToTime(r.Form.Get("start_date"))
	endDate, err := utils.UtilsService.StringToTime(r.Form.Get("end_date"))
	fmt.Println(startDate, endDate)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	rooms, err := dbrepo.ReservationService.SearchAvailabilityAllRoom(startDate, endDate)
	if err != nil {
		helpers.ServerError(w, err)
	}
	if len(rooms) == 0 {
		rp.App.Session.Put(r.Context(), "error", "Sorry no rooms available on this period.")
		http.Redirect(w, r, "/rooms/search-availability", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["rooms"] = rooms

	rsvn := dbrepo.Reservation{
		StartDate: startDate,
		EndDate:   endDate,
	}

	rp.App.Session.Put(r.Context(), "reservation", rsvn)
	render.Template(w, r, "choose-room.page.html", &templates.TemplateData{
		Data: data,
	})

}

// ChooseRoom choose room for reservation
func (rp *Repository) ChooseRoom(w http.ResponseWriter, r *http.Request) {
	roomID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	res, ok := rp.App.Session.Get(r.Context(), "reservation").(dbrepo.Reservation)
	if !ok {
		helpers.ServerError(w, err)
		return
	}

	res.RoomID = roomID
	rp.App.Session.Put(r.Context(), "reservation", res)
	http.Redirect(w, r, "/rooms/reservation", http.StatusSeeOther)
}

// AvailabilityResponse is availability response in json
func (rp *Repository) AvailabilityResponse(w http.ResponseWriter, r *http.Request) {

}

// Reservation is reservation page render
func (rp *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	rsvn, ok := rp.App.Session.Get(r.Context(), "reservation").(dbrepo.Reservation)
	if !ok {
		helpers.ServerError(w, errors.New("can't get reservation information"))
		return
	}

	// convert time.Time to string for display in form
	sd := rsvn.StartDate.Format("2006-01-02")
	ed := rsvn.EndDate.Format("2006-01-02")

	stringMap := make(map[string]string)
	stringMap["start_date"] = sd
	stringMap["end_date"] = ed

	data := make(map[string]interface{})
	data["reservation"] = rsvn

	render.Template(w, r, "make-reservation.page.html", &templates.TemplateData{
		Form:      forms.New(nil),
		Data:      data,
		StringMap: stringMap,
	})
}

// PostReservation is reservation page render
func (rp *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	sd := r.Form.Get("start_date")
	ed := r.Form.Get("end_date")
	fmt.Println("Start Date: ", sd)
	fmt.Println("End Date: ", ed)

	// convert from string date to time.Time format
	startDate, err := utils.UtilsService.StringToTime(sd)
	endDate, err := utils.UtilsService.StringToTime(ed)

	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	roomID, err := strconv.Atoi(r.Form.Get("room_id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	roomType, err := rooms.RoomTypeService.GetRoomTypeByID()
	if err != nil {

	}

	reservation := dbrepo.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
		RoomID:    roomID,
		Room:      rooms.Room{},
		Status:    "stay",
		StartDate: startDate,
		EndDate:   endDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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

	rsvnID, err := dbrepo.ReservationService.Create(reservation)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	rsvnAllmentStatus := rooms.RoomAllotmentStatus{
		RoomTypeID:    3,
		RoomNoID:      3,
		ReservationID: rsvnID,
		RoomStatusID:  2,
		StartDate:     startDate,
		EndDate:       endDate,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	_, err = rooms.RoomAllotmentStatusService.Creat(rsvnAllmentStatus)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	rp.App.Session.Put(r.Context(), "reservation", reservation)
	rp.App.Session.Put(r.Context(), "success", "Thank you, Please re-check your information for next process :)")
	http.Redirect(w, r, "/rooms/reservation-summary", http.StatusSeeOther)

}

// ReservationSummary for customer recheck information before submit
func (rp *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
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
