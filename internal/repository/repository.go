package repository

import "github.com/psinthorn/go_smallsite/models/reservations"

type DatabaseRepo interface {
	GetAllUsers() bool
	InsertReservation(rsvn reservations.Reservation) (int, error)
}
