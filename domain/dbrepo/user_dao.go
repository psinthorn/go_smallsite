package dbrepo

var UserService userDomainInterface = &User{}

type userDomainInterface interface {
	Create() (string, error)
	GetAllUsers() (string, error)
	GetByID()
	Update()
	Delete()
}

// Register new user
func (u *User) Create() (string, error) {

	return "Implement Me! Register function", nil
}

//  GetAllUsers
func (u *User) GetAllUsers() (string, error) {
	return "Implement Me! To get all users function", nil
}

// GetByID
func (u *User) GetByID() {}

// Login user login page
func (u *User) Update() {}

// Login user login page
func (u *User) Delete() {
	// http.Redirect(w, r, "/login", http.StatusSeeOther)
	//render.Template(w, r, "login.page.html", &models.TemplateData{})
}
