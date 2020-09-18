package main

import (
	"log"

	"github.com/dsouzadyn/expensify-api/middleware"
	"github.com/dsouzadyn/expensify-api/services/category"
	"github.com/dsouzadyn/expensify-api/services/exchangerate"
	"github.com/dsouzadyn/expensify-api/services/health"
	"github.com/dsouzadyn/expensify-api/services/user"
	"github.com/dsouzadyn/expensify-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// SetupServer initializes our server and returns an instance of it
func SetupServer() *gin.Engine {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.DBMigrate()

	r := gin.Default()
	r.GET("/health", health.Handler)
	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/create", user.CreateUserHandler)
		userRoutes.POST("/authenticate", user.AuthenticateUserHandler)
	}
	exchangeRateRoutes := r.Group("/exchangerate")
	{
		exchangeRateRoutes.POST("/create", middleware.AuthMiddleware(), exchangerate.CreateExchangeRateHandler)
	}
	categoryRoutes := r.Group("/category")
	{
		categoryRoutes.POST("/create", middleware.AuthMiddleware(), category.CreateCategoryHandler)
	}
	return r
}
