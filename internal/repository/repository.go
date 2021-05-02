package repository

import (
	"github.com/psinthorn/go_smallsite/domain/reservations"
	"github.com/psinthorn/go_smallsite/domain/users"
)

type DatabaseRepo interface {
	CreateUser() (users.User, error)
	GetAllUsers() (users.User, error)
	//GetUserByEmail(email string) (users.User, error)
	UpdateUserByID(id int) (users.User, error)
	InsertReservation(rsvn reservations.Reservation) (int, error)
}
