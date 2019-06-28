package repositories

import (
	"class-review-backend/env"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func CreateDB() *sqlx.DB {
	environmentVariables := env.Init()
	connection := environmentVariables.DBUser + ":" + environmentVariables.DBPassword + "@tcp(" + environmentVariables.DBServer + ":" + environmentVariables.DBPort + ")/" + environmentVariables.DBName+"?allowNativePasswords=True"
	db, err := sql.Open("mysql", connection)
	if err != nil {
		fmt.Println("Failed to open database: ", err.Error())
	}
	// Create wrapper for Go db instance
	dbx := sqlx.NewDb(db, "mysql")
	if err != nil {
		fmt.Println(err)
		return dbx
	}
	return dbx
}
