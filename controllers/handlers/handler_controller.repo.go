package controllers

import (
	"github.com/psinthorn/go_smallsite/configs"
)

// Repo
var HandlerRepo *Repository

type Repository struct {
	App *configs.AppConfig
	//DB  interface{}
	// DB  repository.DatabaseRepo
	// Reservations reservationsControllerInterface
	// Users        usersControllerInterface
}

// NewRepository is create new handler that holds application config and database connection
func NewHandlerRepository(a *configs.AppConfig) *Repository {
	return &Repository{
		App: a,
		//DB:  dbrepo.NewDBConnectRepo(a, db.SQL),
	}
}

// NewHandlers
func NewHandlers(r *Repository) {
	HandlerRepo = r
}
