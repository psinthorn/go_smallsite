package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/psinthorn/go_smallsite/configs"
	"github.com/psinthorn/go_smallsite/controllers"
	"github.com/psinthorn/go_smallsite/internal/utils"
)

// Routes use to map url with controller func
func routes(app *configs.AppConfig) http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(utils.Middleware.NoSurf)
	mux.Use(SessionLoad)

	//mux.Use(utils.Middleware.WriteToConsole)
	// Static file folder
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// Pages routing section
	mux.Get("/", controllers.PagesController.Home)
	mux.Get("/about", controllers.PagesController.About)
	mux.Get("/contact", controllers.PagesController.Contact)

	// Room routing section
	mux.Get("/room", controllers.PagesController.Rooms)
	mux.Get("/superior", controllers.PagesController.Superior)
	mux.Get("/deluxe", controllers.PagesController.Deluxe)

	mux.Get("/user/login", controllers.UsersController.Login)
	mux.Get("/user/logout", controllers.UsersController.Logout)

	// Reservation routing section
	mux.Get("/search-availability", controllers.ReservationsController.SearchAvailability)
	mux.Post("/search-availability", controllers.ReservationsController.PostSearchAvailability)
	mux.Post("/search-availability-response", controllers.ReservationsController.AvailabilityResponse)
	mux.Get("/reservation", controllers.ReservationsController.Reservation)
	mux.Post("/reservation", controllers.ReservationsController.PostReservation)
	mux.Get("/reservation-summary", controllers.ReservationsController.ReservationSummary)

	// Admin routing section
	mux.Get("/admin/dashboard", controllers.PagesController.Contact)

	return mux

}
