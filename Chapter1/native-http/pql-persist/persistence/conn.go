package persist

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DbConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
	SSLMode  string
}

func StartDB() {
	config := DbConfig{
		Host:     "127.0.0.1",
		Port:     5432,
		Name:     "postgres",
		User:     "fady",
		Password: "fady",
		SSLMode:  "disable",
	}

	configString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.User, config.Password, config.Name, config.SSLMode)

	db, err := sql.Open(
		"postgres",
		configString)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("error connecting to database: %v", err)
	}

	log.Println("Successfully connected to database!")
}
