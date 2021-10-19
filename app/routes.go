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

	// Section: General Pages routing
	mux.Get("/", controllers.HandlerRepo.Home)
	mux.Get("/about", controllers.HandlerRepo.About)
	mux.Get("/contact", controllers.HandlerRepo.Contact)

	// Section: Room routing
	mux.Get("/rooms", controllers.HandlerRepo.Rooms)
	mux.Get("/rooms/superior", controllers.HandlerRepo.Superior)
	mux.Get("/rooms/deluxe", controllers.HandlerRepo.Deluxe)

	// Section: users routing
	mux.Get("/users/login", controllers.HandlerRepo.Login)
	mux.Post("/users/login", controllers.HandlerRepo.PostLogin)
	mux.Get("/users/logout", controllers.HandlerRepo.Logout)

	// Section: Reservation routing
	mux.Route("/rooms", func(mux chi.Router) {
		// search form
		mux.Get("/search-availability", controllers.HandlerRepo.SearchAvailability)
		// search all room availability
		mux.Post("/search-availability", controllers.HandlerRepo.PostSearchAvailability)
		// choose available room for make reservation
		mux.Get("/reservation/choose-room/{id}/{type}/{no}", controllers.HandlerRepo.ChooseRoom)

		// serch room available by room type and return as json format
		mux.Post("/search-availability-response", controllers.HandlerRepo.AvailabilityJson)
		// searc availability by room type
		mux.Get("/reservation-by-room-type", controllers.HandlerRepo.ReservationByRoomType)

		// reservation form
		mux.Get("/reservation", controllers.HandlerRepo.Reservation)
		// create new reservation
		mux.Post("/reservation", controllers.HandlerRepo.PostReservation)
		// show summary reservation
		mux.Get("/reservation-summary", controllers.HandlerRepo.ReservationSummary)

	})

	// Section: Promotions routing
	mux.Route("/promotions", func(mux chi.Router) {
		// search form
		// search form
		mux.Get("/", controllers.HandlerRepo.PromotionTypes)
		//mux.Get("/promotion-details/{id}", controllers.HandlerRepo.PromotionDetails)
		mux.Get("/promotion-choose-room/{type}/{id}", controllers.HandlerRepo.PromotionRoomType)

		mux.Get("/lists", controllers.HandlerRepo.PromotionsList)

		// // search all room availability
		// mux.Post("/search-promotion-availability", controllers.HandlerRepo.PostSearchAvailability)
		// // choose available room for make reservation
		// mux.Get("/reseration/choose-room/{id}/{type}/{no}", controllers.HandlerRepo.ChooseRoom)

		// // serch room available by room type and return as json format
		// mux.Post("/search-availability-response", controllers.HandlerRepo.AvailabilityJson)
		// // searc availability by room type
		// mux.Get("/reservation-by-room-type", controllers.HandlerRepo.ReservationByRoomType)

		// // reservation form
		// mux.Get("/reservation", controllers.HandlerRepo.Reservation)
		// // create new reservation
		// mux.Post("/reservation", controllers.HandlerRepo.PostReservation)
		// // show summary reservation
		// mux.Get("/reservation-summary", controllers.HandlerRepo.ReservationSummary)

	})

	// Administrator Section
	// this section is required authentication to get full access authorization

	mux.Route("/admin", func(mux chi.Router) {

		// // Authentication middleware
		// // all to below routes is need to authorize by this middleware
		// mux.Use(utils.Middleware.Auth)

		// Dasboard Section
		// show summary dasboard
		mux.Get("/dashboard", controllers.HandlerRepo.AdminDashBoard)

		// Section: Content
		// Creat content
		mux.Get("/contents/create", controllers.HandlerRepo.ContentForm)
		mux.Post("/contetns", controllers.HandlerRepo.PostContent)
		// Get content
		mux.Get("/contents/{id}", controllers.HandlerRepo.ShowContent)
		mux.Get("/contents", controllers.HandlerRepo.ContentLists)
		// Edit content
		mux.Get("/contents/edit", controllers.HandlerRepo.EditContentForm)
		mux.Post("/contetns/edit/{id}", controllers.HandlerRepo.EditContent)
		// Delete content
		mux.Post("/contents/delete/{id}", controllers.HandlerRepo.DeleteContent)

		// Section: Reservation
		// Add new reservation
		mux.Get("/reservations/form", controllers.HandlerRepo.ReservationAddForm)
		mux.Post("/reservations", controllers.HandlerRepo.ReservationAdd)
		//mux.Get("/reservations/{id}", controllers.HandlerRepo.Promotion)
		// Show all reservation
		mux.Get("/reservations", controllers.HandlerRepo.ReservationLists)
		mux.Get("/reservations/new-reservations", controllers.HandlerRepo.NewReservationLists)
		mux.Get("/reservations/calendar", controllers.HandlerRepo.ReservationCalendar)

		// Edit Reservation
		mux.Get("/reservations/edit/{id}", controllers.HandlerRepo.ReservationEditForm)
		mux.Post("/reservations/edit/{id}", controllers.HandlerRepo.ReservationEdit)
		// Delete Reservation
		mux.Post("/reservations/delete/{id}", controllers.HandlerRepo.ReservationDelete)

		// Rooms Section
		// this section will show all routes that concern about rooms

		// Room type
		// show all room type
		mux.Get("/rooms/roomtype", controllers.HandlerRepo.AddNewRoomTypeForm)
		// show form afor add room type
		mux.Get("/rooms/roomtype/new", controllers.HandlerRepo.AddNewRoomTypeForm)
		// add new room type
		mux.Post("/rooms/roomtype", controllers.HandlerRepo.AddNewRoomType)

		// Rooms
		// show form for add room
		mux.Get("/rooms/new", controllers.HandlerRepo.AddNewRoomForm)
		// add rooms
		mux.Post("/rooms", controllers.HandlerRepo.RoomGetAll)
		// show all rooms
		mux.Get("/rooms", controllers.HandlerRepo.RoomGetAll)

		// Room Status
		// add new room status
		mux.Get("/rooms/room-status/new", controllers.HandlerRepo.AddNewRoomStatusForm)
		// add new room status
		mux.Post("/rooms/room-status", controllers.HandlerRepo.AddNewRoomStatus)
		// show all rooms status
		mux.Get("/rooms/room-status", controllers.HandlerRepo.AddNewRoomStatusForm)

		// Section: Promotion
		mux.Get("/promotions", controllers.HandlerRepo.PromotionsList)
		mux.Get("/promotions/new", controllers.HandlerRepo.PromotionForm)
		mux.Get("/promotions/{id}", controllers.HandlerRepo.Promotion)
		mux.Get("/promotions/{id}/edit", controllers.HandlerRepo.PromotionTypes)
		mux.Get("/promotions/types", controllers.HandlerRepo.PromotionTypesList)
		// mux.Delete("/promotions/{id}", controller.)

		// Section: User
		// Control and manage all users
		// this section show all routes about user management
		mux.Get("/users/register", controllers.HandlerRepo.AddNewUserForm)
		// add new user
		mux.Post("/users", controllers.HandlerRepo.AddNewUser)
		// show all user
		mux.Get("/users", controllers.HandlerRepo.GetAllUsers)

	})

	return mux

}
