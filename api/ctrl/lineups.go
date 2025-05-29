package ctrl

import (
	"fmt"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/5amCurfew/statsbomb-open-data/api/lib"
	"github.com/5amCurfew/statsbomb-open-data/api/models"
	"github.com/gin-gonic/gin"
)

func GetLineUps(c *gin.Context) {
	if c.Param("match_id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "match id is required"})
		return
	}

	gcsClient, exists := c.Get("gcsClient")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "GCS client not found in context"})
		return
	}
	client, _ := gcsClient.(*storage.Client)

	// Read from GCS
	path := fmt.Sprintf("lineups/%s.json", c.Param("match_id"))
	data, err := lib.ReadGCSFile(path, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Parse into []LineUp model
	competition, err := models.ParseLineUps(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid JSON structure for line up"})
		return
	}

	c.JSON(http.StatusOK, competition)
}
