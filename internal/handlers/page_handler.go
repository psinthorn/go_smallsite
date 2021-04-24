package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/psinthorn/go_smallsite/internal/configs"
	"github.com/psinthorn/go_smallsite/internal/models"
	"github.com/psinthorn/go_smallsite/internal/renders"
)

// Repo
var Repo *Repository

type Repository struct {
	App *configs.AppConfig
}

// NewRepository
func NewRepository(a *configs.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is home page render
func (rp *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	fmt.Println(remoteIP)
	rp.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	stringMap := make(map[string]string)
	stringMap["greet"] = "Hello Go"

	renders.RenderTemplate(w, r, "home.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About is about page render
func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Room is room page render
func (rp *Repository) Rooms(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, r, "rooms.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Superior is room page render
func (rp *Repository) Superior(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, r, "room-superior.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Superior is room page render
func (rp *Repository) Deluxe(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, r, "room-deluxe.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// // CheckAvailability is check-availability page render
// func (rp *Repository) CheckAlotment(w http.ResponseWriter, r *http.Request) {
// 	stringMap := make(map[string]string)
// 	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
// 	stringMap["remote_ip"] = remoteIP

// 	renders.RenderTemplate(w, "check-alotment.page.html", &models.TemplateData{
// 		StringMap: stringMap,
// 	})

// }

// CheckAvailability is check-availability page render
func (rp *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{
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

	out, _ := json.MarshalIndent(resp, "", "     ")

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

// Reservation is reservation page render
func (rp *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Reservation is reservation page render
func (rp *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}
