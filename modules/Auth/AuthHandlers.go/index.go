package AuthHandlers

import (
	"log"
	"net/http"
	"todo/models"
	"todo/modules/Auth/AuthService"

	"github.com/gin-gonic/gin"
)

type SigninRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	user         models.User
	access_token string
}

func Signin(c *gin.Context) {
	var body SigninRequest
	var user models.User
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err, token := AuthService.Login(&user, body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"access_token": token,
			"user":         user,
		})
	}
	log.Println(user)
}

func Signup(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err, token := AuthService.Signup(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"access_token": token,
			"user":         user,
		})
	}
	log.Println(user)
}
