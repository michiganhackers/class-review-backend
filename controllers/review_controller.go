package controllers

import (
	"class-review-backend/models"
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

	rc.Routes.Public.GET("/review", rc.getAllReviews)
	rc.Routes.Public.GET("/review/:id", rc.getReviewById)
	rc.Routes.Private.POST("/review", rc.postReview)
	rc.Routes.Private.PUT("/review/:id", rc.updateReview)
	rc.Routes.Private.DELETE("/review/:id", rc.deleteReview)
	return rc

}

func (rc *ReviewController) getAllReviews(c *gin.Context) {
	reviews, err := rc.Services.ReviewService.GetAllReviews()
	if err != nil {
		log.Println("Reviews not found")
		c.JSON(http.StatusNotFound, "Reviews not found")
		return
	}
	c.JSON(http.StatusOK, reviews)
	return
}

func (rc *ReviewController) getReviewById(c *gin.Context) {
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
	review, err := rc.Services.ReviewService.GetReviewById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "No review with provided id")
		return
	}
	c.JSON(http.StatusOK, review)
	return
}

func (rc *ReviewController) postReview(c *gin.Context) {
	var reviewInput models.Review
	err := c.BindJSON(&reviewInput)
	// TODO: Error checking for body
	if err != nil {
		log.Println("Invalid request body")
		c.JSON(http.StatusBadRequest, "Invalid request body")
		return
	}
	err = rc.Services.ReviewService.PostReview(&reviewInput)
	if err != nil {
		log.Println("POST request failed")
		c.JSON(http.StatusNotFound, "POST request failed")
		return
	}
	c.JSON(http.StatusOK, "OK")
	return
}

func (rc *ReviewController) updateReview(c *gin.Context) {
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
	var reviewInput models.Review
	err = c.BindJSON(&reviewInput)
	// TODO: Error checking for body
	if err != nil {
		log.Println("Invalid request body")
		c.JSON(http.StatusBadRequest, "Invalid request body")
		return
	}
	review, err := rc.Services.ReviewService.UpdateReview(&reviewInput, id)
	if err != nil {
		log.Println("PUT request failed")
		c.JSON(http.StatusNotFound, "PUT request failed")
		return
	}
	c.JSON(http.StatusOK, review)
	return
}

func (rc *ReviewController) deleteReview(c *gin.Context) {
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
	err = rc.Services.ReviewService.DeleteReview(id)
	if err != nil {
		log.Println("DELETE request failed")
		c.JSON(http.StatusNotFound, "DELETE request failed")
		return
	}
	c.JSON(http.StatusOK, "OK")
	return
}