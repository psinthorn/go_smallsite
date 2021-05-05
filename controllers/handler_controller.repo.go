package controllers

import (
	"github.com/psinthorn/go_smallsite/configs"
	"github.com/psinthorn/go_smallsite/datasources/drivers"
	repository "github.com/psinthorn/go_smallsite/domain"
	"github.com/psinthorn/go_smallsite/domain/dbrepo"
)

// Repo
var HandlerRepo *Repository

type Repository struct {
	App *configs.AppConfig
	DB  repository.DatabaseRepo
	// Pages        pagesControllerInterface
	// Reservations reservationsControllerInterface
	// Users        usersControllerInterface
}

// NewRepository is create new handler that holds application config and database connection
func NewHandlerRepository(a *configs.AppConfig, db *drivers.DbConn) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewDBConnectRepo(a, db.SQL),
	}
}

// NewHandlers
func NewHandlers(r *Repository) {
	HandlerRepo = r
}
