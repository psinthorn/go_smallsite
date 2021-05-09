package main

import (
	"encoding/gob"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/psinthorn/go_smallsite/domain/dbrepo"
)

// CreateSession and store session to AppConfig.Session
func CreateSession() {

	// register all models to session maybe use or not use
	gob.Register(dbrepo.User{})
	gob.Register(dbrepo.Room{})
	gob.Register(dbrepo.RoomStatus{})
	gob.Register(dbrepo.RoomAllotmentStatus{})
	gob.Register(dbrepo.Reservation{})

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
