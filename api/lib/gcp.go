package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

const (
	bucketName = "statsbomb-open-data-api-data"
)

func GCPStorage(c *gin.Context, path string) {
	filePath := path
	gcsClient, _ := c.Get("gcsClient")
	client, _ := gcsClient.(*storage.Client)

	// Read the file from GCS
	rc, err := client.Bucket(bucketName).Object(filePath).NewReader(context.Background())
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
