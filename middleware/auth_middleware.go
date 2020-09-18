package middleware

import (
	"net/http"

	"github.com/dsouzadyn/expensify-api/utils"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware is the authentication middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := utils.VerifyToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			// Abort ensures pending handlers are not called
			// Only necessary in middleware
			c.Abort()
			return
		}
		c.Next()
	}
}
