package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/psinthorn/go_smallsite/pkg/handlers"
)

// Routes use to map url with controller func
func routes() http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// Static file folder
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Get("/", handlers.Home)
	mux.Get("/about", handlers.About)

	return mux

}
