package handlers

import (
	"fmt"
	"net/http"

	"github.com/psinthorn/go_smallsite/pkg/configs"
	"github.com/psinthorn/go_smallsite/pkg/models"
	"github.com/psinthorn/go_smallsite/pkg/renders"
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

	renders.RenderTemplate(w, "home.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About is about page render
func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Room is room page render
func (rp *Repository) Rooms(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, "rooms.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Superior is room page render
func (rp *Repository) Superior(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, "superior.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Superior is room page render
func (rp *Repository) Deluxe(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, "deluxe.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// CheckAvailability is check-availability page render
func (rp *Repository) CheckAvailability(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, "check-availability.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Reservation is reservation page render
func (rp *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, "reservation.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Reservation is reservation page render
func (rp *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	renders.RenderTemplate(w, "contact.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}
