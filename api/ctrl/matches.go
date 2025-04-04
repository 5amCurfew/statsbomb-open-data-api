package ctrl

import (
	"fmt"
	"net/http"

	"github.com/5amCurfew/statsbomb-open-data/api/lib"
	"github.com/gin-gonic/gin"
)

func GetMatches(c *gin.Context) {
	if c.Param("competition_id") == "" || c.Param("season_id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "competition and season ids are required"})
		return
	}

	lib.GCPStorage(c, fmt.Sprintf("matches/%s/%s.json", c.Param("competition_id"), c.Param("season_id")))
}
