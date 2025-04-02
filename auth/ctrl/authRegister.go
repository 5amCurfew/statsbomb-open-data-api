package ctrl

import (
	"net/http"

	"github.com/5amCurfew/statsbomb-open-data-api/auth/models"
	"github.com/gin-gonic/gin"
)

// Register input model
type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// auth/register route
func AuthRegister(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	candidate := models.User{
		Email:    input.Email,
		Password: input.Password,
	}

	_, err := candidate.Register()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration successful"})
}
