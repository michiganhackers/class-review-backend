package controllers

import (
	"class-review-backend/services"
	"log"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"

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
}

var tokenCache *cache.Cache

func AuthenticationRequired(auths ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id_token := c.GetHeader("ID-Token")
		_, err := authenticate(id_token)
		if err != nil {
			log.Println("could not authenticate id token -- " + err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "could not authenticate id token -- " + err.Error()})
			c.Abort()
			return
		}
		// add token to cache, set it to expire after an hour (I'm basically just using this as a list)
		// should I be using their id or their id_token? Does it make a difference?
		tokenCache.Add(id_token, "dummy", time.Hour)
		c.Next()
	}
}

func DefaultControllers(r *gin.Engine, services *services.Services) *Controllers {
	// create a new cache with default expiration 1 hour and cleanup time 3 hours
	tokenCache = cache.New(time.Hour, 3*time.Hour)
	routes := &Routes{
		Private: r.Group("/"),
		Public:  r.Group("/"),
	}

	routes.Private.Use(AuthenticationRequired())

	controllers := &Controllers{
		ProfessorController: DefaultProfessorController(routes, services),
		ReviewController:    DefaultReviewController(routes, services),
		CourseController:    DefaultCourseController(routes, services),
	}
	return controllers
}
