package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Local(c *gin.Context, path string) {

	filePath := fmt.Sprintf("../data/%s", path)

	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to read data %v", err)})
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
