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
}

func Init() *EnvironmentVariables {

	// set config variables
	environmentVariables := EnvironmentVariables{
		DBName:     GetEnvironmentVariableByKey("DB_NAME"),
		DBUser:     GetEnvironmentVariableByKey("DB_USER"),
		DBPassword: GetEnvironmentVariableByKey("DB_PASSWORD"),
		DBServer:   GetEnvironmentVariableByKey("DB_SERVER"),
	}
	return &environmentVariables
}

func GetEnvironmentVariableByKey(key string) string {
	s := os.Getenv(key)
	if s == "" {
		log.Println("Failed to grab environment variable with key ", key)
	}
	return s
}
