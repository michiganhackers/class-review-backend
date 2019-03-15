package controllers

import (
	"class-review-backend/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewController struct {
	Services *services.Services
	Routes   *gin.RouterGroup
}

func DefaultReviewController(eng *gin.Engine, services *services.Services) *ReviewController {
	rc := &ReviewController{
		Routes:   eng.Group("/review"),
		Services: services,
	}

	rc.Routes.GET("/review/:id", rc.getReview)
	return rc
}

func (rc *ReviewController) getReview(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		log.Println("No id in url")
		c.JSON(http.StatusBadRequest, "No id in url")
		return
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Println("Bad id param")
		c.JSON(http.StatusBadRequest, "Bad id param")
		return
	}
	review, err := rc.Services.ReviewService.GetReview(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "No review with provided id")
		return
	}
	c.JSON(http.StatusOK, review)
	return
}
