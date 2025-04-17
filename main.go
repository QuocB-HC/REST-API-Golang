package main

import (
	// "net/http"
	// "time"

	"github.com/QuocB-HC/REST-API-Golang/config"
	// "github.com/QuocB-HC/REST-API-Golang/models"
	"github.com/QuocB-HC/REST-API-Golang/routes"
)

func main() {
	// now := time.Now().UTC()

	// item := models.TodoItem{
	// 	ID:          1,
	// 	Title:       "Learn Go",
	// 	Description: "Learn Go programming language",
	// 	Status:      "Doing",
	// 	CreatedAt:   &now,
	// 	UpdatedAt:   &now,
	// }

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": item,
	// 	})
	// })
	config.ConnectDB()

	r := routes.Router()

	r.Run(":3000")
}
