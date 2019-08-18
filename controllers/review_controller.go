package controllers

import (
	"class-review-backend/models"
	"class-review-backend/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
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
	if err != nil || !bodyIsValid(&reviewInput) {
		log.Println("Invalid request body")
		c.JSON(http.StatusBadRequest, "Invalid request body")
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(reviewInput.UserEmail), 12)
	if err != nil {
		log.Println("POST request failed")
		c.JSON(http.StatusNotFound, "POST request failed")
		return
	}
	reviewInput.UserEmail = string(hash)
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
	if err != nil || !bodyIsValid(&reviewInput) {
		log.Println("Invalid request body")
		c.JSON(http.StatusBadRequest, "Invalid request body")
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(reviewInput.UserEmail), 12)
	if err != nil {
		log.Println("PUT request failed")
		c.JSON(http.StatusNotFound, "PUT request failed")
		return
	}
	reviewInput.UserEmail = string(hash)
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

func bodyIsValid(body *models.Review) bool {
	if body.Rating > 5 || body.Difficulty > 5 || body.Interest > 5 {
		return false
	} else if body.UserEmail == "" {
		return false
	} else if body.Semester != nil &&
			  (len(*body.Semester) < 6 || ((*body.Semester)[:2] != "FA" && (*body.Semester)[:2] != "WN" &&
			  (*body.Semester)[:2] != "SP" && (*body.Semester)[:2] != "SU" && (*body.Semester)[:2] != "SS")) {
		return false
	}

	return true
}