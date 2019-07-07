package controllers

import (
	"class-review-backend/env"
	"class-review-backend/services"

	"github.com/gin-gonic/gin"
)

type Controllers struct {
	ReviewController     *ReviewController
	CourseController     *CourseController
	EnvironmentVariables *env.EnvironmentVariables
}

func DefaultControllers(r *gin.Engine, services *services.Services, environmentVariables *env.EnvironmentVariables) *Controllers {
	controllers := &Controllers{
		ReviewController:     DefaultReviewController(r, services),
		CourseController:     DefaultCourseController(r, services),
		EnvironmentVariables: environmentVariables,
	}
	return controllers
}
