package ctrl

import (
	"fmt"
	"net/http"

	"github.com/5amCurfew/statsbomb-open-data/api/lib"
	"github.com/gin-gonic/gin"
)

func GetCompetitions(c *gin.Context) {
	lib.GCPStorage(c, "competitions/all.json")
}

func GetCompetition(c *gin.Context) {
	if c.Param("competition_id") == "" || c.Param("season_id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "competition and season ids are required"})
		return
	}

	lib.GCPStorage(c, fmt.Sprintf("competitions/%s_%s.json", c.Param("competition_id"), c.Param("season_id")))
}
