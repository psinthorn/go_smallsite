package dbrepo

import (
	"database/sql"

	"github.com/psinthorn/go_smallsite/internal/configs"
	"github.com/psinthorn/go_smallsite/internal/repository"
)

type SQLDBRepo struct {
	App *configs.AppConfig
	DB  *sql.DB
}

// NewDBRepo will hold database connection pool and app configs
func NewDBConnectRepo(ac *configs.AppConfig, conn *sql.DB) repository.DatabaseRepo {
	return &SQLDBRepo{
		App: ac,
		DB:  conn,
	}
}
