package repository

import "github.com/psinthorn/go_smallsite/internal/models"

type DatabaseRepo interface {
	GetAllUsers() bool
	InsertReservation(rsvn models.Reservation) (int, error)
}
