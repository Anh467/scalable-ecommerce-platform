package router

import (
	"github.com/Anh467/scalable-ecommerce-platform/server/user-service/internal/delivery/http/middleware"

	"github.com/Anh467/scalable-ecommerce-platform/server/user-service/internal/delivery/http/handler"

	_ "github.com/Anh467/scalable-ecommerce-platform/server/user-service/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.LoggingMiddleware())

	api := r.Group("/api/v1")
	{
		api.POST("/register", userHandler.Register)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}