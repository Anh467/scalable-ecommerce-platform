package postgresql

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPOSTGRESQL(connString string) *sql.DB {
	db, err := sql.Open("pgx", connString)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database ping error:", err)
	}

	return db
}