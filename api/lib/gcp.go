package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

const (
	bucketName = "statsbomb-open-data-api-data"
)

func GCPStorage(c *gin.Context, path string) {
	// Build the GCS file path (no ".json" extension is added here, it's assumed the path already includes the file extension)
	filePath := path

	// Create a GCS client
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("service-account.json")) // Provide your service account JSON file
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create GCS client: %v", err)})
		return
	}
	defer client.Close()

	// Read the file from GCS
	rc, err := client.Bucket(bucketName).Object(filePath).NewReader(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to read file from GCS: %v", err)})
		return
	}
	defer rc.Close()

	// Read the content into a byte slice
	data, err := io.ReadAll(rc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to read file content: %v", err)})
		return
	}

	// Parse JSON to ensure it's valid
	var jsonData interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Return the JSON data
	c.JSON(http.StatusOK, jsonData)
}
