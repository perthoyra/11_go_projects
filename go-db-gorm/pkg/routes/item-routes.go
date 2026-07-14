package routes

import (
	"fmt"
	"go-db-gorm/pkg/controllers"

	"github.com/gin-gonic/gin"
)

var RegisterItemRoutes = func(c *gin.Engine) {
	// Simple group: v1
	// Should result in routes looking like '../v1/item/'
	// Add more groups below this if needed, ie v2, v3 etc
	{
		v1 := c.Group("/v1")
		v1.GET("/item", controllers.GetAllTodoItems)
		v1.GET("/item/:itemId", controllers.GetTodoItemById)
		v1.POST("/item", controllers.CreateItem)
		v1.PUT("/item/:itemId", controllers.UpdateItem)
		v1.DELETE("/item/:itemId", controllers.DeleteItem)
	}
}

func init() {
	fmt.Println("Initializing routes...")
}
