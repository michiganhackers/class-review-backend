package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentVariables struct {
	DBName       string
	DBUser       string
	DBPassword   string
	DBServer     string
	DBPort       string
	ClientOrigin string
	ClientId     string
}

var Variables *EnvironmentVariables

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get environment variables
	Variables = &EnvironmentVariables{
		DBName:       getEnvironmentVariableByKey("DB_NAME"),
		DBUser:       getEnvironmentVariableByKey("DB_USER"),
		DBPassword:   getEnvironmentVariableByKey("DB_PASSWORD"),
		DBServer:     getEnvironmentVariableByKey("DB_SERVER"),
		DBPort:       getEnvironmentVariableByKey("DB_PORT"),
		ClientOrigin: getEnvironmentVariableByKey("CLIENT_ORIGIN"),
		ClientId:     getEnvironmentVariableByKey("CLIENT_ID"),
	}
}

func getEnvironmentVariableByKey(key string) string {
	s := os.Getenv(key)
	if s == "" {
		log.Println("Failed to grab environment variable with key", key)
	}
	return s
}
