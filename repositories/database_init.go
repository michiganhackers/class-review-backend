package repositories

import (
	"class-review-backend/env"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

func CreateDB() *sqlx.DB {
	connection := env.Variables.DBUser + ":" + env.Variables.DBPassword + "@tcp(" + env.Variables.DBServer + ":" + env.Variables.DBPort + ")/" + env.Variables.DBName
	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Println("Failed to open database: ", err.Error())
	}
	// Create wrapper for Go db instance
	dbx := sqlx.NewDb(db, "mysql")
	if err != nil {
		log.Println(err)
		return dbx
	}
	return dbx
}
