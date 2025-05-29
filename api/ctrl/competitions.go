package ctrl

import (
	"fmt"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/5amCurfew/statsbomb-open-data/api/lib"
	"github.com/5amCurfew/statsbomb-open-data/api/models"
	"github.com/gin-gonic/gin"
)

func GetCompetitions(c *gin.Context) {
	gcsClient, exists := c.Get("gcsClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "GCS client not found in context"})
		return
	}
	client, _ := gcsClient.(*storage.Client)

	// Read from GCS
	data, err := lib.ReadGCSFile("competitions/all.json", client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Parse into []Competition model
	competition, err := models.ParseCompetitions(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid JSON structure for competitions"})
		return
	}

	c.JSON(http.StatusOK, competition)
}

func GetCompetition(c *gin.Context) {
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
	path := fmt.Sprintf("competitions/%s_%s.json", c.Param("competition_id"), c.Param("season_id"))
	data, err := lib.ReadGCSFile(path, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Parse into Competition model
	competition, err := models.ParseCompetition(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid JSON structure for competition"})
		return
	}

	c.JSON(http.StatusOK, competition)
}
