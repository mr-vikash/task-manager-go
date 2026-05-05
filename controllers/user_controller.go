package controllers

import (
	"task-manager/database"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	database.DB.Create(&user)
	c.JSON(200, user)
}

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Preload("Tasks").Find(&users)
	c.JSON(200, users)
}
