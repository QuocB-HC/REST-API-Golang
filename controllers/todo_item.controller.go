package controllers

import (
	"github.com/QuocB-HC/REST-API-Golang/config"
	"github.com/QuocB-HC/REST-API-Golang/models"
	"github.com/gin-gonic/gin"
)


func GetAllItem(c *gin.Context) {
	var item_list []models.TodoItem
	config.DB.Find((&item_list))
	c.JSON(200,  item_list)
}

func CreateNewItem(c *gin.Context) {
	var item models.TodoItem

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(400, gin.H{"Can't create new item": err.Error()})
		return
	}

	config.DB.Create(&item)
	c.JSON(201, item)
}