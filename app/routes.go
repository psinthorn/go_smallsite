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

	mux.Get("/users/getall", controllers.HandlerRepo.GetAllUsers)
	mux.Get("/users/login", controllers.HandlerRepo.Login)
	mux.Post("/users/login", controllers.HandlerRepo.PostLogin)
	mux.Get("/users/logout", controllers.HandlerRepo.Logout)

	// Reservation routing section
	mux.Get("/rooms/search-availability", controllers.HandlerRepo.SearchAvailability)
	mux.Post("/rooms/search-availability", controllers.HandlerRepo.PostSearchAvailability)
	mux.Post("/rooms/search-availability-response", controllers.HandlerRepo.AvailabilityJson)
	mux.Get("/rooms/reservation-by-room-type", controllers.HandlerRepo.ReservationByRoomType)
	mux.Get("/rooms", controllers.HandlerRepo.Rooms)
	mux.Get("/rooms/reservation", controllers.HandlerRepo.Reservation)
	mux.Post("/rooms/reservation", controllers.HandlerRepo.PostReservation)
	// mux.Post("/rooms/reservation/by-room-type/{type}", controllers.HandlerRepo.PostReservation)
	mux.Get("/rooms/reseration/choose-room/{id}/{type}/{no}", controllers.HandlerRepo.ChooseRoom)
	mux.Get("/rooms/reservation-summary", controllers.HandlerRepo.ReservationSummary)

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(utils.Middleware.Auth)
		// Admin routing section
		mux.Get("/dashboard", controllers.HandlerRepo.AdminDashBoard)

		// Admin user management
		mux.Get("/users/register", controllers.HandlerRepo.AddNewUserForm)
		mux.Post("/users/user/new", controllers.HandlerRepo.AddNewUser)

		// Room Status
		mux.Get("/rooms/room-status", controllers.HandlerRepo.AddNewRoomStatusForm)
		mux.Post("/rooms/room-status/new", controllers.HandlerRepo.AddNewRoomStatus)

		// Room Type
		mux.Get("/rooms/roomtype", controllers.HandlerRepo.AddNewRoomTypeForm)
		mux.Post("/rooms/roomtype/new", controllers.HandlerRepo.AddNewRoomType)

		mux.Get("/rooms", controllers.HandlerRepo.RoomGetAll)
		mux.Get("/rooms/room", controllers.HandlerRepo.AddNewRoomForm)

	})

	return mux

}
