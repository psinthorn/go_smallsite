package handlers

import (
	"fmt"
	"net/http"

	"github.com/psinthorn/go_smallsite/internal/models"
	"github.com/psinthorn/go_smallsite/internal/render"
)

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/psinthorn/go_smallsite/internal/configs"
// 	"github.com/psinthorn/go_smallsite/internal/models"
// 	"github.com/psinthorn/go_smallsite/internal/render"
// )

// // Repo
// var Repo *Repository

// type Repository struct {
// 	App *configs.AppConfig
// }

// // NewRepository
// func NewRepository(a *configs.AppConfig) *Repository {
// 	return &Repository{
// 		App: a,
// 	}
// }

// // NewHandlers
// func NewHandlers(r *Repository) {
// 	Repo = r
// }

// Home is home page render
func (rp *Repository) Home(w http.ResponseWriter, r *http.Request) {
	ok := rp.DBConnect.GetAllUsers()
	if !ok {
		fmt.Println("no users found")
	}

	fmt.Sprintf("return from func is: %s", ok)

	remoteIP := r.RemoteAddr
	fmt.Println(remoteIP)
	rp.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	stringMap := make(map[string]string)
	stringMap["greet"] = "Hello Go"

	render.Template(w, r, "home.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// About is about page render
func (rp *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Template(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Room is room page render
func (rp *Repository) Rooms(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Template(w, r, "rooms.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Superior is room page render
func (rp *Repository) Superior(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Template(w, r, "room-superior.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Superior is room page render
func (rp *Repository) Deluxe(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Template(w, r, "room-deluxe.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Contact is contact page render
func (rp *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Template(w, r, "contact.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}
