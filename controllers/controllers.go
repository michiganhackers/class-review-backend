package controllers

import (
	"michiganhackers/class-review-backend/services"

	"github.com/gin-gonic/gin"
)

type Controllers struct {
	ReviewController *ReviewController
}

func DefaultControllers(r *gin.Engine, services *services.Services) *Controllers {
	controllers := &Controllers{
		ReviewController: DefaultReviewController(r, services),
	}
	return controllers
}
