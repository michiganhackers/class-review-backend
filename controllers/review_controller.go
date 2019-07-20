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
	Routes   *Routes
}

func DefaultReviewController(routes *Routes, services *services.Services) *ReviewController {
	rc := &ReviewController{
		Routes:   routes,
		Services: services,
	}

	rc.Routes.Public.GET("/review/:id", rc.getReview)
	return rc

}

func (rc *ReviewController) getReview(c *gin.Context) {
	// is this where I should check if the user's token is valid?
	id_token := c.GetHeader("ID-Token")
	_, valid := tokenCache.Get(id_token)
	if !valid {
		log.Println("User token invalid")
		c.JSON(http.StatusForbidden, "User token invalid")
	}

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
