package handlers

import (
	"github.com/psinthorn/go_smallsite/configs"
	dbrepo "github.com/psinthorn/go_smallsite/datasources"
	drivers "github.com/psinthorn/go_smallsite/datasources/drivers"
)

// Repo
var HandlerRepo *Repository

type Repository struct {
	App       *configs.AppConfig
	DBConnect interface{}
}

// NewRepository is create new handler that holds application config and database connection
func NewHandlerRepository(a *configs.AppConfig, db *drivers.DB) *Repository {
	return &Repository{
		App:       a,
		DBConnect: dbrepo.NewDBConnectRepo(a, db.SQL),
	}
}

// NewHandlers
func NewHandlers(r *Repository) {
	HandlerRepo = r
}
