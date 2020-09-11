package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler handles the main route for the Health service
func Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Status OK",
	})
}
