package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(server *gin.Engine) {
	// Health-Check Router ...
	server.GET("/healthcheck", func(context *gin.Context) {
		context.JSON(200, gin.H{"status": "OK"})
	})
	// Public Routes
	
}
