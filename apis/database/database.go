package database

import (
	"errors"

	"github.com/helloDevAman/movie-base/apis/config"
)

var ErrUnsupportedDatabase = errors.New("unsupported database type")

type Database interface {
	Connect() error
	Close()
	GetConnection() interface{}
}

func NewDatabaseConnector(dbType string, cfg *config.Config) (Database, error) {
	switch dbType {
	case "postgres":
		return NewPostgresConnector(cfg)
	default:
		return nil, ErrUnsupportedDatabase
	}
}
