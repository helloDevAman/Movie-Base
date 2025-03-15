package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/helloDevAman/movie-base/apis/config"
	_ "github.com/lib/pq"
)

type PostgresConnector struct {
	cfg *config.Config
	db  *sql.DB
}

func NewPostgresConnector(cfg *config.Config) (*PostgresConnector, error) {
	pg := &PostgresConnector{cfg: cfg}
	if err := pg.Connect(); err != nil {
		return nil, err
	}
	return pg, nil
}

func (p *PostgresConnector) Connect() error {
	// Postgres DSN (Connection String)
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p.cfg.DBHost, p.cfg.DBPort, p.cfg.DBUser, p.cfg.DBPass, p.cfg.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Unable to connect to Postgres: %v", err)
		return err
	}

	// Check the connection
	if err = db.Ping(); err != nil {
		log.Fatalf("Postgres ping failed: %v", err)
		return err
	}
	p.db = db
	return nil
}

func (p *PostgresConnector) Close() {
	if p.db != nil {
		p.db.Close()
		log.Println("Postgres connection closed")
	}
}

func (p *PostgresConnector) GetConnection() *sql.DB {
	return p.db
}
