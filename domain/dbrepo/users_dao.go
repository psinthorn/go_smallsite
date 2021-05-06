package dbrepo

var User userDomainInterface = &user{}

type userDomainInterface interface {
	Register() (string, error)
	GetAllUsers() (string, error)
	GetByID()
	Update()
	Delete()
}

// Register new user
func (u *user) Register() (string, error) {

	return "Implement Me! Register function", nil
}

//  GetAllUsers
func (u *user) GetAllUsers() (string, error) {
	return "Implement Me! To get all users function", nil
}

// GetByID
func (u *user) GetByID() {}

// Login user login page
func (u *user) Update() {}

// Login user login page
func (u *user) Delete() {
	// http.Redirect(w, r, "/login", http.StatusSeeOther)
	//render.Template(w, r, "login.page.html", &models.TemplateData{})
}
