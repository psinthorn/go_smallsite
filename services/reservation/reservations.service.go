package services

import "github.com/psinthorn/go_smallsite/models/reservations"

var ReservationService reservationService

type reservationService struct{}

type reservationServiceInterface interface {
	CreateReservation(reservations.Reservation) (*reservations.Reservation, error)
	GetAllReservation() (*reservations.Reservation, error)
	GetReservationByID(int) (*reservations.Reservation, error)
	UpdateReservation(int) (*reservations.Reservation, error)
	DeleteReservation(int) (*reservations.Reservation, error)
}

// Create new reservation
func (rsvn *reservationService) Create() (*reservations.Reservation, error) {

	return nil, nil
}

// GetAll get all reservation list
func (rsvn *reservationService) GetAll() (*reservations.Reservation, error) {
	return nil, nil
}

// GetByID get reservation by ID
func (rsvn *reservationService) GetByID(id int) (*reservations.Reservation, error) {
	return nil, nil
}

// Update get reservation by ID
func (rsvn *reservationService) Update(id int) (*reservations.Reservation, error) {
	return nil, nil
}

// GetByID get reservation by ID
func (rsvn *reservationService) Delete(id int) (*reservations.Reservation, error) {
	return nil, nil
}
