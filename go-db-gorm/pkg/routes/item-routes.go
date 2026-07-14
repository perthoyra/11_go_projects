package routes

import (
	"go-db-gorm/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var RegisterItemRoutes = func(c *gin.Engine) {

	c.GET("/item/", controllers.GetAllTodoItems)
	c.GET("/item/{itemId}", controllers.GetTodoItemById)
	c.POST("/item/", controllers.CreateItem)
	c.PUT("/item/{itemId}", controllers.UpdateItem)
	c.DELETE("/item/{itemId}", controllers.DeleteItem)
}
