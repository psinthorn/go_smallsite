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
	mux.Get("/", controllers.HandlerRepo.Home)
	mux.Get("/about", controllers.HandlerRepo.About)
	mux.Get("/contact", controllers.HandlerRepo.Contact)

	// Room routing section
	mux.Get("/room", controllers.HandlerRepo.Rooms)
	mux.Get("/rooms/superior", controllers.HandlerRepo.Superior)
	mux.Get("/rooms/deluxe", controllers.HandlerRepo.Deluxe)

	mux.Get("/users/login", controllers.HandlerRepo.Login)
	mux.Get("/users/logout", controllers.HandlerRepo.Logout)

	// Reservation routing section
	mux.Get("/search-availability", controllers.HandlerRepo.SearchAvailability)
	mux.Post("/search-availability", controllers.HandlerRepo.PostSearchAvailability)
	mux.Post("/search-availability-response", controllers.HandlerRepo.AvailabilityResponse)
	mux.Get("/reservation", controllers.HandlerRepo.Reservation)
	mux.Post("/reservation", controllers.HandlerRepo.PostReservation)
	mux.Get("/reservation-summary", controllers.HandlerRepo.ReservationSummary)

	// Admin routing section
	mux.Get("/admin/dashboard", controllers.HandlerRepo.Contact)

	return mux

}
