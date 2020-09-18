package exchangerate

import (
	"net/http"

	"github.com/dsouzadyn/expensify-api/models"
	"github.com/dsouzadyn/expensify-api/utils"
	"github.com/gin-gonic/gin"
)

func CreateExchangeRateHandler(c *gin.Context) {
	var exchangeRate models.ExchangeRate
	db := utils.DBConn()

	if err := c.ShouldBindJSON(&exchangeRate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := db.Create(&exchangeRate)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": exchangeRate.ID,
	})
}
