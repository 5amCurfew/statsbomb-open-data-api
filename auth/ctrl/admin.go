package ctrl

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/5amCurfew/statsbomb-open-data-api/auth/lib"
	"github.com/5amCurfew/statsbomb-open-data-api/auth/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// admin/user/:identifier route
func GetUser(c *gin.Context) {
	if c.Param("identifier") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user email or ID required"})
		return
	}

	var data []byte
	var user models.User

	identifier := c.Param("identifier")

	id, err := strconv.Atoi(identifier)

	if err != nil {
		user, err = GetUserByEmail(identifier)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data, _ = json.Marshal(user)

		c.JSON(http.StatusOK, gin.H{"message": "success", "data": json.RawMessage(data)})
		return
	}

	user, err = GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, _ = json.Marshal(user)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": json.RawMessage(data)})
}

// admin/token/:token route
func GetToken(c *gin.Context) {
	if c.Param("token") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token required"})
		return
	}

	token, _ := lib.ParseToken(c.Param("token"))

	var data []byte
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to gather claims"})
		return
	}

	user, err := GetUserByID(int(claims["sub"].(float64)))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, _ = json.Marshal(
		map[string]interface{}{
			"claims":            claims,
			"token":             c.Param("token"),
			"tokenIssuedAt":     time.Unix(int64(claims["iat"].(float64)), 0).Format(time.RFC3339),
			"tokenExpirationAt": time.Unix(int64(claims["exp"].(float64)), 0).Format(time.RFC3339),
			"user":              user,
		},
	)
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": json.RawMessage(data)})
}

// Get first user (by ID)
func GetUserByID(uid int) (models.User, error) {
	var u models.User
	if err := lib.DB.First(&u, uid).Error; err != nil {
		return u, errors.New("user id not found")
	}

	u.ClearPassword()

	return u, nil
}

// Get first user (by username)
func GetUserByEmail(email string) (models.User, error) {
	var u models.User
	if err := lib.DB.Where("email = ?", email).First(&u).Error; err != nil {
		return u, errors.New("user email not found")
	}

	u.ClearPassword()

	return u, nil
}
