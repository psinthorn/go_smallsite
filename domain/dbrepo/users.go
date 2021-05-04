package dbrepo

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
