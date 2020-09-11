package utils

import (
	"database/sql"
	"log"
	"os"

	// MySQL Driver
	_ "github.com/go-sql-driver/mysql"
)

// DBConn returns a database connection instance
func DBConn() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DB_DSN"))
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
