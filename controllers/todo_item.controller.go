package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/QuocB-HC/REST-API-Golang/config"
	"github.com/QuocB-HC/REST-API-Golang/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllItem(c *gin.Context) {
	var item_list []models.TodoItem

	if err := config.DB.Find((&item_list)).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to get todo list",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"item list": item_list,
	})
}

func GetByID(c *gin.Context) {
	var item models.TodoItem

	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
		return
	}

	if err := config.DB.Where("id = ?", id).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find todo with id: " + fmt.Sprint(id),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"item": item,
	})
}

func GetByTitle(c *gin.Context) {
	var item models.TodoItem

	title := c.Param("title")

	if err := config.DB.Where("title = ?", title).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Faild to find todo with title: " + title,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"item": item,
	})
}

func CreateNewItem(c *gin.Context) {
	var item models.TodoItemCreation

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
			"system error": err.Error(),
		})
		return
	}

	if strings.TrimSpace(item.Title) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Tittle cannot be empty",
		})
		return
	}

	if err := config.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create new todo",
			"system error": err.Error(),
		})
		return
	}

	GetAllItem(c)
}

func UpdateItem(c *gin.Context) {
	var item models.TodoItemUpdate
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
		return
	}

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
		})
		return
	}
	
	if err := config.DB.Where("id = ?", id).Updates(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update todo with id " + fmt.Sprint(id),
			"system_error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"item": true,
	})
}

func DeleteItem(c *gin.Context) {
	var item models.TodoItem
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
			"system error": err.Error(),
		})
	}

	if err := config.DB.Where("id = ?", id).Delete(item).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "Failed to delete todo with id " + fmt.Sprint(id),
			"system error": err.Error(),
		})
	}
}