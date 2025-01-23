package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sudhi001/Y2Device/handlers"
	"github.com/sudhi001/Y2Device/services"
)

func main() {
	// Initialize services
	ytService := services.NewYouTubeService()

	// Initialize handlers
	videoHandler := handlers.NewVideoHandler(ytService)

	// Set up the router
	r := gin.Default()

	// Define routes
	r.GET("/download", videoHandler.DownloadVideo)

	// Run the server
	r.Run(":8080")
}

