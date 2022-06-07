package main

import (
	"github.com/bboygg/mattermost-webhook-qase/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := gin.Default()
	r.Use(InitMiddleware())

	router := r.Group("/webhook")

	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": true})
	})

	// Initialize routes
	routes.Init(router)

	r.Run()
}

func InitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-type", "application/json")
		c.Next()
	}
}
