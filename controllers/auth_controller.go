package controllers

import (
	"net/http"
	"task-manager/database"
	"task-manager/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *gin.Context) {

	var credential models.Credential
	var user models.User
	var jwtKey = []byte("my_secret_key")

	c.BindJSON(&credential)

	err := database.DB.Where("email", credential.Email).First(&user).Error

	if err != nil {
		c.JSON(404, gin.H{
			"sucess": "falied",
			"error":  err,
		})
		return
	}

	if user.Password != credential.Password {
		c.JSON(404, gin.H{
			"sucess":  "falied",
			"message": "password is incorrect!",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": credential.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "falied",
			"error":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  tokenString,
	})

}
