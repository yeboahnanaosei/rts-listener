package app

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Model represents all our database models
type Model interface {
	Save() error
	Update() error
}

// DB represents database connection
var DB = createDBConnection()

func createDBConnection() *sql.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("We failed to load environment variables")
	}

	connection, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("Establishing database connection failed because:\n%v", err)
	}
	if err = connection.Ping(); err != nil {
		log.Fatalf("Establishing database connection failed because:\n%v", err)
	}

	return connection
}
