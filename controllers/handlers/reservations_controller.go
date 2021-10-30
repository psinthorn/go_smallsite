// routes -> controllers -> services -> domain

package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	domain_mail "github.com/psinthorn/go_smallsite/domain/mail"
	domain_reservation "github.com/psinthorn/go_smallsite/domain/reservations"
	domain "github.com/psinthorn/go_smallsite/domain/rooms"
	"github.com/psinthorn/go_smallsite/domain/templates"
	"github.com/psinthorn/go_smallsite/internal/forms"
	"github.com/psinthorn/go_smallsite/internal/helpers"
	"github.com/psinthorn/go_smallsite/internal/render"
	"github.com/psinthorn/go_smallsite/internal/utils"
)

// CheckAvailability is check-availability page render
func (rp *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Template(w, r, "search-availability.page.html", &templates.TemplateData{
		StringMap: stringMap,
	})

}

// PostSearchAlailability is check-availability page render
func (rp *Repository) PostSearchAvailability(w http.ResponseWriter, r *http.Request) {
	startDate, err := utils.UtilsService.StringToTime(r.Form.Get("start_date"))
	endDate, err := utils.UtilsService.StringToTime(r.Form.Get("end_date"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	var rooms []domain.Room
	rooms, err = domain_reservation.ReservationService.SearchAvailabilityAllRoom(startDate, endDate)
	if err != nil {
		helpers.ServerError(w, err)
	}
	if len(rooms) == 0 {
		rp.App.Session.Put(r.Context(), "error", "Sorry no rooms available on this period.")
		http.Redirect(w, r, "/rooms/search-availability", http.StatusSeeOther)
		return
	}

	rsvn := domain_reservation.Reservation{
		StartDate:       startDate,
		EndDate:         endDate,
		PromotionId:     999,
		PromotionTypeId: 999,
	}

	data := make(map[string]interface{})
	data["rooms"] = rooms

	rp.App.Session.Put(r.Context(), "reservation", rsvn)
	render.Template(w, r, "choose-room.page.html", &templates.TemplateData{
		Data: data,
	})

}

// ChooseRoom choose room for reservation
func (rp *Repository) ChooseRoom(w http.ResponseWriter, r *http.Request) {
	roomID, err := strconv.Atoi(chi.URLParam(r, "id"))
	roomTypeId, err := strconv.Atoi(chi.URLParam(r, "type"))
	roomNo, err := strconv.Atoi(chi.URLParam(r, "no"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	rsvn, ok := rp.App.Session.Get(r.Context(), "reservation").(domain_reservation.Reservation)
	if !ok {
		helpers.ServerError(w, err)
		return
	}
	fmt.Println("choose room")
	fmt.Println(rsvn.IsPromotion)
	rsvn.RoomID = roomID
	rsvn.Room.RoomTypeId = roomTypeId
	rsvn.Room.RoomNo = roomNo

	rp.App.Session.Put(r.Context(), "reservation", rsvn)
	http.Redirect(w, r, "/rooms/reservation", http.StatusSeeOther)
}

// Reservation is reservation page that will render reservation form for cutomer fill-in information
func (rp *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	rsvn, ok := rp.App.Session.Get(r.Context(), "reservation").(domain_reservation.Reservation)
	if !ok {
		helpers.ServerError(w, errors.New("can't get reservation information"))
		return
	}

	fmt.Println(rsvn)

	// fmt.Println("room type ID: ", rsvn.Room.RoomTypeId)
	roomTypeID := rsvn.Room.RoomTypeId
	roomType, err := domain.RoomTypeService.GetByID(roomTypeID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	roomTypeTitle := roomType.Title

	// convert time.Time to string for display in form
	sd := rsvn.StartDate.Format("2006-01-02")
	ed := rsvn.EndDate.Format("2006-01-02")

	stringMap := make(map[string]string)
	stringMap["start_date"] = sd
	stringMap["end_date"] = ed
	stringMap["room_type"] = roomTypeTitle
	stringMap["room_type_id"] = strconv.Itoa(roomTypeID)

	data := make(map[string]interface{})
	data["reservation"] = rsvn

	render.Template(w, r, "make-reservation.page.html", &templates.TemplateData{
		Form:      forms.New(nil),
		Data:      data,
		StringMap: stringMap,
	})
}

// ReservationByRoomType is search and reservation a room by type that render on individual page
func (rp *Repository) ReservationByRoomType(w http.ResponseWriter, r *http.Request) {
	sd := r.URL.Query().Get("sd")
	ed := r.URL.Query().Get("ed")
	startDate, err := utils.UtilsService.StringToTime(sd)
	endDate, err := utils.UtilsService.StringToTime(ed)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	roomTypeID := r.URL.Query().Get("type")
	roomTypeId, _ := strconv.Atoi(roomTypeID)
	roomType, err := domain.RoomTypeService.GetByID(roomTypeId)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	roomTypeTitle := roomType.Title
	stringMap := make(map[string]string)
	// stringMap["start_date"] = sd
	// stringMap["end_date"] = ed
	stringMap["room_type"] = roomTypeTitle
	stringMap["room_type_id"] = strconv.Itoa(roomTypeId)

	data := make(map[string]interface{})
	rsvn := domain_reservation.Reservation{
		RoomID:          999,
		RoomNo:          999,
		PromotionId:     999,
		PromotionTypeId: 999,
		StartDate:       startDate,
		EndDate:         endDate,
	}
	rsvn.Room.RoomTypeId = 999
	data["reservation"] = rsvn
	rp.App.Session.Put(r.Context(), "reservation", rsvn)
	http.Redirect(w, r, "/rooms/reservation", http.StatusSeeOther)

}

// PostReservation id create and save reservation information to reservaton and allotment table on database
func (rp *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	form := forms.New(r.PostForm)
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	sd := r.Form.Get("start_date")
	ed := r.Form.Get("end_date")

	roomTypeTitle := r.Form.Get("room_type")
	roomTypeId := r.Form.Get("room_type_id")
	roomTypeIdInt, err := strconv.Atoi(roomTypeId)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	roomId := r.Form.Get("room_id")
	roomIdInt, err := strconv.Atoi(roomId)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// form.Has("first_name", r)
	form.Required("first_name", "last_name", "email", "start_date", "end_date")
	// minimum require on input field
	form.MinLength("first_name", 3, r)
	// email validation
	form.IsEmail("email")

	// convert from string date to time.Time format
	startDate, err := utils.UtilsService.StringToTime(sd)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	endDate, err := utils.UtilsService.StringToTime(ed)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	var room domain.Room
	if roomIdInt != 0 {
		room, err = domain.RoomService.GetRoomByID(roomIdInt)
		if err != nil {
			helpers.ServerError(w, err)
			return
		}
	}
	//isPromotion, _ := strconv.ParseBool(r.Form.Get("Is_promotion"))
	rsvn, ok := rp.App.Session.Get(r.Context(), "reservation").(domain_reservation.Reservation)
	if !ok {
		helpers.ServerError(w, errors.New("can't get reservation from session"))
		return
	}

	fmt.Println(rsvn)

	amount, _ := strconv.Atoi(r.Form.Get("amount"))
	perNight := rsvn.Amount / 4
	reservation := domain_reservation.Reservation{
		FirstName:       r.Form.Get("first_name"),
		LastName:        r.Form.Get("last_name"),
		Email:           r.Form.Get("email"),
		Phone:           r.Form.Get("phone"),
		RoomID:          roomIdInt,
		RoomTypeName:    roomTypeTitle,
		Room:            room,
		PromotionId:     rsvn.PromotionId,
		PromotionTypeId: rsvn.PromotionTypeId,
		IsPromotion:     rsvn.IsPromotion,
		Amount:          float32(amount),
		Status:          "reservation",
		StartDate:       startDate,
		EndDate:         endDate,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if !form.Valid() {
		stringMap := make(map[string]string)
		stringMap["start_date"] = sd
		stringMap["end_date"] = ed
		stringMap["room_type"] = roomTypeTitle
		stringMap["room_type_id"] = roomTypeId
		stringMap["per_night"] = strconv.Itoa(int(perNight))

		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.Template(w, r, "make-reservation.page.html", &templates.TemplateData{
			Form:      form,
			Data:      data,
			StringMap: stringMap,
		})
		return
	}

	rsvnID, err := domain_reservation.ReservationService.Create(reservation)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	rsvnAllmentStatus := domain.RoomAllotmentStatus{
		RoomTypeID:    roomTypeIdInt,
		RoomNoID:      roomIdInt,
		ReservationID: rsvnID,
		RoomStatusID:  2,
		StartDate:     startDate,
		EndDate:       endDate,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	_, err = domain.RoomAllotmentStatusService.Creat(rsvnAllmentStatus)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	var subject string
	if reservation.IsPromotion {
		subject = "Promotion Reservation information"
	} else {
		subject = "Room reservation information"
	}

	// Send mail confirmation to guest
	mailToGuest := fmt.Sprintf(`
		<strong>Reservation Information</strong><br/><br/>
		Dear %s %s<br/>
		Thank you for your reservationwith us <br/> 
		the below is your reservation conformation <br/>
		<br/>
		<strong>Room type:</strong> %s <br/>
		<strong>Arrival:</strong> %s <br/>
		<strong>Departure:</strong> %s <br/>
		 <br/>
		 <strong>Promotion type:</strong> %s <br/>
`, reservation.FirstName, reservation.LastName, roomTypeTitle,
		reservation.StartDate.Format("2006-02-01"),
		reservation.EndDate.Format("2006-02-01"),
	)

	mailMsg := domain_mail.MailDataTemplate{
		To:       reservation.Email,
		From:     "rsvn@gosmallsitehotel.com",
		Subject:  subject,
		Content:  mailToGuest,
		Template: "drip.html",
	}
	//send message to mail chanle
	rp.App.MailChan <- mailMsg

	// Send mail notification to reservation
	mailToRsvn := fmt.Sprintf(`
	<strong>Reservation Information</strong><br/>
	<br/>
		
		the below is your reservation information <br/>
		<br/>
		<strong>Customer Name:</strong> %s %s <br/>
		<strong>Email:</strong> %s <br/>
		<strong>Phone:</strong> %s <br/>
		<br/>
		<strong>Room type:</strong> %s <br/>
		<strong>Arrival:</strong> %s <br/>
		<strong>Departure:</strong> %s <br/>
		 <br/>
		 <strong>Promotion type:</strong> %s <br/>

`, reservation.FirstName,
		reservation.LastName,
		reservation.Email,
		reservation.Phone,
		roomTypeTitle,
		reservation.StartDate.Format("2006-02-01"),
		reservation.EndDate.Format("2006-02-01"),
	)

	mailNotificationRsvn := domain_mail.MailDataTemplate{
		To:       "rsvn@gosmallsitehotel.com",
		From:     "rsvn@gosmallsitehotel.com",
		Subject:  subject,
		Content:  mailToRsvn,
		Template: "",
	}
	//send message to mail chanle
	rp.App.MailChan <- mailNotificationRsvn

	rp.App.Session.Put(r.Context(), "reservation", reservation)
	rp.App.Session.Put(r.Context(), "success", "Thank you, Please check your email for confirmation :)")
	http.Redirect(w, r, "/rooms/reservation-summary", http.StatusSeeOther)

}

// ReservationSummary is summary page that display reservation information for customer review
func (rp *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := rp.App.Session.Get(r.Context(), "reservation").(domain_reservation.Reservation)
	if !ok {
		rp.App.ErrorLog.Println("can't get reservation information from session")
		rp.App.Session.Put(r.Context(), "error", "can't get reservation information from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	sd := reservation.StartDate.Format("2006-01-02")
	ed := reservation.EndDate.Format("2006-01-02")
	stringMap := make(map[string]string)
	stringMap["start_date"] = sd
	stringMap["end_date"] = ed

	rp.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation
	render.Template(w, r, "reservation-summary.page.html", &templates.TemplateData{
		Data:      data,
		StringMap: stringMap,
	})
}

// API section

type JsonResponse struct {
	Ok        bool   `json: "ok"`
	RoomID    string `json: "room_id"`
	StartDate string `json: "start_date"`
	EndDate   string `json: "end_date"`
	Message   string `json: "message"`
}

func (rp *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {

	startDate, err := utils.UtilsService.StringToTime(r.Form.Get("start_date"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	endDate, err := utils.UtilsService.StringToTime(r.Form.Get("end_date"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	roomID, err := strconv.Atoi(r.Form.Get("room_id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	available, err := domain_reservation.ReservationService.SearchAvailabilityByRoomId(roomID, startDate, endDate)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	sd := r.Form.Get("start_date")
	ed := r.Form.Get("end_date")
	res := JsonResponse{
		Ok:        available,
		RoomID:    strconv.Itoa(roomID),
		StartDate: sd,
		EndDate:   ed,
		Message:   fmt.Sprintf("Room is available"),
	}
	out, err := json.MarshalIndent(res, "", "     ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
