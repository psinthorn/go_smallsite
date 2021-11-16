package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/psinthorn/go_smallsite/configs"
	controllers "github.com/psinthorn/go_smallsite/controllers/handlers"
)

// Routes use to map url with controller func
func AdminRoutes(app *configs.AppConfig) http.Handler {

	mux := chi.NewRouter()

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
		mux.Post("/promotions", controllers.HandlerRepo.AddPromotion)
		mux.Get("/promotions", controllers.HandlerRepo.AdminPromotionsList)
		mux.Get("/promotions/new", controllers.HandlerRepo.PromotionForm)
		mux.Get("/promotions/{id}", controllers.HandlerRepo.Promotion)
		mux.Post("/promotions/{id}/update", controllers.HandlerRepo.UpdatePromotion)
		mux.Get("/promotions/{id}/delete", controllers.HandlerRepo.DeletePromotion)

		// Section: Promotion-Types
		mux.Post("/promotions-types", controllers.HandlerRepo.AddPromotionType)
		mux.Get("/promotions-types", controllers.HandlerRepo.AdminPromotionTypes)
		mux.Get("/promotions-types/new", controllers.HandlerRepo.PromotionTypeForm)
		mux.Get("/promotions-types/{id}", controllers.HandlerRepo.PromotionType)
		mux.Post("/promotions-types/{id}/update", controllers.HandlerRepo.UpdatePromotionType)
		mux.Get("/promotions-types/{id}/delete", controllers.HandlerRepo.DeletePromotionType)

		// // Section: Promotion-Types
		// mux.Post("/promotions-ratetypes", controllers.HandlerRepo.AddPromotionRateType)
		// mux.Get("/promotions-ratetypes", controllers.HandlerRepo.AdminPromotionRateTypes)
		// mux.Get("/promotions-ratetypes/new", controllers.HandlerRepo.PromotionRateTypeForm)
		// mux.Get("/promotions-ratetypes/{id}", controllers.HandlerRepo.PromotionRateType)
		// mux.Post("/promotions-ratetypes/{id}/update", controllers.HandlerRepo.UpdatePromotionRateType)
		// mux.Get("/promotions-ratetypes/{id}/delete", controllers.HandlerRepo.DeletePromotionRateType)

		// Section: Room Rate
		mux.Post("/roomrates", controllers.HandlerRepo.AddPromotionType)
		mux.Get("/roomrate", controllers.HandlerRepo.AdminPromotionTypes)
		mux.Get("/roomrates/new", controllers.HandlerRepo.PromotionTypeForm)
		mux.Get("/roomrates/{id}", controllers.HandlerRepo.PromotionType)
		mux.Post("/roomrate/{id}/update", controllers.HandlerRepo.UpdatePromotionType)
		mux.Get("/roomrate/{id}/delete", controllers.HandlerRepo.DeletePromotionType)

		// Section: Room Rate Type
		// desc: type of room rate like rackrate, wholesale, ota, member, promotion
		mux.Post("/roomrate-types", controllers.HandlerRepo.AddPromotionType)
		mux.Get("/roomrate-type", controllers.HandlerRepo.AdminPromotionTypes)
		mux.Get("/roomrates-types/new", controllers.HandlerRepo.PromotionTypeForm)
		mux.Get("/roomrates-types/{id}", controllers.HandlerRepo.PromotionType)
		mux.Post("/room-types/{id}/update", controllers.HandlerRepo.UpdatePromotionType)
		mux.Get("/roomrate-types/{id}/delete", controllers.HandlerRepo.DeletePromotionType)

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
