package database

import (
	"database/sql"
	"errors"

	"github.com/helloDevAman/movie-base/apis/config"
)

var ErrUnsupportedDatabase = errors.New("unsupported database type")

type Database interface {
	Connect() error
	Close()
	GetConnection() *sql.DB
}

func NewDatabaseConnector(dbType string, cfg *config.Config) (Database, error) {
	switch dbType {
	case "postgres":
		return NewPostgresConnector(cfg)
	default:
		return nil, ErrUnsupportedDatabase
	}
}
