package main

import (
	"class-review-backend/controllers"
	"class-review-backend/env"
	"class-review-backend/repositories"
	"class-review-backend/services"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	env.Init()
	r.Use(controllers.CORS())

	db := repositories.CreateDB()
	repos := repositories.DefaultRepositories(db)
	servs := services.DefaultServices(repos)
	controllers.DefaultControllers(r, servs)

	r.Run(":8080")
}
