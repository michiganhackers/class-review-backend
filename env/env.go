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
}

func Init() *EnvironmentVariables {

	// Get environment variabless
	environmentVariables := EnvironmentVariables{
		DBName:     GetEnvironmentVariableByKey("DB_NAME"),
		DBUser:     GetEnvironmentVariableByKey("DB_USER"),
		DBPassword: GetEnvironmentVariableByKey("DB_PASSWORD"),
		DBServer:   GetEnvironmentVariableByKey("DB_SERVER"),
		DBPort:     GetEnvironmentVariableByKey("DB_PORT"),
	}
	return &environmentVariables
}

func GetEnvironmentVariableByKey(key string) string {
	s := os.Getenv(key)
	if s == "" {
		log.Println("Failed to grab environment variable with key", key)
	}
	return s
}
