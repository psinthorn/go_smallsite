package dbrepo

import "time"

// User is the user model
type user struct {
	ID          int       `json: "id"`
	FirstName   string    `json: "first_name"`
	LastName    string    `json: "last_name"`
	Email       string    `json: "email"`
	Password    string    `json: "password"`
	AccessLevel int       `json: "access_level"`
	CreatedAt   time.Time `json: "created_at"`
	UpdatedAt   time.Time `json: "update_at"`
}
