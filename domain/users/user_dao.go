package users

import (
	"context"
	"fmt"
	"time"

	"github.com/psinthorn/go_smallsite/datasources/drivers"
)

const (
	queryInsertUser = `insert into users (first_name, last_name, email, password, access_level, status, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8) returning id`
)

var UserService userDomainInterface = &User{}

type User user
type userDomainInterface interface {
	Create(User) (int, error)
	GetAllUsers() ([]User, error)
	GetByID(int) (User, error)
	Update(int) (User, error)
	Delete(int) (int, error)
}

// Register new user
func (m *User) Create(u User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	fmt.Println(u)

	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		panic(err)
	}

	var newUserID int
	err = dbConn.SQL.QueryRowContext(ctx, queryInsertUser, u.FirstName, u.LastName, u.Email, u.Password, u.AccessLevel, u.Status, u.CreatedAt, u.UpdatedAt).Scan(&newUserID)
	if err != nil {
		return 0, err
	}
	defer dbConn.SQL.Close()

	fmt.Println(newUserID)

	return newUserID, nil
}

//  GetAllUsers
func (m *User) GetAllUsers() ([]User, error) {
	var users []User
	return users, nil
}

// GetByID
func (u *User) GetByID(id int) (User, error) {
	var user User
	return user, nil
}

// Login user login page
func (u *User) Update(id int) (User, error) {
	var user User
	return user, nil
}

// Login user login page
func (u *User) Delete(id int) (int, error) {

	return 0, nil
	// http.Redirect(w, r, "/login", http.StatusSeeOther)
	//render.Template(w, r, "login.page.html", &models.TemplateData{})
}
