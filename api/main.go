package main

import (
	"net/http"

	"github.com/5amCurfew/statsbomb-open-data/api/ctrl"
	"github.com/5amCurfew/statsbomb-open-data/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error parsing .env")
	}

	log.SetFormatter(&log.JSONFormatter{})

	r := gin.Default()

	public := r.Group("/")
	public.Use(middleware.Public())
	public.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong 🏓"})
	})

	api := r.Group("/api")
	api.Use(middleware.Api())
	api.GET("/competitions/", ctrl.GetCompetitions)
	api.GET("/competition/:competition_id/:season_id", ctrl.GetCompetition)
	api.GET("/lineups/:match_id", ctrl.GetLineUps)

	r.Run(":8080")
}
