package mssql

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func NewMSSQL(connString string) *sql.DB {
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database ping error:", err)
	}

	return db
}