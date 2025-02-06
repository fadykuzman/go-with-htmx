package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
	MaxConns int
	Timeout  time.Duration
}

type DB struct {
	*sql.DB
}

func DefaultDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     "localhost",
		Port:     5432,
		SSLMode:  "disable",
		MaxConns: 20,
		Timeout:  time.Second * 5,
	}
}

func NewConnection(config DatabaseConfig) (*DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User,
		config.Password, config.Name, config.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return &DB{db}, nil
}
