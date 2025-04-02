package ctrl

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/5amCurfew/statsbomb-open-data-api/auth/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Login input model
type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// auth/login route
func AuthLogin(c *gin.Context) {
	var input LoginInput
	var token string
	var tokenErr error

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{
		Email:    input.Email,
		Password: input.Password,
	}

	if pass, _ := u.Login(); !pass {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		return
	}

	token, tokenErr = generateToken(u)
	if tokenErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("failed to generate token %v", tokenErr)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Generate token for a given User
func generateToken(u models.User) (string, error) {
	lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"iss": "5am",
		"sub": u.ID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * time.Duration(lifespan)).Unix(),
		"adm": u.IsAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}
