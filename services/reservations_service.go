package services

import (
	"github.com/psinthorn/go_smallsite/domain/dbrepo"
)

var ReservationService reservationServiceInterface = &reservationService{}

type reservationService struct{}

type reservationServiceInterface interface {
	Create(dbrepo.Reservation) (*dbrepo.Reservation, error)
	GetAll() (*dbrepo.Reservation, error)
	GetByID(int) (*dbrepo.Reservation, error)
	Update(int) (*dbrepo.Reservation, error)
	Delete(int) (*dbrepo.Reservation, error)
}

// Create new reservation
func (s *reservationService) Create(rsvn dbrepo.Reservation) (*dbrepo.Reservation, error) {

	return nil, nil
}

// GetAll get all reservation list
func (s *reservationService) GetAll() (*dbrepo.Reservation, error) {
	return nil, nil
}

// GetByID get reservation by ID
func (rsvn *reservationService) GetByID(id int) (*dbrepo.Reservation, error) {
	return nil, nil
}

// Update get reservation by ID
func (rsvn *reservationService) Update(id int) (*dbrepo.Reservation, error) {
	return nil, nil
}

// GetByID get reservation by ID
func (rsvn *reservationService) Delete(id int) (*dbrepo.Reservation, error) {
	return nil, nil
}
