package controllers

import (
	"class-review-backend/services"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
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
		endpointPath := c.FullPath()
		endpointVerb := c.GetHeader("method")
		IDToken := c.GetHeader("ID-Token")
		// TODO: How are we going to get a user's uniqname given their ID token?
		uniqname, err := authenticate(IDToken)
		// authenticate user
		if err != nil {
			log.Println("could not authenticate id token -- " + err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "could not authenticate id token -- " + err.Error()})
			c.Abort()
			return
		}
		// if the endpoint requires admin access
		// TODO: How to best determine if an endpoint requires admin
		if requiresAdmin(endpointPath) && !IsAdmin(uniqname) {
			log.Println("user requires admin privileges " + err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Endpoint requires admin privileges"})
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
