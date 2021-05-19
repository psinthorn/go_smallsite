// routes -> controllers -> domain(models)

package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/psinthorn/go_smallsite/domain/templates"
	"github.com/psinthorn/go_smallsite/domain/users"
	"github.com/psinthorn/go_smallsite/internal/forms"
	"github.com/psinthorn/go_smallsite/internal/helpers"
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

// AddNewUserForm
func (rp *Repository) AddNewUserForm(w http.ResponseWriter, r *http.Request) {
	var emptyUser users.User
	data := make(map[string]interface{})
	data["user"] = emptyUser

	render.Template(w, r, "admin-user-add-form.page.html", &templates.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})

}

//  AddNewUser
func (rp *Repository) AddNewUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	accessLevel, err := strconv.Atoi(r.Form.Get("access_level"))
	if err != nil {
		panic(err)
	}
	newUser := users.User{
		FirstName:   r.Form.Get("first_name"),
		LastName:    r.Form.Get("last_name"),
		Email:       r.Form.Get("email"),
		Password:    r.Form.Get("password"),
		AccessLevel: accessLevel,
		Status:      r.Form.Get("status"),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Form validation
	form := forms.New(r.PostForm)
	if !form.Valid() {
		data := make(map[string]interface{})
		data["user"] = newUser
		render.Template(w, r, "admin-user-add-form.page.html", &templates.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	_, err = users.UserService.Create(newUser)
	if err != nil {
		log.Fatal(err)
	}

	rp.App.Session.Put(r.Context(), "user", newUser)
	rp.App.Session.Put(r.Context(), "success", "New user is added :)")
	http.Redirect(w, r, "/admin/users/register", http.StatusSeeOther)
}

//  GetAllUsers
func (rp *Repository) GetAllUsers(w http.ResponseWriter, r *http.Request) {

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
