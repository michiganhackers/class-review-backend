package main

import (
	"class-review-backend/controllers"
	"class-review-backend/env"
	"class-review-backend/repositories"
	"class-review-backend/services"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	environmentVariables := env.Init()

	db := repositories.CreateDB(environmentVariables)
	repos := repositories.DefaultRepositories(db)
	servs := services.DefaultServices(repos)
	controllers.DefaultControllers(r, servs, environmentVariables)

	r.Run(":8080")
}
