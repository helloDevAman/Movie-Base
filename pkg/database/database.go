package database

import (
	"database/sql"
)

// Generic interface for any database connection
type Database interface {
	Connect() (*sql.DB, error)
	Close() error
}
