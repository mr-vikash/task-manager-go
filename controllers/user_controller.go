package controllers

import (
	"task-manager/database"
	"task-manager/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var user models.User

	// ✅ Bind & validate JSON
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// ✅ Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	user.Password = string(hashedPassword)

	// ✅ Save to DB
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// ✅ Success response
	c.JSON(201, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Preload("Tasks").Find(&users)
	c.JSON(200, users)
}
