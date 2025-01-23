package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sudhi001/Y2Device/services"
)

type VideoHandler struct {
	ytService *services.YouTubeService
}

func NewVideoHandler(ytService *services.YouTubeService) *VideoHandler {
	return &VideoHandler{ytService: ytService}
}

func (h *VideoHandler) DownloadVideo(c *gin.Context) {
	// Get the video URL from the query parameters
	videoURL := c.Query("url")
	if videoURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "YouTube URL is required"})
		return
	}

	// Download the video using the service
	filePath, title, err := h.ytService.DownloadVideo(videoURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to download video", "details": err.Error()})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, gin.H{
		"message":    "Download successful",
		"video_title": title,
		"file_path":   filePath,
	})
}
