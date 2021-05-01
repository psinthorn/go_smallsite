package main

import (
	"encoding/gob"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/psinthorn/go_smallsite/models/reservations"
	"github.com/psinthorn/go_smallsite/models/users"
)

// CreateSession and store session to AppConfig.Session
func CreateSession() {

	// register all models to session maybe use or not use
	gob.Register(users.User{})
	gob.Register(reservations.RoomStatus{})
	gob.Register(reservations.RoomAlotmentStatus{})
	gob.Register(reservations.Reservation{})

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.IsProduction

	appConfig.Session = session

}

// SessionLoad middleware adds and save session on every requests
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
