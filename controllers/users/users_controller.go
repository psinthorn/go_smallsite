package controllers

import (
	"net/http"

	"github.com/psinthorn/go_smallsite/domain/templates"
	"github.com/psinthorn/go_smallsite/internal/forms"
	"github.com/psinthorn/go_smallsite/internal/render"
)

var (
	UsersController usersControllerInterface = &usersController{}
)

type usersControllerInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
	GetAllUsers()
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type usersController struct{}

// Register new user
func (rp *Repository) Register(w http.ResponseWriter, r *http.Request) {

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
