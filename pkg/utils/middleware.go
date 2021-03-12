package utils

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

var Middleware middleware

type middleware struct{}

// Basice middleware for example
func (mdw *middleware) WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("1st basic middleware")
		next.ServeHTTP(w, r)
	})
}

// NoSurf CSRF token generate
func (mdw *middleware) NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
