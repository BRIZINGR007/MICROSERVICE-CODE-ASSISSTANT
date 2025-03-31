package main

import (
	"log"

	"github.com/BRIZINGR007/app-002-code-assistant/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.SetupRoutes(server)
	// Start Server
	if err := server.Run(":3080"); err != nil {
		log.Fatalf("Failed to  start server  : %v", err)
	}

}
