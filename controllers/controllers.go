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
		IDToken := c.GetHeader("ID-Token")
		err := authenticate(IDToken)
		if err != nil {
			log.Println("could not authenticate id token -- " + err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "could not authenticate id token -- " + err.Error()})
			c.Abort()
			return
		}
		// add token to cache, set it to expire after an 30 mins (I'm basically just using this as a set)
		// I went with empty string rather than nil to avoid confusion b/c .Get() returns nil if the key isn't found
		tokenCache.Add(IDToken, "", 30*time.Minute)
		c.Next()
	}
}

func DefaultControllers(r *gin.Engine, services *services.Services) *Controllers {
	// create a new cache with default expiration 30 mins and cleanup time 30 mins
	tokenCache = cache.New(30*time.Minute, 30*time.Minute)
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
