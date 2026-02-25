package main

import (
	"log"
	"os"

	"github.com/Anh467/scalable-ecommerce-platform/user-service/internal/application/service"
	"github.com/Anh467/scalable-ecommerce-platform/user-service/internal/delivery/http/handler"
	"github.com/Anh467/scalable-ecommerce-platform/user-service/internal/delivery/http/router"
	"github.com/Anh467/scalable-ecommerce-platform/user-service/internal/infrastructure/persistence/postgresql"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

// @title Ecommerce API
// @version 1.0
// @description Scalable Ecommerce API
// @host localhost:8080
// @BasePath /api/v1
func main() {

	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//connStringMSSQL := os.Getenv("DB_CONN_STRING_MSSQL")
	connStringPostGreSQL := os.Getenv("DB_CONN_STRING_POSTGRESQL")

	//dbMSSQL:= mssql.NewMSSQL(connStringMSSQL)
	dbPostGreSQL:= postgresql.NewPOSTGRESQL(connStringPostGreSQL)

	// Repository
	userRepo := postgresql.NewUserRepository(dbPostGreSQL)

	// Service
	userService := service.NewUserService(userRepo)

	// Handler
	userHandler := handler.NewUserHandler(userService)

	// Router
	r := router.SetupRouter(userHandler)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	r.Run(":" + port)
}