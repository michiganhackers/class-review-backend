package repositories

import (
	"class-review-backend/env"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

func CreateDB() *sqlx.DB {
	environmentVariables := env.Init()
	connection := environmentVariables.DBUser + ":" + environmentVariables.DBPassword + "@" + environmentVariables.DBServer + "/" + environmentVariables.DBName
	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Println("Failed to open database: ", err.Error())
	}
	// Create wrapper for Go db instance
	dbx := sqlx.NewDb(db, "mysql")
	return dbx
}
