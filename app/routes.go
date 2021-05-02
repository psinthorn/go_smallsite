package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/psinthorn/go_smallsite/configs"
	"github.com/psinthorn/go_smallsite/internal/handlers"
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

	mux.Get("/", handlers.HandlerRepo.Home)
	mux.Get("/about", handlers.HandlerRepo.About)
	mux.Get("/contact", handlers.HandlerRepo.Contact)
	mux.Get("/user/login", handlers.HandlerRepo.Login)
	mux.Get("/user/logout", handlers.HandlerRepo.Logout)

	mux.Get("/room", handlers.HandlerRepo.Rooms)
	mux.Get("/superior", handlers.HandlerRepo.Superior)
	mux.Get("/deluxe", handlers.HandlerRepo.Deluxe)

	// mux.Get("/check-alotment", handlers.HandlerRepo.CheckAlotment)
	mux.Get("/search-availability", handlers.HandlerRepo.SearchAvailability)
	mux.Post("/search-availability", handlers.HandlerRepo.PostSearchAvailability)
	mux.Post("/search-availability-response", handlers.HandlerRepo.AvailabilityResponse)

	//mux.Get("/reservation", handlers.HandlerRepo.Reservation)
	mux.Get("/make-reservation", handlers.HandlerRepo.Reservation)
	mux.Post("/make-reservation", handlers.HandlerRepo.PostReservation)
	mux.Get("/reservation-summary", handlers.HandlerRepo.ReservationSummary)

	mux.Get("/admin/dashboard", handlers.HandlerRepo.Contact)

	return mux

}
