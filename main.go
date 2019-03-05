package main

import (
	"michiganhackers/class-review-backend/controllers"
	"michiganhackers/class-review-backend/repositories"
	"michiganhackers/class-review-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: including but not limited to: environment variables, db init, auth
func main() {
	r := gin.Default()
	// TODO: remove this
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	repos := repositories.DefaultRepositories()
	servs := services.DefaultServices(repos)
	controllers.DefaultControllers(r, servs)

	r.Run(":8080")
}
