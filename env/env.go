package env

import (
	"log"
	"os"
)

type EnvironmentVariables struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBServer   string
	DBPort     string
	ClientId   string
}

func Init() *EnvironmentVariables {

	// Get environment variables
	environmentVariables := EnvironmentVariables{
		DBName:     getEnvironmentVariableByKey("DB_NAME"),
		DBUser:     getEnvironmentVariableByKey("DB_USER"),
		DBPassword: getEnvironmentVariableByKey("DB_PASSWORD"),
		DBServer:   getEnvironmentVariableByKey("DB_SERVER"),
		DBPort:     getEnvironmentVariableByKey("DB_PORT"),
		ClientId:   getEnvironmentVariableByKey("CLIENT_ID"),
	}
	return &environmentVariables
}

func getEnvironmentVariableByKey(key string) string {
	s := os.Getenv(key)
	if s == "" {
		log.Println("Failed to grab environment variable with key", key)
	}
	return s
}
