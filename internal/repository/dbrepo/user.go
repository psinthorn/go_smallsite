package dbrepo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/psinthorn/go_smallsite/domain/users"
	"golang.org/x/crypto/bcrypt"
)

const (
	InsertUser     = "INSERT INTO users (first_name, last_name, email, password, access_level, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7)"
	GetUserByID    = "SELECT id, first_name, last_name, email, password, access_level, created_at, updated_at FROM users WHERE id=$1;"
	GetUserByEmail = "SELECT id, password FROM users WHERE email = $1"
	UpdateUserByID = "UPDATE users SET first_name = $1, last_name = $2, email = $3, access_level = $4, updated_at = $5 WHERE id = $6"
)

// CreateUser
func (db *SQLDBRepo) CreateUser() (users.User, error) {
	var user users.User
	return user, nil
}

func (db *SQLDBRepo) GetAllUsers() (users.User, error) {
	var allUsers users.User
	return allUsers, nil
}

// GetUserByID
func (db *SQLDBRepo) GetUserByID(id int) (users.User, error) {
	var u users.User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := db.DB.QueryRowContext(ctx, GetUserByEmail, id)
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.AccessLevel, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return u, err
	}
	return u, nil
}

// UpdateUserByID
func (db *SQLDBRepo) UpdateUserByID(id int) (users.User, error) {
	var u users.User
	ctx, concel := context.WithTimeout(context.Background(), 3*time.Second)
	defer concel()

	row := db.DB.QueryRowContext(ctx, UpdateUserByID, u)
	err := row.Scan(&u.FirstName, &u.LastName, &u.Email, &u.AccessLevel, &u.UpdatedAt)
	if err != nil {
		return u, err
	}

	return u, nil
}

// Authenthication
func (db *SQLDBRepo) Authenthicate(email, password string) (int, bool, error) {
	var u users.User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := db.DB.QueryRowContext(ctx, GetUserByEmail, email)
	err := row.Scan(&u.ID, &u.Password)
	if err != nil {
		return u.ID, false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return u.ID, false, errors.New("incorrect password")
	} else {
		if err != nil {
			fmt.Println("You're good to go")
			return 0, false, errors.New("Wrong password")
		}
	}

	return u.ID, true, nil
}
