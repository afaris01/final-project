package main

import (
	"github.com/gin-gonic/gin"
	"github.com/takadev15/todo-apps/controllers"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/takadev15/todo-apps/docs"
)

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/to-do"
	todoRoutes := router.Group("/to-do")
	{
		todoRoutes.GET("/", controllers.AmbilAll)
		todoRoutes.GET("/:id", controllers.AmbilTodo)
		todoRoutes.POST("/", controllers.BuatTodo)
		todoRoutes.PUT("/:id", controllers.UbahTodo)
		todoRoutes.DELETE("/:id", controllers.HapusTodo)
	}
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":443")
}
