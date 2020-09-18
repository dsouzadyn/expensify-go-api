package category

import (
	"net/http"

	"github.com/dsouzadyn/expensify-api/models"
	"github.com/dsouzadyn/expensify-api/utils"
	"github.com/gin-gonic/gin"
)

// CreateCategoryHandler handles category creation
func CreateCategoryHandler(c *gin.Context) {
	var category models.Category
	db := utils.DBConn()

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := db.Create(&category)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": category.ID,
	})
}
