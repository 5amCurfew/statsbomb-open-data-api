package ctrl

import (
	"fmt"
	"net/http"

	"github.com/5amCurfew/statsbomb-open-data/api/lib"
	"github.com/gin-gonic/gin"
)

func GetLineUps(c *gin.Context) {
	if c.Param("match_id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "match id is required"})
		return
	}

	lib.Local(c, fmt.Sprintf("lineups/%s.json", c.Param("match_id")))
}
