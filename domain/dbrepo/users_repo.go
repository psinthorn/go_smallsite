package dbrepo

import "time"

// User is the user model
type User struct {
	ID          int       `json: "id"`
	FirstName   string    `json: "first_name"`
	LastName    string    `json: "last_name"`
	Email       string    `json: "email"`
	Password    string    `json: "password"`
	AccessLevel int       `json: "access_level"`
	CreatedAt   time.Time `json: "created_at"`
	UpdatedAt   time.Time `json: "update_at"`
}

// Register new user
func (rp *sqlDbRepo) Register() {

}

//  GetAllUsers
func (rp *sqlDbRepo) GetAllUsers() {}

// Login user login page
func (rp *sqlDbRepo) Login() {

}

// Login user login page
func (rp *sqlDbRepo) Logout() {
	// http.Redirect(w, r, "/login", http.StatusSeeOther)
	//render.Template(w, r, "login.page.html", &models.TemplateData{})
}
