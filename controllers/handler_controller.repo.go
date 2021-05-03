package controllers

import (
	"github.com/psinthorn/go_smallsite/configs"
	drivers "github.com/psinthorn/go_smallsite/datasources/drivers"
)

// Repo
var HandlerRepo *Repository

type Repository struct {
	App *configs.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepository is create new handler that holds application config and database connection
func NewHandlerRepository(a *configs.AppConfig, db *drivers.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewDBConnectRepo(a, db.SQL),
	}
}

// NewHandlers
func NewHandlers(r *Repository) {
	HandlerRepo = r
}
