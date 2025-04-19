package routes

import (
	"github.com/QuocB-HC/REST-API-Golang/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.GET("", controllers.GetAllItem)
			items.GET("/id/:id", controllers.GetByID)
			items.GET("/title/:title", controllers.GetByTitle)
			items.POST("", controllers.CreateNewItem)
			items.PATCH("/:id", controllers.UpdateItem)
			items.DELETE("/:id", controllers.DeleteItem)
		}
	}

	// r.GET("/google", func(c *gin.Context) {
	// 	c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	// })
	// r.POST("/test", func (c *gin.Context)  {
	// 	c.Redirect(http.StatusFound, "/google")
	// })

	return r
}
