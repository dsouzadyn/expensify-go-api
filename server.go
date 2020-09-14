package main

import (
	"log"
	"os"

	"github.com/dsouzadyn/expensify-api/services/health"
	"github.com/dsouzadyn/expensify-api/services/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// SetupServer initializes our server and returns an instance of it
func SetupServer() *gin.Engine {
	_, present := os.LookupEnv("GOTEST")

	if !present {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	r := gin.Default()
	r.GET("/health", health.Handler)
	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/create", user.CreateUserHandler)
		userRoutes.POST("/authenticate", user.AuthenticateUserHandler)
	}
	return r
}
