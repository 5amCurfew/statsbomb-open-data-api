package ctrl

import (
	"fmt"
	"net/http"

	"github.com/5amCurfew/statsbomb-open-data/api/lib"
	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	if c.Param("match_id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "match id is required"})
		return
	}

	lib.GCPStorage(c, fmt.Sprintf("events/%s.json", c.Param("match_id")))
}
