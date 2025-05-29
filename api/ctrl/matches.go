package ctrl

import (
	"fmt"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/5amCurfew/statsbomb-open-data/api/lib"
	"github.com/5amCurfew/statsbomb-open-data/api/models"
	"github.com/gin-gonic/gin"
)

func GetMatches(c *gin.Context) {
	if c.Param("competition_id") == "" || c.Param("season_id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "competition and season ids are required"})
		return
	}

	gcsClient, exists := c.Get("gcsClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "GCS client not found in context"})
		return
	}
	client, _ := gcsClient.(*storage.Client)

	// Read from GCS
	path := fmt.Sprintf("matches/%s/%s.json", c.Param("competition_id"), c.Param("season_id"))
	data, err := lib.ReadGCSFile(path, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Parse into []Match model
	matches, err := models.ParseMatches(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid JSON structure for match"})
		return
	}

	c.JSON(http.StatusOK, matches)
}
