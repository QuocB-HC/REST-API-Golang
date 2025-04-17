package routes

import (
	"github.com/QuocB-HC/REST-API-Golang/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/item_list", controllers.GetAllItem)
	r.POST("/item", controllers.CreateNewItem)

	return r
}