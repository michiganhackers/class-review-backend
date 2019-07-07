package repositories

import (
	"class-review-backend/env"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

func CreateDB(environmentVariables *env.EnvironmentVariables) *sqlx.DB {
	connection := environmentVariables.DBUser + ":" + environmentVariables.DBPassword + "@tcp(" + environmentVariables.DBServer + ":" + environmentVariables.DBPort + ")/" + environmentVariables.DBName
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
