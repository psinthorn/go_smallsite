package repository

type DatabaseRepo interface {
	Home()
	About()
	Rooms()
	Superior()
	Deluxe()
	Contact()

	Register()
	GetAllUsers()
	Login()
	Logout()

	SearchAvailability()
	PostSearchAvailability()
	AvailabilityResponse()
	Reservation()
	PostReservation()
	ReservationSummary()
}
