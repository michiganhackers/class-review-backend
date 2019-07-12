package main

import (
	"class-review-backend/controllers"
	"class-review-backend/env"
	"class-review-backend/repositories"
	"class-review-backend/services"
	"log"

	"github.com/gin-contrib/cors"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	//cors "github.com/rs/cors/wrapper/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		log.Println("Here!")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	// config := cors.DefaultConfig()
	// config.AddAllowHeaders("ID-Token")
	// config.AllowAllOrigins = true
	config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"*"}
	config.AllowAllOrigins = true
	config.AddAllowHeaders("ID-Token")
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "ID-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		//AllowAllOrigins:  false,
		AllowOriginFunc: func(origin string) bool { return true },
		MaxAge:          86400,
	}))

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env.Init()

	// config := cors.DefaultConfig()
	// config.AddAllowHeaders("ID-Token")
	// config.AllowAllOrigins = true

	db := repositories.CreateDB()
	repos := repositories.DefaultRepositories(db)
	servs := services.DefaultServices(repos)
	controllers.DefaultControllers(r, servs)

	r.Run(":8080")
}
