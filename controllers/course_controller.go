package controllers

import (
	"class-review-backend/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CourseController struct {
	Services *services.Services
	Routes   *Routes
}

func DefaultCourseController(routes *Routes, services *services.Services) *CourseController {
	cc := &CourseController{
		Routes:   routes,
		Services: services,
	}

	cc.Routes.Public.GET("/course/:id", cc.getCourse)
	return cc
}

func (cc *CourseController) getCourse(c *gin.Context) {
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
	course, err := cc.Services.CourseService.GetCourse(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "No course with provided id")
		return
	}
	c.JSON(http.StatusOK, course)
	return
}
