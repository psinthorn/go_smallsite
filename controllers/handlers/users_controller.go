// routes -> controllers -> services -> domain

package controllers

import (
	"log"
	"net/http"

	"github.com/psinthorn/go_smallsite/domain/dbrepo"
	"github.com/psinthorn/go_smallsite/domain/templates"
	"github.com/psinthorn/go_smallsite/internal/forms"
	"github.com/psinthorn/go_smallsite/internal/render"
)

// var (
// 	UsersController usersController
// )

// type usersControllerInterface interface {
// 	Register(w http.ResponseWriter, r *http.Request)
// 	GetAllUsers(w http.ResponseWriter, r *http.Request)
// 	Login(w http.ResponseWriter, r *http.Request)
// 	Logout(w http.ResponseWriter, r *http.Request)
// }

// type usersController struct {
// }

//  GetAllUsers
func (rp *Repository) Register(w http.ResponseWriter, r *http.Request) {
	resp, err := dbrepo.User.Register()
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(resp))
}

//  GetAllUsers
func (rp *Repository) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := dbrepo.User.GetAllUsers()
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte(resp))
}

// Login user login page
func (rp *Repository) Login(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "login.page.html", &templates.TemplateData{
		Form: forms.New(nil),
	})
	rp.App.Session.Put(r.Context(), "success", "Log in success :)")
}

// Login user login page
func (rp *Repository) Logout(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	//render.Template(w, r, "login.page.html", &models.TemplateData{})
}
