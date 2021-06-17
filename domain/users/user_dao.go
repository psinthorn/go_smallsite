package users

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/psinthorn/go_smallsite/datasources/drivers"
	"golang.org/x/crypto/bcrypt"
)

const (
	queryInsertUser     = `insert into users (first_name, last_name, email, password, access_level, status, created_at, updated_at) values ($1,$2,$3,$4,$5,$6,$7,$8) returning id`
	queryGetUserByEmail = `select id, first_name, last_name, email, password, access_level, status, created_at, updated_at from users where email = $1`
	queryUpdateUserById = `update users set first_name = $1, last_name = $2, email = $3, password = $4, access_level = $5, status =$6, updated_at = $7) where id = $8`
)

var UserService userDomainInterface = &User{}

type User user
type userDomainInterface interface {
	Create(User) (int, error)
	GetAllUsers() ([]User, error)
	GetByID(int) (User, error)
	Update(User) error
	Delete(int) (int, error)
	Authenticate(string, string) (User, error)
}

// Register new user
func (m *User) Create(u User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
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

// GetByID
func (u *User) GetByUserEmail(email string) (User, error) {
	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancle()
	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		panic(err)
	}

	var user User
	row := dbConn.SQL.QueryRowContext(ctx, queryGetUserByEmail, email)
	err = row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.AccessLevel,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err != nil {
		return user, err
	}

	return user, nil
}

// Update user by id
func (u *User) Update(user User) error {
	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancle()
	dbConn, err := drivers.ConnectDB("pgx", drivers.PgDsn)
	if err != nil {
		panic(err)
	}

	_, err = dbConn.SQL.ExecContext(ctx, queryUpdateUserById,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.AccessLevel,
		user.Status,
		time.Now(),
		user.ID)
	if err != nil {
		return err
	}

	return nil
}

// Login user login page
func (u *User) Delete(id int) (int, error) {

	return 0, nil
	// http.Redirect(w, r, "/login", http.StatusSeeOther)
	//render.Template(w, r, "login.page.html", &models.TemplateData{})
}

func (u *User) Authenticate(email, password string) (User, error) {

	// get user by email from database
	user, err := u.GetByUserEmail(email)
	if err != nil {
		log.Println(err)
		return user, errors.New("please check your email address!")
	}

	log.Println(user)
	hashedPassword := user.Password
	// compare password
	// return error if compare is false  return id and token if compare is true
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return user, errors.New("incorrect password")
	} else if err != nil {
		return user, err
	}

	return user, nil

}
