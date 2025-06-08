package main

import (
	"net/http"

	"github.com/5amCurfew/statsbomb-open-data-api/auth/ctrl"
	"github.com/5amCurfew/statsbomb-open-data-api/auth/lib"
	"github.com/5amCurfew/statsbomb-open-data-api/auth/middleware"
	"github.com/5amCurfew/statsbomb-open-data-api/auth/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	lib.ConnectToAuthDatabase()
	lib.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	// public routes
	public := r.Group("/")
	public.Use(middleware.PublicMiddleware())
	public.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong 🏓"})
	})

	// auth routes
	auth := r.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	auth.POST("/register", ctrl.AuthRegister)
	auth.POST("/login", ctrl.AuthLogin)

	// admin routes
	admin := r.Group("/admin")
	admin.Use(middleware.AdminMiddleware())
	admin.GET("/user/:identifier", ctrl.GetUser)
	admin.GET("/token/:token", ctrl.GetToken)

	r.Run(":4000")
}
