package dbrepo

import (
	"database/sql"

	"github.com/psinthorn/go_smallsite/configs"
	repository "github.com/psinthorn/go_smallsite/domain"
)

type sqlDbRepo struct {
	App *configs.AppConfig
	DB  *sql.DB
}

// NewDBRepo will hold database connection pool and app configs
// func NewDBConnectRepo(ac *configs.AppConfig, conn *sql.DB) interface{} {
func NewDBConnectRepo(ac *configs.AppConfig, conn *sql.DB) repository.DatabaseRepo {
	return &sqlDbRepo{
		App: ac,
		DB:  conn,
	}
}
