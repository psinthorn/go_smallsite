package dbrepo

import (
	"database/sql"

	"github.com/psinthorn/go_smallsite/internal/repository"
	"github.com/psinthorn/go_smallsite/pkg/configs"
)

type postgresDBRepo struct {
	App *configs.AppConfig
	DB  *sql.DB
}

func NewPostgresDBRepo(conn *sql.DB, ac *configs.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: ac,
		DB:  conn,
	}
}
