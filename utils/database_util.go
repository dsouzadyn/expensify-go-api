package utils

import (
	"log"
	"os"

	"github.com/dsouzadyn/expensify-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBMigrate is a database migration helper function
func DBMigrate() {
	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_DB_DSN")), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// Migrate the models
	db.AutoMigrate(&models.User{}, &models.Budget{}, &models.Expense{}, &models.Category{}, &models.ExchangeRate{})
}

// DBConn returns a database connection instance
func DBConn() *gorm.DB {
	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_DB_DSN")), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
