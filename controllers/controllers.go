package controllers

import (
	"class-review-backend/services"

	"github.com/gin-gonic/gin"
)

type Controllers struct {
	ReviewController    *ReviewController
	CourseController    *CourseController
	ProfessorController *ProfessorController
}

type Routes struct {
	Private *gin.RouterGroup
	Public  *gin.RouterGroup
	Admin   *gin.RouterGroup
}

func AuthenticationRequired(auths ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
	}
}

func DefaultControllers(r *gin.Engine, services *services.Services) *Controllers {

	routes := &Routes{
		Private: r.Group("/"),
		Public:  r.Group("/"),
		Admin:   r.Group("/admin"),
	}

	routes.Private.Use(AuthenticationRequired("private", "admin"))

	controllers := &Controllers{
		ReviewController:    DefaultReviewController(r, services),
		CourseController:    DefaultCourseController(r, services),
		ProfessorController: DefaultProfessorController(r, services),
	}
	return controllers
}
