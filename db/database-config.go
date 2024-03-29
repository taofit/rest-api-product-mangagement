package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func SetDbCon() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	// dbHost := os.Getenv("DB_HOST") // it is not localhost, but the database service name in the docker-compose.yml file
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("DB_PORT")

	dbSqlN := fmt.Sprintf("host=database port=%s user=%s password=%s dbname=%s sslmode=disable", dbPort, dbUser, dbPass, dbName)
	conn, err := sql.Open("postgres", dbSqlN)
	if err != nil {
		return &sql.DB{}, err
	}

	err = conn.Ping()
	if err != nil {
		return &sql.DB{}, err
	}
	log.Println("Database connection established")

	return conn, nil
}
