package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func runMigrations(dbURL string) {
    m, err := migrate.New(
        "file://migrations",
        dbURL,
    )
    if err != nil {
        log.Fatal(err)
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatal(err)
    }
}

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//connStringMSSQL := os.Getenv("DB_CONN_STRING_MSSQL")
	connStringPostGreSQL := os.Getenv("DB_CONN_STRING_POSTGRESQL")
	
    runMigrations(connStringPostGreSQL)
}