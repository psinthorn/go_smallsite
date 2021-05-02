package services

import (
	"database/sql"

	"github.com/psinthorn/go_smallsite/configs"
	"github.com/psinthorn/go_smallsite/models/reservations"
)

var ReservationService reservationServiceInterface = &reservationService{}

type reservationService struct {
	App *configs.AppConfig
	DB  *sql.DB
}

type reservationServiceInterface interface {
	CreateReservation(reservations.Reservation) (*reservations.Reservation, error)
	GetAllReservation() (*reservations.Reservation, error)
	GetReservationByID(int) (*reservations.Reservation, error)
	UpdateReservation(int) (*reservations.Reservation, error)
	DeleteReservation(int) (*reservations.Reservation, error)
}

// Create new reservation
func (rs *reservationService) CreateReservation(rsvn reservations.Reservation) (*reservations.Reservation, error) {

	return nil, nil
}

// GetAll get all reservation list
func (rs *reservationService) GetAllReservation() (*reservations.Reservation, error) {
	return nil, nil
}

// GetByID get reservation by ID
func (rs *reservationService) GetReservationByID(id int) (*reservations.Reservation, error) {
	return nil, nil
}

// Update get reservation by ID
func (rs *reservationService) UpdateReservation(id int) (*reservations.Reservation, error) {
	return nil, nil
}

// GetByID get reservation by ID
func (rs *reservationService) DeleteReservation(id int) (*reservations.Reservation, error) {
	return nil, nil
}
