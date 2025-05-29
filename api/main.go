package main

import (
	"context"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/5amCurfew/statsbomb-open-data/api/ctrl"
	"github.com/5amCurfew/statsbomb-open-data/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error parsing .env")
	}

	log.SetFormatter(&log.JSONFormatter{})

	r := gin.Default()

	// Create a GCS client
	client, err := storage.NewClient(context.Background(), option.WithCredentialsFile("service-account.json"))
	if err != nil {
		log.Warnf("error: failed to create GCS client: %v", err)
		panic("error connecting to GCS")
	}
	defer client.Close()

	public := r.Group("/")
	public.Use(middleware.Public())
	public.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong 🏓"})
	})

	api := r.Group("/api")
	api.Use(middleware.Api(client))
	api.GET("/competitions/", ctrl.GetCompetitions)
	api.GET("/competition/:competition_id/:season_id", ctrl.GetCompetition)
	//api.GET("/events/:match_id", ctrl.GetEvents)
	api.GET("/lineups/:match_id", ctrl.GetLineUps)
	api.GET("/matches/:competition_id/:season_id", ctrl.GetMatches)

	r.Run(":8080")
}
