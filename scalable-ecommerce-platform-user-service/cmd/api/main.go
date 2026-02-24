package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Anh467/scalable-ecommerce-platform-user-service/internal/application/service"
	"github.com/Anh467/scalable-ecommerce-platform-user-service/internal/delivery/http/handler"
	"github.com/Anh467/scalable-ecommerce-platform-user-service/internal/delivery/http/router"
	"github.com/Anh467/scalable-ecommerce-platform-user-service/internal/infrastructure"

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

	connString := os.Getenv("DB_CONN_STRING")

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Database ping error:", err)
	}

	// Repository
	userRepo := infrastructure.NewUserRepository(db)

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