package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/psinthorn/go_smallsite/pkg/configs"
	"github.com/psinthorn/go_smallsite/pkg/handlers"
	"github.com/psinthorn/go_smallsite/pkg/utils"
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

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/room", handlers.Repo.Rooms)
	mux.Get("/superior", handlers.Repo.Superior)
	mux.Get("/deluxe", handlers.Repo.Deluxe)
	mux.Get("/check-alotment", handlers.Repo.CheckAlotment)
	mux.Get("/search-alotment", handlers.Repo.SearchAlotment)
	mux.Get("/reservation", handlers.Repo.Reservation)
	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Get("/contact", handlers.Repo.Contact)

	return mux

}
