package services

import (
	"fmt"
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

type YouTubeService struct {
	client *youtube.Client
}

func NewYouTubeService() *YouTubeService {
	return &YouTubeService{
		client: &youtube.Client{},
	}
}

func (s *YouTubeService) DownloadVideo(videoURL string) (string, string, error) {
	// Get video metadata
	video, err := s.client.GetVideo(videoURL)
	if err != nil {
		return "", "", fmt.Errorf("failed to fetch video metadata: %w", err)
	}

	// Find the highest-quality format
	var highestQuality youtube.Format
	for _, format := range video.Formats {
		if format.QualityLabel != "" && format.MimeType[:5] == "video" {
			highestQuality = format
			break
		}
	}

	if highestQuality.QualityLabel == "" {
		return "", "", fmt.Errorf("no suitable video format found")
	}

	// Get the video stream
	stream, _, err := s.client.GetStream(video, &highestQuality)
	if err != nil {
		return "", "", fmt.Errorf("failed to get video stream: %w", err)
	}
	defer stream.Close()

	// Create the output file
	fileName := fmt.Sprintf("downloads/%s.mp4", sanitizeFileName(video.Title))
	outputFile, err := os.Create(fileName)
	if err != nil {
		return "", "", fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	// Write the stream to the file
	_, err = io.Copy(outputFile, stream)
	if err != nil {
		return "", "", fmt.Errorf("failed to download video: %w", err)
	}

	return fileName, video.Title, nil
}

// Helper function to sanitize the file name
func sanitizeFileName(name string) string {
	return name // Add proper sanitization here if needed (e.g., replacing invalid characters).
}
