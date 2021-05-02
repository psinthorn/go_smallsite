package dbrepo

import (
	"database/sql"

	"github.com/psinthorn/go_smallsite/configs"
)

type SQLDBRepo struct {
	App *configs.AppConfig
	DB  *sql.DB
}

// NewDBRepo will hold database connection pool and app configs
func NewDBConnectRepo(appConfig *configs.AppConfig, conn *sql.DB) interface{} {
	return &SQLDBRepo{
		App: appConfig,
		DB:  conn,
	}
}
