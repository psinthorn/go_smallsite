package controllers

import (
	"net/http"

	"github.com/psinthorn/go_smallsite/configs"
	"github.com/psinthorn/go_smallsite/domain/templates"
	"github.com/psinthorn/go_smallsite/internal/render"
)

var (
	PagesController pagesControllerInterface = &pagesController{}
)

type pagesControllerInterface interface {
	Home(w http.ResponseWriter, r *http.Request)
	About(w http.ResponseWriter, r *http.Request)
	Rooms(w http.ResponseWriter, r *http.Request)
	Superior(w http.ResponseWriter, r *http.Request)
	Deluxe(w http.ResponseWriter, r *http.Request)
	Contact(w http.ResponseWriter, r *http.Request)
}

type pagesController struct {
	App *configs.AppConfig
	DB  interface{}
}

// Home is home page render
func (rp *pagesController) Home(w http.ResponseWriter, r *http.Request) {
	// users, err := rp.DB.GetAllUsers()
	// if err != nil {
	// 	fmt.Println("no users found")
	// }

	// fmt.Println(users)

	// remoteIP := r.RemoteAddr
	// fmt.Println(remoteIP)
	// rp.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	// stringMap := make(map[string]string)
	// stringMap["greet"] = "Hello Go"

	render.Template(w, r, "home.page.html", &templates.TemplateData{
		// StringMap: stringMap,
	})
}

// About is about page render
func (rp *pagesController) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Template(w, r, "about.page.html", &templates.TemplateData{
		StringMap: stringMap,
	})

}

// Room is room page render
func (rp *pagesController) Rooms(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Template(w, r, "rooms.page.html", &templates.TemplateData{
		StringMap: stringMap,
	})

}

// Superior is room page render
func (rp *pagesController) Superior(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Template(w, r, "room-superior.page.html", &templates.TemplateData{
		StringMap: stringMap,
	})

}

// Superior is room page render
func (rp *pagesController) Deluxe(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Template(w, r, "room-deluxe.page.html", &templates.TemplateData{
		StringMap: stringMap,
	})

}

// Contact is contact page render
func (rp *pagesController) Contact(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := rp.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Template(w, r, "contact.page.html", &templates.TemplateData{
		StringMap: stringMap,
	})

}
