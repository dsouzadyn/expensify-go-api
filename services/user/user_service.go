package user

import (
	"net/http"

	"github.com/dsouzadyn/expensify-api/models"
	"github.com/dsouzadyn/expensify-api/utils"
	"github.com/gin-gonic/gin"
)

// Auth defines the login properties
type Auth struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateUserHandler handles creating a new user
func CreateUserHandler(c *gin.Context) {
	var user models.User
	db := utils.DBConn()

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": user.ID,
	})
}

// AuthenticateUserHandler handles user authentication
func AuthenticateUserHandler(c *gin.Context) {
	var authUser Auth
	var user models.User
	db := utils.DBConn()

	if err := c.ShouldBindJSON(&authUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := db.Where("username = ? OR email = ?", authUser.Username, authUser.Username).First(&user)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No user exists with those credentials",
		})
		return
	}

	if models.CheckPasswordHash(authUser.Password, user.Password) {
		token, err := utils.CreateToken(user.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"accessToken": token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Username or password incorrect",
		})
	}

}
