package main

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/psinthorn/go_smallsite/pkg/configs"
)

var (
	appConfig configs.AppConfig
	session   *scs.SessionManager
)

// CreateSession and store session to AppConfig.Session
func CreateSession() {
	// change this to true when in production or create func auto checkig env
	appConfig.IsProduction = false

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
