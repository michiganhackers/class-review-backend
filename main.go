package main

import (
	"class-review-backend/controllers"
	"class-review-backend/env"
	"class-review-backend/repositories"
	"class-review-backend/services"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// TODO: including but not limited to: environment variables, db init, auth
func main() {
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// TODO: remove this
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	environmentVariables := env.Init()
	db := repositories.CreateDB(environmentVariables)
	repos := repositories.DefaultRepositories(db)
	servs := services.DefaultServices(repos)
	controllers.DefaultControllers(r, servs)

	repos.ReviewRepository.TestDB()
	r.Run(":8080")
}
