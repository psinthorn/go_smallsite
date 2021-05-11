package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/psinthorn/go_smallsite/configs"
	controllers "github.com/psinthorn/go_smallsite/controllers/handlers"
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
	mux.Get("/rooms", controllers.HandlerRepo.Rooms)
	mux.Get("/rooms/superior", controllers.HandlerRepo.Superior)
	mux.Get("/rooms/deluxe", controllers.HandlerRepo.Deluxe)

	mux.Get("/users/register", controllers.HandlerRepo.Register)
	mux.Get("/users/getall", controllers.HandlerRepo.GetAllUsers)
	mux.Get("/users/login", controllers.HandlerRepo.Login)
	mux.Get("/users/logout", controllers.HandlerRepo.Logout)

	// Reservation routing section
	mux.Get("/rooms/search-availability", controllers.HandlerRepo.SearchAvailability)
	mux.Post("/rooms/search-availability", controllers.HandlerRepo.PostSearchAvailability)
	mux.Post("/rooms/search-availability-response", controllers.HandlerRepo.AvailabilityResponse)
	mux.Get("/rooms", controllers.HandlerRepo.Rooms)
	mux.Get("/rooms/reservation", controllers.HandlerRepo.Reservation)
	mux.Post("/rooms/reservation", controllers.HandlerRepo.PostReservation)
	mux.Get("/rooms/reservation-summary", controllers.HandlerRepo.ReservationSummary)

	// Admin routing section
	mux.Get("/admin/dashboard", controllers.HandlerRepo.Contact)

	// Room Status
	mux.Get("/admin/rooms/room-status", controllers.HandlerRepo.AddNewRoomStatusForm)
	mux.Post("/admin/rooms/room-status/new", controllers.HandlerRepo.AddNewRoomStatus)

	// Room Type
	mux.Get("/admin/rooms/roomtype", controllers.HandlerRepo.AddNewRoomTypeForm)
	mux.Post("/admin/rooms/roomtype/new", controllers.HandlerRepo.AddNewRoomType)

	mux.Get("/admin/rooms", controllers.HandlerRepo.RoomGetAll)
	mux.Get("/admin/rooms/new", controllers.HandlerRepo.RoomGetForm)
	mux.Post("/admin/rooms/new", controllers.HandlerRepo.RoomCreate)

	return mux

}
