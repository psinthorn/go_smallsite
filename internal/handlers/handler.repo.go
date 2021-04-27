package handlers

import (
	"github.com/psinthorn/go_smallsite/internal/configs"
)

// Repo
var HandlerRepo *Repository

type Repository struct {
	App *configs.AppConfig
}

// NewRepository
func NewRepository(a *configs.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers
func NewHandlers(r *Repository) {
	HandlerRepo = r
}
