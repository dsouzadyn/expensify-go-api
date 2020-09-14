package user

import (
	"net/http"

	"github.com/dsouzadyn/expensify-api/utils"
	"github.com/gin-gonic/gin"
)

// User defines our user's properties
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Auth defines the login properties
type Auth struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateUserHandler handles creating a new user
func CreateUserHandler(c *gin.Context) {
	var json User
	db := utils.DBConn()
	defer db.Close()

	query := "INSERT INTO user (username, email, password) VALUES (?, ?, ?)"

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	passwordHash, err := utils.HashPassword(json.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := stmt.Exec(json.Username, json.Email, passwordHash)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": json.Username,
		"email":    json.Email,
		"id":       id,
	})
}

// AuthenticateUserHandler handles user authentication
func AuthenticateUserHandler(c *gin.Context) {
	var json Auth
	db := utils.DBConn()
	defer db.Close()

	query := "SELECT user_id, password FROM user WHERE username = ? OR email = ?"

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var hashedPassword string
	var userID uint64
	row := db.QueryRow(query, json.Username, json.Username)
	err := row.Scan(&userID, &hashedPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if utils.CheckPasswordHash(json.Password, hashedPassword) {
		token, err := utils.CreateToken(userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"username":    json.Username,
			"accessToken": token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Username or password incorrect",
		})
	}

}
