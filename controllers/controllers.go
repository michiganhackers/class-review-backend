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

func DefaultControllers(r *gin.Engine, services *services.Services) *Controllers {
	controllers := &Controllers{
		ReviewController:    DefaultReviewController(r, services),
		CourseController:    DefaultCourseController(r, services),
		ProfessorController: DefaultProfessorController(r, services),
	}
	return controllers
}
