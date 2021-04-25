package main

import (
	"encoding/gob"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/psinthorn/go_smallsite/internal/models"
)

// CreateSession and store session to AppConfig.Session
func CreateSession() {

	// register reservation model to session
	gob.Register(models.Reservation{})

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
