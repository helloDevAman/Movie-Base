package database

import (
	"database/sql"
	"fmt"

	"github.com/helloDevAman/movie-base/config"
	_ "github.com/lib/pq"
)

// PostgresConnector implements the Database interface
type PostgresConnector struct {
	cfg *config.Config
	db  *sql.DB
}

// ConnectNewPostgresDB initializes the structure but doesn't connect immediately.
func ConnectNewPostgresDB(cfg *config.Config) *PostgresConnector {
	return &PostgresConnector{cfg: cfg}
}

// Connect returns an existing DB connection or creates a new one
func (p *PostgresConnector) Connect() (*sql.DB, error) {
	// If connection already exists, return it
	if p.db != nil {
		return p.db, nil
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p.cfg.DB.Host, p.cfg.DB.Port, p.cfg.DB.User, p.cfg.DB.Pass, p.cfg.DB.Name,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to PostgreSQL: %v", err)
	}

	// Verify the connection
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("unable to ping PostgreSQL: %v", err)
	}

	p.db = db
	return p.db, nil
}

// Close properly closes the database connection
func (p *PostgresConnector) Close() error {
	if p.db == nil {
		return fmt.Errorf("database connection is already closed")
	}

	err := p.db.Close()
	if err != nil {
		return fmt.Errorf("unable to close db: %v", err)
	}

	p.db = nil
	return nil
}
