package persist

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

type DbConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
	SSLMode  string
}

func Connect() error {
	config := DbConfig{
		Host:     "127.0.0.1",
		Port:     5432,
		Name:     "dogs",
		User:     "dogs",
		Password: "dogs",
		SSLMode:  "disable",
	}

	configString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.User, config.Password, config.Name, config.SSLMode)

	db, err := sql.Open(
		"postgres",
		configString)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	log.Println("Successfully connected to database!")
	DB = db
	return nil
}

func GetDB() *sql.DB {
	return DB
}
